
set a (int 2)

print (self)

set c (child)

under c
  under parent
    print a

print $ get c a

set c x (int 3)
print $ get c x

print $ code
  set a 1
  print (get a)
  print $ array
    int a
    array
      int a

set container (map)
set container code $ code
  set a 1
  print (get a)
  print $ array
    int a
    array
      int a

print container

set just-print $ code
  print a

print just-print

eval (self) just-print
eval just-print