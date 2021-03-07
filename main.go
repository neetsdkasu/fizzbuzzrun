// Command program to show FizzBuzz values from 1 to 24
package main

import (
	"fmt"
	. "github.com/neetsdkasu/fizzbuzz"
)

func main() {
	var i uint8
	for i = 1; i < 25; i++ {
		fmt.Println(FizzBuzz(i))
	}
}
