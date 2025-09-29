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
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AssistantResource{}
var _ resource.ResourceWithImportState = &AssistantResource{}

func NewAssistantResource() resource.Resource {
	return &AssistantResource{}
}

// AssistantResource defines the resource implementation.
type AssistantResource struct {
	client *client.VapiClient
}

// AssistantResourceModel describes the resource data model.
type AssistantResourceModel struct {
	ID                           types.String `tfsdk:"id"`
	Name                         types.String `tfsdk:"name"`
	FirstMessage                 types.String `tfsdk:"first_message"`
	SystemMessage                types.String `tfsdk:"system_message"`
	Model                        types.Object `tfsdk:"model"`
	Voice                        types.Object `tfsdk:"voice"`
	ClientMessages               types.List   `tfsdk:"client_messages"`
	ServerMessages               types.List   `tfsdk:"server_messages"`
	SilenceTimeoutSeconds        types.Int64  `tfsdk:"silence_timeout_seconds"`
	MaxDurationSeconds           types.Int64  `tfsdk:"max_duration_seconds"`
	BackgroundSound              types.String `tfsdk:"background_sound"`
	BackgroundDenoisingEnabled   types.Bool   `tfsdk:"background_denoising_enabled"`
	ModelOutputInMessagesEnabled types.Bool   `tfsdk:"model_output_in_messages_enabled"`
	CreatedAt                    types.String `tfsdk:"created_at"`
	UpdatedAt                    types.String `tfsdk:"updated_at"`
}

// AssistantModelModel describes the model configuration
type AssistantModelModel struct {
	ProviderType              types.String  `tfsdk:"provider_type"`
	Model                     types.String  `tfsdk:"model"`
	Temperature               types.Float64 `tfsdk:"temperature"`
	MaxTokens                 types.Int64   `tfsdk:"max_tokens"`
	EmotionRecognitionEnabled types.Bool    `tfsdk:"emotion_recognition_enabled"`
	NumFastTurns              types.Int64   `tfsdk:"num_fast_turns"`
	ToolIds                   types.List    `tfsdk:"tool_ids"`
	FunctionIds               types.List    `tfsdk:"function_ids"`
}

// AssistantVoiceModel describes the voice configuration
type AssistantVoiceModel struct {
	ProviderType    types.String  `tfsdk:"provider_type"`
	VoiceID         types.String  `tfsdk:"voice_id"`
	Speed           types.Float64 `tfsdk:"speed"`
	Stability       types.Float64 `tfsdk:"stability"`
	SimilarityBoost types.Float64 `tfsdk:"similarity_boost"`
	Style           types.Float64 `tfsdk:"style"`
	UseSpeakerBoost types.Bool    `tfsdk:"use_speaker_boost"`
}

func (r *AssistantResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_assistant"
}

func (r *AssistantResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Vapi Assistant resource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Assistant identifier",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Assistant name",
				Required:            true,
			},
			"first_message": schema.StringAttribute{
				MarkdownDescription: "First message the assistant will say",
				Optional:            true,
			},
			"system_message": schema.StringAttribute{
				MarkdownDescription: "System message for the assistant",
				Optional:            true,
			},
			"model": schema.SingleNestedAttribute{
				MarkdownDescription: "Model configuration for the assistant",
				Optional:            true,
				Attributes: map[string]schema.Attribute{
					"provider_type": schema.StringAttribute{
						MarkdownDescription: "Model provider (e.g., openai, anthropic)",
						Required:            true,
					},
					"model": schema.StringAttribute{
						MarkdownDescription: "Model name (e.g., gpt-4, claude-3-sonnet)",
						Required:            true,
					},
					"temperature": schema.Float64Attribute{
						MarkdownDescription: "Temperature for the model",
						Optional:            true,
					},
					"max_tokens": schema.Int64Attribute{
						MarkdownDescription: "Maximum tokens for the model",
						Optional:            true,
					},
					"emotion_recognition_enabled": schema.BoolAttribute{
						MarkdownDescription: "Whether emotion recognition is enabled",
						Optional:            true,
					},
					"num_fast_turns": schema.Int64Attribute{
						MarkdownDescription: "Number of fast turns",
						Optional:            true,
					},
					"tool_ids": schema.ListAttribute{
						MarkdownDescription: "List of tool IDs",
						Optional:            true,
						ElementType:         types.StringType,
					},
					"function_ids": schema.ListAttribute{
						MarkdownDescription: "List of function IDs",
						Optional:            true,
						ElementType:         types.StringType,
					},
				},
			},
			"voice": schema.SingleNestedAttribute{
				MarkdownDescription: "Voice configuration for the assistant",
				Optional:            true,
				Attributes: map[string]schema.Attribute{
					"provider_type": schema.StringAttribute{
						MarkdownDescription: "Voice provider (e.g., elevenlabs, playht)",
						Required:            true,
					},
					"voice_id": schema.StringAttribute{
						MarkdownDescription: "Voice ID",
						Required:            true,
					},
					"speed": schema.Float64Attribute{
						MarkdownDescription: "Voice speed",
						Optional:            true,
					},
					"stability": schema.Float64Attribute{
						MarkdownDescription: "Voice stability",
						Optional:            true,
					},
					"similarity_boost": schema.Float64Attribute{
						MarkdownDescription: "Voice similarity boost",
						Optional:            true,
					},
					"style": schema.Float64Attribute{
						MarkdownDescription: "Voice style",
						Optional:            true,
					},
					"use_speaker_boost": schema.BoolAttribute{
						MarkdownDescription: "Whether to use speaker boost",
						Optional:            true,
					},
				},
			},
			"client_messages": schema.ListAttribute{
				MarkdownDescription: "List of client messages",
				Optional:            true,
				ElementType:         types.StringType,
			},
			"server_messages": schema.ListAttribute{
				MarkdownDescription: "List of server messages",
				Optional:            true,
				ElementType:         types.StringType,
			},
			"silence_timeout_seconds": schema.Int64Attribute{
				MarkdownDescription: "Silence timeout in seconds",
				Optional:            true,
			},
			"max_duration_seconds": schema.Int64Attribute{
				MarkdownDescription: "Maximum duration in seconds",
				Optional:            true,
			},
			"background_sound": schema.StringAttribute{
				MarkdownDescription: "Background sound",
				Optional:            true,
			},
			"background_denoising_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether background denoising is enabled",
				Optional:            true,
			},
			"model_output_in_messages_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether model output in messages is enabled",
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

