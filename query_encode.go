package yandexmarket

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

func mergeQuery(rawURL string, params interface{}) (string, error) {
	if params == nil {
		return rawURL, nil
	}
	values, err := query.Values(params)
	if err != nil {
		return "", err
	}
	if len(values) == 0 {
		return rawURL, nil
	}
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	q := u.Query()
	for k, vs := range values {
		for _, v := range vs {
			q.Add(k, v)
		}
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}
