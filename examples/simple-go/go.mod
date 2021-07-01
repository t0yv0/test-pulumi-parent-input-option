module hello-go

go 1.14

require (
	github.com/pulumi/pulumi-aws/sdk/v4 v4.0.0
	github.com/pulumi/pulumi-xyz/sdk v0.0.1
	github.com/pulumi/pulumi/sdk/v3 v3.6.1-0.20210701180729-76754321a532
)

replace github.com/pulumi/pulumi-xyz/sdk => ../../sdk
