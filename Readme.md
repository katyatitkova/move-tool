# Move tool

A simple CLI for slicing long samples into Ableton Note / Ableton Move presets

## Quick Start for Non-Developers

### Using Move Tool

1. Prepare an audio sample containing up to 16 parts (e.g., a sample with 16 equal-length slices).
2. Open Terminal (macOS) or Command Prompt (Windows).
3. Run the Move Tool command:
   - For macOS:
     ```
     ./move-tool slice -i /path/to/your/sample.wav -n 16 -o /Users/your-username/Desktop
     ```
   - For Windows:
     ```
     move-tool.exe slice -i C:\path\to\your\sample.wav -n 16 -o C:\Users\YourUsername\Desktop
     ```
4. Move Tool will slice the original sample into 16 pieces and create a Move/Note preset with a random name in the specified output directory.
5. If you find this tool helpful, please star the repository!

### Quick start for macOS Users
1. Download the macOS version of Move Tool from [the latest release](https://github.com/alexfedosov/move-tool/releases/latest).
2. Open Finder and navigate to your Downloads folder.
3. Double-click the downloaded file (it should be named something like `move-tool-macos.zip`) to unzip it.
4. Open Terminal (you can find it by pressing Cmd + Space and typing "Terminal").
5. In Terminal, type `cd ` (with a space after it), then drag the folder containing the unzipped `move-tool` into the Terminal window. Press Enter.
6. Now you can use Move Tool. For example, to slice a sample, type:
   ```
   ./move-tool slice -i /path/to/your/sample.wav -n 16 -o /Users/your-username/Desktop
   ```
   Replace `/path/to/your/sample.wav` with the actual path to your audio file, and `/Users/your-username/Desktop` with where you want to save the output.

### Quick start for Windows Users
1. Download the Windows version of Move Tool from [the latest release](https://github.com/alexfedosov/move-tool/releases/latest).
2. Open File Explorer and navigate to your Downloads folder.
3. Right-click the downloaded file (it should be named something like `move-tool-windows.zip`) and select "Extract All". Choose a location to extract the files.
4. Open Command Prompt (you can find it by pressing Win + R, typing "cmd", and pressing Enter).
5. In Command Prompt, type `cd ` (with a space after it), then type the path to the folder where you extracted Move Tool. For example:
   ```
   cd C:\Users\YourUsername\Downloads\move-tool
   ```
6. Now you can use Move Tool. For example, to slice a sample, type:
   ```
   move-tool.exe slice -i C:\path\to\your\sample.wav -n 16 -o C:\Users\YourUsername\Desktop
   ```
   Replace `C:\path\to\your\sample.wav` with the actual path to your audio file, and `C:\Users\YourUsername\Desktop` with where you want to save the output.


## For Developers

- [Overview](#overview)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Ensure that you have [Go 1.22](https://golang.org/dl/) installed.
2. Clone the repo:

    ```sh
    git clone https://github.com/alexfedosov/move-tool.git
    ```

3. Navigate to the project directory:

    ```sh
    cd move-tool
    ```

4. Install dependencies:

    ```sh
    go mod tidy
    ```

## Usage

```sh
go run . slice -i <file path> -n <number of samples> -o <output directory>
```

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/YourFeature`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Open a Pull Request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
