// Copyright (c) 2019, Sylabs Inc. All rights reserved.
// Copyright (c) 2021, NVIDIA CORPORATION. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package network

import (
	"fmt"
	"log"
	"os/exec"
)

const (
	// IBForceKey is the key used in the configuration file to specific in Infiniband should always be used
	IBForceKey = "force_ib"

	// MXMDirKey is the key used in the configuration file to specify where MXM files are installed
	MXMDirKey = "mxm_dir"

	// KNEMDirKey is the key used in the configuration file to specify where knem files are installed
	KNEMDirKey = "knem_dir"
)

// LoadInfiniband is the function called to load the IB component
func LoadInfiniband() (bool, Info) {
	var ib Info

	_, err := exec.LookPath("ibstat")
	if err != nil {
		log.Println("* Infiniband not detected")
		return false, ib
	}

	log.Println("* Infiniband detected, updating the configuration file")
	ib.ID = Infiniband

	return true, ib
}

// IBSave saves the IB configuration on the system into the tool's configuration file.
func IBSave() error {
	return fmt.Errorf("Not implemented")
}
