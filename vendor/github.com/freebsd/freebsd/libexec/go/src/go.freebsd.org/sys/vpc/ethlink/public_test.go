// Test VPC L2 EthLink objects.
//
// SPDX-License-Identifier: BSD-2-Clause-FreeBSD
//
// Copyright (C) 2018 Sean Chittenden <seanc@joyent.com>
// Copyright (c) 2018 Joyent, Inc.
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE AUTHOR AND CONTRIBUTORS ``AS IS'' AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED.  IN NO EVENT SHALL THE AUTHOR OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
// OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
// HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
// LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
// OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
// SUCH DAMAGE.

package ethlink_test

import (
	"syscall"
	"testing"

	"github.com/freebsd/freebsd/libexec/go/src/go.freebsd.org/sys/vpc"
	"github.com/freebsd/freebsd/libexec/go/src/go.freebsd.org/sys/vpc/ethlink"
	"github.com/freebsd/freebsd/libexec/go/src/go.freebsd.org/sys/vpc/vpctest"
	"github.com/sean-/seed"
)

func init() {
	seed.MustInit()
}

// TestEthLink_CreateCommitDestroy is intended to verify the basic lifecycle
// functionality of a switch.
func TestEthLink_CreateCommitDestroy(t *testing.T) {
	var cfg ethlink.Config
	{
		cfg.ID = vpc.GenID(vpc.ObjTypeLinkEth)
		cfg.ID.Node = [6]byte{}
	}

	existingIfaces, err := vpctest.GetAllInterfaces()
	if err != nil {
		t.Fatalf("unable to get all interfaces")
	}

	func() { // Create + close switch w/o commit
		sw, err := ethlink.Create(cfg)
		if err != nil {
			t.Fatalf("unable to create switch: %v", err)
		}

		{ // Get the before/after
			ifacesAfterCreate, err := vpctest.GetAllInterfaces()
			if err != nil {
				t.Fatalf("unable to get all interfaces")
			}
			_, newIfaces, _ := existingIfaces.Difference(ifacesAfterCreate)
			if len(newIfaces) != 1 {
				t.Fatalf("one interface should have been added")
			}
		}

		// Close the descriptor, the switch should still exist
		if err := sw.Close(); err != nil {
			t.Fatalf("unable to close switch: %v", err)
		}

		{ // Make sure the iface is still available
			ifacesAfterClose, err := vpctest.GetAllInterfaces()
			if err != nil {
				t.Fatalf("unable to get all interfaces")
			}
			_, n, _ := existingIfaces.Difference(ifacesAfterClose)
			if len(n) != 0 {
				t.Fatalf("no new interfaces should have been created")
			}
		}
	}()

	func() { // Create switch scope
		sw, err := ethlink.Create(cfg)
		if err != nil {
			t.Fatalf("unable to create switch: %v", err)
		}

		{ // Get the before/after
			ifacesAfterCreate, err := vpctest.GetAllInterfaces()
			if err != nil {
				t.Fatalf("unable to get all interfaces")
			}
			_, newIfaces, _ := existingIfaces.Difference(ifacesAfterCreate)
			if len(newIfaces) != 1 {
				t.Fatalf("one interface should have been added")
			}
		}

		// Attempt to be nice and clean up after ourselves in the event of a failure.
		defer func() {
			if err := sw.Destroy(); err != nil {
				t.Fatalf("unable to destroy switch in defer: %v", err)
			}

			if err := sw.Close(); err != nil {
				t.Fatalf("unable to close switch in defer: %v", err)
			}
		}()

		// Commit the switch
		if err := sw.Commit(); err != nil {
			t.Fatalf("unable to commit switch: %v", err)
		}

		// Close the descriptor, the switch should still exist
		if err := sw.Close(); err != nil {
			t.Fatalf("unable to close switch: %v", err)
		}

		{ // Make sure the iface is still available
			ifacesAfterClose, err := vpctest.GetAllInterfaces()
			if err != nil {
				t.Fatalf("unable to get all interfaces")
			}
			_, newIfaces, _ := existingIfaces.Difference(ifacesAfterClose)
			if len(newIfaces) != 1 {
				t.Fatalf("one interface should have persisted added")
			}
		}
	}()

	func() { // Open + Close switch scope
		sw, err := ethlink.Open(cfg)
		if err != nil {
			t.Fatalf("unable to open switch: %v", err)
		}

		// Attempt to be nice and clean up after ourselves in the event of a failure.
		defer func() {
			if err := sw.Destroy(); err != nil {
				t.Fatalf("unable to destroy switch in defer: %v", err)
			}

			if err := sw.Close(); err != nil {
				t.Fatalf("unable to close switch in defer: %v", err)
			}
		}()

		if err := sw.Close(); err != nil {
			t.Fatalf("unable to close switch: %v", err)
		}
	}()

	{ // Make sure the iface is still available
		ifacesAfterOpenClose, err := vpctest.GetAllInterfaces()
		if err != nil {
			t.Fatalf("unable to get all interfaces")
		}
		o, n, _ := existingIfaces.Difference(ifacesAfterOpenClose)
		if len(o) != 0 || len(n) != 0 {
			t.Fatalf("no interfaces should have been added or removed: %d/%d", len(o), len(n))
		}
	}

	func() { // Open + Destroy switch scope
		sw, err := ethlink.Open(cfg)
		if err != nil {
			t.Fatalf("unable to open switch: %v", err)
		}

		// Destroy
		if err := sw.Destroy(); err != nil {
			t.Fatalf("unable to destroy switch: %v", err)
		}

		if err := sw.Close(); err != nil {
			t.Fatalf("unable to close switch: %v", err)
		}
	}()

	{ // Make sure the iface is still available
		ifacesAfterDestroy, err := vpctest.GetAllInterfaces()
		if err != nil {
			t.Fatalf("unable to get all interfaces")
		}
		o, n, _ := existingIfaces.Difference(ifacesAfterDestroy)
		if len(o) != 0 || len(n) != 0 {
			t.Fatalf("interface count didn't return to original values")
		}
	}
}

