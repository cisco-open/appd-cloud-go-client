#!/bin/bash

openapiversion="6.2.1"
if ![ -z "$1" ]
    then
        openapiversion=$1
fi

while IFS= read -r url || [[ -n "$url" ]]; do
    title=$( curl -s $url | jq '. | {title: .info.title}' )
    title=$( echo $title | sed 's/.*://' | tr -dc '[:alnum:]\n\r' )
    title=$( echo $title | sed -e "s/^AppDynamics//" -e "s/API$//" -e "s/Service$//" )
    title=$( echo $title | tr '[:upper:]' '[:lower:]' )
    echo $title

    version=$( curl -s $url | jq '. | {title: .info.version}' )
    version=$( echo $version | sed 's/.*://' | sed 's/\..*//' | tr -dc '[:alnum:]\n\r' )
    echo $version

    
    docker run --rm \
        -v ${PWD}:/local openapitools/openapi-generator-cli:v${openapiversion} generate \
        -i ${url} \
        -g go \
        -o /local/apis/v${version}/${title} \
        -c /local/config.yaml \
        --template-dir /local/custom-templates \
        --ignore-file-override /local/.openapi-generator-ignore \
        --additional-properties=packageName=${title} \
        --global-property apis,apiTests,models,modelTests,supportingFiles="client.go:README.md:response.go:utils.go:openapi.yaml"
done < apis.txt



if ![ -z "$2" ]
    then
        addurl=$2

        title=$( curl -s $2 | jq '. | {title: .info.title}' )
        title=$( echo $title | sed 's/.*://' | tr -dc '[:alnum:]\n\r' )
        title=$( echo $title | sed -e "s/^AppDynamics//" -e "s/API$//" -e "s/Service$//" )
        title=$( echo $title | tr '[:upper:]' '[:lower:]' )
        echo $title

        version=$( curl -s $2 | jq '. | {title: .info.version}' )
        version=$( echo $version | sed 's/.*://' | sed 's/\..*//' | tr -dc '[:alnum:]\n\r' )
        echo $version

        docker run --rm \
                -v ${PWD}:/local openapitools/openapi-generator-cli:v${openapiversion} generate \
                -i ${addurl} \
                -g go \
                -o /local/apis/v${version}/${title} \
                -c /local/config.yaml \
                --template-dir /local/custom-templates \
                --ignore-file-override /local/.openapi-generator-ignore \
                --additional-properties=packageName=${title} \
                --global-property apis,apiTests,models,modelTests,supportingFiles="client.go:README.md:response.go:utils.go:openapi.yaml"
fi


