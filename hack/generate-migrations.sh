#!/bin/bash
set -e
go-bindata -pkg migration -o db/migration/migration.go -prefix "db/migration/sql/" db/migration/sql