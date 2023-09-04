package options

import (
	"github.com/spf13/cobra"
)

// BomOptions defines the options for the `bom` command
type BomOptions struct {
	HuggingFaceAPIKey string
	Name              string
	Version           string
	Format            string
	Output            string
	OpenAPIKey        string
}

// AddFlags adds command line flags for the BomOptions struct
func (o *BomOptions) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&o.Name, "name", "", "", "name of generated SBOM document")
	cmd.Flags().StringVarP(&o.Version, "version", "", "", "version of generated SBOM document")
	cmd.Flags().StringVarP(&o.Output, "output", "o", "", "path to write the BOM")
	cmd.Flags().StringVarP(&o.Format, "format", "f", "cyclonedx-json", "the output format [cyclonedx-json]")
	cmd.Flags().StringVar(&o.HuggingFaceAPIKey, "hf-key", "", "HuggingFace API key ModelCard source")
	cmd.Flags().StringVar(&o.OpenAPIKey, "openapi-key", "", "OpenAPI key for the ModelCard parsing service")
}
