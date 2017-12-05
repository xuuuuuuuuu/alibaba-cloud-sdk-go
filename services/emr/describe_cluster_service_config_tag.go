package emr

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) DescribeClusterServiceConfigTag(request *DescribeClusterServiceConfigTagRequest) (response *DescribeClusterServiceConfigTagResponse, err error) {
	response = CreateDescribeClusterServiceConfigTagResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeClusterServiceConfigTagWithChan(request *DescribeClusterServiceConfigTagRequest) (<-chan *DescribeClusterServiceConfigTagResponse, <-chan error) {
	responseChan := make(chan *DescribeClusterServiceConfigTagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeClusterServiceConfigTag(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) DescribeClusterServiceConfigTagWithCallback(request *DescribeClusterServiceConfigTagRequest, callback func(response *DescribeClusterServiceConfigTagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClusterServiceConfigTagResponse
		var err error
		defer close(result)
		response, err = client.DescribeClusterServiceConfigTag(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type DescribeClusterServiceConfigTagRequest struct {
	*requests.RpcRequest
	ConfigTag       string `position:"Query" name:"ConfigTag"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ResourceOwnerId string `position:"Query" name:"ResourceOwnerId"`
}

type DescribeClusterServiceConfigTagResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	ConfigTagList []struct {
		Tag       string `json:"Tag" xml:"Tag"`
		TagDesc   string `json:"TagDesc" xml:"TagDesc"`
		ValueList []struct {
			Value     string `json:"Value" xml:"Value"`
			ValueDesc string `json:"ValueDesc" xml:"ValueDesc"`
		} `json:"ValueList" xml:"ValueList"`
	} `json:"ConfigTagList" xml:"ConfigTagList"`
}

func CreateDescribeClusterServiceConfigTagRequest() (request *DescribeClusterServiceConfigTagRequest) {
	request = &DescribeClusterServiceConfigTagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DescribeClusterServiceConfigTag", "", "")
	return
}

func CreateDescribeClusterServiceConfigTagResponse() (response *DescribeClusterServiceConfigTagResponse) {
	response = &DescribeClusterServiceConfigTagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}