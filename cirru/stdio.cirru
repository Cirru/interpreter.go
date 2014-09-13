
set a $ string 1
print a

print (string 1)

print
  table
    a (float 4)
    b $ table
      a $ float 5
      b $ float 6
      c $ table
        x $ float 7

print
  array
    float 1
    float 2
    array
      float 3
      float 4

print
  array
    float 1
    table
      a $ float 2
      b $ array
        float 3

print
  float 1
  float 2

print $ fn ()
  set a 1
  print (get a)
  print $ array
    float a
    array
      float a

set container (table)
set-table container f $ fn ()
  set a 1
  print (get a)
  print $ array
    float a
    array
      float a

print container