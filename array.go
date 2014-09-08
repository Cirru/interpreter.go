
package interpreter

func (env *Env) array(xs []interface{}) (ret unitype) {
  ret.Type = uniArray
  hold := map[unitype]unitype{}
  for index, item := range xs {
    hold[uni(index)] = env.get([]interface{}{item})
  }
  ret.Value = &hold
  return
}
