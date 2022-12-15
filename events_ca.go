//go:build ca
// +build ca

package main

var _ eventEngine = (*EventEngine)(nil)

type EventEngine struct {
}

func (EventEngine) raiseEvent() string {
	return "I send emails about 🍁 Canada things"
}
