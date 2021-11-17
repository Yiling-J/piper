//nolint
package config

const Barfoo = "barfoo"

const Foo = "foo"

type foobar struct {
	Bar string

	Barfoo string

	Baz foobarBaz

	Foo string

	Qux foobarQux
}

var Foobar = foobar{

	Bar: "foobar.bar",

	Barfoo: "foobar.barfoo",

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
