package yandexmarket

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/returns"
)

func (c *Client) CancelReturn(ctx context.Context, campaignId int64, body *models.CancelReturnRequest) (*models.CancelReturnResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(returns.CancelReturnEndpoint, utils.PathInt("campaignId", campaignId))
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

	var response models.CancelReturnResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) CreateReturn(ctx context.Context, campaignId int64, body *models.CreateReturnRequest) (*models.CreateReturnResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(returns.CreateReturnEndpoint, utils.PathInt("campaignId", campaignId))
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

	var response models.CreateReturnResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetReturn(ctx context.Context, campaignId int64, orderId int64, returnId int64) (*models.GetReturnResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(returns.GetReturnEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId), utils.PathInt("returnId", returnId))
	var err error

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetReturnResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetReturnApplication(ctx context.Context, campaignId int64, orderId int64, returnId int64) error {
	requestURL := c.BaseURL + utils.BuildPath(returns.GetReturnApplicationEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId), utils.PathInt("returnId", returnId))
	var err error

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, reqBody)
	if err != nil {
		return err
	}

	if err = c.Send(req, nil); err != nil {
		return err
	}
	return nil

}

func (c *Client) GetReturnAvailableDecisions(ctx context.Context, businessId int64, body *models.GetReturnAvailableDecisionsRequest) (*models.GetReturnAvailableDecisionsResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(returns.GetReturnAvailableDecisionsEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.GetReturnAvailableDecisionsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetReturnPhoto(ctx context.Context, campaignId int64, orderId int64, returnId int64, itemId int64, imageHash string) error {
	requestURL := c.BaseURL + utils.BuildPath(returns.GetReturnPhotoEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId), utils.PathInt("returnId", returnId), utils.PathInt("itemId", itemId), utils.PathStringParam("imageHash", imageHash))
	var err error

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, reqBody)
	if err != nil {
		return err
	}

	if err = c.Send(req, nil); err != nil {
		return err
	}
	return nil

}

func (c *Client) GetReturns(ctx context.Context, campaignId int64, params *models.GetReturnsParams) (*models.GetReturnsResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(returns.GetReturnsEndpoint, utils.PathInt("campaignId", campaignId))
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

	var response models.GetReturnsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) SetReturnDecision(ctx context.Context, campaignId int64, orderId int64, returnId int64, body *models.SetReturnDecisionRequest) (*models.EmptyApiResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(returns.SetReturnDecisionEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId), utils.PathInt("returnId", returnId))
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

	var response models.EmptyApiResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) SubmitReturnDecision(ctx context.Context, campaignId int64, orderId int64, returnId int64, body *models.SubmitReturnDecisionRequest) (*models.EmptyApiResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(returns.SubmitReturnDecisionEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId), utils.PathInt("returnId", returnId))
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

	var response models.EmptyApiResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
