#!/bin/bash
# Bash script used with CircleCI to release new versions

if ! [ ${CIRCLE_BRANCH} == "master" ]; then
  echo "Circle Branch is not master ... exiting"
  exit 0
fi

echo "Circle Branch is master ... building artifacts"

release=$(git describe --always --tags)
sha=`echo ${CIRCLE_SHA1} | cut -c1-6`
content_type="application/zip"
date=`date -R`

systems="linux windows darwin"
for sys in $systems
do
  name=maestro-${sys}.zip
  zip -j ${CIRCLE_ARTIFACTS}/${name} ./bin/${sys}/maestro* README.md ./scripts/maestro_autocomplete CHANGELOG.md
  s3artifact -bucket $AWS_BUCKET -name ${release}/${name} ${CIRCLE_ARTIFACTS}/${name}
  s3artifact -bucket $AWS_BUCKET -name latest/${name} ${CIRCLE_ARTIFACTS}/${name}
done

# Upload Web Page to S3 Bucket
pandoc -f markdown -t html README.md > ${CIRCLE_ARTIFACTS}/index.html
s3artifact -content-type text/html -bucket $AWS_BUCKET -name index.html ${CIRCLE_ARTIFACTS}/index.html
s3artifact -content-type image/gif -bucket $AWS_BUCKET -name demo.gif ${CIRCLE_ARTIFACTS}/demo.gif


aws_endpoint=https://s3.amazonaws.com/${AWS_BUCKET}
# Check for official release ie v0.5.1 not v0.5.1-kj34kdf
if [[ $release =~ ^v([0-9]+).([0-9]+).([0-9]+)$ ]]; then

  current_version=$(curl -s ${aws_endpoint}/LATEST)
  # If the version in S3 is not the latest then update it
  if [[ $current_version != $release ]]; then
    echo $release > ${CIRCLE_ARTIFACTS}/LATEST
    s3artifact -bucket $AWS_BUCKET -name LATEST -acl public-read ${CIRCLE_ARTIFACTS}/LATEST
  fi
fi
