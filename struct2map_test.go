package struct2map

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestJsonTagNameParser(t *testing.T) {
	type A struct {
		AuthKey     string `protobuf:"bytes,1,opt,name=authKey,proto3" json:"authKey,omitempty"`
		AccountName string `protobuf:"bytes,2,opt,name=accountName,proto3" json:"accountName,omitempty"`
		Age         string `json:"age"`
	}
	a := A{}
	at := reflect.TypeOf(a)

	tests := []struct {
		field string
		want  string
	}{
		{"AuthKey", "authKey"},
		{"AccountName", "accountName"},
		{"Age", "age"},
	}

	for _, tt := range tests {
		ats, ok := at.FieldByName(tt.field)
		assert.Equal(t, true, ok)
		tagVal, ok := ats.Tag.Lookup("json")
		assert.Equal(t, true, ok)
		name := JsonTagNameParser(tagVal)
		assert.Equal(t, tt.want, name)
	}

}

func TestProtobufTagNameParser(t *testing.T) {
	type A struct {
		AuthKey     string `protobuf:"bytes,1,opt,name=authKey,proto3" json:"authKey,omitempty"`
		AccountName string `protobuf:"bytes,2,opt,name=accountName,proto3" json:"accountName,omitempty"`
		Age         string `protobuf:"bytes,3,opt,name=age,proto3"`
	}
	a := A{}
	at := reflect.TypeOf(a)

	tests := []struct {
		field string
		want  string
	}{
		{"AuthKey", "authKey"},
		{"AccountName", "accountName"},
		{"Age", "age"},
	}

	for _, tt := range tests {
		ats, ok := at.FieldByName(tt.field)
		assert.Equal(t, true, ok)
		tagVal, ok := ats.Tag.Lookup("protobuf")
		assert.Equal(t, true, ok)
		name := ProtobufTagNameParser(tagVal)
		assert.Equal(t, tt.want, name)
	}
}

func TestStruct2Map(t *testing.T) {
	type A struct {
		AuthKey     string `protobuf:"bytes,1,opt,name=authKey,proto3" json:"authKey,omitempty"`
		AccountName string `protobuf:"bytes,2,opt,name=accountName,proto3" json:"accountName,omitempty"`
		Age         int    `protobuf:"bytes,3,opt,name=age,proto3" json:"age,omitempty"`
	}

	tests := []struct {
		name          string
		inStruct      interface{}
		tag           string
		tagNameParser func(string) string
		want          map[string]interface{}
	}{
		{
			"json tag struct",
			A{AuthKey: "d1234", AccountName: "", Age: 7},
			"json",
			JsonTagNameParser,
			map[string]interface{}{
				"authKey": "d1234",
				"age":     7,
			},
		},
		{
			"protobuf tag struct",
			&A{AuthKey: "d1234", AccountName: "", Age: 7},
			"protobuf",
			ProtobufTagNameParser,
			map[string]interface{}{
				"authKey": "d1234",
				"age":     7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := Struct2Map(tt.inStruct, tt.tag, tt.tagNameParser, true)
			assert.Equal(t, tt.want, ret)
		})
	}
}
