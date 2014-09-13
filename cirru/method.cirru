
set cat $ table
  name $ string Kitty
  show-name $ fn ()
    print $ get-table this name

print cat

cat show-name

print $ cat name