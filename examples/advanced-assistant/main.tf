terraform {
  required_providers {
    vapi = {
      source = "faruk/vapi"
    }
  }
}

provider "vapi" {
  # Configuration can be set via environment variables:
  # export VAPI_API_KEY="your-api-token"
  # export VAPI_URL="https://api.vapi.ai"  # Optional
}

# Advanced assistant with full configuration
resource "vapi_assistant" "advanced" {
  name           = "Advanced AI Assistant"
  first_message  = "Welcome! I'm your advanced AI assistant. I can help you with a variety of tasks."
  system_message = <<EOF
You are an advanced AI assistant with expertise in multiple domains. 
You should:
- Be helpful, harmless, and honest
- Provide detailed explanations when appropriate
- Ask clarifying questions when needed
- Maintain a professional but friendly tone
EOF

  # Model configuration
  model = {
    provider                    = "openai"
    model                       = "gpt-4"
    temperature                 = 0.7
    max_tokens                  = 1000
    emotion_recognition_enabled = true
    num_fast_turns              = 2
  }

  # Voice configuration
  voice = {
    provider          = "elevenlabs"
    voice_id          = "21m00Tcm4TlvDq8ikWAM" # Example voice ID
    speed             = 1.0
    stability         = 0.75
    similarity_boost  = 0.8
    use_speaker_boost = true
  }

  # Timeout and duration settings
  silence_timeout_seconds = 30
  max_duration_seconds    = 600

  # Background settings
  background_sound                 = "office"
  background_denoising_enabled     = true
  model_output_in_messages_enabled = true

  # Message configuration
  client_messages = [
    "conversation-update",
    "function-call",
    "hang",
    "speech-update"
  ]

  server_messages = [
    "conversation-update",
    "end-of-call-report",
    "function-call",
    "hang",
    "speech-update"
  ]
}
