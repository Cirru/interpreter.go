
package cirruGopher

import (
  "fmt"
  "encoding/json"
  "strings"
)

func stringifyObject(data Object) string {
  if &data.Tag == nil {
    return "nil"
  }
  switch data.Tag {
    case "string":
      if stringValue, ok := data.Value.(string); ok {
        return "\"" + stringValue + "\""
      }
    case "int":
      if intValue, ok := data.Value.(int); ok {
        return fmt.Sprintf("%d", intValue)
      }
    case "float":
      if floatValue, ok := data.Value.(float64); ok {
        return fmt.Sprintf("%g", floatValue)
      }
    case "bool":
      if value, ok := data.Value.(bool); ok {
        if value {
          return "#t"
        }
        return "#f"
      }
    case "array":
      list := []string{}
      if anArray, ok := data.Value.([]Object); ok {
        for _, item := range anArray {
          list = append(list, stringifyObject(item))
        }
      }
      stringValue := strings.Join(list, ", ")
      return "[" + stringValue + "]"
    case "map":
      list := []string{}
      // debugPrint("string is", data.Value)
      if aMap, ok := data.Value.(*Env); ok {
        for key, value := range *aMap {
          hold := "\"" + key + "\": " + stringifyObject(value)
          list = append(list, hold)
        }
      }
      stringValue := strings.Join(list, ", ")
      return "{" + stringValue + "}"
    default: return "<unknown>"
  }
  return ""
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
  ret.Tag = "string"
  ret.Value = x
  return
}

func generateMap(m *Env) (ret Object) {
  ret.Tag = "map"
  ret.Value = m
  return
}
