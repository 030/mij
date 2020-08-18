package mij

import (
	"os"
	"testing"
)

// https://stackoverflow.com/a/34102842/2777965
func (d *dockerImage) TestMain(m *testing.M) {
	d.setup()
	code := m.Run()
	d.shutdown()
	os.Exit(code)
}
