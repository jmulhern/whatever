#!/usr/bin/env bash

make build

./bin/whatever encrypt

mode=${1:-"quick"}

NAME=$(basename  "${PWD}")
AWS_ACCOUNT_ID=$(aws sts get-caller-identity --query "Account" --output text)
AWS_REGION=$(aws configure get region)
NOW=$(date +%Y%m%d.%H%M%S)

echo "${AWS_ACCOUNT_ID} ${AWS_REGION} ${mode}"

if [ "${mode}" == 'quick' ]; then
    docker build -t "${NAME}:local" .
    docker tag "${NAME}:local" "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:latest"
    docker tag "${NAME}:local" "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:${NOW}"
    docker push "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:latest"
    docker push "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:${NOW}"
    echo -n "${NOW}" > what.tag
    (cd ../ever && cdk deploy 'whatever-*')

elif [ "${mode}" == 'full' ]; then
    aws ecr get-login-password --region "${AWS_REGION}" | docker login --username AWS --password-stdin "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com"
    docker build -t "${NAME}:local" .
    docker tag "${NAME}:local" "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:latest"
    docker tag "${NAME}:local" "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:${NOW}"
    docker push "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:latest"
    docker push "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:${NOW}"
    echo -n "${NOW}" > what.tag
    (cd ../ever && cdk diff 'whatever-*')
    (cd ../ever && cdk deploy 'whatever-*')
fi




