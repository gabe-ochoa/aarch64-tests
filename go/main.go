package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	arch := runtime.GOARCH
	os := runtime.GOOS
	s := fmt.Sprintf("Hi. I'm an app running on: %s. %s", arch, os)
	_, err := fmt.Print(s)
	if err != nil {
		log.Fatal("Failed to print line")
	}
}
