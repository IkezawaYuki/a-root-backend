package entity

type PaymentType int

const (
	PaymentTypeUnknown PaymentType = -1
	PaymentTypeNone    PaymentType = 0
	PaymentTypeStripe  PaymentType = 1
)

func ConvertToPaymentType(s *string) PaymentType {
	if s == nil {
		return PaymentTypeUnknown
	}
	switch *s {
	case "none":
		return PaymentTypeNone
	case "stripe":
		return PaymentTypeStripe
	default:
		return PaymentTypeUnknown
	}
}
