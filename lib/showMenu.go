package lib

import (
	"fmt"
	"strings"
)

func ShowMenu(){
	MENU := []string{
		"Register",
		"Login",
		"Forgot Password\n",
		"Exit\n",
	}

	fmt.Println("--- Welcome to system ---")
	fmt.Println("")
	
	for i := range len(MENU) {
		if strings.Contains(MENU[i], "Exit") {
			fmt.Printf("0. %v\n", MENU[i])
		} else if i <= len(MENU) {
			fmt.Printf("%d. %v\n", i+1, MENU[i])
		}
	}
}