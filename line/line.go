package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const firstMile = 1052

func main() {
	sc := bufio.NewScanner(os.Stdin)
	mileage := 0.0
	fee := 0
	for sc.Scan() {
		inputs := strings.Split(sc.Text(), " ")
		if len(inputs) != 2 {
			os.Exit(1)
		}
		mile, err := strconv.ParseFloat(inputs[1], 64)
		if err != nil {
			os.Exit(1)
		}
		mile *= 10
		mileage += mile
	}
	for {
		total := mileage / 10
		if total <= firstMile {
			fee += 410
			break
		}
	}
	fmt.Println(fee)
}
