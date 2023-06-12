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

package entname

import (
	"errors"
	"strings"
	"testing"

	"github.com/openconfig/entity-naming/internal/namer"
	"github.com/openconfig/entity-naming/oc"
)

const fakeVendor = Vendor("fake")

var devParams = &DeviceParams{Vendor: fakeVendor}

func TestLoopbackInterface(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		const want = "fakeLoopback0"
		setFakeNamer(&fakeNamer{LoopbackInterfaceFn: func(int) (string, error) {
			return want, nil
		}})
		got, err := LoopbackInterface(devParams, 0)
		if err != nil {
			t.Errorf("LoopbackInterface(%v,0) got error %v", devParams, err)
		}
		if got != want {
			t.Errorf("LoopbackInterface(%v,0) got %q, want %q", devParams, got, want)
		}
	})

	t.Run("negative index", func(t *testing.T) {
		setFakeNamer(&fakeNamer{LoopbackInterfaceFn: func(int) (string, error) {
			return "", nil
		}})
		_, err := LoopbackInterface(devParams, -1)
		if wantErr := "negative"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Errorf("LoopbackInterface(%v,0) got error %v, want substring %q", devParams, err, wantErr)
		}
	})

	t.Run("error", func(t *testing.T) {
		const wantErr = "fakeLoopbackErr"
		setFakeNamer(&fakeNamer{LoopbackInterfaceFn: func(int) (string, error) {
			return "", errors.New(wantErr)
		}})
		_, err := LoopbackInterface(devParams, 0)
		if err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Errorf("LoopbackInterface(%v,0) got error %v, want substring %q", devParams, err, wantErr)
		}
	})
}

func TestAggregateInterface(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		const want = "fakeAggregate0"
		setFakeNamer(&fakeNamer{AggregateInterfaceFn: func(int) (string, error) {
			return want, nil
		}})
		got, err := AggregateInterface(devParams, 0)
		if err != nil {
			t.Errorf("AggregateInterface(%v,0) got error %v", devParams, err)
		}
		if got != want {
			t.Errorf("AggregateInterface(%v,0) got %q, want %q", devParams, got, want)
		}
	})

	t.Run("negative index", func(t *testing.T) {
		setFakeNamer(&fakeNamer{AggregateInterfaceFn: func(int) (string, error) {
			return "", nil
		}})
		_, err := AggregateInterface(devParams, -1)
		if wantErr := "negative"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Errorf("AggregateInterface(%v,0) got error %v, want substring %q", devParams, err, wantErr)
		}
	})

	t.Run("error", func(t *testing.T) {
		const wantErr = "fakeAggregateErr"
		setFakeNamer(&fakeNamer{AggregateInterfaceFn: func(int) (string, error) {
			return "", errors.New(wantErr)
		}})
		_, err := AggregateInterface(devParams, 0)
		if err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Errorf("AggregateInterface(%v,0) got error %v, want substring %q", devParams, err, wantErr)
		}
	})
}

func TestAggregateMemberInterface(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		const want = "fakeAggregateMember0"
		setFakeNamer(&fakeNamer{AggregateMemberInterfaceFn: func(int) (string, error) {
			return want, nil
		}})
		got, err := AggregateMemberInterface(devParams, 0)
		if err != nil {
			t.Errorf("AggregateMemberInterface(%v,0) got error %v", devParams, err)
		}
		if got != want {
			t.Errorf("AggregateMemberInterface(%v,0) got %q, want %q", devParams, got, want)
		}
	})

	t.Run("negative index", func(t *testing.T) {
		setFakeNamer(&fakeNamer{AggregateMemberInterfaceFn: func(int) (string, error) {
			return "", nil
		}})
		_, err := AggregateMemberInterface(devParams, -1)
		if wantErr := "negative"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Errorf("AggregateMemberInterface(%v,0) got error %v, want substring %q", devParams, err, wantErr)
		}
	})

	t.Run("error", func(t *testing.T) {
		const wantErr = "fakeAggregateMemberErr"
		setFakeNamer(&fakeNamer{AggregateMemberInterfaceFn: func(int) (string, error) {
			return "", errors.New(wantErr)
		}})
		_, err := AggregateMemberInterface(devParams, 0)
		if err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Errorf("AggregateMemberInterface(%v,0) got error %v, want substring %q", devParams, err, wantErr)
		}
	})
}
func TestLinecard(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		const want = "fakeLinecard0"
		setFakeNamer(&fakeNamer{LinecardFn: func(int) (string, error) {
			return want, nil
		}})
		got, err := Linecard(devParams, 0)
		if err != nil {
			t.Errorf("Linecard(%v,0) got error %v", devParams, err)
		}
		if got != want {
			t.Errorf("Linecard(%v,0) got %q, want %q", devParams, got, want)
		}
	})

	t.Run("negative index", func(t *testing.T) {
		setFakeNamer(&fakeNamer{LinecardFn: func(int) (string, error) {
			return "", nil
		}})
		_, err := Linecard(devParams, -1)
		if wantErr := "negative"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Errorf("Linecard(%v,0) got error %v, want substring %q", devParams, err, wantErr)
		}
	})

	t.Run("error", func(t *testing.T) {
		const wantErr = "LinecardErr"
		setFakeNamer(&fakeNamer{LinecardFn: func(int) (string, error) {
			return "", errors.New(wantErr)
		}})
		_, err := Linecard(devParams, 0)
		if err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Errorf("Linecard(%v,0) got error %v, want substring %q", devParams, err, wantErr)
		}
	})
}

