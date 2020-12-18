package werr

import (
	"errors"
	"fmt"
	"testing"
)

func TestRetriable_Unwrap(t *testing.T) {
	err1 := errors.New("err1")
	err2 := NewRetriable(err1)
	var err3 error = err2
	if errors.Unwrap(err3) != err1 {
		t.Fatal("errors.Unwrap(err3) != err1")
	}
	if !errors.Is(err3, err1) {
		t.Fatal("!errors.Is(err3, err1)")
	}

	err4 := fmt.Errorf("err4: %w", err3)
	if !errors.Is(err4, err1) {
		t.Fatal("!errors.Is(err4, err1)")
	}

	if err5, ok := AsRetriable(err4); !ok || err5 != err2 {
		t.Fatal("!ok || err5 != err2")
	}
	if err6, ok := AsRetriable(err3); !ok || err6 != err2 {
		t.Fatal("!ok || err6 != err2")
	}
}
