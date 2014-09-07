
package interpreter

type cirruTypeName int

const (
  cirruTypeInt cirruTypeName = iota
  cirruTypeFloat
  cirruTypeBool
  cirruTypeString
  cirruTypeRegexp
  cirruTypeTable
  cirruTypeArray
  cirruTypeFunction
  cirruTypeCode
)