package tpl

var Main string = `package main

import (
	skyclilib "github.com/avran02/sky-cli-lib"
	"github.com/avran02/sky-cli-plugins/plugin/tpl"
)

// must be named GetPluginConfig
func GetPluginConfig() skyclilib.PluginConfiger {
	return &PluginConfig{
		commands: []skyclilib.OsCommand{
			initGo,
		},
		fs: &skyclilib.Folder{
			IsOptional: false,
			FolderStructure: map[string]interface{}{
				"main.go": mainGo,
				"tpl": skyclilib.Folder{
					IsOptional: false,
				},
			},
		},
	}
}

// main object of plugin. Describes how plugin should generate commands, files and folders
type PluginConfig struct {
	commands []skyclilib.OsCommand
	fs       *skyclilib.Folder
}

// Return slice of os commands that will be executed before generation
func (p *PluginConfig) GetOsCommands() []skyclilib.OsCommand {
	return p.commands
}

// Return JSON-like structure describing virtual file system
func (p *PluginConfig) GetVirtualFs() *skyclilib.Folder {
	return p.fs
}

var (
	// files
	mainGo = skyclilib.File{
		IsOptional:     false,
		RequiredValues: map[string]string{},
		UserValues:     map[string]string{},
		Tpl:            tpl.Main(),
	}

	// commands
	initGo = skyclilib.OsCommand{
		Name: "go",
		Args: []skyclilib.CommandArg{
			{
				Name:            "mod",
				NeedGetFromUser: false,
			},
			{
				Name:            "init",
				NeedGetFromUser: false,
			},
			{
				Name:            "Project name",
				NeedGetFromUser: true,
			},
		},
	}
)
`
