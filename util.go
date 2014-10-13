
package interpreter

import (
  "github.com/Cirru/writer"
  "fmt"
)

func stringifyUnitype(data unitype) string {
  tree := showUnitype(data)
  if raw, ok := tree.(string); ok {
    return raw
  } else {
    lines := []interface{}{tree}
    return writer.MakeCode(lines)
  }
}

func showUnitype(data unitype) interface{} {
  switch data.Type {
    case uniString:
      if stringValue, ok := data.Value.(string); ok {
        return fmt.Sprintf(":%s", stringValue)
      }
    case uniFloat:
      if value, ok := data.Value.(float64); ok {
        return fmt.Sprintf("%g", value)
      }
    case uniBool:
      if value, ok := data.Value.(bool); ok {
        if value {
          return "#true"
        } else {
          return "#false"
        }
      }
    case uniArray:
      list := []interface{}{"array"}
      value, _ := data.Value.(*mapping)
      for _, item := range *value {
        list = append(list, showUnitype(item))
      }
      return list
    case uniTable:
      list := []interface{}{"table"}
      if value, ok := data.Value.(*mapping); ok {
        for k, v := range *value {
          pair := []interface{}{showUnitype(k), showUnitype(v)}
          list = append(list, pair)
        }
      }
      return list
    case uniRegexp:
      return fmt.Sprintf("/%s", data.Value)
    case uniFn:
      if fnContext, ok := data.Value.(context); ok {
        args := transformCode(fnContext.args)
        code := transformCode(fnContext.code)
        return []interface{}{"fn", args, code}
      }
    case uniNil:
      return "@nil"
    default:
      panic("unknown structure")
  }
  return []interface{}{}
}

func transformCode(xs []interface{}) []interface{} {
  hold := []interface{}{}
  for _, item := range xs {
    if buffer, ok := item.(token); ok {
      hold = append(hold, buffer.Text)
    } else if list, ok := item.(sequence); ok {
      hold = append(hold, transformCode(list))
    } else {
      panic("cannot handle code item")
    }
  }
  return hold
}

func uni(x interface{}) (ret unitype) {
  if x == nil {
    return unitype{uniNil, nil}
  }
  switch value := x.(type) {
  case float64:
    ret = unitype{uniFloat, value}
  case string:
    ret = unitype{uniString, value}
  case bool:
    ret = unitype{uniBool, value}
  case *mapping:
    ret = unitype{uniTable, value}
  default: panic("not implemented in uni")
  }
  return
}

func newFileScope() *scope {
  obj := &object{}
  scp := &scope{}
  scp.closure = obj
  return scp
}

func newScope(parent *scope) *scope {
  obj := &object{}
  return &scope{parent, obj}
}