func (r *AssistantResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AssistantResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AssistantResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Convert Terraform model to API model
	assistant := &client.Assistant{
		Name: data.Name.ValueString(),
	}

	if !data.FirstMessage.IsNull() {
		assistant.FirstMessage = data.FirstMessage.ValueString()
	}

	// Handle system message - it needs to be in the model object
	if !data.SystemMessage.IsNull() {
		if assistant.Model == nil {
			assistant.Model = &client.AssistantModel{
				Provider: "openai",      // Default provider
				Model:    "gpt-4o-mini", // Default model
			}
		}
		assistant.Model.SystemPrompt = data.SystemMessage.ValueString()
	}

	if !data.Model.IsNull() {
		var modelData AssistantModelModel
		resp.Diagnostics.Append(data.Model.As(ctx, &modelData, basetypes.ObjectAsOptions{})...)
		if resp.Diagnostics.HasError() {
			return
		}

		// Preserve any existing model if system message was set
		if assistant.Model == nil {
			assistant.Model = &client.AssistantModel{}
		}
		
		assistant.Model.Provider = modelData.ProviderType.ValueString()
		assistant.Model.Model = modelData.Model.ValueString()

		if !modelData.Temperature.IsNull() {
			temp := modelData.Temperature.ValueFloat64()
			assistant.Model.Temperature = &temp
		}

		if !modelData.MaxTokens.IsNull() {
			maxTokens := int(modelData.MaxTokens.ValueInt64())
			assistant.Model.MaxTokens = &maxTokens
		}

		if !modelData.EmotionRecognitionEnabled.IsNull() {
			emotionRecognition := modelData.EmotionRecognitionEnabled.ValueBool()
			assistant.Model.EmotionRecognitionEnabled = &emotionRecognition
		}

		if !modelData.NumFastTurns.IsNull() {
			numFastTurns := int(modelData.NumFastTurns.ValueInt64())
			assistant.Model.NumFastTurns = &numFastTurns
		}

		if !modelData.ToolIds.IsNull() {
			var toolIds []string
			resp.Diagnostics.Append(modelData.ToolIds.ElementsAs(ctx, &toolIds, false)...)
			if resp.Diagnostics.HasError() {
				return
			}
			assistant.Model.ToolIds = toolIds
		}

		if !modelData.FunctionIds.IsNull() {
			var functionIds []string
			resp.Diagnostics.Append(modelData.FunctionIds.ElementsAs(ctx, &functionIds, false)...)
			if resp.Diagnostics.HasError() {
				return
			}
			assistant.Model.FunctionIds = functionIds
		}
	}

	if !data.Voice.IsNull() {
		var voiceData AssistantVoiceModel
		resp.Diagnostics.Append(data.Voice.As(ctx, &voiceData, basetypes.ObjectAsOptions{})...)
		if resp.Diagnostics.HasError() {
			return
		}

		assistant.Voice = &client.AssistantVoice{
			Provider: voiceData.ProviderType.ValueString(),
			VoiceID:  voiceData.VoiceID.ValueString(),
		}

		if !voiceData.Speed.IsNull() {
			speed := voiceData.Speed.ValueFloat64()
			assistant.Voice.Speed = &speed
		}

		if !voiceData.Stability.IsNull() {
			stability := voiceData.Stability.ValueFloat64()
			assistant.Voice.Stability = &stability
		}

		if !voiceData.SimilarityBoost.IsNull() {
			similarityBoost := voiceData.SimilarityBoost.ValueFloat64()
			assistant.Voice.SimilarityBoost = &similarityBoost
		}

		if !voiceData.Style.IsNull() {
			style := voiceData.Style.ValueFloat64()
			assistant.Voice.Style = &style
		}

		if !voiceData.UseSpeakerBoost.IsNull() {
			useSpeakerBoost := voiceData.UseSpeakerBoost.ValueBool()
			assistant.Voice.UseSpeakerBoost = &useSpeakerBoost
		}
	}

	if !data.ClientMessages.IsNull() {
		var clientMessages []string
		resp.Diagnostics.Append(data.ClientMessages.ElementsAs(ctx, &clientMessages, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		assistant.ClientMessages = clientMessages
	}

	if !data.ServerMessages.IsNull() {
		var serverMessages []string
		resp.Diagnostics.Append(data.ServerMessages.ElementsAs(ctx, &serverMessages, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		assistant.ServerMessages = serverMessages
	}

	if !data.SilenceTimeoutSeconds.IsNull() {
		silenceTimeout := int(data.SilenceTimeoutSeconds.ValueInt64())
		assistant.SilenceTimeoutSeconds = &silenceTimeout
	}

	if !data.MaxDurationSeconds.IsNull() {
		maxDuration := int(data.MaxDurationSeconds.ValueInt64())
		assistant.MaxDurationSeconds = &maxDuration
	}

	if !data.BackgroundSound.IsNull() {
		assistant.BackgroundSound = data.BackgroundSound.ValueString()
	}

	if !data.BackgroundDenoisingEnabled.IsNull() {
		backgroundDenoising := data.BackgroundDenoisingEnabled.ValueBool()
		assistant.BackgroundDenoisingEnabled = &backgroundDenoising
	}

	if !data.ModelOutputInMessagesEnabled.IsNull() {
		modelOutputInMessages := data.ModelOutputInMessagesEnabled.ValueBool()
		assistant.ModelOutputInMessagesEnabled = &modelOutputInMessages
	}

	// Create the assistant
	createdAssistant, err := r.client.CreateAssistant(assistant)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create assistant, got error: %s", err))
		return
	}

	// Update the model with the created assistant data
	data.ID = types.StringValue(createdAssistant.ID)
	
	// For now, set timestamp fields to null since VAPI API may not return them consistently
	// This prevents "unknown value" errors while keeping the fields available
	data.CreatedAt = types.StringNull()
	data.UpdatedAt = types.StringNull()

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssistantResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AssistantResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get the assistant from the API
	assistant, err := r.client.GetAssistant(data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read assistant, got error: %s", err))
		return
	}

	// Update the model with the assistant data
	data.Name = types.StringValue(assistant.Name)
	data.FirstMessage = types.StringValue(assistant.FirstMessage)
	if assistant.Model != nil {
		data.SystemMessage = types.StringValue(assistant.Model.SystemPrompt)
	} else {
		data.SystemMessage = types.StringValue("")
	}
	// For now, set timestamp fields to null since VAPI API may not return them consistently
	// This prevents "unknown value" errors while keeping the fields available
	data.CreatedAt = types.StringNull()
	data.UpdatedAt = types.StringNull()

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssistantResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AssistantResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Convert Terraform model to API model (similar to Create)
	assistant := &client.Assistant{
		Name: data.Name.ValueString(),
	}

	if !data.FirstMessage.IsNull() {
		assistant.FirstMessage = data.FirstMessage.ValueString()
	}

	// Handle system message - it needs to be in the model object
	if !data.SystemMessage.IsNull() {
		if assistant.Model == nil {
			assistant.Model = &client.AssistantModel{
				Provider: "openai",      // Default provider
				Model:    "gpt-4o-mini", // Default model
			}
		}
		assistant.Model.SystemPrompt = data.SystemMessage.ValueString()
	}

	// Update the assistant
	updatedAssistant, err := r.client.UpdateAssistant(data.ID.ValueString(), assistant)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update assistant, got error: %s", err))
		return
	}

	// Update the model with the updated assistant data
	data.UpdatedAt = types.StringValue(updatedAssistant.UpdatedAt)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssistantResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AssistantResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete the assistant
	err := r.client.DeleteAssistant(data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete assistant, got error: %s", err))
		return
	}
}

func (r *AssistantResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
