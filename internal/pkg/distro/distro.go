// Copyright (c) 2019, Sylabs Inc. All rights reserved.
// Copyright (c) 2021, NVIDIA CORPORATION. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package distro

import (
	"strings"
)

// ID represents a Linux distribution
type ID struct {
	// Name is the name of the Linux distribution, e.g., ubuntu
	Name string

	// Version is the version of the Linux distribution, e.g., 7, 19.04
	Version string

	// Codename is the codename of the Linux distribution, e.g., disco (can be empty)
	Codename string
}

// ParseDescr parses the description string of a Linux distribution
// (e.g., centos:6) to a ID structure
func ParseDescr(descr string) ID {
	id := ID{
		Name:     "",
		Version:  "",
		Codename: "",
	}
	tokens := strings.Split(descr, ":")
	if len(tokens) != 2 {
		return id
	}

	id.Name = tokens[0]
	if id.Name == "ubuntu" {
		id.Codename = tokens[1]
		id.Version = ubuntuCodenameToVersion(id.Codename)
	} else {
		id.Version = tokens[1]
	}

	return id
}
