package api

import (
	"path/filepath"

	"github.com/Sirupsen/logrus"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"crypto/rand"
	"github.com/docker/notary"
	"github.com/docker/notary/client"
	"github.com/docker/notary/client/changelist"
	"github.com/docker/notary/cryptoservice"
	"github.com/docker/notary/storage"
	"github.com/docker/notary/trustmanager"
	"github.com/docker/notary/trustpinning"
	"github.com/docker/notary/tuf/data"
	"github.com/docker/notary/utils"
)

// NewServer creates a new instance of a Client API server with a configured
// upstream Notary Server.
func NewServer(upstream string, upstreamCAPath string, grpcSrv *grpc.Server) (*grpc.Server, error) {
	srv := &Server{
		upstream:       upstream,
		upstreamCAPath: upstreamCAPath,
	}
	RegisterNotaryServer(grpcSrv, srv)
	return grpcSrv, nil
}

type Server struct {
	upstream       string
	upstreamCAPath string
}

func (srv *Server) Initialize(ctx context.Context, initMessage *InitMessage) (*BasicResponse, error) {
	r, err := srv.initRepo(data.GUN(initMessage.Gun))
	if err != nil {
		return nil, err
	}

	roles := make([]data.RoleName, len(initMessage.ServerManagedRoles.Roles))
	for index, role := range initMessage.ServerManagedRoles.Roles {
		roles[index] = data.RoleName(role)
	}

	err = r.Initialize(initMessage.RootKeyIDs, roles...)
	if err != nil {
		return nil, err
	}

	return &BasicResponse{
		Success: true,
	}, nil
}

func (srv *Server) Publish(ctx context.Context, gun *GunMessage) (*BasicResponse, error) {
	r, err := srv.initRepo(data.GUN(gun.Gun))
	if err != nil {
		return nil, err
	}

	err = r.Publish()
	if err != nil {
		return nil, err
	}

	return &BasicResponse{
		Success: true,
	}, nil
}

func (srv *Server) AddTarget(ctx context.Context, t *Target) (*BasicResponse, error) {
	r, err := srv.initRepo(data.GUN(t.GetGun()))
	if err != nil {
		return nil, err
	}
	if err := r.AddTarget(
		&client.Target{
			Name:   t.GetName(),
			Hashes: data.Hashes(t.Hashes),
			Length: t.Length,
		},
	); err != nil {
		return nil, err
	}
	if err := publishRepo(r); err != nil {
		return nil, err
	}

	return &BasicResponse{
		Success: true,
	}, nil
}

func (srv *Server) RemoveTarget(ctx context.Context, t *Target) (*BasicResponse, error) {
	r, err := srv.initRepo(data.GUN(t.GetGun()))
	if err != nil {
		return nil, err
	}
	if err := r.RemoveTarget(
		t.GetName(), "targets",
	); err != nil {
		return nil, err
	}
	if err := publishRepo(r); err != nil {
		return nil, err
	}

	return &BasicResponse{
		Success: true,
	}, nil
}

func (srv *Server) ListTargets(ctx context.Context, message *RoleNameListMessage) (*TargetWithRoleNameListResponse, error) {
	r, err := srv.initRepo(data.GUN(message.Gun))
	if err != nil {
		return nil, err
	}

	roles := make([]data.RoleName, len(message.Roles))
	for index, role := range message.Roles {
		roles[index] = data.RoleName(role)
	}

	targets, err := r.ListTargets(roles...)
	if err != nil {
		return nil, err
	}

	resTargets := make([]*TargetWithRole, len(targets))
	for index, target := range targets {
		resTargets[index] = &TargetWithRole{
			Target: &Target{
				Gun:    message.Gun,
				Name:   target.Name,
				Length: target.Length,
				Hashes: target.Hashes,
			},
			Role: target.Role.String(),
		}
	}

	return &TargetWithRoleNameListResponse{
		TargetWithRoleNameList: &TargetWithRoleNameList{
			Targets: resTargets,
		},
		Success: true,
	}, nil
}

// GetTargetByName returns a target by the given name.
func (srv *Server) GetTargetByName(ctx context.Context, message *TargetByNameAction) (*TargetWithRoleResponse, error) {
	r, err := srv.initRepo(data.GUN(message.Gun))
	if err != nil {
		return nil, err
	}

	roles := make([]data.RoleName, len(message.Roles.Roles))
	for index, role := range message.Roles.Roles {
		roles[index] = data.RoleName(role)
	}

	target, err := r.GetTargetByName(message.Name, roles...)
	if err != nil {
		return nil, err
	}

	return &TargetWithRoleResponse{
		TargetWithRole: &TargetWithRole{
			Target: &Target{
				Gun:    message.Gun,
				Name:   target.Name,
				Length: target.Length,
				Hashes: target.Hashes,
			},
			Role: target.Role.String(),
		},
		Success: true,
	}, nil
}

