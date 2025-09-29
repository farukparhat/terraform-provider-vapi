#!/bin/bash

# Script to run the basic phone number example
# Make sure to set your VAPI_API_KEY environment variable before running this script

set -e

echo "Running basic phone number example..."

# Check if VAPI_API_KEY is set
if [ -z "$VAPI_API_KEY" ]; then
    echo "Error: VAPI_API_KEY environment variable is not set"
    echo "Please set it with: export VAPI_API_KEY=your_api_key"
    exit 1
fi

# Navigate to the basic phone number example directory
cd examples/basic-phone-number

echo "Initializing Terraform..."
terraform init

echo "Planning Terraform changes..."
terraform plan

echo ""
echo "To apply the changes, run:"
echo "cd examples/basic-phone-number && terraform apply"
echo ""
echo "Don't forget to update the phone number in main.tf with your actual phone number!"
