# Accelerator Log

## Options
```json
{
  "applicationName" : "weatherforecast-csharp",
  "projectName" : "weatherforecast-csharp"
}
```
## Log
```
┏ engine (Chain)
┃  Info Running Chain(GeneratorValidationTransform, UniquePath)
┃ ┏ ┏ engine.transformations[0].validated (Combo)
┃ ┃ ┃  Info Combo running as Let
┃ ┃ ┃ engine.transformations[0].validated.delegate (Let)
┃ ┃ ┃ Debug Adding symbol compilerSafeAppName with value 'weatherforecast_csharp'
┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.in (Chain)
┃ ┃ ┃ ┃  Info Running Chain(Merge, UniquePath)
┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.in.transformations[0] (Merge)
┃ ┃ ┃ ┃ ┃  Info Running Merge(Combo, Combo, Combo, Combo, Provenance)
┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.in.transformations[0].sources[0] (Combo)
┃ ┃ ┃ ┃ ┃ ┃  Info Combo running as Chain
┃ ┃ ┃ ┃ ┃ ┃ engine.transformations[0].validated.delegate.in.transformations[0].sources[0].delegate (Chain)
┃ ┃ ┃ ┃ ┃ ┃  Info Running Chain(Include, Exclude, ReplaceText)
┃ ┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.in.transformations[0].sources[0].delegate.transformations[0] (Include)
┃ ┃ ┃ ┃ ┃ ┃ ┃  Info Will include [**]
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug .gitignore matched [**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug .tanzuignore matched [**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug Controllers/WeatherForecastController.cs matched [**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug LICENSE matched [**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug Program.cs matched [**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug Properties/launchSettings.json matched [**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug README.md matched [**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug Sample.csproj matched [**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug StaticFiles/index.html matched [**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug WeatherForecast.cs matched [**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug appsettings.Development.json matched [**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug appsettings.json matched [**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug catalog/catalog-info.yaml matched [**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug config/workload.yaml matched [**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┗ Debug global.json matched [**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.in.transformations[0].sources[0].delegate.transformations[1] (Exclude)
┃ ┃ ┃ ┃ ┃ ┃ ┃  Info Will exclude [README.md, Sample.csproj, config/**, catalog/**]
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug .gitignore didn't match [README.md, Sample.csproj, config/**, catalog/**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug .tanzuignore didn't match [README.md, Sample.csproj, config/**, catalog/**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug Controllers/WeatherForecastController.cs didn't match [README.md, Sample.csproj, config/**, catalog/**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug LICENSE didn't match [README.md, Sample.csproj, config/**, catalog/**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug Program.cs didn't match [README.md, Sample.csproj, config/**, catalog/**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug Properties/launchSettings.json didn't match [README.md, Sample.csproj, config/**, catalog/**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug README.md matched [README.md, Sample.csproj, config/**, catalog/**] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug Sample.csproj matched [README.md, Sample.csproj, config/**, catalog/**] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug StaticFiles/index.html didn't match [README.md, Sample.csproj, config/**, catalog/**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug WeatherForecast.cs didn't match [README.md, Sample.csproj, config/**, catalog/**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug appsettings.Development.json didn't match [README.md, Sample.csproj, config/**, catalog/**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug appsettings.json didn't match [README.md, Sample.csproj, config/**, catalog/**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug catalog/catalog-info.yaml matched [README.md, Sample.csproj, config/**, catalog/**] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug config/workload.yaml matched [README.md, Sample.csproj, config/**, catalog/**] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┗ Debug global.json didn't match [README.md, Sample.csproj, config/**, catalog/**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.in.transformations[0].sources[0].delegate.transformations[2] (ReplaceText)
┃ ┃ ┃ ┃ ┃ ┗ ┗  Info Will replace [Sample->weatherforecast-csha...(truncated)]
┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.in.transformations[0].sources[1] (Combo)
┃ ┃ ┃ ┃ ┃ ┃  Info Combo running as Chain
┃ ┃ ┃ ┃ ┃ ┃ engine.transformations[0].validated.delegate.in.transformations[0].sources[1].delegate (Chain)
┃ ┃ ┃ ┃ ┃ ┃  Info Running Chain(Include, RewritePath)
┃ ┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.in.transformations[0].sources[1].delegate.transformations[0] (Include)
┃ ┃ ┃ ┃ ┃ ┃ ┃  Info Will include [Sample.csproj]
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug .gitignore didn't match [Sample.csproj] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug .tanzuignore didn't match [Sample.csproj] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug Controllers/WeatherForecastController.cs didn't match [Sample.csproj] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug LICENSE didn't match [Sample.csproj] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug Program.cs didn't match [Sample.csproj] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug Properties/launchSettings.json didn't match [Sample.csproj] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug README.md didn't match [Sample.csproj] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug Sample.csproj matched [Sample.csproj] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug StaticFiles/index.html didn't match [Sample.csproj] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug WeatherForecast.cs didn't match [Sample.csproj] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug appsettings.Development.json didn't match [Sample.csproj] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug appsettings.json didn't match [Sample.csproj] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug catalog/catalog-info.yaml didn't match [Sample.csproj] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug config/workload.yaml didn't match [Sample.csproj] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┗ Debug global.json didn't match [Sample.csproj] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.in.transformations[0].sources[1].delegate.transformations[1] (RewritePath)
┃ ┃ ┃ ┃ ┃ ┗ ┗ Debug Path 'Sample.csproj' matched '^(?<folder>.*/)?(?<filename>([^/]+?|)(?=(?<ext>\.[^/.]*)?)$)' with groups {ext=.csproj, folder=null, filename=Sample.csproj, g0=Sample.csproj, g1=null, g2=Sample.csproj, g3=Sample.csproj, g4=.csproj} and was rewritten to 'weatherforecast-csharp.csproj'
┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.in.transformations[0].sources[2] (Combo)
┃ ┃ ┃ ┃ ┃ ┃  Info Combo running as Chain
┃ ┃ ┃ ┃ ┃ ┃ engine.transformations[0].validated.delegate.in.transformations[0].sources[2].delegate (Chain)
┃ ┃ ┃ ┃ ┃ ┃  Info Running Chain(Include, ReplaceText, Combo)
┃ ┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.in.transformations[0].sources[2].delegate.transformations[0] (Include)
┃ ┃ ┃ ┃ ┃ ┃ ┃  Info Will include [README.md, config/*.yaml, catalog/*.yaml]
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug .gitignore didn't match [README.md, config/*.yaml, catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug .tanzuignore didn't match [README.md, config/*.yaml, catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug Controllers/WeatherForecastController.cs didn't match [README.md, config/*.yaml, catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug LICENSE didn't match [README.md, config/*.yaml, catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug Program.cs didn't match [README.md, config/*.yaml, catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug Properties/launchSettings.json didn't match [README.md, config/*.yaml, catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug README.md matched [README.md, config/*.yaml, catalog/*.yaml] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug Sample.csproj didn't match [README.md, config/*.yaml, catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug StaticFiles/index.html didn't match [README.md, config/*.yaml, catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug WeatherForecast.cs didn't match [README.md, config/*.yaml, catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug appsettings.Development.json didn't match [README.md, config/*.yaml, catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug appsettings.json didn't match [README.md, config/*.yaml, catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug catalog/catalog-info.yaml matched [README.md, config/*.yaml, catalog/*.yaml] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug config/workload.yaml matched [README.md, config/*.yaml, catalog/*.yaml] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┗ Debug global.json didn't match [README.md, config/*.yaml, catalog/*.yaml] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.in.transformations[0].sources[2].delegate.transformations[1] (ReplaceText)
┃ ┃ ┃ ┃ ┃ ┃ ┗  Info Will replace [sample-app->weatherforecast-csha...(truncated)]
┃ ┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.in.transformations[0].sources[2].delegate.transformations[2] (Combo)
┃ ┃ ┃ ┃ ┃ ┃ ┃  Info Combo running as Chain
┃ ┃ ┃ ┃ ┃ ┃ ┃ engine.transformations[0].validated.delegate.in.transformations[0].sources[2].delegate.transformations[2].delegate (Chain)
┃ ┃ ┃ ┃ ┃ ┃ ┃  Info Running Chain(Merge, UniquePath)
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.in.transformations[0].sources[2].delegate.transformations[2].delegate.transformations[0] (Merge)
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃  Info Running Merge(InvokeFragment, Combo)
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.in.transformations[0].sources[2].delegate.transformations[2].delegate.transformations[0].sources[0] (InvokeFragment)
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.in.transformations[0].sources[2].delegate.transformations[2].delegate.transformations[0].sources[0].validated (Combo)
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃  Info Condition (#bsGitRepository != null) evaluated to false
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┗ ┗ null ()
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.in.transformations[0].sources[2].delegate.transformations[2].delegate.transformations[0].sources[1] (Combo)
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃  Info Combo running as Include
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ engine.transformations[0].validated.delegate.in.transformations[0].sources[2].delegate.transformations[2].delegate.transformations[0].sources[1].delegate (Include)
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃  Info Will include [**]
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug README.md matched [**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug catalog/catalog-info.yaml matched [**] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ ┗ ┗ Debug config/workload.yaml matched [**] -> included
┃ ┃ ┃ ┃ ┃ ┗ ┗ ╺ engine.transformations[0].validated.delegate.in.transformations[0].sources[2].delegate.transformations[2].delegate.transformations[1] (UniquePath)
┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.in.transformations[0].sources[3] (Combo)
┃ ┃ ┃ ┃ ┃ ┃  Info Combo running as Chain
┃ ┃ ┃ ┃ ┃ ┃ engine.transformations[0].validated.delegate.in.transformations[0].sources[3].delegate (Chain)
┃ ┃ ┃ ┃ ┃ ┃  Info Running Chain(Include, ReplaceText)
┃ ┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.in.transformations[0].sources[3].delegate.transformations[0] (Include)
┃ ┃ ┃ ┃ ┃ ┃ ┃  Info Will include [**/*.cs]
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug .gitignore didn't match [**/*.cs] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug .tanzuignore didn't match [**/*.cs] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug Controllers/WeatherForecastController.cs matched [**/*.cs] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug LICENSE didn't match [**/*.cs] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug Program.cs matched [**/*.cs] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug Properties/launchSettings.json didn't match [**/*.cs] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug README.md didn't match [**/*.cs] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug Sample.csproj didn't match [**/*.cs] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug StaticFiles/index.html didn't match [**/*.cs] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug WeatherForecast.cs matched [**/*.cs] -> included
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug appsettings.Development.json didn't match [**/*.cs] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug appsettings.json didn't match [**/*.cs] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug catalog/catalog-info.yaml didn't match [**/*.cs] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┃ Debug config/workload.yaml didn't match [**/*.cs] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┗ Debug global.json didn't match [**/*.cs] -> excluded
┃ ┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.in.transformations[0].sources[3].delegate.transformations[1] (ReplaceText)
┃ ┃ ┃ ┃ ┃ ┗ ┗  Info Will replace [Sample->weatherforecast_csha...(truncated)]
┃ ┃ ┃ ┃ ┗ ╺ engine.transformations[0].validated.delegate.in.transformations[0].sources[4] (Provenance)
┃ ┃ ┃ ┃ ┏ engine.transformations[0].validated.delegate.in.transformations[1] (UniquePath)
┃ ┃ ┃ ┃ ┃ Debug Multiple representations for path 'Controllers/WeatherForecastController.cs', will use the one appearing last 
┃ ┃ ┃ ┃ ┃ Debug Multiple representations for path 'WeatherForecast.cs', will use the one appearing last 
┃ ┗ ┗ ┗ ┗ Debug Multiple representations for path 'Program.cs', will use the one appearing last 
┗ ╺ engine.transformations[1] (UniquePath)
```
