#!/bin/bash
oldstring=github.com/brocaar/chirpstack-api
newstring=github.com/fraanx/chirpstack-v3-api

targetlist=$(grep -r $oldstring|cut -d ':' -f 1)

for fname in $targetlist
do
echo file is $fname
sed -i "s#$oldstring#$newstring#g" "$fname"

done

