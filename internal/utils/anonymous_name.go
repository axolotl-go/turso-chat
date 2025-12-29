package utils

import (
	"fmt"
	"math/rand"
)

func RandomName() string {
	adjectives := []string{"Fast", "Silent", "Dark", "Happy", "Blue"}
	animals := []string{"Fox", "Wolf", "Cat", "Owl", "Bear"}

	name := fmt.Sprintf(
		"%s%s_%d",
		adjectives[rand.Intn(len(adjectives))],
		animals[rand.Intn(len(animals))],
		rand.Intn(9999),
	)

	return name
}
