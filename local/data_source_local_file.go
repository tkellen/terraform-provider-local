package local

import (
	"io/ioutil"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceLocalFile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLocalFileRead,

		Schema: map[string]*schema.Schema{
			"filename": {
				Type:        schema.TypeString,
				Description: "Path to the output file",
				Required:    true,
				ForceNew:    true,
			},
			"content": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceLocalFileRead(d *schema.ResourceData, _ interface{}) error {
	path := d.Get("filename").(string)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	d.Set("content", string(content))

	d.SetId("-")

	return nil
}
