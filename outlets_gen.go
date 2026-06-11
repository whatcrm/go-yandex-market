package yandexmarket

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/outlets"
)

func (c *Client) CreateOutlet(ctx context.Context, campaignId int64, body *models.ChangeOutletRequest) (*models.CreateOutletResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(outlets.CreateOutletEndpoint, utils.PathInt("campaignId", campaignId))
	var err error

	var reqBody io.Reader

	var jsonBody []byte
	jsonBody, err = json.Marshal(body)
	if err != nil {
		return nil, err
	}
	reqBody = bytes.NewReader(jsonBody)

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.CreateOutletResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) DeleteOutlet(ctx context.Context, campaignId int64, outletId int64) (*models.EmptyApiResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(outlets.DeleteOutletEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("outletId", outletId))
	var err error

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "DELETE", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.EmptyApiResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetOutlet(ctx context.Context, campaignId int64, outletId int64) (*models.GetOutletResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(outlets.GetOutletEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("outletId", outletId))
	var err error

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetOutletResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetOutlets(ctx context.Context, campaignId int64, params *models.GetOutletsParams) (*models.GetOutletsResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(outlets.GetOutletsEndpoint, utils.PathInt("campaignId", campaignId))
	var err error

	requestURL, err = mergeQuery(requestURL, params)
	if err != nil {
		return nil, err
	}

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetOutletsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) UpdateOutlet(ctx context.Context, campaignId int64, outletId int64, body *models.ChangeOutletRequest) (*models.EmptyApiResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(outlets.UpdateOutletEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("outletId", outletId))
	var err error

	var reqBody io.Reader

	var jsonBody []byte
	jsonBody, err = json.Marshal(body)
	if err != nil {
		return nil, err
	}
	reqBody = bytes.NewReader(jsonBody)

	req, err := http.NewRequestWithContext(ctx, "PUT", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.EmptyApiResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
