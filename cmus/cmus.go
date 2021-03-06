package cmus

import (
	"github.com/guelfey/go.dbus"
	"strings"
)

var data string

func Manager(notify chan int) {
	con, _ := dbus.SessionBus()
	con.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='net.sourceforge.cmus'")
	sig := make(chan *dbus.Signal, 10)
	con.Signal(sig)

	cmus := con.Object("net.sourceforge.cmus", "/net/sourceforge/cmus")
	format := func() {
		artist := ""
		cmus.Call("artist", 0).Store(&artist)
		title := ""
		cmus.Call("title", 0).Store(&title)
		data = "dddddd;1c1c1c; \u266b "
		data += artist;
		if artist != "" && title != "" {
			data += " - "
		}
		data += title
		data += " \u266b ;;"
		notify <- 0
	}

	format()
	for s := range sig {
		switch strings.TrimPrefix(s.Name, "net.sourceforge.cmus.") {
		case "track_change": fallthrough
		case "exit":
			format()
		}
	}
}

func Format() string {
	return data
}
