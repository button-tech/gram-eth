package apiClient

import (
	"encoding/json"
	"errors"
	"github.com/google/go-querystring/query"
	"github.com/imroc/req"
	"net/http"
)

var (
	apiUrl string
)

type apiResponse struct {
	Result     interface{} `json:"result"`
	StatusCode uint        `json:"statusCode"`
	Error      error       `json:"error"`
}

func InitApiClient(url string) {
	apiUrl = url
}

func apiCall(method string, path string, data interface{}) *apiResponse {
	var (
		response   *req.Resp
		respErr    error
		authHeader = req.Header{
			"Content-Type": "application/json",
		}
	)

	uri := apiUrl + path
	if data != nil && method != "GET" {
		var err error
		data, err = serialize(data)
		if err != nil {
			return &apiResponse{
				Result:     response,
				StatusCode: http.StatusNotImplemented,
				Error:      err,
			}
		}
	} else if data != nil {
		v, err := query.Values(data)
		if err != nil {
			return &apiResponse{
				Result:     response,
				StatusCode: http.StatusNotImplemented,
				Error:      err,
			}
		}
		uri += "&" + v.Encode()
	}

	if method == "GET" {
		response, respErr = req.Get(uri, authHeader)
	} else if method == "POST" {
		response, respErr = req.Post(uri, authHeader, data)
	} else if method == "PATCH" {
		response, respErr = req.Patch(uri, authHeader, data)
	} else if method == "DELETE" {
		response, respErr = req.Delete(uri, authHeader)
	}

	return &apiResponse{
		Result:     response,
		StatusCode: uint(response.Response().StatusCode),
		Error:      respErr,
	}
}

func (res *apiResponse) response(v interface{}) (interface{}, error) {
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusCreated {
		if err := res.Result.(*req.Resp).ToJSON(&v); err != nil {
			return nil, err
		}
		return v, nil
	}

	apiError, err := res.Result.(*req.Resp).ToString()
	if err != nil {
		return nil, err
	}

	return nil, errors.New(apiError)
}

func serialize(data interface{}) ([]byte, error) {
	switch data.(type) {
	case []byte:
		return data.([]byte), nil
	default:
		result, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
}
