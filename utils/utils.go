package utils

import (
	"bufio"
	"fmt"
	"strings"
)

func ReadVals(reader *bufio.Reader) []string {
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("something bad happened")
		panic(err)
	}
	text = strings.TrimSuffix(text, "\n")
	vals := strings.Split(text, " ")
	return vals
}
