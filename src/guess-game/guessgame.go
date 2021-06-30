package guessgame

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func ParseInt(input string) (int, error) {
	value, err := strconv.Atoi(input)
	if err != nil {
		fmt.Print(err)
		return 0, errors.New("invalid input")
	}
	return value, nil

}

func GetGuess() (string, error) {
	fmt.Printf("Choose a number: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
		return "", errors.New("invalid input")
	}
	input = strings.TrimSpace(input)
	return input, nil
}

func Check(target *int, guess *int) int {
	if *target == *guess {
		return 0
	}
	if *guess > *target {
		return 1
	}
	return -1
}

func GetTarget() int {
	fmt.Println("Game is setup")
	rand.Seed(time.Now().Unix())
	target := int(math.Floor(rand.Float64() * 100))
	return target
}
