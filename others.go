
package interpreter

func (env *Env) _type(xs []interface{}) (ret Object) {
  value := env.get(xs[0:1])
  if &value != nil {
    ret.Tag = cirruString
    switch value.Tag {
    case 0: ret.Value = "int"
    case 1: ret.Value = "float"
    case 2: ret.Value = "bool"
    case 3: ret.Value = "string"
    case 4: ret.Value = "regexp"
    case 5: ret.Value = "table"
    case 6: ret.Value = "array"
    case 7: ret.Value = "fn"
    case 8: ret.Value = "nil"
    default: panic("unknow type")
    }
    ret.Value = value.Tag
  }
  return
}
