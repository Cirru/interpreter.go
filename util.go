
package cirruGopher

import (
  "fmt"
)

func stringifyObject(data cirruObject) string {
  switch data.Typing {
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
    case "array": return "::array::"
    case "map": return "::map::"
  }
  return ""
}