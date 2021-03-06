package functions

import (
	"github.com/IBM/ibm-cos-sdk-go/service/s3"
	"github.com/IBM/ibm-cos-sdk-go/service/s3/s3iface"
	"github.com/IBM/ibmcloud-cos-cli/config/fields"
	"github.com/IBM/ibmcloud-cos-cli/config/flags"
	"github.com/IBM/ibmcloud-cos-cli/errors"
	"github.com/IBM/ibmcloud-cos-cli/utils"
	"github.com/urfave/cli"
)

// MultipartAbort will end a current active multipart upload instance
// Parameter:
//   	CLI Context Application
// Returns:
//  	Error = zero or non-zero
func MultipartAbort(c *cli.Context) (err error) {
	// check the number of arguments
	if c.NArg() > 0 {
		err = &errors.CommandError{
			CLIContext: c,
			Cause:      errors.InvalidNArg,
		}
		return
	}

	// Load COS Context
	var cosContext *utils.CosContext
	if cosContext, err = GetCosContext(c); err != nil {
		return
	}

	// Initialize AbortMultipartUploadInput
	input := new(s3.AbortMultipartUploadInput)

	// Required parameters and no optional parameter for AbortMultipartUpload
	mandatory := map[string]string{
		fields.Bucket:   flags.Bucket,
		fields.Key:      flags.Key,
		fields.UploadID: flags.UploadID,
	}

	// No optional parameters for AbortMultipartUpload
	options := map[string]string{}

	// Check through user inputs for validation
	if err = MapToSDKInput(c, input, mandatory, options); err != nil {
		return
	}

	// Setting client to do the call
	var client s3iface.S3API
	if client, err = cosContext.GetClient(c.String(flags.Region)); err != nil {
		return
	}

	// AbortMultipartUpload Op
	var output *s3.AbortMultipartUploadOutput
	if output, err = client.AbortMultipartUpload(input); err != nil {
		return
	}

	// Display either in JSON or text
	err = cosContext.GetDisplay(c.Bool(flags.JSON)).Display(input, output, nil)

	// Return
	return
}
