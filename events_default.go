//go:build !ca && !us
// +build !ca,!us

package main

func init() {
	panic("I can't be used on my own")
}

type EventEngine struct{}

func (EventEngine) raiseEvent() string {
	return ""
}

func PlayAnthem() {
}
