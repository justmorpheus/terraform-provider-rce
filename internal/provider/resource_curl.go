package provider

import (
    "context"
    "io/ioutil"
    "net/http"

    "github.com/hashicorp/terraform-plugin-framework/resource"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema"
    "github.com/hashicorp/terraform-plugin-framework/types"
)

type curlResource struct{}

func NewCurlResource() resource.Resource {
    return &curlResource{}
}

func (r *curlResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = "custom_curl"
}

func (r *curlResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
            "url": schema.StringAttribute{
                Description: "The URL to send the GET request to.",
                Required:    true,
            },
            "response": schema.StringAttribute{
                Description: "The response body from the GET request.",
                Computed:    true,
            },
        },
    }
}

type CurlResourceModel struct {
    URL      types.String `tfsdk:"url"`
    Response types.String `tfsdk:"response"`
}

func (r *curlResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
    var plan CurlResourceModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    response, err := http.Get(plan.URL.ValueString())
    if err != nil {
        resp.Diagnostics.AddError("Error making GET request", err.Error())
        return
    }
    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        resp.Diagnostics.AddError("Error reading response body", err.Error())
        return
    }

    var state CurlResourceModel
    state.URL = plan.URL
    state.Response = types.StringValue(string(body))

    diags = resp.State.Set(ctx, state)
    resp.Diagnostics.Append(diags...)
}

func (r *curlResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
    // No-op
}

func (r *curlResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
    // No-op
}

func (r *curlResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
    // No-op
}
