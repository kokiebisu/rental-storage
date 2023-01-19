#! /bin/zsh

function remove_services () {
    local package=("image" "listing" "booking" "slack" "user" "authentication" "authorizer")
    for service in "${services[@]}"
    do
        echo "Removing ${(C)service} service...";
        (cd "packages/service-$service" && pnpm run remove);
    done
}

function remove_composition() {
    local package="composition"
    echo "Removing ${(C)package}...";
    (cd "packages/${package}" && pnpm run remove);
}

function remove_appsync() {
    local package="appsync"
    echo "Removing ${(C)package}...";
    (cd "${package}" && pnpm run remove);
}

remove_services
remove_composition
remove_appsync