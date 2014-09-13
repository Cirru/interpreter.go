
print $ table
  a $ float 5
  b $ array (float 1) (float 2)
  c $ table
    x $ float 1
    y $ array (float 4)

set m $ table
  a $ float 1

set-table m b $ float 2

print m
