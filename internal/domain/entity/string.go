package entity

import "time"

type String struct {
	ID         string
	Length     int
	Durability int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	TemplateString
}

type TemplateString struct {
	ID                string
	Name              string
	DefaultDurability int
	Type              StringType
}

type StringType int

const (
	StringTypeUnknown StringType = iota
	StringTypeA
	StringTypeB
	StringTypeC
	StringTypeD
	StringTypeE
	StringTypeF
)

func (t StringType) String() string {
	switch t {
	case StringTypeA:
		return "A"
	case StringTypeB:
		return "B"
	case StringTypeC:
		return "C"
	case StringTypeD:
		return "D"
	case StringTypeE:
		return "E"
	case StringTypeF:
		return "F"
	default:
		return "Unknown"
	}
}

func (s *String) Consume(amount int) {
	if s.Durability >= amount {
		s.Durability -= amount
	} else {
		s.Durability = 0
	}
}

func (s *String) CanUse() bool {
	return s.Durability > 0
}