// GetAllTargetMetadataByName
func (srv *Server) GetAllTargetMetadataByName(ctx context.Context, message *TargetNameMessage) (*TargetSignedListResponse, error) {
	r, err := srv.initRepo(data.GUN(message.Gun))
	if err != nil {
		return nil, err
	}

	targets, err := r.GetAllTargetMetadataByName(message.Name)
	if err != nil {
		return nil, err
	}

	resTargets := make([]*TargetSigned, len(targets))
	for indexTarget, target := range targets {

		resSignatures := make([]*Signature, len(target.Signatures))
		for indexSig, signature := range target.Signatures {
			resSignatures[indexSig] = &Signature{
				KeyID:  signature.KeyID,
				Method: signature.Method.String(),
			}
		}

		resKeys := make(map[string]*PublicKey, len(target.Role.Keys))
		for keyName, keyPubkey := range target.Role.Keys {
			resKeys[keyName] = &PublicKey{
				Id:        keyPubkey.ID(),
				Algorithm: keyPubkey.Algorithm(),
				Public:    keyPubkey.Public(),
			}
		}

		resTargets[indexTarget] = &TargetSigned{
			Role: &DelegationRole{
				Keys:      resKeys,
				Name:      target.Role.Name.String(),
				Threshold: int32(target.Role.Threshold), // FIXME
				Paths:     target.Role.Paths,
			},
			Target: &Target{
				Gun:    message.Gun,
				Name:   target.Target.Name,
				Length: target.Target.Length,
				Hashes: target.Target.Hashes,
			},
			Signatures: resSignatures,
		}
	}

	return &TargetSignedListResponse{
		TargetSignedList: &TargetSignedList{
			Targets: resTargets,
		},
		Success: true,
	}, nil
}

// GetChangelist returns the list of the repository's unpublished changes
func (srv *Server) GetChangelist(ctx context.Context, message *GunMessage) (*ChangeListResponse, error) {
	r, err := srv.initRepo(data.GUN(message.Gun))
	if err != nil {
		return nil, err
	}

	changelist, err := r.GetChangelist()
	if err != nil {
		return nil, err
	}

	resChangelist := make([]*Change, len(changelist.List()))
	for index, change := range changelist.List() {
		resChangelist[index] = &Change{
			Action:  change.Action(),
			Scope:   change.Scope().String(),
			Type:    change.Type(),
			Path:    change.Path(),
			Content: change.Content(),
		}
	}

	return &ChangeListResponse{
		Changelist: &ChangeList{
			Changes: resChangelist,
		},
		Success: true,
	}, nil
}

func (srv *Server) ListRoles(ctx context.Context, message *GunMessage) (*RoleWithSignaturesListResponse, error) {
	r, err := srv.initRepo(data.GUN(message.Gun))
	if err != nil {
		return nil, err
	}

	roles, err := r.ListRoles()
	if err != nil {
		return nil, err
	}

	resRoles := make([]*RoleWithSignatures, len(roles))
	for index, role := range roles {

		resSignatures := make([]*Signature, len(role.Signatures))
		for indexSig, signature := range role.Signatures {
			resSignatures[indexSig] = &Signature{
				KeyID:  signature.KeyID,
				Method: signature.Method.String(),
			}
		}

		resRoles[index] = &RoleWithSignatures{
			Signatures: resSignatures,
			Role: &Role{
				RootRole: &RootRole{
					KeyIDs:    role.KeyIDs,
					Threshold: int32(role.Threshold), // FIXME
				},
				Name:  role.Name.String(),
				Paths: role.Paths,
			},
		}
	}

	return &RoleWithSignaturesListResponse{
		RoleWithSignaturesList: &RoleWithSignaturesList{
			RoleWithSignatures: resRoles,
		},
		Success: true,
	}, nil
}

func (srv *Server) GetDelegationRoles(ctx context.Context, message *GunMessage) (*RoleListResponse, error) {
	r, err := srv.initRepo(data.GUN(message.Gun))
	if err != nil {
		return nil, err
	}

	roles, err := r.GetDelegationRoles()
	if err != nil {
		return nil, err
	}

	resRoles := make([]*Role, len(roles))
	for index, role := range roles {
		resRoles[index] = &Role{
			RootRole: &RootRole{
				KeyIDs:    role.KeyIDs,
				Threshold: int32(role.Threshold), // FIXME
			},
			Name:  role.Name.String(),
			Paths: role.Paths,
		}
	}

	return &RoleListResponse{
		RoleList: &RoleList{
			Roles: resRoles,
		},
		Success: true,
	}, nil
}

