package provider

import (
	"context"
	"os"

	"terraform-provider-vapi/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure VapiProvider satisfies various provider interfaces.
var _ provider.Provider = &VapiProvider{}

// VapiProvider defines the provider implementation.
type VapiProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// VapiProviderModel describes the provider data model.
type VapiProviderModel struct {
	URL    types.String `tfsdk:"url"`
	ApiKey types.String `tfsdk:"api_key"`
}

func (p *VapiProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "vapi"
	resp.Version = p.version
}

func (p *VapiProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"url": schema.StringAttribute{
				MarkdownDescription: "Vapi API base URL",
				Optional:            true,
			},
			"api_key": schema.StringAttribute{
				MarkdownDescription: "Vapi API key",
				Optional:            true,
				Sensitive:           true,
			},
		},
	}
}

func (p *VapiProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data VapiProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// Example client configuration for data sources and resources
	url := data.URL.ValueString()
	if url == "" {
		url = os.Getenv("VAPI_URL")
		if url == "" {
			url = "https://api.vapi.ai"
		}
	}

	apiKey := data.ApiKey.ValueString()
	if apiKey == "" {
		apiKey = os.Getenv("VAPI_API_KEY")
	}

	if apiKey == "" {
		resp.Diagnostics.AddError(
			"Unable to find API key",
			"API key cannot be an empty string. Please set the api_key in the provider configuration or set the VAPI_API_KEY environment variable.",
		)
		return
	}

	// Example client configuration for data sources and resources
	client := client.NewVapiClient(url, apiKey)
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *VapiProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewAssistantResource,
		NewPhoneNumberResource,
	}
}

func (p *VapiProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		// Add data sources here when needed
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &VapiProvider{
			version: version,
		}
	}
}
