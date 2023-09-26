
# AI BOM (Bill of Materials) JSON Schema

## Overview

This [JSON schema](schema.md) provides a structured format for describing the aspects of an Artificial Intelligence (AI) / Machine Learning (ML) model. This schema aims to help both model developers and users to understand and communicate essential details about a model including its architecture, usage guidelines, and other metadata.

It can be used as a reference for creating a model card or a software bill of materials (SBOM) for an AI/ML model. It can also be used to validate an existing model card or SBOM. This can be integrated to CycloneDX, SPDX SBOM formats or any other lightweight implementation.

## Sections
The following are the schema sections based on the existing documentation in the [../README.md] file of this repository.

### Model Details

- **Name**: The name of the model (Required)
- **Version**: The version of the model (Required)
- **Type**: The type of model, e.g., "text-generation," "image-processing" (Required)
- **Author**: The individual or organization that developed the model (Required)
- **Licenses**: List of software licenses for this model (Required)
- **Libraries**: Libraries that the model is dependent on (Optional)
- **Source (URL)**: The source URL of the model (Required)
- **BOM Generation**: Metadata about how the BOM was generated (Optional)
- **Other References**: Links to additional resources like papers, contact emails, etc. (Optional)
- **Tags**: Labels or tags associated with the model (Optional)

### Model Architecture

- **Datasets**: Datasets used to train the model (Required)
- **Architecture**: The model's architecture, e.g., "GPT-J" (Optional)
- **Architecture Family**: The family of the model's architecture (Optional)
- **Parent Model**: Information about the parent model if applicable (Optional)
- **Base Model**: Information about the base model if applicable (Optional)
- **Input**: The type of input the model accepts (Required)
- **Output**: The type of output the model generates (Required)
- **Hardware**: Information about the hardware used (Optional)
- **Software**: Information about the software used (Optional)
- **Software Required for Execution**: Indicates whether software files are part of the core files (Required)

### Usage

- **Intended Use**: A description of the intended use of the model (Required)
- **Out of Scope Usage**: Descriptions of what the model is not intended for (Required)
- **Misuse or Malicious Use**: What constitutes misuse or malicious use of the model (Required)

### Considerations

- **Environmental Impact**: Information about the model's environmental impact (Optional)
- **Ethical Considerations**: Ethical considerations and biases (Optional)

### Attestations

- **Attestation**: A digital signature for the authenticity and integrity of the BOM (Optional)



