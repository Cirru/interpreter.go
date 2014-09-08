
set swap $ macro (a b)
  set tmp (get outer (get a))
  set outer (get a) (get outer (get b))
  set outer (get b) tmp

set x $ int 1
set y $ int 2

print x y

expand swap x y
print x y

swap x y
print x y