package dwm

import (
	"github.com/guelfey/go.dbus"
)

var s = make(chan string)

func Status(status string) {
	con, _ := dbus.SessionBus()
	dwm := con.Object("org.suckless.dwm", "/org/suckless/dwm")
	dwm.Call("set_status", 0, status);
}

