//go:build us
// +build us

package main

import "log"

var _ eventEngine = (*EventEngine)(nil)

type EventEngine struct {
}

func (EventEngine) raiseEvent() string {
	return "I send Emails about 🦅 MURICA things!!"
}

func PlayAnthem() {
	log.Println("OH SAY CAN YOU SEEEEEEEEEE, BY THE DAWNS EARLY LIIIIGHT!")
}
