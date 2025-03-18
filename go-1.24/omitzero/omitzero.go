package omitzero

import "time"

// BEGIN TYPES OMIT
type Money struct {
	Amount   int
	Currency string
}

type Product struct {
	Discount   Money
	CreatedAt  time.Time
	ModifiedAt time.Time
}

// END TYPES OMIT

// BEGIN ZERO OMIT
func (m Money) IsZero() bool {
	return m.Amount == 0 && m.Currency == ""
}

// END ZERO OMIT
