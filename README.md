# Clipboard to Citrix

This project reads text from the clipboard and simulates typing it into a Citrix session using keyboard events.

## Prerequisites

- Go 1.24 or later
- A Citrix session running on your machine

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/vpanal/clipboard2citrix.git
    cd clipboard2citrix
    ```

2. Install the dependencies:

    ```sh
    go mod tidy
    ```

## Usage

1. Build the project:

    ```sh
    go build -o clipboard2citrix
    ```

2. Run the executable:

    ```sh
    ./clipboard2citrix
    ```

## How It Works

1. The program reads text from the clipboard.
2. It simulates an `Alt+Tab` keyboard event to switch to the Citrix session.
3. It types out the text from the clipboard character by character.

## Code Overview

- `main.go`: The main file that contains the logic for reading from the clipboard and simulating keyboard events.
- `go.mod`: The Go module file that lists the dependencies.
- `go.sum`: The Go checksum file that ensures the integrity of the dependencies.

## Dependencies

- `github.com/atotto/clipboard`: For reading from the clipboard.
- `github.com/micmonay/keybd_event`: For simulating keyboard events.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.