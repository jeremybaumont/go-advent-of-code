package intcode

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRunIntcodeProgram(t *testing.T) {
	tests := map[string]struct {
		inputProgram  []int64
		outputProgram []int64
	}{
		"add opcode: 1 + 1 = 2": {
			inputProgram:  []int64{1, 0, 0, 0, 99},
			outputProgram: []int64{2, 0, 0, 0, 99},
		},
		"mul opcode: 3 * 2 = 6": {
			inputProgram:  []int64{2, 3, 0, 3, 99},
			outputProgram: []int64{2, 3, 0, 6, 99},
		},
		"mul opcode: 99 * 99 = 9801": {
			inputProgram:  []int64{2, 4, 4, 5, 99, 0},
			outputProgram: []int64{2, 4, 4, 5, 99, 9801},
		},
		"add and mul": {
			inputProgram:  []int64{1, 1, 1, 4, 99, 5, 6, 0, 99},
			outputProgram: []int64{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
		"puzzle example": {
			inputProgram:  []int64{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			outputProgram: []int64{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := RunIntcodeProgram(tc.inputProgram)
			if err != nil {
				t.Fatalf("got unexpected error: %s", err)
			}
			diff := cmp.Diff(tc.outputProgram, got)
			if diff != "" {
				t.Fatalf("got %+v, but expected %+v", got, tc.outputProgram)
			}
		})
	}
}
