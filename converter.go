package jyxconverter

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
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

//MarshalXML is a method used to Marshall a Map to xml
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

type xmlEntry struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:"-"`
	Content string     `xml:",innerxml"`
	Nodes   []xmlEntry `xml:",any"`
	Value   string     `xml:",chardata"`
}

func decodeElement(entry xmlEntry) interface{} {
	//fmt.Println("Decoding entry:", entry)
	if entry.Nodes == nil {
		return entry.Value
	}
	m := Map{}
	for _, node := range entry.Nodes {
		m[node.XMLName.Local] = decodeElement(node)
	}
	return m
}

//UnmarshalXML is a method used to Unmarshal a xml to a Map
func (m *Map) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = Map{}
	e := xmlEntry{}
	var err error
	for err = d.Decode(&e); err == nil; err = d.Decode(&e) {
		(*m)[e.XMLName.Local] = decodeElement(e)
		e = xmlEntry{}
	}
	if err != nil && err != io.EOF {
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
	if err != nil {
		return nil, err
	}

	result := convertInterface(yamlMap)
	return json.MarshalIndent(&result, "", "\t")
}

//YamlToXML converts yaml to xml
func YamlToXML(bytes []byte) ([]byte, error) {

	var yamlMap map[interface{}]interface{}
	err := yaml.Unmarshal(bytes, &yamlMap)
	if err != nil {
		return nil, err
	}

	result := convertInterface(yamlMap)
	return xml.MarshalIndent(&result, "", "\t")
}

//XMLToJSON converts xml to json
func XMLToJSON(bytes []byte) ([]byte, error) {

	var xmlMap Map
	err := xml.Unmarshal(bytes, &xmlMap)
	if err != nil {
		return nil, err
	}

	result := convertInterface(xmlMap)

	return json.MarshalIndent(&result, "", "\t")
}

//XMLToYaml converts xml to yaml
func XMLToYaml(bytes []byte) ([]byte, error) {

	var xmlMap Map
	err := xml.Unmarshal(bytes, &xmlMap)

	if err != nil {
		return nil, err
	}

	result := convertInterface(xmlMap)

	return yaml.Marshal(&result)
}
