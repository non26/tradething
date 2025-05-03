package res

import (
	appresponse "github.com/non26/tradepkg/pkg/bn/app_response"
)

type CloseByClientIds struct {
	Response []CloseByClientIdsResponse
}

type CloseByClientIdsResponse struct {
	Message      string `json:"message"`
	Code         string `json:"code"`
	Symbol       string `json:"symbol"`
	PositionSide string `json:"positionSide"`
	ClientId     string `json:"clientId"`
}

func (c *CloseByClientIds) Add(response CloseByClientIdsResponse) {
	c.Response = append(c.Response, response)
}

func (c *CloseByClientIds) AddFailed(message string, symbol string, positionSide string, clientId string) {
	if symbol == "" {
		symbol = appresponse.Unknown
	}
	if positionSide == "" {
		positionSide = appresponse.Unknown
	}

	c.Response = append(c.Response, CloseByClientIdsResponse{
		Message:      message,
		Code:         appresponse.FailCode,
		Symbol:       symbol,
		PositionSide: positionSide,
		ClientId:     clientId,
	})
}

func (c *CloseByClientIds) AddSuccess(symbol string, positionSide string, clientId string) {
	c.Response = append(c.Response, CloseByClientIdsResponse{
		Message:      appresponse.SuccessMsg,
		Code:         appresponse.SuccessCode,
		Symbol:       symbol,
		PositionSide: positionSide,
		ClientId:     clientId,
	})
}
