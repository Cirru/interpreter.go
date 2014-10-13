
set a :1
print a

print :1

print $ table (a :4)
  b $ table (a 5) (b 6)
    c $ table (x 7)

print $ array 1 2
  array 3 4

print $ array
  , 1
  table (a 2) (b $ array 3)

print 1 2

print $ fn ()
  set a 1
  print (get a)
  print $ array a
    array a

set container (table)
set-table container f $ fn ()
  set a 1
  print (get a)
  print $ array a
    array a

print container