package state

import "fmt"
import "strings"

var landVest = []string{"Korn", "Kylling", "Rev", "Korn"}
var båt = []string{}
var landØst = []string{}


func state() {
	fmt.Print("LandVest: ")
	fmt.Println(strings.Join(landVest, " | "))
	fmt.Print("Båt: ")
	fmt.Println(strings.Join(båt, " | "))
	fmt.Print("LandØst: ")
	fmt.Println(strings.Join(landØst, " | "))
}