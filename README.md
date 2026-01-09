# GayaLang

GayaLang is a simple scripting language designed specifically for testing APIs. It allows developers to write scripts to send requests, handle responses, and automate API testing in a clear and concise way.

<!-- Test -->

## Features

- Write scripts to test API endpoints
- Send GET, POST, and other HTTP requests
- Supports execution context for storing runtime variables
- Lightweight and easy to use from the command line

> More features will be added soon!

## Installation

Make sure you have Go installed. Then clone and build the project:

```bash
git clone https://github.com/yourusername/gayalang.git
cd gayalang
go build ./cmd/gaya
```

## Usage

Run a GayaLang script with:

```bash
go run cmd/gaya/main.go cmd/gaya/test.gaya
```

Example output:

```
Request login succeeded: 200
Request getData succeeded: 200
Execution context: map[]
```

## Sample `test.gaya` Script

```gaya
request login {
    POST "https://webhook.site/ec2b106b-a4c6-4620-9873-14cf4af6fe20"
    expect status = 200
}

request getData {
    GET "https://judge0-be.vercel.app/api/indexasdad"
    expect status = 200
}
```

## Contributing

Contributions are welcome! Feel free to fork the repo and submit improvements.
