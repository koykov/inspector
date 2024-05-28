# Inspector

Inspector is a code-generation framework of special wrappers of arbitrary types and data structures that allows to
read/write/iterate fields without using reflection and produce minimum or zero allocations. 

Each type-wrapper (next "inspector" or "type-inspector") has name of original type with suffix `Inspector` and implements
[Inspector](https://github.com/koykov/inspector/blob/master/inspector.go#L4) interface.

## Intro

The main idea: you may read/write/iterate (inspect) structure of arbitrary type in common way. The historical reason of
development that framework was [dyntpl](https://github.com/koykov/dyntpl) package and many others. Usually this problem
solves using reflection, but that way is extremely slow in general and produces huge amount of allocs on every use of
`reflect.Value` type. There is good library [reflect2](https://github.com/modern-go/reflect2) that solves 'reflect.Value'
problem, but it is also slow. The perfect way if to use `type assertion` together with hardcoded combinations of all
possible paths to type fields. Unfortunately this way isn't a pure "dynamic solution". Let's consider that problem using
example:

Let we have type [`T`](https://github.com/koykov/versus/blob/master/inspector2/types/types.go#L3):
```go
type T struct {
	L1 *L1
}

type L1 struct {
	L2 *L2
}

type L2 struct {
	L3 *L3
}

type L3 struct {
	S string
	I int64
	F float64
}
```
with many nested subtypes. And we need to read data of fields for arbitrary path, eg:
* `obj.L1.L2.L3.S`
* `obj.L1.L2.L3.F`
* ...

