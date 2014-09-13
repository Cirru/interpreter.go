
package interpreter

func (env *scope) add(xs sequence) unitype {
  sum := 0.0
  for _, item := range xs {
    value, _ := env.getValue(item).Value.(float64)
    sum += value
  }
  return uni(sum)
}

func (env *scope) minus(xs sequence) unitype {
  lastValue, _ := env.getValue(xs[0]).Value.(float64)
  for _, item := range xs[1:] {
    value, _ := env.getValue(item).Value.(float64)
    lastValue -= value
  }
  return uni(lastValue)
}

func (env *scope) equal(xs sequence) unitype {
  assert(len(xs) >= 2, "need at least 2 item")
  lastValue, _ := env.getValue(xs[0]).Value.(float64)
  for _, item := range xs[1:] {
    value, _ := env.getValue(item).Value.(float64)
    if lastValue != value {
      return uni(false)
    }
    lastValue = value
  }
  return uni(true)
}

func (env *scope) greatThan(xs sequence) unitype {
  assert(len(xs) >= 2, "need at least 2 item")
  lastValue, _ := env.getValue(xs[0]).Value.(float64)
  for _, item := range xs[1:] {
    value, _ := env.getValue(item).Value.(float64)
    if lastValue <= value {
      return uni(false)
    }
    lastValue = value
  }
  return uni(true)
}

func (env *scope) greatEqual(xs sequence) unitype {
  assert(len(xs) >= 2, "need at least 2 item")
  lastValue, _ := env.getValue(xs[0]).Value.(float64)
  for _, item := range xs[1:] {
    value, _ := env.getValue(item).Value.(float64)
    if lastValue < value {
      return uni(false)
    }
    lastValue = value
  }
  return uni(true)
}

func (env *scope) littleThan(xs sequence) unitype {
  assert(len(xs) >= 2, "need at least 2 item")
  lastValue, _ := env.getValue(xs[0]).Value.(float64)
  for _, item := range xs[1:] {
    value, _ := env.getValue(item).Value.(float64)
    if lastValue >= value {
      return uni(false)
    }
    lastValue = value
  }
  return uni(true)
}

func (env *scope) littleEqual(xs sequence) unitype {
  assert(len(xs) >= 2, "need at least 2 item")
  lastValue, _ := env.getValue(xs[0]).Value.(float64)
  for _, item := range xs[1:] {
    value, _ := env.getValue(item).Value.(float64)
    if lastValue > value {
      return uni(false)
    }
    lastValue = value
  }
  return uni(true)
}