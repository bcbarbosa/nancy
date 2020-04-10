//
// Copyright 2018-present Sonatype Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestBadArgs(t *testing.T) {
	var err error
	cmd := exec.Command("./nancy", "-bad")
	out, err := cmd.CombinedOutput()
	sout := string(out) // because out is []byte
	if err != nil && !strings.Contains(sout, "flag provided but not defined: -bad") {
		fmt.Println(sout) // so we can see the full output
		t.Errorf("%v", err)
	}
}
