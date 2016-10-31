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

func TestTimeIgnored(t *testing.T) {
	d1 := FromTime(time.Date(2016, 10, 4, 1, 0, 0, 0, time.UTC))
	d2 := FromTime(time.Date(2016, 10, 4, 21, 0, 0, 0, time.UTC))
	if d2.After(d1) {
		t.Errorf("%s after %s (when created from times with hour difference)", d2, d1)
	}
	if !d2.Equal(d1) {
		t.Errorf("%s not equal to %s (when created from times with hour difference)", d2, d1)
	}
	if d2.Before(d1) {
		t.Errorf("%s before %s (when created from times with hour difference)", d2, d1)
	}
}

func TestComparisons(t *testing.T) {
	d1 := NewDate(2016, 10, 15)
	d2 := NewDate(2016, 10, 16)
	if !d1.Before(d2) {
		t.Errorf("%s before %s returned false, expected true", d1, d2)
	}
	if d1.Equal(d2) {
		t.Errorf("%s equal %s returned true, expected false", d1, d2)
	}
	if d1.After(d2) {
		t.Errorf("%s after %s returned true, expected false", d1, d2)
	}
}

func TestAdd(t *testing.T) {
	d1 := NewDate(2016, 2, 28)
	d2 := d1.AddDate(0, 0, 2)
	d3 := NewDate(2016, 3, 1)
	if !d2.Equal(d3) {
		t.Errorf("%s +2 days returned %s, expected %s", d1, d2, d3)
	}
}
