// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation is not supported by directory buckets.
//
// Retrieves the policy status for an Amazon S3 bucket, indicating whether the
// bucket is public. In order to use this operation, you must have the
// s3:GetBucketPolicyStatus permission. For more information about Amazon S3
// permissions, see [Specifying Permissions in a Policy].
//
// For more information about when Amazon S3 considers a bucket public, see [The Meaning of "Public"].
//
// The following operations are related to GetBucketPolicyStatus :
//
// [Using Amazon S3 Block Public Access]
//
// [GetPublicAccessBlock]
//
// [PutPublicAccessBlock]
//
// [DeletePublicAccessBlock]
//
// [GetPublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetPublicAccessBlock.html
// [PutPublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutPublicAccessBlock.html
// [DeletePublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeletePublicAccessBlock.html
// [Using Amazon S3 Block Public Access]: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html
// [Specifying Permissions in a Policy]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html
// [The Meaning of "Public"]: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status
func (c *Client) GetBucketPolicyStatus(ctx context.Context, params *GetBucketPolicyStatusInput, optFns ...func(*Options)) (*GetBucketPolicyStatusOutput, error) {
	if params == nil {
		params = &GetBucketPolicyStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketPolicyStatus", params, optFns, c.addOperationGetBucketPolicyStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketPolicyStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketPolicyStatusInput struct {

	// The name of the Amazon S3 bucket whose policy status you want to retrieve.
	//
	// This member is required.
	Bucket *string

	// The account ID of the expected bucket owner. If the account ID that you provide
	// does not match the actual owner of the bucket, the request fails with the HTTP
	// status code 403 Forbidden (access denied).
	ExpectedBucketOwner *string

	noSmithyDocumentSerde
}

func (in *GetBucketPolicyStatusInput) bindEndpointParams(p *EndpointParameters) {
	p.Bucket = in.Bucket
	p.UseS3ExpressControlEndpoint = ptr.Bool(true)
}

type GetBucketPolicyStatusOutput struct {

	// The policy status for the specified bucket.
	PolicyStatus *types.PolicyStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBucketPolicyStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketPolicyStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketPolicyStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetBucketPolicyStatus"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addPutBucketContextMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addOpGetBucketPolicyStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketPolicyStatus(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addGetBucketPolicyStatusUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSerializeImmutableHostnameBucketMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func (v *GetBucketPolicyStatusInput) bucket() (string, bool) {
	if v.Bucket == nil {
		return "", false
	}
	return *v.Bucket, true
}

func newServiceMetadataMiddleware_opGetBucketPolicyStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetBucketPolicyStatus",
	}
}

// getGetBucketPolicyStatusBucketMember returns a pointer to string denoting a
// provided bucket member valueand a boolean indicating if the input has a modeled
// bucket name,
func getGetBucketPolicyStatusBucketMember(input interface{}) (*string, bool) {
	in := input.(*GetBucketPolicyStatusInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addGetBucketPolicyStatusUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getGetBucketPolicyStatusBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}
