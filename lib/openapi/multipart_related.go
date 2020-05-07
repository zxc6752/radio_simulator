/*
 * Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Modifications Copyright 2020 Weiting Hu <zxc6752.cs03@g2.nctu.edu.tw>
 */

package openapi

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"radio_simulator/lib/openapi/common"
	"regexp"
)

type MultipartRelatedBinding struct{}

func (MultipartRelatedBinding) Name() string {
	return "json"
}

func (MultipartRelatedBinding) Bind(req *http.Request, obj interface{}) error {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	return common.Decode(obj, b, req.Header.Get("Content-Type"))
}

func (MultipartRelatedBinding) BindBody(body []byte, obj interface{}) error {
	re, _ := regexp.Compile(`--([a-zA-Z0-9+\-_]+)--`)
	submatch := re.FindSubmatch(body)
	if len(submatch) < 1 {
		return errors.New("cannot parse multipart boundary")
	}
	return common.Decode(obj, body, "multipart/related; boundary="+string(submatch[1]))
}

type MultipartRelatedRender struct {
	Data        interface{}
	contentType string
}

func (r MultipartRelatedRender) Render(w http.ResponseWriter) (err error) {
	payloadBuf := &bytes.Buffer{}
	ct, err := common.MultipartEncode(r.Data, payloadBuf)
	if err != nil {
		panic(err)
	}
	r.contentType = ct
	w.Header().Set("Content-Type", r.contentType)
	_, err = payloadBuf.WriteTo(w)
	if err != nil {
		panic(err)
	}
	return
}

func (r MultipartRelatedRender) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", r.contentType)
}
