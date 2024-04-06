package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	m = map[string]string{
		"reset":            "reset what",
		"reset board":      "board fault",
		"board add":        "where to add",
		"board delete":     "no board at all",
		"reboot backplane": "impossible",
		"backplane abort":  "install first",
	}
	noCommand = "unknown command"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for {
		ok := input.Scan()
		command := input.Text()
		if !ok || command == "" {
			return
		}
		v, ok := m[command]
		if ok {
			fmt.Println(v)
			continue
		}
		matchValue, matchCnt := "", 0
		for key, value := range m {
			part1 := strings.Split(key, " ")
			part2 := strings.Split(command, " ")
			if len(part1) != len(part2) {
				continue
			}
			hasMatch := true
			for i := 0; i < len(part1); i++ {
				if !strings.HasPrefix(part1[i], part2[i]) {
					hasMatch = false
				}
			}
			if hasMatch {
				if matchValue == "" {
					matchValue = value
				}
				matchCnt++
			}
		}
		if matchValue == "" || matchCnt > 1 {
			fmt.Println(noCommand)
		} else {
			fmt.Println(matchValue)
		}
	}
}
