// Copyright 2017-2020 The ShadowEditor Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
//
// For more information, please visit: https://github.com/tengge1/ShadowEditor
// You can also visit: https://gitee.com/tengge1/ShadowEditor

package config

import (
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/tengge1/shadoweditor/helper"
	"github.com/tengge1/shadoweditor/server"
	"github.com/tengge1/shadoweditor/server/system/model"
)

func init() {
	config := Config{}
	server.Mux.UsingContext().Handle(http.MethodGet, "/api/Config/Get", config.Get)
	server.Mux.UsingContext().Handle(http.MethodPost, "/api/Config/Save", config.Save)
}

// Config 配置控制器
type Config struct {
}

// Get 获取配置信息
func (Config) Get(w http.ResponseWriter, r *http.Request) {
	db, err := server.Mongo()
	if err != nil {
		helper.Write(w, err.Error())
		return
	}

	config := model.Config{}
	find, err := db.FindOne(server.ConfigCollectionName, bson.M{}, &config)
	if err != nil {
		helper.WriteJSON(w, server.Result{
			Code: 300,
			Msg:  err.Error(),
		})
		return
	}

	if !find {
		doc1 := bson.M{
			"ID":                  primitive.NewObjectID().Hex(),
			"Initialized":         false,
			"DefaultRegisterRole": "",
		}
		db.InsertOne(server.ConfigCollectionName, doc1)
		db.FindOne(server.ConfigCollectionName, bson.M{}, &config)
	}

	result := Result{
		ID:                   config.ID,
		EnableAuthority:      server.Config.Authority.Enabled,
		Initialized:          config.Initialized,
		DefaultRegisterRole:  config.DefaultRegisterRole,
		IsLogin:              false,
		Username:             "",
		Name:                 "",
		RoleID:               "",
		RoleName:             "",
		DeptID:               "",
		DeptName:             "",
		OperatingAuthorities: []string{},
		EnableRemoteEdit:     server.Config.Remote.Enabled,
		WebSocketServerPort:  server.Config.Remote.WebSocketPort,
	}

	user, _ := server.GetCurrentUser(r)

	if user != nil {
		result.IsLogin = true
		result.Username = user.Username
		result.Name = user.Name
		result.RoleID = user.RoleID
		result.RoleName = user.RoleName
		result.DeptID = user.DeptID
		result.DeptName = user.DeptName
		result.OperatingAuthorities = user.OperatingAuthorities
	}

	helper.WriteJSON(w, server.Result{
		Code: 200,
		Msg:  "Get Successfully!",
		Data: result,
	})
}

// Save save system config.
func (Config) Save(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	defaultRegisterRole := r.FormValue("DefaultRegisterRole")

	db, err := server.Mongo()
	if err != nil {
		helper.WriteJSON(w, server.Result{
			Code: 300,
			Msg:  err.Error(),
		})
		return
	}

	var doc bson.M
	find, _ := db.FindOne(server.ConfigCollectionName, bson.M{}, &doc)

	if !find {
		doc = bson.M{
			"ID":                  primitive.NewObjectID(),
			"Initialized":         false,
			"DefaultRegisterRole": defaultRegisterRole,
		}
		db.InsertOne(server.ConfigCollectionName, doc)
	} else {
		update := bson.M{
			"$set": bson.M{
				"DefaultRegisterRole": defaultRegisterRole,
			},
		}
		db.UpdateOne(server.ConfigCollectionName, bson.M{}, update)
	}

	helper.WriteJSON(w, server.Result{
		Code: 200,
		Msg:  "Save Successfully.",
		Data: bson.M{
			"DefaultRegisterRole": defaultRegisterRole,
		},
	})
}

// Result config to front end
type Result struct {
	ID                   string
	EnableAuthority      bool
	Initialized          bool
	DefaultRegisterRole  string
	IsLogin              bool
	Username             string
	Name                 string
	RoleID               string
	RoleName             string
	DeptID               string
	DeptName             string
	OperatingAuthorities []string
	EnableRemoteEdit     bool
	WebSocketServerPort  int
}
