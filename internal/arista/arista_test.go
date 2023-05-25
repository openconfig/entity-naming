// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package arista

import (
	"strings"
	"testing"
)

func TestLoopbackInterface(t *testing.T) {
	var an = new(Namer)
	tests := []struct {
		desc  string
		index int
		want  string
	}{{
		desc:  "min",
		index: 0,
		want:  "Loopback0",
	}, {
		desc:  "max",
		index: 1000,
		want:  "Loopback1000",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := an.LoopbackInterface(test.index)
			if err != nil {
				t.Fatalf("LoopbackInterface(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("LoopbackInterface(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := an.LoopbackInterface(1001)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("LoopbackInterface(1001) got error %v, want substring %q", err, wantErr)
		}
	})
}
