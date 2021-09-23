# `awssecret2env`

[![CirclCI Build Status](https://circleci.com/gh/runwayml/awssecret2env.svg?style=shield)](https://app.circleci.com/pipelines/github/runwayml/awssecret2env)

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
# Usage: awssecret2env [OPTIONS] <input-file>
awssecret2env secrets.txt
# DB_HOST=<REDACTED>
# DB_USER=<REDACTED>
# DB_PASSWORD=<REDACTED>
```

With no options, `awssecret2env` prints the resulting environment variables to `stdout`. You can specify an output file with the `--output` flag, and optionally add an `export` statement to each line with `--export`.

```bash
awssecret2env --output .env --export secrets.txt

cat .env
# export DB_HOST=<REDACTED>
# export DB_USER=<REDACTED>
# export DB_PASSWORD=<REDACTED>

source .env
# The env vars should now be injected in your shell
```

### Input File

Input files are in the following format:

```bash
# lines beginning with "#" are ignored as a comment
ENV_VAR_NAME=secret-name/secret-key
ENV_VAR_NAME_2=secret-name/secret-key-2
ENV_VAR_NAME_3=other-secret-name/other-key
```

The secret's key is always interpreted as the string following the last `/` character in the line.

> NOTE: Secret **names** may contain `/` characters, but secret **keys** SHOULD NOT.

## Download

Downloaded files must be made executable before they can be run.

* [MacOS (Intel)](https://awssecret2env.s3.amazonaws.com/latest/awssecret2env-macos)
* [MacOS (Apple Silicon)](https://awssecret2env.s3.amazonaws.com/latest/awssecret2env-macos-arm64)
* [Windows](https://awssecret2env.s3.amazonaws.com/latest/awssecret2env-windows)
* [Linux (X86_64)](https://awssecret2env.s3.amazonaws.com/latest/awssecret2env-linux64)
* [Linux (ARM6)](https://awssecret2env.s3.amazonaws.com/latest/awssecret2env-linuxarm6)
* [Linux (ARM7)](https://awssecret2env.s3.amazonaws.com/latest/awssecret2env-linuxarm7)

You can also download and execute `awssecret2env` programmatically.

```bash
PLATFORM=macos # supported platforms: "macos", "macos-arm64", "windows", "linux64", "linuxarm6", or "linuxarm7"
VERSION=latest # supported versions: "latest", "master", "v0.1.0", etc.

wget https://awssecret2env.s3.amazonaws.com/${VERSION}/awssecret2env-${PLATFORM}
chmod +x awssecret2env-${PLATFORM}
mv awssecret2env-${PLATFORM} /usr/local/bin/awssecret2env
```

## Usage

```
Usage: ./build/bin/awssecret2env [OPTIONS] <input-file> ...
Note: <input-file> is a required positional argument.
  -r, --aws-region string   The name of the AWS region where secrets are stored (default "us-east-1")
  -e, --export              Prepends "export" statements in front of the output env variables
  -h, --help                Show this screen
  -o, --output string       Redirects output to a file instead of stdout
```
