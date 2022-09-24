package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	m := make(map[string]int)

	m["quarter_of_a_dozen"] = 3
	m["half_of_a_dozen"] = 6
	m["dozen"] = 12
	m["small_gross"] = 120
	m["gross"] = 144
	m["great_gross"] = 1728

	return m
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	m := make(map[string]int)
	return m
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {

	if units[unit] == 0 {
		return false
	}

	if bill[item] == 0 {
		bill[item] = units[unit]
	} else {
		bill[item] += units[unit]
	}

	return true

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	value, found := GetItem(bill, item)
	unitvalue, flag := GetItem(units, unit)

	if !found || !flag {
		return false
	} else if value-unitvalue == 0 {
		delete(bill, item)
		return true
	} else if value-units[unit] < 0 {
		return false
	} else {
		bill[item] -= units[unit]
		return true
	}

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, found := bill[item]
	return value, found
}
