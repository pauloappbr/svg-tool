package converter

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"

	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

// OutputSpec defines the configuration for a single output file.
type OutputSpec struct {
	Name string
	Size int
}

// DefaultWebAssets returns the standard list of web assets (Favicons, PWA, Apple Touch Icon).
func DefaultWebAssets() []OutputSpec {
	return []OutputSpec{
		{"favicon-16.png", 16},
		{"favicon-32.png", 32},
		{"favicon-512.png", 512},
		{"apple-touch-icon.png", 180},
		{"pwa-android.png", 192},
	}
}

// ProcessSVG reads the input SVG file and generates the specified PNG and ICO files in the output directory.
func ProcessSVG(inputPath, outputDir string, specs []OutputSpec, generateIco bool) error {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", outputDir, err)
	}

	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file %s: %w", inputPath, err)
	}
	defer file.Close()

	icon, err := oksvg.ReadIconStream(file)
	if err != nil {
		return fmt.Errorf("failed to parse SVG data: %w", err)
	}

	origW, origH := icon.ViewBox.W, icon.ViewBox.H
	aspectRatio := origW / origH

	var icoImages []image.Image

	for _, spec := range specs {
		w := spec.Size
		h := int(float64(w) / aspectRatio)

		img := renderSVG(icon, w, h)
		fullPath := filepath.Join(outputDir, spec.Name)

		if err := savePNG(fullPath, img); err != nil {
			return fmt.Errorf("failed to save image %s: %w", spec.Name, err)
		}

		if generateIco && (w == 16 || w == 32 || w == 48) {
			icoImages = append(icoImages, img)
		}
	}

	if generateIco {
		if len(icoImages) == 0 {
			for _, s := range []int{16, 32, 48} {
				icoImages = append(icoImages, renderSVG(icon, s, s))
			}
		}
		if err := saveICONative(filepath.Join(outputDir, "favicon.ico"), icoImages); err != nil {
			return fmt.Errorf("failed to save favicon.ico: %w", err)
		}
	}

	return nil
}

// renderSVG converts the vector icon to a rasterized image.
func renderSVG(icon *oksvg.SvgIcon, w, h int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	scanner := rasterx.NewScannerGV(w, h, img, img.Bounds())
	raster := rasterx.NewDasher(w, h, scanner)
	icon.SetTarget(0, 0, float64(w), float64(h))
	icon.Draw(raster, 1.0)
	return img
}

// savePNG saves the image to disk in PNG format.
func savePNG(filename string, img image.Image) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	b := bufio.NewWriter(f)
	if err := png.Encode(b, img); err != nil {
		return err
	}
	return b.Flush()
}

type icoHeader struct {
	Reserved, Type, Count uint16
}

type icoDirEntry struct {
	Width, Height, Palette, Reserved uint8
	Planes, BitCount                 uint16
	Size, Offset                     uint32
}

// saveICONative implements the ICO format manually to avoid unstable external dependencies.
func saveICONative(filename string, images []image.Image) error {
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.LittleEndian, icoHeader{0, 1, uint16(len(images))}); err != nil {
		return err
	}

	offset := uint32(6 + (16 * len(images)))
	var imageBytes [][]byte
	var entries []icoDirEntry

	for _, img := range images {
		pngBuf := new(bytes.Buffer)
		if err := png.Encode(pngBuf, img); err != nil {
			return err
		}
		data := pngBuf.Bytes()
		imageBytes = append(imageBytes, data)

		bounds := img.Bounds()
		entry := icoDirEntry{
			Width:    uint8(bounds.Dx()),
			Height:   uint8(bounds.Dy()),
			Planes:   1,
			BitCount: 32,
			Size:     uint32(len(data)),
			Offset:   offset,
		}
		if bounds.Dx() >= 256 {
			entry.Width = 0
		}
		if bounds.Dy() >= 256 {
			entry.Height = 0
		}

		entries = append(entries, entry)
		offset += uint32(len(data))
	}

	for _, e := range entries {
		if err := binary.Write(buf, binary.LittleEndian, e); err != nil {
			return err
		}
	}
	for _, data := range imageBytes {
		if _, err := buf.Write(data); err != nil {
			return err
		}
	}

	return os.WriteFile(filename, buf.Bytes(), 0644)
}
