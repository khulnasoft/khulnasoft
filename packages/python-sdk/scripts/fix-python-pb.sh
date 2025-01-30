#!/bin/bash

rm -rf khulnasoft/envd/__pycache__
rm -rf khulnasoft/envd/filesystem/__pycache__
rm -rf khulnasoft/envd/process/__pycache__

sed -i '.bak' 's/from\ process\ import/from khulnasoft.envd.process import/g' khulnasoft/envd/process/* khulnasoft/envd/filesystem/*
sed -i '.bak' 's/from\ filesystem\ import/from khulnasoft.envd.filesystem import/g' khulnasoft/envd/process/* khulnasoft/envd/filesystem/*

rm khulnasoft/envd/process/*.bak
rm khulnasoft/envd/filesystem/*.bak
