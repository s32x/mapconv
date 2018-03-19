package mapconv

import (
	"fmt"
	"log"
	"testing"
	"time"
)

type TestStruct struct {
	StringTest   string `json:"string_test"`
	IntTest      int    `json:"int_test"`
	NestedStruct Data   `json:"nested_struct"`
}

type Data struct {
	TimeTest    time.Time `json:"time_test"`
	Float64Test float64   `json:"float_64_test"`
}

func TestMapConv(t *testing.T) {
	myData := make(map[string]interface{})
	myData["StringTest"] = "hello world"
	myData["IntTest"] = 1234

	moreData := make(map[string]interface{})
	moreData["TimeTest"] = time.Now()
	moreData["Float64Test"] = 20.1234
	myData["NestedStruct"] = moreData

	var result TestStruct
	if err := Mtos(myData, &result); err != nil {
		fmt.Println(err)
	}
	log.Println(result)
}

func TestMapConvJSON(t *testing.T) {
	myData := make(map[string]interface{})
	myData["string_test"] = "hello world"
	myData["int_test"] = 1234

	moreData := make(map[string]interface{})
	moreData["time_test"] = time.Now()
	moreData["float_64_test"] = 20.1234
	myData["nested_struct"] = moreData

	var result TestStruct
	if err := MtosJSON(myData, &result); err != nil {
		fmt.Println(err)
	}
	log.Println(result)
}

type TestOmitemptyStruct struct {
	StringTest   string        `json:"string_test,omitempty"`
	IntTest      int           `json:"int_test,omitempty"`
	NestedStruct OmitemptyData `json:"nested_struct,omitempty"`
}

type OmitemptyData struct {
	TimeTest    time.Time `json:"time_test,omitempty"`
	Float64Test float64   `json:"float_64_test,omitempty"`
}

func TestMapConvJSONWithOmitempty(t *testing.T) {
	myData := make(map[string]interface{})
	myData["string_test"] = "hello world"
	myData["int_test"] = 1234

	moreData := make(map[string]interface{})
	moreData["time_test"] = time.Now()
	moreData["float_64_test"] = 20.1234
	myData["nested_struct"] = moreData

	var result TestOmitemptyStruct
	if err := MtosJSON(myData, &result); err != nil {
		fmt.Println(err)
	}
	log.Println(result)
}
