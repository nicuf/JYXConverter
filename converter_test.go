package jyxconverter

import (
	"io/ioutil"
	"testing"
)

func TestJSONToXMLAndYaml(t *testing.T) {
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

	t.Log("yaml from json is:")
	t.Log(string(yamlBytes))
}

func TestYamlToXMLAndJSON(t *testing.T) {
	yamlBytes, err := ioutil.ReadFile("./test_files/Person.yaml")

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log("yaml from file is:")
	t.Log(string(yamlBytes))
	xmlBytes, err := YamlToXML(yamlBytes)

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log("xml from yaml is:")
	t.Log(string(xmlBytes))

	jsonBytes, err := YamlToJSON(yamlBytes)

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log("json from yaml is:")
	t.Log(string(jsonBytes))
}
