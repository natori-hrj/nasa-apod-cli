package translate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const googleTranslateURL = "https://translate.googleapis.com/translate_a/single"

var httpClient = &http.Client{Timeout: 30 * time.Second}

// ToJapanese translates English text to Japanese using Google Translate.
func ToJapanese(text string) (string, error) {
	params := url.Values{
		"client": {"gtx"},
		"sl":     {"en"},
		"tl":     {"ja"},
		"dt":     {"t"},
		"q":      {text},
	}

	reqURL := googleTranslateURL + "?" + params.Encode()
	resp, err := httpClient.Get(reqURL)
	if err != nil {
		return "", fmt.Errorf("translation request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read translation response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("translation API error (status %d)", resp.StatusCode)
	}

	// Response format: [[["translated text","original text",...],...],...]
	var result []any
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to parse translation response: %w", err)
	}

	translated := ""
	sentences, ok := result[0].([]any)
	if !ok {
		return "", fmt.Errorf("unexpected translation response format")
	}
	for _, s := range sentences {
		parts, ok := s.([]any)
		if !ok || len(parts) == 0 {
			continue
		}
		if text, ok := parts[0].(string); ok {
			translated += text
		}
	}

	if translated == "" {
		return "", fmt.Errorf("empty translation result")
	}
	return translated, nil
}
