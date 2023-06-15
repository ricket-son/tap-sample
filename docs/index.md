# Weatherforecast TAP

A sample C# application to run on TAP

## Install

```bash
tanzu apps workload create weatherforecast-csharp \
    --git-repo https://github.com/ricket-son/tap-sample \
    --sub-path weatherforecast-csharp \
    --git-branch main \
    --type web \
    --label app.kubernetes.io/part-of=weatherforecast-csharp \
    --namespace tap-test-dev \
    --yes
```

