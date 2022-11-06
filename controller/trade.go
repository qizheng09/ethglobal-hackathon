package controller

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"github.com/kataras/iris/v12"
	"hackathon/model"
	"log"
	"net/http"
	"strconv"
)

func TradeCheck(ctx iris.Context) {
	var (
		request  model.TradeCheckRequest
		response model.TradeCheckResponse
	)

	err := ctx.ReadJSON(&request)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, model.CommonResp{Data: err.Error()})
		return
	}

	tokenList := []model.TargetToken{
		{
			TokenType:   "ETH",
			TokenAmount: 0.1,
			TokenValue:  163.1,
			TargetAddress: model.TargetAddress{
				Address:            "0x51E3beAbF739F84D5A1d963cBEc9dEeb2Ac85eb4",
				DaysSinceRegister:  1000,
				CumulativeAuth:     1100,
				SevenDayAuth:       7,
				TodayAuth:          6,
				ThreeDayAccuracy:   99.0,
				LastAuthInterval:   10,
				AcceptTimes:        10,
				LastAcceptInterval: 10,
			},
		},
		{
			TokenType:   "BTC",
			TokenAmount: 0.1,
			TokenValue:  2130.2,
			TargetAddress: model.TargetAddress{
				Address:            "0x89D03a506dB9736D41cF7e3946E8a844E9C09571",
				DaysSinceRegister:  900,
				CumulativeAuth:     231,
				SevenDayAuth:       10,
				TodayAuth:          6,
				ThreeDayAccuracy:   99.3,
				LastAuthInterval:   60,
				AcceptTimes:        30,
				LastAcceptInterval: 15,
			},
		},
	}
	response.Message = "Success"
	response.TokenList = tokenList

	ctx.JSON(response)
}

func tenderlySimulate(from, to, data, gasStr, valueStr string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{Transport: tr}
	buf := new(bytes.Buffer)

	gasValue, _ := strconv.ParseUint(gasStr, 16, 32)
	value, _ := strconv.ParseUint(valueStr, 16, 32)
	json.NewEncoder(buf).Encode(model.SimulateRequest{
		NetworkID: "1",
		From:      from,
		To:        to,
		Input:     data,
		Gas:       gasValue,
		Value:     value,
	})

	req, _ := http.NewRequest(http.MethodPost,
		"https://api.tenderly.co/api/v1/account/sniperusopp/project/project/simulate", buf)
	req.Header.Add("X-Access-Key", "FoTlOglOe3AgUqzOkL2LlI6We3k4xBNR")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

func Hello(ctx iris.Context) {
	ctx.WriteString("world")
}
