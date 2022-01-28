package state

import (
	"fmt"
	"strings"
)

var landVest = []string{"Korn", "Kylling", "Rev",}
var båt = []string{""}
var landØst = []string{""}

func ViewState() {
	fmt.Print("LandVest: ")
	fmt.Println(strings.Join(landVest, " | "))
	fmt.Print("Båt: ")
	fmt.Println(strings.Join(båt, " | "))
	fmt.Print("LandØst: ")
	fmt.Println(strings.Join(landØst, " | "))
	fmt.Println("------------------------")
}

func PutIn() {
	for i, v := range landVest {
    if v == "Kylling" {
        landVest = append(landVest[:i], landVest[i+1:]...)
        break
    }
	}
	båt[0] = båt[0] + "Kylling"
}

func Crossriver() {
	for i, v := range båt {
    if v == "Kylling" {
        båt = append(båt[:i], båt[i+1:]...)
        break
    }
	}
	landØst[0] = landØst[0] + "Kylling"
}