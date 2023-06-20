package main

import "fmt"

// Strategy represents the strategy interface
type Strategy interface {
	CalculatePrice(price float64) float64
}

// DefaultPricingStrategy represents the default pricing strategy
type DefaultPricingStrategy struct{}

// CalculatePrice implements the strategy for DefaultPricingStrategy
func (s DefaultPricingStrategy) CalculatePrice(price float64) float64 {
	// Apply any necessary calculations for default pricing
	// For simplicity, let's assume no modification to the price
	return price
}

// RatingBasedPricingStrategy represents the rating-based pricing strategy
type RatingBasedPricingStrategy struct{}

// CalculatePrice implements the strategy for RatingBasedPricingStrategy
func (s RatingBasedPricingStrategy) CalculatePrice(price float64) float64 {
	// Apply any necessary calculations based on ratings
	// For simplicity, let's assume a fixed discount of 10% for low ratings
	discount := 0.10 * price
	return price - discount
}

// Context represents the context that uses the strategy
type Context struct {
	strategy Strategy
}

// SetStrategy sets the strategy for the context
func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

// CalculatePrice calculates the price using the current strategy
func (c *Context) CalculatePrice(price float64) float64 {
	return c.strategy.CalculatePrice(price)
}

func main() {
	context := Context{}

	// Use DefaultPricingStrategy
	context.SetStrategy(DefaultPricingStrategy{})
	price := 100.0
	calculatedPrice := context.CalculatePrice(price)
	fmt.Printf("Price using DefaultPricingStrategy: %.2f\n", calculatedPrice)

	// Use RatingBasedPricingStrategy
	context.SetStrategy(RatingBasedPricingStrategy{})
	calculatedPrice = context.CalculatePrice(price)
	fmt.Printf("Price using RatingBasedPricingStrategy: %.2f\n", calculatedPrice)
}
