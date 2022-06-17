package helper

import (
	"math/rand"
	"time"
)

func GenerateRandomId() int {
	rand.Seed(time.Now().UTC().UnixNano())

	max := 9999
	min := 1000

	return (rand.Intn(max-min) + min)
}
