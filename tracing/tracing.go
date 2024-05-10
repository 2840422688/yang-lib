package tracing

import (
	"context"
	"github.com/TarsCloud/TarsGo/tars"
	"go.opentelemetry.io/otel/sdk/trace"
	"log"
)

var tp *trace.TracerProvider

func OpenTraceProvider(serviceNameKey string) {
	if tp != nil {
		tp = NewTracerProvider(serviceNameKey, "jaeger")
	}
	return
}

func CloseTraceProvider() {
	if tp != nil {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down traces provider: %v", err)
		}
	}
	return
}

func UseTarsFilterMiddleware() {
	filter := New()
	tars.UseServerFilterMiddleware(filter.BuildServerFilter())
	tars.UseClientFilterMiddleware(filter.BuildClientFilter())
	return
}
