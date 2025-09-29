package provider

import (
	"context"
	"fmt"

	"terraform-provider-vapi/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &PhoneNumberResource{}
var _ resource.ResourceWithImportState = &PhoneNumberResource{}

func NewPhoneNumberResource() resource.Resource {
	return &PhoneNumberResource{}
}

// PhoneNumberResource defines the resource implementation.
type PhoneNumberResource struct {
	client *client.VapiClient
}

// PhoneNumberResourceModel describes the resource data model.
type PhoneNumberResourceModel struct {
	ID                  types.String `tfsdk:"id"`
	Number              types.String `tfsdk:"number"`
	Name                types.String `tfsdk:"name"`
	AssistantID         types.String `tfsdk:"assistant_id"`
	SquadID             types.String `tfsdk:"squad_id"`
	ServerURL           types.String `tfsdk:"server_url"`
	ServerURLSecret     types.String `tfsdk:"server_url_secret"`
	ProviderType        types.String `tfsdk:"provider_type"`
	TwilioAccountSid    types.String `tfsdk:"twilio_account_sid"`
	TwilioAuthToken     types.String `tfsdk:"twilio_auth_token"`
	VonageAPIKey        types.String `tfsdk:"vonage_api_key"`
	VonageAPISecret     types.String `tfsdk:"vonage_api_secret"`
	VonageApplicationID types.String `tfsdk:"vonage_application_id"`
	CreatedAt           types.String `tfsdk:"created_at"`
	UpdatedAt           types.String `tfsdk:"updated_at"`
}

func (r *PhoneNumberResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_phone_number"
}

func (r *PhoneNumberResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Vapi Phone Number resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Phone number identifier",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"number": schema.StringAttribute{
				MarkdownDescription: "Phone number in E.164 format (e.g., +1234567890)",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Display name for the phone number",
				Optional:            true,
			},
			"assistant_id": schema.StringAttribute{
				MarkdownDescription: "Assistant ID to handle calls on this number",
				Optional:            true,
			},
			"squad_id": schema.StringAttribute{
				MarkdownDescription: "Squad ID to handle calls on this number",
				Optional:            true,
			},
			"server_url": schema.StringAttribute{
				MarkdownDescription: "Server URL for webhooks",
				Optional:            true,
			},
			"server_url_secret": schema.StringAttribute{
				MarkdownDescription: "Secret for server URL webhook verification",
				Optional:            true,
				Sensitive:           true,
			},
			"provider_type": schema.StringAttribute{
				MarkdownDescription: "Telephony provider (twilio, vonage)",
				Optional:            true,
			},
			"twilio_account_sid": schema.StringAttribute{
				MarkdownDescription: "Twilio Account SID (required if provider is twilio)",
				Optional:            true,
				Sensitive:           true,
			},
			"twilio_auth_token": schema.StringAttribute{
				MarkdownDescription: "Twilio Auth Token (required if provider is twilio)",
				Optional:            true,
				Sensitive:           true,
			},
			"vonage_api_key": schema.StringAttribute{
				MarkdownDescription: "Vonage API Key (required if provider is vonage)",
				Optional:            true,
				Sensitive:           true,
			},
			"vonage_api_secret": schema.StringAttribute{
				MarkdownDescription: "Vonage API Secret (required if provider is vonage)",
				Optional:            true,
				Sensitive:           true,
			},
			"vonage_application_id": schema.StringAttribute{
				MarkdownDescription: "Vonage Application ID (required if provider is vonage)",
				Optional:            true,
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "Creation timestamp",
				Computed:            true,
				Optional:            true,
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: "Last update timestamp",
				Computed:            true,
				Optional:            true,
			},
		},
	}
}

