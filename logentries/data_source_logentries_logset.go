package logentries

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceLogentriesLogSet() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLogentriesLogSetRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceLogentriesLogSetRead(d *schema.ResourceData, meta interface{}) error {
	return resourceLogentriesLogSetRead(d, meta)
}
