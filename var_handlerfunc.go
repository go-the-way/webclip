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

//
// func generateHandlerFunc(middlewares ...http.HandlerFunc) http.HandlerFunc {
// 	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method != http.MethodPost {
// 			w.WriteHeader(http.StatusMethodNotAllowed)
// 			return
// 		}
// 		appName := r.FormValue("appName")
// 		if appName == "" {
// 			writeJSON(w, `{"err":"missing appName"}`)
// 			return
// 		}
// 		appUrl := r.FormValue("appUrl")
// 		if appUrl == "" {
// 			writeJSON(w, `{"err":"missing appUrl"}`)
// 			return
// 		}
// 		appIcon, _, err := r.FormFile("appIcon")
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		if appIcon == nil {
// 			writeJSON(w, `{"err":"missing appIcon"}`)
// 			return
// 		}
// 		removable := r.FormValue("removable") == "true"
// 		fullScreen := r.FormValue("fullScreen") == "true"
// 		icon := IconReader(appIcon)
// 		mobileConfigContent, err := Generate(appName, appUrl, icon, removable, fullScreen)
// 		if err != nil {
// 			writeJSON(w, `{"err":"`+err.Error()+`"}`)
// 			return
// 		}
// 		if r.URL.Query().Has("xml") {
// 			writeXML(w, mobileConfigContent)
// 			return
// 		}
// 		writeStream(w, mobileConfigContent, appName+".mobileconfig")
// 		return
// 	})
// 	middlewares = append(middlewares, middleware.GzipHandler(h).ServeHTTP)
// 	return walkHandlerFunc(middlewares...)
// }
