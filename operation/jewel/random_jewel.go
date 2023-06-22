package jewel

import "math/rand"

const (
	redPercentage    = 75
	bluePercentage   = 50
	greenPercentage  = 50
	yellowPercentage = 25
	blackPercentage  = 5
)

func RandomRed() uint64 {
	rnd := rand.Intn(100)
	if rnd < redPercentage {
		return uint64(rand.Intn(3) + 1)
	}
	return 0
}

func RandomBlue() uint64 {
	rnd := rand.Intn(100)
	if rnd < bluePercentage {
		return uint64(rand.Intn(2) + 1)
	}
	return 0
}

func RandomGreen() uint64 {
	rnd := rand.Intn(100)
	if rnd < greenPercentage {
		return uint64(rand.Intn(2) + 1)
	}
	return 0
}

func RandomYellow() uint64 {
	rnd := rand.Intn(100)
	if rnd < yellowPercentage {
		return uint64(rand.Intn(2) + 1)
	}
	return 0
}

func RandomBlack() uint64 {
	rnd := rand.Intn(100)
	if rnd < blackPercentage {
		return 1
	}
	return 0
}
