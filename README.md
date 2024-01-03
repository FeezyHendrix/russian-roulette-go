# Russian Roulette

## Overview

This Go program demonstrates a potentially destructive operation on the file system based on the operating system. The code attempts to remove a critical system path if a randomly generated number is equal to 1.

## Prerequisites

- Go installed on your system.
- **Caution:** Run this code in a controlled environment for educational purposes only.

## Usage

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/DestructiveOperationExample.git
    ```

2. Navigate to the project directory:

    ```bash
    cd DestructiveOperationExample
    ```

3. Run the Go program:

    ```bash
    go run main.go
    ```

4. **Important:** This code randomly attempts to remove a critical system path (e.g., "/System/Library/CoreServices" on macOS). Ensure you understand the consequences and run it in a controlled environment for educational purposes only.

## Disclaimer

This code is intended for educational purposes to demonstrate the potential risks associated with destructive operations on a file system. **Do not run this code on a real system or in a production environment.** Understand the consequences and use responsibly.

## License

This project is licensed under the [MIT License](LICENSE).
