// Copyright 2017-2020 The ShadowEditor Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
//
// For more information, please visit: https://github.com/tengge1/ShadowEditor
// You can also visit: https://gitee.com/tengge1/ShadowEditor

package server

// Result present a server handler result.
type Result struct {
	// 200 - ok; 300 -error
	Code int         `json:"Code" bson:"Code"`
	Msg  string      `json:"Msg" bson:"Msg"`
	Data interface{} `json:"Data,omitempty" bson:"Data,omitempty"`
}
