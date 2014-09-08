
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
  env *Env
  args []interface{}
  code []interface{}
}

type Env map[unitype]unitype

type code []interface{}