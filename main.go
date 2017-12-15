package main

import (
	"bytes"
	"crypto/sha256"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var separator bool

func init() {
	flag.BoolVar(&separator, "p", false, "insert space every 5bytes.")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "Usage:<cmd> [options] <source> [byte num]\n")
		flag.PrintDefaults()
	}
}
func main() {
	flag.Parse()
	nonflagArgv := flag.Args()
	if len(nonflagArgv) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	sumString := sum256String(nonflagArgv[0])

	outputMaxByte := 32 // 256(bit) / 8(bit per byte)
	if len(nonflagArgv) > 1 {
		num, err := strconv.Atoi(nonflagArgv[1])
		if err != nil {
			flag.Usage()
			os.Exit(1)
		}
		outputMaxByte = num
	}

	outputMaxChar := outputMaxByte * 2 // 2char per 1byte
	if separator {
		sp := ""
		for i := 0; i < outputMaxChar && i < len(sumString); i++ {
			if i%5 == 0 {
				fmt.Printf(sp)
				sp = " "
			}
			fmt.Printf("%c", sumString[i])
		}
	} else {
		for i := 0; i < outputMaxChar && i < len(sumString); i++ {
			fmt.Printf("%c", sumString[i])
		}
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