func (srv *Server) AddDelegation(ctx context.Context, message *AddDelegationMessage) (*BasicResponse, error) {
	r, err := srv.initRepo(data.GUN(message.Gun))
	if err != nil {
		return nil, err
	}

	delegationKeys := make([]data.PublicKey, len(message.DelegationKeys))
	for index, key := range message.DelegationKeys {
		delegationKeys[index] = data.NewPublicKey(key.Algorithm, key.Public)
	}

	err = r.AddDelegation(data.RoleName(message.Name), delegationKeys, message.Paths)
	if err != nil {
		return nil, err
	}

	return &BasicResponse{
		Success: true,
	}, nil
}

func (srv *Server) AddDelegationRoleAndKeys(ctx context.Context, message *AddDelegationRoleAndKeysMessage) (*BasicResponse, error) {
	r, err := srv.initRepo(data.GUN(message.Gun))
	if err != nil {
		return nil, err
	}

	delegationKeys := make([]data.PublicKey, len(message.DelegationKeys))
	for index, key := range message.DelegationKeys {
		delegationKeys[index] = data.NewPublicKey(key.Algorithm, key.Public)
	}

	err = r.AddDelegationRoleAndKeys(data.RoleName(message.Name), delegationKeys)
	if err != nil {
		return nil, err
	}

	return &BasicResponse{
		Success: true,
	}, nil
}

func (srv *Server) AddDelegationPaths(ctx context.Context, message *AddDelegationPathsMessage) (*BasicResponse, error) {
	r, err := srv.initRepo(data.GUN(message.Gun))
	if err != nil {
		return nil, err
	}

	err = r.AddDelegationPaths(data.RoleName(message.Name), message.Paths)
	if err != nil {
		return nil, err
	}

	return &BasicResponse{
		Success: true,
	}, nil
}

func (srv *Server) RemoveDelegationKeysAndPaths(ctx context.Context, message *RemoveDelegationKeysAndPathsMessage) (*BasicResponse, error) {
	r, err := srv.initRepo(data.GUN(message.Gun))
	if err != nil {
		return nil, err
	}

	err = r.RemoveDelegationKeysAndPaths(data.RoleName(message.Name), message.KeyIDs, message.Paths)
	if err != nil {
		return nil, err
	}

	return &BasicResponse{
		Success: true,
	}, nil
}

func (srv *Server) RemoveDelegationRole(ctx context.Context, message *RemoveDelegationRoleMessage) (*BasicResponse, error) {
	r, err := srv.initRepo(data.GUN(message.Gun))
	if err != nil {
		return nil, err
	}

	err = r.RemoveDelegationRole(data.RoleName(message.Name))
	if err != nil {
		return nil, err
	}

	return &BasicResponse{
		Success: true,
	}, nil
}

func (srv *Server) RemoveDelegationPaths(ctx context.Context, message *RemoveDelegationPathsMessage) (*BasicResponse, error) {
	r, err := srv.initRepo(data.GUN(message.Gun))
	if err != nil {
		return nil, err
	}

	err = r.RemoveDelegationPaths(data.RoleName(message.Name), message.Paths)
	if err != nil {
		return nil, err
	}

	return &BasicResponse{
		Success: true,
	}, nil
}

func (srv *Server) RemoveDelegationKeys(ctx context.Context, message *RemoveDelegationKeysMessage) (*BasicResponse, error) {
	r, err := srv.initRepo(data.GUN(message.Gun))
	if err != nil {
		return nil, err
	}

	err = r.RemoveDelegationKeys(data.RoleName(message.Name), message.KeyIDs)
	if err != nil {
		return nil, err
	}

	return &BasicResponse{
		Success: true,
	}, nil
}

func (srv *Server) ClearDelegationPaths(ctx context.Context, message *RoleNameMessage) (*BasicResponse, error) {
	r, err := srv.initRepo(data.GUN(message.Gun))
	if err != nil {
		return nil, err
	}

	err = r.ClearDelegationPaths(data.RoleName(message.Role))
	if err != nil {
		return nil, err
	}

	return &BasicResponse{
		Success: true,
	}, nil
}

