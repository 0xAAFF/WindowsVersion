package main

import (
	"fmt"
	wv "gowinver/windows/version"
)

func main() {
	ver, err := wv.OSVersion()
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Println("Window Version:", ver)
}
