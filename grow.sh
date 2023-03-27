#!/usr/bin/env bash

mode=${1:-"quick"}

NAME=$(basename  "${PWD}")
AWS_ACCOUNT_ID=$(aws sts get-caller-identity --query "Account" --output text)
AWS_REGION=$(aws configure get region)
NOW=$(date +%Y%m%d.%H%M%S)
BUCKET="whatever-20230327-x"
CLOUDFRONT_DISTRIBUTION_ID="E30072L8WN05EW"

echo "${AWS_ACCOUNT_ID} ${AWS_REGION} ${mode}"

if [ "${mode}" == 'quick' ]; then
    docker build -t "${NAME}:local" .
    docker tag "${NAME}:local" "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:latest"
    docker tag "${NAME}:local" "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:${NOW}"
    docker push "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:latest"
    docker push "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:${NOW}"
    echo -n "${NOW}" > sprout.txt
    npm run build
    aws s3 sync --delete dist "s3://${BUCKET}/dist"
    aws s3 sync --delete public "s3://${BUCKET}/public"
    aws cloudfront create-invalidation --distribution-id="${CLOUDFRONT_DISTRIBUTION_ID}" --paths='/*' > /dev/null
    (cd ../garden && cdk deploy 'whatever-*')

elif [ "${mode}" == 'full' ]; then
    aws ecr get-login-password --region "${AWS_REGION}" | docker login --username AWS --password-stdin "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com"
    docker build -t "${NAME}:local" .
    docker tag "${NAME}:local" "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:latest"
    docker tag "${NAME}:local" "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:${NOW}"
    docker push "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:latest"
    docker push "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:${NOW}"
    echo -n "${NOW}" > sprout.txt
    npm run build
    aws s3 sync --delete dist "s3://${BUCKET}/dist"
    aws s3 sync --delete public "s3://${BUCKET}/public"
    aws cloudfront create-invalidation --distribution-id="${CLOUDFRONT_DISTRIBUTION_ID}" --paths='/*' > /dev/null
    (cd ../garden && cdk diff 'whatever-*')
    (cd ../garden && cdk deploy 'whatever-*')

elif [ "${mode}" == 'build' ]; then
    aws ecr get-login-password --region "${AWS_REGION}" | docker login --username AWS --password-stdin "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com"
    docker build -t "${NAME}:local" .
    docker tag "${NAME}:local" "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:latest"
    docker tag "${NAME}:local" "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:${NOW}"

elif [ "${mode}" == 'push' ]; then
    aws ecr get-login-password --region "${AWS_REGION}" | docker login --username AWS --password-stdin "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com"
    docker build -t "${NAME}:local" .
    docker tag "${NAME}:local" "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:latest"
    docker tag "${NAME}:local" "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:${NOW}"
    docker push "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:latest"
    docker push "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:${NOW}"
    echo -n "${NOW}" > sprout.txt

elif [ "${mode}" == 'diff' ]; then
    (cd ../garden && cdk diff 'whatever-*')

elif [ "${mode}" == 'deploy' ]; then
    (cd ../garden && cdk deploy 'whatever-*')

elif [ "${mode}" == 'backend' ]; then
    aws ecr get-login-password --region "${AWS_REGION}" | docker login --username AWS --password-stdin "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com"
    docker build -t "${NAME}:local" .
    docker tag "${NAME}:local" "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:latest"
    docker tag "${NAME}:local" "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:${NOW}"
    docker push "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:latest"
    docker push "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:${NOW}"
    echo -n "${NOW}" > sprout.txt
    (cd ../garden && cdk deploy 'whatever-*')

elif [ "${mode}" == 'frontend' ]; then
    npm run build
    aws s3 sync --delete dist "s3://${BUCKET}/dist"
    aws s3 sync --delete public "s3://${BUCKET}/public"
    aws cloudfront create-invalidation --distribution-id="${CLOUDFRONT_DISTRIBUTION_ID}" --paths='/*' > /dev/null

elif [ "${mode}" == 'build' ]; then
    docker build -t "${NAME}:local" .
    docker tag "${NAME}:local" "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:latest"
    docker tag "${NAME}:local" "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${NAME}:${NOW}"
fi




