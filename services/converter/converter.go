// Copyright 2014 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package converter

import (
	"github.com/arnoSCC/allmark/common/paths"
	"github.com/arnoSCC/allmark/model"
)

type Converter interface {
	// Convert the supplied item with all paths relative to the supplied base route
	Convert(aliasResolver func(alias string) *model.Item, pathProvider paths.Pather, item *model.Item) (convertedContent string, converterError error)
}
