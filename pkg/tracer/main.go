package tracer

import (
	"context"
	"time"

	"github.com/google/uuid"
)

func CreateTracer() *Trace {
	return &Trace{
		Tracer: &Tracer{},
		Ctx:    context.Background(),
	}
}

func (t *Trace) AddSpan() {
	spanId := uuid.NewString()

	span := Span{SpanId: spanId}

	//assigns last called function as parent if any
	if len(t.Tracer.Span) <= 0 {
		span.ParentSpan = spanId
	} else {
		span.ParentSpan = t.Tracer.Span[len(t.Tracer.Span)-1].SpanId
	}

	//appends span to trace
	t.Tracer.Span = append(t.Tracer.Span, &span)
}

func (s *Span) AddEvent(name string, Attributes map[string]interface{}) {
	s.Events = append(s.Events, Event{
		Name:       name,
		TimeStamp:  time.Now().Format(time.RFC3339),
		Attributes: Attributes,
	})
}
