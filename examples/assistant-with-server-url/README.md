# Assistant with Server URL Example

This example demonstrates how to create a Vapi assistant configured with a `server_url` for receiving webhook events.

## Features

- **Webhook Configuration**: The assistant is configured with a `server_url` to receive webhook events
- **Event Types**: Configures both `server_messages` and `client_messages` to specify which events are sent where
- **Complete Configuration**: Includes model and voice settings for a fully functional assistant

## Usage

1. Set up your environment variables:

   ```bash
   export VAPI_API_KEY="your-vapi-api-key"
   ```

2. (Optional) Set a custom webhook URL:

   ```bash
   export TF_VAR_webhook_url="https://your-webhook-endpoint.com/vapi"
   ```

3. Initialize and apply:
   ```bash
   terraform init
   terraform plan
   terraform apply
   ```

## Webhook Events

The assistant is configured to send the following events to your server:

- `conversation-update`: Updates about the conversation state
- `end-of-call-report`: Summary report when the call ends
- `function-call`: When the assistant calls a function
- `hang`: When the call is hung up
- `speech-update`: Updates about speech recognition

## Variables

- `webhook_url`: The URL where webhook events will be sent (default: "https://yourapp.com/vapi/webhook")

## Outputs

- `assistant_id`: The unique identifier of the created assistant
- `assistant_name`: The name of the assistant
- `server_url`: The configured webhook URL
- `configuration_summary`: A summary of key configuration details

## Testing

To test the webhook functionality:

1. Set up a webhook endpoint that can receive POST requests
2. Configure the `webhook_url` variable to point to your endpoint
3. Create the assistant using Terraform
4. Make a call to test the assistant and observe webhook events at your endpoint

## Clean Up

To remove the created resources:

```bash
terraform destroy
```
