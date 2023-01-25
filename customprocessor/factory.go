package customprocessor

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/processorhelper"
	"go.uber.org/zap"
)

const (
	typeStr   = "custom"
	stability = component.StabilityLevelAlpha
)

func NewFactory() processor.Factory {
	return processor.NewFactory(
		typeStr,
		createDefaultConfig,
		processor.WithTraces(createTracesProcessor, stability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		PanicRandomly: false,
	}
}

func newCustomProcessor(ctx context.Context, config *Config, logger *zap.Logger, next consumer.Traces) (*customProcessor, error) {
	return &customProcessor{
		config: config,
		logger: logger,
		next:   next,
	}, nil
}

func createTracesProcessor(
	ctx context.Context,
	set processor.CreateSettings,
	cfg component.Config,
	next consumer.Traces,
) (processor.Traces, error) {
	oCfg := cfg.(*Config)

	custom, err := newCustomProcessor(ctx, oCfg, set.Logger, next)
	if err != nil {
		// TODO: Placeholder for an error metric in the next PR
		return nil, fmt.Errorf("error creating a custom processor: %w", err)
	}

	return processorhelper.NewTracesProcessor(
		ctx,
		set,
		cfg,
		next,
		custom.processTraces,
		processorhelper.WithCapabilities(custom.Capabilities()),
		processorhelper.WithStart(custom.Start),
		processorhelper.WithShutdown(custom.Shutdown))
}
