// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package vboxmanage provides a wrapper around vboxmanage.
package vboxmanage

import (
	"os"

	"github.com/siderolabs/go-cmd/pkg/cmd"
)

// ConvertFromRaw converts raw disk image to another format
func ConvertFromRaw(outputFmt, options, path string, printf func(string, ...any)) error {
	src := path + ".in"
	dest := path

	printf("converting raw disk image to %s", outputFmt)

	if err := os.Rename(path, src); err != nil {
		return err
	}

	if _, err := cmd.Run("VBoxManage", "convertfromraw", "--format", outputFmt, options, src, dest); err != nil {
		return err
	}

	return os.Remove(src)
}

// Resize an image.
func Resize(file, size string) error {
	if _, err := cmd.Run("VBoxManage", "modifyhd", file, "--resize", size); err != nil {
		return err
	}

	return nil
}
