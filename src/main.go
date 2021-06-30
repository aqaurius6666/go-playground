package main

import (
	"fmt"

	"github.com/aquarius6666/go-playground/src/greeting"
	"github.com/aquarius6666/go-playground/src/mymath"
)

func main() {
	fmt.Println("Hello World")
	greeting.Hello()
	ans := mymath.Sum(3.4, -5.4)
	fmt.Println(ans)
}
