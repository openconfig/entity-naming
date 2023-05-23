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

package naming

import (
	"errors"
	"strings"
	"testing"
)

func TestLoopbackInterface(t *testing.T) {
	const fakeVendor = Vendor("fake")
	var devParams = &DeviceParams{Vendor: fakeVendor}

	t.Run("success", func(t *testing.T) {
		const want = "fakeLoopback0"
		vendorToNamer[fakeVendor] = &fakeNamer{LoopbackInterfaceFn: func(int) (string, error) {
			return want, nil
		}}
		got, err := LoopbackInterface(devParams, 0)
		if err != nil {
			t.Errorf("LoopbackInterface(%v,0) got error %v", devParams, err)
		}
		if got != want {
			t.Errorf("LoopbackInterface(%v,0) got %q, want %q", devParams, got, want)
		}
	})

	t.Run("negative index", func(t *testing.T) {
		vendorToNamer[fakeVendor] = &fakeNamer{LoopbackInterfaceFn: func(int) (string, error) {
			return "", nil
		}}
		_, err := LoopbackInterface(devParams, -1)
		if wantErr := "negative"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Errorf("LoopbackInterface(%v,0) got error %v, want substring %q", devParams, err, wantErr)
		}
	})

	t.Run("error", func(t *testing.T) {
		const wantErr = "fakeLoopbackErr"
		vendorToNamer[fakeVendor] = &fakeNamer{LoopbackInterfaceFn: func(int) (string, error) {
			return "", errors.New(wantErr)
		}}
		_, err := LoopbackInterface(devParams, 0)
		if err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Errorf("LoopbackInterface(%v,0) got error %v, want substring %q", devParams, err, wantErr)
		}
	})
}

type fakeNamer struct {
	LoopbackInterfaceFn func(int) (string, error)
}

func (fn *fakeNamer) LoopbackInterface(index int) (string, error) {
	return fn.LoopbackInterfaceFn(index)
}
