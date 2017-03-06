package utils

import (
	"time"
	"math/rand"
)

func GetRandomNum(min, max int) int {
    rand.Seed(int64(time.Now().Nanosecond()))
    return rand.Intn(max - min) + min
}