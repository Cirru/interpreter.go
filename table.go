
package interpreter

import "fmt"

func (env *scope) getKey(x interface{}) unitype {
  var key unitype
  if tok, ok := x.(token); ok {
    key = uni(tok.Text)
  } else if seq, ok := x.(sequence); ok {
    key = Evaluate(env, seq)
  } else {
    panic("getKey excepts code")
  }
  if key.Type == uniNil {
    panic("got nil key")
  }
  return key
}

func (env *scope) getValue(x interface{}) unitype {
  if tok, ok := x.(token); ok {
    value, ok := (*env)[uni(tok.Text)]
    if !ok {
      fmt.Println(tok.Text)
      panic("get nil result")
    }
    return value
  } else if seq, ok := x.(sequence); ok {
    value := Evaluate(env, seq)
    return value
  }
  panic("getValue excepts code")
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

func (env *scope) table(xs sequence) (ret unitype) {
  ret.Type = uniTable
  hold := &scope{}
  ret.Value = hold
  for _, item := range xs {
    pair, ok := item.(sequence)
    if !ok {
      panic("table excepts sequence")
    }
    key := env.getKey(pair[0])
    value := env.getValue(pair[1])
    (*hold)[key] = value
  }
  return
}

func (env *scope) get(xs sequence) unitype {
  assertLen(xs, 1)
  key := env.getKey(xs[0])
  if value, ok := (*env)[key]; ok {
    return value
  } else {
    return uni(nil)
  }
}

func (env *scope) set(xs sequence) (ret unitype) {
  assertLen(xs, 2)
  key := env.getKey(xs[0])
  value := env.getValue(xs[1])
  (*env)[key] = value
  return value
}

func (env *scope) setTable(xs sequence) unitype {
  assertLen(xs, 3)
  target := env.getScope(xs[0])
  key := env.getKey(xs[1])
  value := env.getValue(xs[2])
  area, _ := target.Value.(*scope)
  (*area)[key] = value
  return value
}

func (env *scope) getTable(xs sequence) unitype {
  assertLen(xs, 2)
  target := env.getScope(xs[0])
  key := env.getKey(xs[1])
  area, _ := target.Value.(*scope)
  return (*area)[key]
}
