package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	pattern := regexp.MustCompile(`^0\.0\.0\.0 [\S]+?$`)
	f, err := os.Open("hosts")
	if err != nil {
		log.Fatalf("Error opening file hosts: %v", err)
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		if pattern.MatchString(line) {
			fields := strings.Split(line, " ")
			fmt.Println(fields[1])
		}
	}
	return
}
