package main

import (
	"fmt"
	"bufio"
	"os/exec"
	"strings"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	for {
		fmt.Print("$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		parts := strings.Split(command[:len(command)-1], " ")

		if parts[0] == "exit" {
			os.Exit(0)
		}
		if parts[0] == "echo" {
			fmt.Println(strings.Join(parts[1:], " "))
			continue	
		}
		if parts[0] == "type" {		
			if parts[1] == "type" || parts[1] == "echo" || parts[1] == "exit" {
				fmt.Println(parts[1], "is a shell builtin")
				continue
			} else if path, err := exec.LookPath(parts[1]); err == nil {
				fmt.Println(parts[1], "is", path)
				continue
			} else {
				fmt.Println(parts[1] + ": not found")
				continue
			}
		}
		if parts[0] != "" {
			if path, _ := exec.LookPath(parts[0]); path != "" {
				out, _ := exec.Command(parts[0], parts[1:]...).Output()
				fmt.Println(out)
				continue
			}
		}

		fmt.Println(command[:len(command)-1] + ": command not found")

	}
}
