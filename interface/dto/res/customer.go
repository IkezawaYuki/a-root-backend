package res

import (
	"IkezawaYuki/a-root-backend/domain/model"
	"time"
)

type Customer struct {
	ID                   int        `json:"id"`
	Name                 string     `json:"name"`
	Email                string     `json:"email"`
	WordpressURL         string     `json:"wordpress_url"`
	FacebookToken        *string    `json:"facebook_token"`
	StartDate            *time.Time `json:"start_date"`
	InstagramAccountID   *string    `json:"instagram_account_id"`
	InstagramAccountName *string    `json:"instagram_account_name"`
	DeleteHashFlag       *bool      `json:"delete_hash_flag"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
}

type Customers struct {
	Customers []*Customer `json:"customers"`
	Paginate
}

func GetCustomer(c *model.Customer) *Customer {
	customer := Customer{
		ID:           int(c.ID),
		Name:         c.Name,
		Email:        c.Email,
		WordpressURL: c.WordpressUrl,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
	}
	return &customer
}

func GetCustomers(customers []*model.Customer, count int) *Customers {
	resp := make([]*Customer, len(customers))
	for i, post := range customers {
		resp[i] = GetCustomer(post)
	}
	return &Customers{
		Customers: resp,
		Paginate: Paginate{
			Count: count,
			Total: len(customers),
		},
	}
}
