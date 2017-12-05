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

func (client *Client) CreateShardingDBInstance(request *CreateShardingDBInstanceRequest) (response *CreateShardingDBInstanceResponse, err error) {
	response = CreateCreateShardingDBInstanceResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateShardingDBInstanceWithChan(request *CreateShardingDBInstanceRequest) (<-chan *CreateShardingDBInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateShardingDBInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateShardingDBInstance(request)
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

func (client *Client) CreateShardingDBInstanceWithCallback(request *CreateShardingDBInstanceRequest, callback func(response *CreateShardingDBInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateShardingDBInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateShardingDBInstance(request)
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

type CreateShardingDBInstanceRequest struct {
	*requests.RpcRequest
	ConfigServer          *[]CreateShardingDBInstanceConfigServer `position:"Query" name:"ConfigServer"  type:"Repeated"`
	AccountPassword       string                                  `position:"Query" name:"AccountPassword"`
	ZoneId                string                                  `position:"Query" name:"ZoneId"`
	NetworkType           string                                  `position:"Query" name:"NetworkType"`
	DBInstanceDescription string                                  `position:"Query" name:"DBInstanceDescription"`
	ResourceOwnerAccount  string                                  `position:"Query" name:"ResourceOwnerAccount"`
	ReplicaSet            *[]CreateShardingDBInstanceReplicaSet   `position:"Query" name:"ReplicaSet"  type:"Repeated"`
	RestoreTime           string                                  `position:"Query" name:"RestoreTime"`
	ResourceOwnerId       string                                  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount          string                                  `position:"Query" name:"OwnerAccount"`
	VSwitchId             string                                  `position:"Query" name:"VSwitchId"`
	ClientToken           string                                  `position:"Query" name:"ClientToken"`
	SrcDBInstanceId       string                                  `position:"Query" name:"SrcDBInstanceId"`
	Engine                string                                  `position:"Query" name:"Engine"`
	OwnerId               string                                  `position:"Query" name:"OwnerId"`
	Period                string                                  `position:"Query" name:"Period"`
	SecurityIPList        string                                  `position:"Query" name:"SecurityIPList"`
	ChargeType            string                                  `position:"Query" name:"ChargeType"`
	VpcId                 string                                  `position:"Query" name:"VpcId"`
	SecurityToken         string                                  `position:"Query" name:"SecurityToken"`
	Mongos                *[]CreateShardingDBInstanceMongos       `position:"Query" name:"Mongos"  type:"Repeated"`
	EngineVersion         string                                  `position:"Query" name:"EngineVersion"`
	StorageEngine         string                                  `position:"Query" name:"StorageEngine"`
}

type CreateShardingDBInstanceConfigServer struct {
	Class   string `name:"Class"`
	Storage string `name:"Storage"`
}
type CreateShardingDBInstanceReplicaSet struct {
	Class   string `name:"Class"`
	Storage string `name:"Storage"`
}
type CreateShardingDBInstanceMongos struct {
	Class string `name:"Class"`
}

type CreateShardingDBInstanceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	OrderId      string `json:"OrderId" xml:"OrderId"`
	DBInstanceId string `json:"DBInstanceId" xml:"DBInstanceId"`
}

func CreateCreateShardingDBInstanceRequest() (request *CreateShardingDBInstanceRequest) {
	request = &CreateShardingDBInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "CreateShardingDBInstance", "", "")
	return
}

func CreateCreateShardingDBInstanceResponse() (response *CreateShardingDBInstanceResponse) {
	response = &CreateShardingDBInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}