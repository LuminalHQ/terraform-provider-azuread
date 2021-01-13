package common

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/sender"
	"github.com/hashicorp/terraform-plugin-sdk/v2/meta"
	"github.com/manicminer/hamilton/auth"
	"github.com/manicminer/hamilton/base"

	"github.com/terraform-providers/terraform-provider-azuread/version"
)

type ClientOptions struct {
	TenantID    string
	Environment azure.Environment

	PartnerID        string
	TerraformVersion string

	AadGraphAuthorizer autorest.Authorizer
	AadGraphEndpoint   string

	MsGraphAuthorizer auth.Authorizer

	SkipProviderReg bool
}

func (o ClientOptions) ConfigureClient(c *base.Client, ar *autorest.Client) {
	c.Authorizer = o.MsGraphAuthorizer
	c.UserAgent = o.userAgent(c.UserAgent)

	ar.Authorizer = o.AadGraphAuthorizer
	ar.Sender = sender.BuildSender("AzureAD")
	ar.UserAgent = o.userAgent(ar.UserAgent)
}

func (o ClientOptions) userAgent(sdkUserAgent string) (userAgent string) {
	tfUserAgent := fmt.Sprintf("HashiCorp Terraform/%s (+https://www.terraform.io) Terraform Plugin SDK/%s", o.TerraformVersion, meta.SDKVersionString())
	providerUserAgent := fmt.Sprintf("%s terraform-provider-azuread/%s", tfUserAgent, version.ProviderVersion)
	userAgent = strings.TrimSpace(fmt.Sprintf("%s %s", sdkUserAgent, providerUserAgent))

	// append the CloudShell version to the user agent if it exists
	if azureAgent := os.Getenv("AZURE_HTTP_USER_AGENT"); azureAgent != "" {
		userAgent = fmt.Sprintf("%s %s", userAgent, azureAgent)
	}

	if o.PartnerID != "" {
		userAgent = fmt.Sprintf("%s pid-%s", userAgent, o.PartnerID)
	}

	log.Printf("[DEBUG] User Agent: %s\n", userAgent)
	return
}