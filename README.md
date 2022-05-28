# truffle
Project to track queue of media that needs to be consumed.

## Examples
```bash
go install github.com/minkezhang/truffle/truffle@latest

# Add sample entry.
truffle add \
  --titles=Sabikui \
  --corpus=anime \
  --score=6.3 \
  --providers=crunchyroll \
  --queued=true \
  --studios=OZ \
  --directors="Atsushi Itagaki" \
  --writers="Sadayuki Murai" \
  --season=1 \
  --episode=4

# Start watching the next season of Sabikui. truffle supports partial title and
# corpus matching.
truffle bump \
  --title=Sabikui \
  --corpus=anime \
  --major

# Re-rate the entry.
truffle patch \
  --title=Sabikui \
  --corpus=anime \
  --score=6.4

truffle get --title=Sabikui

# Search the user database as well as the MAL API for similar entries.
truffle search \
  --title=Sabikui \
  --corpus=anime \
  --apis=truffle\
  --apis=mal

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
