// Copyright 2023 webclip Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package webclip

import (
	"github.com/rwscode/webclip/internal"
	"github.com/rwscode/webclip/internal/generator"
)

type (
	GenerateArg = generator.Arg
)

var (
	IconString         = internal.IconString
	IconFile           = internal.IconFile
	IconUri            = internal.IconUri
	IconReader         = internal.IconReader
	Generator          = generator.NewDefaultGenerator()
	SignedGenerator    = generator.DefaultSignedGenerator()
	Generate           = Generator.Generate
	GenerateFile       = Generator.GenerateFile
	SignedGenerate     = SignedGenerator.Generate
	SignedGenerateFile = SignedGenerator.GenerateFile

	// Serve           = serve
	// ServeRouter     = serveRouter
)
