package build

import (
	"os/exec"
)

func Build() {
	// Build the Android APK
	cmd := exec.Command("android/ndk-build", "NDK_DEBUG=1", "NDK_DEBUG=1")
	cmd.Run()
}
