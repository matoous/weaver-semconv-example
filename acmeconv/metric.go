// Code generated from semantic convention specification. DO NOT EDIT.

// Package httpconv provides types and functionality for OpenTelemetry semantic
// conventions in the "acme" namespace.
package acmeconv

import (
	"context"
	"sync"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
)

var (
	addOptPool = &sync.Pool{New: func() any { return &[]metric.AddOption{} }}
	recOptPool = &sync.Pool{New: func() any { return &[]metric.RecordOption{} }}
)

// ProductTypeAttr is an attribute conforming to the acme.product.type semantic
// conventions. It represents the type of an Acme corp product.
type ProductTypeAttr string

var (
	// ProductTypeHelicopter is a helicopter.
	ProductTypeHelicopter ProductTypeAttr = "helicopter"
	// ProductTypeCar is a car.
	ProductTypeCar ProductTypeAttr = "car"
)

// SalesCount is an instrument used to record metric values conforming to the
// "acme.sales.count" semantic conventions. It represents the count of all sales.
type SalesCount struct {
	metric.Int64Counter
}

// NewSalesCount returns a new SalesCount instrument.
func NewSalesCount(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (SalesCount, error) {
	// Check if the meter is nil.
	if m == nil {
		return SalesCount{noop.Int64Counter{}}, nil
	}

	i, err := m.Int64Counter(
		"acme.sales.count",
		append([]metric.Int64CounterOption{
			metric.WithDescription("Count of all sales"),
			metric.WithUnit("{sale}"),
		}, opt...)...,
	)
	if err != nil {
	    return SalesCount{noop.Int64Counter{}}, err
	}
	return SalesCount{i}, nil
}

// Inst returns the underlying metric instrument.
func (m SalesCount) Inst() metric.Int64Counter {
	return m.Int64Counter
}

// Name returns the semantic convention name of the instrument.
func (SalesCount) Name() string {
	return "acme.sales.count"
}

// Unit returns the semantic convention unit of the instrument
func (SalesCount) Unit() string {
	return "{sale}"
}

// Description returns the semantic convention description of the instrument
func (SalesCount) Description() string {
	return "Count of all sales"
}

// Add adds incr to the existing count.
//
// The productType is the type of an Acme corp product.
func (m SalesCount) Add(
	ctx context.Context,
	incr int64,
	productType ProductTypeAttr,
	attrs ...attribute.KeyValue,
) {
	o := addOptPool.Get().(*[]metric.AddOption)
	defer func() {
		*o = (*o)[:0]
		addOptPool.Put(o)
	}()

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs,
				attribute.String("acme.product.type", string(productType)),
			)...,
		),
	)

	m.Int64Counter.Add(ctx, incr, *o...)
}