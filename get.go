
package interpreter

import "fmt"

func (env *scope) getKey(x interface{}) unitype {
  var key unitype
  if tok, ok := x.(token); ok {
    key = uni(tok.Text)
  } else if seq, ok := x.(sequence); ok {
    key = Evaluate(env, seq)
  } else {
    panic("getKey expects code")
  }
  if key.Type == uniNil {
    panic("got nil key")
  }
  return key
}

func (env *scope) getValue(x interface{}) unitype {
  if tok, ok := x.(token); ok {
    value, ok := (*env)[uni(tok.Text)]
    if ok {
      return value
    }
    parent, ok := (*env)[uni("parent")]
    if !ok {
      panic("get nil result")
    }
    area, ok := parent.Value.(*scope)
    if ok {
      return area.getValue(x)
    } else {
      panic("parent is overwritten")
    }
  } else if seq, ok := x.(sequence); ok {
    value := Evaluate(env, seq)
    return value
  }
  panic("getValue expects code")
}

func (env *scope) getScope(x interface{}) unitype {
  key := env.getKey(x)
  value, ok := (*env)[key]
  if !ok {
    panic("get nil rather than scope")
  }
  if value.Type != uniTable {
    panic("value is not a scope")
  }
  return value
}

func (env *scope) self(xs sequence) unitype {
  fmt.Println("self")
  fmt.Println(*env)
  return uni(env)
}