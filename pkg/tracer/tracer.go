package tracer

import (
	"context"
)

type Trace struct {
	Ctx    context.Context
	Tracer *Tracer
}

type Tracer struct {
	TraceId string
	Span    []*Span
}

type Span struct {
	TraceId      string
	SpanId       string
	FunctionName string
	ParentSpan   string
	Events       []Event
}

type Event struct {
	Name       string
	Attributes map[string]interface{}
	TimeStamp  string
}
