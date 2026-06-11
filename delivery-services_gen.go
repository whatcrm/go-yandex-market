package yandexmarket

import (
	"context"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils/deliveryservices"
)

func (c *Client) GetDeliveryServices(ctx context.Context) (*models.GetDeliveryServicesResponse, error) {
	requestURL := c.BaseURL + deliveryservices.GetDeliveryServicesEndpoint
	var err error

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetDeliveryServicesResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
