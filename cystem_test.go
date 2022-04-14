// Copyright (c) 2022 aiocat
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cystem

import "testing"

func TestRunString(t *testing.T) {
	RunString("echo Test")
}

func TestRunStringSlice(t *testing.T) {
	RunStringSlice([]string{
		"echo",
		"Test",
		"Output",
	})
}
