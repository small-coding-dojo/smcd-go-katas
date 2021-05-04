package profit

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyInputReturnsZero(t *testing.T) {
	advisor := Advisor{}
	want := 0
	got := advisor.StrategyFor([]int{})

	assert.Equal(t, want, got)
}

type TestCase struct {
	input    []int
	expected int
}

func TestDescendingValues(t *testing.T) {
	testCases := []TestCase {
		{[]int{3, 1}, 0},
		{[]int{6, 5, 4, 3, 2, 1}, 0},
	}

	RunTestCases(t, testCases)
}

func RunTestCases(t *testing.T, testCases []TestCase) {

	for _, testCase := range testCases {
		name := fmt.Sprintf("%v returns %d", testCase.input, testCase.expected)
		t.Run(name, func(t *testing.T) {
			advisor := Advisor{}
			got := advisor.StrategyFor(testCase.input)
			assert.Equal(t, testCase.expected, got)
		})
	}
}

func TestAscendingValues(t *testing.T) {
	testCases := []TestCase {
		{[]int{1,3}, 2},
		{[]int{1,2,3}, 3},
		{[]int{1,2,3,4}, 6},
		{[]int{1,2,2,4}, 7},
		{[]int{1, 2, 3, 4, 5, 6}, 15},
	}

	RunTestCases(t, testCases)
}
