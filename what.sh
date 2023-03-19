#!/usr/bin/env bash



mode=${1:-"quick"}

NAME=$(basename  "${PWD}")
AWS_ACCOUNT_ID=$(aws sts get-caller-identity --query "Account" --output text)
AWS_REGION=$(aws configure get region)
NOW=$(date +%Y%m%d.%H%M%S)

echo "${AWS_ACCOUNT_ID} ${AWS_REGION} ${mode}"

if [ "${mode}" == 'quick' ]; then
    make build
    ./bin/whatever encrypt
    docker build -t "${NAME}:local" .
    docker tag "${NAME}:local" "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:latest"
    docker tag "${NAME}:local" "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:${NOW}"
    docker push "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:latest"
    docker push "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:${NOW}"
    echo -n "${NOW}" > what.tag
    (cd ../ever && cdk deploy 'whatever-*')
elif [ "${mode}" == 'full' ]; then
    make build
    ./bin/whatever encrypt
    aws ecr get-login-password --region "${AWS_REGION}" | docker login --username AWS --password-stdin "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com"
    docker build -t "${NAME}:local" .
    docker tag "${NAME}:local" "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:latest"
    docker tag "${NAME}:local" "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:${NOW}"
    docker push "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:latest"
    docker push "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:${NOW}"
    echo -n "${NOW}" > what.tag
    (cd ../ever && cdk diff 'whatever-*')
    (cd ../ever && cdk deploy 'whatever-*')
elif [ "${mode}" == 'diff' ]; then
    (cd ../ever && cdk diff 'whatever-*')
elif [ "${mode}" == 'deploy' ]; then
    (cd ../ever && cdk deploy 'whatever-*')
fi




