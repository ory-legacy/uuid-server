# uuid-server

An RESTful server to create [universally unique identifiers (version 4)](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_.28random.29).

[![Build Status](https://travis-ci.org/ory-am/uuid-server.svg)](https://travis-ci.org/ory-am/uuid-server)

## Install

To install and run this service, do:

```
> go install github.com/ory-am/uuid-server
> <gopath>/bin/uuid-server
```

## Config

To specify which port to listen on, create an environment variable called `PORT`. Per default, this application
listens on port 80.

To specify which host to listen on, create an environment variable called `HOST`. Per default, this application
listens on all hosts (which equals to an empty string).

## Usage

To create a new UUID, create a `POST` request on `/`.
You will receive a JSON response containing:

```
{
    "id": "c8237384-0ebf-4228-b244-f1104d00f1b4"
}
```
