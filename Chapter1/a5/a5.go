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
	n, k := readInt2(sc)
	count := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if (k-i-j) >= 1 && (k-i-j) <= n {
				count++
			}
		}
	}

	fmt.Println(count)
}

func readInt(sc *bufio.Scanner) int {
	sc.Scan()
	rnt, _ := strconv.Atoi(sc.Text())
	return rnt
}

func readInt2(sc *bufio.Scanner) (int, int) {
	sc.Scan()
	ss := strings.Split(sc.Text(), " ")
	r1, _ := strconv.Atoi(ss[0])
	r2, _ := strconv.Atoi(ss[1])
	return r1, r2
}

func readIntArray(sc *bufio.Scanner) []int {
	sc.Scan()
	ss := strings.Split(sc.Text(), " ")
	arr := make([]int, len(ss))
	for i, s := range ss {
		arr[i], _ = strconv.Atoi(s)
	}
	return arr
}
