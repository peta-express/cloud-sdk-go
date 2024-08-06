// +-------------------------------------------------------------------------
// | Copyright (C) 2024 PEG Tech, Inc.
// +-------------------------------------------------------------------------
// | Licensed under the Apache License, Version 2.0 (the "License");
// | you may not use this work except in compliance with the License.
// | You may obtain a copy of the License in the LICENSE file, or at:
// |
// | http://www.apache.org/licenses/LICENSE-2.0
// |
// | Unless required by applicable law or agreed to in writing, software
// | distributed under the License is distributed on an "AS IS" BASIS,
// | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// | See the License for the specific language governing permissions and
// | limitations under the License.
// +-------------------------------------------------------------------------

// Package service provides PetaExpress Service API (API Version 2013-08-30)
package service

import (
	"github.com/peta-express/cloud-sdk-go/config"
	"github.com/peta-express/cloud-sdk-go/logger"
	"github.com/peta-express/cloud-sdk-go/request"
	"github.com/peta-express/cloud-sdk-go/request/data"
)

// PetaExpressService: PetaExpress provides a platform which can make the delivery of computing resources more simple, efficient and reliable, even more environmental.
type PetaExpressService struct {
	Config     *config.Config
	Properties *PetaExpressServiceProperties
}

type PetaExpressServiceProperties struct {
}

func Init(c *config.Config) (*PetaExpressService, error) {
	properties := &PetaExpressServiceProperties{}
	logger.SetLevel(c.LogLevel)
	return &PetaExpressService{Config: c, Properties: properties}, nil
}

// Documentation URL: https://docs.petaexpress.com/api/zone/describe_zones.html
func (s *PetaExpressService) DescribeZones(i *DescribeZonesInput) (*DescribeZonesOutput, error) {
	if i == nil {
		i = &DescribeZonesInput{}
	}
	o := &data.Operation{
		Config:        s.Config,
		Properties:    s.Properties,
		APIName:       "DescribeZones",
		RequestMethod: "GET",
	}

	x := &DescribeZonesOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type DescribeZonesInput struct {
	Status []*string `json:"status" name:"status" location:"params"`
	Zones  []*string `json:"zones" name:"zones" location:"params"`
}

func (v *DescribeZonesInput) Validate() error {

	return nil
}

type DescribeZonesOutput struct {
	Message    *string `json:"message" name:"message"`
	Action     *string `json:"action" name:"action" location:"elements"`
	RetCode    *int    `json:"ret_code" name:"ret_code" location:"elements"`
	TotalCount *int    `json:"total_count" name:"total_count" location:"elements"`
	ZoneSet    []*Zone `json:"zone_set" name:"zone_set" location:"elements"`
}
