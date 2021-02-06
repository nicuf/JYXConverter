jyxconverter
============

Package jyxconverter provides functions to convert yaml, xml and json
between them

* xml <-> json
* xml <-> yaml
* json <-> yaml

i.e. you have a string containing json data and you want to convert it to
xml or yaml

```go
jsonString = "{
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
        \"SocialPages\": {
            \"facebook\": \"fbJhon\",
            \"linkedin\": \"lkdnJhon\"
    }"
```

in order to obtain a xml from it use

```go
xmlBytes, err := JSONToXML([]byte(jsonString))
```

or if yaml is needed then use

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

and the result would be for xml:

```xml
<Person>
    <Name>Jhon</Name>
    <Surname>Dhoe</Surname>
    <Gender>0</Gender>
    <Maried>true</Maried>
    <SocialPages>
        <facebook>fbJhon</facebook>
        <linkedin>lkdnJhon</linkedin>
    </SocialPages>
    <Skills>
            <SkillsElement>programming</SkillsElement>
            <SkillsElement>gardening</SkillsElement>
            <SkillsElement>communication</SkillsElement>
    </Skills>
</Person>
```

and for yaml

```yaml
Person:
  name: Jhon
  surname: Dhoe
  gender: 0
  maried: true
  skills:
    - programming
    - gardening
    - communication
  socialpages:
    facebook: fbJhon
    linkedin: lkdnJhon
```
