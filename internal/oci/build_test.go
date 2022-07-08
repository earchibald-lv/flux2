/*
Copyright 2022 The Flux authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package oci

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"
)

func TestBuild(t *testing.T) {
	g := NewWithT(t)
	testDir := "./testdata/build"

	tmpDir := t.TempDir()
	artifactPath := filepath.Join(tmpDir, "files.tar.gz")

	// test with non-existent path
	err := Build(artifactPath, "testdata/non-existent")
	g.Expect(err).To(HaveOccurred())

	err = Build(artifactPath, testDir)
	g.Expect(err).ToNot(HaveOccurred())

	if _, err := os.Stat(artifactPath); err != nil {
		g.Expect(err).ToNot(HaveOccurred())
	}
}
