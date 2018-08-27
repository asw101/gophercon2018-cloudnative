# README

https://github.com/spf13/cobra

go get -u github.com/spf13/cobra/cobra

cobra init sudoku

cd $HOME/go/src/soduku

cobra add subscriber

cd dir-outside-gopath
cp -r $HOME/go/src/sudoku/ .

go mod init sudoku

export CONNECTION_STRING="..."
export TOPIC_NAME="..."
export SUBSCRIPTION_NAME="..."

. ENV.sh

go run main.go subscriber
