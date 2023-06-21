package gox

import (
	"testing"
)

func TestMax(t *testing.T) {

	if 3 == FMax(1, 3) {
		t.Log(`ok`)
	}

	if 5 == FMax(1, 3, 5, 4, 5) {
		t.Log(`ok`)
	}

}

func TestMin(t *testing.T) {

	if -3 == FMin(5, 545, -3, 43) {
		t.Log(`ok`)
	}
}

func TestSliceContains(t *testing.T) {
	if !FContains([]int{0, 1, 3}, 4) {
		t.Log(`ok`)
	}
	if FContains([]string{"a", "b", "c"}, "c") {
		t.Log(`ok`)
	}

}

func TestFifValue(t *testing.T) {
	if 1 == FifValue(true, 1, 3) {
		t.Log(`ok`)
	}

	if 3 == FifValue(false, 1, 3) {
		t.Log(`ok`)
	}

}
