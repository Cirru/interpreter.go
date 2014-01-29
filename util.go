
package cirruGopher

import (
  "github.com/Cirru/cirru-parser.go"
  "fmt"
  "os"
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
    case "string":
      if stringValue, ok := data.Value.(string); ok {
        return "string \"" + stringValue + "\""
      }
    case "int":
      if intValue, ok := data.Value.(int); ok {
        return "int " + fmt.Sprintf("%d", intValue)
      }
    case "float":
      if floatValue, ok := data.Value.(float64); ok {
        return "float " + fmt.Sprintf("%g", floatValue)
      }
    case "bool":
      if value, ok := data.Value.(bool); ok {
        if value {
          return "bool true"
        }
        return "bool false"
      }
    case "array":
      list := []string{}
      indent := "\n" + repeatBlank(level + 1)
      if anArray, ok := data.Value.(*[]Object); ok {
        for _, item := range *anArray {
          list = append(list, stringifyObject(item, (level + 1)))
        }
      }
      stringValue := strings.Join(list, indent)
      return "array " + indent + stringValue
    case "map":
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
    case "regexp":
      return "regexp " + fmt.Sprintf("%s", data.Value)
    case "code":
      if code, ok := data.Value.(*cirru.List); ok {
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
  ret.Tag = "string"
  ret.Value = x
  return
}

func generateMap(m *Env) (ret Object) {
  ret.Tag = "map"
  ret.Value = m
  return
}

func codeString(xs cirru.List, level int) (ret string) {
  hold := []string{}
  indent := "\n" + repeatBlank(level + 1)
  for _, item := range xs {
    if buffer, ok := item.(cirru.Token); ok {
      hold = append(hold, buffer.Text)
    }
    if list, ok := item.(cirru.List); ok {
      tmpString := indent + codeString(list, (level + 1))
      hold = append(hold, tmpString)
    }
  }
  ret = strings.Join(hold, " ")
  return
}

func stop(text ...interface{}) {
  fmt.Println(text...)
  os.Exit(1)
}