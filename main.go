package main

// IMPORTS
import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"pdf_greyhat_go/api"
	"pdf_greyhat_go/download"
)

func main() {
	// Initialize session and keywords
	sessionCookie := "01931a3ff4929fa0e8d8c93ba9dac24c"
	keywords := []string{"internal-use", "forofficialuseonly", "docs", "documentfile", "file", "note", "writeup", "write-up", "filerecord", "draft", "textfile", "paper", "form", "worksheet", "sensitive", "secret", "nonpublic", "classified", "proprietary", "privileged", "summary", "log", "statement", "record", "analysis", "assessment", "review", "findings", "privatefile", "classifieddoc", "restricteddocument", "internalnote", "internalrecord", "techdocs", "technicalmanual", "specdoc", "engineeringdoc", "systemguide", "implementationguide", "private", "restricted", "classified", "internal-use", "for_official_use_only", "docs", "document_file", "file", "note", "writeup", "write-up", "file_record", "draft", "text_file", "paper", "form", "worksheet", "sensitive", "secret", "non_public", "classified", "proprietary", "privileged", "summary", "log", "statement", "record", "analysis", "assessment", "review", "findings", "private_file", "classified_doc", "restricted_document", "internal_note", "internal_record", "tech_docs", "technical_manual", "spec_doc", "engineering_doc", "system_guide", "implementation_guide"}
	extensions := map[string][]string{
		"json": {"Algemeen Dagblad",
			"Allegro",
			"Axel Springer",
			"Azena",
			"BMW Group",
			"BMW Group Automotive",
			"Bpost",
			"Bühler",
			"CM.com",
			"Canada Post",
			"Capital.com",
			"Cloudways by DigitalOcean",
			"Cross Border Fines",
			"Cyber Security Coalition",
			"DPG Media",
			"De Lijn",
			"De Morgen",
			"De Volkskrant",
			"Delen Private Bank",
			"Digitaal Vlaanderen",
			"DigitalOcean",
			"Donorbox",
			"E-Gor",
			"EURid",
			"Fing",
			"HRS Group",
			"Henkel",
			"Here Technologies",
			"Het Laatste Nieuws",
			"Het Parool",
			"Humo",
			"Kinepolis Group",
			"Lansweeper",
			"Libelle",
			"Mobile Vikings",
			"Moralis",
			"Nestlé",
			"Nexuzhealth",
			"Nexuzhealth Web PACS",
			"OVO",
			"PDQ bug bounty program",
			"PeopleCert",
			"Personio",
			"Port of Antwerp-Bruges",
			"Purolator",
			"RGF BE",
			"RIPE NCC",
			"Randstad",
			"Red Bull",
			"Revolut",
			"SimScale",
			"Sixt",
			"Social Deal",
			"Soundtrack Your Brand",
			"Sqills",
			"Stravito",
			"Suivo bug bounty",
			"Sustainable",
			"Telenet",
			"Tempo-Team",
			"Tomorrowland",
			"Torfs",
			"Trouw",
			"TrueLayer",
			"Twago",
			"Tweakers",
			"UZ Leuven",
			"Ubisoft",
			"VRT",
			"VTM GO",
			"Venly",
			"Vlerick Business School",
			"Voi Scooters",
			"WP Engine",
			"Yacht",
			"Yahoo",
			"e-tracker",
			"token",
			"eHealth Hub VZN KUL"},
		"pdf": {"Algemeen Dagblad",
			"Allegro",
			"Axel Springer",
			"Azena",
			"BMW Group",
			"BMW Group Automotive",
			"Bpost",
			"Bühler",
			"CM.com",
			"Canada Post",
			"Capital.com",
			"Cloudways by DigitalOcean",
			"Cross Border Fines",
			"Cyber Security Coalition",
			"DPG Media",
			"De Lijn",
			"De Morgen",
			"De Volkskrant",
			"Delen Private Bank",
			"Digitaal Vlaanderen",
			"DigitalOcean",
			"Donorbox",
			"E-Gor",
			"EURid",
			"Fing",
			"HRS Group",
			"Henkel",
			"Here Technologies",
			"Het Laatste Nieuws",
			"Het Parool",
			"Humo"},
		"js": {
			"token",
			"secret",
			"key",
			"password",
			"pass",
			"apikey",
			"api_key",
			"apiToken",
			"api_token",
			"apiKey",
			"api_key",
			"apiSecret",
			"api_secret",
			"accessToken",
			"access_token",
			"accessKey",
			"access_key",
			"privateKey",
			"private_key",
			"publicKey",
			"public_key",
			"clientSecret",
			"client_secret",
			"clientId",
			"client_id",
			"refreshToken",
			"refresh_token",
			"sessionToken",
			"session_token",
			"authToken",
			"auth_token",
			"authorization",
			"authorizationToken",
			"authorization_token",
			"bearerToken",
			"bearer_token",
			"jwt",
			"jwtToken",
			"jwt_token",
			"oauth",
			"oauth_token",
			"oauth2",
			"oauth2Token",
			"oauth2_token",
			"consumerKey",
			"consumer_key",
			"consumerSecret",
			"consumer_secret",
			"credentials",
			"credential",
			"secure",
			"secureKey",
			"secure_key",
			"securityKey",
			"security_key",
			"aws_access_key",
			"aws_secret_access_key",
			"s3_bucket",
			"s3_access_key",
			"s3_secret_key",
			"firebase_api_key",
			"firebase_secret",
			"gcp_key",
			"gcp_secret",
			"google_api_key",
			"google_client_secret",
			"azure_key",
			"azure_secret",
			"slack_token",
			"github_token",
			"github_secret",
			"gitlab_token",
			"gitlab_secret",
			"stripe_key",
			"stripe_secret",
			"paypal_key",
			"paypal_secret",
			"twilio_key",
			"twilio_secret",
			"sendgrid_key",
			"sendgrid_secret",
			"smtp_password",
			"smtp_secret",
			"smtp_key",
			"db_password",
			"db_secret",
			"database_password",
			"database_secret",
			"masterKey",
			"master_key",
			"encryptionKey",
			"encryption_key",
			"awsKey",
			"aws_key",
			"awsSecret",
			"aws_secret"},
	}
	for _, keyword := range keywords {
		outputFile := fmt.Sprintf("results-%s.json", keyword)
		fmt.Printf("Searching for files with keyword: %s\n", keyword)
		var files []api.FileInfo
		var err error
		maxRetries := 3
		for retries := 0; retries < maxRetries; retries++ {
			files, err = api.QueryFiles(sessionCookie, []string{keyword}, extensions)
			if err == nil {
				break // Exit the retry loop if successful
			}
			log.Printf("Retry %d/%d for keyword '%s' failed: %v", retries+1, maxRetries, keyword, err)
			time.Sleep(2 * time.Second)
		}
		if err != nil {
			log.Printf("All retries failed for keyword '%s'\n", keyword)
			continue
		}

		// Create a semaphore for concurrent downloads
		var wg sync.WaitGroup
		// Initialize results
		results := make([]map[string]interface{}, 0)
		// RESULTS:
		/*
				{"Filename": "file1.pdf", "URL": "http://example.com/file1", "Matches": 10},
			    {"Filename": "file2.pdf", "URL": "http://example.com/file2", "Matches": 5},
		*/
		mutex := &sync.Mutex{}

		// Set the concurrency limit
		concurrencyLimit := 6
		// use semaphore var to set a maximum number of concurrent goroutines
		semaphore := make(chan struct{}, concurrencyLimit)

		// Creates a timer that triggers every 60 seconds.
		ticker := time.NewTicker(60 * time.Second)
		defer ticker.Stop()

		// ensure to goes periodicly saving it's making to avoid big lost if the process is interrupted
		go func() {
			// A channel that emits a signal every time the ticker fires.
			for range ticker.C {
				// save the file periodically
				mutex.Lock()
				err := saveResults(results, outputFile)
				if err != nil {
					log.Printf("Error saving periodic results for keyword '%s': %v", keyword, err)
				}
				// fmt.Printf("Result added: %+v\n", results) // Add this line for debugging
				// MUTEX write the file but priventing race conditions
				mutex.Unlock()
			}
		}()

		for _, fileInfo := range files {

			fmt.Println("Processing file:", fileInfo.Filename)
			if fileInfo.Size > 50*1024*1024 { // Skip files larger than 50 MB
				fmt.Printf("Skipping large file: %s\n", fileInfo.Filename)
				continue
			}

			// increment the wait counter
			wg.Add(1)
			go func(file api.FileInfo) {
				// DECREMENT the wait routine when it's done
				defer wg.Done()
				// send an empty struct into the sempahore channel
				semaphore <- struct{}{} // Acquire a semaphore slot
				// semaphoro green!
				defer func() { <-semaphore }() // Release slot after processing

				// fmt.Printf("Found file: %s (URL: %s, Size: %d bytes)\n", file.Filename, file.URL, file.Size)

				result := download.ProcessFile(file, extensions) // redefine result
				// redefine the results with the function proces file
				if result != nil {
					// append the result (no overwrite)
					mutex.Lock()
					results = append(results, result)

					// MUTEX write the file but priventing race conditions
					mutex.Unlock()
				}
			}(fileInfo)
		}
		// The file info is a struct
		/* type FileInfo struct {
			URL      string
			Filename string
			Size     int
		} */
		// The results are saved as JSON in results.json, after the whole fucking process ends:
		wg.Wait() // Wait for all goroutines to complete
		mutex.Lock()
		err = saveResults(results, outputFile)
		if err != nil {
			log.Printf("Error saving final results for keyword '%s': %v", keyword, err)
		}
		mutex.Unlock()
	}
}

func saveResults(results []map[string]interface{}, outputFile string) error {
	fmt.Printf("Saving %d results...\n", len(results)) // Debug log

	file, err := os.Create(outputFile) // Create (or overwrite) results.json
	if err != nil {
		return fmt.Errorf("failed to create output file '%s': %w", outputFile, err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Add indentation for readability

	if err := encoder.Encode(results); err != nil {
		return fmt.Errorf("failed to write JSON to file '%s': %w", outputFile, err)
	}
	fmt.Printf("Results saved to %s\n", outputFile)
	return nil
}