func (srv *Server) Witness(ctx context.Context, message *RoleNameListMessage) (*RoleNameListResponse, error) {
	r, err := srv.initRepo(data.GUN(message.Gun))
	if err != nil {
		return nil, err
	}

	roles := make([]data.RoleName, len(message.Roles))
	for index, role := range message.Roles {
		roles[index] = data.RoleName(role)
	}

	currRoles, err := r.Witness(roles...)
	if err != nil {
		return nil, err
	}

	resRoles := make([]string, len(currRoles))
	for index, role := range currRoles {
		resRoles[index] = role.String()
	}

	return &RoleNameListResponse{
		RoleNameList: &RoleNameList{
			Roles: resRoles,
		},
		Success: true,
	}, nil
}

func (srv *Server) RotateKey(ctx context.Context, message *RotateKeyMessage) (*BasicResponse, error) {
	r, err := srv.initRepo(data.GUN(message.Gun))
	if err != nil {
		return nil, err
	}

	err = r.RotateKey(data.RoleName(message.Role), message.ServerManagesKey, message.KeyList)
	if err != nil {
		return nil, err
	}

	return &BasicResponse{
		Success: true,
	}, nil
}

// CryptoService implementation
func (srv *Server) CryptoService(ctx context.Context, message *GunMessage) (*CryptoServiceMessage, error) {
	return nil, ErrNotImplemented
}

func (srv *Server) CryptoServiceCreate(ctx context.Context, message *CryptoServiceCreateMessage) (*PublicKeyResponse, error) {
	r, err := srv.initRepo(data.GUN(message.Gun))
	if err != nil {
		return nil, err
	}

	cs := r.CryptoService()

	pubkey, err := cs.Create(data.RoleName(message.RoleName), data.GUN(message.Gun), message.Algorithm)
	if err != nil {
		return nil, err
	}

	return &PublicKeyResponse{
		Pubkey: &PublicKey{
			Id:        pubkey.ID(),
			Algorithm: pubkey.Algorithm(),
			Public:    pubkey.Public(),
		},
		Success: true,
	}, nil
}

func (srv *Server) CryptoServicePrivateKeySign(ctx context.Context, message *CryptoServicePrivateKeySignMessage) (*SignatureResponse, error) {
	pubkey := data.NewPublicKey(message.Pubkey.Algorithm, message.Pubkey.Public)
	privkey, err := data.NewPrivateKey(pubkey, message.Privkey)
	if err != nil {
		return nil, err
	}

	sig, err := privkey.Sign(rand.Reader, message.Digest, nil)
	if err != nil {
		return nil, err
	}

	return &SignatureResponse{
		Signature: sig,
		Success:   true,
	}, nil
}

func (srv *Server) CryptoServiceAddKey(ctx context.Context, message *CryptoServiceAddKeyMessage) (*BasicResponse, error) {
	r, err := srv.initRepo(data.GUN(message.Gun))
	if err != nil {
		return nil, err
	}

	cs := r.CryptoService()

	pubKey := data.NewPublicKey(message.Key.Pubkey.Algorithm, message.Key.Pubkey.Public)

	privKey, err := data.NewPrivateKey(pubKey, message.Key.Privkey)
	if err != nil {
		return nil, err
	}

	err = cs.AddKey(data.RoleName(message.RoleName), data.GUN(message.Gun), privKey)
	if err != nil {
		return nil, err
	}

	return &BasicResponse{
		Success: true,
	}, nil
}

func (srv *Server) CryptoServiceGetKey(ctx context.Context, message *KeyIDMessage) (*PublicKeyResponse, error) {
	r, err := srv.initRepo(data.GUN(message.Gun))
	if err != nil {
		return nil, err
	}

	cs := r.CryptoService()

	pubkey := cs.GetKey(message.KeyID)

	return &PublicKeyResponse{
		Pubkey: &PublicKey{
			Id:        pubkey.ID(),
			Algorithm: pubkey.Algorithm(),
			Public:    pubkey.Public(),
		},
		Success: true,
	}, nil
}

func (srv *Server) CryptoServiceGetPrivateKey(ctx context.Context, message *KeyIDMessage) (*PrivateKeyResponse, error) {
	r, err := srv.initRepo(data.GUN(message.Gun))
	if err != nil {
		return nil, err
	}

	cs := r.CryptoService()

	privkey, rolename, err := cs.GetPrivateKey(message.KeyID)
	if err != nil {
		return nil, err
	}

	resPrivkey := &PrivateKey{
		Pubkey: &PublicKey{
			Id:        privkey.ID(),
			Algorithm: privkey.Algorithm(),
			Public:    privkey.Public(),
		},
		Privkey: privkey.Private(),
		CryptoSigner: &Signer{
			Pubkey: &PublicKey{
				Id:        data.PublicKey(privkey.CryptoSigner().Public()).ID(),
				Algorithm: data.PublicKey(privkey.CryptoSigner().Public()).Algorithm(),
				Public:    data.PublicKey(privkey.CryptoSigner().Public()).Public(),
			},
		},
		SigAlgorithm: privkey.SignatureAlgorithm().String(),
	}

	return &PrivateKeyResponse{
		Role:    rolename.String(),
		Gun:     message.Gun,
		Privkey: resPrivkey,
		Success: true,
	}, nil
}

