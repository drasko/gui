// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"context"

	"github.com/drasko/gui/gui"
	"github.com/go-kit/kit/endpoint"
)

func indexEndpoint(svc gui.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(indexReq)
		err := svc.Index(ctx, req.token)
		return nil, err
	}
}

func thingsEndpoint(svc gui.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(thingsReq)
		err := svc.Things(ctx, req.token)
		return nil, err
	}
}

func channelsEndpoint(svc gui.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(channelsReq)
		err := svc.Channels(ctx, req.token)
		return nil, err
	}
}
