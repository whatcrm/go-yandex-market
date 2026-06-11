package yandexmarket

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/orders"
)

func (c *Client) AcceptOrderCancellation(ctx context.Context, campaignId int64, orderId int64, body *models.AcceptOrderCancellationRequest) (*models.EmptyApiResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(orders.AcceptOrderCancellationEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId))
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

func (c *Client) CreateOrder(ctx context.Context, campaignId int64, body *models.CreateOrderRequest) (*models.CreateOrderResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(orders.CreateOrderEndpoint, utils.PathInt("campaignId", campaignId))
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

	var response models.CreateOrderResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetBusinessOrders(ctx context.Context, businessId int64, params *models.GetBusinessOrdersParams, body *models.GetBusinessOrdersRequest) (*models.GetBusinessOrdersResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(orders.GetBusinessOrdersEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.GetBusinessOrdersResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetOrder(ctx context.Context, campaignId int64, orderId int64) (*models.GetOrderResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(orders.GetOrderEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId))
	var err error

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetOrderResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetOrderIdentifiersStatus(ctx context.Context, campaignId int64, orderId int64) (*models.GetOrderIdentifiersStatusResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(orders.GetOrderIdentifiersStatusEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId))
	var err error

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetOrderIdentifiersStatusResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetOrderUpdateOptions(ctx context.Context, campaignId int64, body *models.GetOrderUpdateOptionsRequest) (*models.GetOrderUpdateOptionsResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(orders.GetOrderUpdateOptionsEndpoint, utils.PathInt("campaignId", campaignId))
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

	var response models.GetOrderUpdateOptionsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetOrders(ctx context.Context, campaignId int64, params *models.GetOrdersParams) (*models.GetOrdersResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(orders.GetOrdersEndpoint, utils.PathInt("campaignId", campaignId))
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

	var response models.GetOrdersResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) ProvideOrderDigitalCodes(ctx context.Context, campaignId int64, orderId int64, body *models.ProvideOrderDigitalCodesRequest) (*models.EmptyApiResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(orders.ProvideOrderDigitalCodesEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId))
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

func (c *Client) ProvideOrderItemIdentifiers(ctx context.Context, campaignId int64, orderId int64, body *models.ProvideOrderItemIdentifiersRequest) (*models.ProvideOrderItemIdentifiersResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(orders.ProvideOrderItemIdentifiersEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId))
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

	var response models.ProvideOrderItemIdentifiersResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) SetOrderBoxLayout(ctx context.Context, campaignId int64, orderId int64, body *models.SetOrderBoxLayoutRequest) (*models.SetOrderBoxLayoutResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(orders.SetOrderBoxLayoutEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId))
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

	var response models.SetOrderBoxLayoutResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) SetOrderShipmentBoxes(ctx context.Context, campaignId int64, orderId int64, shipmentId int64, body *models.SetOrderShipmentBoxesRequest) (*models.SetOrderShipmentBoxesResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(orders.SetOrderShipmentBoxesEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId), utils.PathInt("shipmentId", shipmentId))
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

	var response models.SetOrderShipmentBoxesResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) UpdateExternalOrderId(ctx context.Context, campaignId int64, orderId int64, body *models.UpdateExternalOrderIdRequest) (*models.EmptyApiResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(orders.UpdateExternalOrderIdEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId))
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

func (c *Client) UpdateOrder(ctx context.Context, campaignId int64, body *models.UpdateOrderRequest) (*models.UpdateOrderResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(orders.UpdateOrderEndpoint, utils.PathInt("campaignId", campaignId))
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

	var response models.UpdateOrderResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) UpdateOrderItems(ctx context.Context, campaignId int64, orderId int64, body *models.UpdateOrderItemRequest) error {
	requestURL := c.BaseURL + utils.BuildPath(orders.UpdateOrderItemsEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId))
	var err error

	var reqBody io.Reader

	var jsonBody []byte
	jsonBody, err = json.Marshal(body)
	if err != nil {
		return err
	}
	reqBody = bytes.NewReader(jsonBody)

	req, err := http.NewRequestWithContext(ctx, "PUT", requestURL, reqBody)
	if err != nil {
		return err
	}

	if err = c.Send(req, nil); err != nil {
		return err
	}
	return nil

}

func (c *Client) UpdateOrderStatus(ctx context.Context, campaignId int64, orderId int64, body *models.UpdateOrderStatusRequest) (*models.UpdateOrderStatusResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(orders.UpdateOrderStatusEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId))
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

	var response models.UpdateOrderStatusResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) UpdateOrderStatuses(ctx context.Context, campaignId int64, body *models.UpdateOrderStatusesRequest) (*models.UpdateOrderStatusesResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(orders.UpdateOrderStatusesEndpoint, utils.PathInt("campaignId", campaignId))
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

	var response models.UpdateOrderStatusesResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
