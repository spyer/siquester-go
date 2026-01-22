package siq

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

const (
	contentFileName = "content.xml"
)

// Document represents a SIQ document with package content and media collections.
type Document struct {
	Package *Package
	// Media collections
	Images map[string][]byte
	Audio  map[string][]byte
	Video  map[string][]byte
	HTML   map[string][]byte
	// Source file path (if loaded from file)
	FilePath string
}

// NewDocument creates a new empty document.
func NewDocument(name, author string) *Document {
	pkg := NewPackage(name, author)
	pkg.ID = uuid.New().String()
	return &Document{
		Package: pkg,
		Images:  make(map[string][]byte),
		Audio:   make(map[string][]byte),
		Video:   make(map[string][]byte),
		HTML:    make(map[string][]byte),
	}
}

// LoadFromFile loads a document from a .siq/.zip file or a folder.
func LoadFromFile(path string) (*Document, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("failed to stat path: %w", err)
	}
	// If it's a directory, load from folder
	if stat.IsDir() {
		return LoadFromFolder(path)
	}
	// Otherwise load as archive (.siq or .zip)
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()
	doc, err := Load(f, stat.Size())
	if err != nil {
		return nil, err
	}
	doc.FilePath = path
	return doc, nil
}

// LoadFromFolder loads a document from an extracted folder.
func LoadFromFolder(folderPath string) (*Document, error) {
	doc := &Document{
		Images:   make(map[string][]byte),
		Audio:    make(map[string][]byte),
		Video:    make(map[string][]byte),
		HTML:     make(map[string][]byte),
		FilePath: folderPath,
	}
	// Read content.xml
	contentPath := filepath.Join(folderPath, contentFileName)
	contentData, err := os.ReadFile(contentPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read content.xml: %w", err)
	}
	var pkg Package
	if err := xml.Unmarshal(contentData, &pkg); err != nil {
		return nil, fmt.Errorf("failed to parse content.xml: %w", err)
	}
	doc.Package = &pkg
	// Read media folders
	mediaFolders := map[string]*map[string][]byte{
		"Images": &doc.Images,
		"Audio":  &doc.Audio,
		"Video":  &doc.Video,
		"Html":   &doc.HTML,
	}
	for folderName, collection := range mediaFolders {
		mediaPath := filepath.Join(folderPath, folderName)
		if _, err := os.Stat(mediaPath); os.IsNotExist(err) {
			continue
		}
		entries, err := os.ReadDir(mediaPath)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			filePath := filepath.Join(mediaPath, entry.Name())
			data, err := os.ReadFile(filePath)
			if err != nil {
				continue
			}
			(*collection)[entry.Name()] = data
		}
	}
	return doc, nil
}

// Load loads a document from a reader.
func Load(r io.ReaderAt, size int64) (*Document, error) {
	zipReader, err := zip.NewReader(r, size)
	if err != nil {
		return nil, fmt.Errorf("failed to read zip: %w", err)
	}
	doc := &Document{
		Images: make(map[string][]byte),
		Audio:  make(map[string][]byte),
		Video:  make(map[string][]byte),
		HTML:   make(map[string][]byte),
	}
	for _, file := range zipReader.File {
		if file.Name == contentFileName {
			pkg, err := readPackageFromZip(file)
			if err != nil {
				return nil, err
			}
			doc.Package = pkg
		} else {
			// Handle media files
			collection, name := parseMediaPath(file.Name)
			if collection != "" && name != "" {
				data, err := readFileFromZip(file)
				if err != nil {
					return nil, err
				}
				switch collection {
				case CollectionImages:
					doc.Images[name] = data
				case CollectionAudio:
					doc.Audio[name] = data
				case CollectionVideo:
					doc.Video[name] = data
				case CollectionHTML:
					doc.HTML[name] = data
				}
			}
		}
	}
	if doc.Package == nil {
		return nil, errors.New("content.xml not found in package")
	}
	return doc, nil
}

func readPackageFromZip(file *zip.File) (*Package, error) {
	rc, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open content.xml: %w", err)
	}
	defer rc.Close()
	data, err := io.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("failed to read content.xml: %w", err)
	}
	var pkg Package
	if err := xml.Unmarshal(data, &pkg); err != nil {
		return nil, fmt.Errorf("failed to parse content.xml: %w", err)
	}
	return &pkg, nil
}

func readFileFromZip(file *zip.File) ([]byte, error) {
	rc, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %w", file.Name, err)
	}
	defer rc.Close()
	return io.ReadAll(rc)
}

func parseMediaPath(path string) (collection, name string) {
	parts := strings.SplitN(path, "/", 2)
	if len(parts) != 2 {
		return "", ""
	}
	collection = parts[0]
	name = parts[1]
	// URL-decode the file name (ZIP files may contain URL-encoded names)
	if decoded, err := url.PathUnescape(name); err == nil {
		name = decoded
	}
	// Normalize collection name
	switch strings.ToLower(collection) {
	case "images":
		collection = CollectionImages
	case "audio":
		collection = CollectionAudio
	case "video":
		collection = CollectionVideo
	case "html":
		collection = CollectionHTML
	default:
		return "", ""
	}
	return collection, name
}

