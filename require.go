
package interpreter

import (
  "github.com/Cirru/parser"
  "path"
  "os"
  "io/ioutil"
)

var moduleCenter scope

func (env *scope) require(xs sequence) (ret unitype) {
  if token, ok := xs[0].(token); ok {
    name := token.Text
    if cache, ok := moduleCenter[uni(name)]; ok {
      ret = cache
      return
    } else {
      var filepath string
      if name[0] == '.' {
        if baseFile, ok := (*env)[uni("filepath")].Value.(string); ok {
          filepath = path.Join(path.Dir(baseFile), name)
        }
      } else {
        moduleRoot := os.Getenv("cirru_path")
        filepath = path.Join(moduleRoot, name)
      }
      fileScope := scope{}
      exports := scope{}
      fileScope[uni("filepath")] = uni(filepath)
      ret = uni(&exports)
      fileScope[uni("exports")] = ret
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
        if seq, ok := line.(sequence); ok {
          Evaluate(&fileScope, seq)
        }
      }
      return
    }
  }
  return
}