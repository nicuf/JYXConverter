package jyxconverter

import (
	"io/ioutil"
	"testing"
)

func Test(t *testing.T) {
	jsonBytes, err := ioutil.ReadFile("./test_files/Person.json")

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log("json from file is:")
	t.Log(string(jsonBytes))
	xmlBytes, err := JSONToXML(jsonBytes)

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log("xml from json is:")
	t.Log(string(xmlBytes))

	yamlBytes, err := JSONToYaml(jsonBytes)

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log("yamlBytes from json is:")
	t.Log(string(yamlBytes))
}
