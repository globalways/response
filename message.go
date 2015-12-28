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
	err_code_public_invalid_param = -iota - 1
	err_code_public_internal_error
	err_code_public_invalid_json
	err_code_public_invalid_form
	Err_Code_Public_Not_Allow_Get
	Err_Code_Public_Not_Allow_Http
	err_code_public_entity_new
	err_code_public_entity_get
	err_code_public_entity_delete
	err_code_public_entity_update
	err_code_public_entity_find
	err_code_public_bi_error
	err_code_public_invalid_auth
)

var (
	errMsgs = map[int]string{
		OK:                             "OK!",
		opt_no_more_data:               "没有更多数据了.",
		Opt_update_app:                 "app需要更新啦.",
		err_code_public_internal_error: "服务器内部处理错误: %s",
		err_code_public_invalid_param:  "参数错误: %s",
		err_code_public_invalid_json:   "请求体json错误: %s",
		err_code_public_invalid_form:   "请求体form错误: %s",
		Err_Code_Public_Not_Allow_Get:  "API不接受GET请求.",
		Err_Code_Public_Not_Allow_Http: "API不接受HTTP连接.",
		err_code_public_entity_new:     "新建实体错误: %s",
		err_code_public_entity_get:     "获取实体出错: %s",
		err_code_public_entity_delete:  "删除实体错误: %s",
		err_code_public_entity_update:  "更新实体错误: %s",
		err_code_public_entity_find:    "获取实体列表错误: %s",
		err_code_public_bi_error:       "温馨提示: %s",
		err_code_public_invalid_auth:   "请求授权错误: %s",
	}
)

func ErrMsg(code int) string {
	if v, ok := errMsgs[code]; ok {
		return v
	}

	return "Something wrong: %s"
}
