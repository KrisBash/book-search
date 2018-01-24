#!/bin/bash
source ./BuildEnv/getVersion.sh

echo "Update k8s specs to reference version: ${fullVersion}"
grep -rl 'fullVersion' ./k8s-apps | xargs sed -i "s/{{fullVersion}}/${fullVersion}/g"
grep -rl 'registryUrl' ./k8s-apps | xargs sed -i "s/{{registryUrl}}/${registryUrl}/g"

kubectl apply -R -f ./k8s-apps/