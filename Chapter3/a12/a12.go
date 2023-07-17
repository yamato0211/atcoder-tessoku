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
	buf := make([]byte, 4096)
	sc.Buffer(buf, 2000000)
	n, k := readInt2(sc)
	a := readIntArray(sc)
	left := 0
	right := 1000000000
	for left < right {
		mid := (left + right) / 2
		sum := 0
		for i := 0; i < n; i++ {
			sum += mid / a[i]
		}
		if sum >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	fmt.Println(left)
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
