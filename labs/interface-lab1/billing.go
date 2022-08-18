package main

import "fmt"

type CustomerBill struct {
	name   string
	drinks []int
}

func newCustomer(name string) *CustomerBill {
	return &CustomerBill{name: name}
}

func (b *CustomerBill) add(price, quantity int) {
	b.drinks = append(b.drinks, price*quantity)
}

func (b *CustomerBill) finalizeBill() int {
	total := 0
	for _, d := range b.drinks {
		total = total + d
	}
	fmt.Println(fmt.Sprintf("Total due for '%s' is %d: ", b.name, total))
	b.drinks = nil
	return total
}

func main() {
	bill := newCustomer("Customer one")

	bill.add(100, 2)
	bill.add(200, 1)

	bill.finalizeBill()
}
