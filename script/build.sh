#!/bin/bash
read -p 'if you want build one of services please enter the name, otherwise just skip it':  appname

if test -z "$appname" 
then
      docker-compose build
else
      docker-compose build $appname
fi