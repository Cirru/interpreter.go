
// Package cirruGopher is a small interpreter of Cirru.
// It is based on `cirru-grammar`.
package cirruGopher

import (
  "fmt"
  "io/ioutil"
  "github.com/jiyinyiyong/cirru-grammar"
  "encoding/json"
)

// Interpret takes result from `cirru.Parse` and run in context.
func Interpret() error {
  filename := "code/demo.cr"
  codeByte, err := ioutil.ReadFile(filename)
  if err != nil {
    panic(err)
  }
  code := string(codeByte)
  ast := cirru.Parse(code, filename)
  var globalEnv Env
  for _, line := range ast {
    if codeLine, ok := line.([]interface{}); ok {
      Evaluate(&globalEnv, codeLine)
    }
  }
  return nil
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
  fmt.Println(list...)
}
