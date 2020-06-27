# `awssecret2env`

Convert secrets stored in AWS Secrets Manager to environment variables.

## Example

Given a file like `secrets.txt` that maps environment variable names to secret names in AWS Secret Manager, `awssecret2env` replaces the secret names with their values stored in AWS, and prints the resulting env to `stdout`.

```bash
# secrets.txt
DB_HOST=db/dev/DB_HOST
DB_USER=db/dev/DB_USER
DB_PASSWORD=db/dev/DB_PASSWORD
```

```bash
# Usage: awssecret2env <input-file>
awssecret2env secrets.txt > .env

# cat .env
# DB_HOST=<REDACTED>
# DB_USER=<REDACTED>
# DB_PASSWORD=<REDACTED>
```

## Download

Downloaded files must be made executable before they can be run.

* [MacOS](https://awssecret2env.s3.amazonaws.com/master/awssecret2env-macos)
* [Windows](https://awssecret2env.s3.amazonaws.com/master/awssecret2env-windows)
* [Linux X86_64](https://awssecret2env.s3.amazonaws.com/master/awssecret2env-linux64)
* [Linux ARM6](https://awssecret2env.s3.amazonaws.com/master/awssecret2env-linuxarm6)
* [Linux ARM7](https://awssecret2env.s3.amazonaws.com/master/awssecret2env-linuxarm7)

You can also download and execute `awssecret2env` programmatically.

```bash
PLATFORM=macos # supported platforms: "macos", "windows", "linux64", "linuxarm6", or "linuxarm7"

wget https://awssecret2env.s3.amazonaws.com/master/awssecret2env-${PLATFORM}
chmod +x awssecret2env-${PLATFORM}
mv awssecret2env-${PLATFORM} /usr/local/bin/awssecret2env
```
