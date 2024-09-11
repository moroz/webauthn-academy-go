package layout

import (
	"context"

	"github.com/moroz/webauthn-academy-go/config"
	"github.com/moroz/webauthn-academy-go/types"
)

func fetchFlash(ctx context.Context) []types.FlashMessage {
	if values, ok := ctx.Value(config.FlashContextKey).([]types.FlashMessage); ok {
		return values
	}
	return []types.FlashMessage{}
}
