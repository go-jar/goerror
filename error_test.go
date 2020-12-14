package goerror

import (
	"fmt"
	"testing"
)

func TestGoerror(t *testing.T) {
	e := New("TEST_ERROR", "test goerror")

	fmt.Println(e.Errno(), e.Msg())
	fmt.Println(e.Error())
}
