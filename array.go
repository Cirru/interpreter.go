
package interpreter

func (env *scope) array(xs sequence) (ret unitype) {
  ret.Type = uniArray
  hold := map[unitype]unitype{}
  for index, item := range xs {
    hold[uni(index)] = env.get(sequence{item})
  }
  ret.Value = &hold
  return
}
