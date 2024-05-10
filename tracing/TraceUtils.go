package tracing

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/TarsCloud/TarsGo/tars/util/current"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"runtime"
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

//func ReportErrorToTracing(ctx context.Context, name string, err string) error {
//	span := trace.SpanFromContext(ctx)
//	span.AddEvent("log", trace.WithAttributes(attribute.String("error", err)))
//	span.End()
//	return errors.New(err)
//}

func ReportErrorToTracing(ctx context.Context, name string, err string) error {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("error", err))
	return errors.New(err)
}

//	func ReportHttpRequestToTracing(ctx context.Context, name string, Body map[string]interface{}) {
//		j, err := json.Marshal(Body)
//		if err != nil {
//			return
//		}
//		span := trace.SpanFromContext(ctx)
//		span.AddEvent("log", trace.WithAttributes(attribute.String("response", string(j))))
//		span.End()
//	}
func ReportHttpRequestToTracing(ctx context.Context, name string, Body map[string]interface{}) {
	j, err := json.Marshal(Body)
	if err != nil {
		return
	}
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("response", string(j)))
}

type CurrentStackInfo struct {
	FullStack       []uintptr
	CurrentStack    uintptr
	CurrentFn       string
	CurrentFnLine   int
	CurrentFileName string
	OK              bool
}

//func ReportServiceInfoToTracing(ctx context.Context) error {
//	Stack := CurrentStackInfo{}
//	if Stack.CurrentStack, Stack.CurrentFileName, Stack.CurrentFnLine, Stack.OK = runtime.Caller(1); !Stack.OK {
//		return errors.New("获取当前调用栈失败")
//	}
//	Stack.CurrentFn = runtime.FuncForPC(Stack.CurrentStack).Name()
//	span := trace.SpanFromContext(ctx)
//	span.AddEvent("log",trace.WithAttributes(
//		attribute.String("CurrentFnName", Stack.CurrentFn),
//		attribute.Int("CurrentFnLine", Stack.CurrentFnLine),
//		attribute.String("CurrentFileName", Stack.CurrentFileName),
//	))
//	span.End()
//	return nil
//}

func ReportServiceInfoToTracing(ctx context.Context) error {
	Stack := CurrentStackInfo{}
	if Stack.CurrentStack, Stack.CurrentFileName, Stack.CurrentFnLine, Stack.OK = runtime.Caller(1); !Stack.OK {
		return errors.New("获取当前调用栈失败")
	}
	Stack.CurrentFn = runtime.FuncForPC(Stack.CurrentStack).Name()
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(
		attribute.String("CurrentFnName", Stack.CurrentFn),
		attribute.Int("CurrentFnLine", Stack.CurrentFnLine),
		attribute.String("CurrentFileName", Stack.CurrentFileName),
	)
	return nil
}

func StartTraceAndSpan(tarsCtx context.Context, name string) (ctx context.Context, span trace.Span) {
	ctx, span = otel.Tracer("").Start(tarsCtx, name)
	return
}

func StopTraceAndSpan(tarsCtx context.Context) (ctx context.Context, span trace.Span) {
	span = trace.SpanFromContext(tarsCtx)
	span.End()
	return
}
