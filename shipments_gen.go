package yandexmarket

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/shipments"
)

func (c *Client) ConfirmShipment(ctx context.Context, campaignId int64, shipmentId int64, body *models.ConfirmShipmentRequest) (*models.EmptyApiResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(shipments.ConfirmShipmentEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("shipmentId", shipmentId))
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

func (c *Client) DownloadShipmentAct(ctx context.Context, campaignId int64, shipmentId int64) error {
	requestURL := c.BaseURL + utils.BuildPath(shipments.DownloadShipmentActEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("shipmentId", shipmentId))
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

func (c *Client) DownloadShipmentDiscrepancyAct(ctx context.Context, campaignId int64, shipmentId int64) error {
	requestURL := c.BaseURL + utils.BuildPath(shipments.DownloadShipmentDiscrepancyActEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("shipmentId", shipmentId))
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

func (c *Client) DownloadShipmentInboundAct(ctx context.Context, campaignId int64, shipmentId int64) error {
	requestURL := c.BaseURL + utils.BuildPath(shipments.DownloadShipmentInboundActEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("shipmentId", shipmentId))
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

func (c *Client) DownloadShipmentPalletLabels(ctx context.Context, campaignId int64, shipmentId int64, params *models.DownloadShipmentPalletLabelsParams) error {
	requestURL := c.BaseURL + utils.BuildPath(shipments.DownloadShipmentPalletLabelsEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("shipmentId", shipmentId))
	var err error

	requestURL, err = mergeQuery(requestURL, params)
	if err != nil {
		return err
	}

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

func (c *Client) DownloadShipmentReceptionTransferAct(ctx context.Context, campaignId int64, params *models.DownloadShipmentReceptionTransferActParams) error {
	requestURL := c.BaseURL + utils.BuildPath(shipments.DownloadShipmentReceptionTransferActEndpoint, utils.PathInt("campaignId", campaignId))
	var err error

	requestURL, err = mergeQuery(requestURL, params)
	if err != nil {
		return err
	}

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

func (c *Client) DownloadShipmentTransportationWaybill(ctx context.Context, campaignId int64, shipmentId int64) error {
	requestURL := c.BaseURL + utils.BuildPath(shipments.DownloadShipmentTransportationWaybillEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("shipmentId", shipmentId))
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

func (c *Client) GetShipment(ctx context.Context, campaignId int64, shipmentId int64, params *models.GetShipmentParams) (*models.GetShipmentResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(shipments.GetShipmentEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("shipmentId", shipmentId))
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

	var response models.GetShipmentResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetShipmentOrdersInfo(ctx context.Context, campaignId int64, shipmentId int64) (*models.GetShipmentOrdersInfoResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(shipments.GetShipmentOrdersInfoEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("shipmentId", shipmentId))
	var err error

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetShipmentOrdersInfoResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) SearchShipments(ctx context.Context, campaignId int64, params *models.SearchShipmentsParams, body *models.SearchShipmentsRequest) (*models.SearchShipmentsResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(shipments.SearchShipmentsEndpoint, utils.PathInt("campaignId", campaignId))
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

	req, err := http.NewRequestWithContext(ctx, "PUT", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.SearchShipmentsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) SetShipmentPalletsCount(ctx context.Context, campaignId int64, shipmentId int64, body *models.SetShipmentPalletsCountRequest) (*models.EmptyApiResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(shipments.SetShipmentPalletsCountEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("shipmentId", shipmentId))
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

func (c *Client) TransferOrdersFromShipment(ctx context.Context, campaignId int64, shipmentId int64, body *models.TransferOrdersFromShipmentRequest) (*models.EmptyApiResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(shipments.TransferOrdersFromShipmentEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("shipmentId", shipmentId))
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
