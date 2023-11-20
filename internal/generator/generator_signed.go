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
	"io"
	"os"
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

func (s *SignedGenerator) Generate(arg Arg) (mobileConfig string, err error) {
	// TODO implement me
	panic("implement me")
}

func (s *SignedGenerator) GenerateFile(arg Arg) (file string, err error) {
	var (
		mobileConfig string
		tempFile     *os.File
	)
	if mobileConfig, err = s.Generate(arg); err != nil {
		return
	}
	if tempFile, err = os.CreateTemp("", arg.Label+".mobileconfig"); err != nil {
		return "", err
	}
	defer func() { _ = tempFile.Close() }()
	if _, err = io.WriteString(tempFile, mobileConfig); err != nil {
		return
	}
	return tempFile.Name(), nil
}
