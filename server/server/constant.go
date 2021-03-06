// Copyright 2017-2020 The ShadowEditor Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
//
// For more information, please visit: https://github.com/tengge1/ShadowEditor
// You can also visit: https://gitee.com/tengge1/ShadowEditor

package server

// Collection name is the table name we store data in mongo.
// History suffix is the suffix we add to a scene collection to store history scene data.
// Version field is the field we add to a scene record, and it is the scene version number.

const (
	// CategoryCollectionName category collection name
	CategoryCollectionName string = "_Category"
	// SceneCollectionName scene collection name
	SceneCollectionName string = "_Scene"
	// MeshCollectionName mesh collection name
	MeshCollectionName string = "_Mesh"
	// MapCollectionName map collection name
	MapCollectionName string = "_Map"
	// MaterialCollectionName material collection name
	MaterialCollectionName string = "_Material"
	// AudioCollectionName audio collection name
	AudioCollectionName string = "_Audio"
	// AnimationCollectionName animation collection name
	AnimationCollectionName string = "_Animation"
	// ParticleCollectionName particle collection name
	ParticleCollectionName string = "_Particle"
	// PrefabCollectionName prefab collection name
	PrefabCollectionName string = "_Prefab"
	// CharacterCollectionName character collection name
	CharacterCollectionName string = "_Character"
	// ScreenshotCollectionName screenshot collection name
	ScreenshotCollectionName string = "_Screenshot"
	// VideoCollectionName video collection name
	VideoCollectionName string = "_Video"
	// FileCollectionName file collection name
	FileCollectionName string = "_File"
	// ConfigCollectionName config collection name
	ConfigCollectionName string = "_Config"
	// RoleCollectionName role collection name
	RoleCollectionName string = "_Role"
	// UserCollectionName user collection name
	UserCollectionName string = "_User"
	// OperatingAuthorityCollectionName operating authority collectionname
	OperatingAuthorityCollectionName string = "_OperatingAuthority"
	// DepartmentCollectionName department collection name
	DepartmentCollectionName string = "_Department"
	// PluginCollectionName plugin collection name
	PluginCollectionName string = "_Plugin"
	// TypefaceCollectionName typeface collection name
	TypefaceCollectionName string = "_Typeface"
	// ExportSceneListCollectionName export scene list collection name
	ExportSceneListCollectionName string = "_ExportSceneList"
	// HistorySuffix history suffix
	HistorySuffix string = "_history"
	// VersionField version field
	VersionField string = "_version"
)
