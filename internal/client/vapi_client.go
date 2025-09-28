package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// VapiClient represents a client for the Vapi API
type VapiClient struct {
	BaseURL    string
	Token      string
	HTTPClient *http.Client
}

// NewVapiClient creates a new Vapi API client
func NewVapiClient(baseURL, token string) *VapiClient {
	return &VapiClient{
		BaseURL: baseURL,
		Token:   token,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// Assistant represents a Vapi assistant
type Assistant struct {
	ID                     string                 `json:"id,omitempty"`
	Name                   string                 `json:"name"`
	FirstMessage           string                 `json:"firstMessage,omitempty"`
	SystemMessage          string                 `json:"systemMessage,omitempty"`
	Model                  *AssistantModel        `json:"model,omitempty"`
	Voice                  *AssistantVoice        `json:"voice,omitempty"`
	ClientMessages         []string               `json:"clientMessages,omitempty"`
	ServerMessages         []string               `json:"serverMessages,omitempty"`
	SilenceTimeoutSeconds  *int                   `json:"silenceTimeoutSeconds,omitempty"`
	MaxDurationSeconds     *int                   `json:"maxDurationSeconds,omitempty"`
	BackgroundSound        string                 `json:"backgroundSound,omitempty"`
	BackgroundDenoisingEnabled *bool             `json:"backgroundDenoisingEnabled,omitempty"`
	ModelOutputInMessagesEnabled *bool           `json:"modelOutputInMessagesEnabled,omitempty"`
	TransportConfigurations []map[string]interface{} `json:"transportConfigurations,omitempty"`
	CreatedAt              string                 `json:"createdAt,omitempty"`
	UpdatedAt              string                 `json:"updatedAt,omitempty"`
}

// AssistantModel represents the model configuration for an assistant
type AssistantModel struct {
	Provider             string                 `json:"provider"`
	Model                string                 `json:"model"`
	Temperature          *float64               `json:"temperature,omitempty"`
	MaxTokens            *int                   `json:"maxTokens,omitempty"`
	EmotionRecognitionEnabled *bool            `json:"emotionRecognitionEnabled,omitempty"`
	NumFastTurns         *int                   `json:"numFastTurns,omitempty"`
	ToolIds              []string               `json:"toolIds,omitempty"`
	FunctionIds          []string               `json:"functionIds,omitempty"`
}

// AssistantVoice represents the voice configuration for an assistant
type AssistantVoice struct {
	Provider    string   `json:"provider"`
	VoiceID     string   `json:"voiceId"`
	Speed       *float64 `json:"speed,omitempty"`
	Stability   *float64 `json:"stability,omitempty"`
	SimilarityBoost *float64 `json:"similarityBoost,omitempty"`
	Style       *float64 `json:"style,omitempty"`
	UseSpeakerBoost *bool `json:"useSpeakerBoost,omitempty"`
}

// CreateAssistant creates a new assistant
func (c *VapiClient) CreateAssistant(assistant *Assistant) (*Assistant, error) {
	url := fmt.Sprintf("%s/assistant", c.BaseURL)
	
	jsonData, err := json.Marshal(assistant)
	if err != nil {
		return nil, fmt.Errorf("error marshaling assistant: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("API error: %d - %s", resp.StatusCode, string(body))
	}

	var createdAssistant Assistant
	if err := json.Unmarshal(body, &createdAssistant); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return &createdAssistant, nil
}

// GetAssistant retrieves an assistant by ID
func (c *VapiClient) GetAssistant(id string) (*Assistant, error) {
	url := fmt.Sprintf("%s/assistant/%s", c.BaseURL, id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("assistant not found")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %d - %s", resp.StatusCode, string(body))
	}

	var assistant Assistant
	if err := json.Unmarshal(body, &assistant); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return &assistant, nil
}

// UpdateAssistant updates an existing assistant
func (c *VapiClient) UpdateAssistant(id string, assistant *Assistant) (*Assistant, error) {
	url := fmt.Sprintf("%s/assistant/%s", c.BaseURL, id)

	jsonData, err := json.Marshal(assistant)
	if err != nil {
		return nil, fmt.Errorf("error marshaling assistant: %w", err)
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("assistant not found")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %d - %s", resp.StatusCode, string(body))
	}

	var updatedAssistant Assistant
	if err := json.Unmarshal(body, &updatedAssistant); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return &updatedAssistant, nil
}

// DeleteAssistant deletes an assistant by ID
func (c *VapiClient) DeleteAssistant(id string) error {
	url := fmt.Sprintf("%s/assistant/%s", c.BaseURL, id)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("assistant not found")
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error: %d - %s", resp.StatusCode, string(body))
	}

	return nil
}

// ListAssistants retrieves all assistants
func (c *VapiClient) ListAssistants() ([]Assistant, error) {
	url := fmt.Sprintf("%s/assistant", c.BaseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %d - %s", resp.StatusCode, string(body))
	}

	var assistants []Assistant
	if err := json.Unmarshal(body, &assistants); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return assistants, nil
}
