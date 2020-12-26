// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

// +build !test

package api

import (
	"context"
	"time"

	"github.com/drasko/gui/gui"
	"github.com/go-kit/kit/metrics"
)

var _ gui.Service = (*metricsMiddleware)(nil)

type metricsMiddleware struct {
	counter metrics.Counter
	latency metrics.Histogram
	svc     gui.Service
}

// MetricsMiddleware instruments adapter by tracking request count and latency.
func MetricsMiddleware(svc gui.Service, counter metrics.Counter, latency metrics.Histogram) gui.Service {
	return &metricsMiddleware{
		counter: counter,
		latency: latency,
		svc:     svc,
	}
}

func (mm *metricsMiddleware) Index(ctx context.Context, token string) error {
	defer func(begin time.Time) {
		mm.counter.With("method", "index").Add(1)
		mm.latency.With("method", "index").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mm.svc.Index(ctx, token)
}

func (mm *metricsMiddleware) Things(ctx context.Context, token string) error {
	defer func(begin time.Time) {
		mm.counter.With("method", "things").Add(1)
		mm.latency.With("method", "things").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mm.svc.Things(ctx, token)
}

func (mm *metricsMiddleware) Channels(ctx context.Context, token string) error {
	defer func(begin time.Time) {
		mm.counter.With("method", "channels").Add(1)
		mm.latency.With("method", "channels").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mm.svc.Channels(ctx, token)
}
