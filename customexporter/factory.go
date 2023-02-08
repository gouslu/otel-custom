package customexporter

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.uber.org/zap"
)

const (
	typeStr   = "custom"
	stability = component.StabilityLevelAlpha
)

func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		typeStr,
		createDefaultConfig,
		exporter.WithTraces(createTracesExporter, stability),
		exporter.WithMetrics(createMetricsExporter, stability),
		exporter.WithLogs(createLogsExporter, stability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		PanicRandomly: false,
	}
}

func newCustomExporter(ctx context.Context, config *Config, logger *zap.Logger) (*customExporter, error) {
	return &customExporter{
		config: config,
		logger: logger,
	}, nil
}

func createTracesExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
) (exporter.Traces, error) {
	oCfg := cfg.(*Config)

	custom, err := newCustomExporter(ctx, oCfg, set.Logger)
	if err != nil {
		// TODO: Placeholder for an error metric in the next PR
		return nil, fmt.Errorf("error creating a custom processor: %w", err)
	}

	return exporterhelper.NewTracesExporter(
		ctx,
		set,
		cfg,
		custom.processTraces,
		exporterhelper.WithCapabilities(custom.Capabilities()),
		exporterhelper.WithStart(custom.Start),
		exporterhelper.WithShutdown(custom.Shutdown))
}

func createMetricsExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
) (exporter.Metrics, error) {
	oCfg := cfg.(*Config)

	custom, err := newCustomExporter(ctx, oCfg, set.Logger)
	if err != nil {
		// TODO: Placeholder for an error metric in the next PR
		return nil, fmt.Errorf("error creating a custom processor: %w", err)
	}

	return exporterhelper.NewMetricsExporter(
		ctx,
		set,
		cfg,
		custom.processMetrics,
		exporterhelper.WithCapabilities(custom.Capabilities()),
		exporterhelper.WithStart(custom.Start),
		exporterhelper.WithShutdown(custom.Shutdown))
}

func createLogsExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
) (exporter.Logs, error) {
	oCfg := cfg.(*Config)

	custom, err := newCustomExporter(ctx, oCfg, set.Logger)
	if err != nil {
		// TODO: Placeholder for an error metric in the next PR
		return nil, fmt.Errorf("error creating a custom processor: %w", err)
	}

	return exporterhelper.NewLogsExporter(
		ctx,
		set,
		cfg,
		custom.processLogs,
		exporterhelper.WithCapabilities(custom.Capabilities()),
		exporterhelper.WithStart(custom.Start),
		exporterhelper.WithShutdown(custom.Shutdown))
}
