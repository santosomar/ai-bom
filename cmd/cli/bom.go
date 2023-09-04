package cli

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"

	"github.com/manifest-cyber/ai-bom/cmd/cli/options"
	"github.com/manifest-cyber/ai-bom/pkg/domain"
	"github.com/manifest-cyber/ai-bom/pkg/huggingface"
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

	file, err := client.StreamFile(ctx, repoID, "README.md", "", "model", revision)
	if err != nil {
		return err
	}
	defer file.Close()

	// Parse the file contents
	contents, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// Print the file contents to stdout
	//nolint:forbidigo
	fmt.Println(string(contents))
	return nil
}
