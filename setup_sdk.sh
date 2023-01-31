#!/bin/bash

set -e
testDir="go-sdk-test"

GREEN="\033[0;32m"
NC="\033[0m"

if [ ! $PLIVO_API_PROD_HOST ] || [ ! $PLIVO_API_DEV_HOST ] || [ ! $PLIVO_AUTH_ID ] || [ ! $PLIVO_AUTH_TOKEN ]; then
    echo "Environment variables not properly set! Please refer to Local Development section in README!"
    exit 126
fi

cd /usr/src/app

echo "Setting plivo-api endpoint to dev..."
find /usr/src/app/*.go -type f -exec sed -i "s/$PLIVO_API_PROD_HOST/$PLIVO_API_DEV_HOST/g" {} \;

if [ ! -d $testDir ]; then
    echo "Creating test dir..."
    mkdir -p $testDir
fi

if [ ! -f $testDir/test.go ]; then
    echo "Creating test file..."
    cd $testDir
    echo -e "package main\n" > test.go
    echo -e "import (" >> test.go
    echo -e "\t\"fmt\"" >> test.go
    echo -e "\t\"github.com/plivo/plivo-go\"" >> test.go
    echo -e ")\n" >> test.go
    echo "func main() {" >> test.go
	echo -e "\t_, err := plivo.NewClient(\"\", \"\", &plivo.ClientOptions{})" >> test.go
	echo -e "\tif err != nil {" >> test.go
	echo -e "\t\tfmt.Printf(\"Error: %s\", err.Error())" >> test.go
	echo -e "\t}" >> test.go
    echo -e "}" >> test.go
    cd -
fi

if [ ! -f $testDir/go.mod ]; then
    echo "Setting up test package..."
    cd $testDir
    go mod init SDKtest
    go mod edit -replace github.com/plivo/plivo-go=/usr/src/app
    go mod tidy
else 
    echo "package has been setup at $testDir"
fi

echo -e "\n\nSDK setup complete!"
echo "To test your changes:"
echo -e "\t1. Add your test code in <path_to_cloned_sdk>/$testDir/test.go on host (or /usr/src/app/$testDir/test.go in the container)"
echo -e "\t2. Run a terminal in the container using: $GREEN docker exec -it $HOSTNAME /bin/bash$NC"
echo -e "\t3. Navigate to the test directory: $GREEN cd /usr/src/app/$testDir$NC"
echo -e "\t4. Run your test file: $GREEN go run test.go$NC"
echo -e "\t5. For running unit tests, run on host: $GREEN make test CONTAINER=$HOSTNAME$NC"

# To keep the container running post setup
/bin/bash