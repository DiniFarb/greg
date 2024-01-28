$version = "v0.0.2-alpha"
$url = "https://github.com/DiniFarb/greg/releases/download/$version/greg.zip"
"Install greg from $url"
$installDir = "C:\Program Files\Greg"
"Install dir: $installDir"
$appName = "greg.exe"
$pathEnv = "C:\Program Files\Greg"
try {
    New-Item -ItemType Directory -Path $installDir -Force -ErrorAction Stop
    Invoke-WebRequest $url -OutFile "$version.zip" -ErrorAction Stop
    Expand-Archive -Path "$version.zip" -DestinationPath $installDir -Force -ErrorAction Stop
    "Add $pathEnv to PATH Environment Variable"
    $envPath = [Environment]::GetEnvironmentVariable("Path", "Machine")
    if ($envPath -notlike "*$pathEnv*") {
        $newPath = "$envPath;$pathEnv"
        [Environment]::SetEnvironmentVariable("Path", $newPath, "Machine")
    }
    "DONE! :) You can now run greg from any directory but you need to restart your terminal first"
} catch  {
    $ex = $_.Exception
    Write-Error "Failed to install greg: $ex"
    exit 1
}
