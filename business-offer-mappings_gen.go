package yandexmarket

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/businessoffermappings"
)

func (c *Client) AddOffersToArchive(ctx context.Context, businessId int64, body *models.AddOffersToArchiveRequest) (*models.AddOffersToArchiveResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(businessoffermappings.AddOffersToArchiveEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.AddOffersToArchiveResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) DeleteOffers(ctx context.Context, businessId int64, body *models.DeleteOffersRequest) (*models.DeleteOffersResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(businessoffermappings.DeleteOffersEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.DeleteOffersResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) DeleteOffersFromArchive(ctx context.Context, businessId int64, body *models.DeleteOffersFromArchiveRequest) (*models.DeleteOffersFromArchiveResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(businessoffermappings.DeleteOffersFromArchiveEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.DeleteOffersFromArchiveResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GenerateOfferBarcodes(ctx context.Context, businessId int64, body *models.GenerateOfferBarcodesRequest) (*models.GenerateOfferBarcodesResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(businessoffermappings.GenerateOfferBarcodesEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.GenerateOfferBarcodesResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetOfferMappings(ctx context.Context, businessId int64, params *models.GetOfferMappingsParams, body *models.GetOfferMappingsRequest) (*models.GetOfferMappingsResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(businessoffermappings.GetOfferMappingsEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.GetOfferMappingsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) UpdateOfferMappings(ctx context.Context, businessId int64, params *models.UpdateOfferMappingsParams, body *models.UpdateOfferMappingsRequest) (*models.UpdateOfferMappingsResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(businessoffermappings.UpdateOfferMappingsEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.UpdateOfferMappingsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
