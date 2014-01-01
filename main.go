package main

import (
	"os"
	"os/signal"
	"status/cmus"
	"status/dwm"
	"status/date"
	"syscall"
)

func signals() {
	sig := make(chan os.Signal, 5)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	<-sig
	dwm.Status("")
	os.Exit(0)
}

func main() {
	go signals()

	notify := make(chan int)

	go date.Manager(notify)
	go cmus.Manager(notify)

	old := ""
	for {
		<-notify
		format := formatAll()
		if format != old {
			old = format
			dwm.Status(format)
		}
	}
}

func formatAll() string {
	return cmus.Format() + date.Format();
}
