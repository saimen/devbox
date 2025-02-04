// Copyright 2023 Jetpack Technologies Inc and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package envir

const (
	DevboxCache              = "DEVBOX_CACHE"
	devboxCLICloudShell      = "DEVBOX_CLI_CLOUD_SHELL"
	DevboxDebug              = "DEVBOX_DEBUG"
	DevboxFeaturePrefix      = "DEVBOX_FEATURE_"
	DevboxGateway            = "DEVBOX_GATEWAY"
	DevboxDoNotUpgradeConfig = "DEVBOX_DONT_UPGRADE_CONFIG"
	DevboxRegion             = "DEVBOX_REGION"
	DevboxSearchHost         = "DEVBOX_SEARCH_HOST"
	DevboxShellEnabled       = "DEVBOX_SHELL_ENABLED"
	DevboxShellStartTime     = "DEVBOX_SHELL_START_TIME"
	DevboxVM                 = "DEVBOX_VM"

	LauncherVersion = "LAUNCHER_VERSION"

	SSHTTY = "SSH_TTY"

	XDGDataHome   = "XDG_DATA_HOME"
	XDGConfigHome = "XDG_CONFIG_HOME"
	XDGCacheHome  = "XDG_CACHE_HOME"
	XDGStateHome  = "XDG_STATE_HOME"
)

// system
const (
	Env   = "ENV"
	Home  = "HOME"
	Path  = "PATH"
	Shell = "SHELL"
	User  = "USER"
)
