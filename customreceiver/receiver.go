package customreceiver

import (
	"context"
	"fmt"
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type customReceiver struct {
	host         component.Host
	cancel       context.CancelFunc
	logger       *zap.Logger
	nextConsumer consumer.Traces
	config       *Config
}

func (customRcvr *customReceiver) Start(ctx context.Context, host component.Host) error {
	customRcvr.host = host
	ctx = context.Background()
	ctx, customRcvr.cancel = context.WithCancel(ctx)

	go func() {
		customRcvr.startListen(ctx)
		<-ctx.Done()
	}()

	return nil
}

func (customRcvr *customReceiver) Shutdown(ctx context.Context) error {
	customRcvr.cancel()
	return nil
}

func (customRcvr *customReceiver) startListen(ctx context.Context) {
	http.HandleFunc("/", customRcvr.hello(ctx))
	http.ListenAndServe((":" + customRcvr.config.Port), nil)
}

func (customRcvr *customReceiver) hello(ctx context.Context) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello\n")
		customRcvr.nextConsumer.ConsumeTraces(ctx, ptrace.NewTraces())
	}
}
