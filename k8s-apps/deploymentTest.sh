#!/bin/bash
source ./BuildEnv/getVersion.sh

expectedPods=3
runningPods=`kubectl get pods |grep $fullVersion|grep Running |wc -l`

res=0
x=0
while [ $x -le 10 ]
do
  echo "Current runnning pods: $runningPods. Expected pods: $expectedPods "
  if [ $runningPods -gt 2 ]
  then
      echo "All expected pods are running"
      res=1
      x=10
  fi
  sleep 60
  runningPods=`kubectl get pods |grep $fullVersion|grep Running |wc -l`
  x=$(( $x + 1 ))
done

if [ $res -eq 1 ]
then
    echo "All pods online as expected. Continuing"
else
    echo "Maximum wait reached and not all pods are running"
    exit 100
fi

#Container enum
images=( "book-api" "book-website" "cache-db" )

#Enum and purge repositories
echo "Dumping pod logs..."

activePods=`kubectl get pods |grep  $fullVersion|awk '{print $1}'`
for pod in $activePods ; do
    echo "Logs for pod: pod "
    kubectl logs $pod
done
