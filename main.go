package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	argv := os.Args
	if len(argv) == 1 {
		fmt.Fprint(os.Stderr, "usage:<cmd> <need to hash> [byte num]\n")
		os.Exit(1)
	}
	sumString := sum256String(argv[1])

	limit := math.MaxInt32
	if len(argv) > 2 {
		limit, _ = strconv.Atoi(argv[2])
	}
	sp := ""
	for i := 0; i < limit*2 && i < len(sumString); i++ {
		if i%5 == 0 {
			fmt.Printf(sp)
			sp = " "
		}
		fmt.Printf("%c", sumString[i])
	}
	fmt.Printf("\n")
}

func sum256String(str string) string {
	argByBytes := []byte(str)
	sum := sha256.Sum256(argByBytes)
	var buffer bytes.Buffer
	for _, b := range sum {
		buffer.WriteString(fmt.Sprintf("%02X", b))
	}
	return buffer.String()
}
