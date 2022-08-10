package main

import (
	"math/rand"
	"time"
)

func rand5() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(5) + 1
}

func rand13() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(13) + 1
}

func rand13To5() int {
	r := rand13()
	if r > 10 {
		return rand13To5()
	} else {
		return r%5 + 1
	}
}

func rand5To13() int {
	r := (rand5() - 1) + ((rand5() - 1) * 5)
	if r >= 13 {
		return rand5To13()
	} else {
		return r + 1
	}
}
