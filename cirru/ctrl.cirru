
set a #false

set b $ if a
  print :true
  print :false

print a b

set c $ block
  print ":first line"
  print ":second line"
  , a

print c

if a
  block $ print ":true line"
  block $ print ":false line"