package api

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type FileInfo struct {
	URL          string `json:"url"`
	Filename     string `json:"filename"`
	FullPath     string `json:"fullPath"`
	Size         int    `json:"size"`
	LastModified int    `json:"lastModified"`
}

type ApiResponse struct {
	Files []FileInfo `json:"files"`
}

func QueryFiles(sessionCookie string, keywords []string, extensions map[string][]string) ([]FileInfo, error) {
	apiURL := "https://buckets.grayhatwarfare.com/api/v2/files"
	var allFiles []FileInfo

	// Convert map keys (extensions) to a single string
	extensionKeys := getMapKeys(extensions)
	extensionsParam := strings.Join(extensionKeys, ",")

	// Pagination variables (local to the function)
	start := 0
	limit := 1000

	// transport for the api
	transport := &http.Transport{
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
		DisableKeepAlives:   false,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: false},
	}

	client := &http.Client{Transport: transport}

	go func() {
		for {
			time.Sleep(60 * time.Second) // Run every 60 seconds
			transport.CloseIdleConnections()
			fmt.Println("Idle connections closed")
		}
	}()

	for {
		// Build query parameters
		params := url.Values{}
		params.Set("keywords", joinKeywords(keywords))
		params.Set("extensions", extensionsParam)
		params.Set("limit", fmt.Sprintf("%d", limit))
		params.Set("start", fmt.Sprintf("%d", start))

		// Build the full URL
		fullURL := fmt.Sprintf("%s?%s", apiURL, params.Encode())
		// Add headers (including the authorization token)

		// Create the request to the api
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}
		// Add headers (including the authorization token)
		req.Header.Set("Authorization", " Bearer "+sessionCookie)

		// Resend the request
		resp, err := doRequestWithRetry(client, req, 3) // 3 retries
		if err != nil {
			return nil, fmt.Errorf("failed to send request after retries: %w", err)
		}

		defer resp.Body.Close()

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}

		// Check for successful status code
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("API error: %s", resp.Status)
		}

		// Parse the JSON response
		var apiResponse ApiResponse
		err = json.Unmarshal(body, &apiResponse)
		if err != nil {
			return nil, fmt.Errorf("failed to parse JSON: %w", err)
		}

		// Append results to the aggregate list
		allFiles = append(allFiles, apiResponse.Files...)

		// Stop if fewer than the requested limit of files is returned
		if len(apiResponse.Files) < limit {
			break
		}

		// Move to the next page
		start += limit
	}
	return allFiles, nil
}

// Helper function to get keys of a map
func getMapKeys(m map[string][]string) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// Helper function to join keywords into a single string
func joinKeywords(keywords []string) string {
	return url.QueryEscape(strings.Join(keywords, " "))
}

func doRequestWithRetry(client *http.Client, req *http.Request, retries int) (*http.Response, error) {
	for i := 0; i < retries; i++ {
		resp, err := client.Do(req)
		if err == nil {
			return resp, nil
		}
		fmt.Printf("Retry %d/%d failed: %v\n", i+1, retries, err)
		time.Sleep(2 * time.Second) // Wait before retrying
	}
	return nil, fmt.Errorf("all retries failed")
}

// func saveResultsToFile(results []FileInfo, outputFile string) error {
// 	// Open the file for appending (or create it if it doesn't exist)
// 	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		return fmt.Errorf("failed to open file: %w", err)
// 	}
// 	defer file.Close()

// 	// Write each result as a JSON object, one per line
// 	for _, result := range results {
// 		jsonBytes, err := json.Marshal(result)
// 		if err != nil {
// 			return fmt.Errorf("failed to marshal result: %w", err)
// 		}

// 		// Write JSON line to the file
// 		_, err = file.Write(append(jsonBytes, '\n'))
// 		if err != nil {
// 			return fmt.Errorf("failed to write to file: %w", err)
// 		}
// 	}

// 	return nil
// }
