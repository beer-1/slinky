# API Metrics

## Overview

The API Metrics package provides a set of metrics that will be implemented by default for all providers that inherit from the Base Provider and implement an API-based provider. These metrics are intended to be used by the provider to track the usage of the provider's APIs and the resources it manages.

## API Metrics

The following metrics are provided by the API Metrics package for implementations that use the API Query Handler:

```golang
// APIMetrics is an interface that defines the API for metrics collection for providers
// that implement the APIQueryHandler.
type APIMetrics interface {
	// AddProviderResponse increments the number of ticks with a fully successful provider update.
	// This increments the number of responses by provider, id (i.e. currency pair), and status.
	AddProviderResponse(providerName, id string, status Status)

	// ObserveProviderResponseTime records the time it took for a provider to respond for
	// within a single interval. Note that if the provider is not atomic, this will be the
	// time it took for all of the requests to complete.
	ObserveProviderResponseLatency(providerName string, duration time.Duration)
}
```

### AddProviderResponse

The `AddProviderResponse` metric is used to track the number of ticks with a fully successful provider update. Specifically, this tracks how often providers return good responses within the configured interval.

### ObserveProviderResponseTime

The `ObserveProviderResponseTime` metric is used to track the time it took for a provider to respond. Specifically, provider's must return a response within the configured interval. If the response time is very close to the configured interval, this could indicate that the provider is taking too long to respond, may be timing out, and consuming more resources than necessary.
