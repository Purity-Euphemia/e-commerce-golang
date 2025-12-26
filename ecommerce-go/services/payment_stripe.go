package services

import (
	"fmt"

	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/paymentintent"
)

func StripePayment(amount int64, currency, stripeKey string) (string, error) {
	stripe.Key = stripeKey

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount), // in cents
		Currency: stripe.String(currency),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		return "", err
	}

	return pi.ClientSecret, nil
}
