"Uninstalling greg..."
try {
    $dirToRemove = "C:\Program Files\Greg"
    $path = [Environment]::GetEnvironmentVariable("Path", "Machine")
    $dirs = $path -split ";"
    "Remove $dirToRemove from PATH Environment Variable"
    $dirs = $dirs | Where-Object { $_ -ne $dirToRemove } -ErrorAction Stop
    $newPath = $dirs -join ";"
    "Remove $dirToRemove and all its contents"
    Remove-Item -Path $dirToRemove -Recurse -Force -ErrorAction Stop
    [Environment]::SetEnvironmentVariable("Path", $newPath, "Machine")
    "Done! :) hope you enjoyed using greg!"
} catch {
    $ex = $_.Exception
    Write-Error "Failed to uninstall greg: $ex"
}