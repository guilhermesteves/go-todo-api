package util

import (
	"log"
)

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
