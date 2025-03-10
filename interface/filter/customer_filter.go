package filter

import (
	"IkezawaYuki/a-root-backend/domain/entity"
	"gorm.io/gorm"
)

type CustomerFilter struct {
	Zero *bool
	All  *bool

	ID                   *int
	Email                *string
	PartialName          *string
	PartialWordpressUrl  *string
	PaymentType          *entity.PaymentType
	PaymentStatus        *entity.PaymentStatus
	InstagramTokenStatus *entity.InstagramTokenStatus

	Limit  *int
	Offset *int
}

func (c *CustomerFilter) GenerateMods(db *gorm.DB) *gorm.DB {
	if c.ID != nil {
		db = db.Where("id = ?", *c.ID)
	}
	if c.Email != nil {
		db = db.Where("email = ?", *c.Email)
	}
	if c.PartialName != nil {
		db = db.Where("partial_name like '%?%'", *c.PartialName)
	}
	if c.InstagramTokenStatus != nil {
		db = db.Where("instagram_token_status = ?", *c.InstagramTokenStatus)
	}
	if c.Limit != nil {
		db = db.Limit(*c.Limit)
		if c.Offset != nil {
			db = db.Offset(*c.Offset)
		}
	}
	return db
}
