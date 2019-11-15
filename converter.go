package jyxconverter

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"

	"gopkg.in/yaml.v2"
)

//Map is a interface used for xml marshaling
type Map map[string]interface{}

func getElementTokens(element interface{}, name string) []xml.Token {
	tokens := []xml.Token{}

	startElement := xml.StartElement{Name: xml.Name{"", name}}
	tokens = append(tokens, startElement)

	switch valueType := element.(type) {
	case int:
		tokens = append(tokens, xml.CharData(string(valueType)))
	case float64:
		tokens = append(tokens, xml.CharData(strconv.FormatFloat(valueType, 'f', 6, 64)))
	case string:
		tokens = append(tokens, xml.CharData(valueType))
	case bool:
		tokens = append(tokens, xml.CharData(strconv.FormatBool(valueType)))
	case []interface{}:
		var innerSlice []interface{}
		innerSlice, ok := element.([]interface{})
		if !ok {
			panic("value is not a Slice")
		}
		for _, element := range innerSlice {
			tokens = append(tokens, getElementTokens(element, name+"Element")...)
		}
	case Map:
		tokens = append(tokens, getMapTokens(valueType)...)
	case map[string]interface{}:
		tokens = append(tokens, getMapTokens(valueType)...)
	default:
		fmt.Println("default case")
		fmt.Printf("type is %T\n value is %v", element, element)
	}

	tokens = append(tokens, xml.EndElement{startElement.Name})

	return tokens
}

func getMapTokens(m Map) []xml.Token {
	tokens := []xml.Token{}
	for key, value := range m {
		tokens = append(tokens, getElementTokens(value, key)...)
	}

	return tokens
}

func convertInterface(object interface{}) interface{} {
	result := object
	switch valueType := object.(type) {
	case map[interface{}]interface{}:
		mapResult := Map{}
		for key, value := range valueType {
			mapResult[fmt.Sprint(key)] = convertInterface(value)
		}
		return mapResult
	case []interface{}:
		sliceResult := []interface{}{}
		for _, value := range valueType {
			sliceResult = append(sliceResult, convertInterface(value))
		}
		return sliceResult
	default:
		return result
	}
}

//MarshalXML is a method used to Marshall a MAp to xml
func (m Map) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	tokens := []xml.Token{start}
	tokens = append(tokens, getMapTokens(m)...)
	tokens = append(tokens, xml.EndElement{start.Name})

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	err := e.Flush()
	if err != nil {
		return err
	}

	return nil
}

//JSONToYaml converts json into yaml
func JSONToYaml(bytes []byte) ([]byte, error) {

	var result Map
	err := json.Unmarshal(bytes, &result)

	if err != nil {
		return nil, err
	}

	return yaml.Marshal(&result)
}

//JSONToXML converts json to xml
func JSONToXML(bytes []byte) ([]byte, error) {

	var result Map
	err := json.Unmarshal(bytes, &result)

	if err != nil {
		return nil, err
	}

	return xml.MarshalIndent(&result, "", "\t")
}

//YamlToJSON converts yaml to json
func YamlToJSON(bytes []byte) ([]byte, error) {

	var yamlMap map[interface{}]interface{}
	err := yaml.Unmarshal(bytes, &yamlMap)
	result := convertInterface(yamlMap)
	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(&result, "", "\t")
}

//YamlToXML converts yaml to xml
func YamlToXML(bytes []byte) ([]byte, error) {

	var yamlMap map[interface{}]interface{}
	err := yaml.Unmarshal(bytes, &yamlMap)
	result := convertInterface(yamlMap)

	if err != nil {
		return nil, err
	}

	return xml.MarshalIndent(&result, "", "\t")
}
