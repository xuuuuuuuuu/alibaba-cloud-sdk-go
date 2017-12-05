package cloudapi

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

func (client *Client) DescribeApiLatencyData(request *DescribeApiLatencyDataRequest) (response *DescribeApiLatencyDataResponse, err error) {
	response = CreateDescribeApiLatencyDataResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeApiLatencyDataWithChan(request *DescribeApiLatencyDataRequest) (<-chan *DescribeApiLatencyDataResponse, <-chan error) {
	responseChan := make(chan *DescribeApiLatencyDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApiLatencyData(request)
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

func (client *Client) DescribeApiLatencyDataWithCallback(request *DescribeApiLatencyDataRequest, callback func(response *DescribeApiLatencyDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApiLatencyDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeApiLatencyData(request)
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

type DescribeApiLatencyDataRequest struct {
	*requests.RpcRequest
	EndTime   string `position:"Query" name:"EndTime"`
	StartTime string `position:"Query" name:"StartTime"`
	GroupId   string `position:"Query" name:"GroupId"`
	ApiId     string `position:"Query" name:"ApiId"`
}

type DescribeApiLatencyDataResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	CallLatencys []struct {
		ItemTime  string `json:"ItemTime" xml:"ItemTime"`
		ItemValue string `json:"ItemValue" xml:"ItemValue"`
	} `json:"CallLatencys" xml:"CallLatencys"`
}

func CreateDescribeApiLatencyDataRequest() (request *DescribeApiLatencyDataRequest) {
	request = &DescribeApiLatencyDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiLatencyData", "", "")
	return
}

func CreateDescribeApiLatencyDataResponse() (response *DescribeApiLatencyDataResponse) {
	response = &DescribeApiLatencyDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}