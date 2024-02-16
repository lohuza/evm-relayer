package pkg

import "net/http"

type Response struct {
	Code    int32  `json:"-"`
	Payload any    `json:"payload"`
	Error   string `json:"error"`
}

func NewResponseFromError(err error) Response {
	response := Response{
		Code:    400,
		Payload: nil,
		Error:   err.Error(),
	}

	if relayerErr, ok := err.(*RelayerError); ok {
		response.Code = relayerErr.Code
	}

	return response
}

func NewOkResponse(payload any) Response {
	return Response{
		Code:    http.StatusOK,
		Payload: payload,
		Error:   EmptyString,
	}
}

func NewCreatedResponse(payload any) Response {
	return Response{
		Code:    http.StatusCreated,
		Payload: payload,
		Error:   EmptyString,
	}
}
