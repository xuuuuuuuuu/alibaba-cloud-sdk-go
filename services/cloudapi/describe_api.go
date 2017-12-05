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

func (client *Client) DescribeApi(request *DescribeApiRequest) (response *DescribeApiResponse, err error) {
	response = CreateDescribeApiResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeApiWithChan(request *DescribeApiRequest) (<-chan *DescribeApiResponse, <-chan error) {
	responseChan := make(chan *DescribeApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApi(request)
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

func (client *Client) DescribeApiWithCallback(request *DescribeApiRequest, callback func(response *DescribeApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApiResponse
		var err error
		defer close(result)
		response, err = client.DescribeApi(request)
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

type DescribeApiRequest struct {
	*requests.RpcRequest
	GroupId string `position:"Query" name:"GroupId"`
	ApiId   string `position:"Query" name:"ApiId"`
}

type DescribeApiResponse struct {
	*responses.BaseResponse
	RequestId            string `json:"RequestId" xml:"RequestId"`
	RegionId             string `json:"RegionId" xml:"RegionId"`
	ApiId                string `json:"ApiId" xml:"ApiId"`
	ApiName              string `json:"ApiName" xml:"ApiName"`
	GroupId              string `json:"GroupId" xml:"GroupId"`
	GroupName            string `json:"GroupName" xml:"GroupName"`
	Visibility           string `json:"Visibility" xml:"Visibility"`
	AuthType             string `json:"AuthType" xml:"AuthType"`
	ResultType           string `json:"ResultType" xml:"ResultType"`
	ResultSample         string `json:"ResultSample" xml:"ResultSample"`
	FailResultSample     string `json:"FailResultSample" xml:"FailResultSample"`
	CreatedTime          string `json:"CreatedTime" xml:"CreatedTime"`
	ModifiedTime         string `json:"ModifiedTime" xml:"ModifiedTime"`
	Description          string `json:"Description" xml:"Description"`
	Mock                 string `json:"Mock" xml:"Mock"`
	MockResult           string `json:"MockResult" xml:"MockResult"`
	AllowSignatureMethod string `json:"AllowSignatureMethod" xml:"AllowSignatureMethod"`
	RequestConfig        struct {
		RequestProtocol     string `json:"RequestProtocol" xml:"RequestProtocol"`
		RequestHttpMethod   string `json:"RequestHttpMethod" xml:"RequestHttpMethod"`
		RequestPath         string `json:"RequestPath" xml:"RequestPath"`
		BodyFormat          string `json:"BodyFormat" xml:"BodyFormat"`
		PostBodyDescription string `json:"PostBodyDescription" xml:"PostBodyDescription"`
		RequestMode         string `json:"RequestMode" xml:"RequestMode"`
	} `json:"RequestConfig" xml:"RequestConfig"`
	ServiceConfig struct {
		ServiceProtocol     string `json:"ServiceProtocol" xml:"ServiceProtocol"`
		ServiceAddress      string `json:"ServiceAddress" xml:"ServiceAddress"`
		ServiceHttpMethod   string `json:"ServiceHttpMethod" xml:"ServiceHttpMethod"`
		ServicePath         string `json:"ServicePath" xml:"ServicePath"`
		ServiceTimeout      int    `json:"ServiceTimeout" xml:"ServiceTimeout"`
		ContentTypeCatagory string `json:"ContentTypeCatagory" xml:"ContentTypeCatagory"`
		ContentTypeValue    string `json:"ContentTypeValue" xml:"ContentTypeValue"`
		Mock                string `json:"Mock" xml:"Mock"`
		MockResult          string `json:"MockResult" xml:"MockResult"`
		ServiceVpcEnable    string `json:"ServiceVpcEnable" xml:"ServiceVpcEnable"`
		VpcConfig           struct {
			Name       string `json:"Name" xml:"Name"`
			VpcId      string `json:"VpcId" xml:"VpcId"`
			InstanceId string `json:"InstanceId" xml:"InstanceId"`
			Port       int    `json:"Port" xml:"Port"`
		} `json:"VpcConfig" xml:"VpcConfig"`
	} `json:"ServiceConfig" xml:"ServiceConfig"`
	OpenIdConnectConfig struct {
		OpenIdApiType    string `json:"OpenIdApiType" xml:"OpenIdApiType"`
		IdTokenParamName string `json:"IdTokenParamName" xml:"IdTokenParamName"`
		PublicKeyId      string `json:"PublicKeyId" xml:"PublicKeyId"`
		PublicKey        string `json:"PublicKey" xml:"PublicKey"`
	} `json:"OpenIdConnectConfig" xml:"OpenIdConnectConfig"`
	ErrorCodeSamples []struct {
		Code        string `json:"Code" xml:"Code"`
		Message     string `json:"Message" xml:"Message"`
		Description string `json:"Description" xml:"Description"`
	} `json:"ErrorCodeSamples" xml:"ErrorCodeSamples"`
	ResultDescriptions []struct {
		Id          string `json:"Id" xml:"Id"`
		Pid         string `json:"Pid" xml:"Pid"`
		HasChild    bool   `json:"HasChild" xml:"HasChild"`
		Key         string `json:"Key" xml:"Key"`
		Name        string `json:"Name" xml:"Name"`
		Mandatory   bool   `json:"Mandatory" xml:"Mandatory"`
		Type        string `json:"Type" xml:"Type"`
		Description string `json:"Description" xml:"Description"`
	} `json:"ResultDescriptions" xml:"ResultDescriptions"`
	SystemParameters []struct {
		ParameterName        string `json:"ParameterName" xml:"ParameterName"`
		ServiceParameterName string `json:"ServiceParameterName" xml:"ServiceParameterName"`
		Location             string `json:"Location" xml:"Location"`
		DemoValue            string `json:"DemoValue" xml:"DemoValue"`
		Description          string `json:"Description" xml:"Description"`
	} `json:"SystemParameters" xml:"SystemParameters"`
	CustomSystemParameters []struct {
		ParameterName        string `json:"ParameterName" xml:"ParameterName"`
		ServiceParameterName string `json:"ServiceParameterName" xml:"ServiceParameterName"`
		Location             string `json:"Location" xml:"Location"`
		DemoValue            string `json:"DemoValue" xml:"DemoValue"`
		Description          string `json:"Description" xml:"Description"`
	} `json:"CustomSystemParameters" xml:"CustomSystemParameters"`
	ConstantParameters []struct {
		ServiceParameterName string `json:"ServiceParameterName" xml:"ServiceParameterName"`
		ConstantValue        string `json:"ConstantValue" xml:"ConstantValue"`
		Location             string `json:"Location" xml:"Location"`
		Description          string `json:"Description" xml:"Description"`
	} `json:"ConstantParameters" xml:"ConstantParameters"`
	RequestParameters []struct {
		ApiParameterName  string `json:"ApiParameterName" xml:"ApiParameterName"`
		Location          string `json:"Location" xml:"Location"`
		ParameterType     string `json:"ParameterType" xml:"ParameterType"`
		Required          string `json:"Required" xml:"Required"`
		DefaultValue      string `json:"DefaultValue" xml:"DefaultValue"`
		DemoValue         string `json:"DemoValue" xml:"DemoValue"`
		MaxValue          int64  `json:"MaxValue" xml:"MaxValue"`
		MinValue          int64  `json:"MinValue" xml:"MinValue"`
		MaxLength         int64  `json:"MaxLength" xml:"MaxLength"`
		MinLength         int64  `json:"MinLength" xml:"MinLength"`
		RegularExpression string `json:"RegularExpression" xml:"RegularExpression"`
		JsonScheme        string `json:"JsonScheme" xml:"JsonScheme"`
		EnumValue         string `json:"EnumValue" xml:"EnumValue"`
		DocShow           string `json:"DocShow" xml:"DocShow"`
		DocOrder          int    `json:"DocOrder" xml:"DocOrder"`
		Description       string `json:"Description" xml:"Description"`
	} `json:"RequestParameters" xml:"RequestParameters"`
	ServiceParameters []struct {
		ServiceParameterName string `json:"ServiceParameterName" xml:"ServiceParameterName"`
		Location             string `json:"Location" xml:"Location"`
		ParameterType        string `json:"ParameterType" xml:"ParameterType"`
	} `json:"ServiceParameters" xml:"ServiceParameters"`
	ServiceParametersMap []struct {
		ServiceParameterName string `json:"ServiceParameterName" xml:"ServiceParameterName"`
		RequestParameterName string `json:"RequestParameterName" xml:"RequestParameterName"`
	} `json:"ServiceParametersMap" xml:"ServiceParametersMap"`
	DeployedInfos []struct {
		StageName        string `json:"StageName" xml:"StageName"`
		EffectiveVersion string `json:"EffectiveVersion" xml:"EffectiveVersion"`
		DeployedStatus   string `json:"DeployedStatus" xml:"DeployedStatus"`
	} `json:"DeployedInfos" xml:"DeployedInfos"`
}

func CreateDescribeApiRequest() (request *DescribeApiRequest) {
	request = &DescribeApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApi", "", "")
	return
}

func CreateDescribeApiResponse() (response *DescribeApiResponse) {
	response = &DescribeApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}