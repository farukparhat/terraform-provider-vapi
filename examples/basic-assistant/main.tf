terraform {
  required_providers {
    vapi = {
      source = "terraform-provider-vapi/vapi"
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

# Import and update an existing phone number to use this assistant
# First, import the existing phone number using: terraform import vapi_phone_number.existing 0e48e372-6939-410b-8ce1-90c147c963b4
resource "vapi_phone_number" "existing" {
  # Note: After importing, you only need to specify the fields you want to update
  # The number field should match what was imported and will be managed by Terraform
  number       = "+1234567890"  # This should match the imported phone number exactly
  name         = "Terraform"     # You can update the name if desired
  assistant_id = vapi_assistant.basic.id  # This assigns the new assistant
}

# Outputs
output "assistant_id" {
  description = "The ID of the basic assistant"
  value       = vapi_assistant.basic.id
}

output "phone_number_id" {
  description = "The ID of the phone number"
  value       = vapi_phone_number.existing.id
}

output "assignment_status" {
  description = "Confirmation that assistant is assigned to phone number"
  value       = "Assistant ${vapi_assistant.basic.id} assigned to phone number ${vapi_phone_number.existing.id}"
}
