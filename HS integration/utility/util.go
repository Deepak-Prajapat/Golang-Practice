package utility

import (
	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
	"io"
	"net/http"
	"reflect"
)

var Log *logrus.Entry

var R *render.Render

// SetupService ...
func SetupService(log *logrus.Entry, r *render.Render) {
	Log = log
	R = r
}

func IsBlank(v interface{}) bool {
	if v == 0 {
		return true
	}
	if v == "" {
		return true
	}
	if v == nil {
		return true
	}
	return false
}

func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)

	if s.IsNil() {
		return nil
	}

	ret := make([]interface{}, s.Len())

	if s.Kind() != reflect.Slice {
		ret = append(ret, slice)
		return ret
	}

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func HTTPRequest(method string, url string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(method, url, body)
	return req
}
