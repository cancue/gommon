package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

func Wait() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
