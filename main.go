package main

import "log"

func main() {
	v := EventEngine{}
	log.Println("Event - ", v.raiseEvent())
}
