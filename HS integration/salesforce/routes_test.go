package salesforce_test

import (
	"bytes"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"hsintegration/api"
	"hsintegration/salesforce"
	"hsintegration/test"
	u "hsintegration/utility"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"
	"time"
)

//New Work For Testing
type mockRequest struct{}

var responseType string
var sfAuthorized bool

func (d *mockRequest) SendRequest(req *api.Request) (*http.Response, error) {
	request, err := req.PrepareRequest()
	if err != nil {
		return nil, err
	}

	tempResponseType := responseType
	if responseType[:2] == "SF" && !sfAuthorized {
		tempResponseType = "SF_Authorization_Token"
		sfAuthorized = true
	}
	//Prepare with go-vcr
	r, err := recorder.New(path.Join("../test/fixtures/http/", tempResponseType))

	if err != nil {
		log.Println("an error occurred while setting up recorder in http.client")
	}

	defer r.Stop() // Make sure recorder is stopped once done with it
	client := &http.Client{Timeout: 5 * time.Second, Transport: r}
	response, err := client.Do(request)
	if response == nil {
		return nil, errors.New("failed to get any response")
	}

	return response, nil
}

func TestGetAccountInfo(t *testing.T) {
	test.PrepareRendrer()
	test.PrepareSFConfig()

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	type want struct {
		statusCode int
		errMessage string
	}
	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
		resType string //http response recorder file name
	}{
		{
			name: "without account id in url",
			args: args{
				w: httptest.NewRecorder(),
				r: u.HTTPRequest("GET", "localhost:3000/sf/account?id=", nil),
			},
			want: want{
				statusCode: 502,
				errMessage: "account id is missing",
			},
			wantErr: true,
		},
		{
			name: "account id is not between 15 to 18",
			args: args{
				w: httptest.NewRecorder(),
				r: u.HTTPRequest("GET", "sf/account?id=0015j00000XjY", nil),
			},
			want: want{
				statusCode: 502,
				errMessage: "please provide a valid ID for account\nmsg: ID's length should between 15 to 18",
			},
			wantErr: true,
		},
		{
			name: "valid id and returns salesforce account record",
			args: args{
				w: httptest.NewRecorder(),
				r: u.HTTPRequest("GET", "sf/account?id=0015j00000XjYbQAAV", nil),
			},
			wantErr: false,
			want: want{
				statusCode: 200,
			},
			resType: "SF_TestGetAccountInfo_200_account_fetched",
		},
		{
			name: "no record found in salesforce",
			args: args{
				w: httptest.NewRecorder(),
				r: u.HTTPRequest("GET", "sf/account?id=0015j00000XjYbQAA", nil),
			},
			wantErr: true,
			want: want{
				statusCode: 409,
				errMessage: "AccountManager.NoRecordFoundException: no record found for id = 0015j00000XjYbQAA",
			},
			resType: "SF_TestGetAccountInfo_409_record_not_found",
		},
	}

	api.RequestService = new(mockRequest) //enable mocking of request.Send
	for _, tt := range tests {
		responseType = tt.resType
		sfAuthorized = false

		t.Run(tt.name, func(t *testing.T) {
			salesforce.GetAccountInfo(tt.args.w, tt.args.r)

			if tt.wantErr {
				bodyBytes, _ := io.ReadAll(tt.args.w.Result().Body)
				assert.Equal(t, tt.want.errMessage, string(bodyBytes))
			}

			responseStatusCode := tt.args.w.Result().StatusCode
			assert.Equal(t, tt.want.statusCode, responseStatusCode)
		})
	}
}

//For Update an account in salesforce
func TestUpdateAccount(t *testing.T) {
	test.PrepareRendrer()
	test.PrepareSFConfig()

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	type want struct {
		statusCode int
		errMessage string
	}
	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
		resType string //http response recorder file name
	}{
		{
			name: "successful update record",
			args: args{
				w: httptest.NewRecorder(),
				r: u.HTTPRequest("PUT", "localhost:3000/sf/account", bytes.NewReader([]byte(`{
								"id" : "0015j00000XjYbQAAV",
								"Name" : "Name updated in test method"
							}`))),
			},
			want: want{
				statusCode: 200,
			},
			wantErr: false,
			resType: "SF_TestUpdateAccount_200_record_updated",
		},
		{
			name: "error from salesforce (wrong account id)",
			args: args{
				w: httptest.NewRecorder(),
				r: u.HTTPRequest("PUT", "localhost:3000/sf/account", bytes.NewReader([]byte(`{
								"id" : "0015j00000XjYbQA",
								"Name" : "Name updated in test method"
							}`))),
			},
			want: want{
				statusCode: 409,
				errMessage: "AccountManager.BaseException: Error while updating the record",
			},
			wantErr: true,
			resType: "SF_TestUpdateAccount_409_update_failed",
		},
	}

	api.RequestService = new(mockRequest)
	for _, tt := range tests {
		responseType = tt.resType
		sfAuthorized = false

		t.Run(tt.name, func(t *testing.T) {
			salesforce.UpdateAccount(tt.args.w, tt.args.r)

			if tt.wantErr {
				bodyBytes, _ := io.ReadAll(tt.args.w.Result().Body)
				assert.Equal(t, tt.want.errMessage, string(bodyBytes))
			}

			responseStatusCode := tt.args.w.Result().StatusCode
			assert.Equal(t, tt.want.statusCode, responseStatusCode)
		})
	}
}
