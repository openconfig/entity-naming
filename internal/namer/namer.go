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

// Package namer provides a namer strategy interface for naming entities.
package namer

// Namer is a strategy interface for naming entities.
type Namer interface {
	// LoopbackInterface returns the name of the loopback interface with the
	// specified zero-based index, or an error if no such name exists.
	// This method will never be called with a negative index.
	LoopbackInterface(index int) (string, error)

	// AggregatePort returns the name of the aggregate port with the specified
	// zero-based index, or an error if no such name exists.
	// This method will never be called with a negative index.
	AggregatePort(index int) (string, error)

	// AggregateInterface returns the name of the interface bound to the aggregate
	// port with the specified zero-based index, or an error if no such name
	// exists. This method will never be called with a negative index.
	AggregateInterface(index int) (string, error)

	// Linecard returns the name of the linecard component with the specified
	// zero-based index, or an error if no such name exists.
	// This method will never be called with a negative index.
	Linecard(index int) (string, error)

	// ControllerCard returns the name of the controller card component with the
	// specified zero-based index, or an error if no such name exists.
	// This method will never be called with a negative index.
	ControllerCard(index int) (string, error)

	// Fabric returns the name of the fabric component with the specified
	// zero-based index, or an error if no such name exists.
	// This method will never be called with a negative index.
	Fabric(index int) (string, error)
}
