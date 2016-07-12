package date

import (
	"fmt"
	"testing"
	"time"
)

func checkValid(t *testing.T, y, m, d int) {
	da := NewDate(y, m, d)
	sda := da.String()
	xsda := fmt.Sprintf("%04d-%02d-%02d", y, m, d)
	if sda != xsda {
		t.Errorf("Date.String() failed, expected %s, got %s", xsda, sda)
	}
	json, err := da.MarshalJSON()
	if err != nil {
		t.Errorf("Date.MarshalJSON() failed to serialize %s", sda)
	}
	sjson := string(json)
	if len(sjson) != 12 {
		t.Errorf("Date.MarshalJSON() serialized %s to %s", sda, sjson)
	}
	uda := Date{}
	err = uda.UnmarshalJSON(json)
	if err != nil {
		t.Errorf("Date.UnmarshalJSON() failed to deserialize %s", sjson)
	}
	if time.Time(uda) != time.Time(da) {
		t.Errorf("Dates unequal after JSON round trip: %s <=> %s", da.String(), uda.String())
	}

}

func TestDate(t *testing.T) {
	for y := 1999; y <= 2001; y++ {
		for m := 1; m <= 12; m++ {
			maxd := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}[m-1]
			checkValid(t, y, m, maxd)
			checkValid(t, y, m, 1)
		}
	}
}
