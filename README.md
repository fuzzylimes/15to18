## 15to18
CLI tool for converting between Salesforce 15 and 18 Ids.

### Usage
Either provide your SF Ids in a string in the CLI command, or pass in the path to a file. The tool will scan your input for any SF15 ID and convert them to 18 IDs. No need to sanitize your data first - the tool will find all instances and print out the conversion for you.

#### Via CLI
You can convert any ids by including them as command line arguments

```
> 15to18 a124t000000KDCA

< a124t000000KDCA -> a124t000000KDCAAA4

> 15to18 a124t000000KDC 001i000001AWbWu, 5003000000D8cuI

< a124t000000KDCA -> a124t000000KDCAAA4
  001i000001AWbWu -> 001i000001AWbWuAAL
  5003000000D8cuI -> 5003000000D8cuIAAD
```

#### Via File
To convert all 15 digit SF Ids in a text file, use the `-file` parameter

```
> 15to18 -file test/sample.json

< a124t000000KDCA -> a124t000000KDCAAA4
  001i000001AWbWu -> 001i000001AWbWuAAL
  5003000000D8cuI -> 5003000000D8cuIAAR
```