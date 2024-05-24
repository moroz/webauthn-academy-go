package types

import "time"

type UserTokenContext string

const (
	UserTokenContext_Session UserTokenContext = "session"
)

type UserToken struct {
	ID         int              `db:"id"`
	UserId     int              `db:"user_id"`
	Token      []byte           `db:"token"`
	Context    UserTokenContext `db:"context"`
	InsertedAt time.Time        `db:"inserted_at"`
}
