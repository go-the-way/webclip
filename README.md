# webclip
The Web Clip implementation in Go

Features
---
- **Gzip compression embedded**

Functions
---
- **IconString** `The app icon base64 string`
- **IconFile** `The app icon from file`
- **IconUri** `The app icon from uri resource`
- **IconReader** `The app icon from reader`
- **DefaultGenerator** `Generator without signed`
- **SignedGenerator** `Generator with ivi signed`
- **WrappedGenerateHandlerFunc** `Get the wrapped generate handler func`
- **Serve** `Serve embedded http server`
- **ServeRoute** `Serve embedded http server with custom routers`

Embedded Routers
---

- **GET: /webclip/generate** `The Generate function called`

**_request_**

`Content-Type: multipart/form-data`

`Form Value: appName` The Web Clip app name

`Form Value: appUrl` The Web Clip app url

`Form File: appIcon` The Web Clip app icon stream

`Form Value: removable` The Web Clip app removable

`Form Value: fullScreen` The Web Clip app fullScreen

**_response_**

*stream with fileName=${appName}.mobileconfig*

- **GET: /webclip/generate?xml** `The Generate function called, rendering xml`