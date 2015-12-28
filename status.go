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

import "fmt"

type Status struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"msg,omitempty"`
}

// 新建格式化状态码
func NewStatusf(code int, format string, args ...interface{}) *Status {
	return &Status{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}

// 新建状态码
func NewStatus(code int) *Status {
	return NewStatusf(code, ErrMsg(code))
}

// 正常状态码
func NewStatusOK() *Status {
	return NewStatus(OK)
}

// app有更新
func NewStatusNewVersion() *Status {
	return NewStatus(Opt_update_app)
}

// 服务器内部错误状态码
func NewStatusInternalError(msg string) *Status {
	return NewStatusf(err_code_public_internal_error, fmt.Sprintf(ErrMsg(err_code_public_internal_error), msg))
}

func NewStatusBIError(msg string) *Status {
	return NewStatusf(err_code_public_bi_error, fmt.Sprintf(ErrMsg(err_code_public_bi_error), msg))
}

func NewStatusNewEntityf(msg string) *Status {
	return NewStatusf(err_code_public_entity_new, fmt.Sprintf(ErrMsg(err_code_public_entity_new), msg))
}

func NewStatusNewEntity() *Status {
	return NewStatus(err_code_public_entity_new)
}

func NewStatusGetEntity() *Status {
    return NewStatus(err_code_public_entity_get)
}

func NewStatusGetEntityf(msg string) *Status {
    return NewStatusf(err_code_public_entity_get, fmt.Sprintf(ErrMsg(err_code_public_entity_get), msg))
}

func NewStatusUpdateEntity() *Status {
	return NewStatus(err_code_public_entity_update)
}

func NewStatusUpdateEntityf(msg string) *Status {
	return NewStatusf(err_code_public_entity_update, fmt.Sprintf(ErrMsg(err_code_public_entity_update), msg))
}

func NewStatusDeleteEntity() *Status {
	return NewStatus(err_code_public_entity_delete)
}

func NewStatusDeleteEntityf(msg string) *Status {
	return NewStatusf(err_code_public_entity_delete, fmt.Sprintf(ErrMsg(err_code_public_entity_delete), msg))
}

func NewStatusFindEntities() *Status {
    return NewStatus(err_code_public_entity_find)
}

func NewStatusFindEntitiesf(msg string) *Status {
    return NewStatusf(err_code_public_entity_find, fmt.Sprintf(ErrMsg(err_code_public_entity_find), msg))
}

func NewStatusNoMoreData() *Status {
	return NewStatus(opt_no_more_data)
}

func NewStatusNoMoreDataf(msg string) *Status {
	return NewStatusf(opt_no_more_data, fmt.Sprintf(ErrMsg(opt_no_more_data), msg))
}

func (status *Status) StatusOK() bool {
	return status.Code == OK
}

func (status *Status) StatusNoMoreData() bool {
	return status.Code == opt_no_more_data
}
