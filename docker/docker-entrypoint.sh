#!/usr/bin/env bash
set -e

file_env() {
   local var="$1"
   local fileVar="${var}_FILE"
   local def="${2:-}"

   if [ "${!var:-}" ] && [ "${!fileVar:-}" ]; then
      echo >&2 "error: both $var and $fileVar are set (but are exclusive)"
      exit 1
   fi
   local val="$def"
   if [ "${!var:-}" ]; then
      val="${!var}"
   elif [ "${!fileVar:-}" ]; then
      val="$(< "${!fileVar}")"
   fi
   export "$var"="$val"
   unset "$fileVar"
}

buildGo() {
	# if [ ! -f "quepasa.db" ]; then
		go build -o quepasa main.go
		#echo "error: quepasa not found"
		#exit 1
	# else
	# 	echo "quepasa founded"
	# fi
}

file_env "PGPASSWORD"
file_env "SIGNING_SECRET"

# go build -o quepasa main.go
buildGo
exec "$@"