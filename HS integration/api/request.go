package api

import (
	"bytes"
	"errors"
	u "hsintegration/utility"
	"io"
	"net/http"
	"time"
)

type HTTPService interface {
	SendRequest(req *Request) (*http.Response, error)
}

var RequestService HTTPService = &Request{}

type Request struct {
	ResourcePath        string            // Resource Path
	URIParams           string            // URL Parameters
	Method              string            // Request Method
	Body                []byte            // Request Body
	For                 string            // Either HS or HM
	ContentType         string            // Content Type header
	ApiKey              map[string]string // API key taken via url
	BearerAuthorization string            // Authorization header

}

// SendRequest ...
func (req *Request) SendRequest(r *Request) (*http.Response, error) {
	request, err := r.PrepareRequest()
	if err != nil {
		return nil, err
	}

	// Create http client with 5 seconds of time out
	client := &http.Client{Timeout: 5 * time.Second}
	response, err := client.Do(request)
	if response == nil {
		return nil, errors.New("failed to get any response")
	}

	return response, nil
}

func (req *Request) Send() (*http.Response, error) {
	return RequestService.SendRequest(req)
}

func (req *Request) url() string {
	return req.endpoint() + req.ResourcePath + req.URIParams
}

func (req *Request) body() io.Reader {
	if u.IsBlank(req.Body) {
		return nil
	}
	return bytes.NewReader(req.Body)
}

// PrepareRequest ...
func (req *Request) PrepareRequest() (*http.Request, error) {
	request, err := req.newRequest()

	if req.ApiKey != nil {
		q := request.URL.Query()
		//TODO: Update it with URIPARAMS
		for key, value := range req.ApiKey {
			q.Add(key, value)
		}
		request.URL.RawQuery = q.Encode()
	}
	request.Body.Close()

	if req.ContentType != "" {
		request.Header.Set("Content-Type", req.ContentType)
	} else {
		request.Header.Set("Content-Type", "application/json")
	}

	if !u.IsBlank(req.BearerAuthorization) {
		request.Header.Set("Authorization", req.BearerToken())
	}
	return request, err
}

func (req *Request) newRequest() (*http.Request, error) {
	request, err := http.NewRequest(req.Method, req.url(), req.body())
	return request, err
}

func (req *Request) endpoint() string {
	switch req.For {
	case "HS":
		return "https://api.hubapi.com/contacts/v1/"
	case "SF_AUTH":
		return "https://login.salesforce.com/services/oauth2/token"
	case "SF_INSTANCE":
		return "https://concretio89-dev-ed.my.salesforce.com"
	default:
		return ""
	}
}

func (req *Request) BearerToken() string {
	return "Bearer " + req.BearerAuthorization
}
