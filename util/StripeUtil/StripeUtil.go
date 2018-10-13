package StripeUtil

import (
	"errors"
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func ChargeCard(token string) (*stripe.Charge, error) {
	stripeKeyError := setStripeKey()

	if stripeKeyError != nil {
		return nil, stripeKeyError
	}

	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(25),
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		Description: stripe.String("Rick Palomus LLC Call"),
	}

	params.SetSource(token)

	return charge.New(params)
}

func setStripeKey() error {
	if stripe.Key == "" {
		stripeKey := os.Getenv("STRIPE_KEY")

		if stripe.Key == "" {
			return errors.New("No Stripe Key")
		}

		stripe.Key = stripeKey
	}

	return nil
}
