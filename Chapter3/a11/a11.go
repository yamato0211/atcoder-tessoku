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
	n, x := readInt2(sc)
	a := readIntArray(sc)
	right := n - 1
	left := 0
	for left <= right {
		mid := (right + left) / 2
		if a[mid] > x {
			right = mid - 1
		}
		if a[mid] < x {
			left = mid + 1
		}
		if a[mid] == x {
			fmt.Println(mid + 1)
			return
		}
	}
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
