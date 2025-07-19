package entity

type Room struct {
	ID       string
	Caller   User
	Callee   User
	Balance  int
	String   String
	Distance float64
}

func (r *Room) AddBalance(amount int) {
	if amount > 0 {
		r.Balance += amount
	}
}

func (r *Room) DeductBalance(amount int) {
	if amount > 0 && r.Balance >= amount {
		r.Balance -= amount
	} else if r.Balance < amount {
		r.Balance = 0
	}
}

func (r *Room) CanAfford(amount int) bool {
	return r.Balance >= amount
}

func (r *Room) CanUseString() bool {
	return r.String.CanUse()
}

func (r *Room) UseString(amount int) {
	if r.String.Durability >= amount {
		r.String.Consume(amount)
	} else {
		r.String.Durability = 0
	}
}

func (r *Room) UseTelephoneCard(card *TelephoneCard, amount int) {
	if card.CanUse() && card.Balance >= amount {
		card.Use(amount)
		r.AddBalance(amount)
	} else {
		card.Balance = 0
	}
}
