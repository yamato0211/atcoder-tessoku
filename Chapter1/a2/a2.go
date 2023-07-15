package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	nx := strings.Split(sc.Text(), " ")
	N, _ := strconv.Atoi(nx[0])
	X, _ := strconv.Atoi(nx[1])

	sc.Scan()
	as := strings.Split(sc.Text(), " ")
	A := make([]int, N)
	for i, a := range as {
		A[i], _ = strconv.Atoi(a)
	}

	ans := false
	for i := 0; i < N; i++ {
		if A[i] == X {
			ans = true
		}
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
