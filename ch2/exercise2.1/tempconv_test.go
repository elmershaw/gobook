package tempconv

import "testing"

func TestTempconv(t *testing.T) {
	var c Celsius = 100
	t.Log(CToK(c))
	t.Log(CToF(c))

	var f Fahrenheit = 212
	t.Log(FToC(f))
	t.Log(FToK(f))
}
