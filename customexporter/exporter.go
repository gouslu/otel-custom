package customexporter

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type customExporter struct {
	config *Config
	logger *zap.Logger
}

// Start the redaction processor
func (s *customExporter) Start(_ context.Context, _ component.Host) error {
	return nil
}

// Shutdown the redaction processor
func (s *customExporter) Shutdown(context.Context) error {
	return nil
}

func (s *customExporter) processTraces(ctx context.Context, batch ptrace.Traces) error {
	return nil
}

func (s *customExporter) processMetrics(ctx context.Context, batch pmetric.Metrics) error {
	return nil
}

func (s *customExporter) processLogs(ctx context.Context, batch plog.Logs) error {
	return nil
}

// Capabilities specifies what this processor does, such as whether it mutates data
func (s *customExporter) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: true}
}
