#!/bin/bash

set -e

function vscode_prepare() {
    echo "Setting up vscode configurations...";
    (mkdir -p .vscode && cp config/settings.json .vscode/settings.json);
}


vscode_prepare