func TestControllerCard(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		const want = "fakeControllerCard0"
		setFakeNamer(&fakeNamer{ControllerCardFn: func(int) (string, error) {
			return want, nil
		}})
		got, err := ControllerCard(devParams, 0)
		if err != nil {
			t.Errorf("ControllerCard(%v,0) got error %v", devParams, err)
		}
		if got != want {
			t.Errorf("ControllerCard(%v,0) got %q, want %q", devParams, got, want)
		}
	})

	t.Run("negative index", func(t *testing.T) {
		setFakeNamer(&fakeNamer{ControllerCardFn: func(int) (string, error) {
			return "", nil
		}})
		_, err := ControllerCard(devParams, -1)
		if wantErr := "negative"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Errorf("ControllerCard(%v,0) got error %v, want substring %q", devParams, err, wantErr)
		}
	})

	t.Run("error", func(t *testing.T) {
		const wantErr = "ControllerCardErr"
		setFakeNamer(&fakeNamer{ControllerCardFn: func(int) (string, error) {
			return "", errors.New(wantErr)
		}})
		_, err := ControllerCard(devParams, 0)
		if err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Errorf("ControllerCard(%v,0) got error %v, want substring %q", devParams, err, wantErr)
		}
	})
}

func TestFabric(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		const want = "fakeFabric0"
		setFakeNamer(&fakeNamer{FabricFn: func(int) (string, error) {
			return want, nil
		}})
		got, err := Fabric(devParams, 0)
		if err != nil {
			t.Errorf("Fabric(%v,0) got error %v", devParams, err)
		}
		if got != want {
			t.Errorf("Fabric(%v,0) got %q, want %q", devParams, got, want)
		}
	})

	t.Run("negative index", func(t *testing.T) {
		setFakeNamer(&fakeNamer{FabricFn: func(int) (string, error) {
			return "", nil
		}})
		_, err := Fabric(devParams, -1)
		if wantErr := "negative"; err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Errorf("Fabric(%v,0) got error %v, want substring %q", devParams, err, wantErr)
		}
	})

	t.Run("error", func(t *testing.T) {
		const wantErr = "FabricErr"
		setFakeNamer(&fakeNamer{FabricFn: func(int) (string, error) {
			return "", errors.New(wantErr)
		}})
		_, err := Fabric(devParams, 0)
		if err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Errorf("Fabric(%v,0) got error %v, want substring %q", devParams, err, wantErr)
		}
	})
}

