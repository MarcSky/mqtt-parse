package main

import (
	"flag"
	"fmt"
	"strings"
)

func parse(subscribe, topic string) bool {
	s := strings.Split(subscribe, "/")
	t := strings.Split(topic, "/")
	if len(s) == 0 || len(t) == 0 {
		return false
	}

	if len(s) < len(t) && s[len(s)-1] != "#" {
		return false
	} else if len(s) > len(t) {
		return false
	}

	for i := range s {
		if s[i] == "" || t[i] == "" {
			return false
		}

		if s[i] == "#" {
			return true
		}

		if s[i] == "+" {
			continue
		}

		if strings.HasPrefix(s[i], "+") {
			if strings.Contains(t[i], s[i][1:len(s[i])]) {
				return true
			}
			return false
		}

		if strings.HasSuffix(s[i], "+") {
			if strings.Contains(t[i], s[i][0:len(s[i])-1]) {
				return true
			}
			return false
		}

		if s[i] != t[i] {
			return false
		}
	}

	return true
}

func main() {
	subscribe := flag.String("subscribe", "", "subscribe")
	topic := flag.String("topic", "", "subscribe")
	flag.Parse()

	fmt.Println("result", parse(*subscribe, *topic))
}
