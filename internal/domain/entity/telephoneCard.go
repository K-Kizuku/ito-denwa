package entity

import "time"

type TelephoneCard struct {
	ID        string
	Balance   int
	History   []int
	CreatedAt time.Time
	UpdatedAt time.Time
	TemplateTelephoneCard
}

type TemplateTelephoneCard struct {
	ID             string
	Name           string
	DefaultBalance int
	Type           CardType
}

type CardType int

const (
	CardTypeUnknown CardType = iota
	CardTypeA
	CardTypeB
	CardTypeC
	CardTypeD
	CardTypeE
	CardTypeF
)

func (c CardType) String() string {
	switch c {
	case CardTypeA:
		return "A"
	case CardTypeB:
		return "B"
	case CardTypeC:
		return "C"
	case CardTypeD:
		return "D"
	case CardTypeE:
		return "E"
	case CardTypeF:
		return "F"
	default:
		return "Unknown"
	}
}

func (t *TelephoneCard) Use(amount int) {
	if t.Balance >= amount {
		t.Balance -= amount
	} else {
		t.Balance = 0
	}
}

func (t *TelephoneCard) CanUse() bool {
	return t.Balance > 0
}

func (t *TelephoneCard) AddHistory(amount int) {
	if amount > 0 {
		t.History = append(t.History, amount)
	}
}
