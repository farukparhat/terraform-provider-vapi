terraform {
  required_providers {
    vapi = {
      source = "faruk/vapi"
    }
  }
}

provider "vapi" {
  # Configuration options:
  # url   = "https://api.vapi.ai"  # Optional, defaults to https://api.vapi.ai
  # token = "your-api-token"       # Optional, can be set via VAPI_API_KEY environment variable
}

# Basic assistant with minimal configuration
resource "vapi_assistant" "basic" {
  name = "Basic Assistant"

  # Optional: First message the assistant will say
  first_message = "Hello! How can I help you today?"

  # Optional: System message to guide the assistant's behavior
  system_message = "You are a helpful assistant. Be concise and friendly in your responses."
}
