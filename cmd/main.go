package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/GO-Trainee/GlebL-innotaxi-userservice"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := innoTaxi_userService.Run(ctx, stop); err != nil {
		fmt.Println(err)
	}
}
