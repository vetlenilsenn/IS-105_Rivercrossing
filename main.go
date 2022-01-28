package main

import (
	state "github.com/vetlenilsenn/IS-105_Rivercrossing/state"
	"bufio"
	"os"
	"fmt"
)


func main() {

	scanner := bufio.NewScanner(os.Stdin)


	for {
		
		state.ViewState()
		fmt.Println("Hva vil du sende over? --> kylling, rev eller korn")
		scanner.Scan()
		text := scanner.Text()
		if(text != "kylling"){
			fmt.Print("\033[H\033[2J")
			fmt.Println("Du kan ikke sende over det!")
			continue
		}
		state.PutIn()
		state.ViewState()
		state.Crossriver()
		state.ViewState()
		if(text == "kylling"){
			break
		}
	}


}