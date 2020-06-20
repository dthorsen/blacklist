package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	resp, err := http.Get("https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts")
	if err != nil {
		log.Fatalf("failed to fetch SteveBlack list: %v", err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("failed to fetch SteveBlack list: %v", resp.Status)
	}
	pattern := regexp.MustCompile(`^0\.0\.0\.0 [\S]+?$`)

	outputFile, err := os.Create("blocklist.txt")
	if err != nil {
		log.Fatalf("failed to create output file blocklist.txt: %v", err)
	}

	s := bufio.NewScanner(resp.Body)
	for s.Scan() {
		line := s.Text()
		if pattern.MatchString(line) {
			fields := strings.Split(line, " ")
			if fields[1] == "0.0.0.0" {
				continue
			}
			_, err = outputFile.WriteString(fields[1] + "\n")
			if err != nil {
				log.Fatalf("failed to write to output file blocklist.txt: %v", err)
			}
		}
	}
	outputFile.Close()
	resp.Body.Close()
	fmt.Println("converted blocklist successfully")
	return
}
