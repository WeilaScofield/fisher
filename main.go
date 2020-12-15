package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	s := rand.Intn(100)

	fmt.Println(s)
}
