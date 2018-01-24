#!/bin/bash
source ./BuildEnv/getVersion.sh

echo "Update k8s specs to reference version: ${fullVersion}"
grep -rl 'fullVersion' ./k8s-services | xargs sed -i "s/{{fullVersion}}/${fullVersion}/g"
kubectl apply -R -f ./k8s-services/