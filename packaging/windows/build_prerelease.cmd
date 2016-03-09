:: Build keybase.exe with prerelease options
pushd %GOPATH%\src\github.com\keybase\client\go\keybase
set GOARCH=386
go generate
for /f %%i in ('winresource.exe -cv') do set KEYBASE_VERSION=%%i
echo %KEYBASE_VERSION%
for /f %%i in ('winresource.exe -cb') do set KEYBASE_BUILD=%%i
echo %KEYBASE_BUILD%
go build -a -tags "prerelease production" -ldflags="-X github.com/keybase/client/go/libkb.CustomBuild=%KEYBASE_BUILD%"

:: Then build kbfsdokan
pushd %GOPATH%\src\github.com\keybase\kbfs\kbfsdokan
go build -tags "production prerelease"
popd

:: Then the desktop:
pushd  %GOPATH%\src\github.com\keybase\client\desktop
start /WAIT npm i & exit 0
npm run package -- --arch ia32 --platform win32 --appVersion %KEYBASE_VERSION%
popd