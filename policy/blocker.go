package policy

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// BlockUntilKilled blocks the current thread until a kill signal is received, either of SIGINT, SIGTERM
//
//goland:noinspection GoUnusedExportedFunction
func BlockUntilKilled() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done // Will block here until the user hits ctrl+c
	fmt.Println("")
}
