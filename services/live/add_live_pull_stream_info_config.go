package live

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

func (client *Client) AddLivePullStreamInfoConfig(request *AddLivePullStreamInfoConfigRequest) (response *AddLivePullStreamInfoConfigResponse, err error) {
	response = CreateAddLivePullStreamInfoConfigResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AddLivePullStreamInfoConfigWithChan(request *AddLivePullStreamInfoConfigRequest) (<-chan *AddLivePullStreamInfoConfigResponse, <-chan error) {
	responseChan := make(chan *AddLivePullStreamInfoConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddLivePullStreamInfoConfig(request)
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

func (client *Client) AddLivePullStreamInfoConfigWithCallback(request *AddLivePullStreamInfoConfigRequest, callback func(response *AddLivePullStreamInfoConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddLivePullStreamInfoConfigResponse
		var err error
		defer close(result)
		response, err = client.AddLivePullStreamInfoConfig(request)
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

type AddLivePullStreamInfoConfigRequest struct {
	*requests.RpcRequest
	EndTime       string `position:"Query" name:"EndTime"`
	StreamName    string `position:"Query" name:"StreamName"`
	StartTime     string `position:"Query" name:"StartTime"`
	DomainName    string `position:"Query" name:"DomainName"`
	SourceUrl     string `position:"Query" name:"SourceUrl"`
	AppName       string `position:"Query" name:"AppName"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

type AddLivePullStreamInfoConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateAddLivePullStreamInfoConfigRequest() (request *AddLivePullStreamInfoConfigRequest) {
	request = &AddLivePullStreamInfoConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "AddLivePullStreamInfoConfig", "", "")
	return
}

func CreateAddLivePullStreamInfoConfigResponse() (response *AddLivePullStreamInfoConfigResponse) {
	response = &AddLivePullStreamInfoConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}