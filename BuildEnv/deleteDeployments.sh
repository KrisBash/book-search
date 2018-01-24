#!/bin/bash
source ./BuildEnv/getVersion.sh

#vars
tag=$fullVersion

#Container enum
images=(  "frontend" "book-api" "cache-db" "book-website")

#Enum and purge repositories
for i in "${images[@]}"
do
    echo "Enumerate active pod for $i"
    kubectl get pods |grep $i |grep  $tag
    echo "Enumerating stale pods for $i"
    stalePods=`kubectl get deployment |grep $i |grep -v $tag|awk '{print $1}'`
    echo $stalePods
    for pod in $stalePods ; do
        echo "Deleting pod $pod"
        kubectl delete deployment $pod 
        #kubectl delete pods $pod
    done
done

