package date

import (
	"time"
)

func Manager(notify chan int) {
	for {
		notify <- 0
		time.Sleep(time.Second)
	}
}

func Format() string {
	return "dddddd;303030;" + time.Now().Format(" 15:04 ")
}
