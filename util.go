
package cirruGopher

import (
  "fmt"
  "encoding/json"
)

func stringifyObject(data Object) string {
  if &data.Tag == nil {
    return "nil"
  }
  switch data.Tag {
    case "string":
      if stringValue, ok := data.Value.(string); ok {
        return stringValue
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
    case "array": return "::array::"
    case "map": return "::map::"
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
