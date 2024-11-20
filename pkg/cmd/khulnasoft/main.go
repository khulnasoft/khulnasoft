// Copyright 2016-2018, Pulumi Corporation.
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

package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"

	"github.com/khulnasoft/khulnasoft/pkg/v3/version"
)

// panicHandler displays an emergency error message to the user and a stack trace to
// report the panic.
//
// finished should be set to false when the handler is deferred and set to true as the
// last statement in the scope. This trick is necessary to avoid catching and then
// discarding a panic(nil).
func panicHandler(finished *bool) {
	if panicPayload := recover(); !*finished {
		stack := string(debug.Stack())
		fmt.Fprintln(os.Stderr, "================================================================================")
		fmt.Fprintln(os.Stderr, "The Pulumi CLI encountered a fatal error. This is a bug!")
		fmt.Fprintln(os.Stderr, "We would appreciate a report: https://github.com/khulnasoft/khulnasoft/issues/")
		fmt.Fprintln(os.Stderr, "Please provide all of the text below in your report.")
		fmt.Fprintln(os.Stderr, "================================================================================")
		fmt.Fprintf(os.Stderr, "Pulumi Version:   %s\n", version.Version)
		fmt.Fprintf(os.Stderr, "Go Version:       %s\n", runtime.Version())
		fmt.Fprintf(os.Stderr, "Go Compiler:      %s\n", runtime.Compiler)
		fmt.Fprintf(os.Stderr, "Architecture:     %s\n", runtime.GOARCH)
		fmt.Fprintf(os.Stderr, "Operating System: %s\n", runtime.GOOS)
		fmt.Fprintf(os.Stderr, "Panic:            %s\n\n", panicPayload)
		fmt.Fprintln(os.Stderr, stack)
		os.Exit(1)
	}
}

func main() {
	finished := new(bool)
	defer panicHandler(finished)

	if err := NewPulumiCmd().Execute(); err != nil {
		os.Exit(1)
	}
	*finished = true
}