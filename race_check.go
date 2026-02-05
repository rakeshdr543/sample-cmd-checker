package main

import (
	"fmt"
	"time"
)

var count int

func main() {

	count++

	time.Sleep(1 * time.Second)
	fmt.Println(count)
}
