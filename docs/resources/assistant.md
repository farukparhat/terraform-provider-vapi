---
page_title: "vapi_assistant Resource - terraform-provider-vapi"
subcategory: ""
description: |-
  Manages a Vapi assistant resource.
---

# vapi_assistant (Resource)

Manages a Vapi assistant. Assistants are AI-powered conversational agents that can be configured with different models, voices, and behaviors to create customized voice AI experiences.

## Example Usage

### Basic Assistant

```terraform
resource "vapi_assistant" "basic" {
  name          = "Basic Assistant"
  first_message = "Hello! How can I help you today?"
  system_message = "You are a helpful assistant."
}
```

### Advanced Assistant with Full Configuration

```terraform
resource "vapi_assistant" "advanced" {
  name          = "Advanced Assistant"
  first_message = "Welcome! I'm your advanced AI assistant."
  system_message = "You are an expert AI assistant with multiple capabilities."

  model = {
    provider    = "openai"
    model       = "gpt-4"
    temperature = 0.7
    max_tokens  = 1000
    emotion_recognition_enabled = true
    num_fast_turns = 2
    tool_ids = ["tool1", "tool2"]
    function_ids = ["func1", "func2"]
  }

  voice = {
    provider         = "11labs"
    voice_id         = "21m00Tcm4TlvDq8ikWAM"
    speed           = 1.0
    stability       = 0.75
    similarity_boost = 0.8
    style           = 0.5
    use_speaker_boost = true
  }

  client_messages = ["conversation-update", "function-call"]
  server_messages = ["conversation-update", "end-of-call-report"]

  silence_timeout_seconds = 30
  max_duration_seconds    = 600
  background_sound = "office"
  background_denoising_enabled = true
  model_output_in_messages_enabled = true
}
```

### Assistant with Custom Timeouts

```terraform
resource "vapi_assistant" "timed" {
  name          = "Timed Assistant"
  first_message = "Hi! I have limited time to chat."

  silence_timeout_seconds = 15
  max_duration_seconds    = 300
}
```

### Assistant with Webhook Server URL

```terraform
resource "vapi_assistant" "webhook_assistant" {
  name          = "Webhook Assistant"
  first_message = "Hello! I'm configured with webhook support."
  system_message = "You are a helpful assistant with webhook event reporting."

  # Configure webhook URL for receiving events
  server_url = "https://yourapp.com/vapi/webhook"

  # Specify which events to send to the server
  server_messages = [
    "conversation-update",
    "end-of-call-report",
    "function-call"
  ]
}
```

## Schema

### Required

- `name` (String) The name of the assistant.

### Optional

- `background_denoising_enabled` (Boolean) Whether background denoising is enabled.
- `background_sound` (String) Background sound setting for the assistant.
- `client_messages` (List of String) List of client messages to send during the conversation.
- `first_message` (String) The first message the assistant will say when the conversation starts.
- `max_duration_seconds` (Number) Maximum duration of the conversation in seconds.
- `model` (Object) Configuration for the AI model used by the assistant. See [model](#nested-schema-for-model) below.
- `model_output_in_messages_enabled` (Boolean) Whether model output should be included in messages.
- `server_messages` (List of String) List of server messages to receive during the conversation.
- `server_url` (String) Server URL for webhook events. When set, the assistant will send configured events to this endpoint.
- `silence_timeout_seconds` (Number) Timeout in seconds before ending the conversation due to silence.
- `system_message` (String) System message that guides the assistant's behavior and personality.
- `voice` (Object) Configuration for the voice used by the assistant. See [voice](#nested-schema-for-voice) below.

### Read-Only

- `created_at` (String) Timestamp when the assistant was created.
- `id` (String) The unique identifier of the assistant.
- `updated_at` (String) Timestamp when the assistant was last updated.

### Nested Schema for `model`

Optional:

- `emotion_recognition_enabled` (Boolean) Whether emotion recognition is enabled for the model.
- `function_ids` (List of String) List of function IDs available to the model.
- `max_tokens` (Number) Maximum number of tokens the model can generate.
- `model` (String) The specific model to use (e.g., "gpt-4", "claude-3-sonnet").
- `num_fast_turns` (Number) Number of fast turns for the model.
- `provider` (String) The model provider (e.g., "openai", "anthropic").
- `temperature` (Number) Temperature setting for the model, controlling randomness (0.0-2.0).
- `tool_ids` (List of String) List of tool IDs available to the model.

### Nested Schema for `voice`

Optional:

- `provider` (String) The voice provider (e.g., "11labs", "playht").
- `similarity_boost` (Number) Similarity boost setting for the voice.
- `speed` (Number) Speed of the voice.
- `stability` (Number) Stability setting for the voice.
- `style` (Number) Style setting for the voice.
- `use_speaker_boost` (Boolean) Whether speaker boost is enabled.
- `voice_id` (String) The specific voice ID to use.

## Import

Import is supported using the following syntax:

```shell
terraform import vapi_assistant.example "assistant-id-here"
```
