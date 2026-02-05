package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	pass := flag.String("pass", "", "Password to check")
	flag.Parse()

	if *pass == "" {
		fmt.Println("Please provide a password using --pass")
		os.Exit(1)
	}

	result := checkPassword(*pass)

	fmt.Println("Password Strength Check")
	fmt.Println("")
	fmt.Println("Strength :", result.Strength)
	fmt.Println("Score    :", result.Score, "/ 5")

	if len(result.Reasons) > 0 {
		fmt.Println("\nSuggestions:")
		for _, r := range result.Reasons {
			fmt.Println("-", r)
		}
	} else {
		fmt.Println("\nYour password looks good!")
	}
}
