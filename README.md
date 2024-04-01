# MosDNS Web GUI (Development Version)

**Important Note:** This project is currently in its early stages of development and is not yet fully functional. Use with caution and expect potential bugs or incomplete features.

MosDNS Web GUI is a program that aims to provide a web interface for managing and interacting with MosDNS, a plugin-based DNS forwarder/traffic splitter. This project is a work in progress and currently offers basic functionalities.

### Features (In Development)

* Viewing and editing MosDNS configuration
* Updating GeoIP and GeoSite data
* Managing rule files (whitelists, blocklists, etc.)
* Displaying MosDNS status and logs

### Installation

#### Prerequisites

* Go 1.13 or later
* MosDNS installed and running

#### Installation Steps

1. Clone or download this repository to your local machine.
2. Open a terminal and navigate to the project directory.
3. Run the following command to build and run the program:

```
go run cmd/mosdns-web-gui/main.go
```


4. Access the web interface in your browser at `http://localhost:8080` (or the port you configured).

**Note:** You might need to adjust configuration file paths and other settings in the code to match your specific environment.

### Usage

The web interface provides sections for:

* **Configuration:** View and edit MosDNS configuration options.
* **Geo Data Update:** Trigger the update of GeoIP and GeoSite data files.
* **Rule Management:** Add, edit, and delete rules for whitelists, blocklists, etc.
* **Status:** View the current status of MosDNS.
* **Logs:** View MosDNS logs and filter by log level.

### Testing

This program has been successfully tested on macOS and Linux. However, it's always recommended to test the installation and functionality on your specific system before deploying it in a production environment.

### Contributing

Contributions are welcome! Please see the `CONTRIBUTING.md` file for details.

### License

MosDNS Web GUI is licensed under the GNU General Public License v3.0. Please see the `LICENSE` file for details.

### Author

This project is developed by [inabakumori](https://github.com/inabakumori).

**Please note that this project is still under development and may not be suitable for production use yet.**
