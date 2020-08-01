package inspector

type TypeCastFn func(typ string, val interface{}) (interface{}, bool)

var (
	typeCastFnRegistry = make([]TypeCastFn, 0)
)

func RegisterTypeCastFn(fn TypeCastFn) {
	for _, f := range typeCastFnRegistry {
		if &f == &fn {
			return
		}
	}
	typeCastFnRegistry = append(typeCastFnRegistry, fn)
}

func TypeCast(typ string, val interface{}) (r interface{}, ok bool) {
	for _, fn := range typeCastFnRegistry {
		r, ok = fn(typ, val)
		if ok {
			return
		}
	}
	r, ok = nil, false
	return
}
