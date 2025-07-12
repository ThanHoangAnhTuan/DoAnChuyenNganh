package middlewares

import (
	"context"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/tracing"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func TracingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		ctx := c.Request.Context()
		ctx = trace.ContextWithRemoteSpanContext(ctx, extractSpanContext(c))

		spanName := c.Request.Method + " " + c.Request.URL.Path
		ctx, span := tracing.StartSpan(ctx, spanName,
			trace.WithSpanKind(trace.SpanKindServer),
			trace.WithAttributes(
				attribute.String("http.method", c.Request.Method),
				attribute.String("http.url", c.Request.URL.String()),
				attribute.String("http.user_agent", c.Request.UserAgent()),
			),
		)

		// Add trace ID to span attributes for Jaeger visibility
		if traceID := tracing.TraceIDFromContext(ctx); traceID != "" {
			span.SetAttributes(attribute.String("trace.id", traceID))
		}
		defer span.End()

		c.Request = c.Request.WithContext(ctx)

		c.Next()

		duration := time.Since(start)
		span.SetAttributes(
			attribute.Int("http.status_code", c.Writer.Status()),
			attribute.Int64("http.duration_ms", duration.Milliseconds()),
		)

		// Add trace ID to response headers for client
		traceID := tracing.TraceIDFromContext(ctx)
		if traceID != "" {
			c.Header("X-Trace-ID", traceID)
		}

		if c.Writer.Status() >= 400 {
			span.SetStatus(codes.Error, "")
		} else {
			span.SetStatus(codes.Ok, "")
		}
	}
}

func extractSpanContext(c *gin.Context) trace.SpanContext {
	// Extract traceparent header (W3C Trace Context)
	traceparent := c.GetHeader("traceparent")
	if traceparent != "" {
		// Parse traceparent header: 00-<trace-id>-<span-id>-<trace-flags>
		parts := strings.Split(traceparent, "-")
		if len(parts) == 4 {
			traceID := parts[1]
			spanID := parts[2]
			// Convert hex strings to trace IDs
			if traceIDBytes, err := trace.TraceIDFromHex(traceID); err == nil {
				if spanIDBytes, err := trace.SpanIDFromHex(spanID); err == nil {
					return trace.NewSpanContext(trace.SpanContextConfig{
						TraceID: traceIDBytes,
						SpanID:  spanIDBytes,
					})
				}
			}
		}
	}

	// Extract Jaeger headers
	uberTraceID := c.GetHeader("uber-trace-id")
	if uberTraceID != "" {
		// Parse Jaeger trace ID: <trace-id>:<span-id>:<parent-span-id>:<flags>
		parts := strings.Split(uberTraceID, ":")
		if len(parts) >= 2 {
			traceID := parts[0]
			spanID := parts[1]
			// Convert hex strings to trace IDs
			if traceIDBytes, err := trace.TraceIDFromHex(traceID); err == nil {
				if spanIDBytes, err := trace.SpanIDFromHex(spanID); err == nil {
					return trace.NewSpanContext(trace.SpanContextConfig{
						TraceID: traceIDBytes,
						SpanID:  spanIDBytes,
					})
				}
			}
		}
	}

	// Extract custom trace ID header
	customTraceID := c.GetHeader("X-Trace-ID")
	if customTraceID != "" {
		// Generate a new span context with the provided trace ID
		if traceIDBytes, err := trace.TraceIDFromHex(customTraceID); err == nil {
			return trace.NewSpanContext(trace.SpanContextConfig{
				TraceID: traceIDBytes,
			})
		}
	}

	return trace.SpanContext{}
}

func AddSpanToContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		traceID := tracing.TraceIDFromContext(ctx)
		if traceID != "" {
			c.Set("trace_id", traceID)
		}

		spanID := tracing.SpanIDFromContext(ctx)
		if spanID != "" {
			c.Set("span_id", spanID)
		}

		c.Next()
	}
}

func StartChildSpan(ctx context.Context, operationName string, attributes ...attribute.KeyValue) (context.Context, trace.Span) {
	ctx, span := tracing.StartSpan(ctx, operationName,
		trace.WithSpanKind(trace.SpanKindInternal),
		trace.WithAttributes(attributes...),
	)
	return ctx, span
}
