
set a $ bool false

set b $ if a
  print $ string true
  print $ string false

print a b

set c $ block
  print $ string "first line"
  print $ string "second line"
  , a

print c

if a
  block $ print $ string "true line"
  block $ print $ string "false line"