// Copyright (c) 2019, Sylabs Inc. All rights reserved.
// Copyright (c) 2021, NVIDIA CORPORATION. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package network

import (
	"fmt"
	"log"
)

const (
	// Infiniband is the ID used to identify Infiniband
	Infiniband = "IB"
	// Default is the ID used to identify the default networking configuration
	Default = "default"
)

// SaveFn is a function of a component to save the network configuration in a configuration file
type SaveFn func() error

// Info is a structure storing the details about the network on the system
type Info struct {
	ID   string
	Save SaveFn
}

// Detect is the function called to detect the network on the system and load the corresponding networking component
func Detect() Info {
	loaded, comp := LoadDefault()
	if !loaded {
		log.Fatalln("unable to find a default network configuration")
	}

	loaded, ibComp := LoadInfiniband()
	if loaded {
		return ibComp
	}

	return comp
}

func (i *Info) Display() {
	fmt.Println("Network configuration:")
	fmt.Printf("\tID: %s\n", i.ID)
}
