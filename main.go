package main

import "log"

func main() {
	v := EventEngine{}
	log.Println("Event - ", v.raiseEvent())
}

func AdditionalFunc() {
	log.Println("I'm not tied to a specific implementation but I need to be called when Canada frobs widgets")
}
