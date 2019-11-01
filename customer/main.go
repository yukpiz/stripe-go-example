package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

func main() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	c, err := customer.New(&stripe.CustomerParams{
		Params: stripe.Params{
			Context: context.Background(),
			Metadata: map[string]string{
				"custom1": "(」・ω・)」うー！",
				"custom2": "(／・ω・)／にゃー！",
			},
		},
		Email: stripe.String("yukpiz@gmail.com"),
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", c)
}
