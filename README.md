# Systofi

[![Go Version](https://img.shields.io/badge/Go-1.25%2B-blue)](https://go.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A rofi-powered systemd service manager for Linux.

## Features

- Lists all systemd services
- Color-coded status indicators (active, failed, inactive)
- Context-aware action menu based on service state
- Supports: Start, Stop, Restart, Reload, Enable, Disable, Status
- Privileged operations via polkit (pkexec)

## Prerequisites

- Linux with systemd
- Go 1.25 or later
- [rofi](https://github.com/davatorium/rofi)
- polkit

### Install Dependencies

```bash
# Ubuntu/Debian
sudo apt install golang-go rofi polkitd-pkla

# Arch Linux
sudo pacman -S go rofi polkit
```

## Installation

### Binary Release

```bash
# Download latest release
wget https://github.com/natrium404/systofi/releases/latest/download/systofi-linux-amd64
chmod +x systofi-linux-amd64
sudo mv systofi-linux-amd64 /usr/local/bin/systofi
```

### From Source

```bash
go install github.com/natrium404/systofi@latest
```

Or clone and build:

```bash
git clone https://github.com/natrium404/systofi.git
cd systofi
go build -o systofi .
```

### Running

Run the binary:

```bash
systofi
```

**Bind to a key (Niri WM example)**:

```kdl
MOD+S hotkey-overlay-title="System Services" { spawn-sh "systofi"; }
```

## Configuration

Systofi uses the following environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `SYSTOFI_ROFI_THEME` | Path to rofi theme file | None (system default) |

Example:

```bash
export SYSTOFI_ROFI_THEME="$HOME/.config/rofi/main.rasi"
systofi
```

## Usage

1. **Launch**: Run `systofi` from your terminal or bind it to a keyboard shortcut
2. **Select Service**: Use arrow keys to browse, type to filter, Enter to select
3. **Choose Action**: Select the action you want to perform (Start, Stop, Restart, etc.)
4. **Confirm**: For privileged actions, authenticate via polkit prompt if needed

### Service List Icons

| Icon | Meaning |
|------|---------|
| <span style="color:#00FF88">●</span> | Active |
| <span style="color:#FF5555">✗</span> | Failed |
| <span style="color:#bbbbbb">○</span> | Inactive/Other |

### Available Actions

Actions shown depend on the current state of the service:

- **Active services**: Stop, Restart/Reload, Enable, Disable, Status
- **Failed services**: Start, Restart, Enable, Disable, Status
- **Inactive services**: Start, Enable, Disable, Status

## Security

Systofi uses `pkexec` (via polkit) to execute privileged systemctl commands. You will be prompted for authentication when performing actions that require root privileges (Start, Stop, Restart, Enable, Disable).

## Contributing

Contributions are welcome:

- Fork the repository
- Create a feature branch: `git checkout -b feature/amazing-feature`
- Commit your changes: `git commit -m 'Add amazing feature'`
- Push to the branch: `git push origin feature/amazing-feature`
- Open a Pull Request

### Development

Clone and setup:

```bash
git clone https://github.com/natrium404/systofi.git
cd systofi
go mod tidy
```

Build and test:

```bash
go build ./...
go test ./...
```

Run dev version:

```bash
go run .
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

Copyright (c) 2026 Natrium

## Acknowledgments

- [rofi](https://github.com/davatorium/rofi) - The application launcher and window manager
- [systemd](https://systemd.io/) - The system and service manager

## Fin

![Fin](https://i.pinimg.com/originals/eb/ec/d4/ebecd4010e549f33371d741d46b9b607.gif)
