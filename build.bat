@echo off
setlocal EnableDelayedExpansion
cd /D "%~dp0"

for %%a in (%*) do set "%%~a=1"
if not "%release%"=="1" set debug=1

set common_flags=
set debug_flags=
set release_flags=

if "%debug%"=="1" set flags=%common_flags% %debug_flags%
if "%release%"=="1" set flags=%common_flags% %release_flags%

if not exist out mkdir out

go build -o out/redding-fgc-website.exe %flags% || exit /b 1
