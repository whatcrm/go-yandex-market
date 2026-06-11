package yandexmarket

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/content"
)

func (c *Client) GetCategoryContentParameters(ctx context.Context, categoryId int64, params *models.GetCategoryContentParametersParams) (*models.GetCategoryContentParametersResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(content.GetCategoryContentParametersEndpoint, utils.PathInt("categoryId", categoryId))
	var err error

	requestURL, err = mergeQuery(requestURL, params)
	if err != nil {
		return nil, err
	}

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetCategoryContentParametersResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetOfferCardsContentStatus(ctx context.Context, businessId int64, params *models.GetOfferCardsContentStatusParams, body *models.GetOfferCardsContentStatusRequest) (*models.GetOfferCardsContentStatusResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(content.GetOfferCardsContentStatusEndpoint, utils.PathInt("businessId", businessId))
	var err error

	requestURL, err = mergeQuery(requestURL, params)
	if err != nil {
		return nil, err
	}

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

	var response models.GetOfferCardsContentStatusResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) UpdateOfferContent(ctx context.Context, businessId int64, body *models.UpdateOfferContentRequest) (*models.UpdateOfferContentResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(content.UpdateOfferContentEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.UpdateOfferContentResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
