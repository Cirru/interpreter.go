
package interpreter

func (env *scope) _type(xs sequence) (ret unitype) {
  value := env.getValue(xs[0])
  if &value != nil {
    ret.Type = uniString
    switch value.Type {
    case 0: ret.Value = "float"
    case 1: ret.Value = "bool"
    case 2: ret.Value = "string"
    case 3: ret.Value = "regexp"
    case 4: ret.Value = "table"
    case 5: ret.Value = "array"
    case 6: ret.Value = "fn"
    case 7: ret.Value = "nil"
    default: panic("unknow type")
    }
    ret.Value = value.Type
  }
  return
}
