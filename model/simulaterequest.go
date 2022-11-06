package model

type SimulateRequest struct {
	NetworkID      string `json:"network_id"`
	From           string `json:"from"`
	To             string `json:"to"`
	Input          string `json:"input"`
	Gas            uint64 `json:"gas"`
	GasPrice       string `json:"gas_price"`
	Value          uint64 `json:"value"`
	SaveIfFails    bool   `json:"save_if_fails"`
	Save           bool   `json:"save"`
	SimulationType string `json:"SimulationType"`
}
