#!/bin/bash
#
# Copyright 2023 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -eu

cd oc
git clone https://github.com/openconfig/public.git
go install github.com/openconfig/ygot/generator@latest

generator \
	-output_file=oc.go \
	-exclude_modules=ietf-interfaces \
	-compress_paths=true \
	-trim_enum_openconfig_prefix \
	-package_name=oc \
	-path=public \
	public/release/models/optical-transport/openconfig-terminal-device.yang \
	public/release/models/optical-transport/openconfig-transport-types.yang

goimports -w oc.go
gofmt -w -s oc.go

rm -rf public
