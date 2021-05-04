package profit

type Advisor struct{}

func (a Advisor) StrategyFor(stockQuotes []int) int {
	if len(stockQuotes) < 2 {
		return 0
	}
	result := 0
	sellingPrice := stockQuotes[len(stockQuotes)-1]

	for i := 0; i < len(stockQuotes)-1; i++ {
		result += sellingPrice - stockQuotes[i]
	}

	if result < 0 {
		result = 0
	}
	return result
}
