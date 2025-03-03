package req

type CustomerQuery struct {
	InstagramTokenStatus *string `query:"instagramTokenStatus"`
	PartialName          *string `query:"partialName"`
	Email                *string `query:"email"`
	PaymentType          *string `query:"paymentType"`
	PaymentStatus        *string `query:"paymentStatus"`
	Pagination
}

type CreateCustomerBody struct {
	Name           string `json:"name" example:"yuki"`
	Email          string `json:"email" example:"yuki@gmail.com"`
	Password       string `json:"password" example:"123456"`
	WordpressURL   string `json:"wordpress_url" example:"example.com"`
	DeleteHashFlag int    `json:"delete_hash_flag" example:"0"`
}

type UpdateCustomerBody struct {
	Name string `json:"name"`
}
