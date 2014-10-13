
package interpreter

// import "fmt"

func (env *scope) getSymbol(x interface{}) string {
  if tok, ok := x.(token); ok {
    return tok.Text
  }
  seq, _ := x.(sequence)
  result := Evaluate(env, seq)
  if result.Type != uniString {
    panic("get key expects string")
  }
  ret, ok := result.Value.(string)
  if !ok {
    panic("key not in string")
  }
  return ret
}

func (env *scope) getValue(x interface{}) unitype {
  if tok, ok := x.(token); ok {
    value, ok := (*env.closure)[tok.Text]
    if ok {
      return value
    }
    if env.parent == nil {
      return uni(nil)
    }
    return env.parent.getValue(x)
  } else if seq, ok := x.(sequence); ok {
    value := Evaluate(env, seq)
    return value
  }
  panic("getValue expects code")
}

func (env *scope) setValue(key string, value unitype) unitype {
  (*env.closure)[key] = value
  return value
}