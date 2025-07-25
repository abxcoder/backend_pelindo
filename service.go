package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type ExternalService interface {
	GetNegaras() ([]Negara, error)
	GetPelabuhans(idNegara string) ([]Pelabuhan, error)
	GetBarangs(idPelabuhan string) ([]Barang, error)
}

type externalService struct {
	client  *http.Client
	baseURL string
}

func NewExternalService() ExternalService {
	baseURL := os.Getenv("EXTERNAL_API_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080/api/v1"
	}

	return &externalService{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: baseURL,
	}
}

func (s *externalService) GetNegaras() ([]Negara, error) {
	url := fmt.Sprintf("%s/negaras", s.baseURL)

	fmt.Printf("Fetching negaras from: %s\n", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Pelabuhan-Service/1.0")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request to %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned status code: %d for URL: %s, response: %s", resp.StatusCode, url, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	fmt.Printf("External API Response for negaras: %s\n", string(body))

	var apiResponse struct {
		Status  string   `json:"status"`
		Message string   `json:"message"`
		Data    []Negara `json:"data"`
	}

	if err := json.Unmarshal(body, &apiResponse); err == nil && apiResponse.Status == "success" {
		return s.validateNegaras(apiResponse.Data), nil
	}

	var negaras []Negara
	if err := json.Unmarshal(body, &negaras); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w, response body: %s", err, string(body))
	}

	return s.validateNegaras(negaras), nil
}

func (s *externalService) GetPelabuhans(idNegara string) ([]Pelabuhan, error) {
	url := fmt.Sprintf("%s/pelabuhans?id_negara=%s", s.baseURL, idNegara)

	fmt.Printf("Fetching pelabuhans from: %s\n", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Pelabuhan-Service/1.0")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request to %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned status code: %d for URL: %s, response: %s", resp.StatusCode, url, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	fmt.Printf("External API Response for pelabuhans (id_negara=%s): %s\n", idNegara, string(body))

	var apiResponse struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Data    []Pelabuhan `json:"data"`
	}

	if err := json.Unmarshal(body, &apiResponse); err == nil && apiResponse.Status == "success" {
		// Jika berhasil parse sebagai wrapped response
		return s.validatePelabuhans(apiResponse.Data, idNegara), nil
	}

	var pelabuhans []Pelabuhan
	if err := json.Unmarshal(body, &pelabuhans); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w, response body: %s", err, string(body))
	}

	return s.validatePelabuhans(pelabuhans, idNegara), nil
}

func (s *externalService) GetBarangs(idPelabuhan string) ([]Barang, error) {
	url := fmt.Sprintf("%s/barangs?id_pelabuhan=%s", s.baseURL, idPelabuhan)

	fmt.Printf("Fetching barangs from: %s\n", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Pelabuhan-Service/1.0")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request to %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned status code: %d for URL: %s, response: %s", resp.StatusCode, url, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	fmt.Printf("External API Response for barangs (id_pelabuhan=%s): %s\n", idPelabuhan, string(body))

	var apiResponse struct {
		Status  string   `json:"status"`
		Message string   `json:"message"`
		Data    []Barang `json:"data"`
	}

	if err := json.Unmarshal(body, &apiResponse); err == nil && apiResponse.Status == "success" {
		return s.validateBarangs(apiResponse.Data), nil
	}

	var barangs []Barang
	if err := json.Unmarshal(body, &barangs); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w, response body: %s", err, string(body))
	}

	return s.validateBarangs(barangs), nil
}

func (s *externalService) validateNegaras(negaras []Negara) []Negara {
	var validNegaras []Negara
	for _, negara := range negaras {
		if negara.IDNegara > 0 &&
			strings.TrimSpace(negara.NamaNegara) != "" &&
			strings.TrimSpace(negara.KodeNegara) != "" {

			negara.NamaNegara = strings.TrimSpace(strings.ReplaceAll(negara.NamaNegara, "\r\n", ""))
			negara.KodeNegara = strings.TrimSpace(strings.ReplaceAll(negara.KodeNegara, "\r\n", ""))
			validNegaras = append(validNegaras, negara)
		}
	}
	fmt.Printf("Validated %d out of %d negaras\n", len(validNegaras), len(negaras))
	return validNegaras
}

func (s *externalService) validatePelabuhans(pelabuhans []Pelabuhan, idNegara string) []Pelabuhan {
	var validPelabuhans []Pelabuhan
	for _, pelabuhan := range pelabuhans {
		if strings.TrimSpace(pelabuhan.IDPelabuhan) != "" &&
			strings.TrimSpace(pelabuhan.NamaPelabuhan) != "" &&
			pelabuhan.IDNegara == idNegara {
			// Clean up data
			pelabuhan.NamaPelabuhan = strings.TrimSpace(strings.ReplaceAll(pelabuhan.NamaPelabuhan, "\r\n", ""))
			validPelabuhans = append(validPelabuhans, pelabuhan)
		}
	}
	fmt.Printf("Validated %d out of %d pelabuhans for id_negara=%s\n", len(validPelabuhans), len(pelabuhans), idNegara)
	return validPelabuhans
}

func (s *externalService) validateBarangs(barangs []Barang) []Barang {
	var validBarangs []Barang
	for _, barang := range barangs {
		if barang.IDBarang > 0 &&
			strings.TrimSpace(barang.NamaBarang) != "" &&
			barang.Harga > 0 {
			barang.NamaBarang = strings.TrimSpace(strings.ReplaceAll(barang.NamaBarang, "\r\n", ""))
			barang.Description = strings.TrimSpace(strings.ReplaceAll(barang.Description, "\r\n", ""))
			validBarangs = append(validBarangs, barang)
		}
	}
	fmt.Printf("Validated %d out of %d barangs\n", len(validBarangs), len(barangs))
	return validBarangs
}
