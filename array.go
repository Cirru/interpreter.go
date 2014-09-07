
package interpreter

func (env *Env) array(xs []interface{}) (ret Object) {
  ret.Tag = cirruArray
  hold := []Object{}
  for _, item := range xs {
    list := []interface{}{item}
    hold = append(hold, env.get(list))
  }
  tmp := []interface{}{}
  tmp = append(tmp, &hold)
  ret.Value = tmp[0]
  return
}
