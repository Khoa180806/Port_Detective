# Port Detective automated installer for Windows PowerShell
# Usage: iwr -useb https://raw.githubusercontent.com/Khoa180806/Port_Detective/master/scripts/install.ps1 | iex

$ErrorActionPreference = 'Stop'

$Owner = "Khoa180806"
$Repo = "Port_Detective"
$Binary = "pd.exe"

Write-Host "==> Detecting system architecture..." -ForegroundColor Cyan

$Arch = "x86_64"
if ($env:PROCESSOR_ARCHITECTURE -eq "ARM64") {
    $Arch = "arm64"
}

Write-Host "==> Fetching release assets..." -ForegroundColor Cyan
$ReleaseUrl = "https://api.github.com/repos/$Owner/$Repo/releases/latest"

$DownloadUrl = ""
$Tag = "latest"

try {
    $Response = Invoke-RestMethod -Uri $ReleaseUrl -UseBasicParsing -TimeoutSec 10
    if ($Response.tag_name) {
        $Tag = $Response.tag_name
    }
    # Match Windows asset matching the system architecture
    $Asset = $Response.assets | Where-Object { $_.name -like "*Windows*${Arch}*.zip" } | Select-Object -First 1
    if ($Asset -and $Asset.browser_download_url) {
        $DownloadUrl = $Asset.browser_download_url
        $ZipName = $Asset.name
    }
} catch {
    Write-Host "Notice: Release API lookup failed, falling back to direct release URL." -ForegroundColor Yellow
}

if (-not $DownloadUrl) {
    # Fallback using goreleaser project_name naming convention: port-detective_Windows_x86_64.zip
    $ZipName = "port-detective_Windows_${Arch}.zip"
    $DownloadUrl = "https://github.com/$Owner/$Repo/releases/download/$Tag/$ZipName"
}

$TempDir = Join-Path $env:TEMP "port-detective-install"
if (Test-Path $TempDir) {
    Remove-Item $TempDir -Recurse -Force
}
New-Item -ItemType Directory -Path $TempDir | Out-Null

$ZipPath = Join-Path $TempDir $ZipName
Write-Host "==> Downloading $ZipName ($Tag)..." -ForegroundColor Cyan
Invoke-WebRequest -Uri $DownloadUrl -OutFile $ZipPath -UseBasicParsing

Write-Host "==> Extracting files..." -ForegroundColor Cyan
Expand-Archive -Path $ZipPath -DestinationPath $TempDir -Force

# Locate target directory in user PATH (WindowsApps is standard in user PATH by default on Win10/11)
$TargetDir = Join-Path $env:LOCALAPPDATA "Microsoft\WindowsApps"
if (-not (Test-Path $TargetDir)) {
    # Fallback to user profile bin directory and append to PATH
    $TargetDir = Join-Path $env:USERPROFILE "bin"
    if (-not (Test-Path $TargetDir)) {
        New-Item -ItemType Directory -Path $TargetDir | Out-Null
    }
    $UserPath = [Environment]::GetEnvironmentVariable("PATH", [EnvironmentVariableTarget]::User)
    if ($UserPath -notlike "*$TargetDir*") {
        [Environment]::SetEnvironmentVariable("PATH", "$UserPath;$TargetDir", [EnvironmentVariableTarget]::User)
        $env:PATH = "$env:PATH;$TargetDir"
    }
}

$SourceExe = Get-ChildItem -Path $TempDir -Filter $Binary -Recurse | Select-Object -First 1
if (-not $SourceExe) {
    $SourceExe = Get-ChildItem -Path $TempDir -Filter "*.exe" -Recurse | Select-Object -First 1
}

if (-not $SourceExe) {
    Write-Error "Could not find binary executable in extracted archive."
    exit 1
}

$DestPath = Join-Path $TargetDir "pd.exe"
Copy-Item -Path $SourceExe.FullName -Destination $DestPath -Force

Remove-Item $TempDir -Recurse -Force -ErrorAction SilentlyContinue

Write-Host ""
Write-Host "=========================================================" -ForegroundColor Green
Write-Host "  Port Detective ($Tag) installed successfully!" -ForegroundColor Green
Write-Host "  Destination: $DestPath" -ForegroundColor Gray
Write-Host "  Try running: pd --help or pd check 8080" -ForegroundColor Yellow
Write-Host "=========================================================" -ForegroundColor Green
