
package cirruGopher

import (
  "github.com/Cirru/cirru-grammar"
  "fmt"
  "os"
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
        return "(string \"" + stringValue + "\")"
      }
    case "int":
      if intValue, ok := data.Value.(int); ok {
        return "(int " + fmt.Sprintf("%d", intValue) + ")"
      }
    case "float":
      if floatValue, ok := data.Value.(float64); ok {
        return "(float " + fmt.Sprintf("%g", floatValue) + ")"
      }
    case "bool":
      if value, ok := data.Value.(bool); ok {
        if value {
          return "(bool true)"
        }
        return "(bool false)"
      }
    case "array":
      list := []string{}
      if anArray, ok := data.Value.(*[]Object); ok {
        for _, item := range *anArray {
          list = append(list, stringifyObject(item))
        }
      }
      stringValue := strings.Join(list, " ")
      return "(array " + stringValue + ")"
    case "map":
      list := []string{}
      // debugPrint("string is", data.Value)
      if aMap, ok := data.Value.(*Env); ok {
        for key, value := range *aMap {
          hold := "(\"" + key + "\" " + stringifyObject(value) + ")"
          list = append(list, hold)
        }
      }
      stringValue := strings.Join(list, " ")
      return "(map " + stringValue + ")"
    case "regexp":
      return "(regexp " + fmt.Sprintf("%s", data.Value) + ")"
    case "code":
      if code, ok := data.Value.(*cirru.List); ok {
        return "(code " + codeString(*code) + ")"
      }
    default: return "(unknown)"
  }
  return "(no-type)"
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

func codeString(xs cirru.List) (ret string) {
  hold := []string{}
  for _, item := range xs {
    if buffer, ok := item.(cirru.Token); ok {
      hold = append(hold, buffer.Text)
    }
    if list, ok := item.(cirru.List); ok {
      tmpString := codeString(list)
      hold = append(hold, "(" + tmpString + ")")
    }
  }
  ret = strings.Join(hold, " ")
  return
}

func stop(text ...interface{}) {
  fmt.Println(text...)
  os.Exit(1)
}