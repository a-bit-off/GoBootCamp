package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if args, err := getArgs(); err != nil {
		fmt.Println(err)
	} else {
		if errString := MyXargs(args); errString != "" {
			fmt.Printf("%s", errString)
		}
	}
}

func getArgs() ([]string, error) {
	args := os.Args[1:]
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanLines)
	for in.Scan() {
		text, err := in.Text(), in.Err()
		if err != nil {
			return nil, err
		}
		args = append(args, strings.Split(text, " ")...)
	}
	return args, nil
}

func MyXargs(args []string) string {
	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output)
	}
	return ""
}
