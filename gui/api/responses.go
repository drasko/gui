// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"html/template"
	"net/http"

	"github.com/mainflux/mainflux"
)

var (
	_ mainflux.Response = (*guiRes)(nil)
)

type guiRes struct {
	template *template.Template
	name     string
	data     interface{}
}

func (res guiRes) Code() int {
	return http.StatusCreated
}

func (res guiRes) Headers() map[string]string {
	return map[string]string{}
}

func (res guiRes) Empty() bool {
	return false
}
