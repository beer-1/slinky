package oracle

import (
	"context"

	"github.com/skip-mev/slinky/service/servers/oracle/types"
	"google.golang.org/grpc"
)

// OracleClient defines the interface that will be utilized by the application
// to query the oracle service. This interface is meant to be implemented by
// the gRPC client that connects to the oracle service.
//
//go:generate mockery --name OracleClient --filename mock_oracle_client.go
type OracleClient interface { //nolint
	types.OracleClient

	// Start starts the oracle client. This should connect to the remote oracle
	// service and return an error if the connection fails.
	Start() error

	// Stop stops the oracle client.
	Stop() error
}

// NoOpClient is a no-op implementation of the OracleClient interface. This
// implementation is used when the oracle service is disabled.
type NoOpClient struct{}

// Start is a no-op.
func (NoOpClient) Start() error {
	return nil
}

// Stop is a no-op.
func (NoOpClient) Stop() error {
	return nil
}

// Prices is a no-op.
func (NoOpClient) Prices(
	_ context.Context,
	_ *types.QueryPricesRequest,
	_ ...grpc.CallOption,
) (*types.QueryPricesResponse, error) {
	return nil, nil
}