func (srv *Server) CryptoServiceRemoveKey(ctx context.Context, message *KeyIDMessage) (*BasicResponse, error) {
	r, err := srv.initRepo(data.GUN(message.Gun))
	if err != nil {
		return nil, err
	}

	cs := r.CryptoService()

	err = cs.RemoveKey(message.KeyID)
	if err != nil {
		return nil, err
	}

	return &BasicResponse{
		Success: true,
	}, nil
}

func (srv *Server) CryptoServiceListKeys(ctx context.Context, message *RoleNameMessage) (*KeyIDsListResponse, error) {
	r, err := srv.initRepo(data.GUN(message.Gun))
	if err != nil {
		return nil, err
	}

	cs := r.CryptoService()

	keys := cs.ListKeys(data.RoleName(message.Role))

	return &KeyIDsListResponse{
		KeyIDs:  keys,
		Success: true,
	}, nil
}

func (srv *Server) CryptoServiceListAllKeys(ctx context.Context, message *GunMessage) (*SigningKeyIDsToRolesResponse, error) {
	r, err := srv.initRepo(data.GUN(message.Gun))
	if err != nil {
		return nil, err
	}

	cs := r.CryptoService()

	keys := cs.ListAllKeys()

	resKeys := make(map[string]string, len(keys))
	for keyID, role := range keys {
		resKeys[keyID] = role.String()
	}

	return &SigningKeyIDsToRolesResponse{
		KeyIDs:  resKeys,
		Success: true,
	}, nil
}

func (srv *Server) SetLegacyVersions(ctx context.Context, message *VersionMessage) (*google_protobuf.Empty, error) {
	return nil, ErrNotImplemented
}

func publishRepo(r *client.NotaryRepository) error {
	if err := r.Publish(); err != nil {
		if _, ok := err.(client.ErrRepoNotInitialized); !ok {
			return err
		}
		if err := initializeRepo(r); err != nil {
			return err
		}
		return r.Publish()
	}
	return nil
}

func initializeRepo(r *client.NotaryRepository) error {
	rootKeyList := r.CryptoService().ListKeys(data.CanonicalRootRole)
	var rootKeyID string
	if len(rootKeyList) < 1 {
		rootPublicKey, err := r.CryptoService().Create(data.CanonicalRootRole, "", data.ECDSAKey)
		if err != nil {
			return err
		}
		rootKeyID = rootPublicKey.ID()
	} else {
		// Chooses the first root key available, which is initialization specific
		// but should return the HW one first.
		rootKeyID = rootKeyList[0]
	}
	return r.Initialize([]string{rootKeyID})
}

func (srv *Server) initRepo(gun data.GUN) (*client.NotaryRepository, error) {
	logrus.Errorf("initializing with upstream ca file %s", srv.upstreamCAPath)
	baseDir := "var/lib/clientapi"
	rt, err := utils.GetReadOnlyAuthTransport(
		srv.upstream,
		[]string{gun.String()},
		"",
		"",
		srv.upstreamCAPath,
	)
	if err != nil {
		return nil, err
	}

	keyStore, err := trustmanager.NewKeyFileStore(filepath.Join(baseDir, notary.PrivDir), retriever)
	if err != nil {
		return nil, err
	}

	cryptoService := cryptoservice.NewCryptoService(keyStore)

	remoteStore, err := storage.NewHTTPStore(
		srv.upstream+"/v2/"+gun.String()+"/_trust/tuf/",
		"",
		"json",
		"key",
		rt,
	)

	return client.NewNotaryRepository(
		baseDir,
		gun,
		srv.upstream,
		remoteStore, // remote store
		storage.NewMemoryStore(nil),
		trustpinning.TrustPinConfig{},
		cryptoService,
		changelist.NewMemChangelist(),
	)
}

func retriever(keyName, alias string, createNew bool, attempts int) (string, bool, error) {
	return "password", false, nil
}

func DefaultPermissions() map[string][]string {
	return map[string][]string{
		"/api.Notary/AddTarget":    {"push", "pull"},
		"/api.Notary/RemoveTarget": {"push", "pull"},
	}
}
