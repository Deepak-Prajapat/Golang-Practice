package main

import (
	"errors"
	"fmt"
	"testing"
)

type stuServiceMock struct {
	handleGetFn func() (*Student, error)
}

func (mock stuServiceMock) GetStudentById() (*Student, error) {
	//Code
	return mock.handleGetFn()
}

func TestGetStudentByIdWithError(t *testing.T) {
	serviceMock := stuServiceMock{}
	// StudentService = serviceMock
	serviceMock.handleGetFn = func() (*Student, error) {
		fmt.Println("returning error from mock >>> ")
		return nil, errors.New("Did not get student")
	}

	StudentService = serviceMock

	response, err := ResultController.GetResult()

	if err != nil && response == nil {
		if err.Error() != "Did not get student" {
			t.Error("Wrong Error Responding")
		}
	}
}

func TestGetStudentByIdWithoutError(t *testing.T) {
	serviceMock := stuServiceMock{}

	serviceMock.handleGetFn = func() (*Student, error) {
		return &Student{Name: "Deepak", Course: "BIT", Fees: 90000}, nil
	}

	StudentService = serviceMock

	stu, _ := ResultController.GetResult()
	if stu.Name != "Deepak" && stu.Course == "BIT" {
		t.Error("Name is not as expected")
	}
}
