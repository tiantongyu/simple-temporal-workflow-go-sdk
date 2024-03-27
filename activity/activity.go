package activity

import (
	"context"
	"log"
	"math/rand/v2"
)

func GenerateRandomNumber(ctx context.Context) (int, error) {
	log.Printf("Generating a random number between 1 and 100")
	return rand.IntN(100) + 1, nil
}

func PrintEvenNumber(ctx context.Context, number int) error {
	log.Printf("Number %d is even", number)
	return nil
}

func PrintOddNumber(ctx context.Context, number int) error {
	log.Printf("Number %d is odd", number)
	return nil
}
