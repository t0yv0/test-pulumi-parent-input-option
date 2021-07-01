module hello-go

go 1.14

require (
	github.com/pulumi/pulumi/sdk/v3 v3.6.0
	github.com/pulumi/pulumi-xyz/sdk v0.0.1
)

replace github.com/pulumi/pulumi-xyz/sdk => ../../sdk
