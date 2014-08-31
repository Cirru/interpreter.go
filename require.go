
package interpreter

import (
  "github.com/Cirru/parser"
  "path"
  "os"
  "io/ioutil"
)

var moduleCenter Env

func cirruRequire(env *Env, xs []interface{}) (ret Object) {
  if token, ok := xs[0].(parser.Token); ok {
    name := token.Text
    if cache, ok := moduleCenter[name]; ok {
      ret = cache
      return
    } else {
      var filepath string
      if name[0] == '.' {
        if baseFile, ok := (*env)["filepath"].Value.(string); ok {
          filepath = path.Join(path.Dir(baseFile), name)
        }
      } else {
        moduleRoot := os.Getenv("cirru_path")
        filepath = path.Join(moduleRoot, name)
      }
      scope := Env{}
      exports := Env{}
      scope["filepath"] = generateString(filepath)
      ret = generateMap(&exports)
      scope["exports"] = ret
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
      ast := p.ToTree()

      for _, line := range ast {
        if list, ok := line.([]interface{}); ok {
          Evaluate(&scope, list)
        }
      }
      return
    }
  }
  return
}