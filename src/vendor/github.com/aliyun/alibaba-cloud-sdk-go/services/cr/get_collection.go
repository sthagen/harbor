package cr

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

// GetCollection invokes the cr.GetCollection API synchronously
// api document: https://help.aliyun.com/api/cr/getcollection.html
func (client *Client) GetCollection(request *GetCollectionRequest) (response *GetCollectionResponse, err error) {
	response = CreateGetCollectionResponse()
	err = client.DoAction(request, response)
	return
}

// GetCollectionWithChan invokes the cr.GetCollection API asynchronously
// api document: https://help.aliyun.com/api/cr/getcollection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetCollectionWithChan(request *GetCollectionRequest) (<-chan *GetCollectionResponse, <-chan error) {
	responseChan := make(chan *GetCollectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCollection(request)
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

// GetCollectionWithCallback invokes the cr.GetCollection API asynchronously
// api document: https://help.aliyun.com/api/cr/getcollection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetCollectionWithCallback(request *GetCollectionRequest, callback func(response *GetCollectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCollectionResponse
		var err error
		defer close(result)
		response, err = client.GetCollection(request)
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

// GetCollectionRequest is the request struct for api GetCollection
type GetCollectionRequest struct {
	*requests.RoaRequest
	PageSize requests.Integer `position:"Query" name:"PageSize"`
	Page     requests.Integer `position:"Query" name:"Page"`
}

// GetCollectionResponse is the response struct for api GetCollection
type GetCollectionResponse struct {
	*responses.BaseResponse
}

// CreateGetCollectionRequest creates a request to invoke GetCollection API
func CreateGetCollectionRequest() (request *GetCollectionRequest) {
	request = &GetCollectionRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("cr", "2016-06-07", "GetCollection", "/collections", "cr", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetCollectionResponse creates a response to parse from GetCollection response
func CreateGetCollectionResponse() (response *GetCollectionResponse) {
	response = &GetCollectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
