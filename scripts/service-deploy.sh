#! /bin/zsh

function deploy_services () {
    local packages=("image" "listing" "booking" "slack" "user" "authentication" "authorizer")
    for package in "${packages[@]}"
    do
        echo "Deploying ${(C)package} service...";
        (cd "packages/service-$package" && pnpm run deploy);
    done
}

function deploy_composition () {
    local package="composition"
    echo "Deploying ${(C)package}...";
    (cd "packages/$package" && pnpm run deploy);
}

function deploy_appsync() {
    local package="appsync"
    echo "Deploying ${(C)package}...";
    (cd "$package" && pnpm run deploy);
}

deploy_services
deploy_composition
deploy_appsync