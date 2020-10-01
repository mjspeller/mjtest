
go mod init github.com/mjspeller/mjtest
go build ./modules/mjshelper/
#go build Test_terraform_module_test.go
dep init
dep ensure
go get -t -v


# github.com/mjspeller/mjtest/test
package github.com/mjspeller/mjtest/test (test)
	imports ./mjstest/modules/mjshelper: local import "./mjstest/modules/mjshelper" in non-local package
