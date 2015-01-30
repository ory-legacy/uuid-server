# uuid-server

An RESTful server to create [universally unique identifiers (version 4)](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_.28random.29).

[![Build Status](https://travis-ci.org/ory-platform/uuid-server.svg)](https://travis-ci.org/ory-platform/uuid-server)

## Install

To install and run this service, do:

```
> go install github.com/ory-platform/uuid-server
> <gopath>/bin/uuid-server
```

The server exposes some arguments:

| Option | Description |
|--|--|
| `-host=""` | The address to listen on. Leave empty to listen on all interfaces |
| `-port=80` | The port to listen on | 

## Usage

To create a new UUID, create a `POST` request on `/`.
You will receive a JSON response containing:

```
{
    "id": "c8237384-0ebf-4228-b244-f1104d00f1b4"
}
```