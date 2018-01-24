#!/bin/bash
source ./BuildEnv/getVersion.sh

#vars
tag=$fullVersion
registry=$registryUrl
sudo az acr login --name $registryName

#Container enum
images=(  "frontend" "book-api" "cache-db" "book-website")

#Enum and purge repositories
for i in "${images[@]}"
do
    echo "Enumerate active repository for $i"
    az acr repository list -n $registryName |grep $i |grep  $tag |tr -d \" |tr -d ,
    echo "Enumerating stale repositories for $i"
    staleRepos=`az acr repository list -n $registryName |grep $i |grep -v $tag |tr -d \" |tr -d ,`
    echo $staleRepos
    for repo in $staleRepos
    do
       echo "Deleting stale repositories for $i"
       az acr repository delete -n $registryName --repository $repo --yes
    done
done
