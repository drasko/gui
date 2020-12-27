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
		res, err := svc.Index(ctx, req.token)
		return guiRes{
			template: res.Template,
			name:     res.Name,
			data:     res.Data,
		}, err
	}
}

func thingsEndpoint(svc gui.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(thingsReq)
		res, err := svc.Things(ctx, req.token)
		return guiRes{
			template: res.Template,
			name:     res.Name,
			data:     res.Data,
		}, err
	}
}

func channelsEndpoint(svc gui.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(channelsReq)
		res, err := svc.Channels(ctx, req.token)
		return guiRes{
			template: res.Template,
			name:     res.Name,
			data:     res.Data,
		}, err
	}
}
