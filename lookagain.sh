#! /bin/bash

find ./ -name "*.sh" | sed "s/.*\///" | sort -nr
