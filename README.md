# aireqlog

[![GoDoc](https://godoc.org/github.com/EndpointElf/aireqlog?status.svg)](https://godoc.org/github.com/EndpointElf/aireqlog)
[![Build Status](https://github.com/EndpointElf/aireqlog/workflows/Go/badge.svg)](https://github.com/EndpointElf/aireqlog/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/EndpointElf/aireqlog)](https://goreportcard.com/report/github.com/EndpointElf/aireqlog)

This project implements a structured logging system for monitoring and auditing requests, specifically tracking source code sent to AI agents. Initially created for [EndpointElf](https://endpointelf.com/), a tool designed to extract HTTP routes from Go source code, this AI Request Log ensures full transparency and accountability in handling source code with AI-driven processes.

## Overview

In EndpointElf, AI tools assist in various code analysis and generation tasks, often involving (proprietary) source code. To ensure accountability and traceability of requests made to these AI agents, this logging system captures detailed metadata about each request.

## Features

- **Project Metadata**: Logs the closest parent project (identified by a `.git` directory) to contextualize the request.
- **File Information**: Records the filename of the specific source file involved.
- **Git Commit**: Logs the current git commit (HEAD) to provide version control context.
- **AI Engine and Endpoint**: Captures the specific AI engine used and the URL of the endpoint the data was sent to.
- **Timestamping**: Adds a timestamp for precise tracking of when requests are made.

## How It Works

2. **Request Logging**: When a request is sent, the log records project metadata, source file details, AI engine, endpoint URL, and a timestamp.
3. **Storage and Retrieval**: Logs are stored in a structured format, allowing for easy querying and auditing of all requests made.


## Contributing

Contributions are welcome! If you’d like to enhance this system, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License.
