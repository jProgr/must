package must

import (
    "bytes"
    "errors"
    "log"
    "os"
    "strings"
    "testing"
)

var dummyLog bytes.Buffer

func TestMain(tests *testing.M) {
    log.SetOutput(&dummyLog)

    os.Exit(tests.Run())
}

func TestRunsASimpleFunction(_ *testing.T) {
    JustDo(func() error { return nil })
}

func TestOnJustDoItPanicsWithError(test *testing.T) {
    defer recoverAndCheckPanic(test)

    JustDo(func() error { return errors.New("Panic") })

    test.Fatal("This shouldn't be executed")
}

func TestOnJustDoItPanicsAndLogs(test *testing.T) {
    errorString := "hhzqqpaqgrschmlzhwqa"

    defer recoverAndCheckLog(test, errorString)

    JustDo(func() error { return errors.New(errorString) })
}

func TestRunsAFunctionAndReturnsAValue(test *testing.T) {
    value := Must(func() (int, error) {
        return 7, nil
    })

    if value != 7 {
        test.Fatal("value should be 7")
    }
}

func TestIfMustFailsItPanicsWithError(test *testing.T) {
    defer recoverAndCheckPanic(test)

    Must(func() (int, error) { return 7, errors.New("Panic") })

    test.Fatal("This shouldn't be executed")
}

func TestIfMustFailsItPanicsAndLogs(test *testing.T) {
    errorString := "xBBLzpdmGphhzqqpaqgrschmlzhwqa"

    defer recoverAndCheckLog(test, errorString)

    JustDo(func() error { return errors.New(errorString) })
}

func TestItMapsAValue(test *testing.T) {
    value := MustOn(
        6,
        func(_ int) bool { return true },
        func(value int) int { return value + 1 },
        errors.New("Panic"),
    )

    if value != 7 {
        test.Fatal("value should be 7")
    }
}

func TestItDoesNotMapAValueOnFailure(test *testing.T) {
    defer recoverAndCheckPanic(test)

    MustOn(
        6,
        func(_ int) bool { return false },
        func(value int) int {
            test.Fatal("This shouldn't be executed")

            return value + 1
        },
        errors.New("Panic"),
    )

    test.Fatal("This shouldn't be executed")
}

func recoverAndCheckPanic(test *testing.T) {
    err, ok := recover().(string)
    if !ok {
        test.Fatal("It should return a string")
    }
    if err != "Panic\n" {
        test.Fatal("Wrong error")
    }
}

func recoverAndCheckLog(test *testing.T, err string) {
    recover()

    logged := dummyLog.String()
    if !strings.Contains(logged, err) {
        test.Fatal("Wrong log: ", logged)
    }
}
