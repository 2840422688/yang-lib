package tracing

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/TarsCloud/TarsGo/tars/util/current"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type TraceIDs struct {
	TraceID      string
	SpanID       string
	ParentSpanID string
}

func GetTraceIDs(ctx context.Context) TraceIDs {
	span := trace.SpanContextFromContext(ctx)

	tid := TraceIDs{}
	tid.TraceID = span.TraceID().String()
	tid.SpanID = span.SpanID().String()

	return tid
}

func WriteServantLog(ctx context.Context) context.Context {
	ctx_span := trace.SpanContextFromContext(ctx)
	logger := tars.GetLogger("TLOG")
	spanData, err := ctx_span.MarshalJSON()
	if err != nil {
		logger.Errorf("MarshalJSON err：%v", err)
	}
	logger.Info("span", string(spanData))
	ip, ok := current.GetClientIPFromContext(ctx)
	if !ok {
		logger.Error("Error getting ip from context")
	}
	logger.Infof("Get Client Ip : %s from context", ip)
	reqContext, ok := current.GetRequestContext(ctx)
	if !ok {
		logger.Error("Error getting reqcontext from context")
	}
	logger.Infof("Get context from context: %v", reqContext)
	//ok = current.SetResponseContext(tarsCtx, result)
	//if !ok {
	//	logger.Error("error setting respose context")
	//}
	return ctx
}

func ReportErrorToTracing(ctx context.Context, name string, err string) error {
	span := trace.SpanFromContext(ctx)
	span.AddEvent("error", trace.WithAttributes(attribute.String("error", err)))
	span.End()
	return errors.New(err)
}
func ReportHttpRequestToTracing(ctx context.Context, name string, Body map[string]interface{}) {
	j, err := json.Marshal(Body)
	if err != nil {
		return
	}
	span := trace.SpanFromContext(ctx)
	span.AddEvent("info", trace.WithAttributes(attribute.String("response", string(j))))
	span.End()
}

//func StartTrace(parentCtx context.Context, spanName string, opts ...trace.SpanStartOption) (currentCtx context.Context, span trace.Span) {
//	//只有一个tracer，不再指定了
//	//opts = append(opts, trace.WithAttributes(attribute.String("AppVersoin", util.AnyToStr(version.GetAppVersion()))))
//	currentCtx, span = otel.Tracer("").Start(parentCtx, spanName, opts...)
//
//	return
//}

//func EndTrace(currentCtx context.Context, msgs ...any) {
//	span := trace.SpanFromContext(currentCtx)
//	if span == nil {
//		return
//	}
//
//	msgs = append(msgs, WithCallerDepth(1))
//	TraceAddLog(currentCtx, msgs...)
//	span.End()
//
//}
//
//func TraceHandler(parentCtx context.Context, spanName string, handlerFunc func(ctx context.Context) error) error {
//	currentCtx, _ := StartTrace(parentCtx, spanName)
//	defer func() {
//		EndTrace(curCtx, err, WithCallerDepth(1))
//	}()
//
//	err = handleFunc(curCtx)
//	return
//
//}
