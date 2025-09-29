terraform {
  required_providers {
    vapi = {
      source = "local/faruk/vapi"
    }
  }
}

provider "vapi" {
  # Configuration options:
  # url   = "https://api.vapi.ai"  # Optional, defaults to https://api.vapi.ai
  # token = "your-api-token"       # Optional, can be set via VAPI_API_KEY environment variable
}

# Variable for webhook URL
variable "webhook_url" {
  description = "Webhook URL for receiving events from the assistant"
  type        = string
  default     = "https://yourapp.com/vapi/webhook"
}

# Assistant with server_url configuration for webhook events
resource "vapi_assistant" "webhook_assistant" {
  name = "Webhook Assistant"

  # First message the assistant will say
  first_message = "Hello! I'm an assistant configured with webhook support. How can I help you today?"

  # System message to guide the assistant's behavior
  system_message = "You are a helpful assistant with webhook event reporting. Be concise and friendly in your responses."

  # Server URL for webhook events
  server_url = var.webhook_url

  # Optional model configuration
  model = {
    provider_type = "openai"
    model         = "gpt-4o-mini"
    temperature   = 0.7
    max_tokens    = 500
  }

  # Optional voice configuration
  voice = {
    provider_type = "11labs"
    voice_id      = "21m00Tcm4TlvDq8ikWAM"
    speed         = 1.0
    stability     = 0.75
  }

  # Configure which events to send to the server
  server_messages = [
    "conversation-update",
    "end-of-call-report",
    "function-call",
    "hang",
    "speech-update"
  ]

  # Configure which events to send to the client
  client_messages = [
    "conversation-update",
    "function-call",
    "hang",
    "speech-update"
  ]

  # Additional settings
  silence_timeout_seconds         = 30
  max_duration_seconds           = 600
  background_denoising_enabled   = true
  model_output_in_messages_enabled = true
}

# Outputs
output "assistant_id" {
  description = "The ID of the webhook-enabled assistant"
  value       = vapi_assistant.webhook_assistant.id
}

output "assistant_name" {
  description = "The name of the assistant"
  value       = vapi_assistant.webhook_assistant.name
}

output "server_url" {
  description = "The webhook URL configured for the assistant"
  value       = vapi_assistant.webhook_assistant.server_url
}

output "configuration_summary" {
  description = "Summary of the assistant configuration"
  value = {
    id          = vapi_assistant.webhook_assistant.id
    name        = vapi_assistant.webhook_assistant.name
    server_url  = vapi_assistant.webhook_assistant.server_url
    model       = "gpt-4o-mini"
    voice       = "11labs"
  }
}
