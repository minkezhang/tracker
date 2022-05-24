# API

To generate Golang protos, run

```bash
protoc \
  -I api/ \
  --go_out=api/ \
  --go_opt=paths=import \
  --go_opt=module=github.com/minkezhang/tracker/api \
  api/*proto
```
