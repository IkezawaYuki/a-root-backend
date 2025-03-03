package filter

import "gorm.io/gorm"

type PaymentFilter struct {
	ID *string
}

func (f *PaymentFilter) GenerateMods(db *gorm.DB) *gorm.DB {
	if f.ID != nil {
		db = db.Where("id = ?", *f.ID)
	}
	return db
}
