// +build wireinject

package registory

import (
	"github.com/KETAKOM/go-study-app/application/foobarbaz"
	"github.com/google/wire"
)

func InitializeBaz() (foobarbaz.Baz, error) {
	// foo := foobarbaz.ProvideFoo()
	// bar := foobarbaz.ProvideBar(foo)
	// baz, err := foobarbaz.ProvideBaz(bar)

	// if err != nil {
	// 	return foobarbaz.Baz{Z: 0}, err
	// }
	// return baz, nil

	wire.Build(foobarbaz.SuperSet)
	return foobarbaz.Baz{}, nil
	// wire.Build(foobarbaz)
	// return foobarbaz.Baz{}, nil
}
