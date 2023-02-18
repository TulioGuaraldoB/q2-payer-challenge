package service

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/TulioGuaraldoB/q2-payer-challenge/config/env"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/dto"
)

type IAuthorizationService interface {
	CheckAuthorizerApi() (*dto.AuthorizerResponse, error)
}

type authorizationService struct {
	httpClient *http.Client
}

func NewAuthorizationService(httpClient *http.Client) IAuthorizationService {
	return &authorizationService{
		httpClient: httpClient,
	}
}

func bindAuthorizerResponse(authorizerJson []byte) (*dto.AuthorizerResponse, error) {
	authorizerResponse := new(dto.AuthorizerResponse)
	if err := json.Unmarshal(authorizerJson, authorizerResponse); err != nil {
		return nil, err
	}

	return authorizerResponse, nil
}

func (s *authorizationService) CheckAuthorizerApi() (*dto.AuthorizerResponse, error) {
	authorizerApiRequest, err := http.NewRequest(http.MethodGet, env.Env.AuthorizerUrl, nil)
	if err != nil {
		return nil, err
	}

	authorizerApiResponse, err := s.httpClient.Do(authorizerApiRequest)
	if err != nil {
		return nil, err
	}

	authorizerApiJson, err := io.ReadAll(authorizerApiResponse.Body)
	if err != nil {
		return nil, err
	}

	return bindAuthorizerResponse(authorizerApiJson)
}
