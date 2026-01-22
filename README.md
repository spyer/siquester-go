# SIQuester Go

A cross-platform web-based editor for SIGame question packages (.siq files).

## Features

- Create, edit, and save SIGame packages
- Tree-based navigation through rounds, themes, and questions
- Media management (images, audio, video)
- Import from XML and YAML formats
- Export to XML and YAML formats
- Modern web UI with dark/light theme support

## Requirements

- Go 1.21 or later

## Building

```bash
cd siquester-go
go build -o siquester ./cmd/siquester
```

## Running

```bash
./siquester -port 8080
```

Then open http://localhost:8080 in your browser.

## Package Format

SIQ files are ZIP archives containing:
- `content.xml` - Package structure and metadata
- `Images/` - Image files
- `Audio/` - Audio files
- `Video/` - Video files
- `Html/` - HTML content

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | /api/packages | Create new package |
| GET | /api/packages/:id | Get package content |
| PUT | /api/packages/:id | Update package |
| POST | /api/packages/:id/save | Save to disk |
| POST | /api/packages/open | Open .siq file |
| POST | /api/import/xml | Import from XML |
| POST | /api/import/yaml | Import from YAML |
| GET | /api/export/:id/xml | Export to XML |
| GET | /api/export/:id/yaml | Export to YAML |
| POST | /api/media/:id | Upload media |
| GET | /api/media/:id/:type/:name | Get media file |
| DELETE | /api/media/:id/:type/:name | Delete media |

## License

See LICENSE file in the parent directory.
