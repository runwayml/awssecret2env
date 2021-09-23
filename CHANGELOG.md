# Changelog

## v0.3.2

* Add Apple Silicon support and build with Go v1.17.

## v0.3.1

* Fix bug where fetching secrets with non-string values caused `awssecret2env` to crash. JSON secret values of all types are now converted to strings when they are transformed to environment variables.

## v0.3.0

* Fetch secrets concurrently, significantly decreasing execution time especially with input files containing lots of secrets.
* Add support for `# comments` in input files when lines begin with `#`.
* Exit with error if input file contains no secret pairs.

## v0.2.0

* Wrap output values in single quotes so they can be `source`d correctly when they contain special characters.
* Replace `'`character in secrets value with `'"'"'` for proper escaping ([Stack Overflow reference](https://stackoverflow.com/questions/1250079/how-to-escape-single-quotes-within-single-quoted-strings)).

## v0.1.0

Initial public release.
