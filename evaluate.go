
package interpreter

import "fmt"

// Evaluate expressions and return a unitype value
func Evaluate(env *scope, xs sequence) (ret unitype) {
  if len(xs) == 0 {
    return uni(nil)
  }

  if tok, ok := xs[0].(token); ok {
    switch tok.Text {
    case "--":        ret = env.comment(xs[1:])
    case "array":     ret = env.array(xs[1:])
    case "fn":        ret = env.fn(xs[1:])
    case "bool":      ret = env._bool(xs[1:])
    case "call":      ret = env.call(xs[1:])
    case "float":     ret = env.float(xs[1:])
    case "get":       ret = env.get(xs[1:])
    case "table":     ret = env._table(xs[1:])
    case "print":     ret = env._print(xs[1:])
    case "regexp":    ret = env.regexp(xs[1:])
    case "require":   ret = env.require(xs[1:])
    case "set":       ret = env.set(xs[1:])
    case "string":    ret = env._string(xs[1:])
    case "type":      ret = env._type(xs[1:])
    case "set-table": ret = env.setTable(xs[1:])
    case "get-table": ret = env.getTable(xs[1:])
    case "if":        ret = env._if(xs[1:])
    case "block":     ret = env.block(xs[1:])
    case "+":         ret = env.add(xs[1:])
    case "-":         ret = env.minus(xs[1:])
    case "=":         ret = env.equal(xs[1:])
    case ">":         ret = env.greatThan(xs[1:])
    case "<":         ret = env.littleThan(xs[1:])
    case ">=":        ret = env.greatEqual(xs[1:])
    case "<=":        ret = env.littleEqual(xs[1:])
    default:
      ret = userCall(env, xs)
    }
  } else if _, ok := xs[0].(sequence); ok {
    ret = userCall(env, xs)
  }
  return
}

func userCall(env *scope, xs sequence) (ret unitype) {
  head := env.getValue(xs[0])
  ret = uni(nil)
  if head.Type == uniFn {
    ret = env.call(xs)
  } else {
    fmt.Println(xs, head, *env)
    panic("can not find handler")
  }
  return
}