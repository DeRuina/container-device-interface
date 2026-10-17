/*
   Copyright © 2026 The CDI Authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package cdi

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	cdi "tags.cncf.io/container-device-interface/specs-go"
)

func TestDeepCopyDeviceNode(t *testing.T) {
	require.Nil(t, deepCopyDeviceNode(nil))

	mode := os.FileMode(0o660)
	uid, gid := uint32(1000), uint32(1001)
	orig := &cdi.DeviceNode{Path: "/dev/x", FileMode: &mode, UID: &uid, GID: &gid}

	c := deepCopyDeviceNode(orig)
	require.Equal(t, *orig, cdi.DeviceNode{Path: c.Path, FileMode: orig.FileMode, UID: orig.UID, GID: orig.GID})
	require.Equal(t, *orig.FileMode, *c.FileMode)
	require.Equal(t, *orig.UID, *c.UID)
	require.Equal(t, *orig.GID, *c.GID)

	// No pointer is shared: changing the copy leaves the original untouched.
	*c.FileMode, *c.UID, *c.GID = 0o600, 0, 0
	require.Equal(t, os.FileMode(0o660), *orig.FileMode)
	require.Equal(t, uint32(1000), *orig.UID)
	require.Equal(t, uint32(1001), *orig.GID)
}
