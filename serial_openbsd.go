//
// Copyright 2014-2021 Cristian Maglie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package serial

import "golang.org/x/sys/unix"

const devFolder = "/dev"
const regexFilter = "^(cu|tty)\\..*"

// termios manipulation functions

var databitsMap = map[int]uint32{
	0: unix.CS8, // Default to 8 bits
	5: unix.CS5,
	6: unix.CS6,
	7: unix.CS7,
	8: unix.CS8,
}

const tcCMSPAR uint32 = 0 // may be CMSPAR or PAREXT
const tcIUCLC uint32 = 0

const tcCCTS_OFLOW uint32 = 0x00010000
const tcCRTS_IFLOW uint32 = 0x00020000

const tcCRTSCTS uint32 = tcCCTS_OFLOW

const ioctlTcgetattr = unix.TIOCGETA
const ioctlTcsetattr = unix.TIOCSETA

func tcFlush(handle int, action int) error {
     return unix.IoctlSetPointerInt(handle, unix.TIOCFLUSH, action)
}

func toTermiosSpeedType(speed int) int32 {
	return int32(speed)
}

func setTermSettingsBaudrate(speed int, settings *unix.Termios) (error, bool) {
	if speed == 0 {
		speed = 9600
	}
	settings.Ispeed = toTermiosSpeedType(speed)
	settings.Ospeed = toTermiosSpeedType(speed)
	return nil, false
}

func (port *unixPort) setSpecialBaudrate(speed uint32) error {
	// TODO: unimplemented
	return &PortError{code: InvalidSpeed}
}
