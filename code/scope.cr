
set a (int 2)

print (self)

set c (child)

under c
  under parent
    print a

print $ get c a

set c x (int 3)
print $ get c x

print $ code a