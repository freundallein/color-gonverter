package converter

import (
	"testing"
)

func TestConverter32(t *testing.T) {
	data := []uint32{
		0x11223344,
		0x55667788,
		0x99AABBCC,
		0xDDEEFF00,
	}
	expected := []uint32{
		0x22334411,
		0x66778855,
		0xAABBCC99,
		0xEEFF00DD,
	}
	for index, value := range data {
		returned := Converter32(value)
		if returned != expected[index] {
			t.Errorf("Expected 0x%X, got 0x%X", expected[index], returned)
		}
	}
}

func TestConverter64(t *testing.T) {
	data := []uint64{
		0x1122334411223344,
		0x5566778855667788,
		0x99AABBCC99AABBCC,
		0xDDEEFF00DDEEFF00,
	}
	expected := []uint64{
		0x2233441122334411,
		0x6677885566778855,
		0xAABBCC99AABBCC99,
		0xEEFF00DDEEFF00DD,
	}
	for index, value := range data {
		returned := Converter64(value)
		if returned != expected[index] {
			t.Errorf("Expected 0x%X, got 0x%X", expected[index], returned)
		}
	}
}
