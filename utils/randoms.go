package utils

import (
	"math/rand"
	"time"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func CreateRandomInt() int {
	return seededRand.Int()
}
