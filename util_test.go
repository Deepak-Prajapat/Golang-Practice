package utility

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	type fields struct {
	}

	type args struct {
		input string
	}

	type want struct {
		output int
	}
	tests := []struct {
		name    string // test case name
		args    args
		want    want // output we want from the test case
		wantErr bool // to handle test case which should return a error or not

	}{
		{
			name: "right string with int",
			args: args{
				input: "100",
			},
			want: want{
				output: 100,
			},
			wantErr: false,
		},
		{
			name: "blank string",
			args: args{
				input: "",
			},
			want: want{
				output: 0,
			},
			wantErr: false,
		},
		{
			name: "string with long int",
			args: args{
				input: "1234567890121316465",
			},
			want: want{
				output: 1234567890121316465,
			},
			wantErr: false,
		},
		// function failes on below 2 test-cases and gives : nil pointer dereference
		// @fail : test-case
		// {
		//  name: "string with non numeric",
		//  args: args{
		//      input: "xyz1000",
		//  },
		//  want: want{
		//      output: 0,
		//  },
		//  wantErr: false,
		// },

		// @fail : test-case
		// {
		//  name: "string with special chars",
		//  args: args{
		//      input: "#@%^&*",
		//  },
		//  want: want{
		//      output: 0,
		//  },
		//  wantErr: false,
		// },
	}

	// execute all the test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int(tt.args.input)
			assert.Equal(t, tt.want.output, got)
		})
	}

}

/// Test Split...
func TestSplit(t *testing.T) {

	type args struct {
		text string
		char string
	}

	type want struct {
		output      []string
		sliceLength int //// to apply assertion on sliceLength
	}

	testCases := []struct {
		name string
		args args
		want want
	}{

		{
			name: "split two words with space",
			args: args{
				text: "hello world",
				char: " ",
			},
			want: want{
				output:      []string{"hello", "world"},
				sliceLength: 2,
			},
		},
		{
			name: "split three words with space",
			args: args{
				text: "let it be",
				char: " ",
			},
			want: want{
				output:      []string{"let", "it", "be"},
				sliceLength: 3,
			},
		},
		{
			name: "Empty string",
			args: args{
				text: "",
				char: " ",
			},
			want: want{
				output:      []string{""},
				sliceLength: 1,
			},
		},
		{
			name: "empty char",
			args: args{
				text: "hey mate",
				char: "",
			},
			want: want{
				output:      []string{"h", "e", "y", " ", "m", "a", "t", "e"},
				sliceLength: 8,
			},
		},
		{
			name: "Character is not present",
			args: args{
				text: "hey mate",
				char: "o",
			},
			want: want{
				output:      []string{"hey mate"},
				sliceLength: 1,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			output := Split(testCase.args.text, testCase.args.char)

			assert.Equal(t, testCase.want.output, output)           // output
			assert.Equal(t, testCase.want.sliceLength, len(output)) // length
		})
	}

}

func TestOrigin(t *testing.T) {
	type args struct {
		text string
	}
	type want struct {
		output string
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "text that have whatsapp prefix",
			args: args{
				text: "whatsapp:hello",
			},
			want: want{
				output: "WhatsApp",
			},
		},
		{
			name: "line prefix",
			args: args{
				text: "line:hello",
			},
			want: want{
				output: "Line",
			},
		},
		{
			name: "gmb prefix",
			args: args{
				text: "gmb:hello",
			},
			want: want{
				output: "Google",
			},
		},
		{
			name: "fb prefix",
			args: args{
				text: "fb:hello",
			},
			want: want{
				output: "Facebook",
			},
		},
		{
			name: "abc prefix",
			args: args{
				text: "abc:hello",
			},
			want: want{
				output: "Apple Business Chat",
			},
		},
		{
			name: "hm prefix that returns Heymarket",
			args: args{
				text: "hm:hello",
			},
			want: want{
				output: "Heymarket",
			},
		},
		{
			name: "fb prefix",
			args: args{
				text: "fb:hello",
			},
			want: want{
				output: "Facebook",
			},
		},
		{
			name: "for default value, prefix not available",
			args: args{
				text: "hii:hello",
			},
			want: want{
				output: "SMS",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := Origin(tc.args.text)
			assert.Equal(t, tc.want.output, output)
		})
	}
}

