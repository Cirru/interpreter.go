
// A tiny interpreter of Cirru.
package interpreter

import (
  "github.com/Cirru/parser"
  "path"
  "os"
  "io/ioutil"
)

var moduleCenter map[string]unitype

func (env *scope) require(xs sequence) (ret unitype) {
  tok, ok := xs[0].(token)
  if !ok {
    panic("require expects a token")
  }
  name := tok.Text
  if cache, ok := moduleCenter[name]; ok {
    return cache
  }
  var filepath string
  if name[0] == '.' {
    base, ok := (*env.closure)["filepath"].Value.(string)
    if !ok {
      panic("filepath is expects to be a string")
    }
    filepath = path.Join(path.Dir(base), name)
  } else {
    moduleRoot := os.Getenv("cirru_path")
    filepath = path.Join(moduleRoot, name)
  }
  return Interpret(filepath)
}

// Reads file and evaluate.
func Interpret(filepath string) (ret unitype) {
  fileScope := &scope{}
  exports := &scope{}
  (*fileScope.closure)["filepath"] = uni(filepath)
  ret = uni(exports)
  (*fileScope.closure)["exports"] = ret
  moduleCenter[filepath] = ret

  codeByte, err := ioutil.ReadFile(filepath)
  if err != nil {
    panic(err)
  }

  p := parser.NewParser()
  p.Filename(filepath)
  for _, c := range codeByte {
    p.Read(rune(c))
  }
  p.Complete()
  ast := toSequence(p.ToArray())

  for _, line := range ast {
    fileScope.getValue(line)
  }
  return
}

func toSequence(xs []interface{}) (ret sequence) {
  for _, child := range xs {
    if seq, ok := child.([]interface{}); ok {
      ret = append(ret, toSequence(seq))
    } else if t, ok := child.(parser.Token); ok {
      ret = append(ret, token(t))
    } else {
      panic("got unknown type from code")
    }
  }
  return
}