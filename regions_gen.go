package yandexmarket

import (
	"context"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/regions"
)

func (c *Client) GetRegionsCodes(ctx context.Context) (*models.GetRegionsCodesResponse, error) {
	requestURL := c.BaseURL + regions.GetRegionsCodesEndpoint
	var err error

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetRegionsCodesResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) SearchRegionChildren(ctx context.Context, regionId int64, params *models.SearchRegionChildrenParams) (*models.GetRegionWithChildrenResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(regions.SearchRegionChildrenEndpoint, utils.PathInt("regionId", regionId))
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

	var response models.GetRegionWithChildrenResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) SearchRegionsById(ctx context.Context, regionId int64) (*models.GetRegionByIdResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(regions.SearchRegionsByIdEndpoint, utils.PathInt("regionId", regionId))
	var err error

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetRegionByIdResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) SearchRegionsByName(ctx context.Context, params *models.SearchRegionsByNameParams) (*models.GetRegionsResponse, error) {
	requestURL := c.BaseURL + regions.SearchRegionsByNameEndpoint
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

	var response models.GetRegionsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
