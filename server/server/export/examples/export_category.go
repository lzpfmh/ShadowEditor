// Copyright 2017-2020 The ShadowEditor Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
//
// For more information, please visit: https://github.com/tengge1/ShadowEditor
// You can also visit: https://gitee.com/tengge1/ShadowEditor

package examples

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/tengge1/shadoweditor/helper"
	"github.com/tengge1/shadoweditor/server"
)

func exportCategory(path string) {
	port := server.Config.Server.Port

	result, _ := helper.Get(fmt.Sprintf("http://%v/api/Category/List", port))

	dirName := filepath.Join(path, "api", "Category")
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		os.MkdirAll(dirName, 0755)
	}

	// 获取列表
	fileName := filepath.Join(path, "api", "Category", "List")
	ioutil.WriteFile(fileName, []byte(result), 0755)

	// 其他接口
	apiList := []string{
		"/api/Category/Save",
		"/api/Category/Delete",
	}

	data, _ := helper.ToJSON(map[string]interface{}{
		"Code": 200,
		"Msg":  "Execute successfully!",
	})

	for _, i := range apiList {
		fileName = filepath.Join(path, i)
		ioutil.WriteFile(fileName, []byte(data), 0755)
	}
}
