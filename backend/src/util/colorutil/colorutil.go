package colorutil

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandomColor() string {
	// Create a new random source seeded with the current time
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	// Generate a random hex color
	return fmt.Sprintf("#%06X", r.Intn(0xFFFFFF+1))
}
