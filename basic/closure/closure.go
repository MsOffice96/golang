package main

import (
	"log"
	"strings"
)

func main() {
	zip_factory := MakeSuffix(".zip")
	log.Println(zip_factory("golang"))
}

func MakeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
