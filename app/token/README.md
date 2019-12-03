# Introduction
This app uses NodeJS sdk to communicate with peers

# command
protoc \
    --proto_path=protos/definitions \
    --js_out=import_style=commonjs,binary:protos \
    protos/definitions/*.proto


## Task list
- change bookmark data type
- code refactor