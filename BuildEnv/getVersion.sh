#!/bin/bash
major=0
minor=1
patch=0

ver=$(git describe --exact-match 2> /dev/null || echo "`git symbolic-ref HEAD 2> /dev/null | cut -b 12-`-`git log --pretty=format:\"%h\" -1`")

fullVersion="${major}.${minor}.${patch}${ver}"
export fullVersion

echo $fullVersion
