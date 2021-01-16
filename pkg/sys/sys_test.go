// Copyright (c) 2021, NVIDIA CORPORATION. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package sys

import "testing"

func TestSys(t *testing.T) {
	cfg, err := Load()
	if err != nil {
		t.Fatalf("unable to load the system configuration: %s", err)
	}
	cfg.Display()
}
