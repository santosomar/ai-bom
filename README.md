# ai-bom

![robot (1)](https://github.com/manifest-cyber/ai-bom/assets/862262/0ec82e8b-fdc1-47b9-b9b0-55d9119657e1)

## Proposed AI-BOM Model

We analyzed the leading SBOM formats and various model card formats, and conducted extensive research with AI/ML experts and developers. Below is our initial proposed AI-BOM model. No existing SBOM (CycloneDX or SPDX) or model card format perfectly matches the below content, so we acknowledge there is additional work needed with those communities to align & consolidate models.  


### Model Details

**Name**  [Required]  
The name of the model. 

**Version** [Required]  
The version of the model. 

**Type** [Required]  
The type of the model. Samples include: "text-generation," "image-processing."

**Author** [Required]  
The individual or organization that developed the model. Often referred as "Developed By" in HuggingFace and other model cards.

**Licenses** [Required]  
The list of software licenses for this model (e.g. Apache 2.0, GPL 3.0, etc.)

**Libraries** [Required?]  
Any libraries that the model is dependent on. For example, in HuggingFace, models may list dependencies on libraries like PyTorch, Transformers, ONNX, Diffusers, etc. 
<Insert picture>

**Source (URL)** [Required]  
A URL or other path to where this model lives, such as on a public repository. 

**BOM Generation** [Required]  
Information about how and when the BOM was generated, and who generated it: a timestamp, a method/tool, and a person/organization. 

**Other references** [Optional]  
Links to other resources, such as papers, contact email addresses, websites, etc. 

**Tags** [Optional]  
Tags or other labels associated with a model, which can often be scraped from HuggingFace or other repositories.  



### Model Architecture

**Datasets** [Required]  
The list of datasets that were used to train the model. This should include, at a minimum, the **name** and **source (URL)** of each dataset, as well as how it can be used (e.g. license information, whether it's public or commercial, etc.). Optional information could also include the procedures used to train using each dataset. 

**Parent Model** [Optional]  
Information about the model's parent model. If present, this would include the model's `name`, `version`, and `source` (e.g. URL). 

**Base Model** [Optional]  
Information about the model's base (a.k.a "foundation") model. If present, this would include the model's `name`, `version`, and `source` (e.g. URL). 

**Input** [Required]  
Information about the model's parent model. If present, this would 

**Output** [Required]  
Information about the model's parent model. If present, this would 

**Hardware** [Optional]  
Information about the model's parent model. If present, this would 

**Software** [Optional]  
Information about the model's parent model. If present, this would 

**Software included** [Required]  
A boolean (True / False) that indicates whether a model includes software files (e.g. python, go, etc.) as part of the core files. See this how this [model](https://huggingface.co/tiiuae/falcon-7b-instruct/tree/main) includes python scripts as part of its listed files. 


### Usage  

**Intended Use** [Required]  
A description of how the model should be used.

**Out of Scope Usage** [Required]  
A description of what use cases or usage of the model is not in scope for the model, and how it was developed. 

**Misuse or Malicious Use** [Required]  
A description of what usage of the model constitutes misuse or malicious use. 

### Considerations  

**Environmental Impact** [Optional]  
TODO

**Ethical Considerations** [Optional]  
TODO

**Environmental Impact** [Optional]  
TODO
