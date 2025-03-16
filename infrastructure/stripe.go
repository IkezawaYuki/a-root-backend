package infrastructure

import (
	"context"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/client"
)

type StripeDriver interface {
}

func NewStripeDriver(accessToken string) StripeDriver {
	stripeClient := &client.API{}
	stripeClient.Init(accessToken, nil)
	return &stripeDriver{
		client: stripeClient,
	}
}

type stripeDriver struct {
	client *client.API
}

func (d *stripeDriver) GetCustomerByID(ctx context.Context, customerID string) (*stripe.Customer, error) {
	return d.client.Customers.Get(customerID, nil)
}
