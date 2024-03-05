package projects

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestIncomeCalculator(t *testing.T) {

	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	IncomeCalculator()

	_ = w.Close()

	result, _ := io.ReadAll(r)

	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "$35620") {
		t.Error("wrong balance returned")
	}

}
