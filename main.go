/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/wanglet/collector/cmd"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sig
		// log.WithFields(logrus.Fields{}).Debug("Interrupt received, stopping")
		// clean shutdown
		cancel()
	}()

	mainWorkdir, _ := os.Getwd()

	valueCtx := context.WithValue(ctx, "mainWorkdir", mainWorkdir)

	cmd.ExecuteContext(valueCtx)
}
