
package interpreter

import (
  "github.com/Cirru/parser"
  "path"
  "os"
  "io/ioutil"
)

var moduleCenter scope

func (env *scope) require(xs sequence) (ret unitype) {
  tok, ok := xs[0].(token)
  if !ok {
    panic("require expects a token")
  }
  name := tok.Text
  if cache, ok := moduleCenter[uni(name)]; ok {
    return cache
  }
  var filepath string
  if name[0] == '.' {
    base, ok := (*env)[uni("filepath")].Value.(string)
    if !ok {
      panic("filepath is expects to be a string")
    }
    filepath = path.Join(path.Dir(base), name)
  } else {
    moduleRoot := os.Getenv("cirru_path")
    filepath = path.Join(moduleRoot, name)
  }
  fileScope := &scope{}
  exports := &scope{}
  (*fileScope)[uni("filepath")] = uni(filepath)
  ret = uni(exports)
  (*fileScope)[uni("exports")] = ret
  moduleCenter[uni(filepath)] = ret

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
    seq, _ := line.(sequence)
    Evaluate(fileScope, seq)
  }
  return
}