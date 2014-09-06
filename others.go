
package interpreter

func cirruType(env *Env, xs []interface{}) (ret Object) {
  value := cirruGet(env, xs[0:1])
  if &value != nil {
    ret.Tag = cirruTypeString
    switch value.Tag {
    case 0: ret.Value = "int"
    case 1: ret.Value = "float"
    case 2: ret.Value = "bool"
    case 3: ret.Value = "string"
    case 4: ret.Value = "regexp"
    case 5: ret.Value = "map"
    case 6: ret.Value = "array"
    case 7: ret.Value = "function"
    case 8: ret.Value = "code"
    default: panic("unknow type")
    }
    ret.Value = value.Tag
  }
  return
}
