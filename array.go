
package interpreter

func (env *Env) array(xs []interface{}) (ret unitype) {
  ret.Type = cirruArray
  hold := []unitype{}
  for _, item := range xs {
    list := []interface{}{item}
    hold = append(hold, env.get(list))
  }
  tmp := []interface{}{}
  tmp = append(tmp, &hold)
  ret.Value = tmp[0]
  return
}
