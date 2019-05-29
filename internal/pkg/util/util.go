package util

import (
	"log"
)

var DefaultTimeMask = "2006-01-02T15:04:05-0700"

func Check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func CheckIf(condition bool, ifTrue interface{}, ifFalse interface{}) interface{} {
	if condition {
		switch ifTrue.(type) {
		case func() interface{}:
			return ifTrue.(func() interface{})()
		}
		return ifTrue
	} else {
		switch ifFalse.(type) {
		case func() interface{}:
			return ifFalse.(func() interface{})()
		}
		return ifFalse
	}
}
