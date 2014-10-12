
package interpreter

func (env *scope) _table(xs sequence) (ret unitype) {
  ret.Type = uniTable
  hold := &mapping{}
  ret.Value = hold
  for _, item := range xs {
    pair, ok := item.(sequence)
    if !ok {
      panic("table expects sequence")
    }
    key := env.getValue(pair[0])
    value := env.getValue(pair[1])
    (*hold)[key] = value
  }
  return
}

func (env *scope) get(xs sequence) unitype {
  assertLen(xs, 1)
  return env.getValue(xs[0])
}

func (env *scope) set(xs sequence) (ret unitype) {
  assertLen(xs, 2)
  key := env.getSymbol(xs[0])
  value := env.getValue(xs[1])
  env.setValue(key, value)
  return value
}

func (env *scope) setTable(xs sequence) unitype {
  assertLen(xs, 3)
  target := env.getValue(xs[0])
  if target.Type != uniTable {
    panic("value is not a table")
  }
  key := env.getValue(xs[1])
  value := env.getValue(xs[2])
  area, _ := target.Value.(*mapping)
  (*area)[key] = value
  return value
}

func (env *scope) getTable(xs sequence) unitype {
  assertLen(xs, 2)
  target := env.getValue(xs[0])
  if target.Type != uniTable {
    panic("value is not a table")
  }
  key := env.getValue(xs[1])
  area, _ := target.Value.(*mapping)
  return (*area)[key]
}
