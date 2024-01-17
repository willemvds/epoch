package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

var dateFormats = [17]string{
	time.RFC3339Nano, // "2006-01-02T15:04:05.999999999Z07:00"
	time.RFC3339,     // "2006-01-02T15:04:05Z07:00"
	time.UnixDate,    // "Mon Jan _2 15:04:05 MST 2006"
	time.RubyDate,    // "Mon Jan 02 15:04:05 -0700 2006"
	time.RFC822,      // "02 Jan 06 15:04 MST"
	time.RFC822Z,     // "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	time.RFC850,      // "Monday, 02-Jan-06 15:04:05 MST"
	time.RFC1123,     // "Mon, 02 Jan 2006 15:04:05 MST"
	time.RFC1123Z,    // "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	time.Kitchen,     // "3:04PM"
	time.Stamp,       // "Jan _2 15:04:05"
	time.StampMilli,  // "Jan _2 15:04:05.000"
	time.StampMicro,  // "Jan _2 15:04:05.000000"
	time.StampNano,   // "Jan _2 15:04:05.000000000"
	time.DateTime,    // "2006-01-02 15:04:05"
	time.DateOnly,    // "2006-01-02"
	time.TimeOnly,    // "15:04:05"
}

var formatNames = map[string]string{
	time.RFC3339Nano: "RFC3339Nano",
	time.RFC3339:     "RFC3339",
	time.UnixDate:    "UnixDate",
	time.RubyDate:    "RubyDate",
	time.RFC822:      "RFC822",
	time.RFC822Z:     "RFC822Z",
	time.RFC850:      "RFC850",
	time.RFC1123:     "RFC1123",
	time.RFC1123Z:    "RFC1123Z",
	time.Kitchen:     "Kitchen",
	time.Stamp:       "Stamp",
	time.StampMilli:  "StampMilli",
	time.StampMicro:  "StampMicro",
	time.StampNano:   "StampNano",
	time.DateTime:    "DateTime",
	time.DateOnly:    "DateOnly",
	time.TimeOnly:    "TimeOnly",
}

func parseDate(input string) (time.Time, string, error) {
	for _, format := range dateFormats {
		v, err := time.Parse(format, input)
		if err == nil {
			return v, format, nil
		}
	}

	return time.Time{}, "", errors.New("Unable to parse")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: epoch <date string>")
		os.Exit(1)
	}

	v, format, err := parseDate(os.Args[1])
	if err != nil {
		fmt.Printf("Failed to parse %s: %s\n", os.Args[1], err)
		os.Exit(2)
	}

	fmt.Printf("Format used: %s (%s)\n", formatNames[format], format)
	fmt.Println(v.Unix(), "seconds(s) since epoch")
	fmt.Println(v.UnixNano(), "nanoseconds(ns) since epoch")
}
