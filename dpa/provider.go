package dpa

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	dpa "github.com/strick-j/cybr-dpa/pkg/dpa"
)

// Provider implements Dpa as a schema.Provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DPA_HOST", nil),
				Description: "The host of the DPA server.",
			},
			"client_credentials": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DPA_CLIENT_CREDENTIALS", nil),
				Description: "Whether to use client credentials for authentication to the DPA Server.",
			},
			"client_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DPA_CLIENT_URL", nil),
				Description: "The URL of the DPA Authorization Server.",
			},
			"client_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DPA_CLIENT_ID", nil),
				Description: "The client ID for authentication to the DPA Server.",
			},
			"client_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DPA_CLIENT_SECRET", nil),
				Description: "The client secret for authentication to the DPA Server.",
			},
			"application_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DPA_APPLICATION_ID", nil),
				Description: "The application ID for authentication to the DPA Server.",
			},
			"scope": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DPA_SCOPE", nil),
				Description: "The scope(s) for authentication to the DPA Server. This is a list of strings.",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DPA_USERNAME", nil),
				Description: "The username of the DPA server.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DPA_PASSWORD", nil),
				Description: "The password of the DPA server.",
			},
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DPA_TOKEN", nil),
				Description: "The Bearer Token for authentication to the DPA Server.",
			},
		},
		ConfigureFunc: providerConfig,
	}
}

func providerConfig(d *schema.ResourceData) (interface{}, error) {
	if d.Get("client_credentials").(bool) {
		clientID := d.Get("client_id").(string)
		clientSecret := d.Get("client_secret").(string)
		applicationID := d.Get("application_id").(string)
		clientURL := d.Get("client_url").(string)
		scope := d.Get("scope").([]string)

		client, err := dpa.OauthCredClient(clientID, clientSecret, applicationID, clientURL, scope)
		if err != nil {
			return nil, err
		}

		return client, nil
	}

	return nil, nil
}
