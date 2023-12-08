package rands

import (
	"math/rand"
	"time"
)

var (
	randGen = rand.New(rand.NewSource(time.Now().UnixNano()))
)
