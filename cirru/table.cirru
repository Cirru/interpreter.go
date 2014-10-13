
print $ table
  a 5
  b $ array 1 2
  c $ table (x 1) (y $ array 4)

set m $ table (a 1)

set-table m b 2

print m