func TestEthLink_CreateClose(t *testing.T) {
	existingIfaces, err := vpctest.GetAllInterfaces()
	if err != nil {
		t.Fatalf("unable to get existing interfaces")
	}

	var cfg ethlink.Config
	{
		cfg.ID = vpc.GenID(vpc.ObjTypeLinkEth)
		cfg.ID.Node = [6]byte{}
	}

	sw, err := ethlink.Create(cfg)
	if err != nil {
		t.Fatalf("unable to create switch: %v", err)
	}

	// Get the ifaces after create
	ifacesAfterCreate, err := vpctest.GetAllInterfaces()
	if err != nil {
		t.Fatalf("unable to get all interfaces")
	}
	_, newIfaces, _ := existingIfaces.Difference(ifacesAfterCreate)
	if len(newIfaces) != 1 {
		t.Fatalf("one VPC Switch should have been added")
	}

	if err := sw.Close(); err != nil {
		t.Fatalf("unable to close switch: %v", err)
	}

	ifacesAfterClose, err := vpctest.GetAllInterfaces()
	if err != nil {
		t.Fatalf("unable to get all interfaces")
	}
	o, n, _ := existingIfaces.Difference(ifacesAfterClose)
	if len(o) != 0 || len(n) != 0 {
		t.Fatalf("no interfaces should have been added or removed: %d/%d", len(o), len(n))
	}
}

func TestEthLink_CreateDestroyClose(t *testing.T) {
	existingIfaces, err := vpctest.GetAllInterfaces()
	if err != nil {
		t.Fatalf("unable to get existing interfaces")
	}

	var cfg ethlink.Config
	{
		cfg.ID = vpc.GenID(vpc.ObjTypeLinkEth)
		cfg.ID.Node = [6]byte{}
	}

	sw, err := ethlink.Create(cfg)
	if err != nil {
		t.Fatalf("unable to create switch: %v", err)
	}

	// Get the ifaces after create
	ifacesAfterCreate, err := vpctest.GetAllInterfaces()
	if err != nil {
		t.Fatalf("unable to get all interfaces")
	}
	_, newIfaces, _ := existingIfaces.Difference(ifacesAfterCreate)
	if len(newIfaces) != 1 {
		t.Fatalf("one VPC Switch should have been added")
	}

	// Destroy
	if err := sw.Destroy(); err != nil {
		if errno, ok := err.(syscall.Errno); ok && errno != 35 {
			t.Fatalf("unable to destroy switch: %v", err)
		}
	}

	if err := sw.Close(); err != nil {
		t.Fatalf("unable to close switch: %v", err)
	}

	ifacesAfterClose, err := vpctest.GetAllInterfaces()
	if err != nil {
		t.Fatalf("unable to get all interfaces")
	}
	o, n, _ := existingIfaces.Difference(ifacesAfterClose)
	if len(o) != 0 || len(n) != 0 {
		t.Fatalf("no interfaces should have been added or removed: %d/%d", len(o), len(n))
	}
}

func TestEthLink_CreateCommitDestroyClose(t *testing.T) {
	existingIfaces, err := vpctest.GetAllInterfaces()
	if err != nil {
		t.Fatalf("unable to get existing interfaces")
	}

	var cfg ethlink.Config
	{
		cfg.ID = vpc.GenID(vpc.ObjTypeLinkEth)
		cfg.ID.Node = [6]byte{}
	}

	sw, err := ethlink.Create(cfg)
	if err != nil {
		t.Fatalf("unable to create switch: %v", err)
	}

	// Get the ifaces after create
	ifacesAfterCreate, err := vpctest.GetAllInterfaces()
	if err != nil {
		t.Fatalf("unable to get all interfaces")
	}
	_, newIfaces, _ := existingIfaces.Difference(ifacesAfterCreate)
	if len(newIfaces) != 1 {
		t.Fatalf("one VPC Switch should have been added")
	}

	// Commit the switch
	if err := sw.Commit(); err != nil {
		t.Fatalf("unable to commit switch: %v", err)
	}

	// Destroy
	if err := sw.Destroy(); err != nil {
		t.Fatalf("unable to destroy switch: %v", err)
	}

	if err := sw.Close(); err != nil {
		t.Fatalf("unable to close switch: %v", err)
	}

	ifacesAfterClose, err := vpctest.GetAllInterfaces()
	if err != nil {
		t.Fatalf("unable to get all interfaces")
	}
	o, n, _ := existingIfaces.Difference(ifacesAfterClose)
	if len(o) != 0 || len(n) != 0 {
		t.Fatalf("no interfaces should have been added or removed: %d/%d", len(o), len(n))
	}
}
