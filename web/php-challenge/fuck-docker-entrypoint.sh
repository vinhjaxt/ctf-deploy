#!/bin/sh
pm2 resurrect
exec docker-php-entrypoint "$@"
