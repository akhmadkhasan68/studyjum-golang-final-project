package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

type Clien = resty.Client

type ShipperClient struct {
	*resty.Client
}

func NewShipperAggregatorClient(baseUrl, token string) *ShipperClient {
	c := resty.New()
	c.SetBaseURL(baseUrl)
	c.SetHeader("X-API-Key", token)
	c.SetHeader("Content-Type", "application/json")
	c.SetTimeout(180 * time.Second)

	return &ShipperClient{c}
}

func (c *ShipperClient) GetCountries(ctx context.Context, id, limit, page int) (map[string]interface{}, error) {
	ctxWT, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	params := map[string]string{}
	if id != 0 {
		params["country_id"] = strconv.Itoa(id)
	}
	if limit != 0 {
		params["limit"] = strconv.Itoa(limit)
	}
	if page != 0 {
		params["page"] = strconv.Itoa(page)
	}

	r := c.NewRequest()
	r.SetContext(ctxWT)
	r.SetQueryParams(params)

	res, err := r.Get(c.BaseURL + "/v3/location/countries")
	if err != nil {
		return nil, fmt.Errorf("getting failed: %w", err)

	} else if res.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("response error: %s", res.String())
	}

	result := map[string]interface{}{}
	if err := json.Unmarshal(res.Body(), &result); err != nil {
		return nil, fmt.Errorf("unmarshaling failed: %w", err)
	}

	return result, nil
}

func (c *ShipperClient) GetProvices(ctx context.Context, countryID, id, limit, page int) (*GetProvinces, error) {
	ctxWT, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	params := map[string]string{}
	if id != 0 {
		params["province_id"] = strconv.Itoa(id)
	}
	if limit != 0 {
		params["limit"] = strconv.Itoa(limit)
	}
	if page != 0 {
		params["page"] = strconv.Itoa(page)
	}

	r := c.NewRequest()
	r.SetContext(ctxWT)
	r.SetQueryParams(params)
	r.SetPathParam("country_id", strconv.Itoa(countryID))

	res, err := r.Get(c.BaseURL + "/v3/location/country/{country_id}/provinces")
	if err != nil {
		return nil, fmt.Errorf("getting failed: %w", err)

	} else if res.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("response error: %s", res.String())
	}

	result := new(GetProvinces)
	if err := json.Unmarshal(res.Body(), result); err != nil {
		return nil, fmt.Errorf("unmarshaling failed: %w", err)
	}

	return result, nil
}

func (c *ShipperClient) CreateOrder(ctx context.Context, body CreateOrderRequest) (*CreateOrderResponse, error) {
	ctxWT, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	r := c.NewRequest()
	r.SetContext(ctxWT)
	r.SetBody(body)

	res, err := r.Post(c.BaseURL + "/v3/order")
	if err != nil {
		return nil, fmt.Errorf("getting failed: %w", err)

	} else if res.StatusCode() != http.StatusCreated {
		return nil, fmt.Errorf("response error: %s", res.String())
	}

	result := new(CreateOrderResponse)
	if err := json.Unmarshal(res.Body(), result); err != nil {
		return nil, fmt.Errorf("unmarshaling failed: %w", err)
	}

	return result, nil
}

func (c *ShipperClient) CancelOrder(ctx context.Context, OrderID string, reason string) (*CancelOrderResponse, error) {
	ctxWT, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var result *CancelOrderResponse

	r := c.NewRequest()
	r.SetContext(ctxWT)
	r.SetBody(CancelOrderRequest{
		Reason: reason,
	})
	r.SetResult(&result)

	res, err := r.Delete(c.BaseURL + "/v3/order/" + OrderID)
	if err != nil {
		return nil, fmt.Errorf("getting failed: %w", err)

	} else if res.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("response error: %s", res.String())
	}

	return result, nil
}

func (c *ShipperClient) GetOrderDetailByExternalID(ctx context.Context, OrderNumber string) (*GetOrderDetailResponse, error) {
	ctxWT, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var result *GetOrderDetailResponse

	r := c.NewRequest()
	r.SetContext(ctxWT)
	r.SetResult(&result)

	res, err := r.Get(c.BaseURL + "/v3/order/external-id/" + OrderNumber)
	if err != nil {
		return nil, fmt.Errorf("getting failed: %w", err)

	} else if res.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("response error: %s", res.String())
	}

	return result, nil
}
