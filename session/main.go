package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/checkout/session"
)

func main() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	sess, err := session.New(&stripe.CheckoutSessionParams{
		SuccessURL: stripe.String("https://example.com/success"),
		CancelURL:  stripe.String("https://example.com/cancel"),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		Customer: stripe.String(os.Getenv("STRIPE_CUSTOMER")),
		Locale:   stripe.String("ja"),
		SubscriptionData: &stripe.CheckoutSessionSubscriptionDataParams{
			Params: stripe.Params{
				Context: context.Background(),
				Metadata: map[string]string{
					"custom1": "(」・ω・)」うー！",
					"custom2": "(／・ω・)／にゃー！",
				},
			},
			Items: []*stripe.CheckoutSessionSubscriptionDataItemsParams{
				{
					Plan:     stripe.String(os.Getenv("STRIPE_PLAN")),
					Quantity: stripe.Int64(1),
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", sess)
}
