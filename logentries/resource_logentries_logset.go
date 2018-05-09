package logentries

import (
	"log"
	"strings"

	logentries "github.com/depop/logentries"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceLogentriesLogSet() *schema.Resource {

	return &schema.Resource{
		Create: resourceLogentriesLogSetCreate,
		Read:   resourceLogentriesLogSetRead,
		Update: resourceLogentriesLogSetUpdate,
		Delete: resourceLogentriesLogSetDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceLogentriesLogSetCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*logentries.Client)
	res, err := client.LogSet.Create(&logentries.LogSetCreateRequest{
		LogSet: logentries.LogSetFields{
			Name: d.Get("name").(string),
		},
	})

	if err != nil {
		return err
	}

	d.SetId(res.ID)

	return resourceLogentriesLogSetRead(d, meta)
}

func resourceLogentriesLogSetRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*logentries.Client)
	res, err := client.LogSet.Read(&logentries.LogSetReadRequest{
		ID: d.Id(),
	})
	if err != nil {
		if strings.Contains(err.Error(), "No such log set") {
			log.Printf("Logentries LogSet Not Found - Refreshing from State")
			d.SetId("")
			return nil
		}
		return err
	}

	if res == nil {
		d.SetId("")
		return nil
	}

	d.SetId(res.ID)

	return nil
}

func resourceLogentriesLogSetUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*logentries.Client)
	_, err := client.LogSet.Update(&logentries.LogSetUpdateRequest{
		ID: d.Id(),
		LogSet: logentries.LogSetFields{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
		},
	})
	if err != nil {
		return err
	}

	return resourceLogentriesLogRead(d, meta)
}

func resourceLogentriesLogSetDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*logentries.Client)
	_, err := client.LogSet.Delete(&logentries.LogSetDeleteRequest{
		ID: d.Id(),
	})
	return err
}
