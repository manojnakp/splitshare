package component

type Name string

type Response struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
	Data  any    `json:"data"`
}

type Balance struct {
	From   Name    `json:"from"`
	To     Name    `json:"to"`
	Amount float64 `json:"amount"`
}
