# fps-game

Minimal Go program using raylib-go.

On openSUSE Tumbleweed, install the build dependencies:

```sh
sudo zypper install make gcc pkg-config wayland-devel libxkbcommon-devel
```

Use `make run`, `make build`, or `make test`. The Makefile supplies the
header paths from `pkg-config` automatically.
