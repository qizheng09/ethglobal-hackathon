package model

type TradeCheckRequest struct {
	Method  string       `json:"method"`
	Jsonrpc string       `json:"jsonrpc"`
	ID      int64        `json:"id"`
	Params  []TradeParam `json:"params"`
}

type TradeParam struct {
	Origin string    `json:"origin"`
	Data   TradeData `json:"data"`
}

type TradeData struct {
	Method  string       `json:"method"`
	Jsonrpc string       `json:"jsonrpc"`
	ID      int64        `json:"id"`
	Params  []TransParam `json:"params"`
}

type TransParam struct {
	Gas   string `json:"gas"`
	Value string `json:"value"`
	From  string `json:"from"`
	To    string `json:"to"`
	Data  string `json:"data"`
}
