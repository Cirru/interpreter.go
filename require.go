
package cirruGopher

import (
  "github.com/Cirru/cirru-parser.go"
  "path"
  "os"
  "io/ioutil"
)

var moduleCenter Env

func cirruRequire(env *Env, xs cirru.List) (ret Object) {
  if token, ok := xs[0].(cirru.Token); ok {
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
      code := string(codeByte)
      ast := cirru.Parse(code, filepath)

      for _, line := range ast {
        if list, ok := line.(cirru.List); ok {
          Evaluate(&scope, list)
        }
      }
      return
    }
  }
  return
}