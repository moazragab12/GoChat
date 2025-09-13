# GoChat

A real-time chat application built in Go using TCP sockets. GoChat supports both public broadcasting and private messaging with a simple command-line interface.

## Table of Contents
- [Features](#features)
- [Architecture](#architecture)
- [Project Structure](#project-structure)
- [Installation & Setup](#installation--setup)
- [Usage](#usage)
  - [Compile the code](#compile-the-code-without-the-test-file)
  - [Starting the Server](#starting-the-server)
  - [Connecting as a Client](#connecting-as-a-client)
  - [Chat Commands](#chat-commands)
- [Message Format](#message-format)
- [Load Testing](#load-testing)
- [Technical Details](#technical-details)
  - [Server Implementation](#server-implementation)
  - [Client Implementation](#client-implementation)
  - [Performance Characteristics](#performance-characteristics)
- [Example Usage Session](#example-usage-session)
- [Error Handling](#error-handling)
- [Customization](#customization)
- [Limitations](#limitations)
- [Future Enhancements](#future-enhancements)
- [Contributing](#contributing)
- [License](#license)
- [Support](#support)

## Features

* **Real-time messaging**: Instant message delivery using TCP connections
* **Public chat**: Broadcast messages to all connected users
* **Private messaging**: Send direct messages to specific users
* **User join/leave notifications**: Automatic notifications when users connect or disconnect
* **Load testing**: Built-in load testing tool to evaluate server performance
* **Concurrent connections**: Support for multiple simultaneous users

## Architecture

The application consists of several key components:

* **Server**: Manages client connections and message routing
* **Client**: Handles user input and displays received messages
* **Message Model**: JSON-based message structure
* **Load Tester**: Performance testing tool for stress testing

## Project Structure

```
GoChat/
├── main.go          # Entry point and command router
├── Server.go        # Server implementation
├── Client.go        # Client implementation
├── Models.go        # Message data structures
├── Loadtest.go      # Load testing utility
└── README.md        # This file
```

## Installation & Setup

### Prerequisites

* Go 1.16 or higher
* Network connectivity (localhost)

### Installation

1. Clone or download the project files
2. Navigate to the project directory
3. Ensure all Go files are in the same directory

## Usage
### Compile the code without the test file
```
go build -o gochat.exe .\main.go .\Server.go .\Client.go .\Models.go
```
#### Starting the Server

```bash
.\gochat.exe server
```

The server will start listening on port 8080 and display:

```
Server is running on port 8080
```

#### Connecting as a Client

In a new terminal window:

```bash
.\gochat.exe client <username>
```

Replace `<username>` with your desired username.

Example:

```bash
.\gochat.exe client alice
```

### Chat Commands

Once connected, you can:

#### Public Messages

Simply type your message and press Enter:

```
Hello everyone!
```

#### Private Messages

Use the `@username:message` format:

```
@bob:Hello Bob, this is a private message
```

#### Exit Chat

Type `/quit` to disconnect:

```
/quit
```

## Message Format

Messages are transmitted as JSON objects with the following structure:

```json
{
  "from": "username",
  "to": "recipient_username",
  "Content": "message_content"
}
```

* `from`: Sender's username
* `to`: Recipient's username (empty for public messages)
* `Content`: The actual message content

## Load Testing

The application includes a built-in load testing tool to evaluate server performance.

### Running Load Tests

```bash
go run .\Loadtest.go
```

The load tester will:

* Create 300 concurrent client connections
* Measure connection latency
* Measure message sending latency
* Report success rate and performance metrics

### Load Test Output

```
Connection latency (min/avg/max): 1.2ms / 2.5ms / 8.1ms
Message latency (min/avg/max): 0.1ms / 0.3ms / 1.2ms
298/300 clients connected successfully
```

## Technical Details

### Server Implementation

* **Concurrent Handling**: Each client connection runs in its own goroutine
* **Thread Safety**: Uses mutex locks to protect shared client map
* **Connection Management**: Automatic cleanup of disconnected clients
* **Message Routing**: Supports both broadcast and direct messaging

### Client Implementation

* **Dual Goroutines**: Separate threads for sending and receiving messages
* **Input Parsing**: Automatic detection of private message syntax
* **Real-time Display**: Immediate display of incoming messages

### Performance Characteristics

### Performance Characteristics

- **Scalability**: Tested with 300+ concurrent connections  
  > **Note:** The 300-connection test was constrained by Windows and local hardware limits.  
  > The server itself can scale further on more powerful machines or properly tuned systems.
- **Low Latency**: Sub-millisecond message delivery
- **Memory Efficient**: Minimal overhead per connection
- **Graceful Degradation**: Handles connection failures gracefully


## Example Usage Session

### Terminal 1 (Server)

```bash
$ go build -o gochat.exe main.go Client.go Server.go Models.go
$ ./gochat.exe server
Server is running on port 8080
alice has joined the chat
bob has joined the chat
alice left the chat
```

### Terminal 2 (Client - Alice)

```bash
$ ./gochat.exe client alice
Welcome, alice ! Type /quit to exit.
Use '@bob:Hello' for private messages
Hello everyone!
[Server]: bob joined
@bob:Hey Bob!
/quit
```

### Terminal 3 (Client - Bob)

```bash
$ ./gochat.exe client bob
Welcome, bob ! Type /quit to exit.
Use '@bob:Hello' for private messages
[alice]: Hello everyone!
[alice]: Hey Bob!
[Server]: alice left
```

## Error Handling

The application handles several error scenarios:

* **Connection Failures**: Clients retry or report connection issues
* **Malformed Messages**: Invalid JSON messages are ignored
* **User Disconnections**: Automatic cleanup and notification
* **Network Issues**: Graceful degradation when connections are lost

## Customization

### Changing the Port

Modify the port in both `Server.go` and `Client.go`:

```go
// Server.go
listener, err := net.Listen("tcp", ":8080")

// Client.go  
conn, err := net.Dial("tcp", "localhost:8080")
```

### Adjusting Load Test Parameters

In `Loadtest.go`, modify:

```go
numClients := 300  // Number of concurrent clients
serverAddr := "localhost:8080"  // Server address
time.Sleep(2 * time.Second)  // Connection duration
```

## Limitations

* **Single Server**: No clustering or distributed setup
* **No Persistence**: Messages are not stored or logged
* **Basic Authentication**: Username-based identification only
* **No Encryption**: Messages are transmitted in plain text
* **Memory Storage**: Client list stored in memory only

## Future Enhancements

Potential improvements could include:

* Message persistence and history
* User authentication and authorization
* Encrypted communication (TLS/SSL)
* Web-based interface
* File sharing capabilities
* Chat rooms and channels
* Message acknowledgments
* Rate limiting and spam protection

## Contributing

To contribute to GoChat:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

## Support

For issues or questions:

* Check the code comments for implementation details
* Run the load tester to verify server performance
* Test with multiple clients to ensure functionality
