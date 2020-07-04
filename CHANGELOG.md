# Changelog

## v0.3.0

* Fetch secrets concurrently, significantly decreasing execution time especially with input files containing lots of secrets.
* Add support for `# comments` in input files when lines begin with `#`.
* Exit with error if input file contains no secret pairs.

## v0.2.0

* Wrap output values in single quotes so they can be `source`d correctly when they contain special characters.
* Replace `'`character in secrets value with `'"'"'` for proper escaping ([Stack Overflow reference](https://stackoverflow.com/questions/1250079/how-to-escape-single-quotes-within-single-quoted-strings)).

## v0.1.0

Initial public release.
