package main

import (
	"fmt"
	"time"
)

func now() {
	unixNow := time.Now().Unix()

	// Prints output
	fmt.Printf("%v\n", unixNow)

}
func definedTime() {
	// Defining t in UTC for Unix method

	// year int, month Month, day, hour, min, sec, nsec int
	t := time.Date(2020, 11, 14, 11, 30, 32, 0, time.UTC)

	// Calling Unix method
	unix := t.Unix()

	// Prints output
	fmt.Printf("%v\n", unix)

}
func main() {
	now()
	definedTime()
}
