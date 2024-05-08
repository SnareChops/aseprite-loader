package trace

import "log"

var Enabled bool

func Log(v ...any) {
	if Enabled {
		log.Println(v...)
	}
}
