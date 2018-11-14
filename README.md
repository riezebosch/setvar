# setvar

[![Build Status](https://manuel.visualstudio.com/setvar/_apis/build/status/riezebosch.setvar)](https://manuel.visualstudio.com/setvar/_build/latest?definitionId=23)

Transform output from a tool into [pipeline variables](https://docs.microsoft.com/en-us/azure/devops/pipelines/process/variables?view=vsts&tabs=yaml%2Cbatch#set-a-job-scoped-variable-from-a-script) for your Azure Pipelines.

Instead of 
```yaml
  - powershell: echo "##vso[task.setvariable variable=myOutputVar;isOutput=true]this is the value"
```

Just do
```yaml
  - powershell: echo "this is the value" | setvar -name myOutputVar -isOutput
```

This mostly makes sense when the value is the output of another tool.
```yaml
- task: AzureCLI@1
  inputs:
    scriptLocation: inlineScript
    inlineScript: |+
      az image list '[0].name' -o tsv | setvar -name imageName -isOutput=true
```

Because now you don't have to worry anymore about how to capture output into a variable in a specific shell, thereby making your script agent specific.
