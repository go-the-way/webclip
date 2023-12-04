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
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	iviUrl = "https://ivi.cx"
	cfgUrl = "/i/cfg"
)

type SignedGenerator struct {
	IviUrl string
	CfgUrl string
}

func NewSignedGenerator(iviUrl string, cfgUrl string) *SignedGenerator {
	return &SignedGenerator{IviUrl: iviUrl, CfgUrl: cfgUrl}
}

func DefaultSignedGenerator() *SignedGenerator { return NewSignedGenerator(iviUrl, cfgUrl) }

func (s *SignedGenerator) Generate(arg Arg) (buf []byte, err error) {
	reqUrl := s.IviUrl + s.CfgUrl
	resp, rErr := http.Post(reqUrl, "application/json;charset=UTF-8", strings.NewReader(arg.JSON()))
	if rErr != nil {
		err = rErr
		return
	}
	bf, bErr := io.ReadAll(resp.Body)
	if bErr != nil {
		err = bErr
		return
	}
	type result struct {
		Id string `json:"id"`
	}
	var rs result
	// {"id":"656d953d"}
	if err = json.Unmarshal(bf, &rs); err != nil {
		return
	}
	if rs.Id == "" {
		err = errors.New("cfg id is empty")
		return
	}
	reqUrl += "?" + rs.Id
	resp, rErr = http.Get(reqUrl)
	if rErr != nil {
		err = rErr
		return
	}
	return io.ReadAll(resp.Body)
}

func (s *SignedGenerator) GenerateFile(arg Arg, file *os.File) (err error) {
	if file == nil {
		return errors.New("file is nil")
	}
	buf, err := s.Generate(arg)
	if err != nil {
		return err
	}
	_, err = file.Write(buf)
	return
}
