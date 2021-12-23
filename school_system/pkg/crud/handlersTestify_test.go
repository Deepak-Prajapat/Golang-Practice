package crud

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"school/pkg/dbcon"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type dbMockTestify struct {
	mock.Mock
	// handleGetFn func(string, *dbcon.Student) error
}

func (d *dbMockTestify) studentById(stuId string, stu *dbcon.Student) error {
	args := d.Called(stuId, stu)
	// return d.handleGetFn(stuId, stu)
	return args.Error(0)
}

func TestGetStudentByIdTestify(t *testing.T) {

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

	var stu dbcon.Student
	theDBMock := new(dbMockTestify)
	//// successfully fetch student
	theDBMock.On("studentById", "24", &stu).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*dbcon.Student)
		arg.Name = "Testing API"
		arg.Fees = 78943
		arg.ID = 24
	})

	//// Return error for ID 25
	theDBMock.On("studentById", "25", &stu).Return(errors.New("record not found"))
	StudentService = theDBMock

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/studentById?stuId="+tc.ID, nil)

			//// It will work as response writer (w)
			rec := httptest.NewRecorder()

			GetStudentById(rec, req)
			res := rec.Result()

			if tc.ID == "24" {
				var student2 dbcon.Student
				json.NewDecoder(res.Body).Decode(&student2)
				assert.Equal(t, tc.expectedStu, student2)
			} else {
				var errRes ErrorJson
				json.NewDecoder(res.Body).Decode(&errRes)
				assert.Equal(t, tc.err, errRes)
			}
		})
	}

	theDBMock.AssertNumberOfCalls(t, "studentById", 2)
	theDBMock.AssertExpectations(t)
}
