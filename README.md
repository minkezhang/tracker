# tracker
Project to track queue of media that needs to be consumed.

Formerly "Loose Record" Excel spreadsheet.

```bash
go run internal/importer/main.go \
  --input data/database.csv \
  --output data/database.textproto
```
