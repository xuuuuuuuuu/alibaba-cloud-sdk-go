package market

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

func (client *Client) DescribeProduct(request *DescribeProductRequest) (response *DescribeProductResponse, err error) {
	response = CreateDescribeProductResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeProductWithChan(request *DescribeProductRequest) (<-chan *DescribeProductResponse, <-chan error) {
	responseChan := make(chan *DescribeProductResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeProduct(request)
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

func (client *Client) DescribeProductWithCallback(request *DescribeProductRequest, callback func(response *DescribeProductResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeProductResponse
		var err error
		defer close(result)
		response, err = client.DescribeProduct(request)
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

type DescribeProductRequest struct {
	*requests.RpcRequest
	QueryDraft string `position:"Query" name:"QueryDraft"`
	AliUid     string `position:"Query" name:"AliUid"`
	Code       string `position:"Query" name:"Code"`
}

type DescribeProductResponse struct {
	*responses.BaseResponse
	Code             string `json:"Code" xml:"Code"`
	Name             string `json:"Name" xml:"Name"`
	Type             string `json:"Type" xml:"Type"`
	PicUrl           string `json:"PicUrl" xml:"PicUrl"`
	Description      string `json:"Description" xml:"Description"`
	ShortDescription string `json:"ShortDescription" xml:"ShortDescription"`
	Status           string `json:"Status" xml:"Status"`
	AuditStatus      string `json:"AuditStatus" xml:"AuditStatus"`
	AuditFailMsg     string `json:"AuditFailMsg" xml:"AuditFailMsg"`
	AuditTime        int64  `json:"AuditTime" xml:"AuditTime"`
	GmtCreated       int64  `json:"GmtCreated" xml:"GmtCreated"`
	GmtModified      int64  `json:"GmtModified" xml:"GmtModified"`
	ShopInfo         struct {
		Id         int64    `json:"Id" xml:"Id"`
		Name       string   `json:"Name" xml:"Name"`
		Emails     string   `json:"Emails" xml:"Emails"`
		Telephones []string `json:"Telephones" xml:"Telephones"`
		WangWangs  []struct {
			UserName string `json:"UserName" xml:"UserName"`
			Remark   string `json:"Remark" xml:"Remark"`
		} `json:"WangWangs" xml:"WangWangs"`
	} `json:"ShopInfo" xml:"ShopInfo"`
	ProductSkus []struct {
		Name         string `json:"Name" xml:"Name"`
		Code         string `json:"Code" xml:"Code"`
		ChargeType   string `json:"ChargeType" xml:"ChargeType"`
		Constraints  string `json:"Constraints" xml:"Constraints"`
		Hidden       bool   `json:"Hidden" xml:"Hidden"`
		OrderPeriods []struct {
			Name       string `json:"Name" xml:"Name"`
			PeriodType string `json:"PeriodType" xml:"PeriodType"`
		} `json:"OrderPeriods" xml:"OrderPeriods"`
		Modules []struct {
			Id         string `json:"Id" xml:"Id"`
			Name       string `json:"Name" xml:"Name"`
			Code       string `json:"Code" xml:"Code"`
			Properties []struct {
				Name           string `json:"Name" xml:"Name"`
				Key            string `json:"Key" xml:"Key"`
				PropertyValues []struct {
					Value       string `json:"Value" xml:"Value"`
					DisplayName string `json:"DisplayName" xml:"DisplayName"`
					Type        string `json:"Type" xml:"Type"`
				} `json:"PropertyValues" xml:"PropertyValues"`
			} `json:"Properties" xml:"Properties"`
		} `json:"Modules" xml:"Modules"`
	} `json:"ProductSkus" xml:"ProductSkus"`
	ProductExtras []struct {
		Key    string `json:"Key" xml:"Key"`
		Values string `json:"Values" xml:"Values"`
		Label  string `json:"Label" xml:"Label"`
		Order  int    `json:"Order" xml:"Order"`
		Type   string `json:"Type" xml:"Type"`
	} `json:"ProductExtras" xml:"ProductExtras"`
}

func CreateDescribeProductRequest() (request *DescribeProductRequest) {
	request = &DescribeProductRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Market", "2015-11-01", "DescribeProduct", "", "")
	return
}

func CreateDescribeProductResponse() (response *DescribeProductResponse) {
	response = &DescribeProductResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}