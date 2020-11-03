module github.com/dovetail-lab/fabric-client/activity/fabrequest

go 1.14

replace github.com/project-flogo/flow => github.com/yxuco/flow v1.1.1

replace github.com/project-flogo/core => github.com/yxuco/core v1.1.1

replace go.uber.org/multierr => go.uber.org/multierr v1.6.0

require (
	github.com/dovetail-lab/fabric-client/common v1.2.1
	github.com/pkg/errors v0.9.1
	github.com/project-flogo/core v1.1.0
	github.com/stretchr/testify v1.6.1
	go.uber.org/multierr v1.6.0 // indirect
)