func TestShowPhone(t *testing.T) {
	type args struct {
		text string
	}
	type want struct {
		output bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "hm prefix",
			args: args{
				text: "hm:random string",
			},
			want: want{
				output: false,
			},
		},
		{
			name: "abc prefix",
			args: args{
				text: "abc:random string",
			},
			want: want{
				output: false,
			},
		},
		{
			name: "whatsapp prefix",
			args: args{
				text: "whatsapp:random string",
			},
			want: want{
				output: true,
			},
		},
		{
			name: "fb prefix",
			args: args{
				text: "fb:text",
			},
			want: want{
				output: false,
			},
		},
		{
			name: "gmb prefix",
			args: args{
				text: "gmb:text",
			},
			want: want{
				output: false,
			},
		},
		{
			name: "line prefix",
			args: args{
				text: "line:text",
			},
			want: want{
				output: false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := ShowPhone(tc.args.text)
			assert.Equal(t, tc.want.output, output)
		})
	}
}

func TestGetStringInBetween(t *testing.T) {
	type args struct {
		str   string
		start string
		end   string
	}
	type want struct {
		output  interface{}
		isEmpty bool
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "string between characters",
			args: args{
				str:   "Heymarket",
				start: "m",
				end:   "t",
			},
			want: want{
				output:  "arke",
				isEmpty: false,
			},
		},
		{
			name: "string between words",
			args: args{
				str:   "Heymarket is a SMS Service",
				start: "Heymarket",
				end:   "Service",
			},
			want: want{
				output:  " is a SMS ",
				isEmpty: false,
			},
		},
		{
			name: "when start(string) is not present in string",
			args: args{
				str:   "Heymarket is a SMS Service",
				start: "shopify",
				end:   "Service",
			},
			want: want{
				output:  "",
				isEmpty: true,
			},
		},
		{
			name: "when end(string) is not present in string",
			args: args{
				str:   "Heymarket",
				start: "H",
				end:   "n",
			},
			want: want{
				output: "",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := GetStringInBetween(tc.args.str, tc.args.start, tc.args.end)
			assert.Equal(t, tc.want.output, output)
		})
	}
}

func TestShopifyMessage(t *testing.T) {
	type args struct {
		msg string
	}
	type want struct {
		output string
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "a string in input",
			args: args{
				msg: "Hello",
			},
			want: want{
				output: "[Shopify] Hello",
			},
		},
		{
			name: "blank string in input",
			args: args{
				msg: "",
			},
			want: want{
				output: "[Shopify] ",
			},
		},
		{
			name: "string contain numbers",
			args: args{
				msg: "123456",
			},
			want: want{
				output: "[Shopify] 123456",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := ShopifyMessage(tc.args.msg)
			assert.Equal(t, tc.want.output, output)
		})
	}
}

func TestTrim(t *testing.T) {
	type args struct {
		str  string
		char string
	}
	type want struct {
		output string
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Trim first and last space",
			args: args{
				str:  " hello ",
				char: " ",
			},
			want: want{
				output: "hello",
			},
		},
		{
			name: "Trim same char from first and last",
			args: args{
				str:  "#hello#",
				char: "#",
			},
			want: want{
				output: "hello",
			},
		},
		{
			name: "Wrong Character",
			args: args{
				str:  "#hello#",
				char: "@",
			},
			want: want{
				output: "#hello#",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := Trim(tc.args.str, tc.args.char)
			assert.Equal(t, tc.want.output, output)
		})
	}
}

func TestCleanPhone(t *testing.T) {
	type args struct {
		phNo interface{}
	}
	type want struct {
		phone string
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "for blank string in phone number input",
			args: args{
				phNo: "",
			},
			want: want{
				phone: "",
			},
		},
		{
			name: "nil in input",
			args: args{
				phNo: nil,
			},
			want: want{
				phone: "",
			},
		},
		{
			name: "phone number under brackets()",
			args: args{
				phNo: "(1234567890)",
			},
			want: want{
				phone: "1234567890",
			},
		},

		{
			name: "with + in starting",
			args: args{
				phNo: "+1234567890",
			},
			want: want{
				phone: "1234567890",
			},
		},
		{
			name: "with - between country code and numbers",
			args: args{
				phNo: "91-1234567890",
			},
			want: want{
				phone: "911234567890",
			},
		},
		{
			name: "unnecessary .",
			args: args{
				phNo: ".1234567890",
			},
			want: want{
				phone: "1234567890",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := CleanPhone(tc.args.phNo)
			assert.Equal(t, tc.want.phone, output)
		})
	}
}

