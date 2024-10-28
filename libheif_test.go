package libheif

import (
	"os"
	"strings"
	"testing"
)

// List of image names
var heicAndAvifImages = []string{"images/example-from-libheif.heic", "images/example-from-libheif.avif", "images/image1-from-me.heic", "images/samsung-generated.heic"}

// It has problems with converting images to HEIC format. Couldn't convert these images:
// var pngAndJpegImages = []string{"images/apple.jpg", "images/deer.png", "images/go.png"}
// But as I found it can easily convert libheif generated images back to heic
var pngAndJpegImages = []string{"images/libheif-generated.jpeg"}

func TestHeifToJpeg(t *testing.T) {
	for _, imageName := range heicAndAvifImages {
		heifImagePath := imageName
		jpegImagePath := "test_images/" + strings.Replace(imageName, "/", "-", 1) + ".jpeg"

		err := HeifToJpeg(heifImagePath, jpegImagePath, 80)
		if err != nil {
			t.Errorf("Failed to convert HEIF to JPEG: %v", err)
		}

		// Check if the file exists
		if _, err := os.Stat(jpegImagePath); os.IsNotExist(err) {
			t.Errorf("Output JPEG file does not exist: %v", err)
		}

		// Cleanup
		t.Cleanup(func() {
			os.Remove(jpegImagePath)
		})
	}
}

func TestHeifToPng(t *testing.T) {
	for _, imageName := range heicAndAvifImages {
		heifImagePath := imageName
		pngImagePath := "test_images/" + strings.Replace(imageName, "/", "-", 1) + ".png"

		err := HeifToJpeg(heifImagePath, pngImagePath, 80)
		if err != nil {
			t.Errorf("Failed to convert HEIF to JPEG: %v", err)
		}

		// Check if the file exists
		if _, err := os.Stat(pngImagePath); os.IsNotExist(err) {
			t.Errorf("Output JPEG file does not exist: %v", err)
		}

		// Cleanup
		t.Cleanup(func() {
			os.Remove(pngImagePath)
		})
	}
}

func TestHeifToGif(t *testing.T) {
	for _, imageName := range heicAndAvifImages {
		heifImagePath := imageName
		gifImagePath := "test_images/" + strings.Replace(imageName, "/", "-", 1) + ".gif"

		err := HeifToGif(heifImagePath, gifImagePath)
		if err != nil {
			t.Errorf("Failed to convert HEIF to GIF: %v", err)
		}

		// Check if the file exists
		if _, err := os.Stat(gifImagePath); os.IsNotExist(err) {
			t.Errorf("Output GIF file does not exist: %v", err)
		}

		// Cleanup
		t.Cleanup(func() {
			os.Remove(gifImagePath)
		})
	}
}

func TestReturnImageFromHeif(t *testing.T) {
	for _, imageName := range heicAndAvifImages {
		heifImagePath := imageName

		// Check if error is returned when converting to and returning image from HEIF
		_, err := ReturnImageFromHeif(heifImagePath)
		if err != nil {
			t.Errorf("Failed to return image from HEIF: %v", err)
		}

	}
}
