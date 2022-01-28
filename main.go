package main

import (
	state "github.com/vetlenilsenn/IS-105_Rivercrossing/state"
	"bufio"
	"os"
	"fmt"
)


func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\033[H\033[2J")
	fmt.Println("------------------------")
	state.ViewState()
	fmt.Println("Hvem vil du sende over for å løse dette? (korn, kylling, rev)")

	for {
		scanner.Scan()
		text := scanner.Text()
		if(text != "kylling"){
			fmt.Print("\033[H\033[2J")
			fmt.Println("------------------------")
			state.ViewState()
			fmt.Println("Hvem vil du sende over for å løse dette? (korn, kylling, rev)")
			fmt.Println("!!!Du kan ikke sende over det, skriv på nytt!!!")
			continue
		}
		fmt.Print("\033[H\033[2J")
		fmt.Println("Du sender over kylling")
		fmt.Println("------------------------")
		state.ViewState()
		state.PutIn()
		state.ViewState()
		state.Crossriver()
		state.ViewState()
		if(text == "kylling"){
			break
		}
	}


}