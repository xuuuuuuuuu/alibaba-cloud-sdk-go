package cloudwf

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

func (client *Client) ProfileConsume(request *ProfileConsumeRequest) (response *ProfileConsumeResponse, err error) {
	response = CreateProfileConsumeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ProfileConsumeWithChan(request *ProfileConsumeRequest) (<-chan *ProfileConsumeResponse, <-chan error) {
	responseChan := make(chan *ProfileConsumeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ProfileConsume(request)
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

func (client *Client) ProfileConsumeWithCallback(request *ProfileConsumeRequest, callback func(response *ProfileConsumeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ProfileConsumeResponse
		var err error
		defer close(result)
		response, err = client.ProfileConsume(request)
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

type ProfileConsumeRequest struct {
	*requests.RpcRequest
	Gsid      string `position:"Query" name:"Gsid"`
	DataType  string `position:"Query" name:"DataType"`
	EndDate   string `position:"Query" name:"EndDate"`
	BeginDate string `position:"Query" name:"BeginDate"`
}

type ProfileConsumeResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

func CreateProfileConsumeRequest() (request *ProfileConsumeRequest) {
	request = &ProfileConsumeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "ProfileConsume", "", "")
	return
}

func CreateProfileConsumeResponse() (response *ProfileConsumeResponse) {
	response = &ProfileConsumeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}