package foobarbaz

import (
	"errors"

	"github.com/google/wire"
)

type Foo struct {
	X int
}

func ProvideFoo() Foo {
	return Foo{X: 23}
}

type Bar struct {
	Y int
}

func ProvideBar(foo Foo) Bar {
	return Bar{Y: -foo.X}
}

type Baz struct {
	Z int
}

func ProvideBaz(bar Bar) (Baz, error) {
	if bar.Y == 0 {
		return Baz{}, errors.New("cannot provide baz when bar is zero")
	}
	return Baz{Z: bar.Y}, nil
}

var SuperSet = wire.NewSet(ProvideFoo, ProvideBar, ProvideBaz)
