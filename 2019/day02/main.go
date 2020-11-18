package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/jeremybaumont/go-advent-of-code/2019/day02/intcode"
)

var inputFile = flag.String("inputFile", "day02_input.modified.txt", "Relative file path to use as input.")
var part2 = flag.Bool("part2", false, "second puzzle logic.")

func buildSliceIntcodeProgram(s string) ([]int64, error) {
	var intcodeProgram []int64
	integers := strings.Split(s, ",")
	for _, integer := range integers {
		i, err := strconv.ParseInt(integer, 10, 64)
		if err != nil {
			return []int64{}, err
		}
		intcodeProgram = append(intcodeProgram, i)
	}
	return intcodeProgram, nil
}

func buildStringIntcodeProgram(integers []int64) string {
	var stringIntcodeProgram []string
	for _, integer := range integers {
		stringIntcodeProgram = append(stringIntcodeProgram, strconv.FormatInt(integer, 10))
	}
	return strings.Join(stringIntcodeProgram, ",")
}

func main() {

	flag.Parse()

	file, err := os.Open(*inputFile)
	if err != nil {
		os.Exit(2)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var outputProgram []int64
	for {
		line, _, err := reader.ReadLine()

		if string(line) == "" {
			break
		}

		intcodeProgram, err := buildSliceIntcodeProgram(string(line))
		if err != nil {
			fmt.Printf("could not parse the input to build an intcode program: %s", err)
			os.Exit(2)
		}
		if err == io.EOF {
			break
		}
		if *part2 {
			continue
		} else {
			outputProgram, err = intcode.RunIntcodeProgram(intcodeProgram)
			if err != nil {
				fmt.Printf("execution of intcode program goes in error: %s", err)
			}
		}
		stringOutputProgram := buildStringIntcodeProgram(outputProgram)
		fmt.Printf("%s \n", stringOutputProgram)
	}

}
