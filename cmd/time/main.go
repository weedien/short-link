package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	week := int(now.Weekday())
	fmt.Printf("Today is %v", week)
}
