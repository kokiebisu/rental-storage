#! /bin/zsh

function install_dependencies_root() {
    echo "Installing root dependencies...";
    (pnpm install);
}

function install_dependencies_services () {
    local services=("image" "listing" "booking" "slack" "user" "authentication" "authorizer")
    for service in "${services[@]}"
    do
        echo "Installing ${(C)service} service dependencies...";
        (cd "packages/service-$service" && pnpm install);
    done
}

function install_dependencies_appsync() {
    local package="appsync"
    echo "Installing ${(C)package} service dependencies...";
    (cd appsync && pnpm install);
}

install_dependencies_root
install_dependencies_services
install_dependencies_appsync