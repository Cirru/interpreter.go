
set a 1

set f $ fn ()
  set b 2

f

set fibo $ fn (n)
  if (<= n 2)
    , 1
    +
      fibo (- n 1)
      fibo (- n 2)

print $ fibo 10