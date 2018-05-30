#!/bin/bash

function usage(){
	printf "Script args :\n"
	printf "\t--name | -n              : Mandatory: To specify the directory name of the example you want to run;\n"
	printf "\t--action | -a            : Mandatory: Choose the action among [start, stop, restart];\n"
	printf "\t--help | -h              : Optional: To display the help.\n"
	printf "\nexample: ./run-docker.sh -n cors-simple -a start \n\n"
}

function start_app() {
    echo "Starting app..."

   docker-compose -f docker/docker-compose.yml build
   docker-compose -f docker/docker-compose.yml up -d
}

function stop_app() {
    echo "Stopping app..."
    docker-compose -f docker/docker-compose.yml rm -s -f
}

function restart_app() {
   stop_app
   start_app
}


function exec() {

CURRENT_DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
cd $CURRENT_DIR

if [ -z "$DIRECTORY" ]
  then
    echo "Please specify the directory name of the example you want to run"
    echo "More information with the flag --help"
    exit 1
fi

if [ -z "$ARG" ]
  then
    echo "Please specify the action"
    echo "More information with the flag --help"
    exit 1
fi


export CORS_EXAMPLE_CONTEXT=$DIRECTORY


 if [ "$ARG" == "restart" ]
     then
         restart_app
     elif [ "$ARG" == "start" ]
     then
         start_app
     elif [ "$ARG" == "stop" ]
     then
         stop_app
 fi

}

# Fail when not enough params
[[ $# -lt 4 ]] && usage && exit 1


PARAMS=""
while (( "$#" )); do
  case "$1" in
    -h|--help)
      usage
      shift
      ;;
    -n|--name)
      DIRECTORY=$2
      shift 2
      ;;
    -a|--action)
      ARG=$2
      shift 2
      ;;
    --) # end argument parsing
      shift
      break
      ;;
    -*|--*=) # unsupported flags
      echo "Error: Unsupported flag $1" >&2
      exit 1
      ;;
    *) # preserve positional arguments
      PARAMS="$PARAMS $1"
      shift
      ;;
  esac
done


# set positional arguments in their proper place
eval set -- "$PARAMS"
exec




