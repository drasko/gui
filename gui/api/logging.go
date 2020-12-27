// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

// +build !test

package api

import (
	"context"
	"fmt"
	"time"

	"github.com/drasko/gui/gui"
	log "github.com/mainflux/mainflux/logger"
)

var _ gui.Service = (*loggingMiddleware)(nil)

type loggingMiddleware struct {
	logger log.Logger
	svc    gui.Service
}

// LoggingMiddleware adds logging facilities to the adapter.
func LoggingMiddleware(svc gui.Service, logger log.Logger) gui.Service {
	return &loggingMiddleware{logger, svc}
}

func (lm *loggingMiddleware) Index(ctx context.Context, token string) (td gui.TemplateData, err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method index took %s to complete", time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.Index(ctx, token)
}

func (lm *loggingMiddleware) Things(ctx context.Context, token string) (td gui.TemplateData, err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method things took %s to complete", time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.Things(ctx, token)
}

func (lm *loggingMiddleware) Channels(ctx context.Context, token string) (td gui.TemplateData, err error) {
	defer func(begin time.Time) {
		message := fmt.Sprintf("Method channels took %s to complete", time.Since(begin))
		if err != nil {
			lm.logger.Warn(fmt.Sprintf("%s with error: %s.", message, err))
			return
		}
		lm.logger.Info(fmt.Sprintf("%s without errors.", message))
	}(time.Now())

	return lm.svc.Channels(ctx, token)
}
