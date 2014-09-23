// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !plan9,!solaris

// Package fsnotify provides a platform-independent interface for file system notifications.
package fsnotify

import "fmt"

// Event represents a single file system notification.
type Event struct {
	Name string // Relative path to the file or directory.
	Op   Op     // File operation that triggered the event.
}

// Op describes a set of file operations.
type Op uint32

// These are the generalized file operations that can trigger a notification.
const (
	Create Op = 1 << iota
	Write
	Remove
	Rename
	Chmod
)

// NewEvent new *Event
// for user generated custom Event
// example
/*
e := NewEvent(Create|Write, "/tmp/newdir")
*/
func NewEvent(op Op, name string) Event {
	return Event{
		Name: name,
		Op:   op,
	}
}

// String returns a string representation of the event in the form
// "file: REMOVE|WRITE|..."
func (e Event) String() string {
	events := ""

	if e.Op&Create == Create {
		events += "|CREATE"
	}
	if e.Op&Remove == Remove {
		events += "|REMOVE"
	}
	if e.Op&Write == Write {
		events += "|WRITE"
	}
	if e.Op&Rename == Rename {
		events += "|RENAME"
	}
	if e.Op&Chmod == Chmod {
		events += "|CHMOD"
	}

	if len(events) > 0 {
		events = events[1:]
	}

	return fmt.Sprintf("%q: %s", e.Name, events)
}

// Watch a given file path
func (w *Watcher) Watch(path string) error {
	return w.Add(path)
}

// Remove a watch on a file
func (w *Watcher) RemoveWatch(path string) error {
	return w.Remove(path)
}

// OpVar return e.Op as uint32
func (e *Event) OpVar() uint32 {
	return uint32(e.Op)
}

//
// Count return watching path count
func (w *Watcher) Count() int {
	return len(w.watches)
}

// IsDelete
func (e Event) IsDelete() bool {
	if e.Op&Remove == Remove {
		return true
	}
	return false
}

// IsChmod
func (e Event) IsChmod() bool {
	if e.Op&Chmod == Chmod {
		return true
	}
	return false
}

// IsModify
func (e Event) IsModify() bool {
	if e.IsRename() || e.IsWrite() || e.IsCreate() {
		return true
	}
	return false
}

// IsRename
func (e Event) IsRename() bool {
	if e.Op&Rename == Rename {
		return true
	}
	return false
}

// IsWrite
func (e Event) IsWrite() bool {
	if e.Op&Write == Write {
		return true
	}
	return false
}

// IsCreate
func (e Event) IsCreate() bool {
	if e.Op&Create == Create {
		return true
	}
	return false
}
