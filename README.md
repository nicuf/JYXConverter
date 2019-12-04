JYXconverter
============

Package jyxconverter provides functions to convert yaml, xml and json
between them

* xml <-> json
* xml <-> yaml
* json <-> yaml

i.e. you have a string containing json data and you want to convert it to
xml or yaml
```go
jsonString = "{\"oneKey\":\"oneValue\"}"
```
it just have to call
```go
xmlBytes, err := JSONToXML([]byte(jsonString))
```
or
```go
yamlBytes, err := JSONToYaml([]byte(jsonString))
```
if there is no error, you can make a string from the byte array
```go
xmlString := string(xmlBytes)
```
or
```go
yamlString := string(yamlBytes)
```
and the result should be like:
```xml
<oneKey>oneValue</onekey>
```
or
```yaml
oneKey: oneValue
```