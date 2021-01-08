![build](https://github.com/Whyeasy/zally-cleaner/workflows/build/badge.svg)
![status-badge](https://goreportcard.com/badge/github.com/Whyeasy/zally-cleaner)
![Github go.mod Go version](https://img.shields.io/github/go-mod/go-version/Whyeasy/zally-cleaner)

# atlasdb-slow-queries

A Go program which queries the AtlasDB API to retrieve and log slow queries and suggested indexes.

## Requirements

Provide the Group ID, which is the Project ID within AtlasDB: `--groupId` or as env variable `GROUP_ID`.

Provide the Project ID, which is the connection string and the port number: `--projectId` or as env variable `PROJECT_ID`.

Provide the Public Key of the created API key within AtlasDB: `--publicKey` or as env variable `PUBLIC_KEY`.

Provide the Private Key of the created API key within AtlasDB: `--privateKey` or as env variable `PRIVATE_KEY`.

### Optional

Change log format; `--logFormat <string>` or as env variable `LOG_FORMAT`. The default value is `"logfmt"`.

Change log level; `--logLevel` or as env variable `LOG_LEVEL`. The default value is `info`.

Change amount of hours in the past you want to retrieve data for. `--since int`. The default value is `24`.
