#!/bin/bash
read -p 'if you want run one of services please enter the name, otherwise just skip it':  appname

if test -z "$appname" 
then
      docker-compose up -d
else
      docker-compose up -d $appname
fi