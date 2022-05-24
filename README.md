# tracker
Project to track queue of media that needs to be consumed.

Formerly "Loose Record" Excel spreadsheet.

```bash
go run internal/importer/main.go \
  --input data/database.csv \
  --output data/database.textproto
```

## Proposed API

```
* Search(e dpb.Entry): search across all scrapers and allow user to choose from
  variety of IDs
* Add(e dpb.Entry): add after getting scraper-specific ID (return err if scraper
  ID not found in request)
* Patch(e dpb.Entry): manually update data
* Sync(e dpb.Entry): get data from upstream sources
* Delete(e dpb.Entry)
* Bump(e dpb.Entry): increment chapter or episode count
* BumpMajor(e dpb.Entry): increment volume or season, reset chapter and episode
  count to 1
* Filter(title_asc, corpus_asc, ...): return subset of collection
* Recommend()
```

## Proposed UI

TUI (cli)
Web --> gRPC service
