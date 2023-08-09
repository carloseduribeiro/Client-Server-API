package client

type ResponseData []ExchangeDto

type ExchangeDto struct {
	Bid string `json:"bid"`
}

type ResponseDto struct {
	Message string `json:"message"`
}
