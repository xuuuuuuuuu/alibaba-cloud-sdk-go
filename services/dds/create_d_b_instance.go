package dds

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

func (client *Client) CreateDBInstance(request *CreateDBInstanceRequest) (response *CreateDBInstanceResponse, err error) {
	response = CreateCreateDBInstanceResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateDBInstanceWithChan(request *CreateDBInstanceRequest) (<-chan *CreateDBInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateDBInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDBInstance(request)
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

func (client *Client) CreateDBInstanceWithCallback(request *CreateDBInstanceRequest, callback func(response *CreateDBInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDBInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateDBInstance(request)
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

type CreateDBInstanceRequest struct {
	*requests.RpcRequest
	AccountPassword       string `position:"Query" name:"AccountPassword"`
	ZoneId                string `position:"Query" name:"ZoneId"`
	DBInstanceClass       string `position:"Query" name:"DBInstanceClass"`
	DBInstanceStorage     string `position:"Query" name:"DBInstanceStorage"`
	NetworkType           string `position:"Query" name:"NetworkType"`
	BackupId              string `position:"Query" name:"BackupId"`
	DBInstanceDescription string `position:"Query" name:"DBInstanceDescription"`
	RestoreTime           string `position:"Query" name:"RestoreTime"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	BusinessInfo          string `position:"Query" name:"BusinessInfo"`
	ResourceOwnerId       string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	VSwitchId             string `position:"Query" name:"VSwitchId"`
	CouponNo              string `position:"Query" name:"CouponNo"`
	ClientToken           string `position:"Query" name:"ClientToken"`
	SrcDBInstanceId       string `position:"Query" name:"SrcDBInstanceId"`
	Engine                string `position:"Query" name:"Engine"`
	OwnerId               string `position:"Query" name:"OwnerId"`
	Period                string `position:"Query" name:"Period"`
	SecurityIPList        string `position:"Query" name:"SecurityIPList"`
	ChargeType            string `position:"Query" name:"ChargeType"`
	VpcId                 string `position:"Query" name:"VpcId"`
	SecurityToken         string `position:"Query" name:"SecurityToken"`
	EngineVersion         string `position:"Query" name:"EngineVersion"`
	StorageEngine         string `position:"Query" name:"StorageEngine"`
}

type CreateDBInstanceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	OrderId      string `json:"OrderId" xml:"OrderId"`
	DBInstanceId string `json:"DBInstanceId" xml:"DBInstanceId"`
}

func CreateCreateDBInstanceRequest() (request *CreateDBInstanceRequest) {
	request = &CreateDBInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "CreateDBInstance", "", "")
	return
}

func CreateCreateDBInstanceResponse() (response *CreateDBInstanceResponse) {
	response = &CreateDBInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}