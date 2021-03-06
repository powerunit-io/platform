// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package connections ...
package connections

import (
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/managers"
)

// Manager -
type Manager interface {
	managers.Manager
}

// NewManager -
func NewManager(logger *logging.Logger) Manager {
	return Manager(&managers.BaseManager{
		Logger:   logger,
		Services: make(map[string]managers.Service),
	})
}
