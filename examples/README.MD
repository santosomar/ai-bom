# Examples Directory

## Contents
This directory contains several examples of what an AI-BOM looks like, including two base examples and several examples of from actual models. 

**Templates**  
- `template-ai-bom-cdx.json`: This is a sample in valid CycloneDX format, with some extra data added beyond CycloneDX's [example](https://github.com/CycloneDX/specification/blob/master/tools/src/test/resources/1.5/valid-machine-learning-1.5.json) on their Github.  

-  `template-ai-bom.json`: This is a prososed model,based on CycloneDX, that has additional fields that 
CycloneDX doesn't yet account for (for which we've already filed PRs), and has a cleaner modeling for several other fields based on what we observed acrossing Google's, Meta's, and HuggingFace's model cards. We are calling it "AI Data Exchange" (AIDX) for the sake of these examples, but it is our sincere hope *not* to create a new format for AI-BOMs, but rather that CycloneDX and SPDX incorporate our suggestions into their respective formats.


**Examaples**
- `RedPajama-INCITE-7B-Base`: A popular text generation ML model listed on [HuggingFace](https://huggingface.co/togethercomputer/RedPajama-INCITE-7B-Base). In the temporary 'AIDX' format.  

- `faclon-7b-instruct`: A popular text generation ML model listed on [HuggingFace](https://huggingface.co/tiiuae/falcon-7b-instruct). In the temporary 'AIDX' format.  
