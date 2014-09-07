
print $ table
  a $ int 5
  b $ array (int 1) (int 2)
  c $ table
    int 1
    array (int 4)

set m $ table
  a $ int 1

set m b $ int 2

print m