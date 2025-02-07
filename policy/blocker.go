package policy

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

//goland:noinspection GoUnusedExportedFunction
func BlockUntilKilled() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done // Will block here until the user hits ctrl+c
	fmt.Println("")
}