func TestPort(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		const want = "fakePort"
		setFakeNamer(&fakeNamer{
			PortFn: func(*namer.PortParams) (string, error) {
				return want, nil
			},
			IsFixedFormFactorFn: func() bool {
				return false
			},
		})
		got, err := Port(devParams, &PortParams{Speed: oc.IfEthernet_ETHERNET_SPEED_SPEED_1GB})
		if err != nil {
			t.Errorf("Port(%v,{}) got error %v", devParams, err)
		}
		if got != want {
			t.Errorf("Port(%v,{}) got %q, want %q", devParams, got, want)
		}
	})

	t.Run("success - fixed form factor", func(t *testing.T) {
		const want = "fakePortFFF"
		setFakeNamer(&fakeNamer{
			PortFn: func(*namer.PortParams) (string, error) {
				return want, nil
			},
			IsFixedFormFactorFn: func() bool {
				return true
			},
		})
		got, err := Port(devParams, &PortParams{Speed: oc.IfEthernet_ETHERNET_SPEED_SPEED_1GB})
		if err != nil {
			t.Errorf("Port(%v,{}) got error %v", devParams, err)
		}
		if got != want {
			t.Errorf("Port(%v,{}) got %q, want %q", devParams, got, want)
		}
	})

	t.Run("error", func(t *testing.T) {
		const wantErr = "PortErr"
		setFakeNamer(&fakeNamer{
			PortFn: func(*namer.PortParams) (string, error) {
				return "", errors.New(wantErr)
			},
			IsFixedFormFactorFn: func() bool {
				return false
			},
		})
		_, err := Port(devParams, &PortParams{Speed: oc.IfEthernet_ETHERNET_SPEED_SPEED_1GB})
		if err == nil || !strings.Contains(err.Error(), wantErr) {
			t.Errorf("Port(%v,{}) got error %v, want substring %q", devParams, err, wantErr)
		}
	})

	badParamsTests := []struct {
		desc       string
		portParams *PortParams
		fixedForm  bool
		wantErr    string
	}{{
		desc:       "negative slot",
		portParams: &PortParams{SlotIndex: -1, Speed: oc.IfEthernet_ETHERNET_SPEED_SPEED_1GB},
		wantErr:    "slot",
	}, {
		desc:       "negative pic",
		portParams: &PortParams{PICIndex: -2, Speed: oc.IfEthernet_ETHERNET_SPEED_SPEED_1GB},
		wantErr:    "pic",
	}, {
		desc:       "negative port",
		portParams: &PortParams{PortIndex: -3, Speed: oc.IfEthernet_ETHERNET_SPEED_SPEED_1GB},
		wantErr:    "port",
	}, {
		desc:       "negative channel",
		portParams: &PortParams{ChannelIndex: -4, Speed: oc.IfEthernet_ETHERNET_SPEED_SPEED_1GB},
		wantErr:    "channel",
	}, {
		desc:       "non-zero slot on fixed form factor",
		portParams: &PortParams{SlotIndex: 1, Speed: oc.IfEthernet_ETHERNET_SPEED_SPEED_1GB},
		fixedForm:  true,
		wantErr:    "non-zero slot",
	}, {
		desc:       "non-zero channel on unchannelized port",
		portParams: &PortParams{ChannelIndex: 1, Speed: oc.IfEthernet_ETHERNET_SPEED_SPEED_1GB},
		fixedForm:  true,
		wantErr:    "non-zero channel",
	}, {
		desc:       "non-zero channel on unchannelizable port",
		portParams: &PortParams{ChannelIndex: 1, ChannelState: Unchannelizable, Speed: oc.IfEthernet_ETHERNET_SPEED_SPEED_1GB},
		fixedForm:  true,
		wantErr:    "non-zero channel",
	}}
	for _, test := range badParamsTests {
		t.Run(test.desc, func(t *testing.T) {
			setFakeNamer(&fakeNamer{
				PortFn: func(*namer.PortParams) (string, error) {
					return "fakePort", nil
				},
				IsFixedFormFactorFn: func() bool {
					return test.fixedForm
				},
			})
			_, err := Port(devParams, test.portParams)
			if err == nil || !strings.Contains(err.Error(), test.wantErr) {
				t.Errorf("Port(%v,{}) got error %v, want substring %q", devParams, err, test.wantErr)
			}
		})
	}
}

func setFakeNamer(fn *fakeNamer) {
	namerFactories[fakeVendor] = func(string) namer.Namer { return fn }
}

var _ namer.Namer = (*fakeNamer)(nil)

type fakeNamer struct {
	LoopbackInterfaceFn, AggregateInterfaceFn, AggregateMemberInterfaceFn,
	LinecardFn, ControllerCardFn, FabricFn func(int) (string, error)
	PortFn              func(*namer.PortParams) (string, error)
	IsFixedFormFactorFn func() bool
}

func (fn *fakeNamer) LoopbackInterface(index int) (string, error) {
	return fn.LoopbackInterfaceFn(index)
}

func (fn *fakeNamer) AggregateInterface(index int) (string, error) {
	return fn.AggregateInterfaceFn(index)
}

func (fn *fakeNamer) AggregateMemberInterface(index int) (string, error) {
	return fn.AggregateMemberInterfaceFn(index)
}

func (fn *fakeNamer) Linecard(index int) (string, error) {
	return fn.LinecardFn(index)
}

func (fn *fakeNamer) ControllerCard(index int) (string, error) {
	return fn.ControllerCardFn(index)
}

func (fn *fakeNamer) Fabric(index int) (string, error) {
	return fn.FabricFn(index)
}

func (fn *fakeNamer) Port(pp *namer.PortParams) (string, error) {
	return fn.PortFn(pp)
}

func (fn *fakeNamer) IsFixedFormFactor() bool {
	return fn.IsFixedFormFactorFn()
}
