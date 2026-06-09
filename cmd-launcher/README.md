# CMD Launcher

A lightweight wrapper that automatically and transparently launches a companion `.cmd` script sharing the same name and directory in the background.

For example:

- Rename the executable to `launch.exe`, and double-clicking it will run `launch.cmd`.
- Rename the executable to `start.exe`, and double-clicking it will run `start.cmd`.

## Build with Custom Icon

1. Drop your icon as `./app.ico`.
2. Execute the shell script below:

```shell
go install github.com/akavel/rsrc@latest
rsrc -arch amd64 -ico app.ico -o rsrc_windows_amd64.syso
rsrc -arch arm64 -ico app.ico -o rsrc_windows_arm64.syso
make build
```
