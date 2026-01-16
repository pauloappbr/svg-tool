$BinaryName = "svg-tool.exe"
$InstallDir = "$env:USERPROFILE\bin"
$SourcePath = "cmd/svg-tool/main.go"

Write-Host "Compiling $BinaryName..." -ForegroundColor Cyan
go build -o $BinaryName $SourcePath

if (-Not (Test-Path $BinaryName)) {
    Write-Error "Build failed. Check if Go is installed and the path is correct."
    exit 1
}

# Create installation directory if it doesn't exist
if (-Not (Test-Path $InstallDir)) {
    New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null
}

# Move the executable
Move-Item -Path $BinaryName -Destination "$InstallDir\$BinaryName" -Force

Write-Host "Binary installed to $InstallDir" -ForegroundColor Green

# Add to User PATH persistently if not present
$UserPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($UserPath -notlike "*$InstallDir*") {
    Write-Host "Adding $InstallDir to User PATH..." -ForegroundColor Yellow
    [Environment]::SetEnvironmentVariable("Path", "$UserPath;$InstallDir", "User")
    Write-Host "PATH updated. You may need to restart your terminal for changes to take effect." -ForegroundColor Green
} else {
    Write-Host "$InstallDir is already in your PATH." -ForegroundColor Green
}

Write-Host "Installation successful!"