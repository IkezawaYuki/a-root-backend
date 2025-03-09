package entity

type PaymentStatus int

const (
	PaymentStatusUnknown PaymentStatus = -1
	PaymentStatusNone    PaymentStatus = 0
	PaymentStatusPending PaymentStatus = 1
	PaymentStatusSuccess PaymentStatus = 2
	PaymentStatusFailed  PaymentStatus = 3
)

func ConvertToPaymentStatus(s *string) PaymentStatus {
	if s == nil {
		return PaymentStatusUnknown
	}
	switch *s {
	case "none":
		return PaymentStatusNone
	case "pending":
		return PaymentStatusPending
	case "success":
		return PaymentStatusSuccess
	case "failed":
		return PaymentStatusFailed
	default:
		return PaymentStatusUnknown
	}
}

func (p *PaymentStatus) ToString() string {
	switch *p {
	case PaymentStatusNone:
		return "none"
	case PaymentStatusPending:
		return "pending"
	case PaymentStatusSuccess:
		return "success"
	case PaymentStatusFailed:
		return "failed"
	default:
		return ""
	}
}
