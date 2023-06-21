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

func Default[T Number | string](value T, dft T) T {
	var zero T
	if value == zero {
		return dft
	}
	return value
}

func NilDefault[T Number | string](p *T, dft T) T {
	if p == nil {
		return dft
	}
	return *p
}

func AsPointer[T Number | string | bool](v T) *T {
	return &v
}

func SliceIntersect[T Int | Uint | string](a, b []T) []T {
	var ret []T
	set := make(map[T]struct{}, len(a))
	for _, s := range a {
		set[s] = struct{}{}
	}

	for _, s := range b {
		if _, ok := set[s]; ok {
			ret = append(ret, s)
			delete(set, s)
		}
	}
	return ret
}
