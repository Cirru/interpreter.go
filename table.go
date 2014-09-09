
package interpreter

func (env *scope) table(xs sequence) (ret unitype) {
  ret.Type = uniTable
  hold := scope{}
  for _, item := range xs {
    if pair, ok := item.(sequence); ok {
      name := pair[0]
      var key string
      if token, ok := name.(token); ok {
        key = token.Text
      }
      value := env.get(pair[1:2])
      hold[uni(key)] = value
    }
  }
  ret.Value = &hold
  return
}

func (env *scope) set(xs sequence) (ret unitype) {
  if len(xs) != 2 {
    panic("get only accepts 2 arguemnts")
  }
  value := env.get(xs[1:2])
  if tok, ok := xs[0].(token); ok {
    (*env)[uni(tok.Text)] = value
    return value
  } else {
    variable := env.get([]interface{}{xs})
    if variable.Type == uniString {
      if name, ok := variable.Value.(string); ok {
        (*env)[uni(name)] = value
        return value
      } else {
        panic("get no string in set")
      }
    } else {
      panic("and not string")
    }
  }
}

func (env *scope) setTable(xs sequence) unitype {
  if len(xs) != 3 {
    panic("setTable accepts 3 arguemnts")
  }
  hold := env.get(xs[0:1])
  if hold.Type == uniTable {
    if area, ok := hold.Value.(*scope); ok {
      value := env.get(xs[1:2])
      if tok, ok := xs[0].(token); ok {
        (*area)[uni(tok.Text)] = value
        return value
      } else {
        variable := env.get([]interface{}{xs})
        if variable.Type == uniString {
          if name, ok := variable.Value.(string); ok {
            (*area)[uni(name)] = value
            return value
          } else {
            panic("get no string in set")
          }
        } else {
          panic("and not string")
        }
      }
    } else {
      panic("not getting scope from table")
    }
  } else {
    panic("setTable expects a table")
  }
}

func (env *scope) get(xs sequence) unitype {
  if len(xs) != 1 {
    panic("get only accepts 1 arguemnt")
  }
  name := xs[0]
  if tok, ok := name.(token); ok {
    if value, ok := (*env)[uni(tok.Text)]; ok {
      return value
    }
    if parent, ok := (*env)[uni("parent")]; ok {
      if p, ok := parent.Value.(*scope); ok {
        return p.get(xs)
      } else {
        panic("parent is not a scope")
      }
    } else {
      return uni(nil)
    }
  } else if seq, ok := name.(sequence); ok {
    return Evaluate(env, seq)
  }
  return uni(nil)
}

func (env *scope) getTable(xs sequence) unitype {
  if len(xs) != 2 {
    panic("getTable accepts 2 arguemnts")
  }
  item := env.get(xs[0:1])
  if item.Type == uniTable {
    if area, ok := item.Value.(*scope); ok {
      return area.get(xs[1:2])
    } else {
      panic("not getting scope from area")
    }
  } else {
    panic("getTable expects a table")
  }
}
