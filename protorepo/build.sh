
set -e

CURRENT_BRANCH=${CIRCLE_BRANCH-"branch-not-available"}

rm -rf ./golang
rm -rf ./python
mkdir golang

# Helper for adding a directory to the stack and echoing the result
function enterDir {
  echo "Entering $1"
  pushd $1 > /dev/null
}

# Helper for popping a directory off the stack and echoing the result
function leaveDir {
  echo "Leaving `pwd`"
  popd > /dev/null
}

# Iterates through all of the languages listed in the services .protolangs file
# and compiles them individually
function buildProtoForTypes {
  target=${1%/}
  currentDir="$1"

  if [ -f .protolangs ]; then
    for line in $(< .protolangs)
      do
        buildLange $line $currentDir
    done
  fi
}

function buildLange {
  targetLang="$1"
  currentDir="$2"

  case $targetLang in
    go)
      mkdir ../golang/$currentDir
      protoc -I /usr/local/include \
        -I ${GOPATH}/src \
        -I ./ \
        -I ../golang/$currentDir \
        --go_out=plugins=grpc:../golang/$currentDir \
        --gorm_out=plugins=grpc,engine=mysql:../golang/$currentDir \
        --proto_path=. \
        *.proto 
      ;;
    # python) python3 -m grpc_tools.protoc -I ../python --python_out=../python --grpc_python_out=../python *.proto;;
  esac
}


# Enters the directory and starts the build / compile process for the services
# protobufs
function buildDir {
  currentDir="$1"
  echo "Building directory \"$currentDir\""

  enterDir $currentDir

  buildProtoForTypes $currentDir

  leaveDir
}

# Finds all directories in the repository and iterates through them calling the
# compile process for each one
function buildAll {
  echo "Buidling service's protocol buffers"
  # mkdir -p $REPOPATH
  for d in */; do
    buildDir $d;
  done
}

buildAll