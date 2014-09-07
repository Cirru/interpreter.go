
package interpreter

import (
  "github.com/Cirru/parser"
  "github.com/Cirru/writer"
  "fmt"
)

func repeatBlank(n int) (ret string) {
  ret = ""
  for i := 1; i <= n; i++ {
    ret += "  "
  }
  return
}

func stringifyObject(data Object) string {
  tree := transformObject(data)
  lines := []interface{}{tree}
  return writer.MakeCode(lines)
}

func transformObject(data Object) []interface{} {
  switch data.Tag {
    case cirruTypeString:
      if stringValue, ok := data.Value.(string); ok {
        return []interface{}{"string", stringValue}
      }
    case cirruTypeInt:
      if value, ok := data.Value.(int); ok {
        str := fmt.Sprintf("%d", value)
        return []interface{}{"int", str}
      }
    case cirruTypeFloat:
      if value, ok := data.Value.(float64); ok {
        str := fmt.Sprintf("%g", value)
        return []interface{}{"float", str}
      }
    case cirruTypeBool:
      if value, ok := data.Value.(bool); ok {
        if value {
          return []interface{}{"bool", "true"}
        }
        return []interface{}{"bool", "false"}
      }
    case cirruTypeArray:
      list := []interface{}{"array"}
      if value, ok := data.Value.(*[]Object); ok {
        for _, item := range *value {
          list = append(list, transformObject(item))
        }
      }
      return list
    case cirruTypeMap:
      list := []interface{}{"map"}
      if value, ok := data.Value.(*Env); ok {
        for k, v := range *value {
          pair := []interface{}{k, transformObject(v)}
          list = append(list, pair)
        }
      }
      return list
    case cirruTypeRegexp:
      str := fmt.Sprintf("%s", data.Value)
      return []interface{}{"regexp", str}
    case cirruTypeCode:
      if value, ok := data.Value.(*[]interface{}); ok {
        return []interface{}{"code", transformCode(*value)}
      }
    default:
      panic("unknown structure")
  }
  return []interface{}{}
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

func transformCode(xs []interface{}) []interface{} {
  hold := []interface{}{}
  for _, item := range xs {
    if buffer, ok := item.(parser.Token); ok {
      hold = append(hold, buffer.Text)
    }
    if list, ok := item.([]interface{}); ok {
      hold = append(hold, transformCode(list))
    }
  }
  return hold
}
