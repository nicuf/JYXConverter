package jyxconverter

import (
	"fmt"
	"io/ioutil"
	"testing"
)

const jsonTestString string = `
{ 
	\"Person\": {
		\"Name\": \"Jhon\",
		\"Surname\": \"Dhoe\",
		\"Gender\": 0,
		\"Maried\": true,
		\"Skills\": [
			\"programming\",
			\"gardening\",
			\"communication\"
		],
		\"Address\": {
			\"City\": \"Wien\",
			\"HouseNmbr\": \"34\",
			\"ApartmentNmbr\": \"90\",
			\"Province\": {
				\"Country\": \"Austria\",
				\"County\": \"WienCounty\"
			}
		},
		\"SocialPages\": {
			\"facebook\": \"fbJhon\",
			\"linkedin\": \"lkdnJhon\"
		}
	}
}`

const yamlTestString string = `
Person:
	name: Jhon
	surname: Dhoe
	gender: 0
	maried: true
	skills:
		- programming
		- gardening
		- communication
	address:
		city: Wien
		housenmbr: "34"
		apartmentnmbr: "90"
		province: 
			country: Austria
			county: WienCounty
	socialpages:
		facebook: fbJhon
		linkedin: lkdnJhon
`

const xmlTestString = `
<Person>
	<Name>Jhon</Name>
	<Surname>Dhoe</Surname>
	<Gender>0</Gender>
	<Maried>true</Maried>
	<Address>
		<City>Wien</City>
		<HouseNmbr>34</HouseNmbr>
		<ApartmentNmbr>90</ApartmentNmbr>
		<Province>
			<Country>Austria</Country>
			<County>WienCounty</County>
		</Province>
	</Address>
	<SocialPages>
		<facebook>fbJhon</facebook>
		<linkedin>lkdnJhon</linkedin>
	</SocialPages>
	<Skills>
			<SkillsElement>programming</SkillsElement>
			<SkillsElement>gardening</SkillsElement>
			<SkillsElement>communication</SkillsElement>
	</Skills>
</Person>`

var result string

func BenchmarkJSONToXML(b *testing.B) {
	var r []byte
	jsonBytes := []byte(jsonTestString)
	for n := 0; n < b.N; n++ {
		r, _ = JSONToXML(jsonBytes)
	}
	result = string(r)
}

func BenchmarkJSONToYaml(b *testing.B) {
	var r []byte
	jsonBytes := []byte(jsonTestString)
	for n := 0; n < b.N; n++ {
		r, _ = JSONToYaml(jsonBytes)
	}
	result = string(r)
}

func BenchmarkYamlToXML(b *testing.B) {
	var r []byte
	yamlBytes := []byte(yamlTestString)
	for n := 0; n < b.N; n++ {
		r, _ = YamlToXML(yamlBytes)
	}
	result = string(r)
}

func BenchmarkYamlToJSON(b *testing.B) {
	var r []byte
	yamlBytes := []byte(yamlTestString)
	for n := 0; n < b.N; n++ {
		r, _ = YamlToXML(yamlBytes)
	}
	result = string(r)
}

func BenchmarkXMLToJSON(b *testing.B) {
	var r []byte
	xmlBytes := []byte(xmlTestString)
	for n := 0; n < b.N; n++ {
		r, _ = XMLToJSON(xmlBytes)
	}
	result = string(r)
}

func BenchmarkXMLToYaml(b *testing.B) {
	var r []byte
	xmlBytes := []byte(xmlTestString)
	for n := 0; n < b.N; n++ {
		r, _ = XMLToYaml(xmlBytes)
	}
	result = string(r)
}

