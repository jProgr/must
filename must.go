package must

import "log"

func Must[T any](f func() (T, error)) T {
    v, err := f()
    if err != nil {
        log.Panicln(err)
    }

    return v
}

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

func JustDo(f func() error) {
    err := f()
    if err != nil {
        log.Panicln(err)
    }
}
