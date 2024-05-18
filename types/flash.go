package types

type FlashMessageSeverity string

const (
	FlashMessageSeverity_Error   FlashMessageSeverity = "error"
	FlashMessageSeverity_Info    FlashMessageSeverity = "info"
	FlashMessageSeverity_Success FlashMessageSeverity = "success"
)

type FlashMessage struct {
	Severity FlashMessageSeverity
	Content  string
}
