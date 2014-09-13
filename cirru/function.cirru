
set f $ fn (a b c)
  print a b c

call f (int 1) (int 2) (int 3)

f (int 1) (int 2) (int 3)

print $ string "simulating macro"

set swap $ fn (a b)
  set tmp (get-table outer (get a))
  set-table outer (get a) (get-table outer (get b))
  set-table outer (get b) tmp

set x $ int 1
set y $ int 2

call swap (string x) (string y)
print x y

swap (string x) (string y)
print x y