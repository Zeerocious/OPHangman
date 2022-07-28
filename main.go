package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func main() {

	for {

		file, _ := os.Open("OP.txt")
		scanner := bufio.NewScanner(file)
		var onepiece []string

		for scanner.Scan() {
			onepiece = append(onepiece, scanner.Text())
		}
		rand.Seed(time.Now().UnixNano())
		guess := onepiece[rand.Intn(len(onepiece)-1)]
		var answer []string
		var hidden []string // The hidden word used in hangman, displayed as [_ _ _ _ _ _]
		var used []string
		for i := 0; i < len(guess); i++ {
			answer = append(answer, strings.ToLower(string(guess[i])))
			if regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(string(guess[i])) {
				hidden = append(hidden, "_")
			} else {
				hidden = append(hidden, string(guess[i]))
			}

		}
		fmt.Println("Guess the word!")

		for tries := 5; tries > 0; tries-- {
			var letter string
			triesduplicate := false
			fmt.Printf("\n\n%s\n\n", hidden)
			fmt.Printf("You have %d tries. Guess a letter: ", tries)
			fmt.Scan(&letter)
			letter = strings.ToLower(letter)
			for !regexp.MustCompile(`^[a-zA-Z0-9]${1}`).MatchString(letter) || contains(used, letter) {
				fmt.Printf("\nIncorrect input or duplicate, please guess a letter: ")

				fmt.Scan(&letter)
				letter = strings.ToLower(letter)
			}
			used = append(used, letter)
			for i := 0; i < len(guess); i++ {

				if answer[i] == letter {
					if triesduplicate == false {
						tries++
					}
					triesduplicate = true
					hidden[i] = letter
				}
			}
			if stringSlicesEqual(answer, hidden) {
				fmt.Printf("\nCONGRATS!! You guessed %s correctly\n", answer)
				break
			}
			if tries == 1 {
				fmt.Printf("\nYou lost. The correct answer was %s\n", answer)
				break
			}

		}

		var command string
		for command != "y" && command != "n" {
			fmt.Printf("\nDo you wish to continue? Y/N: ")
			fmt.Scan(&command)
			command = strings.ToLower(command)
		}

		if command == "n" {
			fmt.Printf("\nGood Game!\n")
			os.Exit(0)
		}
	}
}

///fmt.Println(guess)
/// make a word with _ as the letter and the word size is length of guess. check if the user guessed a letter correctily, then add that letter to the _ word
/// and append the letter to an empty list to check if it has been used before.
