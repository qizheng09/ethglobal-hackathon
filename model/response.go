package model

type CommonResp struct {
	Message interface{}
	Data    interface{}
}

type TradeCheckResponse struct {
	Message   interface{}
	TokenList []TargetToken `json:"tokenList"`
}

type TargetToken struct {
	TokenType     string        `json:"tokenType"`
	TokenAmount   float64       `json:"tokenAmount"`
	TokenValue    float64       `json:"tokenValue"`
	TargetAddress TargetAddress `json:"targetAddress"`
}

type TargetAddress struct {
	Address            string  `json:"address"`
	DaysSinceRegister  int64   `json:"daysSinceRegister"`
	CumulativeAuth     int64   `json:"cumulativeAuth"`
	SevenDayAuth       int64   `json:"sevenDayAuth"`
	TodayAuth          int64   `json:"todayAuth"`
	ThreeDayAccuracy   float32 `json:"threeDayAccuracy"`
	LastAuthInterval   int64   `json:"lastAuthInterval"`
	AcceptTimes        int64   `json:"acceptTimes"`
	LastAcceptInterval int64   `json:"lastAcceptInterval"`
}
