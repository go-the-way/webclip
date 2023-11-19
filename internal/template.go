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

package internal

import (
	"strings"
	"text/template"

	"github.com/google/uuid"
)

type generateArg struct {
	AppName, AppUrl       string
	AppIcon               string
	Removable, FullScreen bool
	UUID                  string
}

func Generate(appName, appUrl string, appIcon Icon, removable, fullScreen bool) (xml string, err error) {
	var base64Icon string
	if base64Icon, err = appIcon.Base64(); err != nil {
		return
	}
	return generatePListXml(generateArg{
		AppName:    appName,
		AppUrl:     appUrl,
		AppIcon:    base64Icon,
		Removable:  removable,
		FullScreen: fullScreen,
		UUID:       uuid.New().String(),
	})
}

func generatePListXml(arg generateArg) (xml string, err error) {
	var (
		tpl *template.Template
	)
	if tpl, err = template.New("").Parse(pListXml); err != nil {
		return
	}
	var buf strings.Builder
	if err = tpl.Execute(&buf, arg); err != nil {
		return
	}
	xml = buf.String()
	return
}
