# Decorator Pattern Examples in Go

This repository showcases examples of the decorator pattern implemented in the Go programming language. The decorator pattern is a structural design pattern that allows behavior to be added to individual objects, either statically or dynamically, without affecting the behavior of other objects from the same class.

## Middleware Example

### Files:
- **middleware.go:** Illustrates the decorator pattern in the context of HTTP middleware. It includes a basic HTTP server with logging and authentication middlewares.

### Usage:
1. Uncomment the `main` function in `middleware.go`.
2. Run the program to start the HTTP server.
3. Open your web browser and navigate to [http://localhost:8080](http://localhost:8080).
4. Check the console for request logging.

## Order Processing Example

### Files:
- **orderprocess.go:** Demonstrates the decorator pattern for processing orders. It includes a base order processor, gift wrap decorator, and discount decorator.

### Usage:
1. Run the `main` function in `orderprocess.go`.
2. See how decorators are applied to the base order processor.
3. Observe the total cost of the decorated order printed to the console.

## Pizza Topping Example

### Files:
- **pizza.go:** Shows the decorator pattern applied to a pizza ordering scenario. It includes an interface `IPizza` and toppings like `TomatoTopping` and `CheeseTopping`.

### Usage:
1. Uncomment the `main` function in `pizza.go`.
2. Run the program to see how toppings are added to a base pizza to calculate the total cost.

## Taxi Trip Example

### Files:
- **taxi.go:** Applies the decorator pattern to a taxi trip scenario with a base taxi and long/short trip decorators.

### Usage:
1. Uncomment the `main` function in `taxi.go`.
2. Run the program to observe how long and short trip decorators modify the base taxi price.

Feel free to explore, modify, and experiment with these examples to deepen your understanding of the decorator pattern in Go.
