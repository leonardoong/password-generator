package main

import (
	"fmt"
	"strings"

	"crypto/rand"
	"math/big"
	mRand "math/rand"
	"time"
)

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func main() {
	var (
		minLowerCase   int
		minUpperCase   int
		minNumber      int
		minSpecialChar int
		passwordLength int
	)

	mRand.Seed(time.Now().Unix())
	fmt.Printf("Enter min lower case : ")
	fmt.Scanln(&minLowerCase)
	fmt.Printf("Enter min upper case : ")
	fmt.Scanln(&minUpperCase)
	fmt.Printf("Enter min number : ")
	fmt.Scanln(&minNumber)
	fmt.Printf("Enter min special char : ")
	fmt.Scanln(&minSpecialChar)
	fmt.Printf("Enter password length : ")
	fmt.Scanln(&passwordLength)

	fmt.Println("Generating password...")
	fmt.Println(generatePassword(passwordLength, minLowerCase, minUpperCase, minNumber, minSpecialChar))

	fmt.Println(generatePasswordCrypto(passwordLength, minLowerCase, minUpperCase, minNumber, minSpecialChar))

}

func generatePasswordCrypto(passwordLength, minLowerCase, minUpperCase, minNumber, minSpecialChar int) string {
	var password strings.Builder

	if passwordLength < (minLowerCase + minUpperCase + minNumber + minSpecialChar) {
		return "Password length too short"
	}

	// Set special character
	for i := 0; i < minSpecialChar; i++ {
		length := int64(len(specialCharSet))
		nBig, err := rand.Int(rand.Reader, big.NewInt(length))
		if err != nil {
			panic(err)
		}
		random := nBig.Int64()
		password.WriteString(string(specialCharSet[random]))
	}

	// Set number
	for i := 0; i < minNumber; i++ {
		length := int64(len(numberSet))
		nBig, err := rand.Int(rand.Reader, big.NewInt(length))
		if err != nil {
			panic(err)
		}
		random := nBig.Int64()
		password.WriteString(string(numberSet[random]))
	}

	// Set upper case
	for i := 0; i < minUpperCase; i++ {
		length := int64(len(upperCharSet))
		nBig, err := rand.Int(rand.Reader, big.NewInt(length))
		if err != nil {
			panic(err)
		}
		random := nBig.Int64()
		password.WriteString(string(upperCharSet[random]))
	}

	// Set lower case
	remainingLength := passwordLength - minUpperCase - minNumber - minSpecialChar
	for i := 0; i < remainingLength; i++ {
		length := int64(len(allCharSet))
		nBig, err := rand.Int(rand.Reader, big.NewInt(length))
		if err != nil {
			panic(err)
		}
		random := nBig.Int64()
		password.WriteString(string(allCharSet[random]))
	}

	inRune := []rune(password.String())
	mRand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	return password.String()
}

func generatePassword(passwordLength, minLowerCase, minUpperCase, minNumber, minSpecialChar int) string {
	var password strings.Builder

	if passwordLength < (minLowerCase + minUpperCase + minNumber + minSpecialChar) {
		return "Password length too short"
	}

	// Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := mRand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	// Set number
	for i := 0; i < minNumber; i++ {
		random := mRand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	// Set upper case
	for i := 0; i < minUpperCase; i++ {
		random := mRand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	// Set lower case
	remainingLength := passwordLength - minUpperCase - minNumber - minSpecialChar
	for i := 0; i < remainingLength; i++ {
		random := mRand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}

	inRune := []rune(password.String())
	mRand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	return string(inRune)
}
