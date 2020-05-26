package strategy

// StrategyBuilder
type StrategyBuilder func(moves []Move) Strategy

// Strategy returns how good the Strategy evaluates the Move to be.
// Usually on the range [0,+1].
//
// 1.0 indicates the highest possible signal to do the move.
// 0.0 indicates no signal to do the move.
type Strategy func(move Move) float64

// Constant returns a constant value.
func Constant(c float64) Strategy {
	return func(move Move) float64 {
		return c
	}
}

// Builder converts a raw Strategy into a Builder.
func Builder(s Strategy) StrategyBuilder {
	return func([]Move) Strategy {
		return s
	}
}

// Vector combines a StrategyBuilder and how strongly to weigh the strategy.
type VectorBuilder struct {
	Builder StrategyBuilder

	// Weight is how strongly to consider this strategy.
	// Higher means stronger.
	// 0.0 means don't consider.
	// Negative means to invert it.
	Weight float64
}

// Vector combines a Strategy and how strongly to weigh the strategy.
type Vector struct {
	Strategy Strategy
	Weight   float64
}
