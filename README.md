JYXconverter
============

Package jyxconverter provides functions to convert yaml, xml and json
between them

* xml <-> json
* xml <-> yaml
* json <-> yaml

i.e. you have a string containing json data and you want to convert it to
xml or yaml

    jsonString = {"oneKey":"oneValue"}

it just have to call

    xmlBytes, err := JSONToXML([]byte(jsonString))

or

    yamlBytes, err := JSONToYaml([]byte(jsonString))

if there is no error, you can make a string from the byte array

    xmlString := string(xmlBytes)

or

    yamlString := string(yamlBytes)

and the result should be like:

    <oneKey>oneValue</onekey>

or

    oneKey: oneValue
