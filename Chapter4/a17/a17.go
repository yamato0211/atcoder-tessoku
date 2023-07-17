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
	n := readInt(sc)
	sc.Scan()
	sa := strings.Split(sc.Text(), " ")
	a := make([]int, len(sa)+2)
	for i := 2; i <= n; i++ {
		a[i], _ = strconv.Atoi(sa[i-2])
	}
	sc.Scan()
	sb := strings.Split(sc.Text(), " ")
	b := make([]int, len(sb)+3)
	for i := 3; i <= n; i++ {
		b[i], _ = strconv.Atoi(sb[i-3])
	}
	dp := make([]int, n+1)
	dp[2] = a[2]
	for i := 3; i <= n; i++ {
		if dp[i-1]+a[i] <= dp[i-2]+b[i] {
			dp[i] = dp[i-1] + a[i]
		} else {
			dp[i] = dp[i-2] + b[i]
		}
	}
	place := n
	route := make([]int, 0)
	for {
		route = append([]int{place}, route...)
		if place == 1 {
			break
		}

		if dp[place-1]+a[place] == dp[place] {
			place--
		} else {
			place -= 2
		}
	}
	fmt.Println(len(route))
	for _, v := range route {
		fmt.Printf("%d ", v)
	}
	fmt.Print("\n")
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
