module github.com/dovetail-lab/fabric-client/trigger/eventlistener

go 1.14

replace github.com/project-flogo/core => github.com/yxuco/core v1.1.1

replace go.uber.org/multierr => go.uber.org/multierr v1.6.0

require (
	github.com/dovetail-lab/fabric-client/common v1.2.1
	github.com/golang/protobuf v1.4.3
	github.com/hyperledger/fabric-protos-go v0.0.0-20200707132912-fee30f3ccd23
	github.com/hyperledger/fabric-sdk-go v1.0.0-beta3.0.20201022122854-2f93a3201bb4
	github.com/pkg/errors v0.9.1
	github.com/project-flogo/core v1.1.0
)
