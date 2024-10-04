# simple-datagram
Datagram is the unit of data exchanged between a pair of internet modules (includes the internet header). This program is built for enhancing the understanding of datagram based on RFC 791.

This repository contains a simple example of how to create and send datagrams using the UDP protocol in Go. The project includes two parts: a UDP server and a UDP client that sends datagrams to the server.

## Overview 
A datagram is a unit of data exchanged between two devices over a network using a connectionless protocol such as UDP (User Datagram Protocol). Unlike TCP, which establishes a connection before data transfer, UDP simply sends data packets (datagrams) without requiring a connection. This makes UDP faster but less reliable, as there is no guarantee of packet delivery.

This project demonstrates how to:
- Create a UDP server that listens for incoming datagrams.
- Create a UDP clients that sends datagrams to the server.

## Project Structure
- `server.go`: The UDP server code that listens for incoming datagrams on a specified port.
- `client.go`: The UDP client code that sends a datagram message to the server.

## How It Works
1. The server listens on port `8080` for incoming datagrams.
2. The client sends a datagram message (`"Hello, this is a datagram!"`) to the server.
3. The server receives and prints the message along with the sender's address.

## Running Example 
### Prerequisites 
- `Go` installed on your machine.

### Steps
1. **Run The Server**:
   - Open a terminal and navigate to the directory containing the `server.go` file.
   - Start the server using the following command:
   ```
   go run server.go
   ```
   - The server will start listening on `localhost:8080` for incoming UDP datagrams.
2. **Run The Client**:
   - Open another terminal and navigate to the directory containing the `client.go` file:
   - Send a datagram message to the server using the following command:
   ```
   go run client.go
   ```
3. **Output**:
   - In the terminal server, you will see output like this when it receives a datagram: 
   ```
   Listening on localhost:8080
   Received Hello, this is a datagram from 127.0.0.1:xxxxx
   ```
   - The client will display:
   ```
   Sent message: Hello, this is a datagram!
   ```

## Code Explanation 
### Server Code (`server.go`)
The server is implemented using Go's `net` package:
- `net.ResolveUDPaddr`: Resolves the address for UDP communication.
- `net.ListenUDP`: Binds the UDP address and starts listening for incoming datagrams.
- `conn.ReadFromUDP`: Reads the incoming datagram into a buffer.
### Client Code (`client.go`)
The client sends a datagram using:
- `net.ResolveUDPaddr`: Resolves the server's UDP address.
- `net.DialUDP`: Creates a UDP connection to the server. 
- `conn.Write`: Sends a datagram message over the UDP connection.

## Notes 
- UDP is a connectionless protocol, meaning there is no guarantee that the datagram will reach the destination. It's faster but less reliable compared to TCP.
- This example uses `localhost` communication, but you can modify the IP and port to test on different machines over a network.

## License 
This project is open-source and licensed under the **MIT License**.
