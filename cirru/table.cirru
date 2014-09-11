
print $ table
  a $ int 5
  b $ array (int 1) (int 2)
  c $ table
    x $ int 1
    y $ array (int 4)

set m $ table
  a $ int 1

set-table m b $ int 2

print m
