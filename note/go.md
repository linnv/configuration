```
err := json.Indent(&out, []byte(in), "", "\t")
    if err != nil {
        return in
    }
```

:GoImpl SmartQC/taskgen/asr.Asr
