// Copyright (c) 2018 HyperHQ Inc.
//
// SPDX-License-Identifier: Apache-2.0
//

package cache

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	vc "github.com/kata-containers/runtime/virtcontainers"
	"github.com/kata-containers/runtime/virtcontainers/factory/direct"
)

func TestTemplateFactory(t *testing.T) {
	assert := assert.New(t)

	testDir, _ := ioutil.TempDir("", "vmfactory-tmp-")
	hyperConfig := vc.HypervisorConfig{
		KernelPath: testDir,
		ImagePath:  testDir,
	}
	vmConfig := vc.VMConfig{
		HypervisorType:   vc.MockHypervisor,
		AgentType:        vc.NoopAgentType,
		HypervisorConfig: hyperConfig,
	}

	ctx := context.Background()

	// New
	f := New(ctx, 2, direct.New(ctx, vmConfig))

	// Config
	assert.Equal(f.Config(), vmConfig)

	// GetBaseVM
	_, err := f.GetBaseVM(ctx)
	assert.Nil(err)

	// CloseFactory
	f.CloseFactory(ctx)
}
