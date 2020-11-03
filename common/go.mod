module github.com/dovetail-lab/fabric-client/common

go 1.14

replace github.com/project-flogo/core => github.com/yxuco/core v1.1.1

replace github.com/project-flogo/flow => github.com/yxuco/flow v1.1.1

require (
	github.com/hyperledger/fabric-sdk-go v1.0.0-beta3.0.20201022122854-2f93a3201bb4
	github.com/pkg/errors v0.9.1
	github.com/project-flogo/core v1.1.0
	github.com/stretchr/testify v1.6.1
	github.com/xeipuuv/gojsonschema v1.2.0
)
