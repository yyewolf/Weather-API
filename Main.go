//go:generate goversioninfo
package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go hostAPI()

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
