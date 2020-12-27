// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

// Package gui contains the domain concept definitions needed to support
// Mainflux gui adapter service functionality.
package gui

import (
	"context"
	"html/template"

	"github.com/mainflux/mainflux"
)

const (
	templateDir = "web/template"
)

// Service specifies coap service API.
type Service interface {
	Index(ctx context.Context, token string) (TemplateData, error)
	Things(ctx context.Context, token string) (TemplateData, error)
	Channels(ctx context.Context, token string) (TemplateData, error)
}

var _ Service = (*guiService)(nil)

type guiService struct {
	things mainflux.ThingsServiceClient
}

type TemplateData struct {
	Template *template.Template
	Name     string
	Data     interface{}
}

// New instantiates the HTTP adapter implementation.
func New(things mainflux.ThingsServiceClient) Service {
	return &guiService{
		things: things,
	}
}

func (gs *guiService) Index(ctx context.Context, token string) (TemplateData, error) {
	tmpl, err := template.ParseGlob(templateDir + "/*")
	if err != nil {
		return TemplateData{}, err
	}

	data := struct {
		Name string
	}{"John Smith"}

	return TemplateData{
		Template: tmpl,
		Name:     "index",
		Data:     data,
	}, nil
}

func (gs *guiService) Things(ctx context.Context, token string) (TemplateData, error) {
	tmpl, err := template.ParseGlob(templateDir + "/*")
	if err != nil {
		return TemplateData{}, err
	}

	data := struct {
		Name string
	}{"John Smith"}

	return TemplateData{
		Template: tmpl,
		Name:     "things",
		Data:     data,
	}, nil
}

func (gs *guiService) Channels(ctx context.Context, token string) (TemplateData, error) {
	tmpl, err := template.ParseGlob(templateDir + "/*")
	if err != nil {
		return TemplateData{}, err
	}

	data := struct {
		Name string
	}{"John Smith"}

	return TemplateData{
		Template: tmpl,
		Name:     "channels",
		Data:     data,
	}, nil
}
