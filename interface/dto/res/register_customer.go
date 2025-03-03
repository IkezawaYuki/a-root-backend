package res

import "IkezawaYuki/a-root-backend/domain/model"

type RegisterCustomer struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	WordpressURL string `json:"wordpress_url"`
}

func GetRegisterCustomers(customer *model.Customer) *RegisterCustomer {
	return &RegisterCustomer{
		Name:         customer.Name,
		Email:        customer.Email,
		Password:     customer.Password,
		WordpressURL: customer.WordpressUrl,
	}
}
