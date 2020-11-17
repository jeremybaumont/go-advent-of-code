package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/jeremybaumont/go-advent-of-code/2019/day01/fuel"
)

var inputFile = flag.String("inputFile", "day01_input.txt", "Relative file path to use as input.")
var part2 = flag.Bool("part2", false, "second puzzle logic.")

func main() {

	flag.Parse()

	file, err := os.Open(*inputFile)
	if err != nil {
		os.Exit(2)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var requiredFuel int64
	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}
		moduleMass, err := strconv.ParseInt(string(line), 10, 64)
		if err != nil {
			fmt.Printf("Could not parse the line: %s", string(line))
		}
		if *part2 {
			requiredFuel = fuel.CalculateFuelRequiredWithFuelForFuel(moduleMass)
		} else {
			requiredFuel = fuel.CalculateFuelRequired(moduleMass)
		}

		fmt.Printf("%d \n", requiredFuel)
	}

}
