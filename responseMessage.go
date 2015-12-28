// Copyright 2015 mint.zhao.chiu@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
package response

import (
	"encoding/json"
	"fmt"
)

type ResponseMessage struct {
	Status *Status     `json:"status"`
	Body   interface{} `json:"body,omitempty"`
	Total  int64       `json:"total,omitempty"`
}

func (rsp *ResponseMessage) String() string {
	return fmt.Sprintf("[STATUS]:<%v, %v>, [BODY]:<%+v>, [TOTAL]:<%v>", rsp.Status.Code, rsp.Status.Message, rsp.Body, rsp.Total)
}

func UnmarshalResponseMessage(data []byte) *ResponseMessage {
	rsp := new(ResponseMessage)
	if err := json.Unmarshal(data, rsp); err != nil {
		return nil
	}

	return rsp
}

func NewResponseMessage(status *Status, body interface{}, total int64) *ResponseMessage {
	return &ResponseMessage{
		Status: status,
		Body:   body,
		Total:  total,
	}
}

// 新建格式化客户端返回值
func NewResponseMsgf(code int, format string, args ...interface{}) *ResponseMessage {
	return &ResponseMessage{
		Status: NewStatusf(code, format, args...),
	}
}

// 新建固定客户端返回值
func NewResponseMsg(code int) *ResponseMessage {
	return NewResponseMsgf(code, ErrMsg(code))
}

// 服务器内部错误
func NewResponseMsgInternalError(msg string) *ResponseMessage {
	return &ResponseMessage{
		Status: NewStatusInternalError(msg),
	}
}

// json请求参数格式错误
func NewResponseMsgInvalidJson(msg string) *ResponseMessage {
	return NewResponseMsgf(err_code_public_invalid_json, fmt.Sprintf(ErrMsg(err_code_public_invalid_json), msg))
}

// form请求参数格式错误
func NewResponseMsgInvalidForm(msg string) *ResponseMessage {
	return NewResponseMsgf(err_code_public_invalid_form, fmt.Sprintf(ErrMsg(err_code_public_invalid_form), msg))
}

// 请求参数错误
func NewResponseMsgInvalidParam(msg string) *ResponseMessage {
	return NewResponseMsgf(err_code_public_invalid_param, fmt.Sprintf(ErrMsg(err_code_public_invalid_param), msg))
}

// API请求验证错误
func NewResponseMsgInvalidAuth(msg string) *ResponseMessage {
	return NewResponseMsgf(err_code_public_invalid_auth, fmt.Sprintf(ErrMsg(err_code_public_invalid_auth), msg))
}

// 请求OK
func NewResponseMsgOK() *ResponseMessage {
	return NewResponseMsg(OK)
}
