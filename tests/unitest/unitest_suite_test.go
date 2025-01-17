// SPDX-License-Identifier: Apache-2.0
// Copyright 2022 Authors of KubeArmor

package unitest_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUnitest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Unitest Suite")
}
