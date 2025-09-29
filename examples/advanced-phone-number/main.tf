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

# Variables for configuration
variable "twilio_account_sid" {
  description = "Twilio Account SID"
  type        = string
  sensitive   = true
}

variable "twilio_auth_token" {
  description = "Twilio Auth Token"
  type        = string
  sensitive   = true
}

variable "webhook_secret" {
  description = "Webhook secret for verification"
  type        = string
  sensitive   = true
}

variable "phone_number" {
  description = "Phone number in E.164 format"
  type        = string
  default     = "+1234567890"
}

# Create an advanced assistant for phone calls
resource "vapi_assistant" "advanced_phone_assistant" {
  name          = "Advanced Phone Support Assistant"
  first_message = "Hello! Welcome to our support line. I'm an AI assistant here to help you. How can I assist you today?"
  system_message = "You are an advanced customer support AI assistant for phone conversations. You can help with technical issues, billing questions, and general inquiries. Be professional, empathetic, and thorough in your responses while keeping them conversational for voice interaction."

  model = {
    provider_type               = "openai"
    model                      = "gpt-4"
    temperature                = 0.7
    max_tokens                 = 1000
    emotion_recognition_enabled = true
    num_fast_turns             = 2
  }

  voice = {
    provider_type    = "11labs"
    voice_id         = "21m00Tcm4TlvDq8ikWAM"
    speed           = 1.0
    stability       = 0.75
    similarity_boost = 0.8
    style           = 0.5
    use_speaker_boost = true
  }

  client_messages = ["conversation-update", "function-call"]
  server_messages = ["conversation-update", "end-of-call-report"]

  silence_timeout_seconds         = 30
  max_duration_seconds           = 1800  # 30 minutes
  background_denoising_enabled   = true
  model_output_in_messages_enabled = true
}

# Phone number with Twilio provider configuration
resource "vapi_phone_number" "twilio_support_line" {
  number             = var.phone_number
  name               = "Twilio Customer Support Line"
  assistant_id       = vapi_assistant.advanced_phone_assistant.id
  provider_type      = "twilio"
  twilio_account_sid = var.twilio_account_sid
  twilio_auth_token  = var.twilio_auth_token
  
  # Webhook configuration
  server_url        = "https://yourapp.com/vapi/webhook"
  server_url_secret = var.webhook_secret
}

# Alternative phone number for different regions or use cases
resource "vapi_phone_number" "backup_support_line" {
  number       = "+1987654321"  # Replace with your backup number
  name         = "Backup Support Line"
  assistant_id = vapi_assistant.advanced_phone_assistant.id
}

# Outputs
output "primary_phone_number_id" {
  description = "The ID of the primary Twilio phone number"
  value       = vapi_phone_number.twilio_support_line.id
}

output "primary_phone_number" {
  description = "The primary phone number"
  value       = vapi_phone_number.twilio_support_line.number
}

output "backup_phone_number_id" {
  description = "The ID of the backup phone number"
  value       = vapi_phone_number.backup_support_line.id
}

output "backup_phone_number" {
  description = "The backup phone number"
  value       = vapi_phone_number.backup_support_line.number
}

output "assistant_id" {
  description = "The assistant ID handling both phone numbers"
  value       = vapi_assistant.advanced_phone_assistant.id
}