func TestIsBlank(t *testing.T) {
	type args struct {
		param interface{}
	}
	type want struct {
		isBlank bool
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "check for a blank string",
			args: args{
				param: "",
			},
			want: want{
				isBlank: true,
			},
		},
		{
			name: "hello message in string",
			args: args{
				param: "Hello",
			},
			want: want{
				isBlank: false,
			},
		},
		{
			name: "pass nil input",
			args: args{
				param: nil,
			},
			want: want{
				isBlank: true,
			},
		},
		{
			name: "pass numbers input",
			args: args{
				param: 987,
			},
			want: want{
				isBlank: false,
			},
		},
		{
			name: "pass 0 in input",
			args: args{
				param: 0,
			},
			want: want{
				isBlank: true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := IsBlank(tc.args.param)
			assert.Equal(t, tc.want.isBlank, output)
		})
	}
}

func TestPhoneValid(t *testing.T) {
	type args struct {
		v interface{}
	}
	type want struct {
		output bool
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "nil in input",
			args: args{
				v: nil,
			},
			want: want{
				output: false,
			},
		},
		{
			name: "empty string in input",
			args: args{
				v: "",
			},
			want: want{
				output: false,
			},
		},
		{
			name: "less then 10 digit",
			args: args{
				v: "987654321",
			},
			want: want{
				output: false,
			},
		},
		{
			name: "10 digit number",
			args: args{
				v: "9876543210",
			},
			want: want{
				output: true,
			},
		},
		{
			name: "more then 10 digit number",
			args: args{
				v: "987654321012",
			},
			want: want{
				output: true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := PhoneValid(tc.args.v)
			assert.Equal(t, tc.want.output, output)
		})
	}
}

func TestE164Phone(t *testing.T) {
	type args struct {
		v interface{}
	}
	type want struct {
		output string
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "blank string in input",
			args: args{
				v: "",
			},
			want: want{
				output: "",
			},
		},
		{
			name: "for less then 11 digit",
			args: args{
				v: "321654987",
			},
			want: want{
				output: "321654987",
			},
		},
		{
			name: "for more then 11 digit",
			args: args{
				v: "321654987321",
			},
			want: want{
				output: "321654987321",
			},
		},
		{
			name: "nil in input",
			args: args{
				v: nil,
			},
			want: want{
				output: "",
			},
		},
		{
			name: "without Country Code",
			args: args{
				v: "9876543210",
			},
			want: want{
				output: "19876543210",
			},
		},
		{
			name: "with country code",
			args: args{
				v: "19876543210",
			},
			want: want{
				output: "19876543210",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := E164Phone(tc.args.v)
			assert.Equal(t, tc.want.output, output)
		})
	}
}

//// Float642Int
func TestFloat642Int(t *testing.T) {
	type args struct {
		Value interface{}
	}
	type want struct {
		output  int
		varType string
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Float64 to Int",
			args: args{
				Value: 89.5,
			},
			want: want{
				output:  89,
				varType: "int",
			},
		},
		{
			name: "2nd test Float64 to Int",
			args: args{
				Value: 90.1,
			},
			want: want{
				output:  90,
				varType: "int",
			},
		},
		{
			name: "negative value",
			args: args{
				Value: -1540.90,
			},
			want: want{
				output:  -1540,
				varType: "int",
			},
		},
		{
			name: "blank string in input",
			args: args{
				Value: "",
			},
			want: want{
				output:  0,
				varType: "int",
			},
		},
		{
			name: "Pass 0",
			args: args{
				Value: 0,
			},
			want: want{
				output:  0,
				varType: "int",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := Float642Int(tc.args.Value)
			assert.Equal(t, tc.want.output, output)
			assert.Equal(t, fmt.Sprintf("%T", tc.want.output), fmt.Sprintf("%T", output))
		})
	}
}

