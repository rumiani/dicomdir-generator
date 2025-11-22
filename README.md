# DICOMDIR Generator

A cross-platform desktop application for generating DICOMDIR files from DICOM folders. Built with Go and Fyne, this tool helps medical imaging professionals and patients create DICOMDIR files required by hospitals and radiology departments.

## What is DICOMDIR?

DICOMDIR is an index file that catalogs all DICOM images in a folder structure, following the DICOM standard for media exchange. Many hospitals require this file to efficiently browse and import medical imaging studies.

## Features

- **Simple folder selection** - Choose any folder containing DICOM files (.dcm)
- **Automatic DICOMDIR generation** - Creates the index file directly in your DICOM folder
- **File counter** - Displays the total number of DICOM files found
- **Cross-platform** - Works on Linux, macOS, and Windows
- **Clean interface** - Simple, intuitive UI built with Fyne
- **No file modification** - Your original DICOM files remain unchanged

## Installation

### Prerequisites

- **Go** 1.21 or later
- **Java Runtime Environment (JRE)** - Required by dcm4che

### Setup Steps

1. **Clone the repository**
   ```bash
   git clone https://github.com/rumiani/dicomdir-generator.git
   cd dicomdir-generator
   ```

2. **Verify assets**
   
   Ensure the `dcm4che-5.34.1` folder is present in the `assets/` directory (included in the repo).

3. **Install dependencies**
   ```bash
   go mod download
   ```

## Usage

### Running the Application

```bash
go run main.go
```

### Steps to Generate DICOMDIR

1. Click **"Select DICOM Folder"** and choose your folder containing .dcm files
2. The app will display the number of DICOM files found
3. Click **"Generate DICOMDIR"**
4. The DICOMDIR file will be created in the selected folder
5. Send the entire folder (including DICOMDIR) to your hospital

### Building an Executable

To create a standalone executable:

```bash
go build -o dicomdir-generator
```

Then run:
```bash
./dicomdir-generator
```

## Project Structure

```
dicomdir-generator/
├── main.go                    # Application entry point
├── go.mod                     # Go module dependencies
├── ui/                        # User interface components
│   ├── app.go                # Main application window
│   └── widgets.go            # UI widgets and controls
├── backend/                   # Core functionality
    └── dcm4che-5.34.1/       # DCM4CHE toolkit (Java-based)
│   ├── generator.go          # DICOMDIR generation logic
    └── utils.go              # Helper functions
```

## How It Works

1. The application uses the **dcmdir** utility from the dcm4che toolkit to generate the DICOMDIR file
2. Users select a folder through the Fyne-based GUI
3. The app recursively scans for all `.dcm` files and displays the count
4. When "Generate DICOMDIR" is clicked, the app executes dcmdir to create the index
5. The DICOMDIR file is created directly in the selected folder alongside your DICOM files

## Technical Details

### Dependencies

- **[Fyne](https://fyne.io/)** - Modern, cross-platform GUI toolkit for Go
- **[dcm4che](https://www.dcm4che.org/)** - Java-based DICOM toolkit for medical imaging

### Requirements

- The selected folder must contain `.dcm` files (either directly or in subfolders)
- Java must be installed and accessible in your system PATH
- Sufficient disk space for the DICOMDIR file (typically <1MB)

## Troubleshooting

**"Java not found" error**
- Ensure JRE is installed: `java -version`
- Add Java to your system PATH

**No DICOM files found**
- Verify your files have the `.dcm` extension
- Check that the folder path is correct

**Permission denied**
- Ensure you have write permissions in the selected folder

## Support & Contribution

If you find this project helpful, please consider:
- ⭐ **Starring the repository** on [GitHub](https://github.com/rumiani/dicomdir-generator)
- 🐛 **Reporting issues** or suggesting features
- 🤝 **Contributing** via pull requests

## License

[Add your license information here]

## Acknowledgments

Built with dcm4che, the open-source DICOM toolkit maintained by the medical imaging community.