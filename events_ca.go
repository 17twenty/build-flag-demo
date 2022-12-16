//go:build ca
// +build ca

package main

import "log"

var _ eventEngine = (*EventEngine)(nil)

type EventEngine struct {
}

func (EventEngine) raiseEvent() string {
	log.Println("I also frob widgets and call AdditionalFunc()")
	AdditionalFunc()
	return "I send emails about 🍁 Canada things"
}

func PlayAnthem() {
	log.Println("Oh Canada! Oh Canada!")
}
