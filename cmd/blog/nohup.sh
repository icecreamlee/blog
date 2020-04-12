#!/bin/sh
nohup ./blog >> blog.log 2>&1 &
echo $! > blog.pid