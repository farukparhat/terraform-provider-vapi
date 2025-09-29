#!/bin/bash

# Script to run the assistant with server URL example
echo "🚀 Running Assistant with Server URL - Terraform VAPI Provider"
echo "============================================================="

# Navigate to project root
cd "$(dirname "$0")"

# Load environment variables
if [ -f ".env" ]; then
    source .env
    echo "✅ Environment loaded"
else
    echo "❌ .env file not found. Please create one with your VAPI_API_KEY"
    exit 1
fi

# Check if API key is set
if [ -z "$VAPI_API_KEY" ]; then
    echo "❌ VAPI_API_KEY not set in .env file"
    exit 1
fi

echo "🔑 API Key: ${VAPI_API_KEY:0:10}..."

# Set default webhook URL if not provided
if [ -z "$TF_VAR_webhook_url" ]; then
    export TF_VAR_webhook_url="https://yourapp.com/vapi/webhook"
    echo "🌐 Using default webhook URL: $TF_VAR_webhook_url"
else
    echo "🌐 Using provided webhook URL: $TF_VAR_webhook_url"
fi

# Navigate to assistant with server URL example
cd examples/assistant-with-server-url

# Install provider locally
echo "🔧 Installing provider locally..."
cd ../..
make install > /dev/null 2>&1
cd examples/assistant-with-server-url

# Set up development override
export TF_CLI_CONFIG_FILE=.terraformrc

echo "📋 Planning the deployment..."
terraform plan

if [ $? -eq 0 ]; then
    echo ""
    read -p "🚀 Apply the configuration? (y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        echo "🎯 Creating assistant with server URL..."
        terraform apply -auto-approve

        if [ $? -eq 0 ]; then
            echo ""
            echo "🎉 Success! Your assistant with server URL has been created!"
            echo "📝 Check your VAPI dashboard to see the new assistant"
            echo "🔗 The assistant is configured to send webhooks to: $TF_VAR_webhook_url"
            echo ""
            echo "📊 Configuration Summary:"
            terraform output configuration_summary
        else
            echo "❌ Failed to create assistant. Check the error above."
        fi
    else
        echo "⏹️  Operation cancelled"
    fi
else
    echo "❌ Planning failed. Please check your configuration."
fi
