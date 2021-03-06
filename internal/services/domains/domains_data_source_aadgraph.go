package domains

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
)

func domainsDataSourceReadAadGraph(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tenantId := meta.(*clients.Client).TenantID
	client := meta.(*clients.Client).Domains.AadClient

	includeUnverified := d.Get("include_unverified").(bool)
	onlyDefault := d.Get("only_default").(bool)
	onlyInitial := d.Get("only_initial").(bool)

	results, err := client.List(ctx, "")
	if err != nil {
		return tf.ErrorDiagF(err, "Listing domains")
	}

	d.SetId("domains-" + tenantId) // todo this should be more unique

	domains := flattenDomainsAad(results.Value, includeUnverified, onlyDefault, onlyInitial)
	if len(domains) == 0 {
		return tf.ErrorDiagF(nil, "No domains were returned for the provided filters")
	}

	tf.Set(d, "domains", domains)

	return nil
}

func flattenDomainsAad(input *[]graphrbac.Domain, includeUnverified, onlyDefault, onlyInitial bool) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	domains := make([]interface{}, 0)
	for _, v := range *input {
		if v.Name == nil {
			log.Printf("[DEBUG] Domain Name was nil - skipping")
			continue
		}

		domainName := *v.Name

		authenticationType := "undefined"
		if v.AuthenticationType != nil {
			authenticationType = *v.AuthenticationType
		}

		isDefault := false
		if v.IsDefault != nil {
			isDefault = *v.IsDefault
		}

		isInitial := false
		if v.AdditionalProperties["isInitial"] != nil {
			isInitial = v.AdditionalProperties["isInitial"].(bool)
		}

		isVerified := false
		if v.IsVerified != nil {
			isVerified = *v.IsVerified
		}

		// Filters
		if !isDefault && onlyDefault {
			// skip all domains except the initial domain
			log.Printf("[DEBUG] Skipping %q since the filter requires the default domain", domainName)
			continue
		}

		if !isInitial && onlyInitial {
			// skip all domains except the initial domain
			log.Printf("[DEBUG] Skipping %q since the filter requires the initial domain", domainName)
			continue
		}

		if !isVerified && !includeUnverified {
			//skip unverified domains
			log.Printf("[DEBUG] Skipping %q since the filter requires verified domains", domainName)
			continue
		}

		domain := map[string]interface{}{
			"authentication_type": authenticationType,
			"domain_name":         domainName,
			"is_default":          isDefault,
			"is_initial":          isInitial,
			"is_verified":         isVerified,
		}

		domains = append(domains, domain)
	}

	return domains
}
