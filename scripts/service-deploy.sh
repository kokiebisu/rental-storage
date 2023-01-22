#! /bin/zsh

function deploy_services () {
    local packages=("image" "listing" "booking" "slack" "user" "authentication" "authorizer")
    for package in "${packages[@]}"
    do
        name=`echo ${package:0:1} | tr  '[a-z]' '[A-Z]'`${package:1}
        echo "Deploying ${name} service...";
        (cd "packages/service-$package" && pnpm run deploy);
    done
}

function deploy_composition () {
    local package="composition"
    name=`echo ${package:0:1} | tr  '[a-z]' '[A-Z]'`${package:1}
    echo "Deploying ${name}...";
    (cd "packages/$package" && pnpm run deploy);
}

function deploy_appsync() {
    local package="appsync"
    name=`echo ${package:0:1} | tr  '[a-z]' '[A-Z]'`${package:1}
    echo "Deploying ${name}...";
    (cd "$package" && pnpm run deploy);
}

deploy_services
deploy_composition
deploy_appsync