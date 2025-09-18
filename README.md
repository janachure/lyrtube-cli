# lyrtube-cli

```
██╗  ██╗   ██╗██████╗ ████████╗██╗   ██╗██████╗ ███████╗
██║  ╚██╗ ██╔╝██╔══██╗╚══██╔══╝██║   ██║██╔══██╗██╔════╝
██║   ╚████╔╝ ██████╔╝   ██║   ██║   ██║██████╔╝█████╗  
██║    ╚██╔╝  ██╔══██╗   ██║   ██║   ██║██╔══██╗██╔══╝  
███████╗██║   ██║  ██║   ██║   ╚██████╔╝██████╔╝███████╗
╚══════╝╚═╝   ╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚═════╝ ╚══════╝
   🎵 A powerful song downloader built in Go 🎵
```

**lyrtube-cli** is a powerful command-line audio downloader built in Go, powered by [yt-dlp](https://github.com/yt-dlp/yt-dlp) and [lrclib.net](https://lrclib.net). It allows you to download high-quality audio from YouTube and other supported platforms with automatic lyrics embedding.

## ⚠️ Prerequisites

Before using lyrtube-cli, make sure you have the following installed on your system:

- **Python 3** - Required by yt-dlp
- **FFmpeg** - Required for audio processing and lyrics embedding

### Installing Prerequisites

**macOS (using Homebrew):**
```bash
brew install python3 ffmpeg
```

**Ubuntu/Debian:**
```bash
sudo apt update
sudo apt install python3 ffmpeg
```

**Windows:**
- Download Python 3 from [python.org](https://www.python.org/downloads/)
- Download FFmpeg from [ffmpeg.org](https://ffmpeg.org/download.html)

## ✨ Features

- 🎵 **High-quality audio downloads** from YouTube and other platforms
- 📝 **Automatic lyrics fetching** from lrclib.net
- 🎯 **Multiple lyrics modes**: embed into audio files or save as .lrc files
- 📁 **Playlist support** with individual track processing
- ⚙️ **Configurable settings** for output directory, audio format, and quality
- 🎨 **Beautiful terminal UI** with progress indicators
- 🖼️ **Thumbnail embedding** into audio files
- 📊 **Metadata extraction** and embedding

## 🚀 Installation

### Option 1: Download Pre-built Binary

1. Go to the [Releases](https://github.com/cesp99/lyrtube-cli/releases) page
2. Download the appropriate binary for your operating system
3. Make it executable (Linux/macOS): `chmod +x lyrtube-cli`
4. Move to your PATH or run directly

### Option 2: Build from Source

**Prerequisites:**
- Go 1.24.5 or later

**Steps:**
```bash
# Clone the repository
git clone https://github.com/cesp99/lyrtube-cli.git
cd lyrtube-cli

# Build the application
go build -o lyrtube-cli

# (Optional) Install globally
sudo mv lyrtube-cli /usr/local/bin/
```

## 📖 Usage

### First Run Setup

On your first run, lyrtube-cli will guide you through an interactive setup wizard:

```bash
./lyrtube-cli download "https://www.youtube.com/watch?v=example"
```

The wizard will ask you to configure:
- **Output directory** - Where to save downloaded files
- **Audio format** - mp3, m4a, flac, etc.
- **Audio quality** - 128k, 192k, 320k, best
- **Lyrics mode** - none, lrc (separate file), embedded (into audio)

### Basic Commands

#### Download a Single Video
```bash
./lyrtube-cli download "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
```

#### Download a Playlist
```bash
./lyrtube-cli download "https://www.youtube.com/playlist?list=example"
```

#### Download with Custom Options
```bash
./lyrtube-cli download "https://www.youtube.com/watch?v=example" \
  --out "/path/to/output" \
  --format "mp3" \
  --quality "320k"
```

### Configuration Management

#### View Current Configuration
```bash
./lyrtube-cli config show
```

#### Edit Configuration
```bash
./lyrtube-cli config edit
```

### Command Line Options

| Flag | Short | Description | Example |
|------|-------|-------------|---------|
| `--out` | `-o` | Output directory | `-o /path/to/music` |
| `--format` | `-f` | Audio format | `-f mp3` |
| `--quality` | `-q` | Audio quality | `-q 320k` |

## ⚙️ Configuration

The configuration file is stored at:
- **Linux/macOS**: `~/.lyrtube-cli/config.yaml`
- **Windows**: `%USERPROFILE%\.lyrtube-cli\config.yaml`

### Configuration Options

```yaml
output_dir: "~/Music/lyrtube-cli"
audio_format: "mp3"
audio_quality: "192k"
lyrics_mode: "embedded"  # none, lrc, embedded
```

### Lyrics Modes

- **`none`** - Don't download lyrics
- **`lrc`** - Save lyrics as separate .lrc files
- **`embedded`** - Embed lyrics directly into audio files (recommended)

## 🛠️ Development

### Project Structure

```
lyrtube-cli/
├── cmd/                 # CLI commands
│   ├── download.go     # Download command
│   ├── config.go       # Configuration commands
│   └── root.go         # Root command
├── internal/
│   ├── config/         # Configuration management
│   ├── downloader/     # Audio downloading logic
│   ├── embedded/       # Embedded yt-dlp binary
│   ├── lyrics/         # Lyrics fetching and processing
│   ├── processor/      # Lyrics processing workflow
│   └── ui/             # Terminal UI components
└── main.go             # Application entry point
```

### Building

```bash
# Build for current platform
go build -o lyrtube-cli

# Build for multiple platforms
GOOS=linux GOARCH=amd64 go build -o lyrtube-cli-linux-amd64
GOOS=windows GOARCH=amd64 go build -o lyrtube-cli-windows-amd64.exe
GOOS=darwin GOARCH=amd64 go build -o lyrtube-cli-darwin-amd64
```


## 🤝 Contributing

We welcome contributions! Here's how you can help:

1. **Fork the repository**
2. **Create a feature branch**: `git checkout -b feature/amazing-feature`
3. **Make your changes** and add tests if applicable
4. **Commit your changes**: `git commit -m 'Add amazing feature'`
5. **Push to the branch**: `git push origin feature/amazing-feature`
6. **Open a Pull Request**

### Contribution Guidelines

- Follow Go best practices and conventions
- Add tests for new functionality
- Update documentation as needed

## 📄 License

This project is licensed under the **GPL-3.0 License** - see the [LICENSE](LICENSE) file for details.

Everyone is free to contribute, modify, and distribute this software under the terms of the GPL-3.0 license.

## 🙏 Acknowledgments

- [yt-dlp](https://github.com/yt-dlp/yt-dlp) - The powerful media downloader that powers this tool
- [lrclib.net](https://lrclib.net) - Free lyrics database API
- [Cobra](https://github.com/spf13/cobra) - CLI framework for Go
- [FFmpeg](https://ffmpeg.org) - Media processing toolkit

## 📞 Support

If you encounter any issues or have questions:

1. Check the [Issues](https://github.com/cesp99/lyrtube-cli/issues) page
2. Create a new issue if your problem isn't already reported
3. Provide detailed information about your system and the error

---

**Made with ❤️ and Go**