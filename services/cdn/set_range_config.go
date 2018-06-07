package cdn

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

// SetRangeConfig invokes the cdn.SetRangeConfig API synchronously
// api document: https://help.aliyun.com/api/cdn/setrangeconfig.html
func (client *Client) SetRangeConfig(request *SetRangeConfigRequest) (response *SetRangeConfigResponse, err error) {
	response = CreateSetRangeConfigResponse()
	err = client.DoAction(request, response)
	return
}

// SetRangeConfigWithChan invokes the cdn.SetRangeConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/setrangeconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetRangeConfigWithChan(request *SetRangeConfigRequest) (<-chan *SetRangeConfigResponse, <-chan error) {
	responseChan := make(chan *SetRangeConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetRangeConfig(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// SetRangeConfigWithCallback invokes the cdn.SetRangeConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/setrangeconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetRangeConfigWithCallback(request *SetRangeConfigRequest, callback func(response *SetRangeConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetRangeConfigResponse
		var err error
		defer close(result)
		response, err = client.SetRangeConfig(request)
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

// SetRangeConfigRequest is the request struct for api SetRangeConfig
type SetRangeConfigRequest struct {
	*requests.RpcRequest
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	Enable        string           `position:"Query" name:"Enable"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// SetRangeConfigResponse is the response struct for api SetRangeConfig
type SetRangeConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetRangeConfigRequest creates a request to invoke SetRangeConfig API
func CreateSetRangeConfigRequest() (request *SetRangeConfigRequest) {
	request = &SetRangeConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "SetRangeConfig", "", "")
	return
}

// CreateSetRangeConfigResponse creates a response to parse from SetRangeConfig response
func CreateSetRangeConfigResponse() (response *SetRangeConfigResponse) {
	response = &SetRangeConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
