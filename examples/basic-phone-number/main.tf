terraform {
  required_providers {
    vapi = {
      source = "terraform-provider-vapi/vapi"
    }
  }
}

provider "vapi" {
  # Configuration will be taken from environment variables:
  # VAPI_API_KEY
  # VAPI_URL (optional, defaults to https://api.vapi.ai)
}

# Create a basic assistant for phone calls
resource "vapi_assistant" "phone_assistant" {
  name          = "Phone Support Assistant"
  first_message = "Hello! Thank you for calling. How can I help you today?"
  system_message = "You are a helpful customer support assistant for a phone conversation. Be concise and clear in your responses."

  model = {
    provider_type = "openai"
    model    = "gpt-4o-mini"
    temperature = 0.7
  }

  voice = {
    provider_type = "11labs"
    voice_id = "21m00Tcm4TlvDq8ikWAM"
    speed    = 1.0
  }
}

# Create a basic phone number
resource "vapi_phone_number" "support_line" {
  number       = "+1234567890"  # Replace with your actual phone number
  name         = "Customer Support Line"
  assistant_id = vapi_assistant.phone_assistant.id
}

# Output the phone number details
output "phone_number_id" {
  description = "The ID of the created phone number"
  value       = vapi_phone_number.support_line.id
}

output "phone_number" {
  description = "The phone number"
  value       = vapi_phone_number.support_line.number
}

output "assistant_id" {
  description = "The assistant ID handling this phone number"
  value       = vapi_phone_number.support_line.assistant_id
}
