// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

func StatusVPCStateV2(ctx context.Context, conn *ec2.Client, id string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		output, err := FindVPCByIDV2(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, string(output.State), nil
	}
}

func StatusVPCIPv6CIDRBlockAssociationStateV2(ctx context.Context, conn *ec2.Client, id string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		output, _, err := FindVPCIPv6CIDRBlockAssociationByIDV2(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output.Ipv6CidrBlockState, string(output.Ipv6CidrBlockState.State), nil
	}
}

func StatusVPCAttributeValueV2(ctx context.Context, conn *ec2.Client, id string, attribute string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		attributeValue, err := FindVPCAttributeV2(ctx, conn, id, attribute)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return attributeValue, strconv.FormatBool(attributeValue), nil
	}
}
