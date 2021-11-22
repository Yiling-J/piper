//nolint
package config

const Barfoo = "barfoo"

const Foo = "foo"

type foobar struct {
	Bar string

	Baz foobarBaz

	Foo string

	Qux foobarQux
}

var Foobar = foobar{

	Bar: "foobar.bar",

	Baz: FoobarBaz,

	Foo: "foobar.foo",

	Qux: FoobarQux,
}

type foobarBaz struct {
	Xyz string
}

var FoobarBaz = foobarBaz{

	Xyz: "foobar.baz.xyz",
}

type foobarQux struct {
	Xyz string
}

var FoobarQux = foobarQux{

	Xyz: "foobar.qux.xyz",
}

const Foofoo = "foofoo"

type quux struct {
	Foo string
}

var Quux = quux{

	Foo: "quux.foo",
}

type types struct {
	Bool string

	Duration string

	Float string

	Int string

	Intslice string

	String string

	Stringslice string
}

var Types = types{

	Bool: "types.bool",

	Duration: "types.duration",

	Float: "types.float",

	Int: "types.int",

	Intslice: "types.int_slice",

	String: "types.string",

	Stringslice: "types.string_slice",
}
