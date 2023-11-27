package main

type ITaxi interface {
	getPrice() int
}

type Taxi struct{}

func (p *Taxi) getPrice() int {
	return 50
}

type LongTaxiTrip struct {
	taxi ITaxi
}

func (c *LongTaxiTrip) getPrice() int {
	taxiPrice := c.taxi.getPrice()

	return taxiPrice - 10
}

type ShortTaxiTrip struct {
	taxi ITaxi
}

func (s ShortTaxiTrip) getPrice() int {
	taxiPrice := s.taxi.getPrice()
	return taxiPrice + 40
}
