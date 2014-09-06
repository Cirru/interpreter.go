
package interpreter

func cirruArray(env *Env, xs []interface{}) (ret Object) {
  ret.Tag = cirruTypeArray
  hold := []Object{}
  for _, item := range xs {
    list := []interface{}{item}
    hold = append(hold, cirruGet(env, list))
  }
  tmp := []interface{}{}
  tmp = append(tmp, &hold)
  ret.Value = tmp[0]
  return
}
