# City Journey Indexer

Go backend for [city-journey-nft](https://github.com/m15flores/city-journey-nft). Reads on-chain city data directly from the deployed contract via `go-ethereum` and persists it to SQLite, exposing it through a small HTTP API.

## What it does

- Reads all `MintNFT` events emitted by the contract on Arbitrum One, and for each token, the full `CityData` struct (name, coordinates, date range)
- Persists that data to a local SQLite database
- Serves it over HTTP, with endpoints to list all cities, fetch one by id, find which city a given date falls into, or find which cities overlap a given date range

## Architecture decisions

**SQLite over Postgres.** The dataset is 10 rows, static, with no concurrent writes. Running a Postgres instance for that would be operational overhead with no real benefit here.

**Plain `net/http` over a framework.** The API has three routes. The goal was to learn Go's HTTP fundamentals (routing, handlers, `http.Error`, JSON encoding) before reaching for a framework that abstracts them away.

**Historical reads only, no live subscription.** The collection is fully minted, there are no more events to come, so a WebSocket subscription would have nothing to listen for.

## API

| Method | Path | Description |
|---|---|---|
| GET | `/cities` | All cities |
| GET | `/cities/{id}` | One city by token id |
| GET | `/cities?at=YYYY-MM-DD` | The city whose date range contains this date |
| GET | `/cities?from=YYYY-MM-DD&to=YYYY-MM-DD` | Cities whose date range overlaps this one |

### Conventions

- Date ranges are closed on the left, open on the right: the day a move happens belongs to the city being moved **to**, not the one left behind.
- `toDate = 0` means "still there". Used for the most recent city, which has no end date yet.
- `at` and `from`/`to` are mutually exclusive; combining them returns `400`.
- A `from`/`to` range with no matching cities returns `200` with `[]`, not `404`.
- `at` with no matching city returns `404`, since it targets a single resource.

## Project structure

| Path | Contents |
|---|---|
| `cmd/indexer/` | entrypoint: wires everything together and starts the server |
| `internal/model/` | the `City` type, shared across packages |
| `internal/onchain/` | reads events and contract state via `go-ethereum` |
| `internal/store/` | SQLite persistence and queries |
| `internal/api/` | HTTP handlers |
| `abi/` | contract ABI used to generate onchain bindings |

## Running it

​```bash
go run ./cmd/indexer
​```

Reads the contract, populates `cities.db`, and starts the API on `:8080`.

## Testing

​```bash
go test ./... -v
​```

`internal/store` and `internal/api` are covered, including the two trickiest edge cases in the data: the date boundary between two consecutive cities, and the `toDate = 0` city. `internal/onchain` isn't covered by automated tests. Exercising `ReadMintedCities` would require either a live RPC connection or a mocked Ethereum client, and neither was justified for a one-time indexing job against a collection that's already fully minted.

| Package | Coverage |
|---|---|
| `internal/store` | 81.2% |
| `internal/api` | 90.5% |

The uncovered lines in `internal/store` are error-handling branches for SQLite failures (a broken write, a malformed row) that aren't triggered by the current tests. Forcing them would need a corrupted database file or an incompatible schema, which wasn't worth the setup for this project.