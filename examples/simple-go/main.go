package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
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

		bucketResourceInput := sp.Bucket.ApplyT(func(x *s3.Bucket) pulumi.Resource {
			return x
		}).(pulumi.ResourceInput)

		_, err = s3.NewBucketObject(
			ctx,
			"go.sum",
			&s3.BucketObjectArgs{
				Bucket: sp.Bucket.ApplyT(func(x *s3.Bucket) pulumi.IDOutput {
					return x.ID()
				}),
				Source: pulumi.NewFileAsset("go.sum"),
			},
			pulumi.ParentInput(bucketResourceInput),
		)

		if err != nil {
			return err
		}

		ctx.Export("bucket", sp.Bucket)
		ctx.Export("url", sp.WebsiteUrl)
		return nil
	})
}
