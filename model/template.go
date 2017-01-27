package model

import (
	"github.com/jinzhu/gorm"
	"github.com/rancher/go-rancher/client"
)

type Template struct {
	Name           string `json:"name"`
	Category       string `json:"category"`
	IsSystem       string `json:"isSystem"`
	Description    string `json:"description"`
	DefaultVersion string `json:"defaultVersion" yaml:"version"`
	Path           string `json:"path"`
	Maintainer     string `json:"maintainer"`
	License        string `json:"license"`
	ProjectURL     string `json:"projectURL"`
	//Labels                map[string]string      `json:"labels"`
	UpgradeFrom string `json:"upgradeFrom"`

	// TODO
	FolderName    string `json:"folderName"`
	Catalog       string `json:"catalogId"`
	EnvironmentId string `json:"environmentId"`
	//Prefix        string `json:"prefix"`
	Base         string `json:"templateBase"`
	Icon         []byte `json:"icon"`
	IconFilename string `json:"iconFilename"`
}

type TemplateModel struct {
	gorm.Model
	Template
}

type TemplateResource struct {
	client.Resource
	Template

	VersionLinks map[string]string `json:"versionLinks"`
}

type TemplateCollection struct {
	client.Collection
	Data []TemplateResource `json:"data,omitempty"`
}
