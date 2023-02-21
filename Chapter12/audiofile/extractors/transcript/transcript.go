package transcript

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/marianina8/audiofile/models"
)

func Extract(m *models.Audio) error {
	apiKey := os.Getenv("ASSEMBLYAI_API_KEY")
	if apiKey == "" {
		fmt.Println("missing ASSEMBLYAI_API_KEY.  Skipping transcript exraction")
		return nil
	}
	const UPLOAD_URL = "https://api.assemblyai.com/v2/upload"

	// Load file
	data, err := os.ReadFile(m.Path)
	if err != nil {
		return err
	}

	// Setup HTTP client and set header
	client := &http.Client{}
	req, _ := http.NewRequest("POST", UPLOAD_URL, bytes.NewBuffer(data))
	req.Header.Set("authorization", apiKey)
	res, err := client.Do(req)
	defer res.Body.Close()

	if err != nil {
		log.Fatalln(err)
	}

	// Decode json and store it in a map
	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	// Print the upload_url
	fmt.Println(result["upload_url"])

	var AUDIO_URL = fmt.Sprintf("%s", result["upload_url"])
	fmt.Println("AUDIO_URL: ", AUDIO_URL)
	const TRANSCRIPT_URL = "https://api.assemblyai.com/v2/transcript"

	// Prepare json data
	values := map[string]string{"audio_url": AUDIO_URL}
	jsonData, err := json.Marshal(values)

	if err != nil {
		log.Fatalln(err)
	}

	// Setup HTTP client and set header
	client = &http.Client{}
	req, _ = http.NewRequest("POST", TRANSCRIPT_URL, bytes.NewBuffer(jsonData))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", apiKey)
	res, err = client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&result)

	// Print the id of the transcribed audio
	fmt.Println(result["id"])
	var resultId = fmt.Sprintf("%s", result["id"])
	// New endpoint
	var POLLING_URL = TRANSCRIPT_URL + "/" + resultId

	for {
		// Send GET request
		client = &http.Client{}
		req, _ = http.NewRequest("GET", POLLING_URL, nil)
		req.Header.Set("content-type", "application/json")
		req.Header.Set("authorization", apiKey)
		res, err = client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		defer res.Body.Close()

		json.NewDecoder(res.Body).Decode(&result)

		// Check status and print the transcribed text
		if result["status"] == "completed" {
			fmt.Println("Status is completed...")
			fmt.Println(result["text"])
			m.Metadata.Transcript = fmt.Sprintf("%s", result["text"])
			fmt.Println("m.Metadata.Transcript: ", m.Metadata.Transcript)

			break
		} else {
		}
	}

	return nil
}
