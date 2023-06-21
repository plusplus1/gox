package gox

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Float interface {
	~float64 | ~float32
}

type Number interface {
	Int | Uint | Float
}

func FContains[T Number | string](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func FMax[T Number](a T, data ...T) T {
	num := a
	for _, d := range data {
		if d > num {
			num = d
		}
	}
	return num
}

func FMin[T Number](a T, data ...T) T {
	num := a
	for _, d := range data {
		if d < num {
			num = d
		}
	}
	return num
}

func FifValue[T Number | string](expr bool, trueValue, falseValue T) T {
	if expr {
		return trueValue
	}
	return falseValue
}
