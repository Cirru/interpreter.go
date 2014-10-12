
package interpreter

import "github.com/Cirru/parser"

type unitypeName int

const (
  uniFloat unitypeName = iota
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
  closure *scope
  args sequence
  code sequence
}

type object map[string]unitype

type scope struct {
  parent *scope
  closure *object
}

type mapping map[unitype]unitype

type sequence []interface{}

type token parser.Token