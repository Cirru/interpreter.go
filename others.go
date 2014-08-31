
package interpreter

func cirruType(env *Env, xs []interface{}) (ret Object) {
  value := cirruGet(env, xs[0:1])
  if &value != nil {
    ret.Tag = "string"
    ret.Value = value.Tag
  }
  return
}
