package main

import (
	"flag"
	"fmt"
	"strings"
)

const (
	del1 = "#"
	plus = "+"
)

func parse(s, t []string) bool {
	var lenS, lenT = len(s) ,len(t)
	if lenS < lenT && s[lenS-1] != del1 {
		return false
	} else if lenS > lenT {
		return false
	}

	for i := range s {
		if len(s[i]) == 0 || len(t[i]) == 0 {
			return false
		}

		if s[i] == del1 {
			if i < lenS-1 {
				return false
			}
			return true
		}

		if s[i] == plus {
			continue
		}

		if s[i][0] == 0x2B {
			if strings.Contains(t[i], s[i][1:len(s[i])]) {
				return true
			}
			return false
		}

		if s[i][len(s[i])-1] == 0x2B {
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
	var subscribe, topic string
	flag.StringVar(&subscribe, "s", "", "subscribe")
	flag.StringVar(&topic, "t", "", "topic")
	flag.Parse()

	s := strings.Split(subscribe, "/")
	t := strings.Split(topic, "/")
	fmt.Println(parse(s, t))
}