func TestMarshal(t *testing.T) {

	type tempStruct struct {
		Name string `json:",omitempty"`
		Fees int    `json:",omitempty"`
	}

	type args struct {
		strct tempStruct
	}
	type want struct {
		output []byte
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "correct struct to json",
			args: args{
				strct: tempStruct{
					Name: "Raghav",
					Fees: 5000,
				},
			},
			want: want{
				output: []byte(`{"Name":"Raghav","Fees":5000}`),
			},
		},
		{
			name: "struct with omitempty json",
			args: args{
				strct: tempStruct{
					Name: "Tushar",
				},
			},
			want: want{
				output: []byte(`{"Name":"Tushar"}`),
			},
		},
		{
			name: "empty struct",
			args: args{
				strct: tempStruct{},
			},
			want: want{
				output: []byte(`{}`),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := Marshal(tc.args.strct)
			assert.Equal(t, tc.want.output, output)
		})
	}
}

func TestUnmarshal(t *testing.T) {

	type tempStruct struct {
		Name string
		Fees int
	}

	type args struct {
		jsn []byte
	}

	type want struct {
		output interface{}
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "correct json in byte form",
			args: args{
				jsn: []byte(`{"Name":"Raghav","Fees":5000}`),
			},
			want: want{
				output: &tempStruct{
					Name: "Raghav",
					Fees: 5000,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var v tempStruct
			output := Unmarshal(tc.args.jsn, &v)
			assert.Equal(t, tc.want.output, output)
		})
	}
}

func TestToString(t *testing.T) {
	type args struct {
		value interface{}
	}
	type want struct {
		output  string
		varType string
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "for nil input",
			args: args{
				value: nil,
			},
			want: want{
				output:  "",
				varType: "string",
			},
		},
		{
			name: "blank string",
			args: args{
				value: "",
			},
			want: want{
				output:  "",
				varType: "string",
			},
		},
		{
			name: "int to string",
			args: args{
				value: 9,
			},
			want: want{
				output:  "9",
				varType: "string",
			},
		},
		{
			name: "float value to string",
			args: args{
				value: 9.2,
			},
			want: want{
				output:  "9.2",
				varType: "string",
			},
		},
		{
			name: "struct to string",
			args: args{
				value: struct {
					Name string
					Fees int
				}{
					Name: "Hey",
					Fees: 123,
				},
			},
			want: want{
				output:  "{Hey 123}",
				varType: "string",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := ToString(tc.args.value)
			assert.Equal(t, tc.want.output, output)
			assert.Equal(t, tc.want.varType, fmt.Sprintf("%T", output))
		})
	}
}

func TestInt64(t *testing.T) {

	type args struct {
		input string
	}

	type want struct {
		output int64
	}
	tests := []struct {
		name    string // test case name
		args    args
		want    want // output we want from the test case
		wantErr bool // to handle test case which should return a error or not

	}{
		{
			name: "right string with int",
			args: args{
				input: "100",
			},
			want: want{
				output: 100,
			},
			wantErr: false,
		},
		{
			name: "blank string",
			args: args{
				input: "",
			},
			want: want{
				output: 0,
			},
			wantErr: false,
		},
		{
			name: "string with long int",
			args: args{
				input: "1234567890121316465",
			},
			want: want{
				output: 1234567890121316465,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int64(tt.args.input)
			assert.Equal(t, tt.want.output, got)
			assert.Equal(t, "int64", fmt.Sprintf("%T", got))
		})
	}
}

