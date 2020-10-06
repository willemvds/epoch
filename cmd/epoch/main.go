package main

import (
	"os"
	"fmt"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: epoch <date string>")
		os.Exit(1)
	}

	v, err := time.Parse(time.RFC3339Nano, os.Args[1])
	if err != nil {
		fmt.Printf("Failed to parse %s: %s\n", os.Args[1], err)
		os.Exit(2)
	}

	fmt.Println(v.Unix(), "seconds(s) since epoch")
	fmt.Println(v.UnixNano(), "nanoseconds(ns) since epoch")
}
