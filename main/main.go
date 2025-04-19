package main

import (
	"fmt"

	"github.com/imkk000/play-go-multiple-modules/addition"
	"github.com/imkk000/play-go-multiple-modules/multiplication"
)

func main() {
	fmt.Println(addition.Add[int64](1, 2))
	fmt.Println(multiplication.Product[int64](3, 4))
}
