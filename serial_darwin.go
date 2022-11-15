//
// Copyright 2014-2021 Cristian Maglie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package serial

import "golang.org/x/sys/unix"

const devFolder = "/dev"
const regexFilter = "^(cu|tty)\\..*"

const ioctlTcgetattr = unix.TIOCGETA
const ioctlTcsetattr = unix.TIOCSETA

func tcFlush(handle int, action int) error {
     return unix.IoctlSetPointerInt(handle, unix.TIOCFLUSH, action)
}

func setTermSettingsBaudrate(speed int, settings *unix.Termios) (error, bool) {
	baudrate, ok := baudrateMap[speed]
	if !ok {
		return nil, true
	}
	settings.Ispeed = toTermiosSpeedType(baudrate)
	settings.Ospeed = toTermiosSpeedType(baudrate)
	return nil, false
}

func (port *unixPort) setSpecialBaudrate(speed uint32) error {
	const kIOSSIOSPEED = 0x80045402
	return unix.IoctlSetPointerInt(port.handle, kIOSSIOSPEED, int(speed))
}
