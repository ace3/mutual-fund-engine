# NOBI Investment gRPC Backend

## Project Structure

- `cmd/server/main.go`: Entry point for the gRPC server
- `internal/config/`: Loads environment variables
- `internal/database/`: MySQL connection
- `internal/grpc/`: gRPC server and handler logic
- `internal/models/`: Data models
- `internal/services/`: Business logic
- `internal/utils/`: Utility functions (e.g., rounding)
- `migrations/`: SQL migration files
- `proto/nobi/`: Protobuf definitions
- `pkg/pb/nobi/`: Generated gRPC code

## Setup

1. Copy `.env.example` to `.env` and fill in your DB credentials
2. Run migrations to set up the database:
   ```sh
   make migrate
   ```
3. Generate gRPC code from `proto/nobi/nobi.proto`:
   ```sh
   make proto
   ```
4. (Optional) If you see duplicate generated files in `pkg/pb/nobi/pkg/pb/nobi/`, clean them up:
   ```sh
   rm -rf pkg/pb/nobi/pkg
   ```
5. Build and run the server:
   ```sh
   make run
   ```

## Example gRPC Payloads

### AddUser
```json
{
  "name": "Arman",
  "username": "arman123"
}
```

### TopUp
```json
{
  "user_id": 1,
  "amount_rupiah": 5000
}
```

### Withdraw
```json
{
  "user_id": 1,
  "amount_rupiah": 6000
}
```

### UpdateTotalBalance
```json
{
  "current_balance": 20000
}
```
