package main

type eventEngine interface {
	raiseEvent() string
}