func TestJSONToXMLAndYamlPositiveFlow(t *testing.T) {
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

func TestYamlToXMLAndJSONPositiveFlow(t *testing.T) {
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

func TestXMLToYamlAndJSONPositiveFlow(t *testing.T) {
	xmlBytes, err := ioutil.ReadFile("./test_files/Person.xml")

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log("xml from file is:")
	t.Log(string(xmlBytes))
	yamlBytes, err := XMLToYaml(xmlBytes)

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log("yaml from xml is:")
	t.Log(string(yamlBytes))

	jsonBytes, err := XMLToJSON(xmlBytes)

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log("json from xml is:")
	t.Log(string(jsonBytes))
}

func TestEmptyXMLToJsonShouldReturnError(t *testing.T) {
	xmlBytes := []byte{}
	_, err := XMLToJSON(xmlBytes)

	if err == nil {
		t.Errorf("Trying to convert empty string from xml to json should return an error!")
	}
}

func TestEmptyYamlToJsonShouldReturnErrorOrEmptyJSON(t *testing.T) {
	yamlBytes := []byte{}
	jsonBytes, err := YamlToJSON(yamlBytes)

	if err == nil && len(jsonBytes) > 2 {
		t.Errorf("Trying to convert empty string from yaml to json should return an error!")
	}
}

func TestEmptyJsonToXMLShouldReturnError(t *testing.T) {
	jsonBytes := []byte{}
	_, err := XMLToJSON(jsonBytes)

	if err == nil {
		t.Errorf("Trying to convert empty string from json to xml should return an error!")
	}
}

func ExampleJSONToXML() {
	jsonTestString := `
{ 
	"Person": {
		"Name": "Jhon",
		"Surname": "Dhoe",
		"Gender": 0,
		"Maried": true
	}
}`
	xmlBytes, err := JSONToXML([]byte(jsonTestString))

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(xmlBytes))
	//Output:
	//<Person>
	//	<Name>Jhon</Name>
	//	<Surname>Dhoe</Surname>
	//	<Gender>0.000000</Gender>
	//	<Maried>true</Maried>
	//</Person>
}

func ExampleJSONToYaml() {
	jsonTestString := `
{ 
	"Person": {
		"Name": "Jhon",
		"Surname": "Dhoe",
		"Gender": 0,
		"Maried": true
	}
}`
	yamlBytes, err := JSONToYaml([]byte(jsonTestString))

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(yamlBytes))
	//Output:
	//Person:
	//	Name: Jhon
	//	Surname: Dhoe
	//	Gender: 0
	//	Maried: true
}

func ExampleYamlToJSON() {
	yamlTestString := `
Person:
  Name: Jhon
  Surname: Dhoe
  Gender: 0
  Maried: true
`
	jsonBytes, err := YamlToJSON([]byte(yamlTestString))

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonBytes))
	//Output:
	//{
	//	"Person": {
	//		"Name": "Jhon",
	//		"Surname": "Dhoe",
	//		"Gender": 0,
	//		"Maried": true
	//	}
	//}
}

func ExampleYamlToXML() {
	yamlTestString := `
Person:
  Name: Jhon
  Surname: Dhoe
  Gender: 0
  Maried: true
`
	xmlBytes, err := YamlToXML([]byte(yamlTestString))

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(xmlBytes))
	//Output:
	//<Person>
	//	<Name>Jhon</Name>
	//	<Surname>Dhoe</Surname>
	//	<Gender>0.000000</Gender>
	//	<Maried>true</Maried>
	//</Person>
}

func ExampleXMLToYaml() {
	xmlTestString := `
	<Person>
		<Name>Jhon</Name>
		<Surname>Dhoe</Surname>
		<Gender>0</Gender>
		<Maried>true</Maried>
	</Person>
	`

	yamlBytes, err := XMLToYaml([]byte(xmlTestString))

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(yamlBytes))
	//Output:
	//Person:
	//	Name: Jhon
	//	Surname: Dhoe
	//	Gender: 0
	//	Maried: true
}

func ExampleXMLToJSON() {
	xmlTestString := `
	<Person>
		<Name>Jhon</Name>
		<Surname>Dhoe</Surname>
		<Gender>0</Gender>
		<Maried>true</Maried>
	</Person>
	`

	jsonBytes, err := XMLToJSON([]byte(xmlTestString))

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonBytes))
	//Output:
	//{
	//	"Person": {
	//		"Name": "Jhon",
	//		"Surname": "Dhoe",
	//		"Gender": 0,
	//		"Maried": true
	//	}
	//}
}
