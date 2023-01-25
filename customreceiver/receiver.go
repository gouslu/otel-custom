package customreceiver

import (
	"context"
	"fmt"
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
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

	customRcvr.startListen()

	ctx.Done()
	return nil
}

func (customRcvr *customReceiver) Shutdown(ctx context.Context) error {
	customRcvr.cancel()
	return nil
}

func (customRcvr *customReceiver) startListen() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe((":" + customRcvr.config.Port), nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
