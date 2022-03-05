package hubspot_test

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hsintegration/api"
	"hsintegration/dbaccess"
	"hsintegration/hubspot"
	"hsintegration/test"
	u "hsintegration/utility"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
	"time"
)

const (
	dbDriver = "postgres"
	dbHost   = "localhost"
	dbPort   = 5432

	dbName = "test_db"
	dbUser = "postgres"
	dbPass = "root"

	sslMode = "disable"

	fixtureDir = "../test/fixtures"
)

var (
	fixtures *testfixtures.Loader
)

func TestMain(m *testing.M) {
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", dbHost, dbPort, dbUser, dbPass, dbName)
	config = fmt.Sprintf("%s sslmode=%s", config, sslMode)

	db, err := gorm.Open(postgres.Open(config))
	if err != nil {
		u.Log.Error(err.Error())
	}

	dbaccess.SetupService(db, nil)

	if err != nil {
		log.Fatalf("Error Connecting to db:  %s", err)
	}

	dbb, _ := db.DB()
	fixtures, err = testfixtures.New(
		testfixtures.Database(dbb),
		testfixtures.Dialect(dbDriver),
		testfixtures.Directory(fixtureDir),
	)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func PrepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		fmt.Println("error occurred in db", err)
	}
}

func TestGetAllContacts(t *testing.T) {

	test.PrepareHsConfig()
	test.PrepareRendrer()
	PrepareTestDatabase()

	// To access mock method instead of original
	api.RequestService = new(mockRequest)

	//http test request for testing:
	//hs api key: 1a213976-e94b-4906-9ddd-6bef32c22728

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	type want struct {
		statusCode         int
		errResponseMessage string
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
		resType string
	}{
		{
			name: "Without API key",
			args: args{
				w: httptest.NewRecorder(),
				r: u.HTTPRequest("GET", "localhost:3000/hubspot/contacts?hapikey=", nil),
			},
			want: want{
				statusCode:         400,
				errResponseMessage: "Please provide an api key",
			},
			wantErr: true,
		},
		{
			name: "Wrong API key",
			args: args{
				w: httptest.NewRecorder(),
				r: u.HTTPRequest("GET", "localhost:3000/hubspot/contacts?hapikey=a213976-e94b-4906-9ddd-6bef32c22728", nil),
			},
			want: want{
				statusCode:         400,
				errResponseMessage: "api key is not available in database",
			},
			wantErr: true,
		},
		{
			name: "provided correct api to get response",
			args: args{
				w: httptest.NewRecorder(),
				r: u.HTTPRequest("GET", "localhost:3000/hubspot/contacts?hapikey=1a213976-e94b-4906-9ddd-6bef32c22728", nil),
			},
			want: want{
				statusCode: 200,
			},
			wantErr: false,
			resType: "HS_TestGetAllContact_200_contacts_fetched",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			responseType = tt.resType
			hubspot.GetAllContacts(tt.args.w, tt.args.r)

			body := tt.args.w.Result().Body
			bytesResponse, _ := io.ReadAll(body)

			//Assertions
			assert.Equal(t, tt.want.statusCode, tt.args.w.Result().StatusCode)

			if tt.wantErr {
				assert.Equal(t, tt.want.errResponseMessage, string(bytesResponse))
			}
		})
	}
}

//New Work For Testing
type mockRequest struct{}

var responseType string

func (d *mockRequest) SendRequest(req *api.Request) (*http.Response, error) {
	request, err := req.PrepareRequest()
	if err != nil {
		return nil, err
	}

	//Prepare with go-vcr
	r, err := recorder.New(path.Join("../test/fixtures/http/", responseType))

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

func TestCreateContact(t *testing.T) {
	test.PrepareHsConfig()
	test.PrepareRendrer()
	PrepareTestDatabase()

	body := []byte(`{
  "properties": [
    {
      "property": "email",
      "value": "New@gmail.com"
    },
    {
      "property": "firstname",
      "value": "Black"
    },
    {
      "property": "lastname",
      "value": "Day"
    },
    {
      "property": "website",
      "value": "http://hubspot.com"
    },
    {
      "property": "company",
      "value": "HubSpot"
    },
    {
      "property": "phone",
      "value": "9765321987"
    },
    {
      "property": "address",
      "value": "25 First Street"
    },
    {
      "property": "city",
      "value": "Ajmer"
    },
    {
      "property": "state",
      "value": "MA"
    },
    {
      "property": "zip",
      "value": "02139"
    }
  ]
}`)
	body2 := []byte(`{
  "properties": [
    {
      "property": "email",
      "value": "freshcontact2@gmail.com"
    },
    {
      "property": "firstname",
      "value": "Boy"
    },
    {
      "property": "lastname",
      "value": "Day"
    },
    {
      "property": "website",
      "value": "http://hubspot.com"
    },
    {
      "property": "company",
      "value": "HubSpot"
    },
    {
      "property": "phone",
      "value": "9765321987"
    },
    {
      "property": "address",
      "value": "25 First Street"
    },
    {
      "property": "city",
      "value": "Ajmer"
    },
    {
      "property": "state",
      "value": "MA"
    },
    {
      "property": "zip",
      "value": "02139"
    }
  ]
}`)
	//http test request for testing
	req, err := http.NewRequest("POST", "localhost:3000/hubspot/contact?hapikey=1a213976-e94b-4906-9ddd-6bef32c22728", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal("error", err)
	}

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	type want struct {
		statusCode int
	}
	tests := []struct {
		name       string
		args       args
		setup      func()
		resType    string
		want       want
		wantErr    bool
		errMessage string
	}{
		{
			name: "contact already exists",
			args: args{
				w: httptest.NewRecorder(),
				r: req,
			},
			resType: "HS_TestCreateContact_409_Contact_Exists",
			want:    want{statusCode: 208},
		},
		{
			name: "fresh contact to create",
			args: args{
				w: httptest.NewRecorder(),
				r: u.HTTPRequest("POST", "localhost:3000/hubspot/contact?hapikey=1a213976-e94b-4906-9ddd-6bef32c22728", bytes.NewBuffer(body2)),
			},
			resType: "HS_TestCreateContact_200_created",
			want:    want{statusCode: 200},
		},
		{
			name: "incorrect API key",
			args: args{
				w: httptest.NewRecorder(),
				r: u.HTTPRequest("POST", "localhost:3000/hubspot/contact?hapikey=321654-987654", bytes.NewBuffer(body2)),
			},
			want:       want{statusCode: 502},
			wantErr:    true,
			errMessage: "api key is not available in database",
		},
		{
			name: "missing API key",
			args: args{
				w: httptest.NewRecorder(),
				r: u.HTTPRequest("POST", "localhost:3000/hubspot/contact", bytes.NewBuffer(body2)),
			},
			want:       want{statusCode: 502},
			wantErr:    true,
			errMessage: "api key is missing",
		},
	}

	// To access mock method instead of original
	api.RequestService = new(mockRequest)
	//Ended

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			responseType = tt.resType
			hubspot.CreateContact(tt.args.w, tt.args.r)

			if tt.wantErr {
				bodyBytes, _ := io.ReadAll(tt.args.w.Result().Body)
				assert.Equal(t, tt.errMessage, string(bodyBytes))
			}

			responseStatusCode := tt.args.w.Result().StatusCode
			assert.Equal(t, tt.want.statusCode, responseStatusCode)
		})
	}
}
