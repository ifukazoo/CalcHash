package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
)

func main() {
	argv := os.Args
	if len(argv) == 1 {
		fmt.Fprint(os.Stderr, "usage:<cmd> <need to hash> [byte num]\n")
		os.Exit(1)
	}
	bytes := []byte(argv[1])
	sum := sha256.Sum256(bytes)
	num := 0
	if len(argv) > 2 {
		num, _ = strconv.Atoi(argv[2])
	}
	if num != 0 {
		for i := 0; i < num && i < len(sum); i++ {
			fmt.Printf("%02X", sum[i])
		}
	} else {
		// 全部出力
		for _, b := range sum {
			fmt.Printf("%02X", b)
		}
	}
	os.Exit(0)
}