// SaveToFile saves the document to a .siq file.
func (d *Document) SaveToFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()
	if err := d.Save(f); err != nil {
		return err
	}
	d.FilePath = path
	return nil
}

// Save saves the document to a writer.
func (d *Document) Save(w io.Writer) error {
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()
	// Write content.xml
	contentWriter, err := zipWriter.Create(contentFileName)
	if err != nil {
		return fmt.Errorf("failed to create content.xml: %w", err)
	}
	if _, err := contentWriter.Write([]byte(xml.Header)); err != nil {
		return fmt.Errorf("failed to write xml header: %w", err)
	}
	encoder := xml.NewEncoder(contentWriter)
	encoder.Indent("", "  ")
	if err := encoder.Encode(d.Package); err != nil {
		return fmt.Errorf("failed to encode package: %w", err)
	}
	// Write quality.marker (empty file required by canonical SIGame client)
	if _, err := zipWriter.Create("quality.marker"); err != nil {
		return fmt.Errorf("failed to create quality.marker: %w", err)
	}
	// Write media files
	for name, data := range d.Images {
		if err := writeMediaToZip(zipWriter, CollectionImages, name, data); err != nil {
			return err
		}
	}
	for name, data := range d.Audio {
		if err := writeMediaToZip(zipWriter, CollectionAudio, name, data); err != nil {
			return err
		}
	}
	for name, data := range d.Video {
		if err := writeMediaToZip(zipWriter, CollectionVideo, name, data); err != nil {
			return err
		}
	}
	for name, data := range d.HTML {
		if err := writeMediaToZip(zipWriter, CollectionHTML, name, data); err != nil {
			return err
		}
	}
	return nil
}

func writeMediaToZip(zw *zip.Writer, collection, name string, data []byte) error {
	path := filepath.Join(collection, name)
	w, err := zw.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create %s: %w", path, err)
	}
	if _, err := w.Write(data); err != nil {
		return fmt.Errorf("failed to write %s: %w", path, err)
	}
	return nil
}

// ToBytes serializes the document to bytes.
func (d *Document) ToBytes() ([]byte, error) {
	var buf bytes.Buffer
	if err := d.Save(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// LoadFromBytes loads a document from bytes.
func LoadFromBytes(data []byte) (*Document, error) {
	r := bytes.NewReader(data)
	return Load(r, int64(len(data)))
}

// GetMedia returns media data by type and name.
func (d *Document) GetMedia(mediaType, name string) ([]byte, bool) {
	var collection map[string][]byte
	switch mediaType {
	case ContentTypeImage, CollectionImages:
		collection = d.Images
	case ContentTypeAudio, CollectionAudio:
		collection = d.Audio
	case ContentTypeVideo, CollectionVideo:
		collection = d.Video
	case ContentTypeHTML, CollectionHTML:
		collection = d.HTML
	default:
		return nil, false
	}
	data, ok := collection[name]
	return data, ok
}

// SetMedia sets media data by type and name.
func (d *Document) SetMedia(mediaType, name string, data []byte) error {
	switch mediaType {
	case ContentTypeImage, CollectionImages:
		d.Images[name] = data
	case ContentTypeAudio, CollectionAudio:
		d.Audio[name] = data
	case ContentTypeVideo, CollectionVideo:
		d.Video[name] = data
	case ContentTypeHTML, CollectionHTML:
		d.HTML[name] = data
	default:
		return fmt.Errorf("unknown media type: %s", mediaType)
	}
	return nil
}

// DeleteMedia deletes media by type and name.
func (d *Document) DeleteMedia(mediaType, name string) bool {
	var collection map[string][]byte
	switch mediaType {
	case ContentTypeImage, CollectionImages:
		collection = d.Images
	case ContentTypeAudio, CollectionAudio:
		collection = d.Audio
	case ContentTypeVideo, CollectionVideo:
		collection = d.Video
	case ContentTypeHTML, CollectionHTML:
		collection = d.HTML
	default:
		return false
	}
	if _, ok := collection[name]; ok {
		delete(collection, name)
		return true
	}
	return false
}

// ListMedia returns a list of media files by type.
func (d *Document) ListMedia(mediaType string) []string {
	var collection map[string][]byte
	switch mediaType {
	case ContentTypeImage, CollectionImages:
		collection = d.Images
	case ContentTypeAudio, CollectionAudio:
		collection = d.Audio
	case ContentTypeVideo, CollectionVideo:
		collection = d.Video
	case ContentTypeHTML, CollectionHTML:
		collection = d.HTML
	default:
		return nil
	}
	names := make([]string, 0, len(collection))
	for name := range collection {
		names = append(names, name)
	}
	return names
}

// GetCollectionForType returns the collection name for a content type.
func GetCollectionForType(contentType string) string {
	switch contentType {
	case ContentTypeImage:
		return CollectionImages
	case ContentTypeAudio:
		return CollectionAudio
	case ContentTypeVideo:
		return CollectionVideo
	case ContentTypeHTML:
		return CollectionHTML
	default:
		return ""
	}
}
