package profit

type Advisor struct{}

func (a Advisor) StrategyFor(stockQuotes []int) int {
	if len(stockQuotes) < 2 {
		return 0
	}

	result := stockQuotes[1] - stockQuotes[0]

	if result < 0 {
		result = 0
	}
	return result
}
