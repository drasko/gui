// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

// Package gui contains the domain concept definitions needed to support
// Mainflux gui adapter service functionality.
package gui

import (
	"context"

	"github.com/mainflux/mainflux"
)

// Service specifies coap service API.
type Service interface {
	Index(ctx context.Context, token string) error
	Things(ctx context.Context, token string) error
	Channels(ctx context.Context, token string) error
}

var _ Service = (*guiService)(nil)

type guiService struct {
	things mainflux.ThingsServiceClient
}

// New instantiates the HTTP adapter implementation.
func New(things mainflux.ThingsServiceClient) Service {
	return &guiService{
		things: things,
	}
}

func (gs *guiService) Index(ctx context.Context, token string) error {
	return nil
}

func (gs *guiService) Things(ctx context.Context, token string) error {
	return nil
}

func (gs *guiService) Channels(ctx context.Context, token string) error {
	return nil
}
