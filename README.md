# Terraform Provider for Vapi

This Terraform provider enables you to manage [Vapi.ai](https://vapi.ai) resources using Infrastructure as Code (IaC) principles. Currently, the provider supports managing AI assistants.

## Features

- **Assistant Management**: Create, read, update, and delete Vapi assistants
- **Full Configuration Support**: Configure models, voices, timeouts, and behavior settings
- **Environment Variable Support**: Use environment variables for sensitive configuration
- **Import Support**: Import existing assistants into Terraform state

## Requirements

- [Terraform](https://terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.21 (for building the provider)
- A Vapi.ai account and API token

## Installation

### Using Terraform Registry (Recommended)

Add the provider to your Terraform configuration:

```hcl
terraform {
  required_providers {
    vapi = {
      source  = "faruk/vapi"
      version = "~> 1.0"
    }
  }
}
```

### Building from Source

1. Clone the repository:

   ```bash
   git clone https://github.com/faruk/terraform-provider-vapi.git
   cd terraform-provider-vapi
   ```

2. Build and install the provider:

   ```bash
   go build -o terraform-provider-vapi
   mkdir -p ~/.terraform.d/plugins/local/faruk/vapi/1.0.0/darwin_amd64/
   cp terraform-provider-vapi ~/.terraform.d/plugins/local/faruk/vapi/1.0.0/darwin_amd64/
   ```

   Note: Adjust the path based on your OS and architecture.

## Configuration

### Provider Configuration

```hcl
provider "vapi" {
  url   = "https://api.vapi.ai"  # Optional, defaults to this value
  token = "your-api-token"       # Optional, can be set via VAPI_API_KEY env var
}
```

### Environment Variables

You can configure the provider using environment variables:

```bash
export VAPI_API_KEY="your-api-token"
export VAPI_URL="https://api.vapi.ai"  # Optional
```

### Getting Your API Token

1. Log in to your [Vapi.ai dashboard](https://dashboard.vapi.ai)
2. Navigate to the API section
3. Generate or copy your API token
4. Use it in your provider configuration or set the `VAPI_API_KEY` environment variable

## Usage

### Basic Assistant

```hcl
resource "vapi_assistant" "basic" {
  name          = "My Assistant"
  first_message = "Hello! How can I help you today?"
  system_message = "You are a helpful assistant."
}
```

### Advanced Assistant with Full Configuration

```hcl
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
  }

  voice = {
    provider    = "elevenlabs"
    voice_id    = "21m00Tcm4TlvDq8ikWAM"
    speed       = 1.0
    stability   = 0.75
  }

  silence_timeout_seconds = 30
  max_duration_seconds    = 600
  background_denoising_enabled = true
}
```

## Resource Reference

### vapi_assistant

Manages a Vapi assistant.

#### Argument Reference

**Required:**

- `name` (String) - The name of the assistant

**Optional:**

- `first_message` (String) - The first message the assistant will say
- `system_message` (String) - System message to guide the assistant's behavior
- `model` (Object) - Model configuration
  - `provider` (String) - Model provider (e.g., "openai", "anthropic")
  - `model` (String) - Model name (e.g., "gpt-4", "claude-3-sonnet")
  - `temperature` (Number) - Temperature for the model (0.0-2.0)
  - `max_tokens` (Number) - Maximum tokens for the model
  - `emotion_recognition_enabled` (Boolean) - Enable emotion recognition
  - `num_fast_turns` (Number) - Number of fast turns
  - `tool_ids` (List of String) - List of tool IDs
  - `function_ids` (List of String) - List of function IDs
- `voice` (Object) - Voice configuration
  - `provider` (String) - Voice provider (e.g., "elevenlabs", "playht")
  - `voice_id` (String) - Voice ID
  - `speed` (Number) - Voice speed
  - `stability` (Number) - Voice stability
  - `similarity_boost` (Number) - Voice similarity boost
  - `style` (Number) - Voice style
  - `use_speaker_boost` (Boolean) - Enable speaker boost
- `client_messages` (List of String) - List of client messages to send
- `server_messages` (List of String) - List of server messages to receive
- `silence_timeout_seconds` (Number) - Silence timeout in seconds
- `max_duration_seconds` (Number) - Maximum call duration in seconds
- `background_sound` (String) - Background sound setting
- `background_denoising_enabled` (Boolean) - Enable background denoising
- `model_output_in_messages_enabled` (Boolean) - Enable model output in messages

#### Attributes Reference

- `id` (String) - The assistant ID
- `created_at` (String) - Creation timestamp
- `updated_at` (String) - Last update timestamp

#### Import

Assistants can be imported using their ID:

```bash
terraform import vapi_assistant.example "assistant-id-here"
```

## Examples

See the [examples](./examples/) directory for more detailed examples:

- [Basic Assistant](./examples/basic-assistant/) - Minimal configuration
- [Advanced Assistant](./examples/advanced-assistant/) - Full configuration with all options

## Development

### Building the Provider

```bash
go build -o terraform-provider-vapi
```

### Running Tests

```bash
go test ./...
```

### Running Acceptance Tests

```bash
make testacc
```

**Warning**: Acceptance tests create real resources and may incur costs.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for your changes
5. Ensure all tests pass
6. Submit a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- [GitHub Issues](https://github.com/faruk/terraform-provider-vapi/issues) - Bug reports and feature requests
- [Vapi.ai Documentation](https://docs.vapi.ai) - Official Vapi API documentation
- [Terraform Documentation](https://terraform.io/docs) - Terraform usage and best practices
