
package interpreter

func (env *scope) array(xs sequence) (ret unitype) {
  ret.Type = uniArray
  list := &map[unitype]unitype{}
  for index, item := range xs {
    (*list)[uni(index)] = env.getValue(item)
  }
  ret.Value = list
  return
}
