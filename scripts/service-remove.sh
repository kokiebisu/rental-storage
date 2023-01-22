#! /bin/zsh

function remove_services () {
    local packages=("image" "listing" "booking" "slack" "user" "authentication" "authorizer")
    for package in "${packages[@]}"
    do
        name=`echo ${package:0:1} | tr  '[a-z]' '[A-Z]'`${package:1}
        echo "Removing ${name} service...";
        (cd "packages/service-${package}" && pnpm run remove);
    done
}

function remove_composition() {
    local package="composition"
    name=`echo ${package:0:1} | tr  '[a-z]' '[A-Z]'`${package:1}
    echo "Removing ${name}...";
    (cd "packages/${package}" && pnpm run remove);
}

function remove_appsync() {
    local package="appsync"
    name=`echo ${package:0:1} | tr  '[a-z]' '[A-Z]'`${package:1}
    echo "Removing ${name}...";
    (cd "${package}" && pnpm run remove);
}

remove_services
remove_composition
remove_appsync