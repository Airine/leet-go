package utils

import (
    "testing"
)

func f1(arg1 string, arg2 int) (string, int) {
    return arg1, arg2
}

func TestFunctionCall(t *testing.T) {
    output := Call(f1, "test", "1")
    if output[0] != "test" || output[1] != "1" {
        t.Fail()
    }
}
