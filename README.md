
Cirru Interpreter in Go
------

Run Cirru in Go.

See [master/code/](https://github.com/Cirru/interpreter/tree/master/code) for demos.

Check `cirru/` directories to see what Cirru look like.
Notice that this repo is still working in progress experimentally.

Running `go build bin/cli.go` generates a binary.

[![GoDoc](https://godoc.org/github.com/Cirru/interpreter?status.png)](https://godoc.org/github.com/Cirru/interpreter)

### Compact literal syntax

```
a      -- variables start with letter
1.2    -- numbers start with digits
#true  -- bool values start with sharp
:str   -- strings start with colon
/^x    -- regular expressions start with slash
@nil   -- special values starts with at
```

### License

MIT