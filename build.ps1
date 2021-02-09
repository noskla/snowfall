if ( $null -eq (Get-Command "go.exe" -ErrorAction SilentlyContinue) ) {
    Write-Host "Go binary not found. Install golang using 'choco install go' " -NoNewline
    Write-Host "or by visiting https://golang.org/"
    Exit
}
if ( $null -eq (Get-Command "sass.exe" -ErrorAction SilentlyContinue) ) {
    Write-Host "SASS binary not found. Install SASS using 'choco install sass' " -NoNewline
    Write-Host "or by visiting https://sass-lang.com/install"
    Exit
}

if ( -not (Test-Path -Path '.\dist' -PathType Container) ) {
    New-Item -ItemType Directory -Force -Path '.\dist'
} else {
    Remove-Item -Path '.\dist\*' -Recurse
}
 
Write-Host "Compiling... " -NoNewline
$env:GOARCH = 'amd64'
ForEach ($os in 'windows','linux','freebsd') {
    $env:GOOS = $os
    $filepath = ".\dist\snowfall.$os"
    if ( $os -eq 'windows' ) {
        $filepath += '.exe'
    }
    go.exe build -o $filepath
    if ($?) {
        Write-Host "$os ok," -NoNewline
    } else {
        Write-Host "$os failed," -NoNewline
    }
}
Write-Host "Done"

Write-Host "Copying assets..." -NoNewline
Copy-Item -Path '.\templates' -Destination '.\dist\' -Recurse
Copy-Item -Path '.\static' -Destination '.\dist\' -Recurse
Write-Host "Done"

Write-Host "Compiling SASS..." -NoNewline
New-Item -ItemType Container -Force -Path ".\dist\static\css" | Out-Null
$sassfiles = Get-ChildItem -Path ".\sass\"
ForEach ($sass in $sassfiles) {
    sass.exe --no-source-map --style compressed ".\sass\$($sass.Name)" ".\dist\static\css\$($sass.BaseName).css"
}
Write-Host "OK"

Write-Host "Binaries are ready in .\dist\ directory."
