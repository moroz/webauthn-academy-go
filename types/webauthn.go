package types

import (
	"database/sql/driver"
	"errors"
	"strings"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/moroz/webauthn-academy-go/db/queries"
)

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

func BuildFrameworkCredential(wc queries.WebauthnCredential) webauthn.Credential {
	var transports []protocol.AuthenticatorTransport
	for _, value := range wc.Transport {
		transports = append(transports, protocol.AuthenticatorTransport(value))
	}

	attestationType := ""
	if wc.AttestationType != nil {
		attestationType = *wc.AttestationType
	}

	attachment := protocol.AuthenticatorAttachment("")
	if wc.AuthenticatorAttachment != nil {
		attachment = protocol.AuthenticatorAttachment(*wc.AuthenticatorAttachment)
	}

	return webauthn.Credential{
		ID: wc.WebauthnID,
		Authenticator: webauthn.Authenticator{
			AAGUID:       wc.AuthenticatorAAGUID,
			SignCount:    uint32(wc.AuthenticatorSignCount),
			CloneWarning: wc.AuthenticatorCloneWarning,
			Attachment:   attachment,
		},
		PublicKey: wc.PublicKey,
		Flags: webauthn.CredentialFlags{
			UserPresent:    wc.UserPresent,
			UserVerified:   wc.UserVerified,
			BackupEligible: wc.BackupEligible,
			BackupState:    wc.BackupState,
		},
		AttestationType: attestationType,
		Transport:       transports,
	}
}
