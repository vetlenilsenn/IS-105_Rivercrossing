package state

import (
	"fmt"
	"strings"
)

var landVest = []string{"Korn", "Kylling", "Rev", "Korn"}
var båt = []string{}
var landØst = []string{}

func ViewState() {
	fmt.Print("LandVest: ")
	fmt.Println(strings.Join(landVest, " | "))
	fmt.Print("Båt: ")
	fmt.Println(strings.Join(båt, " | "))
	fmt.Print("LandØst: ")
	fmt.Println(strings.Join(landØst, " | "))
}