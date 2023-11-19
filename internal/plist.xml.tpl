<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
    <dict>
        <key>HasRemovalPasscode</key>
        <false/>
        <key>PayloadContent</key>
        <array>
            <dict>
                <key>Label</key>
                <string>{{.AppName}}</string>
                <key>URL</key>
                <string>{{.AppUrl}}</string>
                <key>Icon</key>
                <data>{{.AppIcon}}</data>
                <key>IsRemovable</key>
                <{{.Removable}}/>
                <key>FullScreen</key>
                <{{.FullScreen}}/>
                <key>PayloadDescription</key>
                <string>Configures settings for a web clip</string>
                <key>PayloadDisplayName</key>
                <string>Web Clip</string>
                <key>PayloadIdentifier</key>
                <string>com.apple.webClip.managed.{{.UUID}}</string>
                <key>PayloadType</key>
                <string>com.apple.webClip.managed</string>
                <key>PayloadUUID</key>
                <string>{{.UUID}}</string>
                <key>PayloadVersion</key>
                <integer>1</integer>
                <key>Precomposed</key>
                <true/>
            </dict>
        </array>
        <key>PayloadDisplayName</key>
        <string>Web Clip for {{.AppName}}</string>
        <key>PayloadIdentifier</key>
        <string>wg-3.{{.UUID}}</string>
        <key>PayloadRemovalDisallowed</key>
        <false/>
        <key>PayloadType</key>
        <string>Configuration</string>
        <key>PayloadUUID</key>
        <string>{{.UUID}}</string>
        <key>PayloadVersion</key>
        <integer>1</integer>
    </dict>
</plist>