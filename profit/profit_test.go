package profit

import (
	"testing"
)


func TestProfit(t *testing.T) {
	t.Run("Empty input results in zero as return value", func(t *testing.T) {
		advisor := Advisor{}
		want := 0
		got := advisor.StrategyFor([]int{})

		if got != want {
			t.Errorf("want %d, got %d", want, got)
		}
	})

	t.Run("Two ascending input values returns difference of the two values", func(t *testing.T) {
		advisor := Advisor{}
		want := 2
		got := advisor.StrategyFor([]int{1,3})

		if got != want {
			t.Errorf("want %d, got %d", want, got)
		}
	})

	t.Run("Two descending input values returns 0", func(t *testing.T) {
		advisor := Advisor{}
		want := 0
		got := advisor.StrategyFor([]int{3,1})

		if got != want {
			t.Errorf("want %d, got %d", want, got)
		}
	})
}
