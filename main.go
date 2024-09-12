package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	Reset := "\033[0m"
	Green := "\033[32m"
	Yellow := "\033[33m"
	// Red := "\033[31m"
	// Blue := "\033[34m"
	// Magenta := "\033[35m"
	// Cyan := "\033[36m"
	// Gray := "\033[37m"
	// White := "\033[97m"

	output, err := exec.Command("gh", "pr", "status", "--json", "reviewDecision", "--jq", ".currentBranch.reviewDecision").CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	outputString := string(output)
	isPending := true
	for isPending == true {
		if outputString == "REVIEW_REQUIRED\n" {
			fmt.Println(Yellow + outputString + Reset)
		} else {
			fmt.Println(Green + outputString + Reset)
			isPending = false
		}
		time.Sleep(1 * time.Second)
	}
}
