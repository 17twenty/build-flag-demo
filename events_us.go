//go:build us
// +build us

package main

var _ eventEngine = (*EventEngine)(nil)

type EventEngine struct {
}

func (EventEngine) raiseEvent() string {
	return "I send Emails about 🦅 'MURICA things!!"
}
