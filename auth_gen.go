package yandexmarket

import (
	"context"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils/auth"
)

func (c *Client) GetAuthTokenInfo(ctx context.Context) (*models.GetTokenInfoResponse, error) {
	requestURL := c.BaseURL + auth.GetAuthTokenInfoEndpoint
	var err error

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetTokenInfoResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
