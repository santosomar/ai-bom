package serializer

import (
	"fmt"
	"strings"
	"time"

	cdx "github.com/CycloneDX/cyclonedx-go"
	"github.com/google/uuid"

	"github.com/manifest-cyber/ai-bom/pkg/huggingface"
)

func ConvertHuggingfaceModelToBom(name, version string, modelInfo *huggingface.ModelInfo) (*cdx.BOM, error) {
	var comp *cdx.Component
	if name != "" || version != "" {
		comp = &cdx.Component{
			Type:    cdx.ComponentTypeMachineLearningModel,
			Name:    name,
			Version: version,
		}
	}
	bom := &cdx.BOM{
		BOMFormat:    cdx.BOMFormat,
		SpecVersion:  cdx.SpecVersion1_5,
		Version:      1,
		SerialNumber: fmt.Sprintf("urn:uuid:%s", uuid.New().String()),
		Metadata: &cdx.Metadata{
			Timestamp: time.Now().Format(time.RFC3339),
			Tools: &[]cdx.Tool{
				// TODO: consider moving to CLI level
				{
					Vendor:  "Manifest Cyber",
					Name:    "ai-bom",
					Version: "0.0.1",
				},
			},
			Component: comp,
		},
		Components: &[]cdx.Component{
			*ConvertToProtoComponent(modelInfo),
		},
	}

	return bom, nil
}

func ConvertToProtoComponent(modelInfo *huggingface.ModelInfo) *cdx.Component {
	var licenses cdx.Licenses
	licenses = append(licenses, cdx.LicenseChoice{
		License: &cdx.License{
			ID: modelInfo.CardData.License,
		},
	})

	component := &cdx.Component{
		BOMRef: strings.ToLower(fmt.Sprintf("%s-%s", strings.ReplaceAll(modelInfo.ModelID, "/", "-"), modelInfo.ID)),
		Supplier: &cdx.OrganizationalEntity{
			Name: modelInfo.Author,
		},
		Name:     modelInfo.ModelID,
		Version:  modelInfo.SHA,
		Licenses: &licenses,
		ModelCard: &cdx.MLModelCard{
			ModelParameters: &cdx.MLModelParameters{
				Task:               modelInfo.PipelineTag,
				ArchitectureFamily: modelInfo.Config.ModelType,
				ModelArchitecture:  modelInfo.Config.Architectures[0],
				Datasets:           ConvertDatasets(&modelInfo.CardData.Datasets),
			},
			QuantitativeAnalysis: &cdx.MLQuantitativeAnalysis{
				PerformanceMetrics: ConvertMetrics(modelInfo.ModelIndex),
			},
			Considerations: &cdx.MLModelCardConsiderations{},
		},
	}

	return component
}

func ConvertDatasets(datasets *[]string) *[]cdx.MLDatasetChoice {
	var protoDatasets *[]cdx.MLDatasetChoice

	if datasets == nil || len(*datasets) == 0 {
		return nil
	}

	protoDatasets = &[]cdx.MLDatasetChoice{}

	for _, dataset := range *datasets {
		split := strings.Split(dataset, "/")

		var namespace string
		name := split[0]
		classification := "private"
		var contents *cdx.ComponentDataContents
		if len(split) == 2 {
			namespace = split[0]
			name = split[1]
			classification = "public"
			contents = &cdx.ComponentDataContents{
				URL: fmt.Sprintf("%s/datasets/%s/%s", huggingface.HuggingfaceURL, namespace, name),
			}
		}
		protoDataset := cdx.MLDatasetChoice{
			ComponentData: &cdx.ComponentData{
				Name:           name,
				Type:           "dataset",
				Classification: classification,
				Contents:       contents,
			},
		}
		*protoDatasets = append(*protoDatasets, protoDataset)
	}

	return protoDatasets
}

func ConvertMetrics(metrics *[]huggingface.ModelIndex) *[]cdx.MLPerformanceMetric {
	var protoMetrics []cdx.MLPerformanceMetric

	if metrics == nil {
		return &protoMetrics
	}

	for _, metric := range *metrics {
		for _, result := range metric.Results {
			for _, m := range result.Metrics {
				protoMetric := cdx.MLPerformanceMetric{
					Type:  m.Type,
					Value: fmt.Sprintf("%f", m.Value),
				}
				protoMetrics = append(protoMetrics, protoMetric)
			}
		}
	}

	return &protoMetrics
}
