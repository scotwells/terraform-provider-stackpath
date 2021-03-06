package stackpath

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns the configured provider for managing StackPath resources.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"client_id": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Optional:  true,
			},
			"client_secret": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Optional:  true,
			},
			"access_token": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Optional:  true,
			},
			"stack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"base_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				// Default to the official StackPath API
				Default: defaultBaseURL,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"stackpath_compute_workload":       resourceComputeWorkload(),
			"stackpath_compute_network_policy": resourceComputeNetworkPolicy(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(data *schema.ResourceData) (interface{}, error) {
	config := &Config{
		Stack:   data.Get("stack").(string),
		BaseURL: data.Get("base_url").(string),
	}

	if v, ok := data.GetOk("access_token"); ok {
		config.AccessToken = v.(string)
	}
	if v, ok := data.GetOk("client_id"); ok {
		config.ClientID = v.(string)
	}
	if v, ok := data.GetOk("client_secret"); ok {
		config.ClientSecret = v.(string)
	}

	log.Printf("[INFO] configuring stackpath provider")
	if err := config.LoadAndValidate(); err != nil {
		return nil, fmt.Errorf("unable to validate configuration: %v", err)
	}

	return config, nil
}
