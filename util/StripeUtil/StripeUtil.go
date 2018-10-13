package StripeUtil

import (
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func chargeCard(token string) (*stripe.Charge, error) {
	setStripeKey()

	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(25),
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		Description: stripe.String("Rick Palomus LLC Call"),
	}

	params.SetSource(token)

	return charge.New(params)
}

func setStripeKey() {
	if stripe.Key == "" {
		stripe.Key = os.Getenv("STRIPE_KEY")
	}
}
