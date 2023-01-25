package customreceiver

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const (
	typeStr   = "custom"
	stability = component.StabilityLevelAlpha
)

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig,
		receiver.WithTraces(createCustomReceiver, stability))
}

func createDefaultConfig() component.Config {
	return &Config{}
}

func createCustomReceiver(_ context.Context, params receiver.CreateSettings, baseCfg component.Config, consumer consumer.Traces) (receiver.Traces, error) {
	if consumer == nil {
		return nil, component.ErrNilNextConsumer
	}

	logger := params.Logger
	customRcvrCfg := baseCfg.(*Config)

	traceRcvr := &customReceiver{
		logger:       logger,
		nextConsumer: consumer,
		config:       customRcvrCfg,
	}

	return traceRcvr, nil
}
