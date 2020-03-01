package main

import (
	"errors"
	"fmt"
	"testing"
)

//紧跟在verb之前的[n]符号表示应格式化第n个参数（索引从1开始）。同样的在'*'之前的[n]符号表示采用第n个参数的值作为宽度或精度。在处理完方括号表达式[n]后，除非另有指示，会接着处理参数n+1，n+2……（就是说移动了当前处理位置）
func main() {
	fmt.Printf("%[2]d %[1]d\n", 11, 22)
	fmt.Printf("%[3]*.[2]*[1]f\n", 12.0, 2, 6)
	fmt.Printf("%d %d %#[1]x %#x\n", 16, 17)
}

func TestErrorf(t *testing.T) {
	wrapped := errors.New("inner error")
	for _, test := range []struct {
		err        error
		wantText   string
		wantUnwrap error
	}{{
		err:        fmt.Errorf("%w", wrapped),
		wantText:   "inner error",
		wantUnwrap: wrapped,
	}, {
		err:        fmt.Errorf("added context: %w", wrapped),
		wantText:   "added context: inner error",
		wantUnwrap: wrapped,
	}, {
		err:        fmt.Errorf("%w with added context", wrapped),
		wantText:   "inner error with added context",
		wantUnwrap: wrapped,
	}, {
		err:        fmt.Errorf("%s %w %v", "prefix", wrapped, "suffix"),
		wantText:   "prefix inner error suffix",
		wantUnwrap: wrapped,
	}, {
		err:        fmt.Errorf("%[2]s: %[1]w", wrapped, "positional verb"),
		wantText:   "positional verb: inner error",
		wantUnwrap: wrapped,
	}, {
		err:      fmt.Errorf("%v", wrapped),
		wantText: "inner error",
	}, {
		err:      fmt.Errorf("added context: %v", wrapped),
		wantText: "added context: inner error",
	}, {
		err:      fmt.Errorf("%v with added context", wrapped),
		wantText: "inner error with added context",
	}, {
		err:      fmt.Errorf("%w is not an error", "not-an-error"),
		wantText: "%!w(string=not-an-error) is not an error",
	}, {
		err:      fmt.Errorf("wrapped two errors: %w %w", errString("1"), errString("2")),
		wantText: "wrapped two errors: 1 %!w(fmt_test.errString=2)",
	}, {
		err:      fmt.Errorf("wrapped three errors: %w %w %w", errString("1"), errString("2"), errString("3")),
		wantText: "wrapped three errors: 1 %!w(fmt_test.errString=2) %!w(fmt_test.errString=3)",
	}, {
		err:        fmt.Errorf("%w", nil),
		wantText:   "%!w(<nil>)",
		wantUnwrap: nil, // still nil
	}} {
		if got, want := errors.Unwrap(test.err), test.wantUnwrap; got != want {
			t.Errorf("Formatted error: %v\nerrors.Unwrap() = %v, want %v", test.err, got, want)
		}
		if got, want := test.err.Error(), test.wantText; got != want {
			t.Errorf("err.Error() = %q, want %q", got, want)
		}
	}
}

type errString string

func (e errString) Error() string { return string(e) }
