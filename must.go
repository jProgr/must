package must

import "log"

func Must[T any](f func() (T, error)) T {
    v, err := f()
    if err != nil {
        log.Panicln(err)
    }

    return v
}

func JustDo(f func() error) {
    err := f()
    if err != nil {
        log.Panicln(err)
    }
}
