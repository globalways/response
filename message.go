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

const (
	OK int = 0
)

const (
	opt_no_more_data = OK + 1
	Opt_update_app   = OK + 2
)

const (
	err_code_public_base           = OK - 1
	err_code_public_invalid_param  = err_code_public_base - 1
	err_code_public_internal_error = err_code_public_base - 2
	err_code_public_invalid_json   = err_code_public_base - 3
	err_code_public_invalid_form   = err_code_public_base - 4
	Err_Code_Public_Not_Allow_Get  = err_code_public_base - 5
	Err_Code_Public_Not_Allow_Http = err_code_public_base - 6
	err_code_public_invalid_valid  = err_code_public_base - 7
)

var (
	errMsgs = map[int]string{
		OK:                             "OK!",
		opt_no_more_data:               "没有更多数据了.",
		Opt_update_app:                 "app需要更新啦.",
		err_code_public_base:           "未知错误.",
		err_code_public_invalid_param:  "API请求参数格式错误:",
		err_code_public_internal_error: "API服务器内部错误:",
		err_code_public_invalid_json:   "API请求json结构错误:",
		err_code_public_invalid_form:   "API请求form结构错误:",
		err_code_public_invalid_valid:  "请求参数验证错误:",
		Err_Code_Public_Not_Allow_Get:  "API不接受GET请求.",
		Err_Code_Public_Not_Allow_Http: "API不接受HTTP连接.",
	}
)

func ErrMsg(code int) string {
	if v, ok := errMsgs[code]; ok {
		return v
	}

	return "Something wrong!"
}
