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

package juniper

import (
	"strings"
	"testing"
)

var jn = new(Namer)

func TestLoopbackInterface(t *testing.T) {
	tests := []struct {
		desc  string
		index int
		want  string
	}{{
		desc:  "min",
		index: 0,
		want:  "lo0.0",
	}, {
		desc:  "max",
		index: 16000,
		want:  "lo0.16000",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := jn.LoopbackInterface(test.index)
			if err != nil {
				t.Fatalf("LoopbackInterface(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("LoopbackInterface(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := jn.LoopbackInterface(16001)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("LoopbackInterface(16001) got error %v, want substring %q", err, wantErr)
		}
	})
}

func TestAggregatePort(t *testing.T) {
	tests := []struct {
		desc  string
		index int
		want  string
	}{{
		desc:  "min",
		index: 0,
		want:  "ae0",
	}, {
		desc:  "max",
		index: 1151,
		want:  "ae1151",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := jn.AggregatePort(test.index)
			if err != nil {
				t.Fatalf("AggregatePort(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("AggregatePort(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := jn.AggregatePort(1152)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("AggregatePort(1152) got error %v, want substring %q", err, wantErr)
		}
	})
}

func TestAggregateInterface(t *testing.T) {
	tests := []struct {
		desc  string
		index int
		want  string
	}{{
		desc:  "min",
		index: 0,
		want:  "ae0.0",
	}, {
		desc:  "max",
		index: 1151,
		want:  "ae1151.0",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := jn.AggregateInterface(test.index)
			if err != nil {
				t.Fatalf("AggregateInterface(%v) got error: %v", test.index, err)
			}
			if got != test.want {
				t.Errorf("AggregateInterface(%d) got %q, want %q", test.index, got, test.want)
			}
		})
	}

	t.Run("over max", func(t *testing.T) {
		_, err := jn.AggregateInterface(1152)
		if wantErr := "exceed"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Fatalf("AggregateInterface(1152) got error %v, want substring %q", err, wantErr)
		}
	})
}
