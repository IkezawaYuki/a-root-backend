package model

import (
	"IkezawaYuki/a-root-backend/domain/entity"
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	gorm.Model
	Name                 string                      `gorm:"column:name"`
	Email                string                      `gorm:"column:email"`
	Password             string                      `gorm:"column:password"`
	WordpressUrl         string                      `gorm:"column:wordpress_url"`
	StartDate            *time.Time                  `gorm:"column:start_date"`
	FacebookToken        *string                     `gorm:"column:facebook_token"`
	InstagramTokenStatus entity.InstagramTokenStatus `gorm:"column:instagram_token_status"`
	InstagramAccountID   *string                     `gorm:"column:instagram_account_id"`
	InstagramAccountName *string                     `gorm:"column:instagram_account_name"`
	SubscriptionID       *string                     `gorm:"column:stripe_subscription_id"`
	StripeCustomerID     *string                     `gorm:"column:stripe_customer_id"`
	PaymentType          entity.PaymentType          `gorm:"column:payment_type"`
	PaymentStatus        entity.PaymentStatus        `gorm:"column:payment_status"`
	DeleteHashFlag       entity.DeleteHashFlag       `gorm:"column:delete_hash_flag"`
	DashboardStatus      entity.DashboardStatus      `gorm:"column:dashboard_status"`
}

func (c *Customer) IsLinked() bool {
	return c.FacebookToken != nil && c.InstagramTokenStatus == entity.InstagramTokenStatusActive
}
