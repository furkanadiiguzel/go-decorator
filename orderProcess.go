package main

import "fmt"

// OrderProcessor is the interface for processing orders.
type OrderProcessor interface {
	ProcessOrder() int
}

// BaseOrderProcessor represents the base order without any optional services.
type BaseOrderProcessor struct{}

func (o *BaseOrderProcessor) ProcessOrder() int {
	return 100 // Base cost for the order
}

// OrderDecorator is the interface for order decorators.
type OrderDecorator interface {
	Decorate(OrderProcessor) OrderProcessor
}

// GiftWrapDecorator is a decorator that adds gift wrapping to the order.
type GiftWrapDecorator struct{}

func (g *GiftWrapDecorator) Decorate(base OrderProcessor) OrderProcessor {
	return &struct {
		OrderProcessor
	}{
		OrderProcessor: base,
	}
}

func (g *GiftWrapDecorator) ProcessOrder() int {
	baseCost := g.ProcessOrder()
	return baseCost + 20 // Additional cost for gift wrapping
}

// DiscountDecorator is a decorator that applies a discount to the order.
type DiscountDecorator struct {
	DiscountPercentage int
}

func (d *DiscountDecorator) Decorate(base OrderProcessor) OrderProcessor {
	return &struct {
		OrderProcessor
	}{
		OrderProcessor: base,
	}
}

func (d *DiscountDecorator) ProcessOrder() int {
	baseCost := d.ProcessOrder()
	discountAmount := (baseCost * d.DiscountPercentage) / 100
	return baseCost - discountAmount // Apply discount
}

func main() {
	// Create a base order
	baseOrder := &BaseOrderProcessor{}

	// Decorate the base order with gift wrapping
	orderWithGiftWrap := &GiftWrapDecorator{}
	decoratedOrder1 := orderWithGiftWrap.Decorate(baseOrder)

	// Decorate the order with a discount
	orderWithDiscount := &DiscountDecorator{DiscountPercentage: 10}
	decoratedOrder2 := orderWithDiscount.Decorate(decoratedOrder1)

	// Calculate the total cost of the decorated order
	totalCost := decoratedOrder2.ProcessOrder()

	// Print the total cost
	fmt.Println("Total Cost:", totalCost)
}
