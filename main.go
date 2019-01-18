package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]

	for _, a := range args {
		p := strings.Split(a, "/")
		if len(p) != 2 {
			fmt.Sprintf("ERROR! %v is missing cidr rules. IE: 192.168.1.0/24", a)
			return
		}
		cidr, e := strconv.Atoi(p[1])
		if e != nil || cidr > 32 || cidr < 1 {
			fmt.Sprintf("ERROR! %v is missing cidr rules. IE: 192.168.1.0/24", a)
			return
		}
		ipS := strings.Split(p[0], ".")
		ip := [4]int{}
		for i := 0; i < len(ipS); i++ {
			ip[i], _ = strconv.Atoi(ipS[i])
		}

		o := cidr / 8
		size := int(math.Pow(2, float64(32-cidr)))
		fmt.Println(math.Pow(2, float64((o+1)*8-cidr)))
		start := getStart(ip[o], int(math.Pow(2, float64((o+1)*8-cidr))))
		ip[o] = start
		for o++; o < 4; o++ {
			ip[o] = 0
		}
		fmt.Println(start)
		fmt.Println(strings.Join(genRange(ip, size), ","))

		fmt.Println()
	}
}

func getStart(i int, size int) int {
	end := size
	for i > end {
		end += size
	}
	return end - size
}

func genRange(start [4]int, size int) []string {
	ips := make([]string, 0)
	for ; size > 0; size-- {
		ip := fmt.Sprintf("%v.%v.%v.%v", start[0], start[1], start[2], start[3])
		ips = append(ips, ip)

		for i := 3; i >= 0; i-- {
			start[i]++
			if start[i] <= 255 {
				break
			}
			start[i] = 0
		}
	}

	return ips
}
