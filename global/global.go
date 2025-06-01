package global

import "log"

var Debug = true

func Log(vals ...interface{}) { // variadic function
	if Debug {
		log.Println(vals...)
	}
}
