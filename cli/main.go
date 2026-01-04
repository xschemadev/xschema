package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/xschemadev/xschema/cmd"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	cmd.Execute(ctx)
}
