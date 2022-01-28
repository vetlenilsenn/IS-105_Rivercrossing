package state

import (
	"fmt"
	"strings"
	event "github.com/vetlenilsenn/IS-105_Rivercrossing/event"
)



func ViewState() {
	fmt.Print("LandVest: ")
	fmt.Println(strings.Join(event.landVest(), " | "))
	fmt.Print("Båt: ")
	fmt.Println(strings.Join(event.båt(), " | "))
	fmt.Print("LandØst: ")
	fmt.Println(strings.Join(event.landØst(), " | "))
}