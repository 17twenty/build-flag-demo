//go:build !ca && !us
// +build !ca,!us

package main

import "log"

type EventEngine struct{}

func (EventEngine) raiseEvent() string {
	log.Panic("Not Implemented")
	return ""
}
