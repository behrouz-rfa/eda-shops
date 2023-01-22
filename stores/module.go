package stores

import (
	"context"

	"eda-shops/internal/ddd"
	"eda-shops/internal/monolith"
	"eda-shops/stores/internal/application"
	"eda-shops/stores/internal/grpc"
	"eda-shops/stores/internal/logging"
	"eda-shops/stores/internal/postgres"
	"eda-shops/stores/internal/rest"
)

type Module struct {
}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	domainDispatcher := ddd.NewEventDispatcher()
	stores := postgres.NewStoreRepository("stores.stores", mono.DB())
	participatingStores := postgres.NewParticipatingStoreRepository("stores.stores", mono.DB())
	products := postgres.NewProductRepository("stores.products", mono.DB())

	// setup application
	var app application.App
	app = application.New(stores, participatingStores, products, domainDispatcher)
	app = logging.LogApplicationAccess(app, mono.Logger())

	// setup Driver adapters
	if err := grpc.RegisterServer(ctx, app, mono.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}

	return nil
}
