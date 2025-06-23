# H2O - Comics and Images Web Server

## Welcome to H2O — Your Gateway to Comics & Creative Artworks

Discover a new way to experience visual storytelling with H2O, a platform built for comic lovers and art enthusiasts. Whether you're into manga, webtoons, graphic novels, or digital illustrations, our site is designed to offer a seamless and immersive reading experience — right from your browser.

🎨 **Browse Beautiful Artworks**
Explore a growing collection of high-quality illustrations and digital artwork from independent and aspiring artists. From fantasy to sci-fi, slice-of-life to surreal, there's something to captivate every visual taste.

📚 **Read Comics with Ease**
Enjoy comics in a clean, distraction-free reader. Our intuitive interface supports smooth scrolling, fast loading, and flexible viewing modes — vertical or horizontal, your choice.

🚀 **Fast, Lightweight, and Mobile-Friendly**
Built with performance in mind, the site loads quickly even on slower connections and adapts beautifully to any screen — mobile, tablet, or desktop.

🔍 **Smart Navigation**
Search by genre, artist, or title. Our smart filters help you discover hidden gems and keep track of your favorites.

💡 **For Readers and Creators**
We believe in supporting creators. Artists can showcase their work, gain exposure, and reach new fans.

---

A lightweight web application built with Go and Echo framework for serving comics and images in an organized, browsable format.

## Features

- **Comics Browser**: Browse and read comics organized by title and chapter
- **Images Gallery**: View image collections organized by folders
- **Responsive UI**: Modern web interface with Tailwind CSS
- **File-based Organization**: Automatically organizes content from file system structure
- **Static File Serving**: Efficient serving of assets, comics, and images

## Project Structure

```
H2O/
├── main.go              # Main application entry point
├── go.mod               # Go module dependencies
├── go.sum               # Go module checksums
├── views/               # HTML templates and static assets
│   ├── assets/          # CSS, JS, and image assets
│   ├── comics/          # Comics organized by title/chapter
│   ├── images/          # Image collections
│   └── *.html           # HTML templates
└── README.md            # This file
```

## Prerequisites

- **Go 1.16 or higher** - [Download Go](https://golang.org/dl/)
- **Git** (for cloning the repository)

## Installation

1. **Clone the repository** (if not already done):
   ```bash
   git clone <repository-url>
   cd H2O
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

## Running the Application

### Method 1: Direct Go Run
```bash
go run main.go
```

### Method 2: Build and Run
```bash
# Build the application
go build -o H2O main.go

# Run the built binary
./H2O
```

### Method 3: Development with Hot Reload (Optional)
If you have `go-watcher` installed:
```bash
go install github.com/canthefason/go-watcher/cmd/watcher@latest
watcher -cmd="go run main.go" -recursive
```

## Accessing the Application

Once the application is running, open your web browser and navigate to:

```
http://localhost:100
```

The application will be available on port 100.

## Content Organization

### Comics Structure
Place your comics in the `views/comics/` directory with the following structure:
```
views/comics/
├── Title1/
│   ├── icon.png          # Cover image for the title
│   ├── Chapter1/
│   │   ├── page1.jpg
│   │   ├── page2.jpg
│   │   └── ...
│   └── Chapter2/
│       ├── page1.jpg
│       └── ...
└── Title2/
    └── ...
```

### Images Structure
Place your image collections in the `views/images/` directory:
```
views/images/
├── Collection1/
│   ├── image1.jpg
│   ├── image2.jpg
│   └── ...
└── Collection2/
    ├── photo1.png
    └── ...
```

## API Endpoints

- `GET /` - Home page showing all comics and image collections
- `GET /:title` - Title page showing chapters or images for a specific title
- `GET /:title/:chapter` - Chapter page showing comic pages
- `GET /assets/*` - Static assets (CSS, JS, images)
- `GET /comics/*` - Comic files
- `GET /images/*` - Image files

## Dependencies

- **Echo v4** - Web framework for Go
- **Go 1.16+** - Programming language and runtime

## Development

### Adding New Features
1. Modify `main.go` to add new routes or handlers
2. Update HTML templates in `views/` directory
3. Add any new static assets to `views/assets/`

### Building for Production
```bash
# Build for current platform
go build -o H2O main.go

# Build for specific platform (e.g., Linux)
GOOS=linux GOARCH=amd64 go build -o H2O main.go

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o H2O.exe main.go
```

## Configuration

The application runs on port 100 by default. To change the port, modify line 25 in `main.go`:

```go
if err := e.Start(":8080"); err != http.ErrServerClosed {
    log.Fatal(err)
}
```

## Troubleshooting

### Common Issues

1. **Port already in use**: Change the port number in `main.go`
2. **Missing dependencies**: Run `go mod download`
3. **Permission denied**: Ensure you have write permissions in the project directory
4. **Content not showing**: Check that your comics/images are in the correct directory structure

### Logs
The application logs to stdout. Check the terminal output for any error messages.

## License

This project is licensed under the GNU General Public License v3.0 (GPL-3.0).

### GPL-3.0 License Details

- **Freedom to use**: You can use this software for any purpose
- **Freedom to study**: You can examine the source code and modify it
- **Freedom to share**: You can distribute copies of the software
- **Freedom to improve**: You can distribute modified versions

### Key Requirements

- Any derivative work must also be licensed under GPL-3.0
- Source code must be made available when distributing the software
- License and copyright notices must be preserved

For the full license text, see [LICENSE](LICENSE) file or visit [GNU GPL v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html).

## Contributing

Contributions are welcome! Here's how you can contribute:

1. **Fork the repository** on GitHub
2. **Clone your fork** to your local machine
3. **Make your changes** - add features, fix bugs, or improve documentation
4. **Test your changes** to ensure everything works correctly
5. **Commit and push** your changes to your fork
6. **Submit a pull request** to the main repository