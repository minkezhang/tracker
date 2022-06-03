# Truffle API Proto
---

## Compile

Run from project root.
```bash
protoc -I ./ \
  --go_out=api/ \
  --go_opt=paths=import \
  --go_opt=module=github.com/minkezhang/truffle/api \
  api/*proto
```
