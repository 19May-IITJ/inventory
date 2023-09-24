package id

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func GenerateUniqueID() string {
	// Get current timestamp in nanoseconds
	timestamp := time.Now().Unix()
	// Create a new source with a specific seed
	source := rand.NewSource(timestamp)
	randGen := rand.New(source)

	// Generate a random number
	randomNum := randGen.Intn(10000)

	// Combine timestamp and random number to create a unique ID
	uniqueID := fmt.Sprintf("%d%d", timestamp, randomNum)

	return uniqueID
}

func GenerateUUID() string {
	// Generate a new UUID
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	// Convert UUID to string
	uuidStr := uuidObj.String()

	return uuidStr
}
