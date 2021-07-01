package main

import (
	"github.com/pulumi/pulumi-xyz/sdk/go/xyz"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		sp, err := xyz.NewStaticPage(ctx, "page", &xyz.StaticPageArgs{
			IndexContent: pulumi.String("<html><body>Hello</body></html>"),
		})
		if err != nil {
			return err
		}

		ctx.Export("bucket", sp.Bucket)
		ctx.Export("url", sp.WebsiteUrl)
		return nil
	})
}
