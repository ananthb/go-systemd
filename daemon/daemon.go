// Copyright 2014 Docker, Inc.
// Copyright 2015-2018 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Package daemon provides a Go implementation of the sd_notify protocol.
// It can be used to inform systemd of service start-up completion, watchdog
// events, and other status changes.
//
// https://www.freedesktop.org/software/systemd/man/sd_notify.html#Description
package daemon

import (
	"errors"

	"golang.org/x/sys/unix"
)

// SdBooted returns true if the system was booted with systemd.
// See sd_booted(3) for more information.
func SdBooted() (bool, error) {
	if err := unix.Faccessat(unix.AT_FDCWD, "/run/systemd/systemd/", unix.F_OK, 0); err != nil {
		if errors.Is(err, unix.ENOENT) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
