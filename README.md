# truffle
Project to track queue of media that needs to be consumed.

## Examples
```bash
go install github.com/minkezhang/truffle/truffle@latest

# Add sample entry.
truffle add \
  --title=Sabikui \
  --corpus=anime \
  --score=6.3 \
  --provider=crunchyroll \
  --queued=true \
  --studio=OZ \
  --director="Atsushi Itagaki" \
  --writer="Sadayuki Murai" \
  --season=1 \
  --episode=4

# Start watching the next season of Sabikui. truffle supports partial title and
# corpus matching.
truffle bump \
  --title=Sabikui \
  --corpus=anime \
  --major

# Re-rate the entry and mark the MAL entry as duplicate, which will be filtered
# out in searches.
truffle patch \
  --title=Sabikui \
  --corpus=anime \
  --score=6.4 \
  --link=mal:48414

truffle get --title=Sabikui

# Search the user database as well as the MAL API for similar entries.
truffle search \
  --title=Sabikui \
  --corpus=anime \
  --api=truffle\
  --api=mal

# Delete the entry.
truffle delete --title=Sabikui

```

## Uninstall

```bash
go clean -i github.com/minkezhang/truffle/truffle
```

## Feature Docket

```
* database.Filter(queued bool, score_acending bool)
* database.Recommend()
```
