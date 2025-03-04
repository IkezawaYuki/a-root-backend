package model

import (
	"IkezawaYuki/a-root-backend/domain/entity"
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	gorm.Model
	Name                 string                  `gorm:"column:name"`
	Email                string                  `gorm:"column:email"`
	Password             string                  `gorm:"column:password"`
	WordpressUrl         string                  `gorm:"column:wordpress_url"`
	FacebookToken        *string                 `gorm:"column:facebook_token"`
	StartDate            *time.Time              `gorm:"column:start_date"`
	InstagramTokenStatus *InstagramTokenStatus   `gorm:"column:instagram_token_status"`
	InstagramAccountID   *string                 `gorm:"column:instagram_account_id"`
	InstagramAccountName *string                 `gorm:"column:instagram_account_name"`
	SubscriptionID       *string                 `gorm:"column:stripe_subscription_id"`
	StripeCustomerID     *string                 `gorm:"column:stripe_customer_id"`
	PaymentType          *entity.PaymentType     `gorm:"column:payment_type"`
	PaymentStatus        *PaymentStatus          `gorm:"column:payment_status"`
	DeleteHashFlag       *DeleteHashFlag         `gorm:"column:delete_hash_flag"`
	DashboardStatus      *entity.DashboardStatus `gorm:"column:dashboard_status"`
}

func (c *Customer) IsLinked() bool {
	return c.InstagramTokenStatus != nil && *c.InstagramTokenStatus == InstagramTokenStatusActive
}

type MetaTokenStatus int

const (
	MetaTokenStatusInactive MetaTokenStatus = iota
	MetaTokenStatusActive
)

type DeleteHashFlag bool

const (
	DeleteHashFlagFalse DeleteHashFlag = false
	DeleteHashFlagTrue  DeleteHashFlag = true
)

func (f *DeleteHashFlag) ToBool() bool {
	return bool(*f)
}

type InstagramTokenStatus int

const (
	InstagramTokenStatusUnknown InstagramTokenStatus = -1
	InstagramTokenStatusYet     InstagramTokenStatus = 0
	InstagramTokenStatusActive  InstagramTokenStatus = 1
	InstagramTokenStatusExpired InstagramTokenStatus = 2
)

func (t InstagramTokenStatus) ToString() string {
	switch t {
	case InstagramTokenStatusYet:
		return "yet"
	case InstagramTokenStatusActive:
		return "active"
	case InstagramTokenStatusExpired:
		return "expired"
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
	default:
		return InstagramTokenStatusUnknown
	}
}

type PaymentStatus int

const (
	PaymentStatusUnknown PaymentStatus = -1
	PaymentStatusPending PaymentStatus = 0
	PaymentStatusSuccess PaymentStatus = 1
	PaymentStatusFailed  PaymentStatus = 2
)

func ConvertToPaymentStatus(s *string) PaymentStatus {
	if s == nil {
		return PaymentStatusUnknown
	}
	switch *s {
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
