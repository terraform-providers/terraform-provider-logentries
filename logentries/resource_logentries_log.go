package logentries

import (
	"fmt"
	"log"
	"strings"

	logentries "github.com/depop/logentries"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceLogentriesLog() *schema.Resource {

	return &schema.Resource{
		Create: resourceLogentriesLogCreate,
		Read:   resourceLogentriesLogRead,
		Update: resourceLogentriesLogUpdate,
		Delete: resourceLogentriesLogDelete,

		Schema: map[string]*schema.Schema{
			"token": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"logset_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"filename": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "token",
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(string)
					allowed_values := []string{"token", "syslog", "agent", "api"}
					if !sliceContains(value, allowed_values) {
						errors = append(errors, fmt.Errorf("Invalid log source option: %s (must be one of: %s)", value, allowed_values))
					}
					return
				},
			},
			"type": {
				Type:     schema.TypeString,
				Default:  "",
				Optional: true,
			},
		},
	}
}

func resourceLogentriesLogCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*logentries.Client)

	res, err := client.Log.Create(&logentries.LogCreateRequest{
		logentries.LogCreateRequestFields{
			Name:       d.Get("name").(string),
			SourceType: d.Get("source").(string),
			UserData: logentries.LogUserData{
				LeAgentFilename: d.Get("filename").(string),
			},
			LogsetsInfo: []logentries.LogsetsInfo{
				logentries.LogsetsInfo{
					ID: d.Get("logset_id").(string),
				},
			},
		},
	})

	if err != nil {
		return err
	}

	if d.Get("source").(string) == "token" {
		d.Set("token", res.Tokens[0])
	}

	d.SetId(res.ID)

	return resourceLogentriesLogRead(d, meta)
}

func resourceLogentriesLogRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*logentries.Client)
	res, err := client.Log.Read(&logentries.LogReadRequest{
		ID: d.Id(),
	})
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Printf("Logentries Log Not Found - Refreshing from State")
			d.SetId("")
			return nil
		}
		return err
	}

	if res == nil {
		d.SetId("")
		return nil
	}

	return nil
}

func resourceLogentriesLogUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*logentries.Client)
	_, err := client.Log.Update(&logentries.LogUpdateRequest{
		ID: d.Id(),
		Log: logentries.LogUpdateRequestFields{
			UserData: logentries.LogUserData{
				LeAgentFilename: d.Get("filename").(string),
			},
			LogsetsInfo: []logentries.LogsetsInfo{
				logentries.LogsetsInfo{
					Name: d.Get("name").(string),
				},
			},
		},
	})
	if err != nil {
		return err
	}

	return resourceLogentriesLogRead(d, meta)
}

func resourceLogentriesLogDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*logentries.Client)
	_, err := client.Log.Delete(&logentries.LogDeleteRequest{
		ID: d.Id(),
	})
	return err
}

func sliceContains(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
