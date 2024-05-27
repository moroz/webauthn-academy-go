package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/moroz/webauthn-academy-go/types"
)

type WebauthnCredentialStore struct {
	db *sqlx.DB
}

func NewWebauthnCredentialStore(db *sqlx.DB) WebauthnCredentialStore {
	return WebauthnCredentialStore{db}
}

var listUserWebauthnCredentialsQuery = `select * from webauthn_credentials where user_id = $1`

func (s *WebauthnCredentialStore) ListUserWebauthnCredentials(userId int) ([]types.WebauthnCredential, error) {
	var result []types.WebauthnCredential
	err := s.db.Select(&result, listUserWebauthnCredentialsQuery, userId)
	return result, err
}

var insertCredentialQuery = `insert into webauthn_credentials (webauthn_id, user_id, public_key, attestation_type, transport, user_present, user_verified, backup_eligible, backup_state, display_name, authenticator_aaguid, authenticator_sign_count, authenticator_clone_warning, authenticator_attachment) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) returning *`

func (s *WebauthnCredentialStore) InsertWebauthnCredential(wc *types.WebauthnCredential) (*types.WebauthnCredential, error) {
	var result types.WebauthnCredential
	err := s.db.Get(&result, insertCredentialQuery, wc.WebauthnID, wc.UserID, wc.PublicKey, wc.AttestationType, wc.Transport, wc.UserVerified, wc.BackupEligible, wc.BackupState, wc.DisplayName, wc.AuthenticatorAAGUID, wc.AuthenticatorSignCount, wc.AuthenticatorCloneWarning, wc.AuthenticatorAttachment)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
