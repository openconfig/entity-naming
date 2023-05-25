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

	"github.com/openconfig/entity-naming/internal/namer"
)

const fakeVendor = Vendor("fake")

var devParams = &DeviceParams{Vendor: fakeVendor}

func TestLoopbackInterface(t *testing.T) {
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

func TestAggregatePort(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		const want = "fakeAggregatePort0"
		vendorToNamer[fakeVendor] = &fakeNamer{AggregatePortFn: func(int) (string, error) {
			return want, nil
		}}
		got, err := AggregatePort(devParams, 0)
		if err != nil {
			t.Errorf("AggregatePort(%v,0) got error %v", devParams, err)
		}
		if got != want {
			t.Errorf("AggregatePort(%v,0) got %q, want %q", devParams, got, want)
		}
	})

	t.Run("negative index", func(t *testing.T) {
		vendorToNamer[fakeVendor] = &fakeNamer{AggregatePortFn: func(int) (string, error) {
			return "", nil
		}}
		_, err := AggregatePort(devParams, -1)
		if wantErr := "negative"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Errorf("AggregatePort(%v,0) got error %v, want substring %q", devParams, err, wantErr)
		}
	})

	t.Run("error", func(t *testing.T) {
		const wantErr = "fakeAggregatePortErr"
		vendorToNamer[fakeVendor] = &fakeNamer{AggregatePortFn: func(int) (string, error) {
			return "", errors.New(wantErr)
		}}
		_, err := AggregatePort(devParams, 0)
		if err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Errorf("AggregatePort(%v,0) got error %v, want substring %q", devParams, err, wantErr)
		}
	})
}

func TestAggregateInterface(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		const want = "fakeAggregateIntf0"
		vendorToNamer[fakeVendor] = &fakeNamer{AggregateInterfaceFn: func(int) (string, error) {
			return want, nil
		}}
		got, err := AggregateInterface(devParams, 0)
		if err != nil {
			t.Errorf("AggregateInterface(%v,0) got error %v", devParams, err)
		}
		if got != want {
			t.Errorf("AggregateInterface(%v,0) got %q, want %q", devParams, got, want)
		}
	})

	t.Run("negative index", func(t *testing.T) {
		vendorToNamer[fakeVendor] = &fakeNamer{AggregateInterfaceFn: func(int) (string, error) {
			return "", nil
		}}
		_, err := AggregateInterface(devParams, -1)
		if wantErr := "negative"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Errorf("AggregateInterface(%v,0) got error %v, want substring %q", devParams, err, wantErr)
		}
	})

	t.Run("error", func(t *testing.T) {
		const wantErr = "fakeAggregateIntfErr"
		vendorToNamer[fakeVendor] = &fakeNamer{AggregateInterfaceFn: func(int) (string, error) {
			return "", errors.New(wantErr)
		}}
		_, err := AggregateInterface(devParams, 0)
		if err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Errorf("AggregateInterface(%v,0) got error %v, want substring %q", devParams, err, wantErr)
		}
	})
}

var _ namer.Namer = (*fakeNamer)(nil)

type fakeNamer struct {
	LoopbackInterfaceFn  func(int) (string, error)
	AggregatePortFn      func(int) (string, error)
	AggregateInterfaceFn func(int) (string, error)
}

func (fn *fakeNamer) LoopbackInterface(index int) (string, error) {
	return fn.LoopbackInterfaceFn(index)
}

func (fn *fakeNamer) AggregatePort(index int) (string, error) {
	return fn.AggregatePortFn(index)
}

func (fn *fakeNamer) AggregateInterface(index int) (string, error) {
	return fn.AggregateInterfaceFn(index)
}
