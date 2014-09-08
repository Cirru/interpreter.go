
package interpreter

import (
  "github.com/Cirru/parser"
  "github.com/Cirru/writer"
  "fmt"
)

func stringifyunitype(data unitype) string {
  tree := transformunitype(data)
  lines := []interface{}{tree}
  return writer.MakeCode(lines)
}

func transformunitype(data unitype) []interface{} {
  switch data.Type {
    case uniString:
      if stringValue, ok := data.Value.(string); ok {
        return []interface{}{"string", stringValue}
      }
    case uniInt:
      if value, ok := data.Value.(int); ok {
        str := fmt.Sprintf("%d", value)
        return []interface{}{"int", str}
      }
    case uniFloat:
      if value, ok := data.Value.(float64); ok {
        str := fmt.Sprintf("%g", value)
        return []interface{}{"float", str}
      }
    case uniBool:
      if value, ok := data.Value.(bool); ok {
        if value {
          return []interface{}{"bool", "true"}
        }
        return []interface{}{"bool", "false"}
      }
    case uniArray:
      list := []interface{}{"array"}
      if value, ok := data.Value.(*map[unitype]unitype); ok {
        for _, item := range *value {
          list = append(list, transformunitype(item))
        }
      }
      return list
    case uniTable:
      list := []interface{}{"table"}
      if value, ok := data.Value.(*scope); ok {
        for k, v := range *value {
          pair := []interface{}{k, transformunitype(v)}
          list = append(list, pair)
        }
      }
      return list
    case uniRegexp:
      str := fmt.Sprintf("%s", data.Value)
      return []interface{}{"regexp", str}
    case uniFn:
      if fnContext, ok := data.Value.(context); ok {
        args := transformCode(fnContext.args)
        code := transformCode(fnContext.code)
        return []interface{}{"fn", args, code}
      }
    case uniMacro:
      if macContext, ok := data.Value.(context); ok {
        args := transformCode(macContext.args)
        code := transformCode(macContext.code)
        return []interface{}{"macro", args, code}
      }
    case uniNil:
      return []interface{}{"nil"}
    default:
      panic("unknown structure")
  }
  return []interface{}{}
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

func uni(x interface{}) (ret unitype) {
  switch value := x.(type) {
  case int:
    ret = unitype{uniInt, value}
  case string:
    ret = unitype{uniString, value}
  case *scope:
    ret = unitype{uniTable, value}
  default: panic("not implemented")
  }
  return
}