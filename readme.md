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

The hard requirement is a "dynamic" paths to fields - at any time may require to read value of any field.

Let's check how that problem may be solved ob [showcase](https://github.com/koykov/versus/tree/master/inspector2):

### reflect

Solution https://github.com/koykov/versus/blob/master/inspector2/reflect.go#L8

It's a pure dynamic solution, it [solves the problem](https://github.com/koykov/versus/blob/master/inspector2/reflect_test.go#L7),
but [benchmark](https://github.com/koykov/versus/blob/master/inspector2/reflect_test.go#L11) looks bad:
```
BenchmarkReflect/obj.L1.L2.L3.S-8         	 3119121	       375.6 ns/op	      64 B/op	       5 allocs/op
```
Speed is not acceptable and too many allocations - they will trigger problems with GC.

### reflect2

Solution https://github.com/koykov/versus/blob/master/inspector2/reflect2.go#L12

It's also pure dynamic solution, and it should solve problem with allocations (cost of use `reflect.Value`). It's true,
but benchmarks shows problems:
```
BenchmarkReflect2/obj.L1.L2.L3.S-8         	 2973918	       391.3 ns/op	       0 B/op	       0 allocs/op
```
Allocations problem is solved, but speed is much worse than native reflect due to internal design of `reflect2.frozenConfig`
(it uses `sync.Map` inside and speed reduces due to sync operations).

### inspector

Solution https://github.com/koykov/versus/blob/master/inspector2/inspector_test.go#L13
Benchmark https://github.com/koykov/versus/blob/master/inspector2/inspector_test.go#L18
```
BenchmarkInspector/obj.L1.L2.L3.S-8         	159301698	         7.596 ns/op	       0 B/op	       0 allocs/op
```
And that speed is acceptable. See explanation is the next chapter.

## How it works

As mentioned, the perfect way is using `type assertion` together with hardcoded combinations of all possible paths to
type fields, i.e. zero reflection. Type `T` allows the following paths combinations:
* `obj.L1.L2.L3.S`
* `obj.L1.L2.L3.I`
* `obj.L1.L2.L3.F`
* `obj.L1.L2.L3`
* `obj.L1.L2`
* `obj.L1`

Type is simple, thus has only 6 combinations. The code considers these combinations looks the following https://github.com/koykov/versus/blob/master/inspector2/inspector2_ins/t_ins.go#L31.
So primitive and thus so fast.

But what about big types? They provide hundreds/thousands combinations and write code manually is so big and boring work.
The further support is a problem as well. Therefore, this work was automatized writing special [code-generation tool](https://github.com/koykov/inspector/tree/master/inspc).

## Additional features

Fields values reading is so important feature, but inspectors also provide features:
* [write values to fields](https://github.com/koykov/inspector/blob/master/inspector.go#L12)
* [compare fields values](https://github.com/koykov/inspector/blob/master/inspector.go#L17)
* [iterate iterable fields](https://github.com/koykov/inspector/blob/master/inspector.go#L19) по заданному пути
* [deep comparison of a whole types](https://github.com/koykov/inspector/blob/master/inspector.go#L21) по заданному пути
* [deep copy of types](https://github.com/koykov/inspector/blob/master/inspector.go#L27)
* [reset types](https://github.com/koykov/inspector/blob/master/inspector.go#L35)

## Basic types support

### static

For support of basic types (int, uint, float64, ...) was developed special [static](https://github.com/koykov/inspector/blob/master/static.go)
inspector. It uses in [dyntpl](https://github.com/koykov/dyntpl) and [decoder](https://github.com/koykov/decoder)
packages for primitive types.

### strings

Special for types `string` and `[][]byte` was developed inspector [strings](https://github.com/koykov/inspector/blob/master/strings.go).
See [test/bench](https://github.com/koykov/inspector/blob/master/test/strings_test.go).

### map[string]any

Popular type uses together with `encoding/json` supported by inspector [map\[string\]any](https://github.com/koykov/inspector/blob/master/stranymap.go).
See test/bench https://github.com/koykov/inspector/blob/master/test/stranymap_test.go
