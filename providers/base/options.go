package base

import (
	"go.uber.org/zap"

	"github.com/skip-mev/slinky/oracle/config"
	apihandlers "github.com/skip-mev/slinky/providers/base/api/handlers"
	providermetrics "github.com/skip-mev/slinky/providers/base/metrics"
	wshandlers "github.com/skip-mev/slinky/providers/base/websocket/handlers"
)

// ProviderOption is a function that can be used to modify a provider.
type ProviderOption[K comparable, V any] func(*BaseProvider[K, V])

// WithName sets the name of the provider.
func WithName[K comparable, V any](name string) ProviderOption[K, V] {
	return func(p *BaseProvider[K, V]) {
		if name == "" {
			panic("cannot set empty name")
		}

		p.name = name
	}
}

// WithLogger sets the logger for the provider.
func WithLogger[K comparable, V any](logger *zap.Logger) ProviderOption[K, V] {
	return func(p *BaseProvider[K, V]) {
		if logger == nil {
			panic("cannot set nil logger")
		}

		p.logger = logger.With(zap.String("provider", p.name))
	}
}

// WithAPIQueryHandler sets the APIQueryHandler for the provider. If your provider utilizes a
// API (HTTP) based provider, you should use this option to set the APIQueryHandler.
func WithAPIQueryHandler[K comparable, V any](api apihandlers.APIQueryHandler[K, V]) ProviderOption[K, V] {
	return func(p *BaseProvider[K, V]) {
		if p.api != nil {
			panic("cannot set api query handler twice")
		}

		if api == nil {
			panic("cannot set nil api query handler")
		}

		p.api = api
	}
}

// WithAPIConfig sets the APIConfig for the provider.
func WithAPIConfig[K comparable, V any](cfg config.APIConfig) ProviderOption[K, V] {
	return func(p *BaseProvider[K, V]) {
		if cfg.ValidateBasic() != nil {
			panic("invalid api config")
		}

		p.apiCfg = cfg
	}
}

// WithWebSocketQueryHandler sets the WebSocketQueryHandler for the provider. If your provider
// utilizes a websocket based provider, you should use this option to set the WebSocketQueryHandler.
func WithWebSocketQueryHandler[K comparable, V any](ws wshandlers.WebSocketQueryHandler[K, V]) ProviderOption[K, V] {
	return func(p *BaseProvider[K, V]) {
		if p.ws != nil {
			panic("cannot set web socket query handler twice")
		}

		if ws == nil {
			panic("cannot set nil web socket query handler")
		}

		p.ws = ws
	}
}

// WithWebSocketConfig sets the WebSocketConfig for the provider.
func WithWebSocketConfig[K comparable, V any](cfg config.WebSocketConfig) ProviderOption[K, V] {
	return func(p *BaseProvider[K, V]) {
		if cfg.ValidateBasic() != nil {
			panic("invalid web socket config")
		}

		p.wsCfg = cfg
	}
}

// WithIDs sets the IDs that the provider is responsible for fetching data for.
func WithIDs[K comparable, V any](ids []K) ProviderOption[K, V] {
	return func(p *BaseProvider[K, V]) {
		if ids == nil {
			panic("cannot set nil ids")
		}

		p.ids = ids
	}
}

// WithMetrics sets the metrics implementation for the provider.
func WithMetrics[K comparable, V any](metrics providermetrics.ProviderMetrics) ProviderOption[K, V] {
	return func(p *BaseProvider[K, V]) {
		if metrics == nil {
			panic("cannot set nil metrics")
		}

		p.metrics = metrics
	}
}
