
package interpreter

type unitypeName int

const (
  uniInt unitypeName = iota
  uniFloat
  uniBool
  uniString
  uniRegexp
  uniTable
  uniArray
  uniFn
  uniNil
  uniMacro
)

type unitype struct {
  Type unitypeName
  Value interface{}
}

type context struct {
  env *scope
  args []interface{}
  code []interface{}
}

type scope map[unitype]unitype
