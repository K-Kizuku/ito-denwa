package entity

import "time"

type User struct {
	ID        string
	Name      string
	Tel       string
	Credit    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) AddCredit(amount int) {
	if amount > 0 {
		u.Credit += amount
	}
}

func (u *User) DeductCredit(amount int) {
	if amount > 0 && u.Credit >= amount {
		u.Credit -= amount
	} else if u.Credit < amount {
		u.Credit = 0
	}
}

func (u *User) CanAfford(amount int) bool {
	return u.Credit >= amount
}
