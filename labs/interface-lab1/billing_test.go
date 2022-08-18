package main

import (
	"testing"
)

func Test_normal_hours_pricing(t *testing.T) {
	bill := newCustomer("Customer one")

	bill.add(100, 2)
	bill.add(200, 1)

	if bill.finalizeBill() != 400 {
		t.Errorf("Normal hours pricing calculation is wrong")
	}
}

func Test_happy_hours_pricing(t *testing.T) {
	bill := newCustomer("Customer one")

	// TODO: need to apply happy hours pricing
	bill.add(100, 2)
	bill.add(200, 1)

	if bill.finalizeBill() != 200 {
		t.Errorf("Happy hours pricing calculation is wrong")
	}
}

func Test_normal_and_happy_hours_pricing(t *testing.T) {
	// customer starts at normal hours and stays till happy hours
	bill := newCustomer("Customer one")

	// TODO: need to apply normal hours pricing
	bill.add(100, 2)

	// TODO: need to apply happy hours pricing
	bill.add(200, 1)

	if bill.finalizeBill() != 300 {
		t.Errorf("Final bill calculation is wrong")
	}
}
