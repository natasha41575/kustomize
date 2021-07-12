// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package main_test

import (
	"testing"

	kusttest_test "sigs.k8s.io/kustomize/api/testutils/kusttest"
)

func TestBashedConfigMapPlugin(t *testing.T) {
	th := kusttest_test.MakeEnhancedHarness(t).
		PrepExecPlugin("someteam.example.com", "v1", "BashedConfigMap")
	defer th.Reset()

	m := th.LoadAndRunGenerator(`
apiVersion: someteam.example.com/v1
kind: BashedConfigMap
metadata:
  name: whatever
argsOneLiner: alice myMomsMaidenName
`)
	th.AssertActualEqualsExpected(m, `
apiVersion: v1
data:
  password: myMomsMaidenName
  username: alice
kind: ConfigMap
metadata:
  name: example-configmap-test
`)
}
