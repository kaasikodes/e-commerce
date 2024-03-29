package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomID(length int) (string, error) {
    // Determine how many bytes we need to generate based on the desired length
    numBytes := length / 2 // Each byte is represented by two hexadecimal characters

    // Generate random bytes
    randomBytes := make([]byte, numBytes)
    _, err := rand.Read(randomBytes)
    if err != nil {
        return "", err
    }

    // Convert bytes to hexadecimal string
    randomID := hex.EncodeToString(randomBytes)

    // Ensure that the length of the generated ID is exactly the desired length
    if len(randomID) > length {
        randomID = randomID[:length]
    } else if len(randomID) < length {
        // If the generated ID is shorter than desired, pad it with zeros
        paddingLength := length - len(randomID)
        randomID += generateZeroPadding(paddingLength)
    }

    return randomID, nil
}

func generateZeroPadding(length int) string {
    padding := ""
    for i := 0; i < length; i++ {
        padding += "0"
    }
    return padding
}

