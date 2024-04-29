package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice"
	"github.com/shopspring/decimal"
)

func main() {
	var a decimal.Decimal
	a = decimal.NewFromFloatWithExponent(1.125123123, -2)
	fmt.Println(a)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		stop()
	}()

	if err := run.Run(ctx); err != nil {
		fmt.Println(err)
	}
}
