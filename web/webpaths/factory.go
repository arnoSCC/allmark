// Copyright 2014 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webpaths

import (
	"github.com/arnoSCC/allmark/common/logger"
	"github.com/arnoSCC/allmark/common/paths"
	"github.com/arnoSCC/allmark/common/route"
	"github.com/arnoSCC/allmark/dataaccess"
)

func NewFactory(logger logger.Logger, repository dataaccess.Repository) *PatherFactory {
	return &PatherFactory{
		logger:     logger,
		repository: repository,
	}
}

type PatherFactory struct {
	logger     logger.Logger
	repository dataaccess.Repository
}

func (factory *PatherFactory) Absolute(prefix string) paths.Pather {
	return newAbsoluteWebPathProvider(prefix)
}

func (factory *PatherFactory) Relative(baseRoute route.Route) paths.Pather {
	return newRelativeWebPathProvider(factory.repository, baseRoute)
}
