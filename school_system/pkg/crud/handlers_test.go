package crud

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"school/pkg/dbcon"
	"testing"

	"github.com/stretchr/testify/assert"
)

type dbMock struct {
	// mock.Mock
	handleGetFn func(string, *dbcon.Student) error
}

func (d dbMock) studentById(stuId string, stu *dbcon.Student) error {
	return d.handleGetFn(stuId, stu)
}

func TestGetStudentById(t *testing.T) {

	testCases := []struct {
		Name        string
		ID          string
		expectedStu dbcon.Student
		err         ErrorJson
	}{
		{
			Name: "ID is 24",
			ID:   "24",
			expectedStu: dbcon.Student{
				Name: "Testing API",
				Fees: 78943,
				ID:   24,
			},
		}, {
			Name: "ID is 25",
			ID:   "25",
			err: ErrorJson{
				Title: "record not found",
				Body:  "No Student Found For Particular Id = 25",
			},
		},
	}

	serviceMock := dbMock{}
	//// Mock response for studentById
	serviceMock.handleGetFn = func(s1 string, s2 *dbcon.Student) error {
		if s1 == "24" {
			s2.Name = "Testing API"
			s2.Fees = 78943
			s2.ID = 24
			return nil
		}
		return errors.New("record not found")
	}

	//// To access mock method instead of original
	StudentService = serviceMock

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/studentById?stuId="+tc.ID, nil)
			//// It will work as response writer (w)
			rec := httptest.NewRecorder()

			GetStudentById(rec, req)
			res := rec.Result()

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("Cound not read response: %v", err)
			}

			if tc.ID == "24" {
				var student2 dbcon.Student
				err = json.Unmarshal([]byte(b), &student2)
				if err != nil {
					fmt.Println("error", err.Error())
					panic(err)
				}
				assert.Equal(t, tc.expectedStu, student2)
			} else {
				var errRes ErrorJson
				err = json.Unmarshal([]byte(b), &errRes)
				if err != nil {
					fmt.Println("error", err.Error())
					panic(err)
				}
				assert.Equal(t, tc.err, errRes)
			}
		})
	}
}
