# Tanzu Application Platform Initialize Accelerator

The following files are included:
- `config/workload.yaml`: The definition of a Workload for a TAP Supply Chain
- `catalog/catalog-info.yaml`: The definition of a Component to make the Workload available in the TAP GUI Catalog
- `INSTRUCTIONS.md` : This file, providing instructions for preparing your project
- `accelerator-log.md` : This is the log file from the accelerator engine processing
- `Tiltfile` : A Tiltfile to enable Live Update and Debugging for Java Workloads using the TAP IDE Plugin (only provided if you selected to "Include TAP IDE Support for Java Workloads")

## Preparing your goapi project

Copy the provided `config` and `catalog` directories including the files they contain into your `goapi` project. Copy the `DEPLOYING.md` file to the root directory of `goapi` project. If a `Tiltfile` was included then copy it to the root directory of `goapi` project.

You can now follow the instructions in the `DEPLOYING.md` file to deploy your project.

