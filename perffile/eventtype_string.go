// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// generated by stringer -type=EventType; DO NOT EDIT

package perffile

import "fmt"

const _EventType_name = "EventTypeHardwareEventTypeSoftwareEventTypeTracepointEventTypeHWCacheEventTypeRawEventTypeBreakpoint"

var _EventType_index = [...]uint8{0, 17, 34, 53, 69, 81, 100}

func (i EventType) String() string {
	if i+1 >= EventType(len(_EventType_index)) {
		return fmt.Sprintf("EventType(%d)", i)
	}
	return _EventType_name[_EventType_index[i]:_EventType_index[i+1]]
}
