// Package must provides helper functions for executing tasks which may return an
// error and there is no point in handling it. Every function in the package logs
// and then panics on error.
//
// # Examples
//
// If clipboard.Init (golang.design/x/clipboard) returns an error
// we can't continue. If it succeeds, there is no return value so
// we can just continue.
//
//  JustDo(clipboard.Init)
//
//  // Do something with clipboard…
//
// If os.Open() fails, we can't continue so Must() just panics
// with the error returned by os.Open(). Otherwise, we ditch the nil
// error and we get just the file.
//
//  file := Must(func() (*os.File, error) {
//      return os.Open("some/path.txt")
//  })
//
//  // Do something with file…
//
// We need the second arg of os.Args but it is possible that it
// is not there. In such case, we can't continue so we just panic
// (len(args) > 1 returns false) with the error provided. Otherwise
// we can continue so we extract the arg from os.Args (args[1]) and
// we have such value in arg.
//
//  arg := MustOn(
//      os.Args,
//      func(args []string) bool { return len(args) > 1 },
//      func(args []string) string { return args[1] },
//      errors.New("An arg must be provided"),
//  )
//
//  // Do something with arg…
package must

import "log"

// Must runs f() and returns a value of type T. If an error is returned by
// f(), it panics with that error.
func Must[T any](f func() (T, error)) T {
    v, err := f()
    if err != nil {
        log.Panicln(err)
    }

    return v
}

// MustOn runs mapper() on value of type T and returns the result if evaluator()
// returns true. Otherwise, it panics with err.
func MustOn[T, U any](
    value T,
    evaluator func(value T) bool,
    mapper func(value T) U,
    err any,
) U {
    if !evaluator(value) {
        log.Panicln(err)
    }

    return mapper(value)
}

// JustDo runs f() and if an error is returned, it panics.
func JustDo(f func() error) {
    err := f()
    if err != nil {
        log.Panicln(err)
    }
}
