// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package tracingvalidator

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	"github.com/stretchr/testify/require"
)

// NewSpanValidator creates a Provider that verifies a span was created that matches the specified SpanMatcher.
// The returned Provider can be used to create a client with a tracing provider that will validate spans in unit tests.
func NewSpanValidator(t *testing.T, matcher SpanMatcher, options *tracing.ProviderOptions) tracing.Provider {
	return tracing.NewProvider(func(name, version string) tracing.Tracer {
		tt := matchingTracer{
			matcher: matcher,
		}

		t.Cleanup(func() {
			require.NotNil(t, tt.match, "didn't find a span with name %s", tt.matcher.Name)
			require.True(t, tt.match.ended, "span wasn't ended")
			if tt.matcher.Kind != 0 {
				require.EqualValues(t, tt.matcher.Kind, tt.match.kind, "span kind values don't match")
			}
			require.EqualValues(t, matcher.Status, tt.match.status, "span status values don't match")
			require.ElementsMatch(t, matcher.Attributes, tt.match.attrs, "span attributes don't match")
			require.ElementsMatch(t, matcher.Links, tt.match.links, "span links don't match")
		})

		return tracing.NewTracer(func(ctx context.Context, spanName string, options *tracing.SpanOptions) (context.Context, tracing.Span) {
			kind := tracing.SpanKindInternal
			attrs := []tracing.Attribute{}
			links := []tracing.Link{}
			if options != nil {
				kind = options.Kind
				attrs = append(attrs, options.Attributes...)
				links = append(links, options.Links...)
			}
			return tt.Start(ctx, spanName, kind, attrs, links)
		}, &tracing.TracerOptions{
			SpanFromContext: func(ctx context.Context) tracing.Span {
				return convertSpan(tt.match)
			},
			LinkFromContext: func(ctx context.Context, attrs ...tracing.Attribute) tracing.Link {
				return tracing.Link{Attributes: attrs}
			},
		})
	}, options)
}

// SpanMatcher contains the values to match when a span is ended.
type SpanMatcher struct {
	// Name is the name of the span to match.
	Name string

	// Kind is the kind of the span to match.
	Kind tracing.SpanKind

	// Status is the final status of the span to match.
	Status tracing.SpanStatus

	// Attributes are the final set of attributes of the span to match.
	Attributes []tracing.Attribute

	// Links are the final set of links of the span to match.
	Links []tracing.Link
}

type matchingTracer struct {
	matcher SpanMatcher
	match   *span
}

func (mt *matchingTracer) Start(ctx context.Context, spanName string, kind tracing.SpanKind, attrs []tracing.Attribute, links []tracing.Link) (context.Context, tracing.Span) {
	if spanName != mt.matcher.Name {
		return ctx, tracing.Span{}
	}
	// span name matches our matcher, track it
	mt.match = &span{
		name:  spanName,
		kind:  kind,
		attrs: attrs,
		links: links,
	}

	return ctx, convertSpan(mt.match)
}

func convertSpan(sp *span) tracing.Span {
	return tracing.NewSpan(tracing.SpanImpl{
		End:           sp.End,
		SetStatus:     sp.SetStatus,
		SetAttributes: sp.SetAttributes,
		AddLink:       sp.AddLink,
	})
}

type span struct {
	name   string
	kind   tracing.SpanKind
	status tracing.SpanStatus
	attrs  []tracing.Attribute
	links  []tracing.Link
	desc   string
	ended  bool
}

func (s *span) End() {
	s.ended = true
}

func (s *span) SetAttributes(attrs ...tracing.Attribute) {
	s.attrs = append(s.attrs, attrs...)
}

func (s *span) AddLink(link tracing.Link) {
	s.links = append(s.links, link)
}

func (s *span) SetStatus(code tracing.SpanStatus, desc string) {
	s.status = code
	s.desc = desc
	s.ended = true
}
