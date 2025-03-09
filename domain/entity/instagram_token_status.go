package entity

type InstagramTokenStatus int

const (
	InstagramTokenStatusUnknown InstagramTokenStatus = -1
	InstagramTokenStatusYet     InstagramTokenStatus = 0
	InstagramTokenStatusActive  InstagramTokenStatus = 1
	InstagramTokenStatusExpired InstagramTokenStatus = 2
	InstagramTokenStatusPaused  InstagramTokenStatus = 3
)

func (t InstagramTokenStatus) ToString() string {
	switch t {
	case InstagramTokenStatusYet:
		return "yet"
	case InstagramTokenStatusActive:
		return "active"
	case InstagramTokenStatusExpired:
		return "expired"
	case InstagramTokenStatusPaused:
		return "paused"
	default:
		return "unknown"
	}
}

func ConvertToInstagramTokenStatus(status *string) InstagramTokenStatus {
	if status == nil {
		return InstagramTokenStatusUnknown
	}
	switch *status {
	case "yet":
		return InstagramTokenStatusYet
	case "active":
		return InstagramTokenStatusActive
	case "expired":
		return InstagramTokenStatusExpired
	case "paused":
		return InstagramTokenStatusPaused
	default:
		return InstagramTokenStatusUnknown
	}
}
