package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/brscn2/orders-api/application"
)

func main() {
	app := application.New(application.LoadConfig())

	// ONLY USE context.Background() WHEN DERIVING A CONTEXT
	// cancel derived context and child contexts, call it upon exit
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
