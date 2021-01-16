// Copyright (c) 2019, Sylabs Inc. All rights reserved.
// Copyright (c) 2021, NVIDIA CORPORATION. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package network

// LoadDefault is the function called to load the default component for networking
func LoadDefault() (bool, Info) {
	var network Info
	network.ID = Default

	return true, network
}
