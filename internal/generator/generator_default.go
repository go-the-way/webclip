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

type DefaultGenerator struct{}

func NewDefaultGenerator() *DefaultGenerator { return &DefaultGenerator{} }

func (d *DefaultGenerator) Generate(arg Arg) (mobileConfig string, err error) {
	// TODO implement me
	panic("implement me")
}

func (d *DefaultGenerator) GenerateFile(arg Arg) (file string, err error) {
	// TODO implement me
	panic("implement me")
}

const defTpl = `<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
    <dict>
        <key>ConsentText</key>
        <dict>
            <key>default</key>
            <string>Created with webClip</string>
        </dict>
        <key>PayloadContent</key>
        <array>
            <dict>
                <key>FullScreen</key>
                <{{.FullScreen}}/>
                <key>Icon</key>
                <data>{{.Icon}}</data>
                <key>IsRemovable</key>
                <{{.IsRemovable}}/>
                <key>Label</key>
                <string>{{.Label}}</string>
                <key>PayloadDescription</key>
                <string>{{.PayloadDescription}}</string>
                <key>PayloadDisplayName</key>
                <string>{{.PayloadDisplayName}}</string>
                <key>PayloadIdentifier</key>
                <string>{{.UUID}}</string>
                <key>PayloadOrganization</key>
                <string>{{.PayloadOrganization}}</string>
                <key>PayloadType</key>
                <string>com.apple.webClip.managed</string>
                <key>PayloadUUID</key>
                <string>{{.UUID}}</string>
                <key>PayloadVersion</key>
                <integer>1</integer>
                <key>Precomposed</key>
                <true/>
                <key>URL</key>
                <string>{{.URL}}</string>
            </dict>
        </array>
        <key>PayloadDescription</key>
        <string>{{.PayloadDescription}}</string>
        <key>PayloadDisplayName</key>
        <string>{{.PayloadDisplayName}}</string>
        <key>PayloadIdentifier</key>
        <string>{{.UUID}}</string>
        <key>PayloadOrganization</key>
        <string>{{.PayloadOrganization}}</string>
        <key>PayloadRemovalDisallowed</key>
        <false/>
        <key>PayloadType</key>
        <string>Configuration</string>
        <key>PayloadUUID</key>
        <string>{{.UUID}}</string>
        <key>PayloadVersion</key>
        <integer>1</integer>
    </dict>
</plist>`
