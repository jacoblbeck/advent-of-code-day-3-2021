package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	body, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	bits := ReadBits(strings.NewReader(string(body)))

	power := determinePower(bits)

	fmt.Println(power)

}

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadBits(r io.Reader) []string {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var bits []string

	for scanner.Scan() {
		x := scanner.Text()
		bits = append(bits, x)
	}
	return bits
}

func determinePower(bits []string) int64 {
	var gamma, epsilon string

	for i := 0; i < len(bits[0]); i++ {
		var zero, one int
		for _, bit := range bits {
			if string(bit[i]) == "1" {
				one++
			}
			if string(bit[i]) == "0" {
				zero++
			}
		}
		if one > zero {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		} else {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		}
	}

	g, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	e, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return e * g
}
