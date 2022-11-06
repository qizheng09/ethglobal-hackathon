package model

type SimulateResponse struct {
	Transaction         Transaction       `json:"Transaction"`
	Simulation          Simulation        `json:"simulation"`
	Contracts           []Contract        `json:"contracts"`
	GeneratedAccessList []GeneratedAccess `json:"generated_access_list"`
}

type Transaction struct {
	Hash              string            `json:"hash"`
	BlockHeader       string            `json:"block_hash"`
	BlockNumber       int64             `json:"block_number"`
	From              string            `json:"from"`
	Gas               int64             `json:"gas"`
	GasPrice          int64             `json:"gas_price"`
	GasFeeCap         int64             `json:"gas_fee_cap"`
	GasTipCap         int64             `json:"gas_tip_cap"`
	CumulativeGasUsed int64             `json:"cumulative_gas_used"`
	GasUsed           int64             `json:"gas_used"`
	EffectiveGasPrice int64             `json:"effective_gas_price"`
	Input             string            `json:"input"`
	Nonce             int64             `json:"nonce"`
	To                string            `json:"to"`
	Index             int64             `json:"index"`
	Value             string            `json:"value"`
	AccessList        string            `json:"omitempty,access_list"`
	Status            bool              `json:"status"`
	Addresses         string            `json:"omitempty,addresses"`
	ContractIDs       string            `json:"omitempty,contract_ids"`
	NetworkId         string            `json:"network_id"`
	Timestamp         string            `json:"timestamp"`
	FunctionSelector  string            `json:"function_selector"`
	L1BlockNumber     int64             `json:"l1_block_number"`
	L1Timestamp       int64             `json:"l1_timestamp"`
	Method            string            `json:"method"`
	DecodedInput      string            `json:"omitempty,decoded_input"`
	CallTrace         []CallTrace       `json:"call_trace"`
	TransactionInfo   []TransactionInfo `json:"transaction_info"`
}

type CallTrace struct {
	CallType     string  `json:"call_type"`
	From         string  `json:"from"`
	To           string  `json:"to"`
	Gas          int64   `json:"gas"`
	GasUsed      int64   `json:"gas_used"`
	Subtraces    int64   `json:"subtraces"`
	TraceAddress []int64 `json:"trace_address"`
	Type         string  `json:"type"`
	Input        string  `json:"input"`
	Output       string  `json:"output"`
	OutOff       int64   `json:"outOff"`
	OutLen       int64   `json:"outLen"`
	GasIn        int64   `json:"gas_in"`
	GasCost      int64   `json:"gas_cost"`
	FromBalance  string  `json:"fromBalance"`
	ToBalance    string  `json:"toBalance"`
}

type TransactionInfo struct {
	ContractId      string `json:"contract_id"`
	BlockNumber     int64  `json:"block_number"`
	TransactionId   string `json:"transaction_id"`
	ContractAddress string `json:"contract_address"`
	Method          string `json:"method"`
	Parameters      string `json:"parameters"`
	IntrinsicGas    int64  `json:"intrinsic_gas"`
	RefundGas       int64  `json:"refund_gas"`
	StackTrace      string `json:"stack_trace"`
	CreatedAt       string `json:"created_at"`
}

type Simulation struct {
	ID               string      `json:"id"`
	ProjectID        string      `json:"project_id"`
	OwnerID          string      `json:"owner_id"`
	NetworkID        string      `json:"network_id"`
	BlockNumber      int64       `json:"block_number"`
	TransactionIndex int64       `json:"transaction_index"`
	From             string      `json:"from"`
	To               string      `json:"to"`
	Input            string      `json:"input"`
	Gas              int64       `json:"gas"`
	GasPrice         string      `json:"gas_price"`
	Value            string      `json:"value"`
	Status           bool        `json:"status"`
	AccessList       string      `json:"omitempty,access_list"`
	QueueOrigin      string      `json:"queue_origin"`
	BlockHeader      BlockHeader `json:"block_header"`
	CreatedAt        string      `json:"created_at"`
}

type BlockHeader struct {
	Number           string `json:"number"`
	Hash             string `json:"hash"`
	StateRoot        string `json:"stateRoot"`
	ParentHash       string `json:"parentHash"`
	Sha3Uncles       string `json:"sha3Uncles"`
	TransactionsRoot string `json:"transactionsRoot"`
	ReceiptsRoot     string `json:"receiptsRoot"`
	LogsBloom        string `json:"logsBloom"`
	Timestamp        string `json:"timestamp"`
	Difficulty       string `json:"difficulty"`
	GasLimit         string `json:"gasLimit"`
	GasUsed          string `json:"gasUsed"`
	Miner            string `json:"miner"`
	ExtraData        string `json:"extraData"`
	MixHash          string `json:"mixHash"`
	Nonce            string `json:"nonce"`
	BaseFeePerGas    string `json:"baseFeePerGas"`
	Size             string `json:"size"`
	TotalDifficulty  string `json:"totalDifficulty"`
	Uncles           string `json:"omitempty,uncles"`
	Transactions     string `json:"omitempty,transactions"`
}

type Contract struct {
}

type GeneratedAccess struct {
}
