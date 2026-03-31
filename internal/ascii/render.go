package ascii

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"time"

	"github.com/qeesung/image2ascii/convert"
)

var httpClient = &http.Client{Timeout: 60 * time.Second}

// RenderFromURL downloads an image from the given URL and returns its ASCII art representation.
func RenderFromURL(imageURL string) (string, error) {
	resp, err := httpClient.Get(imageURL)
	if err != nil {
		return "", fmt.Errorf("failed to download image: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("image download failed (status %d)", resp.StatusCode)
	}

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to decode image: %w", err)
	}

	converter := convert.NewImageConverter()
	opts := convert.DefaultOptions
	opts.FixedWidth = 80
	opts.FixedHeight = 40
	opts.Colored = true

	return converter.Image2ASCIIString(img, &opts), nil
}
