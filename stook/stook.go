package stook

import (
	"github.com/fatihusta/medianova-go/client/request"
)

type StookService struct {
	request *request.RequestConfig
}

func NewStookService(reqCfg *request.RequestConfig) *StookService {
	return &StookService{request: reqCfg}
}
