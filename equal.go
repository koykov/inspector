package inspector

import "math"

func EqualFloat64(a, b float64, opts *DEQOptions) bool {
	prec := FloatPrecision
	if opts != nil && opts.Precision > 0 {
		prec = opts.Precision
	}
	return math.Abs(a-b) <= prec
}

func EqualFloat32(a, b float32, opts *DEQOptions) bool {
	return EqualFloat64(float64(a), float64(b), opts)
}
