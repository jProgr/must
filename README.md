# Must

A collecion of functions that I use on personal and one-time scripts. The package provides helper functions for executing tasks which may return an error and there is no point in handling it. Every function in the package logs and then panics on error.

## Examples

If `clipboard.Init` (golang.design/x/clipboard) returns an error we can't continue. If it succeeds, there is no return value so we can just continue.

```go
JustDo(clipboard.Init)

// Do something with clipboard…
```

If `os.Open()` fails, we can't continue so `Must()` just panics with the error returned by `os.Open()`. Otherwise, we ditch the `nil` error and we get just the file.

```go
file := Must(func() (*os.File, error) {
  return os.Open("some/path.txt")
})

// Do something with file…
```

We need the second arg of `os.Args` but it is possible that it is not there. In such case, we can't continue so we just panic (`len(args) > 1` returns `false`) with the error provided. Otherwise we can continue so we extract the arg from `os.Args` (`args[1]`) and we have such value in arg.

```go
arg := MustOn(
  os.Args,
  func(args []string) bool { return len(args) > 1 },
  func(args []string) string { return args[1] },
  errors.New("An arg must be provided"),
)

// Do something with arg…
```
