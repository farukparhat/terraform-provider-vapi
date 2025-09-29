---
page_title: "vapi_phone_number Resource - terraform-provider-vapi"
subcategory: ""
description: |-
  Manages a Vapi phone number resource.
---

# vapi_phone_number (Resource)

Manages a Vapi phone number. Phone numbers can be configured to route incoming calls to specific assistants or squads, and can be integrated with various telephony providers like Twilio and Vonage.

## Example Usage

### Basic Phone Number

```terraform
resource "vapi_phone_number" "basic" {
  number = "+1234567890"
  name   = "Customer Support Line"
}
```

### Phone Number with Assistant

```terraform
resource "vapi_assistant" "support" {
  name          = "Support Assistant"
  first_message = "Hello! How can I help you today?"
  system_message = "You are a helpful customer support assistant."
}

resource "vapi_phone_number" "support_line" {
  number       = "+1234567890"
  name         = "Customer Support Line"
  assistant_id = vapi_assistant.support.id
}
```

### Phone Number with Twilio Provider

```terraform
resource "vapi_phone_number" "twilio_number" {
  number             = "+1234567890"
  name               = "Twilio Support Line"
  provider           = "twilio"
  twilio_account_sid = var.twilio_account_sid
  twilio_auth_token  = var.twilio_auth_token
  assistant_id       = vapi_assistant.support.id
}
```

### Phone Number with Vonage Provider

```terraform
resource "vapi_phone_number" "vonage_number" {
  number                  = "+1234567890"
  name                    = "Vonage Support Line"
  provider                = "vonage"
  vonage_api_key          = var.vonage_api_key
  vonage_api_secret       = var.vonage_api_secret
  vonage_application_id   = var.vonage_application_id
  assistant_id            = vapi_assistant.support.id
}
```

### Phone Number with Webhooks

```terraform
resource "vapi_phone_number" "webhook_number" {
  number              = "+1234567890"
  name                = "Webhook Support Line"
  assistant_id        = vapi_assistant.support.id
  server_url          = "https://yourapp.com/vapi/webhook"
  server_url_secret   = var.webhook_secret
}
```

### Phone Number with Squad

```terraform
resource "vapi_phone_number" "squad_number" {
  number   = "+1234567890"
  name     = "Squad Support Line"
  squad_id = "squad_abc123"
}
```

## Schema

### Required

- `number` (String) Phone number in E.164 format (e.g., +1234567890).

### Optional

- `assistant_id` (String) Assistant ID to handle calls on this number.
- `name` (String) Display name for the phone number.
- `provider` (String) Telephony provider (twilio, vonage).
- `server_url` (String) Server URL for webhooks.
- `server_url_secret` (String, Sensitive) Secret for server URL webhook verification.
- `squad_id` (String) Squad ID to handle calls on this number.
- `twilio_account_sid` (String, Sensitive) Twilio Account SID (required if provider is twilio).
- `twilio_auth_token` (String, Sensitive) Twilio Auth Token (required if provider is twilio).
- `vonage_api_key` (String, Sensitive) Vonage API Key (required if provider is vonage).
- `vonage_api_secret` (String, Sensitive) Vonage API Secret (required if provider is vonage).
- `vonage_application_id` (String) Vonage Application ID (required if provider is vonage).

### Read-Only

- `created_at` (String) Creation timestamp.
- `id` (String) Phone number identifier.
- `updated_at` (String) Last update timestamp.

## Import

Phone numbers can be imported using the phone number ID:

```shell
terraform import vapi_phone_number.example phone_number_id
```

## Notes

- Either `assistant_id` or `squad_id` should be specified to handle incoming calls, but not both.
- When using a specific telephony provider (twilio or vonage), the corresponding credentials must be provided.
- The `number` must be in E.164 format (starting with + followed by country code and number).
- Sensitive fields like authentication tokens and secrets are not exposed in state refresh operations for security reasons.
