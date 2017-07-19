@rem mq_agent -
@rem Copyright (C) 2015-2017 Vutiliti, Inc.

@echo off

:init
    call pathx --cleanup
    call pathx --insert C:\Windows\System32
    call pathx --insert C:\Bin

    where /Q git.exe
    if %ERRORLEVEL% NEQ 0 call setup Git --noversion >NUL

:parse

:main
    call PrintOutput {Header} Vutiliti mq_agent
    echo.
    echo   WRITTEN  in Go
    echo   RUNS     on each mq_agent
    echo   PROVIDES reading meter data and
    echo            uploads it to the bridge
    echo.

    call pathx --append "%UserProfile%\Dropbox\Public\Device-Setup\setup\rtl-sdr-release\x64"

    call PrintOutput {Highlight} Setting up Go..
    REM call setup Go --noversion >NUL

    set "GOARCH="
    set "GOOS="
    REM Clear these out just in case (Go 1.2+ no longer require them)
    set "GOBIN="
    set "GOROOT="
    set "GOPATH="

    pushd "%~dp0..\..\..\.."
    set "GOPATH=%CD%"

    if not exist "%GOPATH%\bin" mkdir "%GOPATH%\bin"
    call pathx --insert "%GOPATH%\bin"
    call pathx --append "%AppData%\Go\bin"

    set "GOPATH=%CD%;%AppData%\Go"
    popd

    call go version
    echo.

    :: Display current Go envars.
    call PrintOutput {Highlight} Go envars..
    setlocal EnableDelayedExpansion
    for /F "tokens=*" %%G in ('set GO') do (
        set "__G=%%G"
        if not "!__G:~0,14!"=="GOOGLE_API_KEY" (
            echo.  %%G
        )
    )
    endlocal
    echo.

    REM call PrintOutput {Highlight} Common Go commands..
    REM echo   go build
    REM echo   go clean ^&^& go build
    REM echo.

    REM call PrintOutput {Highlight} Build commands..
    REM echo   build windows raspi
    REM echo.
