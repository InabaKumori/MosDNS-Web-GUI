# MosDNS Web GUI

MosDNS Web GUI is a program that provides a web interface for managing and interacting with MosDNS, a plugin-based DNS forwarder/traffic splitter.

## Installation

### Prerequisites

* Go 1.13 or later
* MosDNS installed and running

### Installation steps

1. Download the latest release of MosDNS Web GUI for your platform.
2. Extract the downloaded archive.
3. Run the installation script: `sudo ./install.sh`

## Usage

1. After installation, run the program: `/usr/local/mosdns-web-gui/mosdns-web-gui`
2. Access the web interface in your browser at `http://localhost:8080` (or the port you configured).
3. Use the web interface to configure MosDNS, update Geo data, manage rules, view status, and check logs.

## Configuration

The program can be configured by editing the `mosdns.yaml` file located in the `config` directory within the installation directory. Please refer to the MosDNS documentation for details on the configuration options.

## Testing

The program has been tested on various Linux distributions. However, it's always recommended to test the installation and functionality on your specific system before deploying it in a production environment.

## Contributing

Contributions are welcome! Please see the CONTRIBUTING.md file for details.

## License

MosDNS Web GUI is licensed under the GPLv3 license. Please see the LICENSE file for details.
