module github.com/mrWinston/hyprworkspacemgr

go 1.21.5

require github.com/labi-le/hyprland-ipc-client v1.0.2

require github.com/stretchr/testify v1.8.4 // indirect

require (
	github.com/gotk3/gotk3 v0.6.2
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/samber/lo v1.39.0
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/cobra v1.8.0
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/exp v0.0.0-20220303212507-bbda1eaf7a17 // indirect
	golang.org/x/sys v0.13.0 // indirect
)

replace github.com/labi-le/hyprland-ipc-client => github.com/mrwinston/hyprland-ipc-client v1.0.4-0.20240228155123-7e28b73c668b
