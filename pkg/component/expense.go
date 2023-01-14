package component

import "time"

type Expense struct {
	Id     string    `json:"id"`
	Name   Name      `json:"name"`
	Amount float64   `json:"amount"`
	By     Name      `json:"by"`
	When   time.Time `json:"when"`
	Among  []float64 `json:"among"`
}