func TestJSON2Map(t *testing.T) {

	type args struct {
		rawMessage interface{}
	}
	type want struct {
		output map[string]interface{}
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "pass nil argument",
			args: args{
				rawMessage: nil,
			},
			want: want{
				output: nil,
			},
		},
		{
			name: "pass correct raw json",
			args: args{
				rawMessage: json.RawMessage([]byte(`{"Name" : "Hubspot","KD" : 12}`)),
			},
			want: want{
				output: map[string]interface{}{
					"Name": "Hubspot",
					"KD":   float64(12),
				},
			},
		},
		{
			name: "pass correct raw json second",
			args: args{
				rawMessage: json.RawMessage([]byte(`{"Name" : "Heymarket","Fees" : 123}`)),
			},
			want: want{
				output: map[string]interface{}{
					"Name": "Heymarket",
					"Fees": float64(123),
				},
			},
		},
		{
			name: "pass blank json with {}",
			args: args{
				rawMessage: json.RawMessage([]byte(`{}`)),
			},
			want: want{
				output: map[string]interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := JSON2Map(tc.args.rawMessage)
			assert.Equal(t, tc.want.output, output)
			assert.Equal(t, tc.want.output["KD"], output["KD"])
			assert.Equal(t, tc.want.output["Name"], output["Name"])
		})
	}
}

func TestUUID(t *testing.T) {

	testCases := []struct {
		name string
	}{
		{
			name: "should return random value that is not equal to nil",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.NotNil(t, UUID())
		})
	}
}

func TestReadAll(t *testing.T) {

	type args struct {
		req io.ReadCloser
	}
	type want struct {
		output []byte
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "simple input of io.ReadCloser type",
			args: args{
				req: NopCloser([]byte(`{"Name" : "Heymarket","Fees" : 123,}`)), // To get io.ReadCloser of jsonBodyForRequest
			},
			want: want{
				output: []byte(`{"Name" : "Heymarket","Fees" : 123,}`),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := ReadAll(tc.args.req)
			assert.Equal(t, tc.want.output, output)
		})
	}
}

func TestNopCloser(t *testing.T) {
	type args struct {
		body []byte
	}
	type want struct {
		output io.ReadCloser
	}
	body := []byte(`{"Name" : "Heymarket","Rating" : 5}`)
	outputBody := ioutil.NopCloser(bytes.NewBuffer(body)) /// for output assertion

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "JSON string in input",
			args: args{
				body: []byte(`{"Name" : "Heymarket","Rating" : 5}`),
			},
			want: want{
				output: outputBody,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := NopCloser(tc.args.body)
			assert.Equal(t, tc.want.output, output)
		})
	}
}

func TestToInt(t *testing.T) {
	type args struct {
		input interface{}
	}
	type want struct {
		output int
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "int in the type of interface",
			args: args{
				input: 45,
			},
			want: want{
				output: 45,
			},
		},
		{
			name: "nil in the type of interface",
			args: args{
				input: nil,
			},
			want: want{
				output: 0,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := ToInt(tc.args.input)
			assert.Equal(t, tc.want.output, output)
			assert.Equal(t, fmt.Sprintf("%T", tc.want.output), "int")
		})
	}
}

func TestFloat64(t *testing.T) {
	type args struct {
		v interface{}
	}
	type want struct {
		output float64
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "float64 type value in interface",
			args: args{
				v: 52.5,
			},
			want: want{
				output: 52.5,
			},
		},
		{
			name: "nil in the type of interface",
			args: args{
				v: nil,
			},
			want: want{
				output: 0,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := Float64(tc.args.v)
			assert.Equal(t, tc.want.output, output)
			assert.Equal(t, fmt.Sprintf("%T", tc.want.output), "float64")
			assert.NotNil(t, output)
		})
	}
}

func TestConvertMap(t *testing.T) {
	type args struct {
		v interface{}
	}
	type want struct {
		output     map[string]interface{}
		outputType string
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "pass an interface value which holding values in map",
			args: args{
				v: interface{}(map[string]interface{}{
					"Name":   "Heymarket",
					"Rating": 4.5,
				}),
			},
			want: want{
				output: map[string]interface{}{
					"Name":   "Heymarket",
					"Rating": 4.5,
				},
				outputType: "map[string]interface {}",
			},
		},
		{
			name: "for blank interface in input",
			args: args{
				v: interface{}(nil),
			},
			want: want{
				output:     nil,
				outputType: "map[string]interface {}",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := ConvertMap(tc.args.v)
			assert.Equal(t, tc.want.output, output)
			assert.Equal(t, tc.want.outputType, fmt.Sprintf("%T", output))
		})
	}
}
