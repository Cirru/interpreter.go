
set swap $ macro (a b)
  set tmp (get-table outer (get a))
  set-table outer (get a) (get-table outer (get b))
  set-table outer (get b) tmp

set x $ int 1
set y $ int 2

print x y

print (string "ok with set")

expand swap x y
print x y

swap x y
print x y