
package interpreter

import "github.com/Cirru/parser"

type context struct {
  env *Env
  args []interface{}
  code []interface{}
}

func cirruBlock(env *Env, xs []interface{}) (ret Object) {
  ret.Tag = cirruTypeBlock
  if args, ok := xs[0].([]interface{}); ok {
    ret.Value = context{env, args, xs[1:]}
  }
  return
}

func cirruCall(env *Env, xs []interface{}) (ret Object) {
  block := cirruGet(env, xs[0:1])
  // debugPrint("block is", block)
  if block.Tag == cirruTypeBlock {
    if item, ok := block.Value.(context); ok {
      runtime := Env{}
      for i, para := range item.args {
        // println("i is:", i)
        // debugPrint(xs)
        if token, ok := para.(parser.Token); ok {
          runtime[token.Text] = cirruGet(env, xs[i+1:i+2])
        }
      }
      for _, line := range item.code {
        if exp, ok := line.([]interface{}); ok {
          ret = Evaluate(&runtime, exp)
        }
      }
      return
    }
  } else {
    stop("not block")
  }
  return
}