func (r *PhoneNumberResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.VapiClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.VapiClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *PhoneNumberResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data PhoneNumberResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Convert Terraform model to API model
	phoneNumber := &client.PhoneNumber{
		Number: data.Number.ValueString(),
	}

	if !data.Name.IsNull() {
		phoneNumber.Name = data.Name.ValueString()
	}

	if !data.AssistantID.IsNull() {
		phoneNumber.AssistantID = data.AssistantID.ValueString()
	}

	if !data.SquadID.IsNull() {
		phoneNumber.SquadID = data.SquadID.ValueString()
	}

	if !data.ServerURL.IsNull() {
		phoneNumber.ServerURL = data.ServerURL.ValueString()
	}

	if !data.ServerURLSecret.IsNull() {
		phoneNumber.ServerURLSecret = data.ServerURLSecret.ValueString()
	}

	if !data.ProviderType.IsNull() {
		phoneNumber.Provider = data.ProviderType.ValueString()
	}

	if !data.TwilioAccountSid.IsNull() {
		phoneNumber.TwilioAccountSid = data.TwilioAccountSid.ValueString()
	}

	if !data.TwilioAuthToken.IsNull() {
		phoneNumber.TwilioAuthToken = data.TwilioAuthToken.ValueString()
	}

	if !data.VonageAPIKey.IsNull() {
		phoneNumber.VonageAPIKey = data.VonageAPIKey.ValueString()
	}

	if !data.VonageAPISecret.IsNull() {
		phoneNumber.VonageAPISecret = data.VonageAPISecret.ValueString()
	}

	if !data.VonageApplicationID.IsNull() {
		phoneNumber.VonageApplicationID = data.VonageApplicationID.ValueString()
	}

	// Create the phone number
	createdPhoneNumber, err := r.client.CreatePhoneNumber(phoneNumber)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create phone number, got error: %s", err))
		return
	}

	// Update the model with the created phone number data
	data.ID = types.StringValue(createdPhoneNumber.ID)
	
	// For now, set timestamp fields to null since VAPI API may not return them consistently
	// This prevents "unknown value" errors while keeping the fields available
	data.CreatedAt = types.StringNull()
	data.UpdatedAt = types.StringNull()

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PhoneNumberResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data PhoneNumberResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get the phone number from the API
	phoneNumber, err := r.client.GetPhoneNumber(data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read phone number, got error: %s", err))
		return
	}

	// Update the model with the phone number data
	data.Number = types.StringValue(phoneNumber.Number)
	data.Name = types.StringValue(phoneNumber.Name)
	data.AssistantID = types.StringValue(phoneNumber.AssistantID)
	data.SquadID = types.StringValue(phoneNumber.SquadID)
	data.ServerURL = types.StringValue(phoneNumber.ServerURL)
	data.ProviderType = types.StringValue(phoneNumber.Provider)
	data.TwilioAccountSid = types.StringValue(phoneNumber.TwilioAccountSid)
	data.VonageAPIKey = types.StringValue(phoneNumber.VonageAPIKey)
	data.VonageApplicationID = types.StringValue(phoneNumber.VonageApplicationID)
	// For now, set timestamp fields to null since VAPI API may not return them consistently
	// This prevents "unknown value" errors while keeping the fields available
	data.CreatedAt = types.StringNull()
	data.UpdatedAt = types.StringNull()

	// Note: Sensitive fields like secrets and tokens are not updated from API response for security

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PhoneNumberResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data PhoneNumberResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Convert Terraform model to API model for update
	// Note: Number field is immutable and should not be included in updates
	phoneNumber := &client.PhoneNumber{}

	if !data.Name.IsNull() {
		phoneNumber.Name = data.Name.ValueString()
	}

	if !data.AssistantID.IsNull() {
		phoneNumber.AssistantID = data.AssistantID.ValueString()
	}

	if !data.SquadID.IsNull() {
		phoneNumber.SquadID = data.SquadID.ValueString()
	}

	if !data.ServerURL.IsNull() {
		phoneNumber.ServerURL = data.ServerURL.ValueString()
	}

	if !data.ServerURLSecret.IsNull() {
		phoneNumber.ServerURLSecret = data.ServerURLSecret.ValueString()
	}

	if !data.ProviderType.IsNull() {
		phoneNumber.Provider = data.ProviderType.ValueString()
	}

	if !data.TwilioAccountSid.IsNull() {
		phoneNumber.TwilioAccountSid = data.TwilioAccountSid.ValueString()
	}

	if !data.TwilioAuthToken.IsNull() {
		phoneNumber.TwilioAuthToken = data.TwilioAuthToken.ValueString()
	}

	if !data.VonageAPIKey.IsNull() {
		phoneNumber.VonageAPIKey = data.VonageAPIKey.ValueString()
	}

	if !data.VonageAPISecret.IsNull() {
		phoneNumber.VonageAPISecret = data.VonageAPISecret.ValueString()
	}

	if !data.VonageApplicationID.IsNull() {
		phoneNumber.VonageApplicationID = data.VonageApplicationID.ValueString()
	}

	// Update the phone number
	_, err := r.client.UpdatePhoneNumber(data.ID.ValueString(), phoneNumber)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update phone number, got error: %s", err))
		return
	}

	// Update the model with the updated phone number data
	// For now, set timestamp fields to null since VAPI API may not return them consistently
	// This prevents "unknown value" errors while keeping the fields available
	data.CreatedAt = types.StringNull()
	data.UpdatedAt = types.StringNull()

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PhoneNumberResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data PhoneNumberResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete the phone number
	err := r.client.DeletePhoneNumber(data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete phone number, got error: %s", err))
		return
	}
}

func (r *PhoneNumberResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
