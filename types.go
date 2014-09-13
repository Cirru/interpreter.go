
package interpreter

import "github.com/Cirru/parser"

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
)

type unitype struct {
  Type unitypeName
  Value interface{}
}

type context struct {
  env *scope
  args sequence
  code sequence
}

type scope map[unitype]unitype

type sequence []interface{}

type token parser.Token