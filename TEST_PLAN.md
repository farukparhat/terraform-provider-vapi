# Test Plan for Basic Assistant Assignment

## Overview

This document outlines how to test the basic assistant assignment to the phone number ID: `0e48e372-6939-410b-8ce1-90c147c963b4`

## Prerequisites

1. Set your VAPI API key: `export VAPI_API_KEY="your-actual-api-key"`
2. Ensure you have access to the phone number with ID `0e48e372-6939-410b-8ce1-90c147c963b4`

## Testing Steps

### Step 1: Import the Existing Phone Number

First, we need to import the existing phone number into Terraform state:

```bash
cd examples/basic-assistant
export TF_CLI_CONFIG_FILE=.terraformrc
export VAPI_API_KEY="your-actual-api-key"

# Import the existing phone number
terraform import vapi_phone_number.existing 0e48e372-6939-410b-8ce1-90c147c963b4
```

### Step 2: Update Configuration with Actual Values

After importing, you'll need to update the `main.tf` file with the actual phone number details:

1. Run `terraform show` to see the current state of the imported phone number
2. Update the `vapi_phone_number.existing` resource in `main.tf` with the actual values:
   - Replace `number = "+1234567890"` with the actual phone number
   - Replace `name = "Support Line"` with the actual name
   - Keep `assistant_id = vapi_assistant.basic.id` to assign the new assistant

### Step 3: Plan and Apply

```bash
# Check what changes will be made
terraform plan

# Apply the changes to assign the assistant
terraform apply
```

### Step 4: Verify the Assignment

After successful application, you should see outputs showing:

- The assistant ID
- The phone number ID
- A confirmation message about the assignment

## Important Notes

### Schema Changes

⚠️ **Important**: The provider schema has been updated to fix Terraform reserved attribute name conflicts:

- `provider` attribute is now `provider_type` in model and voice configurations
- All examples have been updated to use the new attribute name

### Example Configuration

The basic assistant example now:

1. Creates a basic assistant with minimal configuration
2. Imports and updates an existing phone number to use this assistant
3. Provides outputs to confirm the assignment

## Troubleshooting

### Import Issues

If import fails:

- Verify the phone number ID exists in your VAPI account
- Check your API key is correct
- Ensure you have permissions to access the phone number

### Configuration Issues

If plan/apply fails:

- Make sure you've updated the phone number details with actual values after import
- Check that all `provider` attributes have been changed to `provider_type`

### Provider Issues

If you see "provider not found" errors:

- Ensure you're using `export TF_CLI_CONFIG_FILE=.terraformrc`
- The provider should be built locally using `make build`
