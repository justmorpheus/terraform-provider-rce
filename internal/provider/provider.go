package provider

import (
"context"

"github.com/hashicorp/terraform-plugin-framework/provider"
"github.com/hashicorp/terraform-plugin-framework/provider/schema"
"github.com/hashicorp/terraform-plugin-framework/resource"
"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type customProvider struct{}

func New() provider.Provider {
return &customProvider{}
}

func (p *customProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
resp.TypeName = "custom"
}

func (p *customProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
resp.Schema = schema.Schema{
Attributes: map[string]schema.Attribute{
"example": schema.StringAttribute{
Description: "An example attribute.",
Optional:    true,
},
},
}
}

func (p *customProvider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
// No-op
}

func (p *customProvider) Resources(_ context.Context) []func() resource.Resource {
return []func() resource.Resource{
NewCurlResource,
}
}

func (p *customProvider) DataSources(_ context.Context) []func() datasource.DataSource {
// Returning nil as no data sources are defined yet
return nil
}
