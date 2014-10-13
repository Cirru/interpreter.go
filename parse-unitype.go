
package interpreter

import (
  "strconv"
  "regexp"
)

func parseUnitype(x string) (unitype, bool) {
  switch x[0] {
  case '#': return parseUniBool(x), true
  case ':': return parseUniString(x), true
  case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
    return parseUniFloat(x), true
  case '/': return parseUniRegexp(x), true
  case '@': return parseSpecialUni(x), true
  default: return unitype{uniNil, nil}, false
  }
}

func parseUniBool(x string) unitype {
  text := x[1:]
  switch text {
  case "true", "yes", "t":
    return unitype{uniBool, true}
  case "false", "no", "f":
    return unitype{uniBool, false}
  default: panic("cannot parse as bool")
  }
}

func parseUniFloat(x string) unitype {
  n, err := strconv.ParseFloat(x, 64)
  if err != nil {
    panic(err)
  }
  return unitype{uniFloat, n}
}

func parseUniString(x string) unitype {
  return unitype{uniString, x[1:]}
}

func parseUniRegexp(x string) unitype {
  reg, err := regexp.Compile(x[1:])
  if err != nil {
    panic(err)
  }
  return unitype{uniRegexp, reg}
}

func parseSpecialUni(x string) unitype {
  if x[1:] == "nil" {
    return unitype{uniNil, nil}
  }
  panic("special word not implemented")
}
