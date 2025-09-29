#!/bin/bash

# Setup script for VAPI Terraform Provider
echo "Setting up environment for VAPI Terraform Provider..."

# Check if .env already exists
if [ -f ".env" ]; then
    echo "⚠️  .env file already exists!"
    read -p "Do you want to overwrite it? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo "Setup cancelled."
        exit 1
    fi
fi

# Copy template
cp .env.example .env

echo "✅ Created .env file from template"
echo ""
echo "📝 Next steps:"
echo "1. Edit .env file and replace 'your-api-key-here' with your actual VAPI API key"
echo "2. Get your API key from: https://dashboard.vapi.ai/"
echo "3. The .env file is ignored by git for security"
echo ""
echo "🚀 Usage:"
echo "   source .env  # Load environment variables"
echo "   terraform plan  # Test your configuration"
