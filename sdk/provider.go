package sdk

import "context"

// ProviderStatus represents the health status of a provider.
type ProviderStatus string

const (
	ProviderStatusOK          ProviderStatus = "ok"
	ProviderStatusUnhealthy   ProviderStatus = "unhealthy"
	ProviderStatusUnreachable ProviderStatus = "unreachable"
)

// ProviderTestResult holds the result of a provider health check.
type ProviderTestResult struct {
	Status  ProviderStatus
	Message string
	Error   error
}

// ModelTestResult holds the result of a model support check.
type ModelTestResult struct {
	Supported bool
	Message   string
}

// Provider is the interface that AI backends must implement.
type Provider interface {
	Name() string
	ListModels(ctx context.Context) ([]Model, error)
	Test(ctx context.Context) *ProviderTestResult
	TestModel(ctx context.Context, modelID string) (*ModelTestResult, error)
	DoGenerate(ctx context.Context, params GenerateParams) (*GenerateResult, error)
	DoStream(ctx context.Context, params GenerateParams) (*StreamResult, error)
}
