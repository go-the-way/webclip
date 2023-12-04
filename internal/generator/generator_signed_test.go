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

package generator

import (
	"os"
	"testing"
)

func TestSignedGenerator_Generate(t *testing.T) {
	g := DefaultSignedGenerator()
	buf, err := g.Generate(Arg{
		Label:               "MyApp",
		Icon:                b64,
		URL:                 "http://www.qq.com",
		PayloadDisplayName:  "PayloadDisplayName",
		PayloadDescription:  "PayloadDescription",
		PayloadOrganization: "ivi.cx",
	})
	t.Log(len(buf))
	t.Log(err)
}

func TestSignedGenerator_GenerateFile(t *testing.T) {
	g := DefaultSignedGenerator()
	f, _ := os.Create("/Users/local/Desktop/www/mc.mobileconfig")
	err := g.GenerateFile(Arg{
		Label:               "MyApp",
		Icon:                b64,
		URL:                 "http://www.qq.com",
		PayloadDisplayName:  "PayloadDisplayName",
		PayloadDescription:  "PayloadDescription",
		PayloadOrganization: "ivi.cx",
	}, f)
	_ = f.Close()
	t.Log(err)
}
