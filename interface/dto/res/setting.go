package res

import (
	"IkezawaYuki/a-root-backend/domain/model"
	"time"
)

type Setting struct {
	ID                   int       `json:"id"`
	CustomerID           int       `json:"customer_id"`
	FacebookToken        string    `json:"facebook_token"`
	StartDate            time.Time `json:"start_date"`
	InstagramAccountID   string    `json:"instagram_account_id"`
	InstagramAccountName string    `json:"instagram_account_name"`
	DeleteHashFlag       bool      `json:"delete_hash_flag"`
}

func GetSetting(s *model.Setting) *Setting {
	return &Setting{
		ID:                   int(s.ID),
		CustomerID:           int(s.CustomerID),
		FacebookToken:        s.FacebookToken,
		StartDate:            s.StartDate,
		InstagramAccountID:   s.InstagramAccountID,
		InstagramAccountName: s.InstagramAccountName,
		DeleteHashFlag:       s.DeleteHashFlag.ToBool(),
	}
}
