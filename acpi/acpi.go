package acpi

import (
	"time"
	"strings"
	"os/exec"
)

var data string

var request = make(chan chan string)

func Manager() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			get()
		case req := <-request:
			req <- format()
		}
	}
}

func Format() string {
	req := make(chan string)
	request <- req
	return <-req
}

func format() string {
	return "dddddd;808080; " + data + " "
}

func get() {
	cmd := exec.Command("acpi")
	out, _ := cmd.Output()
	for _, col := range strings.Split(string(out), ", ") {
		if strings.HasSuffix(col, "%") {
			data = col
			return
		}
	}
	data = "??"
}
