package main

import (
	"fmt"

	"github.com/fatih/color"
)

func log(msg string) {
	fmt.Printf(msg)
}

func logError(msg string) {
	log("\n")
	log(fmt.Sprintf("   %s: %s", color.RedString("error"), msg))
	log("\n")
}
