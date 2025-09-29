#!/bin/bash

# Script to run the basic assistant locally
echo "🚀 Running Basic Assistant with Terraform VAPI Provider"
echo "========================================================"

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

# Navigate to basic assistant example
cd examples/basic-assistant

# Set up development override
export TF_CLI_CONFIG_FILE=.terraformrc

echo "📋 Planning the deployment..."
terraform plan

if [ $? -eq 0 ]; then
    echo ""
    read -p "🚀 Apply the configuration? (y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        echo "🎯 Creating assistant..."
        terraform apply -auto-approve
        
        if [ $? -eq 0 ]; then
            echo ""
            echo "🎉 Success! Your basic assistant has been created!"
            echo "📝 Check your VAPI dashboard to see the new assistant"
        else
            echo "❌ Failed to create assistant. Check the error above."
        fi
    else
        echo "⏹️  Operation cancelled"
    fi
else
    echo "❌ Planning failed. Please check your configuration."
fi
