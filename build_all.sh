#!/bin/bash
set -e

APP_NAME="dicomdir-generator"
APP_ID="com.github.rumiani.dicomdir-generator"

echo "Building Linux binary..."
go build -o "${APP_NAME}-linux" .
echo "✓ Linux build complete: ${APP_NAME}-linux"

echo "Building Windows binary..."
# Backup and modify go.mod for cross-compilation
cp go.mod go.mod.backup
sed -i 's/^go 1\.25/go 1.21/' go.mod

fyne-cross windows -name "$APP_NAME" -app-id "$APP_ID"

# Extract Windows binary from fyne-cross output
unzip -j "fyne-cross/dist/windows-amd64/${APP_NAME}.zip" -d . "${APP_NAME}.exe"
mv "${APP_NAME}.exe" "${APP_NAME}-windows.exe"

# Restore original go.mod
mv go.mod.backup go.mod

# Clean up all fyne-cross artifacts
rm -rf fyne-cross

echo "✓ Windows build complete: ${APP_NAME}-windows.exe"
echo ""
echo "✓ Build complete! You now have:"
echo "  - ${APP_NAME}-linux"
echo "  - ${APP_NAME}-windows.exe"
echo ""
echo "All temporary files cleaned up."