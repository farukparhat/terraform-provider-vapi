---
page_title: "Vapi Provider"
subcategory: ""
description: |-
  The Vapi provider enables you to manage Vapi.ai resources using Terraform.
---

# Vapi Provider

The Vapi provider enables you to manage [Vapi.ai](https://vapi.ai) resources using Infrastructure as Code (IaC) principles. Vapi is a platform for building voice AI applications with advanced conversational capabilities.

## Features

- **Assistant Management**: Create, read, update, and delete Vapi assistants
- **Phone Number Management**: Create, read, update, and delete Vapi phone numbers
- **Full Configuration Support**: Configure models, voices, timeouts, and behavior settings
- **Telephony Provider Support**: Integration with Twilio and Vonage for phone number management
- **Environment Variable Support**: Use environment variables for sensitive configuration
- **Import Support**: Import existing assistants and phone numbers into Terraform state

## Example Usage

```terraform
terraform {
  required_providers {
    vapi = {
      source  = "faruk/vapi"
      version = "~> 1.0"
    }
  }
}

provider "vapi" {
  # Configuration options
}

resource "vapi_assistant" "example" {
  name          = "My AI Assistant"
  first_message = "Hello! How can I help you today?"
  system_message = "You are a helpful AI assistant."

  model = {
    provider    = "openai"
    model       = "gpt-4"
    temperature = 0.7
  }

  voice = {
    provider = "elevenlabs"
    voice_id = "21m00Tcm4TlvDq8ikWAM"
  }
}

resource "vapi_phone_number" "example" {
  number       = "+1234567890"
  name         = "Customer Support Line"
  assistant_id = vapi_assistant.example.id
}
```

## Authentication

The provider requires a Vapi.ai API token for authentication. You can obtain this from your [Vapi.ai dashboard](https://dashboard.vapi.ai).

### Environment Variables

You can configure the provider using environment variables:

```bash
export VAPI_API_KEY="your-api-token"
export VAPI_URL="https://api.vapi.ai"  # Optional
```

### Provider Configuration

Alternatively, you can configure the provider directly:

```terraform
provider "vapi" {
  url   = "https://api.vapi.ai"  # Optional, defaults to this value
  token = "your-api-token"       # Optional, can be set via VAPI_API_KEY env var
}
```

## Getting Your API Token

1. Log in to your [Vapi.ai dashboard](https://dashboard.vapi.ai)
2. Navigate to the API section
3. Generate or copy your API token
4. Use it in your provider configuration or set the `VAPI_API_KEY` environment variable

## Schema

### Optional

- `url` (String) Vapi API base URL. Defaults to `https://api.vapi.ai`. Can also be set with the `VAPI_URL` environment variable.
- `token` (String, Sensitive) Vapi API token. Can also be set with the `VAPI_API_KEY` environment variable.
