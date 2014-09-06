
package interpreter

import (
  "github.com/Cirru/parser"
  "fmt"
  "encoding/json"
  "strings"
)

func repeatBlank(n int) (ret string) {
  ret = ""
  for i := 1; i <= n; i++ {
    ret += "  "
  }
  return
}

func stringifyObject(data Object, level int) string {
  switch data.Tag {
    case cirruTypeString:
      if stringValue, ok := data.Value.(string); ok {
        return "string \"" + stringValue + "\""
      }
    case cirruTypeInt:
      if intValue, ok := data.Value.(int); ok {
        return "int " + fmt.Sprintf("%d", intValue)
      }
    case cirruTypeFloat:
      if floatValue, ok := data.Value.(float64); ok {
        return "float " + fmt.Sprintf("%g", floatValue)
      }
    case cirruTypeBool:
      if value, ok := data.Value.(bool); ok {
        if value {
          return "bool true"
        }
        return "bool false"
      }
    case cirruTypeArray:
      list := []string{}
      indent := "\n" + repeatBlank(level + 1)
      if anArray, ok := data.Value.(*[]Object); ok {
        for _, item := range *anArray {
          list = append(list, stringifyObject(item, (level + 1)))
        }
      }
      stringValue := strings.Join(list, indent)
      return "array " + indent + stringValue
    case cirruTypeMap:
      list := []string{}
      indent := "\n" + repeatBlank(level + 1)
      // debugPrint("string is", data.Value)
      if aMap, ok := data.Value.(*Env); ok {
        for key, value := range *aMap {
          hold := "\"" + key + "\" $ " + stringifyObject(value, (level + 1))
          list = append(list, hold)
        }
      }
      stringValue := strings.Join(list, indent)
      return "map " + indent + stringValue
    case cirruTypeRegexp:
      return "regexp " + fmt.Sprintf("%s", data.Value)
    case cirruTypeCode:
      if code, ok := data.Value.(*[]interface{}); ok {
        return "code " + codeString(*code, level)
      }
    default: return "unknown"
  }
  return "nil"
}

func debugPrint(xs ...interface{}) {
  list := []interface{}{}
  for _, item := range xs {
    json, err := json.MarshalIndent(item, "", "  ")
    if err != nil {
      panic(err)
    }
    list = append(list, interface{}(string(json)))
  }
  fmt.Println("")
  fmt.Println(xs)
  fmt.Println(list...)
}

func generateString(x string) (ret Object) {
  ret.Tag = cirruTypeString
  ret.Value = x
  return
}

func generateMap(m *Env) (ret Object) {
  ret.Tag = cirruTypeMap
  ret.Value = m
  return
}

func codeString(xs []interface{}, level int) (ret string) {
  hold := []string{}
  indent := "\n" + repeatBlank(level + 1)
  for _, item := range xs {
    if buffer, ok := item.(parser.Token); ok {
      hold = append(hold, buffer.Text)
    }
    if list, ok := item.([]interface{}); ok {
      tmpString := indent + codeString(list, (level + 1))
      hold = append(hold, tmpString)
    }
  }
  ret = strings.Join(hold, " ")
  return
}

func stop(text ...interface{}) {
  fmt.Println(text...)
  panic("")
}