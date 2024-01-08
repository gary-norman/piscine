#! /bin/bash

find ./ -name "*.sh" | sed "s/.*\///" | sed "s/.sh//" | sort -nr
