package intcode

import "fmt"

func RunIntcodeProgram(intcodeProgram []int64) ([]int64, error) {
	var integer, dst, pos1, pos2 int64
	for index := 0; index < len(intcodeProgram); index++ {
		integer = intcodeProgram[index]
		switch integer {
		case 1:
			dst = intcodeProgram[index+3]
			pos1 = intcodeProgram[index+1]
			pos2 = intcodeProgram[index+2]
			intcodeProgram[dst] = intcodeProgram[pos1] + intcodeProgram[pos2]
			index = index + 3
		case 2:
			dst = intcodeProgram[index+3]
			pos1 = intcodeProgram[index+1]
			pos2 = intcodeProgram[index+2]
			intcodeProgram[dst] = intcodeProgram[pos1] * intcodeProgram[pos2]
			index = index + 3
		case 99:
			return intcodeProgram, nil
		}
	}
	return []int64{}, fmt.Errorf("did not meet exit opcode: 99")
}
