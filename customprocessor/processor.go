package customprocessor

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type customProcessor struct {
	config *Config
	logger *zap.Logger
}

// Start the redaction processor
func (s *customProcessor) Start(_ context.Context, _ component.Host) error {
	return nil
}

// Shutdown the redaction processor
func (s *customProcessor) Shutdown(context.Context) error {
	return nil
}

func (s *customProcessor) processTraces(ctx context.Context, traces ptrace.Traces) (ptrace.Traces, error) {
	return traces, nil
}

func (s *customProcessor) processMetrics(ctx context.Context, metrics pmetric.Metrics) (pmetric.Metrics, error) {
	return metrics, nil
}

func (s *customProcessor) processLogs(ctx context.Context, logs plog.Logs) (plog.Logs, error) {
	return logs, nil
}

// Capabilities specifies what this processor does, such as whether it mutates data
func (s *customProcessor) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: true}
}
