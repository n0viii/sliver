package extensions

/*
	Sliver Implant Framework
	Copyright (C) 2021  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"encoding/json"
	"testing"
)

const (
	sample1 = `{
	"name": "test1",
	"command_name": "test1",
	"version": "1.0.0",
	"extension_author": "test",
	"original_author": "test",
	"repo_url": "https://example.com/",
	"help": "some help",
	"files": [
		{
			"os": "windows",
			"arch": "amd64",
			"path": "foo/test1.dll"
		}
	]
}`

	sample2 = `{
	"name": "test2",
	"command_name": "test2",
	"help": "some help",
	"files": [
		{
			"os": "windows",
			"arch": "amd64",
			"path": "../../../../foo/test1.dll"
		}
	]
}`
	sample3 = `{
	"name": "test3",
	"command_name": "test3",
	"version": "1.0.0",
	"extension_author": "test",
	"original_author": "test",
	"repo_url": "https://example.com/",
	"help": "some help",
	"files": [
		{
			"os": "windows",
			"arch": "amd64",
			"path": "foo/test1.dll"
		}
	]
}`
	sample4 = `[{
		"name": "testmultisingle",
		"command_name": "testmultisingle",
		"version": "1.0.0",
		"extension_author": "test",
		"original_author": "test",
		"repo_url": "https://example.com/",
		"help": "some help",
		"files": [
			{
				"os": "windows",
				"arch": "amd64",
				"path": "foo/test1.dll"
			}
		]
	}]`
	sample5 = `[{
		"name": "testmulti",
		"command_name": "testmulti",
		"version": "1.0.0",
		"extension_author": "test",
		"original_author": "test",
		"repo_url": "https://example.com/",
		"help": "some help",
		"files": [
			{
				"os": "windows",
				"arch": "amd64",
				"path": "foo/test1.dll"
			}
		]
	},{
		"name": "testmulti2",
		"command_name": "testmulti2",
		"version": "1.0.0",
		"extension_author": "test",
		"original_author": "test",
		"repo_url": "https://example.com/",
		"help": "some help",
		"files": [
			{
				"os": "windows",
				"arch": "amd64",
				"path": "foo/test1.dll"
			}
		]
	}]`
)

func TestParseExtensionManifest(t *testing.T) {
	mextManifest, err := ParseMultiManifest([]byte(sample1))
	for _, extManifest := range mextManifest { //should only be a single manfiest here, so should pass
		if err != nil {
			t.Fatalf("Error parsing extension manifest: %s", err)
		}
		if extManifest.Name != "test1" {
			t.Errorf("Expected extension name 'test1', got '%s'", extManifest.Name)
		}
		if extManifest.CommandName != "test1" {
			t.Errorf("Expected extension command name 'test1', got '%s'", extManifest.CommandName)
		}
		if extManifest.Version != "1.0.0" {
			t.Errorf("Expected extension version '1.0.0', got '%s'", extManifest.Version)
		}
		if extManifest.ExtensionAuthor != "test" {
			t.Errorf("Expected extension author 'test', got '%s'", extManifest.ExtensionAuthor)
		}
		if extManifest.OriginalAuthor != "test" {
			t.Errorf("Expected original author 'test', got '%s'", extManifest.OriginalAuthor)
		}
		if extManifest.RepoURL != "https://example.com/" {
			t.Errorf("Expected repo URL 'https://example.com/', got '%s'", extManifest.RepoURL)
		}
		if extManifest.Help != "some help" {
			t.Errorf("Expected help 'some help', got '%s'", extManifest.Help)
		}
		if len(extManifest.Files) != 1 {
			t.Errorf("Expected 1 file, got %d", len(extManifest.Files))
		}
		if extManifest.Files[0].OS != "windows" {
			t.Errorf("Expected OS 'windows', got '%s'", extManifest.Files[0].OS)
		}
		if extManifest.Files[0].Arch != "amd64" {
			t.Errorf("Expected Arch 'amd64', got '%s'", extManifest.Files[0].Arch)
		}
		if extManifest.Files[0].Path != "/foo/test1.dll" {
			t.Errorf("Expected path '/foo/test1.dll', got '%s'", extManifest.Files[0].Path)
		}
	}

	mextManifest2, err := ParseMultiManifest([]byte(sample2)) //should only be a single manfiest here, so should pass
	for _, extManifest2 := range mextManifest2 {
		if err != nil {
			t.Fatalf("Error parsing extension manifest (2): %s", err)
		}
		if extManifest2.Name != "test2" {
			t.Errorf("Expected extension name 'test2', got '%s'", extManifest2.Name)
		}
		if extManifest2.CommandName != "test2" {
			t.Errorf("Expected extension command name 'test2', got '%s'", extManifest2.CommandName)
		}
		if extManifest2.Help != "some help" {
			t.Errorf("Expected help 'some help', got '%s'", extManifest2.Help)
		}
		if len(extManifest2.Files) != 1 {
			t.Errorf("Expected 1 file, got %d", len(extManifest2.Files))
		}
		if extManifest2.Files[0].OS != "windows" {
			t.Errorf("Expected OS 'windows', got '%s'", extManifest2.Files[0].OS)
		}
		if extManifest2.Files[0].Arch != "amd64" {
			t.Errorf("Expected Arch 'amd64', got '%s'", extManifest2.Files[0].Arch)
		}
		if extManifest2.Files[0].Path != "/foo/test1.dll" {
			t.Errorf("Expected path '/foo/test1.dll', got '%s'", extManifest2.Files[0].Path)
		}
	}
}

func TestParseMultipleManifests(t *testing.T) {
	mextManifest, err := ParseMultiManifest([]byte(sample4)) //single manifest in a slice
	for _, extManifest := range mextManifest {
		if err != nil {
			t.Fatalf("Error parsing extension manifest: %s", err)
		}
		if extManifest.Name != "testmultisingle" {
			t.Errorf("Expected extension name 'testmultisingle', got '%s'", extManifest.Name)
		}
		if extManifest.CommandName != "testmultisingle" {
			t.Errorf("Expected extension command name 'testmultisingle', got '%s'", extManifest.CommandName)
		}
		if extManifest.Version != "1.0.0" {
			t.Errorf("Expected extension version '1.0.0', got '%s'", extManifest.Version)
		}
		if extManifest.ExtensionAuthor != "test" {
			t.Errorf("Expected extension author 'test', got '%s'", extManifest.ExtensionAuthor)
		}
		if extManifest.OriginalAuthor != "test" {
			t.Errorf("Expected original author 'test', got '%s'", extManifest.OriginalAuthor)
		}
		if extManifest.RepoURL != "https://example.com/" {
			t.Errorf("Expected repo URL 'https://example.com/', got '%s'", extManifest.RepoURL)
		}
		if extManifest.Help != "some help" {
			t.Errorf("Expected help 'some help', got '%s'", extManifest.Help)
		}
		if len(extManifest.Files) != 1 {
			t.Errorf("Expected 1 file, got %d", len(extManifest.Files))
		}
		if extManifest.Files[0].OS != "windows" {
			t.Errorf("Expected OS 'windows', got '%s'", extManifest.Files[0].OS)
		}
		if extManifest.Files[0].Arch != "amd64" {
			t.Errorf("Expected Arch 'amd64', got '%s'", extManifest.Files[0].Arch)
		}
		if extManifest.Files[0].Path != "/foo/test1.dll" {
			t.Errorf("Expected path '/foo/test1.dll', got '%s'", extManifest.Files[0].Path)
		}
	}

	mextManifest2, err := ParseMultiManifest([]byte(sample5)) //single manifest in a slice
	for _, extManifest := range mextManifest2 {
		if err != nil {
			t.Fatalf("Error parsing extension manifest: %s", err)
		}
		if extManifest.Name != "testmulti" && extManifest.Name != "testmulti2" {
			t.Errorf("Expected extension name 'testmulti or testmulti2', got '%s'", extManifest.Name)
		}
		if extManifest.CommandName != "testmulti" && extManifest.CommandName != "testmulti2" {
			t.Errorf("Expected extension command name 'testmulti or testmulti2', got '%s'", extManifest.CommandName)
		}
		if extManifest.Version != "1.0.0" {
			t.Errorf("Expected extension version '1.0.0', got '%s'", extManifest.Version)
		}
		if extManifest.ExtensionAuthor != "test" {
			t.Errorf("Expected extension author 'test', got '%s'", extManifest.ExtensionAuthor)
		}
		if extManifest.OriginalAuthor != "test" {
			t.Errorf("Expected original author 'test', got '%s'", extManifest.OriginalAuthor)
		}
		if extManifest.RepoURL != "https://example.com/" {
			t.Errorf("Expected repo URL 'https://example.com/', got '%s'", extManifest.RepoURL)
		}
		if extManifest.Help != "some help" {
			t.Errorf("Expected help 'some help', got '%s'", extManifest.Help)
		}
		if len(extManifest.Files) != 1 {
			t.Errorf("Expected 1 file, got %d", len(extManifest.Files))
		}
		if extManifest.Files[0].OS != "windows" {
			t.Errorf("Expected OS 'windows', got '%s'", extManifest.Files[0].OS)
		}
		if extManifest.Files[0].Arch != "amd64" {
			t.Errorf("Expected Arch 'amd64', got '%s'", extManifest.Files[0].Arch)
		}
		if extManifest.Files[0].Path != "/foo/test1.dll" {
			t.Errorf("Expected path '/foo/test1.dll', got '%s'", extManifest.Files[0].Path)
		}
	}
}

func TestParseExtensionManifestErrors(t *testing.T) {
	sample3, err := parseExtensionManifest([]byte(sample3))
	if err != nil {
		t.Fatalf("Failed to parse initial sample3: %s", err)
	}

	missingName := (*sample3)
	missingName.Name = ""
	data, _ := json.Marshal(missingName)
	_, err = ParseMultiManifest(data)
	if err == nil {
		t.Fatalf("Expected missing name error, got none")
	}

	missingCmdName := (*sample3)
	missingCmdName.CommandName = ""
	data, _ = json.Marshal(missingCmdName)
	_, err = ParseMultiManifest(data)
	if err == nil {
		t.Fatalf("Expected missing command name error, got none")
	}

	missingHelp := (*sample3)
	missingHelp.Help = ""
	data, _ = json.Marshal(missingHelp)
	_, err = ParseMultiManifest(data)
	if err == nil {
		t.Fatalf("Expected missing help error, got none")
	}

	missingFiles := (*sample3)
	missingFiles.Files = []*extensionFile{}
	data, _ = json.Marshal(missingFiles)
	_, err = ParseMultiManifest(data)
	if err == nil {
		t.Fatalf("Expected missing files error, got none")
	}

	missingFileOS := (*sample3)
	missingFileOS.Files = []*extensionFile{
		{
			OS:   "",
			Arch: "amd64",
			Path: "foo/test1.dll",
		},
	}
	data, _ = json.Marshal(missingFileOS)
	_, err = ParseMultiManifest(data)
	if err == nil {
		t.Fatalf("Expected missing files.os error, got none")
	}

	missingFileArch := (*sample3)
	missingFileArch.Files = []*extensionFile{
		{
			OS:   "windows",
			Arch: "",
			Path: "foo/test1.dll",
		},
	}
	data, _ = json.Marshal(missingFileArch)
	_, err = ParseMultiManifest(data)
	if err == nil {
		t.Fatalf("Expected missing files.arch error, got none")
	}

	missingFilePath := (*sample3)
	missingFilePath.Files = []*extensionFile{
		{
			OS:   "windows",
			Arch: "amd64",
			Path: "",
		},
	}
	data, _ = json.Marshal(missingFilePath)
	_, err = ParseMultiManifest(data)
	if err == nil {
		t.Fatalf("Expected missing files.path error, got none")
	}

	invalidPaths := []string{
		"../../../../../",
		"/../../../../..",
		".",
		"/",
	}
	for _, invalidPath := range invalidPaths {
		missingFilePath2 := (*sample3)
		missingFilePath2.Files = []*extensionFile{
			{
				OS:   "windows",
				Arch: "amd64",
				Path: invalidPath,
			},
		}
		data, _ = json.Marshal(missingFilePath2)
		_, err = ParseMultiManifest(data)
		if err == nil {
			t.Fatalf("Expected missing files.path error, got none")
		}
	}
}
