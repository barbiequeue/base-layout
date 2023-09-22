package main

import "testing"

func TestConfParamGetType(t *testing.T) {

	expectedInt := 10
	gotInt := configParameter("10").Int()
	if gotInt != expectedInt {
		t.Errorf("configParameter.Int() failed. gotInt %d, expected %d",
			gotInt, expectedInt)
	}

	expectedFloat := 10.1
	gotFloat := configParameter("10.1").Float64()
	if gotFloat != expectedFloat {
		t.Errorf("configParameter.Float64() failed. gotInt %f, expected %f",
			gotFloat, expectedFloat)
	}

	gotBool := configParameter("").Bool()
	if gotBool != true {
		t.Errorf("configParameter.Bool() failed. gotInt %t, expected %t",
			gotBool, true)
	}
}
