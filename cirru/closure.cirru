
set 1 $ float 1
set 2 $ float 2
set 3 $ float 3

set a $ float 1

set f $ fn ()
  set b $ float 2

f

set fibo $ fn (n)
  if (<= n 2)
    , 1
    +
      fibo (- n 1)
      fibo (- n 2)

print $ fibo $ float 10