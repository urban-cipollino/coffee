package entity

type Drink struct {
	Name   string
	Price  int
	Recipe Recipe
}

func (d Drink) IsEnoughMoney(money int) bool {
	return money >= d.Price
}
