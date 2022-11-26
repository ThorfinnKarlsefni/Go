package exercise

import (
	"log"
	"runtime"
)

func TestLog() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}

	where()
	where()
	where()
}
