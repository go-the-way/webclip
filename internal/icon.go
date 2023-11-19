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
	"encoding/base64"
	"io"
	"net/http"
	"os"
)

type (
	Icon interface {
		Base64() (base64 string, err error)
	}
	iconString struct {
		string
		error
	}
)

func IconString(base64 string) Icon { return &iconString{base64, nil} }

func (a *iconString) Base64() (base64 string, err error) { return a.string, a.error }

func IconFile(file string) (i Icon) {
	var (
		is  iconString
		buf []byte
	)
	i = &is
	if buf, is.error = os.ReadFile(file); is.error != nil {
		return
	}
	is.string = base64.StdEncoding.EncodeToString(buf)
	return
}

func IconUri(uri string) (i Icon) {
	var (
		is  iconString
		buf []byte
	)
	i = &is
	resp, err := http.Get(uri)
	if err != nil {
		is.error = err
		return
	}
	buf, is.error = io.ReadAll(resp.Body)
	if is.error != nil {
		return
	}
	is.string = base64.StdEncoding.EncodeToString(buf)
	return
}

func IconReader(reader io.Reader) (i Icon) {
	var (
		is  iconString
		buf []byte
	)
	i = &is
	if buf, is.error = io.ReadAll(reader); is.error != nil {
		return
	}
	is.string = base64.StdEncoding.EncodeToString(buf)
	return
}
