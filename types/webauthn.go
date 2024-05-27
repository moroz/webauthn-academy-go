package types

import (
	"database/sql/driver"
	"errors"
	"strings"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
)

type WebauthnCredential struct {
	ID                        int         `db:"id"`
	WebauthnID                []byte      `db:"webauthn_id"`
	UserID                    int         `db:"user_id"`
	DisplayName               string      `db:"display_name"`
	PublicKey                 []byte      `db:"public_key"`
	AttestationType           string      `db:"attestation_type"`
	Transport                 StringSlice `db:"transport"`
	UserPresent               bool        `db:"user_present"`
	UserVerified              bool        `db:"user_verified"`
	BackupEligible            bool        `db:"backup_eligible"`
	BackupState               bool        `db:"backup_state"`
	LastUsedAt                *time.Time  `db:"last_used_at"`
	LastUsedIP                *string     `db:"last_used_at"`
	LastUsedUserAgent         *string     `db:"last_used_user_agent"`
	AuthenticatorAAGUID       []byte      `db:"authenticator_aaguid"`
	AuthenticatorSignCount    uint32      `db:"authenticator_sign_count"`
	AuthenticatorCloneWarning bool        `db:"authenticator_clone_warning"`
	AuthenticatorAttachment   string      `db:"authenticator_attachment"`
	InsertedAt                time.Time   `db:"inserted_at"`
	UpdatedAt                 time.Time   `db:"updated_at"`
}

type StringSlice []string

// Value implements `driver.Valuer` interface for `StringSlice`
func (a StringSlice) Value() (driver.Value, error) {
	return "{" + strings.Join(a, ",") + "}", nil
}

// Scan implements `sql.Scan` interface for `StringSlice`
func (a *StringSlice) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	literal := strings.Trim(string(b), "{}")
	*a = StringSlice(strings.Split(literal, ","))
	return nil
}

func (wc *WebauthnCredential) ToCredential() webauthn.Credential {
	var transports []protocol.AuthenticatorTransport
	for _, value := range wc.Transport {
		transports = append(transports, protocol.AuthenticatorTransport(value))
	}

	return webauthn.Credential{
		ID: wc.WebauthnID,
		Authenticator: webauthn.Authenticator{
			AAGUID:       wc.AuthenticatorAAGUID,
			SignCount:    wc.AuthenticatorSignCount,
			CloneWarning: wc.AuthenticatorCloneWarning,
			Attachment:   protocol.AuthenticatorAttachment(wc.AuthenticatorAttachment),
		},
		PublicKey: wc.PublicKey,
		Flags: webauthn.CredentialFlags{
			UserPresent:    wc.UserPresent,
			UserVerified:   wc.UserVerified,
			BackupEligible: wc.BackupEligible,
			BackupState:    wc.BackupState,
		},
		AttestationType: wc.AttestationType,
		Transport:       transports,
	}
}
