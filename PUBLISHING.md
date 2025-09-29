# Publishing Guide for terraform-provider-vapi

This guide provides detailed instructions for publishing the Vapi Terraform Provider to the Terraform Registry.

## Prerequisites

Before publishing, ensure you have:

1. **GitHub Repository**: Repository must be public and named `terraform-provider-vapi`
2. **GPG Key**: Required for signing releases
3. **GitHub Account**: With access to the repository
4. **Vapi.ai Account**: For testing the provider

## Repository Structure

The repository is already configured with the required structure:

```
terraform-provider-vapi/
├── docs/                           # Provider documentation
│   ├── index.md                   # Provider overview
│   └── resources/
│       └── assistant.md           # Resource documentation
├── .github/workflows/
│   └── release.yml                # GitHub Actions release workflow
├── .goreleaser.yml                # GoReleaser configuration
├── terraform-registry-manifest.json # Registry metadata
├── public-key.asc                 # GPG public key (if present)
└── ...                           # Other provider files
```

## GPG Key Setup

### 1. Generate a GPG Key (if you don't have one)

```bash
# Generate a new GPG key
gpg --full-generate-key

# Follow the prompts:
# - Select RSA and RSA (default)
# - Use 4096 bits for maximum security
# - Set expiration (recommended: 2 years)
# - Enter your name and email address
# - Create a secure passphrase
```

### 2. Export Your Keys

```bash
# Export public key (for Terraform Registry)
gpg --armor --export "your-email@example.com" > public-key.asc

# Export private key (for GitHub Secrets)
gpg --armor --export-secret-keys "your-email@example.com" > private-key.asc
```

### 3. Add GPG Key to Terraform Registry

1. Sign in to [Terraform Registry](https://registry.terraform.io/) with your GitHub account
2. Go to **User Settings** → **Signing Keys**
3. Add your public key from `public-key.asc`

## GitHub Secrets Configuration

Add these secrets to your GitHub repository (**Settings** → **Secrets and variables** → **Actions**):

| Secret Name       | Value                         | Description                                      |
| ----------------- | ----------------------------- | ------------------------------------------------ |
| `GPG_PRIVATE_KEY` | Contents of `private-key.asc` | Your ASCII-armored GPG private key               |
| `PASSPHRASE`      | Your GPG passphrase           | The passphrase you set when creating the GPG key |

**Important**: Delete the `private-key.asc` file after adding it to GitHub Secrets for security.

## Publishing Process

### 1. Prepare for Release

Ensure your code is ready:

```bash
# Run tests
make test
make testacc

# Format code
make fmt

# Run linter
make lint

# Build locally to verify
make build
```

### 2. Create a Release

```bash
# Ensure you're on main branch with latest changes
git checkout main
git pull origin main

# Create and push a version tag
git tag v1.0.0
git push origin v1.0.0
```

### 3. Monitor the Release

1. **GitHub Actions**: Check the "Actions" tab in your repository to monitor the release workflow
2. **Release Assets**: Verify the release appears in GitHub Releases with proper assets:
   - Binary archives for different platforms
   - Checksums file (`terraform-provider-vapi_v1.0.0_SHA256SUMS`)
   - GPG signature (`terraform-provider-vapi_v1.0.0_SHA256SUMS.sig`)
   - Manifest file (`terraform-provider-vapi_v1.0.0_manifest.json`)

### 4. Publish to Terraform Registry

1. Sign in to [Terraform Registry](https://registry.terraform.io/)
2. Click **Publish** → **Provider**
3. Select your GitHub organization and repository
4. The registry will create a webhook for automatic future releases

## Verification

After publishing, verify your provider:

1. **Registry Page**: Check your provider appears at `https://registry.terraform.io/providers/YOUR_USERNAME/vapi`
2. **Documentation**: Verify documentation renders correctly
3. **Test Installation**: Try using your provider in a test Terraform configuration

## Troubleshooting

### Common Issues

1. **GPG Signing Errors**:

   - Ensure `GPG_PRIVATE_KEY` and `PASSPHRASE` secrets are set correctly
   - Verify your GPG key doesn't require interactive input

2. **Release Workflow Failures**:

   - Check GitHub Actions logs for specific error messages
   - Ensure your repository has proper permissions

3. **Registry Publication Issues**:

   - Verify your repository is public
   - Check that the repository name matches `terraform-provider-{name}`
   - Ensure all required assets are present in the GitHub release

4. **Documentation Not Rendering**:
   - Verify markdown syntax in `docs/` files
   - Check that resource documentation matches your actual resource schema

### Testing Registry Publication

You can test your provider before official publication:

1. Use the [Terraform Registry Doc Preview Tool](https://registry.terraform.io/tools/doc-preview)
2. Test with a private registry if available
3. Use local development builds for functionality testing

## Maintenance

### Updating the Provider

For subsequent releases:

1. Make your changes and test thoroughly
2. Update version references if needed
3. Create a new version tag following semantic versioning
4. The registry will automatically detect and ingest the new version

### Managing Documentation

- Update `docs/index.md` for provider-level changes
- Update `docs/resources/assistant.md` for resource schema changes
- Add new resource documentation files as you add resources

## Resources

- [HashiCorp Provider Publishing Guide](https://developer.hashicorp.com/terraform/registry/providers/publishing)
- [Terraform Registry Doc Preview Tool](https://registry.terraform.io/tools/doc-preview)
- [GoReleaser Documentation](https://goreleaser.com/)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)
