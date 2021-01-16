// Copyright (c) 2019, Sylabs Inc. All rights reserved.
// Copyright (c) 2021, NVIDIA CORPORATION. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package sys

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gvallee/go_sys/internal/pkg/network"
)

// SetConfigFn is a "function pointer" that lets us store the configuration of a given job manager
type SetConfigFn func() error

// GetConfigFn is a "function pointer" that lets us get the configuration of a given job manager
type GetConfigFn func() error

// Config captures some system configuration aspects that are necessary
// to run experiments
type Config struct {
	// HostDistro is the Linux distribution on the host
	HostDistro string

	// BinPath is the path to the current binary
	BinPath string

	// CurPath is the current path
	CurPath string

	// SedBin is the path to the sed binary
	SedBin string

	// Verbose mode is active/inactive
	Verbose bool

	// Debug mode is active/inactive
	Debug bool

	// SudoBin is the path to sudo on the host
	SudoBin string

	// Network holds details about the network configuration
	Network network.Info
}

// ParseDistroID parses the string we use to identify a specific distro into a distribution name and its version
func ParseDistroID(distro string) (string, string) {
	if !strings.Contains(distro, ":") {
		log.Printf("[WARN] %s an invalid distro ID\n", distro)
		return "", ""
	}

	tokens := strings.Split(distro, ":")
	if len(tokens) != 2 {
		log.Printf("[WARN] %s an invalid distro ID\n", distro)
		return "", ""
	}

	return tokens[0], tokens[1]
}

// GetDistroID returns a formatted version of the value of TargetDistro.
//
// This is mainly used to have a standard way to set directory and file names
func GetDistroID(distro string) string {
	return strings.Replace(distro, ":", "_", 1)
}

// CompatibleArch checks whether the local architecture is compatible with a list of architectures.
//
// The list of architectures is for example the output of sy.GetSIFArchs()
func CompatibleArch(list []string) bool {
	for _, arch := range list {
		if arch == runtime.GOARCH {
			return true
		}
	}
	return false
}

func Load() (*Config, error) {
	cfg := new(Config)

	/* Figure out the directory of this binary */
	bin, err := os.Executable()
	if err != nil {
		return cfg, fmt.Errorf("cannot detect the directory of the binary")
	}
	cfg.BinPath = filepath.Dir(bin)
	cfg.CurPath, err = os.Getwd()
	if err != nil {
		return cfg, fmt.Errorf("cannot detect current directory")
	}

	cfg.SudoBin, err = exec.LookPath("sudo")
	if err != nil {
		return cfg, fmt.Errorf("sudo not available: %s", err)
	}

	cfg.SedBin, err = exec.LookPath("sed")
	if err != nil {
		return cfg, fmt.Errorf("sed not available: %s", err)
	}

	// Load the network configuration
	cfg.Network = network.Detect()

	return cfg, nil
}

func (cfg *Config) Display() {
	fmt.Println("System configuration:")
	fmt.Printf("\tPath to the current binary: %s\n", cfg.BinPath)
	fmt.Printf("\tCurrent path: %s\n", cfg.CurPath)
	fmt.Printf("\tPath to the sed binary: %s\n", cfg.SedBin)
	fmt.Printf("\tPath to the sudo binary: %s\n", cfg.SudoBin)

	cfg.Network.Display()
}
