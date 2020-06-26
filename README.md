# AWS Secret Manager 2 Env

Convert secrets stored in AWS Secrets Manager to bash ENVs.

```bash
# secrets.dev input file mapping ENV variable names to secrets stored in AWS Secret Manager
DB_HOST=db/dev/DB_HOST
DB_USER=db/dev/DB_USER
DB_PASSWORD=db/dev/DB_PASSWORD
```

```bash
# Example
awssecret2env secrets.env > .env
```

```bash
DB_HOST=<REDACTED>
DB_USER=<REDACTED>
DB_PASSWORD=<REDACTED>
```

