package cli

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	cdx "github.com/CycloneDX/cyclonedx-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"

	"github.com/manifest-cyber/ai-bom/cmd/cli/options"
	"github.com/manifest-cyber/ai-bom/pkg/domain"
	"github.com/manifest-cyber/ai-bom/pkg/huggingface"
	"github.com/manifest-cyber/ai-bom/pkg/serializer"
)

func BomCommand() *cobra.Command {
	co := &options.BomOptions{}
	c := &cobra.Command{
		Use:        "bom",
		Aliases:    []string{"b"},
		SuggestFor: []string{"bom"},
		Short:      "Generate a BOM from a HuggingFace ModelCard",
		Long:       `Generate a BOM from a HuggingFace ModelCard"`,
		Example: `
ai-bom bom meta-llama/Llama-2-7b        				output to stdout
ai-bom bom meta-llama/Llama-2-7b -o model.json        			output to a file
	 				`,
		SilenceErrors:     true,
		SilenceUsage:      true,
		DisableAutoGenTag: true,
		Args:              cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runBom(cmd.Context(), co, args)
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Parent().PersistentPreRunE(cmd.Parent(), args); err != nil {
				return err
			}

			if err := options.BindConfig(viper.GetViper(), cmd); err != nil {
				return err
			}

			return validateBomOptions(co, args)
		},
	}

	co.AddFlags(c)

	return c
}

func validateBomOptions(opt *options.BomOptions, args []string) error {
	if opt.Name == "" {
		opt.Name = strings.ReplaceAll(args[0], "/", "_")
	}

	if !slices.Contains(domain.BOMOutputs, opt.Format) {
		return fmt.Errorf("output format %s is not supported, must be one of: %v", opt.Format, domain.BOMOutputs)
	}

	return nil
}

// runBom is the main entrypoint for the `bom` command
func runBom(ctx context.Context, opt *options.BomOptions, args []string) error {
	var revision, repoID string
	parts := strings.Split(args[0], ":")
	repoID = parts[0]
	if len(parts) == 2 {
		revision = parts[1]
	}

	client := huggingface.NewHFHubClient(
		huggingface.WithToken(opt.HuggingFaceAPIKey),
	)

	apiClient := huggingface.NewHFAPIClient(
		huggingface.WithAPIToken(opt.HuggingFaceAPIKey),
	)

	// openaiClient := openai.NewCompletionsClient(
	// 	openai.WithToken(opt.OpenAIAPIKey),
	// )

	file, err := client.StreamFile(ctx, repoID, "README.md", "", "model", revision)
	if err != nil {
		return err
	}
	defer file.Close()

	// Parse the file contents
	// contents, err := io.ReadAll(file)
	// if err != nil {
	// 	return err
	// }

	modelInfo, err := apiClient.GetModelInfo(ctx, &huggingface.GetModelInfoOptions{
		RepoID:   repoID,
		Revision: revision,
	})
	if err != nil {
		return err
	}

	comp, err := serializer.ConvertHuggingfaceModelToBom(opt.Name, opt.Version, modelInfo)
	if err != nil {
		return err
	}

	var wr io.Writer
	if opt.Output == "" {
		wr = os.Stdout
	} else {
		file, err := os.Create(opt.Output)
		if err != nil {
			return err
		}

		defer file.Close()
		wr = file
	}

	encoder := cdx.NewBOMEncoder(wr, cdx.BOMFileFormatJSON)
	encoder.SetPretty(true)

	if err := encoder.EncodeVersion(comp, comp.SpecVersion); err != nil {
		return fmt.Errorf("encoding sbom to stream: %w", err)
	}

	// res, err := openaiClient.Completions(ctx, fmt.Sprintf("build a json from this README.md file content: %s", string(contents)))
	// if err != nil {
	// 	return err
	// }
	// //nolint:forbidigo
	// fmt.Println(modelInfo)

	return nil
}
