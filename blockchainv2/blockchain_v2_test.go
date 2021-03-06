/**
 * (C) Copyright IBM Corp. 2020.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package blockchainv2_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/IBM-Blockchain/ibp-go-sdk/blockchainv2"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`BlockchainV2`, func() {
	var testServer *httptest.Server
    Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
				URL: "https://blockchainv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BLOCKCHAIN_URL": "https://blockchainv2/api",
				"BLOCKCHAIN_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := blockchainv2.NewBlockchainV2UsingExternalConfig(&blockchainv2.BlockchainV2Options{
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := blockchainv2.NewBlockchainV2UsingExternalConfig(&blockchainv2.BlockchainV2Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := blockchainv2.NewBlockchainV2UsingExternalConfig(&blockchainv2.BlockchainV2Options{
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BLOCKCHAIN_URL": "https://blockchainv2/api",
				"BLOCKCHAIN_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := blockchainv2.NewBlockchainV2UsingExternalConfig(&blockchainv2.BlockchainV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BLOCKCHAIN_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := blockchainv2.NewBlockchainV2UsingExternalConfig(&blockchainv2.BlockchainV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetComponent(getComponentOptions *GetComponentOptions) - Operation response error`, func() {
		getComponentPath := "/ak/api/v2/components/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getComponentPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["deployment_attrs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["parsed_certs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["cache"]).To(Equal([]string{"skip"}))

					Expect(req.URL.Query()["ca_attrs"]).To(Equal([]string{"included"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetComponent with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetComponentOptions model
				getComponentOptionsModel := new(blockchainv2.GetComponentOptions)
				getComponentOptionsModel.ID = core.StringPtr("testString")
				getComponentOptionsModel.DeploymentAttrs = core.StringPtr("included")
				getComponentOptionsModel.ParsedCerts = core.StringPtr("included")
				getComponentOptionsModel.Cache = core.StringPtr("skip")
				getComponentOptionsModel.CaAttrs = core.StringPtr("included")
				getComponentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetComponent(getComponentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetComponent(getComponentOptions *GetComponentOptions)`, func() {
		getComponentPath := "/ak/api/v2/components/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getComponentPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["deployment_attrs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["parsed_certs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["cache"]).To(Equal([]string{"skip"}))

					Expect(req.URL.Query()["ca_attrs"]).To(Equal([]string{"included"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "myca-2", "type": "fabric-ca", "display_name": "Example CA", "grpcwp_url": "https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "operations_url": "https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443", "msp_id": "Org1", "location": "ibmcloud", "ca_name": "ca", "admin_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="], "node_ou": {"enabled": true}, "ecert": {"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "cacert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}, "state_db": "couchdb", "timestamp": 1537262855753, "resources": {"ca": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "peer": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "orderer": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "statedb": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"ca": {"size": "4GiB", "class": "default"}, "peer": {"size": "4GiB", "class": "default"}, "orderer": {"size": "4GiB", "class": "default"}, "statedb": {"size": "4GiB", "class": "default"}}, "tags": ["Tags"], "tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "version": "Version", "zone": "Zone"}`)
				}))
			})
			It(`Invoke GetComponent successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetComponent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetComponentOptions model
				getComponentOptionsModel := new(blockchainv2.GetComponentOptions)
				getComponentOptionsModel.ID = core.StringPtr("testString")
				getComponentOptionsModel.DeploymentAttrs = core.StringPtr("included")
				getComponentOptionsModel.ParsedCerts = core.StringPtr("included")
				getComponentOptionsModel.Cache = core.StringPtr("skip")
				getComponentOptionsModel.CaAttrs = core.StringPtr("included")
 				getComponentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetComponent(getComponentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetComponent with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetComponentOptions model
				getComponentOptionsModel := new(blockchainv2.GetComponentOptions)
				getComponentOptionsModel.ID = core.StringPtr("testString")
				getComponentOptionsModel.DeploymentAttrs = core.StringPtr("included")
				getComponentOptionsModel.ParsedCerts = core.StringPtr("included")
				getComponentOptionsModel.Cache = core.StringPtr("skip")
				getComponentOptionsModel.CaAttrs = core.StringPtr("included")
				getComponentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetComponent(getComponentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetComponentOptions model with no property values
				getComponentOptionsModelNew := new(blockchainv2.GetComponentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetComponent(getComponentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RemoveComponent(removeComponentOptions *RemoveComponentOptions) - Operation response error`, func() {
		removeComponentPath := "/ak/api/v2/components/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(removeComponentPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RemoveComponent with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the RemoveComponentOptions model
				removeComponentOptionsModel := new(blockchainv2.RemoveComponentOptions)
				removeComponentOptionsModel.ID = core.StringPtr("testString")
				removeComponentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.RemoveComponent(removeComponentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`RemoveComponent(removeComponentOptions *RemoveComponentOptions)`, func() {
		removeComponentPath := "/ak/api/v2/components/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(removeComponentPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"message": "deleted", "type": "fabric-peer", "id": "component-1", "display_name": "My Peer"}`)
				}))
			})
			It(`Invoke RemoveComponent successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.RemoveComponent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RemoveComponentOptions model
				removeComponentOptionsModel := new(blockchainv2.RemoveComponentOptions)
				removeComponentOptionsModel.ID = core.StringPtr("testString")
 				removeComponentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.RemoveComponent(removeComponentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke RemoveComponent with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the RemoveComponentOptions model
				removeComponentOptionsModel := new(blockchainv2.RemoveComponentOptions)
				removeComponentOptionsModel.ID = core.StringPtr("testString")
				removeComponentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.RemoveComponent(removeComponentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RemoveComponentOptions model with no property values
				removeComponentOptionsModelNew := new(blockchainv2.RemoveComponentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.RemoveComponent(removeComponentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteComponent(deleteComponentOptions *DeleteComponentOptions) - Operation response error`, func() {
		deleteComponentPath := "/ak/api/v2/kubernetes/components/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteComponentPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteComponent with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteComponentOptions model
				deleteComponentOptionsModel := new(blockchainv2.DeleteComponentOptions)
				deleteComponentOptionsModel.ID = core.StringPtr("testString")
				deleteComponentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.DeleteComponent(deleteComponentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteComponent(deleteComponentOptions *DeleteComponentOptions)`, func() {
		deleteComponentPath := "/ak/api/v2/kubernetes/components/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteComponentPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"message": "deleted", "type": "fabric-peer", "id": "component-1", "display_name": "My Peer"}`)
				}))
			})
			It(`Invoke DeleteComponent successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteComponent(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteComponentOptions model
				deleteComponentOptionsModel := new(blockchainv2.DeleteComponentOptions)
				deleteComponentOptionsModel.ID = core.StringPtr("testString")
 				deleteComponentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteComponent(deleteComponentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteComponent with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteComponentOptions model
				deleteComponentOptionsModel := new(blockchainv2.DeleteComponentOptions)
				deleteComponentOptionsModel.ID = core.StringPtr("testString")
				deleteComponentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DeleteComponent(deleteComponentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteComponentOptions model with no property values
				deleteComponentOptionsModelNew := new(blockchainv2.DeleteComponentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.DeleteComponent(deleteComponentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCa(createCaOptions *CreateCaOptions) - Operation response error`, func() {
		createCaPath := "/ak/api/v2/kubernetes/components/fabric-ca"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createCaPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCa with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ConfigCADbTlsClient model
				configCaDbTlsClientModel := new(blockchainv2.ConfigCADbTlsClient)
				configCaDbTlsClientModel.Certfile = core.StringPtr("testString")
				configCaDbTlsClientModel.Keyfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateTlsClient model
				configCaIntermediateTlsClientModel := new(blockchainv2.ConfigCAIntermediateTlsClient)
				configCaIntermediateTlsClientModel.Certfile = core.StringPtr("testString")
				configCaIntermediateTlsClientModel.Keyfile = core.StringPtr("testString")

				// Construct an instance of the IdentityAttrs model
				identityAttrsModel := new(blockchainv2.IdentityAttrs)
				identityAttrsModel.HfRegistrarRoles = core.StringPtr("*")
				identityAttrsModel.HfRegistrarDelegateRoles = core.StringPtr("*")
				identityAttrsModel.HfRevoker = core.BoolPtr(true)
				identityAttrsModel.HfIntermediateCA = core.BoolPtr(true)
				identityAttrsModel.HfGenCRL = core.BoolPtr(true)
				identityAttrsModel.HfRegistrarAttributes = core.StringPtr("*")
				identityAttrsModel.HfAffiliationMgr = core.BoolPtr(true)

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv2.BccspPKCS11)
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv2.BccspSW)
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the ConfigCACfgIdentities model
				configCaCfgIdentitiesModel := new(blockchainv2.ConfigCACfgIdentities)
				configCaCfgIdentitiesModel.Passwordattempts = core.Float64Ptr(float64(10))
				configCaCfgIdentitiesModel.Allowremove = core.BoolPtr(false)

				// Construct an instance of the ConfigCACsrCa model
				configCaCsrCaModel := new(blockchainv2.ConfigCACsrCa)
				configCaCsrCaModel.Expiry = core.StringPtr("131400h")
				configCaCsrCaModel.Pathlength = core.Float64Ptr(float64(0))

				// Construct an instance of the ConfigCACsrKeyrequest model
				configCaCsrKeyrequestModel := new(blockchainv2.ConfigCACsrKeyrequest)
				configCaCsrKeyrequestModel.Algo = core.StringPtr("ecdsa")
				configCaCsrKeyrequestModel.Size = core.Float64Ptr(float64(256))

				// Construct an instance of the ConfigCACsrNamesItem model
				configCaCsrNamesItemModel := new(blockchainv2.ConfigCACsrNamesItem)
				configCaCsrNamesItemModel.C = core.StringPtr("US")
				configCaCsrNamesItemModel.ST = core.StringPtr("North Carolina")
				configCaCsrNamesItemModel.L = core.StringPtr("Raleigh")
				configCaCsrNamesItemModel.O = core.StringPtr("Hyperledger")
				configCaCsrNamesItemModel.OU = core.StringPtr("Fabric")

				// Construct an instance of the ConfigCADbTls model
				configCaDbTlsModel := new(blockchainv2.ConfigCADbTls)
				configCaDbTlsModel.Certfiles = []string{"testString"}
				configCaDbTlsModel.Client = configCaDbTlsClientModel
				configCaDbTlsModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the ConfigCAIntermediateEnrollment model
				configCaIntermediateEnrollmentModel := new(blockchainv2.ConfigCAIntermediateEnrollment)
				configCaIntermediateEnrollmentModel.Hosts = core.StringPtr("localhost")
				configCaIntermediateEnrollmentModel.Profile = core.StringPtr("testString")
				configCaIntermediateEnrollmentModel.Label = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateParentserver model
				configCaIntermediateParentserverModel := new(blockchainv2.ConfigCAIntermediateParentserver)
				configCaIntermediateParentserverModel.URL = core.StringPtr("testString")
				configCaIntermediateParentserverModel.Caname = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateTls model
				configCaIntermediateTlsModel := new(blockchainv2.ConfigCAIntermediateTls)
				configCaIntermediateTlsModel.Certfiles = []string{"testString"}
				configCaIntermediateTlsModel.Client = configCaIntermediateTlsClientModel

				// Construct an instance of the ConfigCARegistryIdentitiesItem model
				configCaRegistryIdentitiesItemModel := new(blockchainv2.ConfigCARegistryIdentitiesItem)
				configCaRegistryIdentitiesItemModel.Name = core.StringPtr("admin")
				configCaRegistryIdentitiesItemModel.Pass = core.StringPtr("password")
				configCaRegistryIdentitiesItemModel.Type = core.StringPtr("client")
				configCaRegistryIdentitiesItemModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryIdentitiesItemModel.Affiliation = core.StringPtr("testString")
				configCaRegistryIdentitiesItemModel.Attrs = identityAttrsModel

				// Construct an instance of the ConfigCATlsClientauth model
				configCaTlsClientauthModel := new(blockchainv2.ConfigCATlsClientauth)
				configCaTlsClientauthModel.Type = core.StringPtr("noclientcert")
				configCaTlsClientauthModel.Certfiles = []string{"testString"}

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv2.MetricsStatsd)
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv2.Bccsp)
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model

				// Construct an instance of the ConfigCAAffiliations model
				configCaAffiliationsModel := new(blockchainv2.ConfigCAAffiliations)
				configCaAffiliationsModel.Org1 = []string{"department1"}
				configCaAffiliationsModel.Org2 = []string{"department1"}
				configCaAffiliationsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ConfigCACa model
				configCaCaModel := new(blockchainv2.ConfigCACa)
				configCaCaModel.Keyfile = core.StringPtr("testString")
				configCaCaModel.Certfile = core.StringPtr("testString")
				configCaCaModel.Chainfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCACfg model
				configCaCfgModel := new(blockchainv2.ConfigCACfg)
				configCaCfgModel.Identities = configCaCfgIdentitiesModel

				// Construct an instance of the ConfigCACors model
				configCaCorsModel := new(blockchainv2.ConfigCACors)
				configCaCorsModel.Enabled = core.BoolPtr(true)
				configCaCorsModel.Origins = []string{"*"}

				// Construct an instance of the ConfigCACrl model
				configCaCrlModel := new(blockchainv2.ConfigCACrl)
				configCaCrlModel.Expiry = core.StringPtr("24h")

				// Construct an instance of the ConfigCACsr model
				configCaCsrModel := new(blockchainv2.ConfigCACsr)
				configCaCsrModel.Cn = core.StringPtr("ca")
				configCaCsrModel.Keyrequest = configCaCsrKeyrequestModel
				configCaCsrModel.Names = []blockchainv2.ConfigCACsrNamesItem{*configCaCsrNamesItemModel}
				configCaCsrModel.Hosts = []string{"localhost"}
				configCaCsrModel.Ca = configCaCsrCaModel

				// Construct an instance of the ConfigCADb model
				configCaDbModel := new(blockchainv2.ConfigCADb)
				configCaDbModel.Type = core.StringPtr("postgres")
				configCaDbModel.Datasource = core.StringPtr("host=fake.databases.appdomain.cloud port=31941 user=ibm_cloud password=password dbname=ibmclouddb sslmode=verify-full")
				configCaDbModel.Tls = configCaDbTlsModel

				// Construct an instance of the ConfigCAIdemix model
				configCaIdemixModel := new(blockchainv2.ConfigCAIdemix)
				configCaIdemixModel.Rhpoolsize = core.Float64Ptr(float64(100))
				configCaIdemixModel.Nonceexpiration = core.StringPtr("15s")
				configCaIdemixModel.Noncesweepinterval = core.StringPtr("15m")

				// Construct an instance of the ConfigCAIntermediate model
				configCaIntermediateModel := new(blockchainv2.ConfigCAIntermediate)
				configCaIntermediateModel.Parentserver = configCaIntermediateParentserverModel
				configCaIntermediateModel.Enrollment = configCaIntermediateEnrollmentModel
				configCaIntermediateModel.Tls = configCaIntermediateTlsModel

				// Construct an instance of the ConfigCARegistry model
				configCaRegistryModel := new(blockchainv2.ConfigCARegistry)
				configCaRegistryModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryModel.Identities = []blockchainv2.ConfigCARegistryIdentitiesItem{*configCaRegistryIdentitiesItemModel}

				// Construct an instance of the ConfigCATls model
				configCaTlsModel := new(blockchainv2.ConfigCATls)
				configCaTlsModel.Keyfile = core.StringPtr("testString")
				configCaTlsModel.Certfile = core.StringPtr("testString")
				configCaTlsModel.Clientauth = configCaTlsClientauthModel

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv2.Metrics)
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv2.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv2.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ConfigCACreate model
				configCaCreateModel := new(blockchainv2.ConfigCACreate)
				configCaCreateModel.Cors = configCaCorsModel
				configCaCreateModel.Debug = core.BoolPtr(false)
				configCaCreateModel.Crlsizelimit = core.Float64Ptr(float64(512000))
				configCaCreateModel.Tls = configCaTlsModel
				configCaCreateModel.Ca = configCaCaModel
				configCaCreateModel.Crl = configCaCrlModel
				configCaCreateModel.Registry = configCaRegistryModel
				configCaCreateModel.Db = configCaDbModel
				configCaCreateModel.Affiliations = configCaAffiliationsModel
				configCaCreateModel.Csr = configCaCsrModel
				configCaCreateModel.Idemix = configCaIdemixModel
				configCaCreateModel.BCCSP = bccspModel
				configCaCreateModel.Intermediate = configCaIntermediateModel
				configCaCreateModel.Cfg = configCaCfgModel
				configCaCreateModel.Metrics = metricsModel

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv2.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the StorageObject model
				storageObjectModel := new(blockchainv2.StorageObject)
				storageObjectModel.Size = core.StringPtr("4GiB")
				storageObjectModel.Class = core.StringPtr("default")

				// Construct an instance of the CreateCaBodyConfigOverride model
				createCaBodyConfigOverrideModel := new(blockchainv2.CreateCaBodyConfigOverride)
				createCaBodyConfigOverrideModel.Ca = configCaCreateModel
				createCaBodyConfigOverrideModel.Tlsca = configCaCreateModel

				// Construct an instance of the CreateCaBodyResources model
				createCaBodyResourcesModel := new(blockchainv2.CreateCaBodyResources)
				createCaBodyResourcesModel.Ca = resourceObjectModel

				// Construct an instance of the CreateCaBodyStorage model
				createCaBodyStorageModel := new(blockchainv2.CreateCaBodyStorage)
				createCaBodyStorageModel.Ca = storageObjectModel

				// Construct an instance of the Hsm model
				hsmModel := new(blockchainv2.Hsm)
				hsmModel.Pkcs11endpoint = core.StringPtr("tcp://example.com:666")

				// Construct an instance of the CreateCaOptions model
				createCaOptionsModel := new(blockchainv2.CreateCaOptions)
				createCaOptionsModel.DisplayName = core.StringPtr("My CA")
				createCaOptionsModel.ConfigOverride = createCaBodyConfigOverrideModel
				createCaOptionsModel.Resources = createCaBodyResourcesModel
				createCaOptionsModel.Storage = createCaBodyStorageModel
				createCaOptionsModel.Zone = core.StringPtr("testString")
				createCaOptionsModel.Replicas = core.Float64Ptr(float64(1))
				createCaOptionsModel.Tags = []string{"testString"}
				createCaOptionsModel.Hsm = hsmModel
				createCaOptionsModel.Region = core.StringPtr("testString")
				createCaOptionsModel.Version = core.StringPtr("1.4.6-1")
				createCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateCa(createCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateCa(createCaOptions *CreateCaOptions)`, func() {
		createCaPath := "/ak/api/v2/kubernetes/components/fabric-ca"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createCaPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "component-1", "dep_component_id": "admin", "ca_name": "ca", "display_name": "My CA", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "operations_url": "https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443", "config_override": {"anyKey": "anyValue"}, "location": "ibmcloud", "timestamp": 1537262855753, "resources": {"ca": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"ca": {"size": "4GiB", "class": "default"}}, "tags": ["Tags"], "tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "version": "1.4.6-1", "zone": "Zone"}`)
				}))
			})
			It(`Invoke CreateCa successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateCa(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ConfigCADbTlsClient model
				configCaDbTlsClientModel := new(blockchainv2.ConfigCADbTlsClient)
				configCaDbTlsClientModel.Certfile = core.StringPtr("testString")
				configCaDbTlsClientModel.Keyfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateTlsClient model
				configCaIntermediateTlsClientModel := new(blockchainv2.ConfigCAIntermediateTlsClient)
				configCaIntermediateTlsClientModel.Certfile = core.StringPtr("testString")
				configCaIntermediateTlsClientModel.Keyfile = core.StringPtr("testString")

				// Construct an instance of the IdentityAttrs model
				identityAttrsModel := new(blockchainv2.IdentityAttrs)
				identityAttrsModel.HfRegistrarRoles = core.StringPtr("*")
				identityAttrsModel.HfRegistrarDelegateRoles = core.StringPtr("*")
				identityAttrsModel.HfRevoker = core.BoolPtr(true)
				identityAttrsModel.HfIntermediateCA = core.BoolPtr(true)
				identityAttrsModel.HfGenCRL = core.BoolPtr(true)
				identityAttrsModel.HfRegistrarAttributes = core.StringPtr("*")
				identityAttrsModel.HfAffiliationMgr = core.BoolPtr(true)

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv2.BccspPKCS11)
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv2.BccspSW)
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the ConfigCACfgIdentities model
				configCaCfgIdentitiesModel := new(blockchainv2.ConfigCACfgIdentities)
				configCaCfgIdentitiesModel.Passwordattempts = core.Float64Ptr(float64(10))
				configCaCfgIdentitiesModel.Allowremove = core.BoolPtr(false)

				// Construct an instance of the ConfigCACsrCa model
				configCaCsrCaModel := new(blockchainv2.ConfigCACsrCa)
				configCaCsrCaModel.Expiry = core.StringPtr("131400h")
				configCaCsrCaModel.Pathlength = core.Float64Ptr(float64(0))

				// Construct an instance of the ConfigCACsrKeyrequest model
				configCaCsrKeyrequestModel := new(blockchainv2.ConfigCACsrKeyrequest)
				configCaCsrKeyrequestModel.Algo = core.StringPtr("ecdsa")
				configCaCsrKeyrequestModel.Size = core.Float64Ptr(float64(256))

				// Construct an instance of the ConfigCACsrNamesItem model
				configCaCsrNamesItemModel := new(blockchainv2.ConfigCACsrNamesItem)
				configCaCsrNamesItemModel.C = core.StringPtr("US")
				configCaCsrNamesItemModel.ST = core.StringPtr("North Carolina")
				configCaCsrNamesItemModel.L = core.StringPtr("Raleigh")
				configCaCsrNamesItemModel.O = core.StringPtr("Hyperledger")
				configCaCsrNamesItemModel.OU = core.StringPtr("Fabric")

				// Construct an instance of the ConfigCADbTls model
				configCaDbTlsModel := new(blockchainv2.ConfigCADbTls)
				configCaDbTlsModel.Certfiles = []string{"testString"}
				configCaDbTlsModel.Client = configCaDbTlsClientModel
				configCaDbTlsModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the ConfigCAIntermediateEnrollment model
				configCaIntermediateEnrollmentModel := new(blockchainv2.ConfigCAIntermediateEnrollment)
				configCaIntermediateEnrollmentModel.Hosts = core.StringPtr("localhost")
				configCaIntermediateEnrollmentModel.Profile = core.StringPtr("testString")
				configCaIntermediateEnrollmentModel.Label = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateParentserver model
				configCaIntermediateParentserverModel := new(blockchainv2.ConfigCAIntermediateParentserver)
				configCaIntermediateParentserverModel.URL = core.StringPtr("testString")
				configCaIntermediateParentserverModel.Caname = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateTls model
				configCaIntermediateTlsModel := new(blockchainv2.ConfigCAIntermediateTls)
				configCaIntermediateTlsModel.Certfiles = []string{"testString"}
				configCaIntermediateTlsModel.Client = configCaIntermediateTlsClientModel

				// Construct an instance of the ConfigCARegistryIdentitiesItem model
				configCaRegistryIdentitiesItemModel := new(blockchainv2.ConfigCARegistryIdentitiesItem)
				configCaRegistryIdentitiesItemModel.Name = core.StringPtr("admin")
				configCaRegistryIdentitiesItemModel.Pass = core.StringPtr("password")
				configCaRegistryIdentitiesItemModel.Type = core.StringPtr("client")
				configCaRegistryIdentitiesItemModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryIdentitiesItemModel.Affiliation = core.StringPtr("testString")
				configCaRegistryIdentitiesItemModel.Attrs = identityAttrsModel

				// Construct an instance of the ConfigCATlsClientauth model
				configCaTlsClientauthModel := new(blockchainv2.ConfigCATlsClientauth)
				configCaTlsClientauthModel.Type = core.StringPtr("noclientcert")
				configCaTlsClientauthModel.Certfiles = []string{"testString"}

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv2.MetricsStatsd)
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv2.Bccsp)
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model

				// Construct an instance of the ConfigCAAffiliations model
				configCaAffiliationsModel := new(blockchainv2.ConfigCAAffiliations)
				configCaAffiliationsModel.Org1 = []string{"department1"}
				configCaAffiliationsModel.Org2 = []string{"department1"}
				configCaAffiliationsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ConfigCACa model
				configCaCaModel := new(blockchainv2.ConfigCACa)
				configCaCaModel.Keyfile = core.StringPtr("testString")
				configCaCaModel.Certfile = core.StringPtr("testString")
				configCaCaModel.Chainfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCACfg model
				configCaCfgModel := new(blockchainv2.ConfigCACfg)
				configCaCfgModel.Identities = configCaCfgIdentitiesModel

				// Construct an instance of the ConfigCACors model
				configCaCorsModel := new(blockchainv2.ConfigCACors)
				configCaCorsModel.Enabled = core.BoolPtr(true)
				configCaCorsModel.Origins = []string{"*"}

				// Construct an instance of the ConfigCACrl model
				configCaCrlModel := new(blockchainv2.ConfigCACrl)
				configCaCrlModel.Expiry = core.StringPtr("24h")

				// Construct an instance of the ConfigCACsr model
				configCaCsrModel := new(blockchainv2.ConfigCACsr)
				configCaCsrModel.Cn = core.StringPtr("ca")
				configCaCsrModel.Keyrequest = configCaCsrKeyrequestModel
				configCaCsrModel.Names = []blockchainv2.ConfigCACsrNamesItem{*configCaCsrNamesItemModel}
				configCaCsrModel.Hosts = []string{"localhost"}
				configCaCsrModel.Ca = configCaCsrCaModel

				// Construct an instance of the ConfigCADb model
				configCaDbModel := new(blockchainv2.ConfigCADb)
				configCaDbModel.Type = core.StringPtr("postgres")
				configCaDbModel.Datasource = core.StringPtr("host=fake.databases.appdomain.cloud port=31941 user=ibm_cloud password=password dbname=ibmclouddb sslmode=verify-full")
				configCaDbModel.Tls = configCaDbTlsModel

				// Construct an instance of the ConfigCAIdemix model
				configCaIdemixModel := new(blockchainv2.ConfigCAIdemix)
				configCaIdemixModel.Rhpoolsize = core.Float64Ptr(float64(100))
				configCaIdemixModel.Nonceexpiration = core.StringPtr("15s")
				configCaIdemixModel.Noncesweepinterval = core.StringPtr("15m")

				// Construct an instance of the ConfigCAIntermediate model
				configCaIntermediateModel := new(blockchainv2.ConfigCAIntermediate)
				configCaIntermediateModel.Parentserver = configCaIntermediateParentserverModel
				configCaIntermediateModel.Enrollment = configCaIntermediateEnrollmentModel
				configCaIntermediateModel.Tls = configCaIntermediateTlsModel

				// Construct an instance of the ConfigCARegistry model
				configCaRegistryModel := new(blockchainv2.ConfigCARegistry)
				configCaRegistryModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryModel.Identities = []blockchainv2.ConfigCARegistryIdentitiesItem{*configCaRegistryIdentitiesItemModel}

				// Construct an instance of the ConfigCATls model
				configCaTlsModel := new(blockchainv2.ConfigCATls)
				configCaTlsModel.Keyfile = core.StringPtr("testString")
				configCaTlsModel.Certfile = core.StringPtr("testString")
				configCaTlsModel.Clientauth = configCaTlsClientauthModel

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv2.Metrics)
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv2.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv2.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ConfigCACreate model
				configCaCreateModel := new(blockchainv2.ConfigCACreate)
				configCaCreateModel.Cors = configCaCorsModel
				configCaCreateModel.Debug = core.BoolPtr(false)
				configCaCreateModel.Crlsizelimit = core.Float64Ptr(float64(512000))
				configCaCreateModel.Tls = configCaTlsModel
				configCaCreateModel.Ca = configCaCaModel
				configCaCreateModel.Crl = configCaCrlModel
				configCaCreateModel.Registry = configCaRegistryModel
				configCaCreateModel.Db = configCaDbModel
				configCaCreateModel.Affiliations = configCaAffiliationsModel
				configCaCreateModel.Csr = configCaCsrModel
				configCaCreateModel.Idemix = configCaIdemixModel
				configCaCreateModel.BCCSP = bccspModel
				configCaCreateModel.Intermediate = configCaIntermediateModel
				configCaCreateModel.Cfg = configCaCfgModel
				configCaCreateModel.Metrics = metricsModel

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv2.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the StorageObject model
				storageObjectModel := new(blockchainv2.StorageObject)
				storageObjectModel.Size = core.StringPtr("4GiB")
				storageObjectModel.Class = core.StringPtr("default")

				// Construct an instance of the CreateCaBodyConfigOverride model
				createCaBodyConfigOverrideModel := new(blockchainv2.CreateCaBodyConfigOverride)
				createCaBodyConfigOverrideModel.Ca = configCaCreateModel
				createCaBodyConfigOverrideModel.Tlsca = configCaCreateModel

				// Construct an instance of the CreateCaBodyResources model
				createCaBodyResourcesModel := new(blockchainv2.CreateCaBodyResources)
				createCaBodyResourcesModel.Ca = resourceObjectModel

				// Construct an instance of the CreateCaBodyStorage model
				createCaBodyStorageModel := new(blockchainv2.CreateCaBodyStorage)
				createCaBodyStorageModel.Ca = storageObjectModel

				// Construct an instance of the Hsm model
				hsmModel := new(blockchainv2.Hsm)
				hsmModel.Pkcs11endpoint = core.StringPtr("tcp://example.com:666")

				// Construct an instance of the CreateCaOptions model
				createCaOptionsModel := new(blockchainv2.CreateCaOptions)
				createCaOptionsModel.DisplayName = core.StringPtr("My CA")
				createCaOptionsModel.ConfigOverride = createCaBodyConfigOverrideModel
				createCaOptionsModel.Resources = createCaBodyResourcesModel
				createCaOptionsModel.Storage = createCaBodyStorageModel
				createCaOptionsModel.Zone = core.StringPtr("testString")
				createCaOptionsModel.Replicas = core.Float64Ptr(float64(1))
				createCaOptionsModel.Tags = []string{"testString"}
				createCaOptionsModel.Hsm = hsmModel
				createCaOptionsModel.Region = core.StringPtr("testString")
				createCaOptionsModel.Version = core.StringPtr("1.4.6-1")
 				createCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateCa(createCaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateCa with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ConfigCADbTlsClient model
				configCaDbTlsClientModel := new(blockchainv2.ConfigCADbTlsClient)
				configCaDbTlsClientModel.Certfile = core.StringPtr("testString")
				configCaDbTlsClientModel.Keyfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateTlsClient model
				configCaIntermediateTlsClientModel := new(blockchainv2.ConfigCAIntermediateTlsClient)
				configCaIntermediateTlsClientModel.Certfile = core.StringPtr("testString")
				configCaIntermediateTlsClientModel.Keyfile = core.StringPtr("testString")

				// Construct an instance of the IdentityAttrs model
				identityAttrsModel := new(blockchainv2.IdentityAttrs)
				identityAttrsModel.HfRegistrarRoles = core.StringPtr("*")
				identityAttrsModel.HfRegistrarDelegateRoles = core.StringPtr("*")
				identityAttrsModel.HfRevoker = core.BoolPtr(true)
				identityAttrsModel.HfIntermediateCA = core.BoolPtr(true)
				identityAttrsModel.HfGenCRL = core.BoolPtr(true)
				identityAttrsModel.HfRegistrarAttributes = core.StringPtr("*")
				identityAttrsModel.HfAffiliationMgr = core.BoolPtr(true)

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv2.BccspPKCS11)
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv2.BccspSW)
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the ConfigCACfgIdentities model
				configCaCfgIdentitiesModel := new(blockchainv2.ConfigCACfgIdentities)
				configCaCfgIdentitiesModel.Passwordattempts = core.Float64Ptr(float64(10))
				configCaCfgIdentitiesModel.Allowremove = core.BoolPtr(false)

				// Construct an instance of the ConfigCACsrCa model
				configCaCsrCaModel := new(blockchainv2.ConfigCACsrCa)
				configCaCsrCaModel.Expiry = core.StringPtr("131400h")
				configCaCsrCaModel.Pathlength = core.Float64Ptr(float64(0))

				// Construct an instance of the ConfigCACsrKeyrequest model
				configCaCsrKeyrequestModel := new(blockchainv2.ConfigCACsrKeyrequest)
				configCaCsrKeyrequestModel.Algo = core.StringPtr("ecdsa")
				configCaCsrKeyrequestModel.Size = core.Float64Ptr(float64(256))

				// Construct an instance of the ConfigCACsrNamesItem model
				configCaCsrNamesItemModel := new(blockchainv2.ConfigCACsrNamesItem)
				configCaCsrNamesItemModel.C = core.StringPtr("US")
				configCaCsrNamesItemModel.ST = core.StringPtr("North Carolina")
				configCaCsrNamesItemModel.L = core.StringPtr("Raleigh")
				configCaCsrNamesItemModel.O = core.StringPtr("Hyperledger")
				configCaCsrNamesItemModel.OU = core.StringPtr("Fabric")

				// Construct an instance of the ConfigCADbTls model
				configCaDbTlsModel := new(blockchainv2.ConfigCADbTls)
				configCaDbTlsModel.Certfiles = []string{"testString"}
				configCaDbTlsModel.Client = configCaDbTlsClientModel
				configCaDbTlsModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the ConfigCAIntermediateEnrollment model
				configCaIntermediateEnrollmentModel := new(blockchainv2.ConfigCAIntermediateEnrollment)
				configCaIntermediateEnrollmentModel.Hosts = core.StringPtr("localhost")
				configCaIntermediateEnrollmentModel.Profile = core.StringPtr("testString")
				configCaIntermediateEnrollmentModel.Label = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateParentserver model
				configCaIntermediateParentserverModel := new(blockchainv2.ConfigCAIntermediateParentserver)
				configCaIntermediateParentserverModel.URL = core.StringPtr("testString")
				configCaIntermediateParentserverModel.Caname = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateTls model
				configCaIntermediateTlsModel := new(blockchainv2.ConfigCAIntermediateTls)
				configCaIntermediateTlsModel.Certfiles = []string{"testString"}
				configCaIntermediateTlsModel.Client = configCaIntermediateTlsClientModel

				// Construct an instance of the ConfigCARegistryIdentitiesItem model
				configCaRegistryIdentitiesItemModel := new(blockchainv2.ConfigCARegistryIdentitiesItem)
				configCaRegistryIdentitiesItemModel.Name = core.StringPtr("admin")
				configCaRegistryIdentitiesItemModel.Pass = core.StringPtr("password")
				configCaRegistryIdentitiesItemModel.Type = core.StringPtr("client")
				configCaRegistryIdentitiesItemModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryIdentitiesItemModel.Affiliation = core.StringPtr("testString")
				configCaRegistryIdentitiesItemModel.Attrs = identityAttrsModel

				// Construct an instance of the ConfigCATlsClientauth model
				configCaTlsClientauthModel := new(blockchainv2.ConfigCATlsClientauth)
				configCaTlsClientauthModel.Type = core.StringPtr("noclientcert")
				configCaTlsClientauthModel.Certfiles = []string{"testString"}

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv2.MetricsStatsd)
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv2.Bccsp)
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model

				// Construct an instance of the ConfigCAAffiliations model
				configCaAffiliationsModel := new(blockchainv2.ConfigCAAffiliations)
				configCaAffiliationsModel.Org1 = []string{"department1"}
				configCaAffiliationsModel.Org2 = []string{"department1"}
				configCaAffiliationsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ConfigCACa model
				configCaCaModel := new(blockchainv2.ConfigCACa)
				configCaCaModel.Keyfile = core.StringPtr("testString")
				configCaCaModel.Certfile = core.StringPtr("testString")
				configCaCaModel.Chainfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCACfg model
				configCaCfgModel := new(blockchainv2.ConfigCACfg)
				configCaCfgModel.Identities = configCaCfgIdentitiesModel

				// Construct an instance of the ConfigCACors model
				configCaCorsModel := new(blockchainv2.ConfigCACors)
				configCaCorsModel.Enabled = core.BoolPtr(true)
				configCaCorsModel.Origins = []string{"*"}

				// Construct an instance of the ConfigCACrl model
				configCaCrlModel := new(blockchainv2.ConfigCACrl)
				configCaCrlModel.Expiry = core.StringPtr("24h")

				// Construct an instance of the ConfigCACsr model
				configCaCsrModel := new(blockchainv2.ConfigCACsr)
				configCaCsrModel.Cn = core.StringPtr("ca")
				configCaCsrModel.Keyrequest = configCaCsrKeyrequestModel
				configCaCsrModel.Names = []blockchainv2.ConfigCACsrNamesItem{*configCaCsrNamesItemModel}
				configCaCsrModel.Hosts = []string{"localhost"}
				configCaCsrModel.Ca = configCaCsrCaModel

				// Construct an instance of the ConfigCADb model
				configCaDbModel := new(blockchainv2.ConfigCADb)
				configCaDbModel.Type = core.StringPtr("postgres")
				configCaDbModel.Datasource = core.StringPtr("host=fake.databases.appdomain.cloud port=31941 user=ibm_cloud password=password dbname=ibmclouddb sslmode=verify-full")
				configCaDbModel.Tls = configCaDbTlsModel

				// Construct an instance of the ConfigCAIdemix model
				configCaIdemixModel := new(blockchainv2.ConfigCAIdemix)
				configCaIdemixModel.Rhpoolsize = core.Float64Ptr(float64(100))
				configCaIdemixModel.Nonceexpiration = core.StringPtr("15s")
				configCaIdemixModel.Noncesweepinterval = core.StringPtr("15m")

				// Construct an instance of the ConfigCAIntermediate model
				configCaIntermediateModel := new(blockchainv2.ConfigCAIntermediate)
				configCaIntermediateModel.Parentserver = configCaIntermediateParentserverModel
				configCaIntermediateModel.Enrollment = configCaIntermediateEnrollmentModel
				configCaIntermediateModel.Tls = configCaIntermediateTlsModel

				// Construct an instance of the ConfigCARegistry model
				configCaRegistryModel := new(blockchainv2.ConfigCARegistry)
				configCaRegistryModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryModel.Identities = []blockchainv2.ConfigCARegistryIdentitiesItem{*configCaRegistryIdentitiesItemModel}

				// Construct an instance of the ConfigCATls model
				configCaTlsModel := new(blockchainv2.ConfigCATls)
				configCaTlsModel.Keyfile = core.StringPtr("testString")
				configCaTlsModel.Certfile = core.StringPtr("testString")
				configCaTlsModel.Clientauth = configCaTlsClientauthModel

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv2.Metrics)
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv2.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv2.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ConfigCACreate model
				configCaCreateModel := new(blockchainv2.ConfigCACreate)
				configCaCreateModel.Cors = configCaCorsModel
				configCaCreateModel.Debug = core.BoolPtr(false)
				configCaCreateModel.Crlsizelimit = core.Float64Ptr(float64(512000))
				configCaCreateModel.Tls = configCaTlsModel
				configCaCreateModel.Ca = configCaCaModel
				configCaCreateModel.Crl = configCaCrlModel
				configCaCreateModel.Registry = configCaRegistryModel
				configCaCreateModel.Db = configCaDbModel
				configCaCreateModel.Affiliations = configCaAffiliationsModel
				configCaCreateModel.Csr = configCaCsrModel
				configCaCreateModel.Idemix = configCaIdemixModel
				configCaCreateModel.BCCSP = bccspModel
				configCaCreateModel.Intermediate = configCaIntermediateModel
				configCaCreateModel.Cfg = configCaCfgModel
				configCaCreateModel.Metrics = metricsModel

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv2.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the StorageObject model
				storageObjectModel := new(blockchainv2.StorageObject)
				storageObjectModel.Size = core.StringPtr("4GiB")
				storageObjectModel.Class = core.StringPtr("default")

				// Construct an instance of the CreateCaBodyConfigOverride model
				createCaBodyConfigOverrideModel := new(blockchainv2.CreateCaBodyConfigOverride)
				createCaBodyConfigOverrideModel.Ca = configCaCreateModel
				createCaBodyConfigOverrideModel.Tlsca = configCaCreateModel

				// Construct an instance of the CreateCaBodyResources model
				createCaBodyResourcesModel := new(blockchainv2.CreateCaBodyResources)
				createCaBodyResourcesModel.Ca = resourceObjectModel

				// Construct an instance of the CreateCaBodyStorage model
				createCaBodyStorageModel := new(blockchainv2.CreateCaBodyStorage)
				createCaBodyStorageModel.Ca = storageObjectModel

				// Construct an instance of the Hsm model
				hsmModel := new(blockchainv2.Hsm)
				hsmModel.Pkcs11endpoint = core.StringPtr("tcp://example.com:666")

				// Construct an instance of the CreateCaOptions model
				createCaOptionsModel := new(blockchainv2.CreateCaOptions)
				createCaOptionsModel.DisplayName = core.StringPtr("My CA")
				createCaOptionsModel.ConfigOverride = createCaBodyConfigOverrideModel
				createCaOptionsModel.Resources = createCaBodyResourcesModel
				createCaOptionsModel.Storage = createCaBodyStorageModel
				createCaOptionsModel.Zone = core.StringPtr("testString")
				createCaOptionsModel.Replicas = core.Float64Ptr(float64(1))
				createCaOptionsModel.Tags = []string{"testString"}
				createCaOptionsModel.Hsm = hsmModel
				createCaOptionsModel.Region = core.StringPtr("testString")
				createCaOptionsModel.Version = core.StringPtr("1.4.6-1")
				createCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateCa(createCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateCaOptions model with no property values
				createCaOptionsModelNew := new(blockchainv2.CreateCaOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreateCa(createCaOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ImportCa(importCaOptions *ImportCaOptions) - Operation response error`, func() {
		importCaPath := "/ak/api/v2/components/fabric-ca"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(importCaPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ImportCa with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ImportCaOptions model
				importCaOptionsModel := new(blockchainv2.ImportCaOptions)
				importCaOptionsModel.DisplayName = core.StringPtr("Sample CA")
				importCaOptionsModel.ApiURL = core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:7054")
				importCaOptionsModel.CaName = core.StringPtr("org1CA")
				importCaOptionsModel.TlscaName = core.StringPtr("org1CA")
				importCaOptionsModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				importCaOptionsModel.Location = core.StringPtr("ibmcloud")
				importCaOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443")
				importCaOptionsModel.Tags = []string{"testString"}
				importCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ImportCa(importCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ImportCa(importCaOptions *ImportCaOptions)`, func() {
		importCaPath := "/ak/api/v2/components/fabric-ca"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(importCaPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "component-1", "dep_component_id": "admin", "ca_name": "ca", "display_name": "My CA", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "operations_url": "https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443", "config_override": {"anyKey": "anyValue"}, "location": "ibmcloud", "timestamp": 1537262855753, "resources": {"ca": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"ca": {"size": "4GiB", "class": "default"}}, "tags": ["Tags"], "tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "version": "1.4.6-1", "zone": "Zone"}`)
				}))
			})
			It(`Invoke ImportCa successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ImportCa(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ImportCaOptions model
				importCaOptionsModel := new(blockchainv2.ImportCaOptions)
				importCaOptionsModel.DisplayName = core.StringPtr("Sample CA")
				importCaOptionsModel.ApiURL = core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:7054")
				importCaOptionsModel.CaName = core.StringPtr("org1CA")
				importCaOptionsModel.TlscaName = core.StringPtr("org1CA")
				importCaOptionsModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				importCaOptionsModel.Location = core.StringPtr("ibmcloud")
				importCaOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443")
				importCaOptionsModel.Tags = []string{"testString"}
 				importCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ImportCa(importCaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ImportCa with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ImportCaOptions model
				importCaOptionsModel := new(blockchainv2.ImportCaOptions)
				importCaOptionsModel.DisplayName = core.StringPtr("Sample CA")
				importCaOptionsModel.ApiURL = core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:7054")
				importCaOptionsModel.CaName = core.StringPtr("org1CA")
				importCaOptionsModel.TlscaName = core.StringPtr("org1CA")
				importCaOptionsModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				importCaOptionsModel.Location = core.StringPtr("ibmcloud")
				importCaOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443")
				importCaOptionsModel.Tags = []string{"testString"}
				importCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ImportCa(importCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ImportCaOptions model with no property values
				importCaOptionsModelNew := new(blockchainv2.ImportCaOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ImportCa(importCaOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCa(updateCaOptions *UpdateCaOptions) - Operation response error`, func() {
		updateCaPath := "/ak/api/v2/kubernetes/components/fabric-ca/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateCaPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCa with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ConfigCADbTlsClient model
				configCaDbTlsClientModel := new(blockchainv2.ConfigCADbTlsClient)
				configCaDbTlsClientModel.Certfile = core.StringPtr("testString")
				configCaDbTlsClientModel.Keyfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateTlsClient model
				configCaIntermediateTlsClientModel := new(blockchainv2.ConfigCAIntermediateTlsClient)
				configCaIntermediateTlsClientModel.Certfile = core.StringPtr("testString")
				configCaIntermediateTlsClientModel.Keyfile = core.StringPtr("testString")

				// Construct an instance of the IdentityAttrs model
				identityAttrsModel := new(blockchainv2.IdentityAttrs)
				identityAttrsModel.HfRegistrarRoles = core.StringPtr("*")
				identityAttrsModel.HfRegistrarDelegateRoles = core.StringPtr("*")
				identityAttrsModel.HfRevoker = core.BoolPtr(true)
				identityAttrsModel.HfIntermediateCA = core.BoolPtr(true)
				identityAttrsModel.HfGenCRL = core.BoolPtr(true)
				identityAttrsModel.HfRegistrarAttributes = core.StringPtr("*")
				identityAttrsModel.HfAffiliationMgr = core.BoolPtr(true)

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv2.BccspPKCS11)
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv2.BccspSW)
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the ConfigCACfgIdentities model
				configCaCfgIdentitiesModel := new(blockchainv2.ConfigCACfgIdentities)
				configCaCfgIdentitiesModel.Passwordattempts = core.Float64Ptr(float64(10))
				configCaCfgIdentitiesModel.Allowremove = core.BoolPtr(false)

				// Construct an instance of the ConfigCACsrCa model
				configCaCsrCaModel := new(blockchainv2.ConfigCACsrCa)
				configCaCsrCaModel.Expiry = core.StringPtr("131400h")
				configCaCsrCaModel.Pathlength = core.Float64Ptr(float64(0))

				// Construct an instance of the ConfigCACsrKeyrequest model
				configCaCsrKeyrequestModel := new(blockchainv2.ConfigCACsrKeyrequest)
				configCaCsrKeyrequestModel.Algo = core.StringPtr("ecdsa")
				configCaCsrKeyrequestModel.Size = core.Float64Ptr(float64(256))

				// Construct an instance of the ConfigCACsrNamesItem model
				configCaCsrNamesItemModel := new(blockchainv2.ConfigCACsrNamesItem)
				configCaCsrNamesItemModel.C = core.StringPtr("US")
				configCaCsrNamesItemModel.ST = core.StringPtr("North Carolina")
				configCaCsrNamesItemModel.L = core.StringPtr("Raleigh")
				configCaCsrNamesItemModel.O = core.StringPtr("Hyperledger")
				configCaCsrNamesItemModel.OU = core.StringPtr("Fabric")

				// Construct an instance of the ConfigCADbTls model
				configCaDbTlsModel := new(blockchainv2.ConfigCADbTls)
				configCaDbTlsModel.Certfiles = []string{"testString"}
				configCaDbTlsModel.Client = configCaDbTlsClientModel
				configCaDbTlsModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the ConfigCAIntermediateEnrollment model
				configCaIntermediateEnrollmentModel := new(blockchainv2.ConfigCAIntermediateEnrollment)
				configCaIntermediateEnrollmentModel.Hosts = core.StringPtr("localhost")
				configCaIntermediateEnrollmentModel.Profile = core.StringPtr("testString")
				configCaIntermediateEnrollmentModel.Label = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateParentserver model
				configCaIntermediateParentserverModel := new(blockchainv2.ConfigCAIntermediateParentserver)
				configCaIntermediateParentserverModel.URL = core.StringPtr("testString")
				configCaIntermediateParentserverModel.Caname = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateTls model
				configCaIntermediateTlsModel := new(blockchainv2.ConfigCAIntermediateTls)
				configCaIntermediateTlsModel.Certfiles = []string{"testString"}
				configCaIntermediateTlsModel.Client = configCaIntermediateTlsClientModel

				// Construct an instance of the ConfigCARegistryIdentitiesItem model
				configCaRegistryIdentitiesItemModel := new(blockchainv2.ConfigCARegistryIdentitiesItem)
				configCaRegistryIdentitiesItemModel.Name = core.StringPtr("admin")
				configCaRegistryIdentitiesItemModel.Pass = core.StringPtr("password")
				configCaRegistryIdentitiesItemModel.Type = core.StringPtr("client")
				configCaRegistryIdentitiesItemModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryIdentitiesItemModel.Affiliation = core.StringPtr("testString")
				configCaRegistryIdentitiesItemModel.Attrs = identityAttrsModel

				// Construct an instance of the ConfigCATlsClientauth model
				configCaTlsClientauthModel := new(blockchainv2.ConfigCATlsClientauth)
				configCaTlsClientauthModel.Type = core.StringPtr("noclientcert")
				configCaTlsClientauthModel.Certfiles = []string{"testString"}

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv2.MetricsStatsd)
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv2.Bccsp)
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model

				// Construct an instance of the ConfigCAAffiliations model
				configCaAffiliationsModel := new(blockchainv2.ConfigCAAffiliations)
				configCaAffiliationsModel.Org1 = []string{"department1"}
				configCaAffiliationsModel.Org2 = []string{"department1"}
				configCaAffiliationsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ConfigCACa model
				configCaCaModel := new(blockchainv2.ConfigCACa)
				configCaCaModel.Keyfile = core.StringPtr("testString")
				configCaCaModel.Certfile = core.StringPtr("testString")
				configCaCaModel.Chainfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCACfg model
				configCaCfgModel := new(blockchainv2.ConfigCACfg)
				configCaCfgModel.Identities = configCaCfgIdentitiesModel

				// Construct an instance of the ConfigCACors model
				configCaCorsModel := new(blockchainv2.ConfigCACors)
				configCaCorsModel.Enabled = core.BoolPtr(true)
				configCaCorsModel.Origins = []string{"*"}

				// Construct an instance of the ConfigCACrl model
				configCaCrlModel := new(blockchainv2.ConfigCACrl)
				configCaCrlModel.Expiry = core.StringPtr("24h")

				// Construct an instance of the ConfigCACsr model
				configCaCsrModel := new(blockchainv2.ConfigCACsr)
				configCaCsrModel.Cn = core.StringPtr("ca")
				configCaCsrModel.Keyrequest = configCaCsrKeyrequestModel
				configCaCsrModel.Names = []blockchainv2.ConfigCACsrNamesItem{*configCaCsrNamesItemModel}
				configCaCsrModel.Hosts = []string{"localhost"}
				configCaCsrModel.Ca = configCaCsrCaModel

				// Construct an instance of the ConfigCADb model
				configCaDbModel := new(blockchainv2.ConfigCADb)
				configCaDbModel.Type = core.StringPtr("postgres")
				configCaDbModel.Datasource = core.StringPtr("host=fake.databases.appdomain.cloud port=31941 user=ibm_cloud password=password dbname=ibmclouddb sslmode=verify-full")
				configCaDbModel.Tls = configCaDbTlsModel

				// Construct an instance of the ConfigCAIdemix model
				configCaIdemixModel := new(blockchainv2.ConfigCAIdemix)
				configCaIdemixModel.Rhpoolsize = core.Float64Ptr(float64(100))
				configCaIdemixModel.Nonceexpiration = core.StringPtr("15s")
				configCaIdemixModel.Noncesweepinterval = core.StringPtr("15m")

				// Construct an instance of the ConfigCAIntermediate model
				configCaIntermediateModel := new(blockchainv2.ConfigCAIntermediate)
				configCaIntermediateModel.Parentserver = configCaIntermediateParentserverModel
				configCaIntermediateModel.Enrollment = configCaIntermediateEnrollmentModel
				configCaIntermediateModel.Tls = configCaIntermediateTlsModel

				// Construct an instance of the ConfigCARegistry model
				configCaRegistryModel := new(blockchainv2.ConfigCARegistry)
				configCaRegistryModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryModel.Identities = []blockchainv2.ConfigCARegistryIdentitiesItem{*configCaRegistryIdentitiesItemModel}

				// Construct an instance of the ConfigCATls model
				configCaTlsModel := new(blockchainv2.ConfigCATls)
				configCaTlsModel.Keyfile = core.StringPtr("testString")
				configCaTlsModel.Certfile = core.StringPtr("testString")
				configCaTlsModel.Clientauth = configCaTlsClientauthModel

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv2.Metrics)
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv2.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv2.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ConfigCAUpdate model
				configCaUpdateModel := new(blockchainv2.ConfigCAUpdate)
				configCaUpdateModel.Cors = configCaCorsModel
				configCaUpdateModel.Debug = core.BoolPtr(false)
				configCaUpdateModel.Crlsizelimit = core.Float64Ptr(float64(512000))
				configCaUpdateModel.Tls = configCaTlsModel
				configCaUpdateModel.Ca = configCaCaModel
				configCaUpdateModel.Crl = configCaCrlModel
				configCaUpdateModel.Registry = configCaRegistryModel
				configCaUpdateModel.Db = configCaDbModel
				configCaUpdateModel.Affiliations = configCaAffiliationsModel
				configCaUpdateModel.Csr = configCaCsrModel
				configCaUpdateModel.Idemix = configCaIdemixModel
				configCaUpdateModel.BCCSP = bccspModel
				configCaUpdateModel.Intermediate = configCaIntermediateModel
				configCaUpdateModel.Cfg = configCaCfgModel
				configCaUpdateModel.Metrics = metricsModel

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv2.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the UpdateCaBodyConfigOverride model
				updateCaBodyConfigOverrideModel := new(blockchainv2.UpdateCaBodyConfigOverride)
				updateCaBodyConfigOverrideModel.Ca = configCaUpdateModel

				// Construct an instance of the UpdateCaBodyResources model
				updateCaBodyResourcesModel := new(blockchainv2.UpdateCaBodyResources)
				updateCaBodyResourcesModel.Ca = resourceObjectModel

				// Construct an instance of the UpdateCaOptions model
				updateCaOptionsModel := new(blockchainv2.UpdateCaOptions)
				updateCaOptionsModel.ID = core.StringPtr("testString")
				updateCaOptionsModel.Resources = updateCaBodyResourcesModel
				updateCaOptionsModel.Zone = core.StringPtr("testString")
				updateCaOptionsModel.ConfigOverride = updateCaBodyConfigOverrideModel
				updateCaOptionsModel.Replicas = core.Float64Ptr(float64(1))
				updateCaOptionsModel.Version = core.StringPtr("1.4.6-1")
				updateCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateCa(updateCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateCa(updateCaOptions *UpdateCaOptions)`, func() {
		updateCaPath := "/ak/api/v2/kubernetes/components/fabric-ca/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateCaPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "component-1", "dep_component_id": "admin", "ca_name": "ca", "display_name": "My CA", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "operations_url": "https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443", "config_override": {"anyKey": "anyValue"}, "location": "ibmcloud", "timestamp": 1537262855753, "resources": {"ca": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"ca": {"size": "4GiB", "class": "default"}}, "tags": ["Tags"], "tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "version": "1.4.6-1", "zone": "Zone"}`)
				}))
			})
			It(`Invoke UpdateCa successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateCa(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ConfigCADbTlsClient model
				configCaDbTlsClientModel := new(blockchainv2.ConfigCADbTlsClient)
				configCaDbTlsClientModel.Certfile = core.StringPtr("testString")
				configCaDbTlsClientModel.Keyfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateTlsClient model
				configCaIntermediateTlsClientModel := new(blockchainv2.ConfigCAIntermediateTlsClient)
				configCaIntermediateTlsClientModel.Certfile = core.StringPtr("testString")
				configCaIntermediateTlsClientModel.Keyfile = core.StringPtr("testString")

				// Construct an instance of the IdentityAttrs model
				identityAttrsModel := new(blockchainv2.IdentityAttrs)
				identityAttrsModel.HfRegistrarRoles = core.StringPtr("*")
				identityAttrsModel.HfRegistrarDelegateRoles = core.StringPtr("*")
				identityAttrsModel.HfRevoker = core.BoolPtr(true)
				identityAttrsModel.HfIntermediateCA = core.BoolPtr(true)
				identityAttrsModel.HfGenCRL = core.BoolPtr(true)
				identityAttrsModel.HfRegistrarAttributes = core.StringPtr("*")
				identityAttrsModel.HfAffiliationMgr = core.BoolPtr(true)

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv2.BccspPKCS11)
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv2.BccspSW)
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the ConfigCACfgIdentities model
				configCaCfgIdentitiesModel := new(blockchainv2.ConfigCACfgIdentities)
				configCaCfgIdentitiesModel.Passwordattempts = core.Float64Ptr(float64(10))
				configCaCfgIdentitiesModel.Allowremove = core.BoolPtr(false)

				// Construct an instance of the ConfigCACsrCa model
				configCaCsrCaModel := new(blockchainv2.ConfigCACsrCa)
				configCaCsrCaModel.Expiry = core.StringPtr("131400h")
				configCaCsrCaModel.Pathlength = core.Float64Ptr(float64(0))

				// Construct an instance of the ConfigCACsrKeyrequest model
				configCaCsrKeyrequestModel := new(blockchainv2.ConfigCACsrKeyrequest)
				configCaCsrKeyrequestModel.Algo = core.StringPtr("ecdsa")
				configCaCsrKeyrequestModel.Size = core.Float64Ptr(float64(256))

				// Construct an instance of the ConfigCACsrNamesItem model
				configCaCsrNamesItemModel := new(blockchainv2.ConfigCACsrNamesItem)
				configCaCsrNamesItemModel.C = core.StringPtr("US")
				configCaCsrNamesItemModel.ST = core.StringPtr("North Carolina")
				configCaCsrNamesItemModel.L = core.StringPtr("Raleigh")
				configCaCsrNamesItemModel.O = core.StringPtr("Hyperledger")
				configCaCsrNamesItemModel.OU = core.StringPtr("Fabric")

				// Construct an instance of the ConfigCADbTls model
				configCaDbTlsModel := new(blockchainv2.ConfigCADbTls)
				configCaDbTlsModel.Certfiles = []string{"testString"}
				configCaDbTlsModel.Client = configCaDbTlsClientModel
				configCaDbTlsModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the ConfigCAIntermediateEnrollment model
				configCaIntermediateEnrollmentModel := new(blockchainv2.ConfigCAIntermediateEnrollment)
				configCaIntermediateEnrollmentModel.Hosts = core.StringPtr("localhost")
				configCaIntermediateEnrollmentModel.Profile = core.StringPtr("testString")
				configCaIntermediateEnrollmentModel.Label = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateParentserver model
				configCaIntermediateParentserverModel := new(blockchainv2.ConfigCAIntermediateParentserver)
				configCaIntermediateParentserverModel.URL = core.StringPtr("testString")
				configCaIntermediateParentserverModel.Caname = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateTls model
				configCaIntermediateTlsModel := new(blockchainv2.ConfigCAIntermediateTls)
				configCaIntermediateTlsModel.Certfiles = []string{"testString"}
				configCaIntermediateTlsModel.Client = configCaIntermediateTlsClientModel

				// Construct an instance of the ConfigCARegistryIdentitiesItem model
				configCaRegistryIdentitiesItemModel := new(blockchainv2.ConfigCARegistryIdentitiesItem)
				configCaRegistryIdentitiesItemModel.Name = core.StringPtr("admin")
				configCaRegistryIdentitiesItemModel.Pass = core.StringPtr("password")
				configCaRegistryIdentitiesItemModel.Type = core.StringPtr("client")
				configCaRegistryIdentitiesItemModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryIdentitiesItemModel.Affiliation = core.StringPtr("testString")
				configCaRegistryIdentitiesItemModel.Attrs = identityAttrsModel

				// Construct an instance of the ConfigCATlsClientauth model
				configCaTlsClientauthModel := new(blockchainv2.ConfigCATlsClientauth)
				configCaTlsClientauthModel.Type = core.StringPtr("noclientcert")
				configCaTlsClientauthModel.Certfiles = []string{"testString"}

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv2.MetricsStatsd)
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv2.Bccsp)
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model

				// Construct an instance of the ConfigCAAffiliations model
				configCaAffiliationsModel := new(blockchainv2.ConfigCAAffiliations)
				configCaAffiliationsModel.Org1 = []string{"department1"}
				configCaAffiliationsModel.Org2 = []string{"department1"}
				configCaAffiliationsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ConfigCACa model
				configCaCaModel := new(blockchainv2.ConfigCACa)
				configCaCaModel.Keyfile = core.StringPtr("testString")
				configCaCaModel.Certfile = core.StringPtr("testString")
				configCaCaModel.Chainfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCACfg model
				configCaCfgModel := new(blockchainv2.ConfigCACfg)
				configCaCfgModel.Identities = configCaCfgIdentitiesModel

				// Construct an instance of the ConfigCACors model
				configCaCorsModel := new(blockchainv2.ConfigCACors)
				configCaCorsModel.Enabled = core.BoolPtr(true)
				configCaCorsModel.Origins = []string{"*"}

				// Construct an instance of the ConfigCACrl model
				configCaCrlModel := new(blockchainv2.ConfigCACrl)
				configCaCrlModel.Expiry = core.StringPtr("24h")

				// Construct an instance of the ConfigCACsr model
				configCaCsrModel := new(blockchainv2.ConfigCACsr)
				configCaCsrModel.Cn = core.StringPtr("ca")
				configCaCsrModel.Keyrequest = configCaCsrKeyrequestModel
				configCaCsrModel.Names = []blockchainv2.ConfigCACsrNamesItem{*configCaCsrNamesItemModel}
				configCaCsrModel.Hosts = []string{"localhost"}
				configCaCsrModel.Ca = configCaCsrCaModel

				// Construct an instance of the ConfigCADb model
				configCaDbModel := new(blockchainv2.ConfigCADb)
				configCaDbModel.Type = core.StringPtr("postgres")
				configCaDbModel.Datasource = core.StringPtr("host=fake.databases.appdomain.cloud port=31941 user=ibm_cloud password=password dbname=ibmclouddb sslmode=verify-full")
				configCaDbModel.Tls = configCaDbTlsModel

				// Construct an instance of the ConfigCAIdemix model
				configCaIdemixModel := new(blockchainv2.ConfigCAIdemix)
				configCaIdemixModel.Rhpoolsize = core.Float64Ptr(float64(100))
				configCaIdemixModel.Nonceexpiration = core.StringPtr("15s")
				configCaIdemixModel.Noncesweepinterval = core.StringPtr("15m")

				// Construct an instance of the ConfigCAIntermediate model
				configCaIntermediateModel := new(blockchainv2.ConfigCAIntermediate)
				configCaIntermediateModel.Parentserver = configCaIntermediateParentserverModel
				configCaIntermediateModel.Enrollment = configCaIntermediateEnrollmentModel
				configCaIntermediateModel.Tls = configCaIntermediateTlsModel

				// Construct an instance of the ConfigCARegistry model
				configCaRegistryModel := new(blockchainv2.ConfigCARegistry)
				configCaRegistryModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryModel.Identities = []blockchainv2.ConfigCARegistryIdentitiesItem{*configCaRegistryIdentitiesItemModel}

				// Construct an instance of the ConfigCATls model
				configCaTlsModel := new(blockchainv2.ConfigCATls)
				configCaTlsModel.Keyfile = core.StringPtr("testString")
				configCaTlsModel.Certfile = core.StringPtr("testString")
				configCaTlsModel.Clientauth = configCaTlsClientauthModel

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv2.Metrics)
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv2.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv2.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ConfigCAUpdate model
				configCaUpdateModel := new(blockchainv2.ConfigCAUpdate)
				configCaUpdateModel.Cors = configCaCorsModel
				configCaUpdateModel.Debug = core.BoolPtr(false)
				configCaUpdateModel.Crlsizelimit = core.Float64Ptr(float64(512000))
				configCaUpdateModel.Tls = configCaTlsModel
				configCaUpdateModel.Ca = configCaCaModel
				configCaUpdateModel.Crl = configCaCrlModel
				configCaUpdateModel.Registry = configCaRegistryModel
				configCaUpdateModel.Db = configCaDbModel
				configCaUpdateModel.Affiliations = configCaAffiliationsModel
				configCaUpdateModel.Csr = configCaCsrModel
				configCaUpdateModel.Idemix = configCaIdemixModel
				configCaUpdateModel.BCCSP = bccspModel
				configCaUpdateModel.Intermediate = configCaIntermediateModel
				configCaUpdateModel.Cfg = configCaCfgModel
				configCaUpdateModel.Metrics = metricsModel

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv2.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the UpdateCaBodyConfigOverride model
				updateCaBodyConfigOverrideModel := new(blockchainv2.UpdateCaBodyConfigOverride)
				updateCaBodyConfigOverrideModel.Ca = configCaUpdateModel

				// Construct an instance of the UpdateCaBodyResources model
				updateCaBodyResourcesModel := new(blockchainv2.UpdateCaBodyResources)
				updateCaBodyResourcesModel.Ca = resourceObjectModel

				// Construct an instance of the UpdateCaOptions model
				updateCaOptionsModel := new(blockchainv2.UpdateCaOptions)
				updateCaOptionsModel.ID = core.StringPtr("testString")
				updateCaOptionsModel.Resources = updateCaBodyResourcesModel
				updateCaOptionsModel.Zone = core.StringPtr("testString")
				updateCaOptionsModel.ConfigOverride = updateCaBodyConfigOverrideModel
				updateCaOptionsModel.Replicas = core.Float64Ptr(float64(1))
				updateCaOptionsModel.Version = core.StringPtr("1.4.6-1")
 				updateCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateCa(updateCaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateCa with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ConfigCADbTlsClient model
				configCaDbTlsClientModel := new(blockchainv2.ConfigCADbTlsClient)
				configCaDbTlsClientModel.Certfile = core.StringPtr("testString")
				configCaDbTlsClientModel.Keyfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateTlsClient model
				configCaIntermediateTlsClientModel := new(blockchainv2.ConfigCAIntermediateTlsClient)
				configCaIntermediateTlsClientModel.Certfile = core.StringPtr("testString")
				configCaIntermediateTlsClientModel.Keyfile = core.StringPtr("testString")

				// Construct an instance of the IdentityAttrs model
				identityAttrsModel := new(blockchainv2.IdentityAttrs)
				identityAttrsModel.HfRegistrarRoles = core.StringPtr("*")
				identityAttrsModel.HfRegistrarDelegateRoles = core.StringPtr("*")
				identityAttrsModel.HfRevoker = core.BoolPtr(true)
				identityAttrsModel.HfIntermediateCA = core.BoolPtr(true)
				identityAttrsModel.HfGenCRL = core.BoolPtr(true)
				identityAttrsModel.HfRegistrarAttributes = core.StringPtr("*")
				identityAttrsModel.HfAffiliationMgr = core.BoolPtr(true)

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv2.BccspPKCS11)
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv2.BccspSW)
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the ConfigCACfgIdentities model
				configCaCfgIdentitiesModel := new(blockchainv2.ConfigCACfgIdentities)
				configCaCfgIdentitiesModel.Passwordattempts = core.Float64Ptr(float64(10))
				configCaCfgIdentitiesModel.Allowremove = core.BoolPtr(false)

				// Construct an instance of the ConfigCACsrCa model
				configCaCsrCaModel := new(blockchainv2.ConfigCACsrCa)
				configCaCsrCaModel.Expiry = core.StringPtr("131400h")
				configCaCsrCaModel.Pathlength = core.Float64Ptr(float64(0))

				// Construct an instance of the ConfigCACsrKeyrequest model
				configCaCsrKeyrequestModel := new(blockchainv2.ConfigCACsrKeyrequest)
				configCaCsrKeyrequestModel.Algo = core.StringPtr("ecdsa")
				configCaCsrKeyrequestModel.Size = core.Float64Ptr(float64(256))

				// Construct an instance of the ConfigCACsrNamesItem model
				configCaCsrNamesItemModel := new(blockchainv2.ConfigCACsrNamesItem)
				configCaCsrNamesItemModel.C = core.StringPtr("US")
				configCaCsrNamesItemModel.ST = core.StringPtr("North Carolina")
				configCaCsrNamesItemModel.L = core.StringPtr("Raleigh")
				configCaCsrNamesItemModel.O = core.StringPtr("Hyperledger")
				configCaCsrNamesItemModel.OU = core.StringPtr("Fabric")

				// Construct an instance of the ConfigCADbTls model
				configCaDbTlsModel := new(blockchainv2.ConfigCADbTls)
				configCaDbTlsModel.Certfiles = []string{"testString"}
				configCaDbTlsModel.Client = configCaDbTlsClientModel
				configCaDbTlsModel.Enabled = core.BoolPtr(false)

				// Construct an instance of the ConfigCAIntermediateEnrollment model
				configCaIntermediateEnrollmentModel := new(blockchainv2.ConfigCAIntermediateEnrollment)
				configCaIntermediateEnrollmentModel.Hosts = core.StringPtr("localhost")
				configCaIntermediateEnrollmentModel.Profile = core.StringPtr("testString")
				configCaIntermediateEnrollmentModel.Label = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateParentserver model
				configCaIntermediateParentserverModel := new(blockchainv2.ConfigCAIntermediateParentserver)
				configCaIntermediateParentserverModel.URL = core.StringPtr("testString")
				configCaIntermediateParentserverModel.Caname = core.StringPtr("testString")

				// Construct an instance of the ConfigCAIntermediateTls model
				configCaIntermediateTlsModel := new(blockchainv2.ConfigCAIntermediateTls)
				configCaIntermediateTlsModel.Certfiles = []string{"testString"}
				configCaIntermediateTlsModel.Client = configCaIntermediateTlsClientModel

				// Construct an instance of the ConfigCARegistryIdentitiesItem model
				configCaRegistryIdentitiesItemModel := new(blockchainv2.ConfigCARegistryIdentitiesItem)
				configCaRegistryIdentitiesItemModel.Name = core.StringPtr("admin")
				configCaRegistryIdentitiesItemModel.Pass = core.StringPtr("password")
				configCaRegistryIdentitiesItemModel.Type = core.StringPtr("client")
				configCaRegistryIdentitiesItemModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryIdentitiesItemModel.Affiliation = core.StringPtr("testString")
				configCaRegistryIdentitiesItemModel.Attrs = identityAttrsModel

				// Construct an instance of the ConfigCATlsClientauth model
				configCaTlsClientauthModel := new(blockchainv2.ConfigCATlsClientauth)
				configCaTlsClientauthModel.Type = core.StringPtr("noclientcert")
				configCaTlsClientauthModel.Certfiles = []string{"testString"}

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv2.MetricsStatsd)
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv2.Bccsp)
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model

				// Construct an instance of the ConfigCAAffiliations model
				configCaAffiliationsModel := new(blockchainv2.ConfigCAAffiliations)
				configCaAffiliationsModel.Org1 = []string{"department1"}
				configCaAffiliationsModel.Org2 = []string{"department1"}
				configCaAffiliationsModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ConfigCACa model
				configCaCaModel := new(blockchainv2.ConfigCACa)
				configCaCaModel.Keyfile = core.StringPtr("testString")
				configCaCaModel.Certfile = core.StringPtr("testString")
				configCaCaModel.Chainfile = core.StringPtr("testString")

				// Construct an instance of the ConfigCACfg model
				configCaCfgModel := new(blockchainv2.ConfigCACfg)
				configCaCfgModel.Identities = configCaCfgIdentitiesModel

				// Construct an instance of the ConfigCACors model
				configCaCorsModel := new(blockchainv2.ConfigCACors)
				configCaCorsModel.Enabled = core.BoolPtr(true)
				configCaCorsModel.Origins = []string{"*"}

				// Construct an instance of the ConfigCACrl model
				configCaCrlModel := new(blockchainv2.ConfigCACrl)
				configCaCrlModel.Expiry = core.StringPtr("24h")

				// Construct an instance of the ConfigCACsr model
				configCaCsrModel := new(blockchainv2.ConfigCACsr)
				configCaCsrModel.Cn = core.StringPtr("ca")
				configCaCsrModel.Keyrequest = configCaCsrKeyrequestModel
				configCaCsrModel.Names = []blockchainv2.ConfigCACsrNamesItem{*configCaCsrNamesItemModel}
				configCaCsrModel.Hosts = []string{"localhost"}
				configCaCsrModel.Ca = configCaCsrCaModel

				// Construct an instance of the ConfigCADb model
				configCaDbModel := new(blockchainv2.ConfigCADb)
				configCaDbModel.Type = core.StringPtr("postgres")
				configCaDbModel.Datasource = core.StringPtr("host=fake.databases.appdomain.cloud port=31941 user=ibm_cloud password=password dbname=ibmclouddb sslmode=verify-full")
				configCaDbModel.Tls = configCaDbTlsModel

				// Construct an instance of the ConfigCAIdemix model
				configCaIdemixModel := new(blockchainv2.ConfigCAIdemix)
				configCaIdemixModel.Rhpoolsize = core.Float64Ptr(float64(100))
				configCaIdemixModel.Nonceexpiration = core.StringPtr("15s")
				configCaIdemixModel.Noncesweepinterval = core.StringPtr("15m")

				// Construct an instance of the ConfigCAIntermediate model
				configCaIntermediateModel := new(blockchainv2.ConfigCAIntermediate)
				configCaIntermediateModel.Parentserver = configCaIntermediateParentserverModel
				configCaIntermediateModel.Enrollment = configCaIntermediateEnrollmentModel
				configCaIntermediateModel.Tls = configCaIntermediateTlsModel

				// Construct an instance of the ConfigCARegistry model
				configCaRegistryModel := new(blockchainv2.ConfigCARegistry)
				configCaRegistryModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryModel.Identities = []blockchainv2.ConfigCARegistryIdentitiesItem{*configCaRegistryIdentitiesItemModel}

				// Construct an instance of the ConfigCATls model
				configCaTlsModel := new(blockchainv2.ConfigCATls)
				configCaTlsModel.Keyfile = core.StringPtr("testString")
				configCaTlsModel.Certfile = core.StringPtr("testString")
				configCaTlsModel.Clientauth = configCaTlsClientauthModel

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv2.Metrics)
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv2.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv2.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ConfigCAUpdate model
				configCaUpdateModel := new(blockchainv2.ConfigCAUpdate)
				configCaUpdateModel.Cors = configCaCorsModel
				configCaUpdateModel.Debug = core.BoolPtr(false)
				configCaUpdateModel.Crlsizelimit = core.Float64Ptr(float64(512000))
				configCaUpdateModel.Tls = configCaTlsModel
				configCaUpdateModel.Ca = configCaCaModel
				configCaUpdateModel.Crl = configCaCrlModel
				configCaUpdateModel.Registry = configCaRegistryModel
				configCaUpdateModel.Db = configCaDbModel
				configCaUpdateModel.Affiliations = configCaAffiliationsModel
				configCaUpdateModel.Csr = configCaCsrModel
				configCaUpdateModel.Idemix = configCaIdemixModel
				configCaUpdateModel.BCCSP = bccspModel
				configCaUpdateModel.Intermediate = configCaIntermediateModel
				configCaUpdateModel.Cfg = configCaCfgModel
				configCaUpdateModel.Metrics = metricsModel

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv2.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the UpdateCaBodyConfigOverride model
				updateCaBodyConfigOverrideModel := new(blockchainv2.UpdateCaBodyConfigOverride)
				updateCaBodyConfigOverrideModel.Ca = configCaUpdateModel

				// Construct an instance of the UpdateCaBodyResources model
				updateCaBodyResourcesModel := new(blockchainv2.UpdateCaBodyResources)
				updateCaBodyResourcesModel.Ca = resourceObjectModel

				// Construct an instance of the UpdateCaOptions model
				updateCaOptionsModel := new(blockchainv2.UpdateCaOptions)
				updateCaOptionsModel.ID = core.StringPtr("testString")
				updateCaOptionsModel.Resources = updateCaBodyResourcesModel
				updateCaOptionsModel.Zone = core.StringPtr("testString")
				updateCaOptionsModel.ConfigOverride = updateCaBodyConfigOverrideModel
				updateCaOptionsModel.Replicas = core.Float64Ptr(float64(1))
				updateCaOptionsModel.Version = core.StringPtr("1.4.6-1")
				updateCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateCa(updateCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateCaOptions model with no property values
				updateCaOptionsModelNew := new(blockchainv2.UpdateCaOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateCa(updateCaOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`EditCa(editCaOptions *EditCaOptions) - Operation response error`, func() {
		editCaPath := "/ak/api/v2/components/fabric-ca/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(editCaPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke EditCa with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the EditCaOptions model
				editCaOptionsModel := new(blockchainv2.EditCaOptions)
				editCaOptionsModel.ID = core.StringPtr("testString")
				editCaOptionsModel.DisplayName = core.StringPtr("My CA")
				editCaOptionsModel.ApiURL = core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:7054")
				editCaOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443")
				editCaOptionsModel.CaName = core.StringPtr("ca")
				editCaOptionsModel.Location = core.StringPtr("ibmcloud")
				editCaOptionsModel.Tags = []string{"testString"}
				editCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.EditCa(editCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`EditCa(editCaOptions *EditCaOptions)`, func() {
		editCaPath := "/ak/api/v2/components/fabric-ca/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(editCaPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "component-1", "dep_component_id": "admin", "ca_name": "ca", "display_name": "My CA", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "operations_url": "https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443", "config_override": {"anyKey": "anyValue"}, "location": "ibmcloud", "timestamp": 1537262855753, "resources": {"ca": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"ca": {"size": "4GiB", "class": "default"}}, "tags": ["Tags"], "tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "version": "1.4.6-1", "zone": "Zone"}`)
				}))
			})
			It(`Invoke EditCa successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.EditCa(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EditCaOptions model
				editCaOptionsModel := new(blockchainv2.EditCaOptions)
				editCaOptionsModel.ID = core.StringPtr("testString")
				editCaOptionsModel.DisplayName = core.StringPtr("My CA")
				editCaOptionsModel.ApiURL = core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:7054")
				editCaOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443")
				editCaOptionsModel.CaName = core.StringPtr("ca")
				editCaOptionsModel.Location = core.StringPtr("ibmcloud")
				editCaOptionsModel.Tags = []string{"testString"}
 				editCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.EditCa(editCaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke EditCa with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the EditCaOptions model
				editCaOptionsModel := new(blockchainv2.EditCaOptions)
				editCaOptionsModel.ID = core.StringPtr("testString")
				editCaOptionsModel.DisplayName = core.StringPtr("My CA")
				editCaOptionsModel.ApiURL = core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:7054")
				editCaOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443")
				editCaOptionsModel.CaName = core.StringPtr("ca")
				editCaOptionsModel.Location = core.StringPtr("ibmcloud")
				editCaOptionsModel.Tags = []string{"testString"}
				editCaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.EditCa(editCaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the EditCaOptions model with no property values
				editCaOptionsModelNew := new(blockchainv2.EditCaOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.EditCa(editCaOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreatePeer(createPeerOptions *CreatePeerOptions) - Operation response error`, func() {
		createPeerPath := "/ak/api/v2/kubernetes/components/fabric-peer"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createPeerPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreatePeer with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy model
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel := new(blockchainv2.ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy)
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.RequiredPeerCount = core.Float64Ptr(float64(0))
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.MaxPeerCount = core.Float64Ptr(float64(1))

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv2.BccspPKCS11)
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv2.BccspSW)
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the ConfigObjectEnrollmentComponentCatls model
				configObjectEnrollmentComponentCatlsModel := new(blockchainv2.ConfigObjectEnrollmentComponentCatls)
				configObjectEnrollmentComponentCatlsModel.Cacert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")

				// Construct an instance of the ConfigObjectEnrollmentTlsCatls model
				configObjectEnrollmentTlsCatlsModel := new(blockchainv2.ConfigObjectEnrollmentTlsCatls)
				configObjectEnrollmentTlsCatlsModel.Cacert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")

				// Construct an instance of the ConfigObjectEnrollmentTlsCsr model
				configObjectEnrollmentTlsCsrModel := new(blockchainv2.ConfigObjectEnrollmentTlsCsr)
				configObjectEnrollmentTlsCsrModel.Hosts = []string{"testString"}

				// Construct an instance of the ConfigPeerDeliveryclientAddressOverridesItem model
				configPeerDeliveryclientAddressOverridesItemModel := new(blockchainv2.ConfigPeerDeliveryclientAddressOverridesItem)
				configPeerDeliveryclientAddressOverridesItemModel.From = core.StringPtr("n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.To = core.StringPtr("n3a3ec3-myorderer2.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.CaCertsFile = core.StringPtr("my-data/cert.pem")

				// Construct an instance of the ConfigPeerGossipElection model
				configPeerGossipElectionModel := new(blockchainv2.ConfigPeerGossipElection)
				configPeerGossipElectionModel.StartupGracePeriod = core.StringPtr("15s")
				configPeerGossipElectionModel.MembershipSampleInterval = core.StringPtr("1s")
				configPeerGossipElectionModel.LeaderAliveThreshold = core.StringPtr("10s")
				configPeerGossipElectionModel.LeaderElectionDuration = core.StringPtr("5s")

				// Construct an instance of the ConfigPeerGossipPvtData model
				configPeerGossipPvtDataModel := new(blockchainv2.ConfigPeerGossipPvtData)
				configPeerGossipPvtDataModel.PullRetryThreshold = core.StringPtr("60s")
				configPeerGossipPvtDataModel.TransientstoreMaxBlockRetention = core.Float64Ptr(float64(1000))
				configPeerGossipPvtDataModel.PushAckTimeout = core.StringPtr("3s")
				configPeerGossipPvtDataModel.BtlPullMargin = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileBatchSize = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileSleepInterval = core.StringPtr("1m")
				configPeerGossipPvtDataModel.ReconciliationEnabled = core.BoolPtr(true)
				configPeerGossipPvtDataModel.SkipPullingInvalidTransactionsDuringCommit = core.BoolPtr(false)
				configPeerGossipPvtDataModel.ImplicitCollectionDisseminationPolicy = configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel

				// Construct an instance of the ConfigPeerGossipState model
				configPeerGossipStateModel := new(blockchainv2.ConfigPeerGossipState)
				configPeerGossipStateModel.Enabled = core.BoolPtr(true)
				configPeerGossipStateModel.CheckInterval = core.StringPtr("10s")
				configPeerGossipStateModel.ResponseTimeout = core.StringPtr("3s")
				configPeerGossipStateModel.BatchSize = core.Float64Ptr(float64(10))
				configPeerGossipStateModel.BlockBufferSize = core.Float64Ptr(float64(100))
				configPeerGossipStateModel.MaxRetries = core.Float64Ptr(float64(3))

				// Construct an instance of the ConfigPeerKeepaliveClient model
				configPeerKeepaliveClientModel := new(blockchainv2.ConfigPeerKeepaliveClient)
				configPeerKeepaliveClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveClientModel.Timeout = core.StringPtr("20s")

				// Construct an instance of the ConfigPeerKeepaliveDeliveryClient model
				configPeerKeepaliveDeliveryClientModel := new(blockchainv2.ConfigPeerKeepaliveDeliveryClient)
				configPeerKeepaliveDeliveryClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveDeliveryClientModel.Timeout = core.StringPtr("20s")

				// Construct an instance of the ConfigPeerLimitsConcurrency model
				configPeerLimitsConcurrencyModel := new(blockchainv2.ConfigPeerLimitsConcurrency)
				configPeerLimitsConcurrencyModel.EndorserService = map[string]interface{}{"anyKey": "anyValue"}
				configPeerLimitsConcurrencyModel.DeliverService = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv2.Bccsp)
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model

				// Construct an instance of the ConfigObjectEnrollmentComponent model
				configObjectEnrollmentComponentModel := new(blockchainv2.ConfigObjectEnrollmentComponent)
				configObjectEnrollmentComponentModel.Cahost = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				configObjectEnrollmentComponentModel.Caport = core.Float64Ptr(float64(7054))
				configObjectEnrollmentComponentModel.Caname = core.StringPtr("ca")
				configObjectEnrollmentComponentModel.Catls = configObjectEnrollmentComponentCatlsModel
				configObjectEnrollmentComponentModel.Enrollid = core.StringPtr("admin")
				configObjectEnrollmentComponentModel.Enrollsecret = core.StringPtr("password")
				configObjectEnrollmentComponentModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the ConfigObjectEnrollmentTls model
				configObjectEnrollmentTlsModel := new(blockchainv2.ConfigObjectEnrollmentTls)
				configObjectEnrollmentTlsModel.Cahost = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				configObjectEnrollmentTlsModel.Caport = core.Float64Ptr(float64(7054))
				configObjectEnrollmentTlsModel.Caname = core.StringPtr("tlsca")
				configObjectEnrollmentTlsModel.Catls = configObjectEnrollmentTlsCatlsModel
				configObjectEnrollmentTlsModel.Enrollid = core.StringPtr("admin")
				configObjectEnrollmentTlsModel.Enrollsecret = core.StringPtr("password")
				configObjectEnrollmentTlsModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				configObjectEnrollmentTlsModel.Csr = configObjectEnrollmentTlsCsrModel

				// Construct an instance of the ConfigPeerAdminService model
				configPeerAdminServiceModel := new(blockchainv2.ConfigPeerAdminService)
				configPeerAdminServiceModel.ListenAddress = core.StringPtr("0.0.0.0:7051")

				// Construct an instance of the ConfigPeerAuthentication model
				configPeerAuthenticationModel := new(blockchainv2.ConfigPeerAuthentication)
				configPeerAuthenticationModel.Timewindow = core.StringPtr("15m")

				// Construct an instance of the ConfigPeerChaincodeExternalBuildersItem model
				configPeerChaincodeExternalBuildersItemModel := new(blockchainv2.ConfigPeerChaincodeExternalBuildersItem)
				configPeerChaincodeExternalBuildersItemModel.Path = core.StringPtr("/path/to/directory")
				configPeerChaincodeExternalBuildersItemModel.Name = core.StringPtr("descriptive-build-name")
				configPeerChaincodeExternalBuildersItemModel.EnvironmentWhitelist = []string{"GOPROXY"}

				// Construct an instance of the ConfigPeerChaincodeGolang model
				configPeerChaincodeGolangModel := new(blockchainv2.ConfigPeerChaincodeGolang)
				configPeerChaincodeGolangModel.DynamicLink = core.BoolPtr(false)

				// Construct an instance of the ConfigPeerChaincodeLogging model
				configPeerChaincodeLoggingModel := new(blockchainv2.ConfigPeerChaincodeLogging)
				configPeerChaincodeLoggingModel.Level = core.StringPtr("info")
				configPeerChaincodeLoggingModel.Shim = core.StringPtr("warning")
				configPeerChaincodeLoggingModel.Format = core.StringPtr("%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}")

				// Construct an instance of the ConfigPeerChaincodeSystem model
				configPeerChaincodeSystemModel := new(blockchainv2.ConfigPeerChaincodeSystem)
				configPeerChaincodeSystemModel.Cscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Lscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Escc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Vscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Qscc = core.BoolPtr(true)

				// Construct an instance of the ConfigPeerClient model
				configPeerClientModel := new(blockchainv2.ConfigPeerClient)
				configPeerClientModel.ConnTimeout = core.StringPtr("2s")

				// Construct an instance of the ConfigPeerDeliveryclient model
				configPeerDeliveryclientModel := new(blockchainv2.ConfigPeerDeliveryclient)
				configPeerDeliveryclientModel.ReconnectTotalTimeThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.ConnTimeout = core.StringPtr("2s")
				configPeerDeliveryclientModel.ReConnectBackoffThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.AddressOverrides = []blockchainv2.ConfigPeerDeliveryclientAddressOverridesItem{*configPeerDeliveryclientAddressOverridesItemModel}

				// Construct an instance of the ConfigPeerDiscovery model
				configPeerDiscoveryModel := new(blockchainv2.ConfigPeerDiscovery)
				configPeerDiscoveryModel.Enabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheEnabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheMaxSize = core.Float64Ptr(float64(1000))
				configPeerDiscoveryModel.AuthCachePurgeRetentionRatio = core.Float64Ptr(float64(0.75))
				configPeerDiscoveryModel.OrgMembersAllowedAccess = core.BoolPtr(false)

				// Construct an instance of the ConfigPeerGossip model
				configPeerGossipModel := new(blockchainv2.ConfigPeerGossip)
				configPeerGossipModel.UseLeaderElection = core.BoolPtr(true)
				configPeerGossipModel.OrgLeader = core.BoolPtr(false)
				configPeerGossipModel.MembershipTrackerInterval = core.StringPtr("5s")
				configPeerGossipModel.MaxBlockCountToStore = core.Float64Ptr(float64(100))
				configPeerGossipModel.MaxPropagationBurstLatency = core.StringPtr("10ms")
				configPeerGossipModel.MaxPropagationBurstSize = core.Float64Ptr(float64(10))
				configPeerGossipModel.PropagateIterations = core.Float64Ptr(float64(3))
				configPeerGossipModel.PullInterval = core.StringPtr("4s")
				configPeerGossipModel.PullPeerNum = core.Float64Ptr(float64(3))
				configPeerGossipModel.RequestStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.PublishStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.StateInfoRetentionInterval = core.StringPtr("0s")
				configPeerGossipModel.PublishCertPeriod = core.StringPtr("10s")
				configPeerGossipModel.SkipBlockVerification = core.BoolPtr(false)
				configPeerGossipModel.DialTimeout = core.StringPtr("3s")
				configPeerGossipModel.ConnTimeout = core.StringPtr("2s")
				configPeerGossipModel.RecvBuffSize = core.Float64Ptr(float64(20))
				configPeerGossipModel.SendBuffSize = core.Float64Ptr(float64(200))
				configPeerGossipModel.DigestWaitTime = core.StringPtr("1s")
				configPeerGossipModel.RequestWaitTime = core.StringPtr("1500ms")
				configPeerGossipModel.ResponseWaitTime = core.StringPtr("2s")
				configPeerGossipModel.AliveTimeInterval = core.StringPtr("5s")
				configPeerGossipModel.AliveExpirationTimeout = core.StringPtr("25s")
				configPeerGossipModel.ReconnectInterval = core.StringPtr("25s")
				configPeerGossipModel.Election = configPeerGossipElectionModel
				configPeerGossipModel.PvtData = configPeerGossipPvtDataModel
				configPeerGossipModel.State = configPeerGossipStateModel

				// Construct an instance of the ConfigPeerKeepalive model
				configPeerKeepaliveModel := new(blockchainv2.ConfigPeerKeepalive)
				configPeerKeepaliveModel.MinInterval = core.StringPtr("60s")
				configPeerKeepaliveModel.Client = configPeerKeepaliveClientModel
				configPeerKeepaliveModel.DeliveryClient = configPeerKeepaliveDeliveryClientModel

				// Construct an instance of the ConfigPeerLimits model
				configPeerLimitsModel := new(blockchainv2.ConfigPeerLimits)
				configPeerLimitsModel.Concurrency = configPeerLimitsConcurrencyModel

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv2.MetricsStatsd)
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the MspConfigData model
				mspConfigDataModel := new(blockchainv2.MspConfigData)
				mspConfigDataModel.Keystore = core.StringPtr("testString")
				mspConfigDataModel.Signcerts = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspConfigDataModel.Cacerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				mspConfigDataModel.Intermediatecerts = []string{"testString"}
				mspConfigDataModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv2.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv2.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ConfigObjectEnrollment model
				configObjectEnrollmentModel := new(blockchainv2.ConfigObjectEnrollment)
				configObjectEnrollmentModel.Component = configObjectEnrollmentComponentModel
				configObjectEnrollmentModel.Tls = configObjectEnrollmentTlsModel

				// Construct an instance of the ConfigObjectMsp model
				configObjectMspModel := new(blockchainv2.ConfigObjectMsp)
				configObjectMspModel.Component = mspConfigDataModel
				configObjectMspModel.Tls = mspConfigDataModel
				configObjectMspModel.Clientauth = mspConfigDataModel

				// Construct an instance of the ConfigPeerChaincode model
				configPeerChaincodeModel := new(blockchainv2.ConfigPeerChaincode)
				configPeerChaincodeModel.Golang = configPeerChaincodeGolangModel
				configPeerChaincodeModel.ExternalBuilders = []blockchainv2.ConfigPeerChaincodeExternalBuildersItem{*configPeerChaincodeExternalBuildersItemModel}
				configPeerChaincodeModel.InstallTimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Startuptimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Executetimeout = core.StringPtr("30s")
				configPeerChaincodeModel.System = configPeerChaincodeSystemModel
				configPeerChaincodeModel.Logging = configPeerChaincodeLoggingModel

				// Construct an instance of the ConfigPeerCreatePeer model
				configPeerCreatePeerModel := new(blockchainv2.ConfigPeerCreatePeer)
				configPeerCreatePeerModel.ID = core.StringPtr("john-doe")
				configPeerCreatePeerModel.NetworkID = core.StringPtr("dev")
				configPeerCreatePeerModel.Keepalive = configPeerKeepaliveModel
				configPeerCreatePeerModel.Gossip = configPeerGossipModel
				configPeerCreatePeerModel.Authentication = configPeerAuthenticationModel
				configPeerCreatePeerModel.BCCSP = bccspModel
				configPeerCreatePeerModel.Client = configPeerClientModel
				configPeerCreatePeerModel.Deliveryclient = configPeerDeliveryclientModel
				configPeerCreatePeerModel.AdminService = configPeerAdminServiceModel
				configPeerCreatePeerModel.ValidatorPoolSize = core.Float64Ptr(float64(8))
				configPeerCreatePeerModel.Discovery = configPeerDiscoveryModel
				configPeerCreatePeerModel.Limits = configPeerLimitsModel

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv2.Metrics)
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv2.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectCouchDb model
				resourceObjectCouchDbModel := new(blockchainv2.ResourceObjectCouchDb)
				resourceObjectCouchDbModel.Requests = resourceRequestsModel
				resourceObjectCouchDbModel.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectFabV1 model
				resourceObjectFabV1Model := new(blockchainv2.ResourceObjectFabV1)
				resourceObjectFabV1Model.Requests = resourceRequestsModel
				resourceObjectFabV1Model.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectFabV2 model
				resourceObjectFabV2Model := new(blockchainv2.ResourceObjectFabV2)
				resourceObjectFabV2Model.Requests = resourceRequestsModel
				resourceObjectFabV2Model.Limits = resourceLimitsModel

				// Construct an instance of the StorageObject model
				storageObjectModel := new(blockchainv2.StorageObject)
				storageObjectModel.Size = core.StringPtr("4GiB")
				storageObjectModel.Class = core.StringPtr("default")

				// Construct an instance of the ConfigObject model
				configObjectModel := new(blockchainv2.ConfigObject)
				configObjectModel.Enrollment = configObjectEnrollmentModel
				configObjectModel.Msp = configObjectMspModel

				// Construct an instance of the ConfigPeerCreate model
				configPeerCreateModel := new(blockchainv2.ConfigPeerCreate)
				configPeerCreateModel.Peer = configPeerCreatePeerModel
				configPeerCreateModel.Chaincode = configPeerChaincodeModel
				configPeerCreateModel.Metrics = metricsModel

				// Construct an instance of the CreatePeerBodyStorage model
				createPeerBodyStorageModel := new(blockchainv2.CreatePeerBodyStorage)
				createPeerBodyStorageModel.Peer = storageObjectModel
				createPeerBodyStorageModel.Statedb = storageObjectModel

				// Construct an instance of the Hsm model
				hsmModel := new(blockchainv2.Hsm)
				hsmModel.Pkcs11endpoint = core.StringPtr("tcp://example.com:666")

				// Construct an instance of the PeerResources model
				peerResourcesModel := new(blockchainv2.PeerResources)
				peerResourcesModel.Chaincodelauncher = resourceObjectFabV2Model
				peerResourcesModel.Couchdb = resourceObjectCouchDbModel
				peerResourcesModel.Statedb = resourceObjectModel
				peerResourcesModel.Dind = resourceObjectFabV1Model
				peerResourcesModel.Fluentd = resourceObjectFabV1Model
				peerResourcesModel.Peer = resourceObjectModel
				peerResourcesModel.Proxy = resourceObjectModel

				// Construct an instance of the CreatePeerOptions model
				createPeerOptionsModel := new(blockchainv2.CreatePeerOptions)
				createPeerOptionsModel.MspID = core.StringPtr("Org1")
				createPeerOptionsModel.DisplayName = core.StringPtr("My Peer")
				createPeerOptionsModel.Config = configObjectModel
				createPeerOptionsModel.ConfigOverride = configPeerCreateModel
				createPeerOptionsModel.Resources = peerResourcesModel
				createPeerOptionsModel.Storage = createPeerBodyStorageModel
				createPeerOptionsModel.Zone = core.StringPtr("testString")
				createPeerOptionsModel.StateDb = core.StringPtr("couchdb")
				createPeerOptionsModel.Tags = []string{"testString"}
				createPeerOptionsModel.Hsm = hsmModel
				createPeerOptionsModel.Region = core.StringPtr("testString")
				createPeerOptionsModel.Version = core.StringPtr("1.4.6-1")
				createPeerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreatePeer(createPeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreatePeer(createPeerOptions *CreatePeerOptions)`, func() {
		createPeerPath := "/ak/api/v2/kubernetes/components/fabric-peer"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createPeerPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "component-1", "dep_component_id": "admin", "type": "fabric-peer", "display_name": "My Peer", "grpcwp_url": "https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "operations_url": "https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443", "msp_id": "Org1", "config_override": {"anyKey": "anyValue"}, "node_ou": {"enabled": true}, "ecert": {"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "cacert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}, "location": "ibmcloud", "state_db": "couchdb", "timestamp": 1537262855753, "resources": {"peer": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "statedb": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"peer": {"size": "4GiB", "class": "default"}, "statedb": {"size": "4GiB", "class": "default"}}, "tags": ["Tags"], "tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "version": "1.4.6-1", "zone": "Zone"}`)
				}))
			})
			It(`Invoke CreatePeer successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreatePeer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy model
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel := new(blockchainv2.ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy)
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.RequiredPeerCount = core.Float64Ptr(float64(0))
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.MaxPeerCount = core.Float64Ptr(float64(1))

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv2.BccspPKCS11)
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv2.BccspSW)
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the ConfigObjectEnrollmentComponentCatls model
				configObjectEnrollmentComponentCatlsModel := new(blockchainv2.ConfigObjectEnrollmentComponentCatls)
				configObjectEnrollmentComponentCatlsModel.Cacert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")

				// Construct an instance of the ConfigObjectEnrollmentTlsCatls model
				configObjectEnrollmentTlsCatlsModel := new(blockchainv2.ConfigObjectEnrollmentTlsCatls)
				configObjectEnrollmentTlsCatlsModel.Cacert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")

				// Construct an instance of the ConfigObjectEnrollmentTlsCsr model
				configObjectEnrollmentTlsCsrModel := new(blockchainv2.ConfigObjectEnrollmentTlsCsr)
				configObjectEnrollmentTlsCsrModel.Hosts = []string{"testString"}

				// Construct an instance of the ConfigPeerDeliveryclientAddressOverridesItem model
				configPeerDeliveryclientAddressOverridesItemModel := new(blockchainv2.ConfigPeerDeliveryclientAddressOverridesItem)
				configPeerDeliveryclientAddressOverridesItemModel.From = core.StringPtr("n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.To = core.StringPtr("n3a3ec3-myorderer2.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.CaCertsFile = core.StringPtr("my-data/cert.pem")

				// Construct an instance of the ConfigPeerGossipElection model
				configPeerGossipElectionModel := new(blockchainv2.ConfigPeerGossipElection)
				configPeerGossipElectionModel.StartupGracePeriod = core.StringPtr("15s")
				configPeerGossipElectionModel.MembershipSampleInterval = core.StringPtr("1s")
				configPeerGossipElectionModel.LeaderAliveThreshold = core.StringPtr("10s")
				configPeerGossipElectionModel.LeaderElectionDuration = core.StringPtr("5s")

				// Construct an instance of the ConfigPeerGossipPvtData model
				configPeerGossipPvtDataModel := new(blockchainv2.ConfigPeerGossipPvtData)
				configPeerGossipPvtDataModel.PullRetryThreshold = core.StringPtr("60s")
				configPeerGossipPvtDataModel.TransientstoreMaxBlockRetention = core.Float64Ptr(float64(1000))
				configPeerGossipPvtDataModel.PushAckTimeout = core.StringPtr("3s")
				configPeerGossipPvtDataModel.BtlPullMargin = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileBatchSize = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileSleepInterval = core.StringPtr("1m")
				configPeerGossipPvtDataModel.ReconciliationEnabled = core.BoolPtr(true)
				configPeerGossipPvtDataModel.SkipPullingInvalidTransactionsDuringCommit = core.BoolPtr(false)
				configPeerGossipPvtDataModel.ImplicitCollectionDisseminationPolicy = configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel

				// Construct an instance of the ConfigPeerGossipState model
				configPeerGossipStateModel := new(blockchainv2.ConfigPeerGossipState)
				configPeerGossipStateModel.Enabled = core.BoolPtr(true)
				configPeerGossipStateModel.CheckInterval = core.StringPtr("10s")
				configPeerGossipStateModel.ResponseTimeout = core.StringPtr("3s")
				configPeerGossipStateModel.BatchSize = core.Float64Ptr(float64(10))
				configPeerGossipStateModel.BlockBufferSize = core.Float64Ptr(float64(100))
				configPeerGossipStateModel.MaxRetries = core.Float64Ptr(float64(3))

				// Construct an instance of the ConfigPeerKeepaliveClient model
				configPeerKeepaliveClientModel := new(blockchainv2.ConfigPeerKeepaliveClient)
				configPeerKeepaliveClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveClientModel.Timeout = core.StringPtr("20s")

				// Construct an instance of the ConfigPeerKeepaliveDeliveryClient model
				configPeerKeepaliveDeliveryClientModel := new(blockchainv2.ConfigPeerKeepaliveDeliveryClient)
				configPeerKeepaliveDeliveryClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveDeliveryClientModel.Timeout = core.StringPtr("20s")

				// Construct an instance of the ConfigPeerLimitsConcurrency model
				configPeerLimitsConcurrencyModel := new(blockchainv2.ConfigPeerLimitsConcurrency)
				configPeerLimitsConcurrencyModel.EndorserService = map[string]interface{}{"anyKey": "anyValue"}
				configPeerLimitsConcurrencyModel.DeliverService = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv2.Bccsp)
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model

				// Construct an instance of the ConfigObjectEnrollmentComponent model
				configObjectEnrollmentComponentModel := new(blockchainv2.ConfigObjectEnrollmentComponent)
				configObjectEnrollmentComponentModel.Cahost = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				configObjectEnrollmentComponentModel.Caport = core.Float64Ptr(float64(7054))
				configObjectEnrollmentComponentModel.Caname = core.StringPtr("ca")
				configObjectEnrollmentComponentModel.Catls = configObjectEnrollmentComponentCatlsModel
				configObjectEnrollmentComponentModel.Enrollid = core.StringPtr("admin")
				configObjectEnrollmentComponentModel.Enrollsecret = core.StringPtr("password")
				configObjectEnrollmentComponentModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the ConfigObjectEnrollmentTls model
				configObjectEnrollmentTlsModel := new(blockchainv2.ConfigObjectEnrollmentTls)
				configObjectEnrollmentTlsModel.Cahost = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				configObjectEnrollmentTlsModel.Caport = core.Float64Ptr(float64(7054))
				configObjectEnrollmentTlsModel.Caname = core.StringPtr("tlsca")
				configObjectEnrollmentTlsModel.Catls = configObjectEnrollmentTlsCatlsModel
				configObjectEnrollmentTlsModel.Enrollid = core.StringPtr("admin")
				configObjectEnrollmentTlsModel.Enrollsecret = core.StringPtr("password")
				configObjectEnrollmentTlsModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				configObjectEnrollmentTlsModel.Csr = configObjectEnrollmentTlsCsrModel

				// Construct an instance of the ConfigPeerAdminService model
				configPeerAdminServiceModel := new(blockchainv2.ConfigPeerAdminService)
				configPeerAdminServiceModel.ListenAddress = core.StringPtr("0.0.0.0:7051")

				// Construct an instance of the ConfigPeerAuthentication model
				configPeerAuthenticationModel := new(blockchainv2.ConfigPeerAuthentication)
				configPeerAuthenticationModel.Timewindow = core.StringPtr("15m")

				// Construct an instance of the ConfigPeerChaincodeExternalBuildersItem model
				configPeerChaincodeExternalBuildersItemModel := new(blockchainv2.ConfigPeerChaincodeExternalBuildersItem)
				configPeerChaincodeExternalBuildersItemModel.Path = core.StringPtr("/path/to/directory")
				configPeerChaincodeExternalBuildersItemModel.Name = core.StringPtr("descriptive-build-name")
				configPeerChaincodeExternalBuildersItemModel.EnvironmentWhitelist = []string{"GOPROXY"}

				// Construct an instance of the ConfigPeerChaincodeGolang model
				configPeerChaincodeGolangModel := new(blockchainv2.ConfigPeerChaincodeGolang)
				configPeerChaincodeGolangModel.DynamicLink = core.BoolPtr(false)

				// Construct an instance of the ConfigPeerChaincodeLogging model
				configPeerChaincodeLoggingModel := new(blockchainv2.ConfigPeerChaincodeLogging)
				configPeerChaincodeLoggingModel.Level = core.StringPtr("info")
				configPeerChaincodeLoggingModel.Shim = core.StringPtr("warning")
				configPeerChaincodeLoggingModel.Format = core.StringPtr("%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}")

				// Construct an instance of the ConfigPeerChaincodeSystem model
				configPeerChaincodeSystemModel := new(blockchainv2.ConfigPeerChaincodeSystem)
				configPeerChaincodeSystemModel.Cscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Lscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Escc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Vscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Qscc = core.BoolPtr(true)

				// Construct an instance of the ConfigPeerClient model
				configPeerClientModel := new(blockchainv2.ConfigPeerClient)
				configPeerClientModel.ConnTimeout = core.StringPtr("2s")

				// Construct an instance of the ConfigPeerDeliveryclient model
				configPeerDeliveryclientModel := new(blockchainv2.ConfigPeerDeliveryclient)
				configPeerDeliveryclientModel.ReconnectTotalTimeThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.ConnTimeout = core.StringPtr("2s")
				configPeerDeliveryclientModel.ReConnectBackoffThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.AddressOverrides = []blockchainv2.ConfigPeerDeliveryclientAddressOverridesItem{*configPeerDeliveryclientAddressOverridesItemModel}

				// Construct an instance of the ConfigPeerDiscovery model
				configPeerDiscoveryModel := new(blockchainv2.ConfigPeerDiscovery)
				configPeerDiscoveryModel.Enabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheEnabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheMaxSize = core.Float64Ptr(float64(1000))
				configPeerDiscoveryModel.AuthCachePurgeRetentionRatio = core.Float64Ptr(float64(0.75))
				configPeerDiscoveryModel.OrgMembersAllowedAccess = core.BoolPtr(false)

				// Construct an instance of the ConfigPeerGossip model
				configPeerGossipModel := new(blockchainv2.ConfigPeerGossip)
				configPeerGossipModel.UseLeaderElection = core.BoolPtr(true)
				configPeerGossipModel.OrgLeader = core.BoolPtr(false)
				configPeerGossipModel.MembershipTrackerInterval = core.StringPtr("5s")
				configPeerGossipModel.MaxBlockCountToStore = core.Float64Ptr(float64(100))
				configPeerGossipModel.MaxPropagationBurstLatency = core.StringPtr("10ms")
				configPeerGossipModel.MaxPropagationBurstSize = core.Float64Ptr(float64(10))
				configPeerGossipModel.PropagateIterations = core.Float64Ptr(float64(3))
				configPeerGossipModel.PullInterval = core.StringPtr("4s")
				configPeerGossipModel.PullPeerNum = core.Float64Ptr(float64(3))
				configPeerGossipModel.RequestStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.PublishStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.StateInfoRetentionInterval = core.StringPtr("0s")
				configPeerGossipModel.PublishCertPeriod = core.StringPtr("10s")
				configPeerGossipModel.SkipBlockVerification = core.BoolPtr(false)
				configPeerGossipModel.DialTimeout = core.StringPtr("3s")
				configPeerGossipModel.ConnTimeout = core.StringPtr("2s")
				configPeerGossipModel.RecvBuffSize = core.Float64Ptr(float64(20))
				configPeerGossipModel.SendBuffSize = core.Float64Ptr(float64(200))
				configPeerGossipModel.DigestWaitTime = core.StringPtr("1s")
				configPeerGossipModel.RequestWaitTime = core.StringPtr("1500ms")
				configPeerGossipModel.ResponseWaitTime = core.StringPtr("2s")
				configPeerGossipModel.AliveTimeInterval = core.StringPtr("5s")
				configPeerGossipModel.AliveExpirationTimeout = core.StringPtr("25s")
				configPeerGossipModel.ReconnectInterval = core.StringPtr("25s")
				configPeerGossipModel.Election = configPeerGossipElectionModel
				configPeerGossipModel.PvtData = configPeerGossipPvtDataModel
				configPeerGossipModel.State = configPeerGossipStateModel

				// Construct an instance of the ConfigPeerKeepalive model
				configPeerKeepaliveModel := new(blockchainv2.ConfigPeerKeepalive)
				configPeerKeepaliveModel.MinInterval = core.StringPtr("60s")
				configPeerKeepaliveModel.Client = configPeerKeepaliveClientModel
				configPeerKeepaliveModel.DeliveryClient = configPeerKeepaliveDeliveryClientModel

				// Construct an instance of the ConfigPeerLimits model
				configPeerLimitsModel := new(blockchainv2.ConfigPeerLimits)
				configPeerLimitsModel.Concurrency = configPeerLimitsConcurrencyModel

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv2.MetricsStatsd)
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the MspConfigData model
				mspConfigDataModel := new(blockchainv2.MspConfigData)
				mspConfigDataModel.Keystore = core.StringPtr("testString")
				mspConfigDataModel.Signcerts = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspConfigDataModel.Cacerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				mspConfigDataModel.Intermediatecerts = []string{"testString"}
				mspConfigDataModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv2.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv2.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ConfigObjectEnrollment model
				configObjectEnrollmentModel := new(blockchainv2.ConfigObjectEnrollment)
				configObjectEnrollmentModel.Component = configObjectEnrollmentComponentModel
				configObjectEnrollmentModel.Tls = configObjectEnrollmentTlsModel

				// Construct an instance of the ConfigObjectMsp model
				configObjectMspModel := new(blockchainv2.ConfigObjectMsp)
				configObjectMspModel.Component = mspConfigDataModel
				configObjectMspModel.Tls = mspConfigDataModel
				configObjectMspModel.Clientauth = mspConfigDataModel

				// Construct an instance of the ConfigPeerChaincode model
				configPeerChaincodeModel := new(blockchainv2.ConfigPeerChaincode)
				configPeerChaincodeModel.Golang = configPeerChaincodeGolangModel
				configPeerChaincodeModel.ExternalBuilders = []blockchainv2.ConfigPeerChaincodeExternalBuildersItem{*configPeerChaincodeExternalBuildersItemModel}
				configPeerChaincodeModel.InstallTimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Startuptimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Executetimeout = core.StringPtr("30s")
				configPeerChaincodeModel.System = configPeerChaincodeSystemModel
				configPeerChaincodeModel.Logging = configPeerChaincodeLoggingModel

				// Construct an instance of the ConfigPeerCreatePeer model
				configPeerCreatePeerModel := new(blockchainv2.ConfigPeerCreatePeer)
				configPeerCreatePeerModel.ID = core.StringPtr("john-doe")
				configPeerCreatePeerModel.NetworkID = core.StringPtr("dev")
				configPeerCreatePeerModel.Keepalive = configPeerKeepaliveModel
				configPeerCreatePeerModel.Gossip = configPeerGossipModel
				configPeerCreatePeerModel.Authentication = configPeerAuthenticationModel
				configPeerCreatePeerModel.BCCSP = bccspModel
				configPeerCreatePeerModel.Client = configPeerClientModel
				configPeerCreatePeerModel.Deliveryclient = configPeerDeliveryclientModel
				configPeerCreatePeerModel.AdminService = configPeerAdminServiceModel
				configPeerCreatePeerModel.ValidatorPoolSize = core.Float64Ptr(float64(8))
				configPeerCreatePeerModel.Discovery = configPeerDiscoveryModel
				configPeerCreatePeerModel.Limits = configPeerLimitsModel

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv2.Metrics)
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv2.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectCouchDb model
				resourceObjectCouchDbModel := new(blockchainv2.ResourceObjectCouchDb)
				resourceObjectCouchDbModel.Requests = resourceRequestsModel
				resourceObjectCouchDbModel.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectFabV1 model
				resourceObjectFabV1Model := new(blockchainv2.ResourceObjectFabV1)
				resourceObjectFabV1Model.Requests = resourceRequestsModel
				resourceObjectFabV1Model.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectFabV2 model
				resourceObjectFabV2Model := new(blockchainv2.ResourceObjectFabV2)
				resourceObjectFabV2Model.Requests = resourceRequestsModel
				resourceObjectFabV2Model.Limits = resourceLimitsModel

				// Construct an instance of the StorageObject model
				storageObjectModel := new(blockchainv2.StorageObject)
				storageObjectModel.Size = core.StringPtr("4GiB")
				storageObjectModel.Class = core.StringPtr("default")

				// Construct an instance of the ConfigObject model
				configObjectModel := new(blockchainv2.ConfigObject)
				configObjectModel.Enrollment = configObjectEnrollmentModel
				configObjectModel.Msp = configObjectMspModel

				// Construct an instance of the ConfigPeerCreate model
				configPeerCreateModel := new(blockchainv2.ConfigPeerCreate)
				configPeerCreateModel.Peer = configPeerCreatePeerModel
				configPeerCreateModel.Chaincode = configPeerChaincodeModel
				configPeerCreateModel.Metrics = metricsModel

				// Construct an instance of the CreatePeerBodyStorage model
				createPeerBodyStorageModel := new(blockchainv2.CreatePeerBodyStorage)
				createPeerBodyStorageModel.Peer = storageObjectModel
				createPeerBodyStorageModel.Statedb = storageObjectModel

				// Construct an instance of the Hsm model
				hsmModel := new(blockchainv2.Hsm)
				hsmModel.Pkcs11endpoint = core.StringPtr("tcp://example.com:666")

				// Construct an instance of the PeerResources model
				peerResourcesModel := new(blockchainv2.PeerResources)
				peerResourcesModel.Chaincodelauncher = resourceObjectFabV2Model
				peerResourcesModel.Couchdb = resourceObjectCouchDbModel
				peerResourcesModel.Statedb = resourceObjectModel
				peerResourcesModel.Dind = resourceObjectFabV1Model
				peerResourcesModel.Fluentd = resourceObjectFabV1Model
				peerResourcesModel.Peer = resourceObjectModel
				peerResourcesModel.Proxy = resourceObjectModel

				// Construct an instance of the CreatePeerOptions model
				createPeerOptionsModel := new(blockchainv2.CreatePeerOptions)
				createPeerOptionsModel.MspID = core.StringPtr("Org1")
				createPeerOptionsModel.DisplayName = core.StringPtr("My Peer")
				createPeerOptionsModel.Config = configObjectModel
				createPeerOptionsModel.ConfigOverride = configPeerCreateModel
				createPeerOptionsModel.Resources = peerResourcesModel
				createPeerOptionsModel.Storage = createPeerBodyStorageModel
				createPeerOptionsModel.Zone = core.StringPtr("testString")
				createPeerOptionsModel.StateDb = core.StringPtr("couchdb")
				createPeerOptionsModel.Tags = []string{"testString"}
				createPeerOptionsModel.Hsm = hsmModel
				createPeerOptionsModel.Region = core.StringPtr("testString")
				createPeerOptionsModel.Version = core.StringPtr("1.4.6-1")
 				createPeerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreatePeer(createPeerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreatePeer with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy model
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel := new(blockchainv2.ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy)
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.RequiredPeerCount = core.Float64Ptr(float64(0))
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.MaxPeerCount = core.Float64Ptr(float64(1))

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv2.BccspPKCS11)
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv2.BccspSW)
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the ConfigObjectEnrollmentComponentCatls model
				configObjectEnrollmentComponentCatlsModel := new(blockchainv2.ConfigObjectEnrollmentComponentCatls)
				configObjectEnrollmentComponentCatlsModel.Cacert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")

				// Construct an instance of the ConfigObjectEnrollmentTlsCatls model
				configObjectEnrollmentTlsCatlsModel := new(blockchainv2.ConfigObjectEnrollmentTlsCatls)
				configObjectEnrollmentTlsCatlsModel.Cacert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")

				// Construct an instance of the ConfigObjectEnrollmentTlsCsr model
				configObjectEnrollmentTlsCsrModel := new(blockchainv2.ConfigObjectEnrollmentTlsCsr)
				configObjectEnrollmentTlsCsrModel.Hosts = []string{"testString"}

				// Construct an instance of the ConfigPeerDeliveryclientAddressOverridesItem model
				configPeerDeliveryclientAddressOverridesItemModel := new(blockchainv2.ConfigPeerDeliveryclientAddressOverridesItem)
				configPeerDeliveryclientAddressOverridesItemModel.From = core.StringPtr("n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.To = core.StringPtr("n3a3ec3-myorderer2.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.CaCertsFile = core.StringPtr("my-data/cert.pem")

				// Construct an instance of the ConfigPeerGossipElection model
				configPeerGossipElectionModel := new(blockchainv2.ConfigPeerGossipElection)
				configPeerGossipElectionModel.StartupGracePeriod = core.StringPtr("15s")
				configPeerGossipElectionModel.MembershipSampleInterval = core.StringPtr("1s")
				configPeerGossipElectionModel.LeaderAliveThreshold = core.StringPtr("10s")
				configPeerGossipElectionModel.LeaderElectionDuration = core.StringPtr("5s")

				// Construct an instance of the ConfigPeerGossipPvtData model
				configPeerGossipPvtDataModel := new(blockchainv2.ConfigPeerGossipPvtData)
				configPeerGossipPvtDataModel.PullRetryThreshold = core.StringPtr("60s")
				configPeerGossipPvtDataModel.TransientstoreMaxBlockRetention = core.Float64Ptr(float64(1000))
				configPeerGossipPvtDataModel.PushAckTimeout = core.StringPtr("3s")
				configPeerGossipPvtDataModel.BtlPullMargin = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileBatchSize = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileSleepInterval = core.StringPtr("1m")
				configPeerGossipPvtDataModel.ReconciliationEnabled = core.BoolPtr(true)
				configPeerGossipPvtDataModel.SkipPullingInvalidTransactionsDuringCommit = core.BoolPtr(false)
				configPeerGossipPvtDataModel.ImplicitCollectionDisseminationPolicy = configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel

				// Construct an instance of the ConfigPeerGossipState model
				configPeerGossipStateModel := new(blockchainv2.ConfigPeerGossipState)
				configPeerGossipStateModel.Enabled = core.BoolPtr(true)
				configPeerGossipStateModel.CheckInterval = core.StringPtr("10s")
				configPeerGossipStateModel.ResponseTimeout = core.StringPtr("3s")
				configPeerGossipStateModel.BatchSize = core.Float64Ptr(float64(10))
				configPeerGossipStateModel.BlockBufferSize = core.Float64Ptr(float64(100))
				configPeerGossipStateModel.MaxRetries = core.Float64Ptr(float64(3))

				// Construct an instance of the ConfigPeerKeepaliveClient model
				configPeerKeepaliveClientModel := new(blockchainv2.ConfigPeerKeepaliveClient)
				configPeerKeepaliveClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveClientModel.Timeout = core.StringPtr("20s")

				// Construct an instance of the ConfigPeerKeepaliveDeliveryClient model
				configPeerKeepaliveDeliveryClientModel := new(blockchainv2.ConfigPeerKeepaliveDeliveryClient)
				configPeerKeepaliveDeliveryClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveDeliveryClientModel.Timeout = core.StringPtr("20s")

				// Construct an instance of the ConfigPeerLimitsConcurrency model
				configPeerLimitsConcurrencyModel := new(blockchainv2.ConfigPeerLimitsConcurrency)
				configPeerLimitsConcurrencyModel.EndorserService = map[string]interface{}{"anyKey": "anyValue"}
				configPeerLimitsConcurrencyModel.DeliverService = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv2.Bccsp)
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model

				// Construct an instance of the ConfigObjectEnrollmentComponent model
				configObjectEnrollmentComponentModel := new(blockchainv2.ConfigObjectEnrollmentComponent)
				configObjectEnrollmentComponentModel.Cahost = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				configObjectEnrollmentComponentModel.Caport = core.Float64Ptr(float64(7054))
				configObjectEnrollmentComponentModel.Caname = core.StringPtr("ca")
				configObjectEnrollmentComponentModel.Catls = configObjectEnrollmentComponentCatlsModel
				configObjectEnrollmentComponentModel.Enrollid = core.StringPtr("admin")
				configObjectEnrollmentComponentModel.Enrollsecret = core.StringPtr("password")
				configObjectEnrollmentComponentModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the ConfigObjectEnrollmentTls model
				configObjectEnrollmentTlsModel := new(blockchainv2.ConfigObjectEnrollmentTls)
				configObjectEnrollmentTlsModel.Cahost = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				configObjectEnrollmentTlsModel.Caport = core.Float64Ptr(float64(7054))
				configObjectEnrollmentTlsModel.Caname = core.StringPtr("tlsca")
				configObjectEnrollmentTlsModel.Catls = configObjectEnrollmentTlsCatlsModel
				configObjectEnrollmentTlsModel.Enrollid = core.StringPtr("admin")
				configObjectEnrollmentTlsModel.Enrollsecret = core.StringPtr("password")
				configObjectEnrollmentTlsModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				configObjectEnrollmentTlsModel.Csr = configObjectEnrollmentTlsCsrModel

				// Construct an instance of the ConfigPeerAdminService model
				configPeerAdminServiceModel := new(blockchainv2.ConfigPeerAdminService)
				configPeerAdminServiceModel.ListenAddress = core.StringPtr("0.0.0.0:7051")

				// Construct an instance of the ConfigPeerAuthentication model
				configPeerAuthenticationModel := new(blockchainv2.ConfigPeerAuthentication)
				configPeerAuthenticationModel.Timewindow = core.StringPtr("15m")

				// Construct an instance of the ConfigPeerChaincodeExternalBuildersItem model
				configPeerChaincodeExternalBuildersItemModel := new(blockchainv2.ConfigPeerChaincodeExternalBuildersItem)
				configPeerChaincodeExternalBuildersItemModel.Path = core.StringPtr("/path/to/directory")
				configPeerChaincodeExternalBuildersItemModel.Name = core.StringPtr("descriptive-build-name")
				configPeerChaincodeExternalBuildersItemModel.EnvironmentWhitelist = []string{"GOPROXY"}

				// Construct an instance of the ConfigPeerChaincodeGolang model
				configPeerChaincodeGolangModel := new(blockchainv2.ConfigPeerChaincodeGolang)
				configPeerChaincodeGolangModel.DynamicLink = core.BoolPtr(false)

				// Construct an instance of the ConfigPeerChaincodeLogging model
				configPeerChaincodeLoggingModel := new(blockchainv2.ConfigPeerChaincodeLogging)
				configPeerChaincodeLoggingModel.Level = core.StringPtr("info")
				configPeerChaincodeLoggingModel.Shim = core.StringPtr("warning")
				configPeerChaincodeLoggingModel.Format = core.StringPtr("%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}")

				// Construct an instance of the ConfigPeerChaincodeSystem model
				configPeerChaincodeSystemModel := new(blockchainv2.ConfigPeerChaincodeSystem)
				configPeerChaincodeSystemModel.Cscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Lscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Escc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Vscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Qscc = core.BoolPtr(true)

				// Construct an instance of the ConfigPeerClient model
				configPeerClientModel := new(blockchainv2.ConfigPeerClient)
				configPeerClientModel.ConnTimeout = core.StringPtr("2s")

				// Construct an instance of the ConfigPeerDeliveryclient model
				configPeerDeliveryclientModel := new(blockchainv2.ConfigPeerDeliveryclient)
				configPeerDeliveryclientModel.ReconnectTotalTimeThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.ConnTimeout = core.StringPtr("2s")
				configPeerDeliveryclientModel.ReConnectBackoffThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.AddressOverrides = []blockchainv2.ConfigPeerDeliveryclientAddressOverridesItem{*configPeerDeliveryclientAddressOverridesItemModel}

				// Construct an instance of the ConfigPeerDiscovery model
				configPeerDiscoveryModel := new(blockchainv2.ConfigPeerDiscovery)
				configPeerDiscoveryModel.Enabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheEnabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheMaxSize = core.Float64Ptr(float64(1000))
				configPeerDiscoveryModel.AuthCachePurgeRetentionRatio = core.Float64Ptr(float64(0.75))
				configPeerDiscoveryModel.OrgMembersAllowedAccess = core.BoolPtr(false)

				// Construct an instance of the ConfigPeerGossip model
				configPeerGossipModel := new(blockchainv2.ConfigPeerGossip)
				configPeerGossipModel.UseLeaderElection = core.BoolPtr(true)
				configPeerGossipModel.OrgLeader = core.BoolPtr(false)
				configPeerGossipModel.MembershipTrackerInterval = core.StringPtr("5s")
				configPeerGossipModel.MaxBlockCountToStore = core.Float64Ptr(float64(100))
				configPeerGossipModel.MaxPropagationBurstLatency = core.StringPtr("10ms")
				configPeerGossipModel.MaxPropagationBurstSize = core.Float64Ptr(float64(10))
				configPeerGossipModel.PropagateIterations = core.Float64Ptr(float64(3))
				configPeerGossipModel.PullInterval = core.StringPtr("4s")
				configPeerGossipModel.PullPeerNum = core.Float64Ptr(float64(3))
				configPeerGossipModel.RequestStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.PublishStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.StateInfoRetentionInterval = core.StringPtr("0s")
				configPeerGossipModel.PublishCertPeriod = core.StringPtr("10s")
				configPeerGossipModel.SkipBlockVerification = core.BoolPtr(false)
				configPeerGossipModel.DialTimeout = core.StringPtr("3s")
				configPeerGossipModel.ConnTimeout = core.StringPtr("2s")
				configPeerGossipModel.RecvBuffSize = core.Float64Ptr(float64(20))
				configPeerGossipModel.SendBuffSize = core.Float64Ptr(float64(200))
				configPeerGossipModel.DigestWaitTime = core.StringPtr("1s")
				configPeerGossipModel.RequestWaitTime = core.StringPtr("1500ms")
				configPeerGossipModel.ResponseWaitTime = core.StringPtr("2s")
				configPeerGossipModel.AliveTimeInterval = core.StringPtr("5s")
				configPeerGossipModel.AliveExpirationTimeout = core.StringPtr("25s")
				configPeerGossipModel.ReconnectInterval = core.StringPtr("25s")
				configPeerGossipModel.Election = configPeerGossipElectionModel
				configPeerGossipModel.PvtData = configPeerGossipPvtDataModel
				configPeerGossipModel.State = configPeerGossipStateModel

				// Construct an instance of the ConfigPeerKeepalive model
				configPeerKeepaliveModel := new(blockchainv2.ConfigPeerKeepalive)
				configPeerKeepaliveModel.MinInterval = core.StringPtr("60s")
				configPeerKeepaliveModel.Client = configPeerKeepaliveClientModel
				configPeerKeepaliveModel.DeliveryClient = configPeerKeepaliveDeliveryClientModel

				// Construct an instance of the ConfigPeerLimits model
				configPeerLimitsModel := new(blockchainv2.ConfigPeerLimits)
				configPeerLimitsModel.Concurrency = configPeerLimitsConcurrencyModel

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv2.MetricsStatsd)
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the MspConfigData model
				mspConfigDataModel := new(blockchainv2.MspConfigData)
				mspConfigDataModel.Keystore = core.StringPtr("testString")
				mspConfigDataModel.Signcerts = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspConfigDataModel.Cacerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				mspConfigDataModel.Intermediatecerts = []string{"testString"}
				mspConfigDataModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv2.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv2.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ConfigObjectEnrollment model
				configObjectEnrollmentModel := new(blockchainv2.ConfigObjectEnrollment)
				configObjectEnrollmentModel.Component = configObjectEnrollmentComponentModel
				configObjectEnrollmentModel.Tls = configObjectEnrollmentTlsModel

				// Construct an instance of the ConfigObjectMsp model
				configObjectMspModel := new(blockchainv2.ConfigObjectMsp)
				configObjectMspModel.Component = mspConfigDataModel
				configObjectMspModel.Tls = mspConfigDataModel
				configObjectMspModel.Clientauth = mspConfigDataModel

				// Construct an instance of the ConfigPeerChaincode model
				configPeerChaincodeModel := new(blockchainv2.ConfigPeerChaincode)
				configPeerChaincodeModel.Golang = configPeerChaincodeGolangModel
				configPeerChaincodeModel.ExternalBuilders = []blockchainv2.ConfigPeerChaincodeExternalBuildersItem{*configPeerChaincodeExternalBuildersItemModel}
				configPeerChaincodeModel.InstallTimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Startuptimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Executetimeout = core.StringPtr("30s")
				configPeerChaincodeModel.System = configPeerChaincodeSystemModel
				configPeerChaincodeModel.Logging = configPeerChaincodeLoggingModel

				// Construct an instance of the ConfigPeerCreatePeer model
				configPeerCreatePeerModel := new(blockchainv2.ConfigPeerCreatePeer)
				configPeerCreatePeerModel.ID = core.StringPtr("john-doe")
				configPeerCreatePeerModel.NetworkID = core.StringPtr("dev")
				configPeerCreatePeerModel.Keepalive = configPeerKeepaliveModel
				configPeerCreatePeerModel.Gossip = configPeerGossipModel
				configPeerCreatePeerModel.Authentication = configPeerAuthenticationModel
				configPeerCreatePeerModel.BCCSP = bccspModel
				configPeerCreatePeerModel.Client = configPeerClientModel
				configPeerCreatePeerModel.Deliveryclient = configPeerDeliveryclientModel
				configPeerCreatePeerModel.AdminService = configPeerAdminServiceModel
				configPeerCreatePeerModel.ValidatorPoolSize = core.Float64Ptr(float64(8))
				configPeerCreatePeerModel.Discovery = configPeerDiscoveryModel
				configPeerCreatePeerModel.Limits = configPeerLimitsModel

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv2.Metrics)
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv2.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectCouchDb model
				resourceObjectCouchDbModel := new(blockchainv2.ResourceObjectCouchDb)
				resourceObjectCouchDbModel.Requests = resourceRequestsModel
				resourceObjectCouchDbModel.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectFabV1 model
				resourceObjectFabV1Model := new(blockchainv2.ResourceObjectFabV1)
				resourceObjectFabV1Model.Requests = resourceRequestsModel
				resourceObjectFabV1Model.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectFabV2 model
				resourceObjectFabV2Model := new(blockchainv2.ResourceObjectFabV2)
				resourceObjectFabV2Model.Requests = resourceRequestsModel
				resourceObjectFabV2Model.Limits = resourceLimitsModel

				// Construct an instance of the StorageObject model
				storageObjectModel := new(blockchainv2.StorageObject)
				storageObjectModel.Size = core.StringPtr("4GiB")
				storageObjectModel.Class = core.StringPtr("default")

				// Construct an instance of the ConfigObject model
				configObjectModel := new(blockchainv2.ConfigObject)
				configObjectModel.Enrollment = configObjectEnrollmentModel
				configObjectModel.Msp = configObjectMspModel

				// Construct an instance of the ConfigPeerCreate model
				configPeerCreateModel := new(blockchainv2.ConfigPeerCreate)
				configPeerCreateModel.Peer = configPeerCreatePeerModel
				configPeerCreateModel.Chaincode = configPeerChaincodeModel
				configPeerCreateModel.Metrics = metricsModel

				// Construct an instance of the CreatePeerBodyStorage model
				createPeerBodyStorageModel := new(blockchainv2.CreatePeerBodyStorage)
				createPeerBodyStorageModel.Peer = storageObjectModel
				createPeerBodyStorageModel.Statedb = storageObjectModel

				// Construct an instance of the Hsm model
				hsmModel := new(blockchainv2.Hsm)
				hsmModel.Pkcs11endpoint = core.StringPtr("tcp://example.com:666")

				// Construct an instance of the PeerResources model
				peerResourcesModel := new(blockchainv2.PeerResources)
				peerResourcesModel.Chaincodelauncher = resourceObjectFabV2Model
				peerResourcesModel.Couchdb = resourceObjectCouchDbModel
				peerResourcesModel.Statedb = resourceObjectModel
				peerResourcesModel.Dind = resourceObjectFabV1Model
				peerResourcesModel.Fluentd = resourceObjectFabV1Model
				peerResourcesModel.Peer = resourceObjectModel
				peerResourcesModel.Proxy = resourceObjectModel

				// Construct an instance of the CreatePeerOptions model
				createPeerOptionsModel := new(blockchainv2.CreatePeerOptions)
				createPeerOptionsModel.MspID = core.StringPtr("Org1")
				createPeerOptionsModel.DisplayName = core.StringPtr("My Peer")
				createPeerOptionsModel.Config = configObjectModel
				createPeerOptionsModel.ConfigOverride = configPeerCreateModel
				createPeerOptionsModel.Resources = peerResourcesModel
				createPeerOptionsModel.Storage = createPeerBodyStorageModel
				createPeerOptionsModel.Zone = core.StringPtr("testString")
				createPeerOptionsModel.StateDb = core.StringPtr("couchdb")
				createPeerOptionsModel.Tags = []string{"testString"}
				createPeerOptionsModel.Hsm = hsmModel
				createPeerOptionsModel.Region = core.StringPtr("testString")
				createPeerOptionsModel.Version = core.StringPtr("1.4.6-1")
				createPeerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreatePeer(createPeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreatePeerOptions model with no property values
				createPeerOptionsModelNew := new(blockchainv2.CreatePeerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreatePeer(createPeerOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ImportPeer(importPeerOptions *ImportPeerOptions) - Operation response error`, func() {
		importPeerPath := "/ak/api/v2/components/fabric-peer"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(importPeerPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ImportPeer with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ImportPeerOptions model
				importPeerOptionsModel := new(blockchainv2.ImportPeerOptions)
				importPeerOptionsModel.DisplayName = core.StringPtr("My Peer")
				importPeerOptionsModel.MspID = core.StringPtr("Org1")
				importPeerOptionsModel.GrpcwpURL = core.StringPtr("https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084")
				importPeerOptionsModel.TlsCaRootCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkRpZmZlcmVudCBkYXRhIGhlcmUgaWYgdGhpcyB3YXMgcmVhbAotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==")
				importPeerOptionsModel.Location = core.StringPtr("ibmcloud")
				importPeerOptionsModel.ApiURL = core.StringPtr("grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051")
				importPeerOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443")
				importPeerOptionsModel.Tags = []string{"testString"}
				importPeerOptionsModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				importPeerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ImportPeer(importPeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ImportPeer(importPeerOptions *ImportPeerOptions)`, func() {
		importPeerPath := "/ak/api/v2/components/fabric-peer"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(importPeerPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "component-1", "dep_component_id": "admin", "type": "fabric-peer", "display_name": "My Peer", "grpcwp_url": "https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "operations_url": "https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443", "msp_id": "Org1", "config_override": {"anyKey": "anyValue"}, "node_ou": {"enabled": true}, "ecert": {"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "cacert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}, "location": "ibmcloud", "state_db": "couchdb", "timestamp": 1537262855753, "resources": {"peer": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "statedb": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"peer": {"size": "4GiB", "class": "default"}, "statedb": {"size": "4GiB", "class": "default"}}, "tags": ["Tags"], "tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "version": "1.4.6-1", "zone": "Zone"}`)
				}))
			})
			It(`Invoke ImportPeer successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ImportPeer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ImportPeerOptions model
				importPeerOptionsModel := new(blockchainv2.ImportPeerOptions)
				importPeerOptionsModel.DisplayName = core.StringPtr("My Peer")
				importPeerOptionsModel.MspID = core.StringPtr("Org1")
				importPeerOptionsModel.GrpcwpURL = core.StringPtr("https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084")
				importPeerOptionsModel.TlsCaRootCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkRpZmZlcmVudCBkYXRhIGhlcmUgaWYgdGhpcyB3YXMgcmVhbAotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==")
				importPeerOptionsModel.Location = core.StringPtr("ibmcloud")
				importPeerOptionsModel.ApiURL = core.StringPtr("grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051")
				importPeerOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443")
				importPeerOptionsModel.Tags = []string{"testString"}
				importPeerOptionsModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
 				importPeerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ImportPeer(importPeerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ImportPeer with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ImportPeerOptions model
				importPeerOptionsModel := new(blockchainv2.ImportPeerOptions)
				importPeerOptionsModel.DisplayName = core.StringPtr("My Peer")
				importPeerOptionsModel.MspID = core.StringPtr("Org1")
				importPeerOptionsModel.GrpcwpURL = core.StringPtr("https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084")
				importPeerOptionsModel.TlsCaRootCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkRpZmZlcmVudCBkYXRhIGhlcmUgaWYgdGhpcyB3YXMgcmVhbAotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==")
				importPeerOptionsModel.Location = core.StringPtr("ibmcloud")
				importPeerOptionsModel.ApiURL = core.StringPtr("grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051")
				importPeerOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443")
				importPeerOptionsModel.Tags = []string{"testString"}
				importPeerOptionsModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				importPeerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ImportPeer(importPeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ImportPeerOptions model with no property values
				importPeerOptionsModelNew := new(blockchainv2.ImportPeerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ImportPeer(importPeerOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`EditPeer(editPeerOptions *EditPeerOptions) - Operation response error`, func() {
		editPeerPath := "/ak/api/v2/components/fabric-peer/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(editPeerPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke EditPeer with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the EditPeerOptions model
				editPeerOptionsModel := new(blockchainv2.EditPeerOptions)
				editPeerOptionsModel.ID = core.StringPtr("testString")
				editPeerOptionsModel.DisplayName = core.StringPtr("My Peer")
				editPeerOptionsModel.ApiURL = core.StringPtr("grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051")
				editPeerOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443")
				editPeerOptionsModel.GrpcwpURL = core.StringPtr("https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084")
				editPeerOptionsModel.MspID = core.StringPtr("Org1")
				editPeerOptionsModel.Location = core.StringPtr("ibmcloud")
				editPeerOptionsModel.Tags = []string{"testString"}
				editPeerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.EditPeer(editPeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`EditPeer(editPeerOptions *EditPeerOptions)`, func() {
		editPeerPath := "/ak/api/v2/components/fabric-peer/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(editPeerPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "component-1", "dep_component_id": "admin", "type": "fabric-peer", "display_name": "My Peer", "grpcwp_url": "https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "operations_url": "https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443", "msp_id": "Org1", "config_override": {"anyKey": "anyValue"}, "node_ou": {"enabled": true}, "ecert": {"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "cacert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}, "location": "ibmcloud", "state_db": "couchdb", "timestamp": 1537262855753, "resources": {"peer": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "statedb": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"peer": {"size": "4GiB", "class": "default"}, "statedb": {"size": "4GiB", "class": "default"}}, "tags": ["Tags"], "tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "version": "1.4.6-1", "zone": "Zone"}`)
				}))
			})
			It(`Invoke EditPeer successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.EditPeer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EditPeerOptions model
				editPeerOptionsModel := new(blockchainv2.EditPeerOptions)
				editPeerOptionsModel.ID = core.StringPtr("testString")
				editPeerOptionsModel.DisplayName = core.StringPtr("My Peer")
				editPeerOptionsModel.ApiURL = core.StringPtr("grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051")
				editPeerOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443")
				editPeerOptionsModel.GrpcwpURL = core.StringPtr("https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084")
				editPeerOptionsModel.MspID = core.StringPtr("Org1")
				editPeerOptionsModel.Location = core.StringPtr("ibmcloud")
				editPeerOptionsModel.Tags = []string{"testString"}
 				editPeerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.EditPeer(editPeerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke EditPeer with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the EditPeerOptions model
				editPeerOptionsModel := new(blockchainv2.EditPeerOptions)
				editPeerOptionsModel.ID = core.StringPtr("testString")
				editPeerOptionsModel.DisplayName = core.StringPtr("My Peer")
				editPeerOptionsModel.ApiURL = core.StringPtr("grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051")
				editPeerOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443")
				editPeerOptionsModel.GrpcwpURL = core.StringPtr("https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084")
				editPeerOptionsModel.MspID = core.StringPtr("Org1")
				editPeerOptionsModel.Location = core.StringPtr("ibmcloud")
				editPeerOptionsModel.Tags = []string{"testString"}
				editPeerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.EditPeer(editPeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the EditPeerOptions model with no property values
				editPeerOptionsModelNew := new(blockchainv2.EditPeerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.EditPeer(editPeerOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdatePeer(updatePeerOptions *UpdatePeerOptions) - Operation response error`, func() {
		updatePeerPath := "/ak/api/v2/kubernetes/components/fabric-peer/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updatePeerPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdatePeer with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy model
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel := new(blockchainv2.ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy)
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.RequiredPeerCount = core.Float64Ptr(float64(0))
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.MaxPeerCount = core.Float64Ptr(float64(1))

				// Construct an instance of the ConfigPeerDeliveryclientAddressOverridesItem model
				configPeerDeliveryclientAddressOverridesItemModel := new(blockchainv2.ConfigPeerDeliveryclientAddressOverridesItem)
				configPeerDeliveryclientAddressOverridesItemModel.From = core.StringPtr("n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.To = core.StringPtr("n3a3ec3-myorderer2.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.CaCertsFile = core.StringPtr("my-data/cert.pem")

				// Construct an instance of the ConfigPeerGossipElection model
				configPeerGossipElectionModel := new(blockchainv2.ConfigPeerGossipElection)
				configPeerGossipElectionModel.StartupGracePeriod = core.StringPtr("15s")
				configPeerGossipElectionModel.MembershipSampleInterval = core.StringPtr("1s")
				configPeerGossipElectionModel.LeaderAliveThreshold = core.StringPtr("10s")
				configPeerGossipElectionModel.LeaderElectionDuration = core.StringPtr("5s")

				// Construct an instance of the ConfigPeerGossipPvtData model
				configPeerGossipPvtDataModel := new(blockchainv2.ConfigPeerGossipPvtData)
				configPeerGossipPvtDataModel.PullRetryThreshold = core.StringPtr("60s")
				configPeerGossipPvtDataModel.TransientstoreMaxBlockRetention = core.Float64Ptr(float64(1000))
				configPeerGossipPvtDataModel.PushAckTimeout = core.StringPtr("3s")
				configPeerGossipPvtDataModel.BtlPullMargin = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileBatchSize = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileSleepInterval = core.StringPtr("1m")
				configPeerGossipPvtDataModel.ReconciliationEnabled = core.BoolPtr(true)
				configPeerGossipPvtDataModel.SkipPullingInvalidTransactionsDuringCommit = core.BoolPtr(false)
				configPeerGossipPvtDataModel.ImplicitCollectionDisseminationPolicy = configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel

				// Construct an instance of the ConfigPeerGossipState model
				configPeerGossipStateModel := new(blockchainv2.ConfigPeerGossipState)
				configPeerGossipStateModel.Enabled = core.BoolPtr(true)
				configPeerGossipStateModel.CheckInterval = core.StringPtr("10s")
				configPeerGossipStateModel.ResponseTimeout = core.StringPtr("3s")
				configPeerGossipStateModel.BatchSize = core.Float64Ptr(float64(10))
				configPeerGossipStateModel.BlockBufferSize = core.Float64Ptr(float64(100))
				configPeerGossipStateModel.MaxRetries = core.Float64Ptr(float64(3))

				// Construct an instance of the ConfigPeerKeepaliveClient model
				configPeerKeepaliveClientModel := new(blockchainv2.ConfigPeerKeepaliveClient)
				configPeerKeepaliveClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveClientModel.Timeout = core.StringPtr("20s")

				// Construct an instance of the ConfigPeerKeepaliveDeliveryClient model
				configPeerKeepaliveDeliveryClientModel := new(blockchainv2.ConfigPeerKeepaliveDeliveryClient)
				configPeerKeepaliveDeliveryClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveDeliveryClientModel.Timeout = core.StringPtr("20s")

				// Construct an instance of the ConfigPeerLimitsConcurrency model
				configPeerLimitsConcurrencyModel := new(blockchainv2.ConfigPeerLimitsConcurrency)
				configPeerLimitsConcurrencyModel.EndorserService = map[string]interface{}{"anyKey": "anyValue"}
				configPeerLimitsConcurrencyModel.DeliverService = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the ConfigPeerAdminService model
				configPeerAdminServiceModel := new(blockchainv2.ConfigPeerAdminService)
				configPeerAdminServiceModel.ListenAddress = core.StringPtr("0.0.0.0:7051")

				// Construct an instance of the ConfigPeerAuthentication model
				configPeerAuthenticationModel := new(blockchainv2.ConfigPeerAuthentication)
				configPeerAuthenticationModel.Timewindow = core.StringPtr("15m")

				// Construct an instance of the ConfigPeerChaincodeExternalBuildersItem model
				configPeerChaincodeExternalBuildersItemModel := new(blockchainv2.ConfigPeerChaincodeExternalBuildersItem)
				configPeerChaincodeExternalBuildersItemModel.Path = core.StringPtr("/path/to/directory")
				configPeerChaincodeExternalBuildersItemModel.Name = core.StringPtr("descriptive-build-name")
				configPeerChaincodeExternalBuildersItemModel.EnvironmentWhitelist = []string{"GOPROXY"}

				// Construct an instance of the ConfigPeerChaincodeGolang model
				configPeerChaincodeGolangModel := new(blockchainv2.ConfigPeerChaincodeGolang)
				configPeerChaincodeGolangModel.DynamicLink = core.BoolPtr(false)

				// Construct an instance of the ConfigPeerChaincodeLogging model
				configPeerChaincodeLoggingModel := new(blockchainv2.ConfigPeerChaincodeLogging)
				configPeerChaincodeLoggingModel.Level = core.StringPtr("info")
				configPeerChaincodeLoggingModel.Shim = core.StringPtr("warning")
				configPeerChaincodeLoggingModel.Format = core.StringPtr("%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}")

				// Construct an instance of the ConfigPeerChaincodeSystem model
				configPeerChaincodeSystemModel := new(blockchainv2.ConfigPeerChaincodeSystem)
				configPeerChaincodeSystemModel.Cscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Lscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Escc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Vscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Qscc = core.BoolPtr(true)

				// Construct an instance of the ConfigPeerClient model
				configPeerClientModel := new(blockchainv2.ConfigPeerClient)
				configPeerClientModel.ConnTimeout = core.StringPtr("2s")

				// Construct an instance of the ConfigPeerDeliveryclient model
				configPeerDeliveryclientModel := new(blockchainv2.ConfigPeerDeliveryclient)
				configPeerDeliveryclientModel.ReconnectTotalTimeThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.ConnTimeout = core.StringPtr("2s")
				configPeerDeliveryclientModel.ReConnectBackoffThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.AddressOverrides = []blockchainv2.ConfigPeerDeliveryclientAddressOverridesItem{*configPeerDeliveryclientAddressOverridesItemModel}

				// Construct an instance of the ConfigPeerDiscovery model
				configPeerDiscoveryModel := new(blockchainv2.ConfigPeerDiscovery)
				configPeerDiscoveryModel.Enabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheEnabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheMaxSize = core.Float64Ptr(float64(1000))
				configPeerDiscoveryModel.AuthCachePurgeRetentionRatio = core.Float64Ptr(float64(0.75))
				configPeerDiscoveryModel.OrgMembersAllowedAccess = core.BoolPtr(false)

				// Construct an instance of the ConfigPeerGossip model
				configPeerGossipModel := new(blockchainv2.ConfigPeerGossip)
				configPeerGossipModel.UseLeaderElection = core.BoolPtr(true)
				configPeerGossipModel.OrgLeader = core.BoolPtr(false)
				configPeerGossipModel.MembershipTrackerInterval = core.StringPtr("5s")
				configPeerGossipModel.MaxBlockCountToStore = core.Float64Ptr(float64(100))
				configPeerGossipModel.MaxPropagationBurstLatency = core.StringPtr("10ms")
				configPeerGossipModel.MaxPropagationBurstSize = core.Float64Ptr(float64(10))
				configPeerGossipModel.PropagateIterations = core.Float64Ptr(float64(3))
				configPeerGossipModel.PullInterval = core.StringPtr("4s")
				configPeerGossipModel.PullPeerNum = core.Float64Ptr(float64(3))
				configPeerGossipModel.RequestStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.PublishStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.StateInfoRetentionInterval = core.StringPtr("0s")
				configPeerGossipModel.PublishCertPeriod = core.StringPtr("10s")
				configPeerGossipModel.SkipBlockVerification = core.BoolPtr(false)
				configPeerGossipModel.DialTimeout = core.StringPtr("3s")
				configPeerGossipModel.ConnTimeout = core.StringPtr("2s")
				configPeerGossipModel.RecvBuffSize = core.Float64Ptr(float64(20))
				configPeerGossipModel.SendBuffSize = core.Float64Ptr(float64(200))
				configPeerGossipModel.DigestWaitTime = core.StringPtr("1s")
				configPeerGossipModel.RequestWaitTime = core.StringPtr("1500ms")
				configPeerGossipModel.ResponseWaitTime = core.StringPtr("2s")
				configPeerGossipModel.AliveTimeInterval = core.StringPtr("5s")
				configPeerGossipModel.AliveExpirationTimeout = core.StringPtr("25s")
				configPeerGossipModel.ReconnectInterval = core.StringPtr("25s")
				configPeerGossipModel.Election = configPeerGossipElectionModel
				configPeerGossipModel.PvtData = configPeerGossipPvtDataModel
				configPeerGossipModel.State = configPeerGossipStateModel

				// Construct an instance of the ConfigPeerKeepalive model
				configPeerKeepaliveModel := new(blockchainv2.ConfigPeerKeepalive)
				configPeerKeepaliveModel.MinInterval = core.StringPtr("60s")
				configPeerKeepaliveModel.Client = configPeerKeepaliveClientModel
				configPeerKeepaliveModel.DeliveryClient = configPeerKeepaliveDeliveryClientModel

				// Construct an instance of the ConfigPeerLimits model
				configPeerLimitsModel := new(blockchainv2.ConfigPeerLimits)
				configPeerLimitsModel.Concurrency = configPeerLimitsConcurrencyModel

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv2.MetricsStatsd)
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv2.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv2.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ConfigPeerChaincode model
				configPeerChaincodeModel := new(blockchainv2.ConfigPeerChaincode)
				configPeerChaincodeModel.Golang = configPeerChaincodeGolangModel
				configPeerChaincodeModel.ExternalBuilders = []blockchainv2.ConfigPeerChaincodeExternalBuildersItem{*configPeerChaincodeExternalBuildersItemModel}
				configPeerChaincodeModel.InstallTimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Startuptimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Executetimeout = core.StringPtr("30s")
				configPeerChaincodeModel.System = configPeerChaincodeSystemModel
				configPeerChaincodeModel.Logging = configPeerChaincodeLoggingModel

				// Construct an instance of the ConfigPeerUpdatePeer model
				configPeerUpdatePeerModel := new(blockchainv2.ConfigPeerUpdatePeer)
				configPeerUpdatePeerModel.ID = core.StringPtr("john-doe")
				configPeerUpdatePeerModel.NetworkID = core.StringPtr("dev")
				configPeerUpdatePeerModel.Keepalive = configPeerKeepaliveModel
				configPeerUpdatePeerModel.Gossip = configPeerGossipModel
				configPeerUpdatePeerModel.Authentication = configPeerAuthenticationModel
				configPeerUpdatePeerModel.Client = configPeerClientModel
				configPeerUpdatePeerModel.Deliveryclient = configPeerDeliveryclientModel
				configPeerUpdatePeerModel.AdminService = configPeerAdminServiceModel
				configPeerUpdatePeerModel.ValidatorPoolSize = core.Float64Ptr(float64(8))
				configPeerUpdatePeerModel.Discovery = configPeerDiscoveryModel
				configPeerUpdatePeerModel.Limits = configPeerLimitsModel

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv2.Metrics)
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv2.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectCouchDb model
				resourceObjectCouchDbModel := new(blockchainv2.ResourceObjectCouchDb)
				resourceObjectCouchDbModel.Requests = resourceRequestsModel
				resourceObjectCouchDbModel.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectFabV1 model
				resourceObjectFabV1Model := new(blockchainv2.ResourceObjectFabV1)
				resourceObjectFabV1Model.Requests = resourceRequestsModel
				resourceObjectFabV1Model.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectFabV2 model
				resourceObjectFabV2Model := new(blockchainv2.ResourceObjectFabV2)
				resourceObjectFabV2Model.Requests = resourceRequestsModel
				resourceObjectFabV2Model.Limits = resourceLimitsModel

				// Construct an instance of the ConfigPeerUpdate model
				configPeerUpdateModel := new(blockchainv2.ConfigPeerUpdate)
				configPeerUpdateModel.Peer = configPeerUpdatePeerModel
				configPeerUpdateModel.Chaincode = configPeerChaincodeModel
				configPeerUpdateModel.Metrics = metricsModel

				// Construct an instance of the PeerResources model
				peerResourcesModel := new(blockchainv2.PeerResources)
				peerResourcesModel.Chaincodelauncher = resourceObjectFabV2Model
				peerResourcesModel.Couchdb = resourceObjectCouchDbModel
				peerResourcesModel.Statedb = resourceObjectModel
				peerResourcesModel.Dind = resourceObjectFabV1Model
				peerResourcesModel.Fluentd = resourceObjectFabV1Model
				peerResourcesModel.Peer = resourceObjectModel
				peerResourcesModel.Proxy = resourceObjectModel

				// Construct an instance of the UpdatePeerOptions model
				updatePeerOptionsModel := new(blockchainv2.UpdatePeerOptions)
				updatePeerOptionsModel.ID = core.StringPtr("testString")
				updatePeerOptionsModel.ConfigOverride = configPeerUpdateModel
				updatePeerOptionsModel.Resources = peerResourcesModel
				updatePeerOptionsModel.Zone = core.StringPtr("testString")
				updatePeerOptionsModel.Version = core.StringPtr("1.4.6-1")
				updatePeerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdatePeer(updatePeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdatePeer(updatePeerOptions *UpdatePeerOptions)`, func() {
		updatePeerPath := "/ak/api/v2/kubernetes/components/fabric-peer/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updatePeerPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "component-1", "dep_component_id": "admin", "type": "fabric-peer", "display_name": "My Peer", "grpcwp_url": "https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "operations_url": "https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443", "msp_id": "Org1", "config_override": {"anyKey": "anyValue"}, "node_ou": {"enabled": true}, "ecert": {"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "cacert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}, "location": "ibmcloud", "state_db": "couchdb", "timestamp": 1537262855753, "resources": {"peer": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "statedb": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"peer": {"size": "4GiB", "class": "default"}, "statedb": {"size": "4GiB", "class": "default"}}, "tags": ["Tags"], "tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "version": "1.4.6-1", "zone": "Zone"}`)
				}))
			})
			It(`Invoke UpdatePeer successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdatePeer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy model
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel := new(blockchainv2.ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy)
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.RequiredPeerCount = core.Float64Ptr(float64(0))
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.MaxPeerCount = core.Float64Ptr(float64(1))

				// Construct an instance of the ConfigPeerDeliveryclientAddressOverridesItem model
				configPeerDeliveryclientAddressOverridesItemModel := new(blockchainv2.ConfigPeerDeliveryclientAddressOverridesItem)
				configPeerDeliveryclientAddressOverridesItemModel.From = core.StringPtr("n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.To = core.StringPtr("n3a3ec3-myorderer2.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.CaCertsFile = core.StringPtr("my-data/cert.pem")

				// Construct an instance of the ConfigPeerGossipElection model
				configPeerGossipElectionModel := new(blockchainv2.ConfigPeerGossipElection)
				configPeerGossipElectionModel.StartupGracePeriod = core.StringPtr("15s")
				configPeerGossipElectionModel.MembershipSampleInterval = core.StringPtr("1s")
				configPeerGossipElectionModel.LeaderAliveThreshold = core.StringPtr("10s")
				configPeerGossipElectionModel.LeaderElectionDuration = core.StringPtr("5s")

				// Construct an instance of the ConfigPeerGossipPvtData model
				configPeerGossipPvtDataModel := new(blockchainv2.ConfigPeerGossipPvtData)
				configPeerGossipPvtDataModel.PullRetryThreshold = core.StringPtr("60s")
				configPeerGossipPvtDataModel.TransientstoreMaxBlockRetention = core.Float64Ptr(float64(1000))
				configPeerGossipPvtDataModel.PushAckTimeout = core.StringPtr("3s")
				configPeerGossipPvtDataModel.BtlPullMargin = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileBatchSize = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileSleepInterval = core.StringPtr("1m")
				configPeerGossipPvtDataModel.ReconciliationEnabled = core.BoolPtr(true)
				configPeerGossipPvtDataModel.SkipPullingInvalidTransactionsDuringCommit = core.BoolPtr(false)
				configPeerGossipPvtDataModel.ImplicitCollectionDisseminationPolicy = configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel

				// Construct an instance of the ConfigPeerGossipState model
				configPeerGossipStateModel := new(blockchainv2.ConfigPeerGossipState)
				configPeerGossipStateModel.Enabled = core.BoolPtr(true)
				configPeerGossipStateModel.CheckInterval = core.StringPtr("10s")
				configPeerGossipStateModel.ResponseTimeout = core.StringPtr("3s")
				configPeerGossipStateModel.BatchSize = core.Float64Ptr(float64(10))
				configPeerGossipStateModel.BlockBufferSize = core.Float64Ptr(float64(100))
				configPeerGossipStateModel.MaxRetries = core.Float64Ptr(float64(3))

				// Construct an instance of the ConfigPeerKeepaliveClient model
				configPeerKeepaliveClientModel := new(blockchainv2.ConfigPeerKeepaliveClient)
				configPeerKeepaliveClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveClientModel.Timeout = core.StringPtr("20s")

				// Construct an instance of the ConfigPeerKeepaliveDeliveryClient model
				configPeerKeepaliveDeliveryClientModel := new(blockchainv2.ConfigPeerKeepaliveDeliveryClient)
				configPeerKeepaliveDeliveryClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveDeliveryClientModel.Timeout = core.StringPtr("20s")

				// Construct an instance of the ConfigPeerLimitsConcurrency model
				configPeerLimitsConcurrencyModel := new(blockchainv2.ConfigPeerLimitsConcurrency)
				configPeerLimitsConcurrencyModel.EndorserService = map[string]interface{}{"anyKey": "anyValue"}
				configPeerLimitsConcurrencyModel.DeliverService = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the ConfigPeerAdminService model
				configPeerAdminServiceModel := new(blockchainv2.ConfigPeerAdminService)
				configPeerAdminServiceModel.ListenAddress = core.StringPtr("0.0.0.0:7051")

				// Construct an instance of the ConfigPeerAuthentication model
				configPeerAuthenticationModel := new(blockchainv2.ConfigPeerAuthentication)
				configPeerAuthenticationModel.Timewindow = core.StringPtr("15m")

				// Construct an instance of the ConfigPeerChaincodeExternalBuildersItem model
				configPeerChaincodeExternalBuildersItemModel := new(blockchainv2.ConfigPeerChaincodeExternalBuildersItem)
				configPeerChaincodeExternalBuildersItemModel.Path = core.StringPtr("/path/to/directory")
				configPeerChaincodeExternalBuildersItemModel.Name = core.StringPtr("descriptive-build-name")
				configPeerChaincodeExternalBuildersItemModel.EnvironmentWhitelist = []string{"GOPROXY"}

				// Construct an instance of the ConfigPeerChaincodeGolang model
				configPeerChaincodeGolangModel := new(blockchainv2.ConfigPeerChaincodeGolang)
				configPeerChaincodeGolangModel.DynamicLink = core.BoolPtr(false)

				// Construct an instance of the ConfigPeerChaincodeLogging model
				configPeerChaincodeLoggingModel := new(blockchainv2.ConfigPeerChaincodeLogging)
				configPeerChaincodeLoggingModel.Level = core.StringPtr("info")
				configPeerChaincodeLoggingModel.Shim = core.StringPtr("warning")
				configPeerChaincodeLoggingModel.Format = core.StringPtr("%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}")

				// Construct an instance of the ConfigPeerChaincodeSystem model
				configPeerChaincodeSystemModel := new(blockchainv2.ConfigPeerChaincodeSystem)
				configPeerChaincodeSystemModel.Cscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Lscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Escc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Vscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Qscc = core.BoolPtr(true)

				// Construct an instance of the ConfigPeerClient model
				configPeerClientModel := new(blockchainv2.ConfigPeerClient)
				configPeerClientModel.ConnTimeout = core.StringPtr("2s")

				// Construct an instance of the ConfigPeerDeliveryclient model
				configPeerDeliveryclientModel := new(blockchainv2.ConfigPeerDeliveryclient)
				configPeerDeliveryclientModel.ReconnectTotalTimeThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.ConnTimeout = core.StringPtr("2s")
				configPeerDeliveryclientModel.ReConnectBackoffThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.AddressOverrides = []blockchainv2.ConfigPeerDeliveryclientAddressOverridesItem{*configPeerDeliveryclientAddressOverridesItemModel}

				// Construct an instance of the ConfigPeerDiscovery model
				configPeerDiscoveryModel := new(blockchainv2.ConfigPeerDiscovery)
				configPeerDiscoveryModel.Enabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheEnabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheMaxSize = core.Float64Ptr(float64(1000))
				configPeerDiscoveryModel.AuthCachePurgeRetentionRatio = core.Float64Ptr(float64(0.75))
				configPeerDiscoveryModel.OrgMembersAllowedAccess = core.BoolPtr(false)

				// Construct an instance of the ConfigPeerGossip model
				configPeerGossipModel := new(blockchainv2.ConfigPeerGossip)
				configPeerGossipModel.UseLeaderElection = core.BoolPtr(true)
				configPeerGossipModel.OrgLeader = core.BoolPtr(false)
				configPeerGossipModel.MembershipTrackerInterval = core.StringPtr("5s")
				configPeerGossipModel.MaxBlockCountToStore = core.Float64Ptr(float64(100))
				configPeerGossipModel.MaxPropagationBurstLatency = core.StringPtr("10ms")
				configPeerGossipModel.MaxPropagationBurstSize = core.Float64Ptr(float64(10))
				configPeerGossipModel.PropagateIterations = core.Float64Ptr(float64(3))
				configPeerGossipModel.PullInterval = core.StringPtr("4s")
				configPeerGossipModel.PullPeerNum = core.Float64Ptr(float64(3))
				configPeerGossipModel.RequestStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.PublishStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.StateInfoRetentionInterval = core.StringPtr("0s")
				configPeerGossipModel.PublishCertPeriod = core.StringPtr("10s")
				configPeerGossipModel.SkipBlockVerification = core.BoolPtr(false)
				configPeerGossipModel.DialTimeout = core.StringPtr("3s")
				configPeerGossipModel.ConnTimeout = core.StringPtr("2s")
				configPeerGossipModel.RecvBuffSize = core.Float64Ptr(float64(20))
				configPeerGossipModel.SendBuffSize = core.Float64Ptr(float64(200))
				configPeerGossipModel.DigestWaitTime = core.StringPtr("1s")
				configPeerGossipModel.RequestWaitTime = core.StringPtr("1500ms")
				configPeerGossipModel.ResponseWaitTime = core.StringPtr("2s")
				configPeerGossipModel.AliveTimeInterval = core.StringPtr("5s")
				configPeerGossipModel.AliveExpirationTimeout = core.StringPtr("25s")
				configPeerGossipModel.ReconnectInterval = core.StringPtr("25s")
				configPeerGossipModel.Election = configPeerGossipElectionModel
				configPeerGossipModel.PvtData = configPeerGossipPvtDataModel
				configPeerGossipModel.State = configPeerGossipStateModel

				// Construct an instance of the ConfigPeerKeepalive model
				configPeerKeepaliveModel := new(blockchainv2.ConfigPeerKeepalive)
				configPeerKeepaliveModel.MinInterval = core.StringPtr("60s")
				configPeerKeepaliveModel.Client = configPeerKeepaliveClientModel
				configPeerKeepaliveModel.DeliveryClient = configPeerKeepaliveDeliveryClientModel

				// Construct an instance of the ConfigPeerLimits model
				configPeerLimitsModel := new(blockchainv2.ConfigPeerLimits)
				configPeerLimitsModel.Concurrency = configPeerLimitsConcurrencyModel

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv2.MetricsStatsd)
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv2.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv2.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ConfigPeerChaincode model
				configPeerChaincodeModel := new(blockchainv2.ConfigPeerChaincode)
				configPeerChaincodeModel.Golang = configPeerChaincodeGolangModel
				configPeerChaincodeModel.ExternalBuilders = []blockchainv2.ConfigPeerChaincodeExternalBuildersItem{*configPeerChaincodeExternalBuildersItemModel}
				configPeerChaincodeModel.InstallTimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Startuptimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Executetimeout = core.StringPtr("30s")
				configPeerChaincodeModel.System = configPeerChaincodeSystemModel
				configPeerChaincodeModel.Logging = configPeerChaincodeLoggingModel

				// Construct an instance of the ConfigPeerUpdatePeer model
				configPeerUpdatePeerModel := new(blockchainv2.ConfigPeerUpdatePeer)
				configPeerUpdatePeerModel.ID = core.StringPtr("john-doe")
				configPeerUpdatePeerModel.NetworkID = core.StringPtr("dev")
				configPeerUpdatePeerModel.Keepalive = configPeerKeepaliveModel
				configPeerUpdatePeerModel.Gossip = configPeerGossipModel
				configPeerUpdatePeerModel.Authentication = configPeerAuthenticationModel
				configPeerUpdatePeerModel.Client = configPeerClientModel
				configPeerUpdatePeerModel.Deliveryclient = configPeerDeliveryclientModel
				configPeerUpdatePeerModel.AdminService = configPeerAdminServiceModel
				configPeerUpdatePeerModel.ValidatorPoolSize = core.Float64Ptr(float64(8))
				configPeerUpdatePeerModel.Discovery = configPeerDiscoveryModel
				configPeerUpdatePeerModel.Limits = configPeerLimitsModel

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv2.Metrics)
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv2.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectCouchDb model
				resourceObjectCouchDbModel := new(blockchainv2.ResourceObjectCouchDb)
				resourceObjectCouchDbModel.Requests = resourceRequestsModel
				resourceObjectCouchDbModel.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectFabV1 model
				resourceObjectFabV1Model := new(blockchainv2.ResourceObjectFabV1)
				resourceObjectFabV1Model.Requests = resourceRequestsModel
				resourceObjectFabV1Model.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectFabV2 model
				resourceObjectFabV2Model := new(blockchainv2.ResourceObjectFabV2)
				resourceObjectFabV2Model.Requests = resourceRequestsModel
				resourceObjectFabV2Model.Limits = resourceLimitsModel

				// Construct an instance of the ConfigPeerUpdate model
				configPeerUpdateModel := new(blockchainv2.ConfigPeerUpdate)
				configPeerUpdateModel.Peer = configPeerUpdatePeerModel
				configPeerUpdateModel.Chaincode = configPeerChaincodeModel
				configPeerUpdateModel.Metrics = metricsModel

				// Construct an instance of the PeerResources model
				peerResourcesModel := new(blockchainv2.PeerResources)
				peerResourcesModel.Chaincodelauncher = resourceObjectFabV2Model
				peerResourcesModel.Couchdb = resourceObjectCouchDbModel
				peerResourcesModel.Statedb = resourceObjectModel
				peerResourcesModel.Dind = resourceObjectFabV1Model
				peerResourcesModel.Fluentd = resourceObjectFabV1Model
				peerResourcesModel.Peer = resourceObjectModel
				peerResourcesModel.Proxy = resourceObjectModel

				// Construct an instance of the UpdatePeerOptions model
				updatePeerOptionsModel := new(blockchainv2.UpdatePeerOptions)
				updatePeerOptionsModel.ID = core.StringPtr("testString")
				updatePeerOptionsModel.ConfigOverride = configPeerUpdateModel
				updatePeerOptionsModel.Resources = peerResourcesModel
				updatePeerOptionsModel.Zone = core.StringPtr("testString")
				updatePeerOptionsModel.Version = core.StringPtr("1.4.6-1")
 				updatePeerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdatePeer(updatePeerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdatePeer with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy model
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel := new(blockchainv2.ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy)
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.RequiredPeerCount = core.Float64Ptr(float64(0))
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.MaxPeerCount = core.Float64Ptr(float64(1))

				// Construct an instance of the ConfigPeerDeliveryclientAddressOverridesItem model
				configPeerDeliveryclientAddressOverridesItemModel := new(blockchainv2.ConfigPeerDeliveryclientAddressOverridesItem)
				configPeerDeliveryclientAddressOverridesItemModel.From = core.StringPtr("n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.To = core.StringPtr("n3a3ec3-myorderer2.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.CaCertsFile = core.StringPtr("my-data/cert.pem")

				// Construct an instance of the ConfigPeerGossipElection model
				configPeerGossipElectionModel := new(blockchainv2.ConfigPeerGossipElection)
				configPeerGossipElectionModel.StartupGracePeriod = core.StringPtr("15s")
				configPeerGossipElectionModel.MembershipSampleInterval = core.StringPtr("1s")
				configPeerGossipElectionModel.LeaderAliveThreshold = core.StringPtr("10s")
				configPeerGossipElectionModel.LeaderElectionDuration = core.StringPtr("5s")

				// Construct an instance of the ConfigPeerGossipPvtData model
				configPeerGossipPvtDataModel := new(blockchainv2.ConfigPeerGossipPvtData)
				configPeerGossipPvtDataModel.PullRetryThreshold = core.StringPtr("60s")
				configPeerGossipPvtDataModel.TransientstoreMaxBlockRetention = core.Float64Ptr(float64(1000))
				configPeerGossipPvtDataModel.PushAckTimeout = core.StringPtr("3s")
				configPeerGossipPvtDataModel.BtlPullMargin = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileBatchSize = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileSleepInterval = core.StringPtr("1m")
				configPeerGossipPvtDataModel.ReconciliationEnabled = core.BoolPtr(true)
				configPeerGossipPvtDataModel.SkipPullingInvalidTransactionsDuringCommit = core.BoolPtr(false)
				configPeerGossipPvtDataModel.ImplicitCollectionDisseminationPolicy = configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel

				// Construct an instance of the ConfigPeerGossipState model
				configPeerGossipStateModel := new(blockchainv2.ConfigPeerGossipState)
				configPeerGossipStateModel.Enabled = core.BoolPtr(true)
				configPeerGossipStateModel.CheckInterval = core.StringPtr("10s")
				configPeerGossipStateModel.ResponseTimeout = core.StringPtr("3s")
				configPeerGossipStateModel.BatchSize = core.Float64Ptr(float64(10))
				configPeerGossipStateModel.BlockBufferSize = core.Float64Ptr(float64(100))
				configPeerGossipStateModel.MaxRetries = core.Float64Ptr(float64(3))

				// Construct an instance of the ConfigPeerKeepaliveClient model
				configPeerKeepaliveClientModel := new(blockchainv2.ConfigPeerKeepaliveClient)
				configPeerKeepaliveClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveClientModel.Timeout = core.StringPtr("20s")

				// Construct an instance of the ConfigPeerKeepaliveDeliveryClient model
				configPeerKeepaliveDeliveryClientModel := new(blockchainv2.ConfigPeerKeepaliveDeliveryClient)
				configPeerKeepaliveDeliveryClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveDeliveryClientModel.Timeout = core.StringPtr("20s")

				// Construct an instance of the ConfigPeerLimitsConcurrency model
				configPeerLimitsConcurrencyModel := new(blockchainv2.ConfigPeerLimitsConcurrency)
				configPeerLimitsConcurrencyModel.EndorserService = map[string]interface{}{"anyKey": "anyValue"}
				configPeerLimitsConcurrencyModel.DeliverService = map[string]interface{}{"anyKey": "anyValue"}

				// Construct an instance of the ConfigPeerAdminService model
				configPeerAdminServiceModel := new(blockchainv2.ConfigPeerAdminService)
				configPeerAdminServiceModel.ListenAddress = core.StringPtr("0.0.0.0:7051")

				// Construct an instance of the ConfigPeerAuthentication model
				configPeerAuthenticationModel := new(blockchainv2.ConfigPeerAuthentication)
				configPeerAuthenticationModel.Timewindow = core.StringPtr("15m")

				// Construct an instance of the ConfigPeerChaincodeExternalBuildersItem model
				configPeerChaincodeExternalBuildersItemModel := new(blockchainv2.ConfigPeerChaincodeExternalBuildersItem)
				configPeerChaincodeExternalBuildersItemModel.Path = core.StringPtr("/path/to/directory")
				configPeerChaincodeExternalBuildersItemModel.Name = core.StringPtr("descriptive-build-name")
				configPeerChaincodeExternalBuildersItemModel.EnvironmentWhitelist = []string{"GOPROXY"}

				// Construct an instance of the ConfigPeerChaincodeGolang model
				configPeerChaincodeGolangModel := new(blockchainv2.ConfigPeerChaincodeGolang)
				configPeerChaincodeGolangModel.DynamicLink = core.BoolPtr(false)

				// Construct an instance of the ConfigPeerChaincodeLogging model
				configPeerChaincodeLoggingModel := new(blockchainv2.ConfigPeerChaincodeLogging)
				configPeerChaincodeLoggingModel.Level = core.StringPtr("info")
				configPeerChaincodeLoggingModel.Shim = core.StringPtr("warning")
				configPeerChaincodeLoggingModel.Format = core.StringPtr("%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}")

				// Construct an instance of the ConfigPeerChaincodeSystem model
				configPeerChaincodeSystemModel := new(blockchainv2.ConfigPeerChaincodeSystem)
				configPeerChaincodeSystemModel.Cscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Lscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Escc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Vscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Qscc = core.BoolPtr(true)

				// Construct an instance of the ConfigPeerClient model
				configPeerClientModel := new(blockchainv2.ConfigPeerClient)
				configPeerClientModel.ConnTimeout = core.StringPtr("2s")

				// Construct an instance of the ConfigPeerDeliveryclient model
				configPeerDeliveryclientModel := new(blockchainv2.ConfigPeerDeliveryclient)
				configPeerDeliveryclientModel.ReconnectTotalTimeThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.ConnTimeout = core.StringPtr("2s")
				configPeerDeliveryclientModel.ReConnectBackoffThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.AddressOverrides = []blockchainv2.ConfigPeerDeliveryclientAddressOverridesItem{*configPeerDeliveryclientAddressOverridesItemModel}

				// Construct an instance of the ConfigPeerDiscovery model
				configPeerDiscoveryModel := new(blockchainv2.ConfigPeerDiscovery)
				configPeerDiscoveryModel.Enabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheEnabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheMaxSize = core.Float64Ptr(float64(1000))
				configPeerDiscoveryModel.AuthCachePurgeRetentionRatio = core.Float64Ptr(float64(0.75))
				configPeerDiscoveryModel.OrgMembersAllowedAccess = core.BoolPtr(false)

				// Construct an instance of the ConfigPeerGossip model
				configPeerGossipModel := new(blockchainv2.ConfigPeerGossip)
				configPeerGossipModel.UseLeaderElection = core.BoolPtr(true)
				configPeerGossipModel.OrgLeader = core.BoolPtr(false)
				configPeerGossipModel.MembershipTrackerInterval = core.StringPtr("5s")
				configPeerGossipModel.MaxBlockCountToStore = core.Float64Ptr(float64(100))
				configPeerGossipModel.MaxPropagationBurstLatency = core.StringPtr("10ms")
				configPeerGossipModel.MaxPropagationBurstSize = core.Float64Ptr(float64(10))
				configPeerGossipModel.PropagateIterations = core.Float64Ptr(float64(3))
				configPeerGossipModel.PullInterval = core.StringPtr("4s")
				configPeerGossipModel.PullPeerNum = core.Float64Ptr(float64(3))
				configPeerGossipModel.RequestStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.PublishStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.StateInfoRetentionInterval = core.StringPtr("0s")
				configPeerGossipModel.PublishCertPeriod = core.StringPtr("10s")
				configPeerGossipModel.SkipBlockVerification = core.BoolPtr(false)
				configPeerGossipModel.DialTimeout = core.StringPtr("3s")
				configPeerGossipModel.ConnTimeout = core.StringPtr("2s")
				configPeerGossipModel.RecvBuffSize = core.Float64Ptr(float64(20))
				configPeerGossipModel.SendBuffSize = core.Float64Ptr(float64(200))
				configPeerGossipModel.DigestWaitTime = core.StringPtr("1s")
				configPeerGossipModel.RequestWaitTime = core.StringPtr("1500ms")
				configPeerGossipModel.ResponseWaitTime = core.StringPtr("2s")
				configPeerGossipModel.AliveTimeInterval = core.StringPtr("5s")
				configPeerGossipModel.AliveExpirationTimeout = core.StringPtr("25s")
				configPeerGossipModel.ReconnectInterval = core.StringPtr("25s")
				configPeerGossipModel.Election = configPeerGossipElectionModel
				configPeerGossipModel.PvtData = configPeerGossipPvtDataModel
				configPeerGossipModel.State = configPeerGossipStateModel

				// Construct an instance of the ConfigPeerKeepalive model
				configPeerKeepaliveModel := new(blockchainv2.ConfigPeerKeepalive)
				configPeerKeepaliveModel.MinInterval = core.StringPtr("60s")
				configPeerKeepaliveModel.Client = configPeerKeepaliveClientModel
				configPeerKeepaliveModel.DeliveryClient = configPeerKeepaliveDeliveryClientModel

				// Construct an instance of the ConfigPeerLimits model
				configPeerLimitsModel := new(blockchainv2.ConfigPeerLimits)
				configPeerLimitsModel.Concurrency = configPeerLimitsConcurrencyModel

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv2.MetricsStatsd)
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv2.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv2.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ConfigPeerChaincode model
				configPeerChaincodeModel := new(blockchainv2.ConfigPeerChaincode)
				configPeerChaincodeModel.Golang = configPeerChaincodeGolangModel
				configPeerChaincodeModel.ExternalBuilders = []blockchainv2.ConfigPeerChaincodeExternalBuildersItem{*configPeerChaincodeExternalBuildersItemModel}
				configPeerChaincodeModel.InstallTimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Startuptimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Executetimeout = core.StringPtr("30s")
				configPeerChaincodeModel.System = configPeerChaincodeSystemModel
				configPeerChaincodeModel.Logging = configPeerChaincodeLoggingModel

				// Construct an instance of the ConfigPeerUpdatePeer model
				configPeerUpdatePeerModel := new(blockchainv2.ConfigPeerUpdatePeer)
				configPeerUpdatePeerModel.ID = core.StringPtr("john-doe")
				configPeerUpdatePeerModel.NetworkID = core.StringPtr("dev")
				configPeerUpdatePeerModel.Keepalive = configPeerKeepaliveModel
				configPeerUpdatePeerModel.Gossip = configPeerGossipModel
				configPeerUpdatePeerModel.Authentication = configPeerAuthenticationModel
				configPeerUpdatePeerModel.Client = configPeerClientModel
				configPeerUpdatePeerModel.Deliveryclient = configPeerDeliveryclientModel
				configPeerUpdatePeerModel.AdminService = configPeerAdminServiceModel
				configPeerUpdatePeerModel.ValidatorPoolSize = core.Float64Ptr(float64(8))
				configPeerUpdatePeerModel.Discovery = configPeerDiscoveryModel
				configPeerUpdatePeerModel.Limits = configPeerLimitsModel

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv2.Metrics)
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv2.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectCouchDb model
				resourceObjectCouchDbModel := new(blockchainv2.ResourceObjectCouchDb)
				resourceObjectCouchDbModel.Requests = resourceRequestsModel
				resourceObjectCouchDbModel.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectFabV1 model
				resourceObjectFabV1Model := new(blockchainv2.ResourceObjectFabV1)
				resourceObjectFabV1Model.Requests = resourceRequestsModel
				resourceObjectFabV1Model.Limits = resourceLimitsModel

				// Construct an instance of the ResourceObjectFabV2 model
				resourceObjectFabV2Model := new(blockchainv2.ResourceObjectFabV2)
				resourceObjectFabV2Model.Requests = resourceRequestsModel
				resourceObjectFabV2Model.Limits = resourceLimitsModel

				// Construct an instance of the ConfigPeerUpdate model
				configPeerUpdateModel := new(blockchainv2.ConfigPeerUpdate)
				configPeerUpdateModel.Peer = configPeerUpdatePeerModel
				configPeerUpdateModel.Chaincode = configPeerChaincodeModel
				configPeerUpdateModel.Metrics = metricsModel

				// Construct an instance of the PeerResources model
				peerResourcesModel := new(blockchainv2.PeerResources)
				peerResourcesModel.Chaincodelauncher = resourceObjectFabV2Model
				peerResourcesModel.Couchdb = resourceObjectCouchDbModel
				peerResourcesModel.Statedb = resourceObjectModel
				peerResourcesModel.Dind = resourceObjectFabV1Model
				peerResourcesModel.Fluentd = resourceObjectFabV1Model
				peerResourcesModel.Peer = resourceObjectModel
				peerResourcesModel.Proxy = resourceObjectModel

				// Construct an instance of the UpdatePeerOptions model
				updatePeerOptionsModel := new(blockchainv2.UpdatePeerOptions)
				updatePeerOptionsModel.ID = core.StringPtr("testString")
				updatePeerOptionsModel.ConfigOverride = configPeerUpdateModel
				updatePeerOptionsModel.Resources = peerResourcesModel
				updatePeerOptionsModel.Zone = core.StringPtr("testString")
				updatePeerOptionsModel.Version = core.StringPtr("1.4.6-1")
				updatePeerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdatePeer(updatePeerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdatePeerOptions model with no property values
				updatePeerOptionsModelNew := new(blockchainv2.UpdatePeerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdatePeer(updatePeerOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateOrderer(createOrdererOptions *CreateOrdererOptions) - Operation response error`, func() {
		createOrdererPath := "/ak/api/v2/kubernetes/components/fabric-orderer"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createOrdererPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateOrderer with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv2.BccspPKCS11)
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv2.BccspSW)
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the ConfigObjectEnrollmentComponentCatls model
				configObjectEnrollmentComponentCatlsModel := new(blockchainv2.ConfigObjectEnrollmentComponentCatls)
				configObjectEnrollmentComponentCatlsModel.Cacert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")

				// Construct an instance of the ConfigObjectEnrollmentTlsCatls model
				configObjectEnrollmentTlsCatlsModel := new(blockchainv2.ConfigObjectEnrollmentTlsCatls)
				configObjectEnrollmentTlsCatlsModel.Cacert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")

				// Construct an instance of the ConfigObjectEnrollmentTlsCsr model
				configObjectEnrollmentTlsCsrModel := new(blockchainv2.ConfigObjectEnrollmentTlsCsr)
				configObjectEnrollmentTlsCsrModel.Hosts = []string{"testString"}

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv2.Bccsp)
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model

				// Construct an instance of the ConfigObjectEnrollmentComponent model
				configObjectEnrollmentComponentModel := new(blockchainv2.ConfigObjectEnrollmentComponent)
				configObjectEnrollmentComponentModel.Cahost = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				configObjectEnrollmentComponentModel.Caport = core.Float64Ptr(float64(7054))
				configObjectEnrollmentComponentModel.Caname = core.StringPtr("ca")
				configObjectEnrollmentComponentModel.Catls = configObjectEnrollmentComponentCatlsModel
				configObjectEnrollmentComponentModel.Enrollid = core.StringPtr("admin")
				configObjectEnrollmentComponentModel.Enrollsecret = core.StringPtr("password")
				configObjectEnrollmentComponentModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the ConfigObjectEnrollmentTls model
				configObjectEnrollmentTlsModel := new(blockchainv2.ConfigObjectEnrollmentTls)
				configObjectEnrollmentTlsModel.Cahost = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				configObjectEnrollmentTlsModel.Caport = core.Float64Ptr(float64(7054))
				configObjectEnrollmentTlsModel.Caname = core.StringPtr("tlsca")
				configObjectEnrollmentTlsModel.Catls = configObjectEnrollmentTlsCatlsModel
				configObjectEnrollmentTlsModel.Enrollid = core.StringPtr("admin")
				configObjectEnrollmentTlsModel.Enrollsecret = core.StringPtr("password")
				configObjectEnrollmentTlsModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				configObjectEnrollmentTlsModel.Csr = configObjectEnrollmentTlsCsrModel

				// Construct an instance of the ConfigOrdererAuthentication model
				configOrdererAuthenticationModel := new(blockchainv2.ConfigOrdererAuthentication)
				configOrdererAuthenticationModel.TimeWindow = core.StringPtr("15m")
				configOrdererAuthenticationModel.NoExpirationChecks = core.BoolPtr(false)

				// Construct an instance of the ConfigOrdererKeepalive model
				configOrdererKeepaliveModel := new(blockchainv2.ConfigOrdererKeepalive)
				configOrdererKeepaliveModel.ServerMinInterval = core.StringPtr("60s")
				configOrdererKeepaliveModel.ServerInterval = core.StringPtr("2h")
				configOrdererKeepaliveModel.ServerTimeout = core.StringPtr("20s")

				// Construct an instance of the ConfigOrdererMetricsStatsd model
				configOrdererMetricsStatsdModel := new(blockchainv2.ConfigOrdererMetricsStatsd)
				configOrdererMetricsStatsdModel.Network = core.StringPtr("udp")
				configOrdererMetricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				configOrdererMetricsStatsdModel.WriteInterval = core.StringPtr("10s")
				configOrdererMetricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the MspConfigData model
				mspConfigDataModel := new(blockchainv2.MspConfigData)
				mspConfigDataModel.Keystore = core.StringPtr("testString")
				mspConfigDataModel.Signcerts = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspConfigDataModel.Cacerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				mspConfigDataModel.Intermediatecerts = []string{"testString"}
				mspConfigDataModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv2.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv2.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ConfigObjectEnrollment model
				configObjectEnrollmentModel := new(blockchainv2.ConfigObjectEnrollment)
				configObjectEnrollmentModel.Component = configObjectEnrollmentComponentModel
				configObjectEnrollmentModel.Tls = configObjectEnrollmentTlsModel

				// Construct an instance of the ConfigObjectMsp model
				configObjectMspModel := new(blockchainv2.ConfigObjectMsp)
				configObjectMspModel.Component = mspConfigDataModel
				configObjectMspModel.Tls = mspConfigDataModel
				configObjectMspModel.Clientauth = mspConfigDataModel

				// Construct an instance of the ConfigOrdererDebug model
				configOrdererDebugModel := new(blockchainv2.ConfigOrdererDebug)
				configOrdererDebugModel.BroadcastTraceDir = core.StringPtr("testString")
				configOrdererDebugModel.DeliverTraceDir = core.StringPtr("testString")

				// Construct an instance of the ConfigOrdererGeneral model
				configOrdererGeneralModel := new(blockchainv2.ConfigOrdererGeneral)
				configOrdererGeneralModel.Keepalive = configOrdererKeepaliveModel
				configOrdererGeneralModel.BCCSP = bccspModel
				configOrdererGeneralModel.Authentication = configOrdererAuthenticationModel

				// Construct an instance of the ConfigOrdererMetrics model
				configOrdererMetricsModel := new(blockchainv2.ConfigOrdererMetrics)
				configOrdererMetricsModel.Provider = core.StringPtr("disabled")
				configOrdererMetricsModel.Statsd = configOrdererMetricsStatsdModel

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv2.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the StorageObject model
				storageObjectModel := new(blockchainv2.StorageObject)
				storageObjectModel.Size = core.StringPtr("4GiB")
				storageObjectModel.Class = core.StringPtr("default")

				// Construct an instance of the ConfigObject model
				configObjectModel := new(blockchainv2.ConfigObject)
				configObjectModel.Enrollment = configObjectEnrollmentModel
				configObjectModel.Msp = configObjectMspModel

				// Construct an instance of the ConfigOrdererCreate model
				configOrdererCreateModel := new(blockchainv2.ConfigOrdererCreate)
				configOrdererCreateModel.General = configOrdererGeneralModel
				configOrdererCreateModel.Debug = configOrdererDebugModel
				configOrdererCreateModel.Metrics = configOrdererMetricsModel

				// Construct an instance of the CreateOrdererRaftBodyResources model
				createOrdererRaftBodyResourcesModel := new(blockchainv2.CreateOrdererRaftBodyResources)
				createOrdererRaftBodyResourcesModel.Orderer = resourceObjectModel
				createOrdererRaftBodyResourcesModel.Proxy = resourceObjectModel

				// Construct an instance of the CreateOrdererRaftBodyStorage model
				createOrdererRaftBodyStorageModel := new(blockchainv2.CreateOrdererRaftBodyStorage)
				createOrdererRaftBodyStorageModel.Orderer = storageObjectModel

				// Construct an instance of the Hsm model
				hsmModel := new(blockchainv2.Hsm)
				hsmModel.Pkcs11endpoint = core.StringPtr("tcp://example.com:666")

				// Construct an instance of the CreateOrdererOptions model
				createOrdererOptionsModel := new(blockchainv2.CreateOrdererOptions)
				createOrdererOptionsModel.OrdererType = core.StringPtr("raft")
				createOrdererOptionsModel.MspID = core.StringPtr("Org1")
				createOrdererOptionsModel.DisplayName = core.StringPtr("orderer")
				createOrdererOptionsModel.Config = []blockchainv2.ConfigObject{*configObjectModel}
				createOrdererOptionsModel.ClusterName = core.StringPtr("ordering service 1")
				createOrdererOptionsModel.ClusterID = core.StringPtr("abcde")
				createOrdererOptionsModel.ExternalAppend = core.StringPtr("false")
				createOrdererOptionsModel.ConfigOverride = []blockchainv2.ConfigOrdererCreate{*configOrdererCreateModel}
				createOrdererOptionsModel.Resources = createOrdererRaftBodyResourcesModel
				createOrdererOptionsModel.Storage = createOrdererRaftBodyStorageModel
				createOrdererOptionsModel.SystemChannelID = core.StringPtr("testchainid")
				createOrdererOptionsModel.Zone = []string{"testString"}
				createOrdererOptionsModel.Tags = []string{"testString"}
				createOrdererOptionsModel.Region = []string{"testString"}
				createOrdererOptionsModel.Hsm = hsmModel
				createOrdererOptionsModel.Version = core.StringPtr("1.4.6-1")
				createOrdererOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateOrderer(createOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateOrderer(createOrdererOptions *CreateOrdererOptions)`, func() {
		createOrdererPath := "/ak/api/v2/kubernetes/components/fabric-orderer"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createOrdererPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "component-1", "dep_component_id": "admin", "type": "fabric-peer", "display_name": "orderer", "grpcwp_url": "https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443", "api_url": "grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050", "operations_url": "https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443", "msp_id": "Org1", "config_override": {"anyKey": "anyValue"}, "consenter_proposal_fin": true, "node_ou": {"enabled": true}, "ecert": {"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "cacert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}, "location": "ibmcloud", "timestamp": 1537262855753, "resources": {"orderer": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"orderer": {"size": "4GiB", "class": "default"}}, "system_channel_id": "testchainid", "tags": ["Tags"], "tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "server_tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "client_tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "orderer_type": "raft", "version": "1.4.6-1", "zone": "Zone"}`)
				}))
			})
			It(`Invoke CreateOrderer successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateOrderer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv2.BccspPKCS11)
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv2.BccspSW)
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the ConfigObjectEnrollmentComponentCatls model
				configObjectEnrollmentComponentCatlsModel := new(blockchainv2.ConfigObjectEnrollmentComponentCatls)
				configObjectEnrollmentComponentCatlsModel.Cacert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")

				// Construct an instance of the ConfigObjectEnrollmentTlsCatls model
				configObjectEnrollmentTlsCatlsModel := new(blockchainv2.ConfigObjectEnrollmentTlsCatls)
				configObjectEnrollmentTlsCatlsModel.Cacert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")

				// Construct an instance of the ConfigObjectEnrollmentTlsCsr model
				configObjectEnrollmentTlsCsrModel := new(blockchainv2.ConfigObjectEnrollmentTlsCsr)
				configObjectEnrollmentTlsCsrModel.Hosts = []string{"testString"}

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv2.Bccsp)
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model

				// Construct an instance of the ConfigObjectEnrollmentComponent model
				configObjectEnrollmentComponentModel := new(blockchainv2.ConfigObjectEnrollmentComponent)
				configObjectEnrollmentComponentModel.Cahost = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				configObjectEnrollmentComponentModel.Caport = core.Float64Ptr(float64(7054))
				configObjectEnrollmentComponentModel.Caname = core.StringPtr("ca")
				configObjectEnrollmentComponentModel.Catls = configObjectEnrollmentComponentCatlsModel
				configObjectEnrollmentComponentModel.Enrollid = core.StringPtr("admin")
				configObjectEnrollmentComponentModel.Enrollsecret = core.StringPtr("password")
				configObjectEnrollmentComponentModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the ConfigObjectEnrollmentTls model
				configObjectEnrollmentTlsModel := new(blockchainv2.ConfigObjectEnrollmentTls)
				configObjectEnrollmentTlsModel.Cahost = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				configObjectEnrollmentTlsModel.Caport = core.Float64Ptr(float64(7054))
				configObjectEnrollmentTlsModel.Caname = core.StringPtr("tlsca")
				configObjectEnrollmentTlsModel.Catls = configObjectEnrollmentTlsCatlsModel
				configObjectEnrollmentTlsModel.Enrollid = core.StringPtr("admin")
				configObjectEnrollmentTlsModel.Enrollsecret = core.StringPtr("password")
				configObjectEnrollmentTlsModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				configObjectEnrollmentTlsModel.Csr = configObjectEnrollmentTlsCsrModel

				// Construct an instance of the ConfigOrdererAuthentication model
				configOrdererAuthenticationModel := new(blockchainv2.ConfigOrdererAuthentication)
				configOrdererAuthenticationModel.TimeWindow = core.StringPtr("15m")
				configOrdererAuthenticationModel.NoExpirationChecks = core.BoolPtr(false)

				// Construct an instance of the ConfigOrdererKeepalive model
				configOrdererKeepaliveModel := new(blockchainv2.ConfigOrdererKeepalive)
				configOrdererKeepaliveModel.ServerMinInterval = core.StringPtr("60s")
				configOrdererKeepaliveModel.ServerInterval = core.StringPtr("2h")
				configOrdererKeepaliveModel.ServerTimeout = core.StringPtr("20s")

				// Construct an instance of the ConfigOrdererMetricsStatsd model
				configOrdererMetricsStatsdModel := new(blockchainv2.ConfigOrdererMetricsStatsd)
				configOrdererMetricsStatsdModel.Network = core.StringPtr("udp")
				configOrdererMetricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				configOrdererMetricsStatsdModel.WriteInterval = core.StringPtr("10s")
				configOrdererMetricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the MspConfigData model
				mspConfigDataModel := new(blockchainv2.MspConfigData)
				mspConfigDataModel.Keystore = core.StringPtr("testString")
				mspConfigDataModel.Signcerts = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspConfigDataModel.Cacerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				mspConfigDataModel.Intermediatecerts = []string{"testString"}
				mspConfigDataModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv2.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv2.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ConfigObjectEnrollment model
				configObjectEnrollmentModel := new(blockchainv2.ConfigObjectEnrollment)
				configObjectEnrollmentModel.Component = configObjectEnrollmentComponentModel
				configObjectEnrollmentModel.Tls = configObjectEnrollmentTlsModel

				// Construct an instance of the ConfigObjectMsp model
				configObjectMspModel := new(blockchainv2.ConfigObjectMsp)
				configObjectMspModel.Component = mspConfigDataModel
				configObjectMspModel.Tls = mspConfigDataModel
				configObjectMspModel.Clientauth = mspConfigDataModel

				// Construct an instance of the ConfigOrdererDebug model
				configOrdererDebugModel := new(blockchainv2.ConfigOrdererDebug)
				configOrdererDebugModel.BroadcastTraceDir = core.StringPtr("testString")
				configOrdererDebugModel.DeliverTraceDir = core.StringPtr("testString")

				// Construct an instance of the ConfigOrdererGeneral model
				configOrdererGeneralModel := new(blockchainv2.ConfigOrdererGeneral)
				configOrdererGeneralModel.Keepalive = configOrdererKeepaliveModel
				configOrdererGeneralModel.BCCSP = bccspModel
				configOrdererGeneralModel.Authentication = configOrdererAuthenticationModel

				// Construct an instance of the ConfigOrdererMetrics model
				configOrdererMetricsModel := new(blockchainv2.ConfigOrdererMetrics)
				configOrdererMetricsModel.Provider = core.StringPtr("disabled")
				configOrdererMetricsModel.Statsd = configOrdererMetricsStatsdModel

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv2.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the StorageObject model
				storageObjectModel := new(blockchainv2.StorageObject)
				storageObjectModel.Size = core.StringPtr("4GiB")
				storageObjectModel.Class = core.StringPtr("default")

				// Construct an instance of the ConfigObject model
				configObjectModel := new(blockchainv2.ConfigObject)
				configObjectModel.Enrollment = configObjectEnrollmentModel
				configObjectModel.Msp = configObjectMspModel

				// Construct an instance of the ConfigOrdererCreate model
				configOrdererCreateModel := new(blockchainv2.ConfigOrdererCreate)
				configOrdererCreateModel.General = configOrdererGeneralModel
				configOrdererCreateModel.Debug = configOrdererDebugModel
				configOrdererCreateModel.Metrics = configOrdererMetricsModel

				// Construct an instance of the CreateOrdererRaftBodyResources model
				createOrdererRaftBodyResourcesModel := new(blockchainv2.CreateOrdererRaftBodyResources)
				createOrdererRaftBodyResourcesModel.Orderer = resourceObjectModel
				createOrdererRaftBodyResourcesModel.Proxy = resourceObjectModel

				// Construct an instance of the CreateOrdererRaftBodyStorage model
				createOrdererRaftBodyStorageModel := new(blockchainv2.CreateOrdererRaftBodyStorage)
				createOrdererRaftBodyStorageModel.Orderer = storageObjectModel

				// Construct an instance of the Hsm model
				hsmModel := new(blockchainv2.Hsm)
				hsmModel.Pkcs11endpoint = core.StringPtr("tcp://example.com:666")

				// Construct an instance of the CreateOrdererOptions model
				createOrdererOptionsModel := new(blockchainv2.CreateOrdererOptions)
				createOrdererOptionsModel.OrdererType = core.StringPtr("raft")
				createOrdererOptionsModel.MspID = core.StringPtr("Org1")
				createOrdererOptionsModel.DisplayName = core.StringPtr("orderer")
				createOrdererOptionsModel.Config = []blockchainv2.ConfigObject{*configObjectModel}
				createOrdererOptionsModel.ClusterName = core.StringPtr("ordering service 1")
				createOrdererOptionsModel.ClusterID = core.StringPtr("abcde")
				createOrdererOptionsModel.ExternalAppend = core.StringPtr("false")
				createOrdererOptionsModel.ConfigOverride = []blockchainv2.ConfigOrdererCreate{*configOrdererCreateModel}
				createOrdererOptionsModel.Resources = createOrdererRaftBodyResourcesModel
				createOrdererOptionsModel.Storage = createOrdererRaftBodyStorageModel
				createOrdererOptionsModel.SystemChannelID = core.StringPtr("testchainid")
				createOrdererOptionsModel.Zone = []string{"testString"}
				createOrdererOptionsModel.Tags = []string{"testString"}
				createOrdererOptionsModel.Region = []string{"testString"}
				createOrdererOptionsModel.Hsm = hsmModel
				createOrdererOptionsModel.Version = core.StringPtr("1.4.6-1")
 				createOrdererOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateOrderer(createOrdererOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateOrderer with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv2.BccspPKCS11)
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv2.BccspSW)
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))

				// Construct an instance of the ConfigObjectEnrollmentComponentCatls model
				configObjectEnrollmentComponentCatlsModel := new(blockchainv2.ConfigObjectEnrollmentComponentCatls)
				configObjectEnrollmentComponentCatlsModel.Cacert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")

				// Construct an instance of the ConfigObjectEnrollmentTlsCatls model
				configObjectEnrollmentTlsCatlsModel := new(blockchainv2.ConfigObjectEnrollmentTlsCatls)
				configObjectEnrollmentTlsCatlsModel.Cacert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")

				// Construct an instance of the ConfigObjectEnrollmentTlsCsr model
				configObjectEnrollmentTlsCsrModel := new(blockchainv2.ConfigObjectEnrollmentTlsCsr)
				configObjectEnrollmentTlsCsrModel.Hosts = []string{"testString"}

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv2.Bccsp)
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model

				// Construct an instance of the ConfigObjectEnrollmentComponent model
				configObjectEnrollmentComponentModel := new(blockchainv2.ConfigObjectEnrollmentComponent)
				configObjectEnrollmentComponentModel.Cahost = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				configObjectEnrollmentComponentModel.Caport = core.Float64Ptr(float64(7054))
				configObjectEnrollmentComponentModel.Caname = core.StringPtr("ca")
				configObjectEnrollmentComponentModel.Catls = configObjectEnrollmentComponentCatlsModel
				configObjectEnrollmentComponentModel.Enrollid = core.StringPtr("admin")
				configObjectEnrollmentComponentModel.Enrollsecret = core.StringPtr("password")
				configObjectEnrollmentComponentModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the ConfigObjectEnrollmentTls model
				configObjectEnrollmentTlsModel := new(blockchainv2.ConfigObjectEnrollmentTls)
				configObjectEnrollmentTlsModel.Cahost = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				configObjectEnrollmentTlsModel.Caport = core.Float64Ptr(float64(7054))
				configObjectEnrollmentTlsModel.Caname = core.StringPtr("tlsca")
				configObjectEnrollmentTlsModel.Catls = configObjectEnrollmentTlsCatlsModel
				configObjectEnrollmentTlsModel.Enrollid = core.StringPtr("admin")
				configObjectEnrollmentTlsModel.Enrollsecret = core.StringPtr("password")
				configObjectEnrollmentTlsModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				configObjectEnrollmentTlsModel.Csr = configObjectEnrollmentTlsCsrModel

				// Construct an instance of the ConfigOrdererAuthentication model
				configOrdererAuthenticationModel := new(blockchainv2.ConfigOrdererAuthentication)
				configOrdererAuthenticationModel.TimeWindow = core.StringPtr("15m")
				configOrdererAuthenticationModel.NoExpirationChecks = core.BoolPtr(false)

				// Construct an instance of the ConfigOrdererKeepalive model
				configOrdererKeepaliveModel := new(blockchainv2.ConfigOrdererKeepalive)
				configOrdererKeepaliveModel.ServerMinInterval = core.StringPtr("60s")
				configOrdererKeepaliveModel.ServerInterval = core.StringPtr("2h")
				configOrdererKeepaliveModel.ServerTimeout = core.StringPtr("20s")

				// Construct an instance of the ConfigOrdererMetricsStatsd model
				configOrdererMetricsStatsdModel := new(blockchainv2.ConfigOrdererMetricsStatsd)
				configOrdererMetricsStatsdModel.Network = core.StringPtr("udp")
				configOrdererMetricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				configOrdererMetricsStatsdModel.WriteInterval = core.StringPtr("10s")
				configOrdererMetricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the MspConfigData model
				mspConfigDataModel := new(blockchainv2.MspConfigData)
				mspConfigDataModel.Keystore = core.StringPtr("testString")
				mspConfigDataModel.Signcerts = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspConfigDataModel.Cacerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				mspConfigDataModel.Intermediatecerts = []string{"testString"}
				mspConfigDataModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv2.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv2.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ConfigObjectEnrollment model
				configObjectEnrollmentModel := new(blockchainv2.ConfigObjectEnrollment)
				configObjectEnrollmentModel.Component = configObjectEnrollmentComponentModel
				configObjectEnrollmentModel.Tls = configObjectEnrollmentTlsModel

				// Construct an instance of the ConfigObjectMsp model
				configObjectMspModel := new(blockchainv2.ConfigObjectMsp)
				configObjectMspModel.Component = mspConfigDataModel
				configObjectMspModel.Tls = mspConfigDataModel
				configObjectMspModel.Clientauth = mspConfigDataModel

				// Construct an instance of the ConfigOrdererDebug model
				configOrdererDebugModel := new(blockchainv2.ConfigOrdererDebug)
				configOrdererDebugModel.BroadcastTraceDir = core.StringPtr("testString")
				configOrdererDebugModel.DeliverTraceDir = core.StringPtr("testString")

				// Construct an instance of the ConfigOrdererGeneral model
				configOrdererGeneralModel := new(blockchainv2.ConfigOrdererGeneral)
				configOrdererGeneralModel.Keepalive = configOrdererKeepaliveModel
				configOrdererGeneralModel.BCCSP = bccspModel
				configOrdererGeneralModel.Authentication = configOrdererAuthenticationModel

				// Construct an instance of the ConfigOrdererMetrics model
				configOrdererMetricsModel := new(blockchainv2.ConfigOrdererMetrics)
				configOrdererMetricsModel.Provider = core.StringPtr("disabled")
				configOrdererMetricsModel.Statsd = configOrdererMetricsStatsdModel

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv2.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the StorageObject model
				storageObjectModel := new(blockchainv2.StorageObject)
				storageObjectModel.Size = core.StringPtr("4GiB")
				storageObjectModel.Class = core.StringPtr("default")

				// Construct an instance of the ConfigObject model
				configObjectModel := new(blockchainv2.ConfigObject)
				configObjectModel.Enrollment = configObjectEnrollmentModel
				configObjectModel.Msp = configObjectMspModel

				// Construct an instance of the ConfigOrdererCreate model
				configOrdererCreateModel := new(blockchainv2.ConfigOrdererCreate)
				configOrdererCreateModel.General = configOrdererGeneralModel
				configOrdererCreateModel.Debug = configOrdererDebugModel
				configOrdererCreateModel.Metrics = configOrdererMetricsModel

				// Construct an instance of the CreateOrdererRaftBodyResources model
				createOrdererRaftBodyResourcesModel := new(blockchainv2.CreateOrdererRaftBodyResources)
				createOrdererRaftBodyResourcesModel.Orderer = resourceObjectModel
				createOrdererRaftBodyResourcesModel.Proxy = resourceObjectModel

				// Construct an instance of the CreateOrdererRaftBodyStorage model
				createOrdererRaftBodyStorageModel := new(blockchainv2.CreateOrdererRaftBodyStorage)
				createOrdererRaftBodyStorageModel.Orderer = storageObjectModel

				// Construct an instance of the Hsm model
				hsmModel := new(blockchainv2.Hsm)
				hsmModel.Pkcs11endpoint = core.StringPtr("tcp://example.com:666")

				// Construct an instance of the CreateOrdererOptions model
				createOrdererOptionsModel := new(blockchainv2.CreateOrdererOptions)
				createOrdererOptionsModel.OrdererType = core.StringPtr("raft")
				createOrdererOptionsModel.MspID = core.StringPtr("Org1")
				createOrdererOptionsModel.DisplayName = core.StringPtr("orderer")
				createOrdererOptionsModel.Config = []blockchainv2.ConfigObject{*configObjectModel}
				createOrdererOptionsModel.ClusterName = core.StringPtr("ordering service 1")
				createOrdererOptionsModel.ClusterID = core.StringPtr("abcde")
				createOrdererOptionsModel.ExternalAppend = core.StringPtr("false")
				createOrdererOptionsModel.ConfigOverride = []blockchainv2.ConfigOrdererCreate{*configOrdererCreateModel}
				createOrdererOptionsModel.Resources = createOrdererRaftBodyResourcesModel
				createOrdererOptionsModel.Storage = createOrdererRaftBodyStorageModel
				createOrdererOptionsModel.SystemChannelID = core.StringPtr("testchainid")
				createOrdererOptionsModel.Zone = []string{"testString"}
				createOrdererOptionsModel.Tags = []string{"testString"}
				createOrdererOptionsModel.Region = []string{"testString"}
				createOrdererOptionsModel.Hsm = hsmModel
				createOrdererOptionsModel.Version = core.StringPtr("1.4.6-1")
				createOrdererOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateOrderer(createOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateOrdererOptions model with no property values
				createOrdererOptionsModelNew := new(blockchainv2.CreateOrdererOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreateOrderer(createOrdererOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ImportOrderer(importOrdererOptions *ImportOrdererOptions) - Operation response error`, func() {
		importOrdererPath := "/ak/api/v2/components/fabric-orderer"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(importOrdererPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ImportOrderer with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ImportOrdererOptions model
				importOrdererOptionsModel := new(blockchainv2.ImportOrdererOptions)
				importOrdererOptionsModel.ClusterName = core.StringPtr("ordering service 1")
				importOrdererOptionsModel.DisplayName = core.StringPtr("orderer")
				importOrdererOptionsModel.MspID = core.StringPtr("Org1")
				importOrdererOptionsModel.GrpcwpURL = core.StringPtr("https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443")
				importOrdererOptionsModel.TlsCaRootCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkRpZmZlcmVudCBkYXRhIGhlcmUgaWYgdGhpcyB3YXMgcmVhbAotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==")
				importOrdererOptionsModel.Location = core.StringPtr("ibmcloud")
				importOrdererOptionsModel.ApiURL = core.StringPtr("grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				importOrdererOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443")
				importOrdererOptionsModel.SystemChannelID = core.StringPtr("testchainid")
				importOrdererOptionsModel.Tags = []string{"testString"}
				importOrdererOptionsModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				importOrdererOptionsModel.ServerTlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				importOrdererOptionsModel.ClientTlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				importOrdererOptionsModel.ClusterID = core.StringPtr("testString")
				importOrdererOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ImportOrderer(importOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ImportOrderer(importOrdererOptions *ImportOrdererOptions)`, func() {
		importOrdererPath := "/ak/api/v2/components/fabric-orderer"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(importOrdererPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "component-1", "dep_component_id": "admin", "type": "fabric-peer", "display_name": "orderer", "grpcwp_url": "https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443", "api_url": "grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050", "operations_url": "https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443", "msp_id": "Org1", "config_override": {"anyKey": "anyValue"}, "consenter_proposal_fin": true, "node_ou": {"enabled": true}, "ecert": {"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "cacert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}, "location": "ibmcloud", "timestamp": 1537262855753, "resources": {"orderer": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"orderer": {"size": "4GiB", "class": "default"}}, "system_channel_id": "testchainid", "tags": ["Tags"], "tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "server_tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "client_tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "orderer_type": "raft", "version": "1.4.6-1", "zone": "Zone"}`)
				}))
			})
			It(`Invoke ImportOrderer successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ImportOrderer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ImportOrdererOptions model
				importOrdererOptionsModel := new(blockchainv2.ImportOrdererOptions)
				importOrdererOptionsModel.ClusterName = core.StringPtr("ordering service 1")
				importOrdererOptionsModel.DisplayName = core.StringPtr("orderer")
				importOrdererOptionsModel.MspID = core.StringPtr("Org1")
				importOrdererOptionsModel.GrpcwpURL = core.StringPtr("https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443")
				importOrdererOptionsModel.TlsCaRootCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkRpZmZlcmVudCBkYXRhIGhlcmUgaWYgdGhpcyB3YXMgcmVhbAotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==")
				importOrdererOptionsModel.Location = core.StringPtr("ibmcloud")
				importOrdererOptionsModel.ApiURL = core.StringPtr("grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				importOrdererOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443")
				importOrdererOptionsModel.SystemChannelID = core.StringPtr("testchainid")
				importOrdererOptionsModel.Tags = []string{"testString"}
				importOrdererOptionsModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				importOrdererOptionsModel.ServerTlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				importOrdererOptionsModel.ClientTlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				importOrdererOptionsModel.ClusterID = core.StringPtr("testString")
 				importOrdererOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ImportOrderer(importOrdererOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ImportOrderer with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ImportOrdererOptions model
				importOrdererOptionsModel := new(blockchainv2.ImportOrdererOptions)
				importOrdererOptionsModel.ClusterName = core.StringPtr("ordering service 1")
				importOrdererOptionsModel.DisplayName = core.StringPtr("orderer")
				importOrdererOptionsModel.MspID = core.StringPtr("Org1")
				importOrdererOptionsModel.GrpcwpURL = core.StringPtr("https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443")
				importOrdererOptionsModel.TlsCaRootCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkRpZmZlcmVudCBkYXRhIGhlcmUgaWYgdGhpcyB3YXMgcmVhbAotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==")
				importOrdererOptionsModel.Location = core.StringPtr("ibmcloud")
				importOrdererOptionsModel.ApiURL = core.StringPtr("grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				importOrdererOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443")
				importOrdererOptionsModel.SystemChannelID = core.StringPtr("testchainid")
				importOrdererOptionsModel.Tags = []string{"testString"}
				importOrdererOptionsModel.TlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				importOrdererOptionsModel.ServerTlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				importOrdererOptionsModel.ClientTlsCert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				importOrdererOptionsModel.ClusterID = core.StringPtr("testString")
				importOrdererOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ImportOrderer(importOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ImportOrdererOptions model with no property values
				importOrdererOptionsModelNew := new(blockchainv2.ImportOrdererOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ImportOrderer(importOrdererOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`EditOrderer(editOrdererOptions *EditOrdererOptions) - Operation response error`, func() {
		editOrdererPath := "/ak/api/v2/components/fabric-orderer/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(editOrdererPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke EditOrderer with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the EditOrdererOptions model
				editOrdererOptionsModel := new(blockchainv2.EditOrdererOptions)
				editOrdererOptionsModel.ID = core.StringPtr("testString")
				editOrdererOptionsModel.ClusterName = core.StringPtr("ordering service 1")
				editOrdererOptionsModel.DisplayName = core.StringPtr("orderer")
				editOrdererOptionsModel.ApiURL = core.StringPtr("grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				editOrdererOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443")
				editOrdererOptionsModel.GrpcwpURL = core.StringPtr("https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443")
				editOrdererOptionsModel.MspID = core.StringPtr("Org1")
				editOrdererOptionsModel.ConsenterProposalFin = core.BoolPtr(true)
				editOrdererOptionsModel.Location = core.StringPtr("ibmcloud")
				editOrdererOptionsModel.SystemChannelID = core.StringPtr("testchainid")
				editOrdererOptionsModel.Tags = []string{"testString"}
				editOrdererOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.EditOrderer(editOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`EditOrderer(editOrdererOptions *EditOrdererOptions)`, func() {
		editOrdererPath := "/ak/api/v2/components/fabric-orderer/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(editOrdererPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "component-1", "dep_component_id": "admin", "type": "fabric-peer", "display_name": "orderer", "grpcwp_url": "https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443", "api_url": "grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050", "operations_url": "https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443", "msp_id": "Org1", "config_override": {"anyKey": "anyValue"}, "consenter_proposal_fin": true, "node_ou": {"enabled": true}, "ecert": {"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "cacert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}, "location": "ibmcloud", "timestamp": 1537262855753, "resources": {"orderer": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"orderer": {"size": "4GiB", "class": "default"}}, "system_channel_id": "testchainid", "tags": ["Tags"], "tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "server_tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "client_tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "orderer_type": "raft", "version": "1.4.6-1", "zone": "Zone"}`)
				}))
			})
			It(`Invoke EditOrderer successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.EditOrderer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EditOrdererOptions model
				editOrdererOptionsModel := new(blockchainv2.EditOrdererOptions)
				editOrdererOptionsModel.ID = core.StringPtr("testString")
				editOrdererOptionsModel.ClusterName = core.StringPtr("ordering service 1")
				editOrdererOptionsModel.DisplayName = core.StringPtr("orderer")
				editOrdererOptionsModel.ApiURL = core.StringPtr("grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				editOrdererOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443")
				editOrdererOptionsModel.GrpcwpURL = core.StringPtr("https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443")
				editOrdererOptionsModel.MspID = core.StringPtr("Org1")
				editOrdererOptionsModel.ConsenterProposalFin = core.BoolPtr(true)
				editOrdererOptionsModel.Location = core.StringPtr("ibmcloud")
				editOrdererOptionsModel.SystemChannelID = core.StringPtr("testchainid")
				editOrdererOptionsModel.Tags = []string{"testString"}
 				editOrdererOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.EditOrderer(editOrdererOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke EditOrderer with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the EditOrdererOptions model
				editOrdererOptionsModel := new(blockchainv2.EditOrdererOptions)
				editOrdererOptionsModel.ID = core.StringPtr("testString")
				editOrdererOptionsModel.ClusterName = core.StringPtr("ordering service 1")
				editOrdererOptionsModel.DisplayName = core.StringPtr("orderer")
				editOrdererOptionsModel.ApiURL = core.StringPtr("grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				editOrdererOptionsModel.OperationsURL = core.StringPtr("https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443")
				editOrdererOptionsModel.GrpcwpURL = core.StringPtr("https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443")
				editOrdererOptionsModel.MspID = core.StringPtr("Org1")
				editOrdererOptionsModel.ConsenterProposalFin = core.BoolPtr(true)
				editOrdererOptionsModel.Location = core.StringPtr("ibmcloud")
				editOrdererOptionsModel.SystemChannelID = core.StringPtr("testchainid")
				editOrdererOptionsModel.Tags = []string{"testString"}
				editOrdererOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.EditOrderer(editOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the EditOrdererOptions model with no property values
				editOrdererOptionsModelNew := new(blockchainv2.EditOrdererOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.EditOrderer(editOrdererOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateOrderer(updateOrdererOptions *UpdateOrdererOptions) - Operation response error`, func() {
		updateOrdererPath := "/ak/api/v2/kubernetes/components/fabric-orderer/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateOrdererPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateOrderer with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ConfigOrdererAuthentication model
				configOrdererAuthenticationModel := new(blockchainv2.ConfigOrdererAuthentication)
				configOrdererAuthenticationModel.TimeWindow = core.StringPtr("15m")
				configOrdererAuthenticationModel.NoExpirationChecks = core.BoolPtr(false)

				// Construct an instance of the ConfigOrdererKeepalive model
				configOrdererKeepaliveModel := new(blockchainv2.ConfigOrdererKeepalive)
				configOrdererKeepaliveModel.ServerMinInterval = core.StringPtr("60s")
				configOrdererKeepaliveModel.ServerInterval = core.StringPtr("2h")
				configOrdererKeepaliveModel.ServerTimeout = core.StringPtr("20s")

				// Construct an instance of the ConfigOrdererMetricsStatsd model
				configOrdererMetricsStatsdModel := new(blockchainv2.ConfigOrdererMetricsStatsd)
				configOrdererMetricsStatsdModel.Network = core.StringPtr("udp")
				configOrdererMetricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				configOrdererMetricsStatsdModel.WriteInterval = core.StringPtr("10s")
				configOrdererMetricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv2.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv2.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ConfigOrdererDebug model
				configOrdererDebugModel := new(blockchainv2.ConfigOrdererDebug)
				configOrdererDebugModel.BroadcastTraceDir = core.StringPtr("testString")
				configOrdererDebugModel.DeliverTraceDir = core.StringPtr("testString")

				// Construct an instance of the ConfigOrdererGeneralUpdate model
				configOrdererGeneralUpdateModel := new(blockchainv2.ConfigOrdererGeneralUpdate)
				configOrdererGeneralUpdateModel.Keepalive = configOrdererKeepaliveModel
				configOrdererGeneralUpdateModel.Authentication = configOrdererAuthenticationModel

				// Construct an instance of the ConfigOrdererMetrics model
				configOrdererMetricsModel := new(blockchainv2.ConfigOrdererMetrics)
				configOrdererMetricsModel.Provider = core.StringPtr("disabled")
				configOrdererMetricsModel.Statsd = configOrdererMetricsStatsdModel

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv2.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the ConfigOrdererUpdate model
				configOrdererUpdateModel := new(blockchainv2.ConfigOrdererUpdate)
				configOrdererUpdateModel.General = configOrdererGeneralUpdateModel
				configOrdererUpdateModel.Debug = configOrdererDebugModel
				configOrdererUpdateModel.Metrics = configOrdererMetricsModel

				// Construct an instance of the UpdateOrdererBodyResources model
				updateOrdererBodyResourcesModel := new(blockchainv2.UpdateOrdererBodyResources)
				updateOrdererBodyResourcesModel.Orderer = resourceObjectModel
				updateOrdererBodyResourcesModel.Proxy = resourceObjectModel

				// Construct an instance of the UpdateOrdererOptions model
				updateOrdererOptionsModel := new(blockchainv2.UpdateOrdererOptions)
				updateOrdererOptionsModel.ID = core.StringPtr("testString")
				updateOrdererOptionsModel.ConfigOverride = configOrdererUpdateModel
				updateOrdererOptionsModel.Resources = updateOrdererBodyResourcesModel
				updateOrdererOptionsModel.Zone = core.StringPtr("testString")
				updateOrdererOptionsModel.Version = core.StringPtr("1.4.6-1")
				updateOrdererOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateOrderer(updateOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateOrderer(updateOrdererOptions *UpdateOrdererOptions)`, func() {
		updateOrdererPath := "/ak/api/v2/kubernetes/components/fabric-orderer/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateOrdererPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "component-1", "dep_component_id": "admin", "type": "fabric-peer", "display_name": "orderer", "grpcwp_url": "https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443", "api_url": "grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050", "operations_url": "https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443", "msp_id": "Org1", "config_override": {"anyKey": "anyValue"}, "consenter_proposal_fin": true, "node_ou": {"enabled": true}, "ecert": {"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "cacert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}, "location": "ibmcloud", "timestamp": 1537262855753, "resources": {"orderer": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"orderer": {"size": "4GiB", "class": "default"}}, "system_channel_id": "testchainid", "tags": ["Tags"], "tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "server_tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "client_tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "orderer_type": "raft", "version": "1.4.6-1", "zone": "Zone"}`)
				}))
			})
			It(`Invoke UpdateOrderer successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateOrderer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ConfigOrdererAuthentication model
				configOrdererAuthenticationModel := new(blockchainv2.ConfigOrdererAuthentication)
				configOrdererAuthenticationModel.TimeWindow = core.StringPtr("15m")
				configOrdererAuthenticationModel.NoExpirationChecks = core.BoolPtr(false)

				// Construct an instance of the ConfigOrdererKeepalive model
				configOrdererKeepaliveModel := new(blockchainv2.ConfigOrdererKeepalive)
				configOrdererKeepaliveModel.ServerMinInterval = core.StringPtr("60s")
				configOrdererKeepaliveModel.ServerInterval = core.StringPtr("2h")
				configOrdererKeepaliveModel.ServerTimeout = core.StringPtr("20s")

				// Construct an instance of the ConfigOrdererMetricsStatsd model
				configOrdererMetricsStatsdModel := new(blockchainv2.ConfigOrdererMetricsStatsd)
				configOrdererMetricsStatsdModel.Network = core.StringPtr("udp")
				configOrdererMetricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				configOrdererMetricsStatsdModel.WriteInterval = core.StringPtr("10s")
				configOrdererMetricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv2.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv2.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ConfigOrdererDebug model
				configOrdererDebugModel := new(blockchainv2.ConfigOrdererDebug)
				configOrdererDebugModel.BroadcastTraceDir = core.StringPtr("testString")
				configOrdererDebugModel.DeliverTraceDir = core.StringPtr("testString")

				// Construct an instance of the ConfigOrdererGeneralUpdate model
				configOrdererGeneralUpdateModel := new(blockchainv2.ConfigOrdererGeneralUpdate)
				configOrdererGeneralUpdateModel.Keepalive = configOrdererKeepaliveModel
				configOrdererGeneralUpdateModel.Authentication = configOrdererAuthenticationModel

				// Construct an instance of the ConfigOrdererMetrics model
				configOrdererMetricsModel := new(blockchainv2.ConfigOrdererMetrics)
				configOrdererMetricsModel.Provider = core.StringPtr("disabled")
				configOrdererMetricsModel.Statsd = configOrdererMetricsStatsdModel

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv2.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the ConfigOrdererUpdate model
				configOrdererUpdateModel := new(blockchainv2.ConfigOrdererUpdate)
				configOrdererUpdateModel.General = configOrdererGeneralUpdateModel
				configOrdererUpdateModel.Debug = configOrdererDebugModel
				configOrdererUpdateModel.Metrics = configOrdererMetricsModel

				// Construct an instance of the UpdateOrdererBodyResources model
				updateOrdererBodyResourcesModel := new(blockchainv2.UpdateOrdererBodyResources)
				updateOrdererBodyResourcesModel.Orderer = resourceObjectModel
				updateOrdererBodyResourcesModel.Proxy = resourceObjectModel

				// Construct an instance of the UpdateOrdererOptions model
				updateOrdererOptionsModel := new(blockchainv2.UpdateOrdererOptions)
				updateOrdererOptionsModel.ID = core.StringPtr("testString")
				updateOrdererOptionsModel.ConfigOverride = configOrdererUpdateModel
				updateOrdererOptionsModel.Resources = updateOrdererBodyResourcesModel
				updateOrdererOptionsModel.Zone = core.StringPtr("testString")
				updateOrdererOptionsModel.Version = core.StringPtr("1.4.6-1")
 				updateOrdererOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateOrderer(updateOrdererOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateOrderer with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ConfigOrdererAuthentication model
				configOrdererAuthenticationModel := new(blockchainv2.ConfigOrdererAuthentication)
				configOrdererAuthenticationModel.TimeWindow = core.StringPtr("15m")
				configOrdererAuthenticationModel.NoExpirationChecks = core.BoolPtr(false)

				// Construct an instance of the ConfigOrdererKeepalive model
				configOrdererKeepaliveModel := new(blockchainv2.ConfigOrdererKeepalive)
				configOrdererKeepaliveModel.ServerMinInterval = core.StringPtr("60s")
				configOrdererKeepaliveModel.ServerInterval = core.StringPtr("2h")
				configOrdererKeepaliveModel.ServerTimeout = core.StringPtr("20s")

				// Construct an instance of the ConfigOrdererMetricsStatsd model
				configOrdererMetricsStatsdModel := new(blockchainv2.ConfigOrdererMetricsStatsd)
				configOrdererMetricsStatsdModel.Network = core.StringPtr("udp")
				configOrdererMetricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				configOrdererMetricsStatsdModel.WriteInterval = core.StringPtr("10s")
				configOrdererMetricsStatsdModel.Prefix = core.StringPtr("server")

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv2.ResourceLimits)
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv2.ResourceRequests)
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")

				// Construct an instance of the ConfigOrdererDebug model
				configOrdererDebugModel := new(blockchainv2.ConfigOrdererDebug)
				configOrdererDebugModel.BroadcastTraceDir = core.StringPtr("testString")
				configOrdererDebugModel.DeliverTraceDir = core.StringPtr("testString")

				// Construct an instance of the ConfigOrdererGeneralUpdate model
				configOrdererGeneralUpdateModel := new(blockchainv2.ConfigOrdererGeneralUpdate)
				configOrdererGeneralUpdateModel.Keepalive = configOrdererKeepaliveModel
				configOrdererGeneralUpdateModel.Authentication = configOrdererAuthenticationModel

				// Construct an instance of the ConfigOrdererMetrics model
				configOrdererMetricsModel := new(blockchainv2.ConfigOrdererMetrics)
				configOrdererMetricsModel.Provider = core.StringPtr("disabled")
				configOrdererMetricsModel.Statsd = configOrdererMetricsStatsdModel

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv2.ResourceObject)
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel

				// Construct an instance of the ConfigOrdererUpdate model
				configOrdererUpdateModel := new(blockchainv2.ConfigOrdererUpdate)
				configOrdererUpdateModel.General = configOrdererGeneralUpdateModel
				configOrdererUpdateModel.Debug = configOrdererDebugModel
				configOrdererUpdateModel.Metrics = configOrdererMetricsModel

				// Construct an instance of the UpdateOrdererBodyResources model
				updateOrdererBodyResourcesModel := new(blockchainv2.UpdateOrdererBodyResources)
				updateOrdererBodyResourcesModel.Orderer = resourceObjectModel
				updateOrdererBodyResourcesModel.Proxy = resourceObjectModel

				// Construct an instance of the UpdateOrdererOptions model
				updateOrdererOptionsModel := new(blockchainv2.UpdateOrdererOptions)
				updateOrdererOptionsModel.ID = core.StringPtr("testString")
				updateOrdererOptionsModel.ConfigOverride = configOrdererUpdateModel
				updateOrdererOptionsModel.Resources = updateOrdererBodyResourcesModel
				updateOrdererOptionsModel.Zone = core.StringPtr("testString")
				updateOrdererOptionsModel.Version = core.StringPtr("1.4.6-1")
				updateOrdererOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateOrderer(updateOrdererOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateOrdererOptions model with no property values
				updateOrdererOptionsModelNew := new(blockchainv2.UpdateOrdererOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateOrderer(updateOrdererOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SubmitBlock(submitBlockOptions *SubmitBlockOptions) - Operation response error`, func() {
		submitBlockPath := "/ak/api/v2/kubernetes/components/testString/config"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(submitBlockPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SubmitBlock with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the SubmitBlockOptions model
				submitBlockOptionsModel := new(blockchainv2.SubmitBlockOptions)
				submitBlockOptionsModel.ID = core.StringPtr("testString")
				submitBlockOptionsModel.B64Block = core.StringPtr("bWFzc2l2ZSBiaW5hcnkgb2YgYSBjb25maWcgYmxvY2sgd291bGQgYmUgaGVyZSBpZiB0aGlzIHdhcyByZWFs")
				submitBlockOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.SubmitBlock(submitBlockOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SubmitBlock(submitBlockOptions *SubmitBlockOptions)`, func() {
		submitBlockPath := "/ak/api/v2/kubernetes/components/testString/config"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(submitBlockPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "text/plain")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "myca-2", "type": "fabric-ca", "display_name": "Example CA", "grpcwp_url": "https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "operations_url": "https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443", "msp_id": "Org1", "location": "ibmcloud", "ca_name": "ca", "admin_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="], "node_ou": {"enabled": true}, "ecert": {"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "cacert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}, "state_db": "couchdb", "timestamp": 1537262855753, "resources": {"ca": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "peer": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "orderer": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "statedb": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"ca": {"size": "4GiB", "class": "default"}, "peer": {"size": "4GiB", "class": "default"}, "orderer": {"size": "4GiB", "class": "default"}, "statedb": {"size": "4GiB", "class": "default"}}, "tags": ["Tags"], "tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "version": "Version", "zone": "Zone"}`)
				}))
			})
			It(`Invoke SubmitBlock successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.SubmitBlock(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SubmitBlockOptions model
				submitBlockOptionsModel := new(blockchainv2.SubmitBlockOptions)
				submitBlockOptionsModel.ID = core.StringPtr("testString")
				submitBlockOptionsModel.B64Block = core.StringPtr("bWFzc2l2ZSBiaW5hcnkgb2YgYSBjb25maWcgYmxvY2sgd291bGQgYmUgaGVyZSBpZiB0aGlzIHdhcyByZWFs")
 				submitBlockOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.SubmitBlock(submitBlockOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke SubmitBlock with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the SubmitBlockOptions model
				submitBlockOptionsModel := new(blockchainv2.SubmitBlockOptions)
				submitBlockOptionsModel.ID = core.StringPtr("testString")
				submitBlockOptionsModel.B64Block = core.StringPtr("bWFzc2l2ZSBiaW5hcnkgb2YgYSBjb25maWcgYmxvY2sgd291bGQgYmUgaGVyZSBpZiB0aGlzIHdhcyByZWFs")
				submitBlockOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.SubmitBlock(submitBlockOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SubmitBlockOptions model with no property values
				submitBlockOptionsModelNew := new(blockchainv2.SubmitBlockOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.SubmitBlock(submitBlockOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ImportMsp(importMspOptions *ImportMspOptions) - Operation response error`, func() {
		importMspPath := "/ak/api/v2/components/msp"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(importMspPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ImportMsp with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ImportMspOptions model
				importMspOptionsModel := new(blockchainv2.ImportMspOptions)
				importMspOptionsModel.MspID = core.StringPtr("Org1")
				importMspOptionsModel.DisplayName = core.StringPtr("My Peer")
				importMspOptionsModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				importMspOptionsModel.IntermediateCerts = []string{"testString"}
				importMspOptionsModel.Admins = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				importMspOptionsModel.TlsRootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				importMspOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ImportMsp(importMspOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ImportMsp(importMspOptions *ImportMspOptions)`, func() {
		importMspPath := "/ak/api/v2/components/msp"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(importMspPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "component-1", "type": "fabric-peer", "display_name": "My Peer", "msp_id": "Org1", "timestamp": 1537262855753, "tags": ["Tags"], "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="], "intermediate_certs": ["IntermediateCerts"], "admins": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="], "scheme_version": "v1", "tls_root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}`)
				}))
			})
			It(`Invoke ImportMsp successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ImportMsp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ImportMspOptions model
				importMspOptionsModel := new(blockchainv2.ImportMspOptions)
				importMspOptionsModel.MspID = core.StringPtr("Org1")
				importMspOptionsModel.DisplayName = core.StringPtr("My Peer")
				importMspOptionsModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				importMspOptionsModel.IntermediateCerts = []string{"testString"}
				importMspOptionsModel.Admins = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				importMspOptionsModel.TlsRootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
 				importMspOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ImportMsp(importMspOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ImportMsp with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ImportMspOptions model
				importMspOptionsModel := new(blockchainv2.ImportMspOptions)
				importMspOptionsModel.MspID = core.StringPtr("Org1")
				importMspOptionsModel.DisplayName = core.StringPtr("My Peer")
				importMspOptionsModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				importMspOptionsModel.IntermediateCerts = []string{"testString"}
				importMspOptionsModel.Admins = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				importMspOptionsModel.TlsRootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				importMspOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ImportMsp(importMspOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ImportMspOptions model with no property values
				importMspOptionsModelNew := new(blockchainv2.ImportMspOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ImportMsp(importMspOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`EditMsp(editMspOptions *EditMspOptions) - Operation response error`, func() {
		editMspPath := "/ak/api/v2/components/msp/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(editMspPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke EditMsp with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the EditMspOptions model
				editMspOptionsModel := new(blockchainv2.EditMspOptions)
				editMspOptionsModel.ID = core.StringPtr("testString")
				editMspOptionsModel.MspID = core.StringPtr("Org1")
				editMspOptionsModel.DisplayName = core.StringPtr("My Peer")
				editMspOptionsModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editMspOptionsModel.IntermediateCerts = []string{"testString"}
				editMspOptionsModel.Admins = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editMspOptionsModel.TlsRootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editMspOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.EditMsp(editMspOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`EditMsp(editMspOptions *EditMspOptions)`, func() {
		editMspPath := "/ak/api/v2/components/msp/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(editMspPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "component-1", "type": "fabric-peer", "display_name": "My Peer", "msp_id": "Org1", "timestamp": 1537262855753, "tags": ["Tags"], "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="], "intermediate_certs": ["IntermediateCerts"], "admins": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="], "scheme_version": "v1", "tls_root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}`)
				}))
			})
			It(`Invoke EditMsp successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.EditMsp(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EditMspOptions model
				editMspOptionsModel := new(blockchainv2.EditMspOptions)
				editMspOptionsModel.ID = core.StringPtr("testString")
				editMspOptionsModel.MspID = core.StringPtr("Org1")
				editMspOptionsModel.DisplayName = core.StringPtr("My Peer")
				editMspOptionsModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editMspOptionsModel.IntermediateCerts = []string{"testString"}
				editMspOptionsModel.Admins = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editMspOptionsModel.TlsRootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
 				editMspOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.EditMsp(editMspOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke EditMsp with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the EditMspOptions model
				editMspOptionsModel := new(blockchainv2.EditMspOptions)
				editMspOptionsModel.ID = core.StringPtr("testString")
				editMspOptionsModel.MspID = core.StringPtr("Org1")
				editMspOptionsModel.DisplayName = core.StringPtr("My Peer")
				editMspOptionsModel.RootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editMspOptionsModel.IntermediateCerts = []string{"testString"}
				editMspOptionsModel.Admins = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editMspOptionsModel.TlsRootCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editMspOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.EditMsp(editMspOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the EditMspOptions model with no property values
				editMspOptionsModelNew := new(blockchainv2.EditMspOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.EditMsp(editMspOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMspCertificate(getMspCertificateOptions *GetMspCertificateOptions) - Operation response error`, func() {
		getMspCertificatePath := "/ak/api/v2/components/msps/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getMspCertificatePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["cache"]).To(Equal([]string{"skip"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetMspCertificate with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetMspCertificateOptions model
				getMspCertificateOptionsModel := new(blockchainv2.GetMspCertificateOptions)
				getMspCertificateOptionsModel.MspID = core.StringPtr("testString")
				getMspCertificateOptionsModel.Cache = core.StringPtr("skip")
				getMspCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetMspCertificate(getMspCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetMspCertificate(getMspCertificateOptions *GetMspCertificateOptions)`, func() {
		getMspCertificatePath := "/ak/api/v2/components/msps/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getMspCertificatePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["cache"]).To(Equal([]string{"skip"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"msps": [{"msp_id": "Org1", "root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="], "admins": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="], "tls_root_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="]}]}`)
				}))
			})
			It(`Invoke GetMspCertificate successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetMspCertificate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetMspCertificateOptions model
				getMspCertificateOptionsModel := new(blockchainv2.GetMspCertificateOptions)
				getMspCertificateOptionsModel.MspID = core.StringPtr("testString")
				getMspCertificateOptionsModel.Cache = core.StringPtr("skip")
 				getMspCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetMspCertificate(getMspCertificateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetMspCertificate with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetMspCertificateOptions model
				getMspCertificateOptionsModel := new(blockchainv2.GetMspCertificateOptions)
				getMspCertificateOptionsModel.MspID = core.StringPtr("testString")
				getMspCertificateOptionsModel.Cache = core.StringPtr("skip")
				getMspCertificateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetMspCertificate(getMspCertificateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetMspCertificateOptions model with no property values
				getMspCertificateOptionsModelNew := new(blockchainv2.GetMspCertificateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetMspCertificate(getMspCertificateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`EditAdminCerts(editAdminCertsOptions *EditAdminCertsOptions) - Operation response error`, func() {
		editAdminCertsPath := "/ak/api/v2/kubernetes/components/testString/certs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(editAdminCertsPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke EditAdminCerts with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the EditAdminCertsOptions model
				editAdminCertsOptionsModel := new(blockchainv2.EditAdminCertsOptions)
				editAdminCertsOptionsModel.ID = core.StringPtr("testString")
				editAdminCertsOptionsModel.AppendAdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editAdminCertsOptionsModel.RemoveAdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editAdminCertsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.EditAdminCerts(editAdminCertsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`EditAdminCerts(editAdminCertsOptions *EditAdminCertsOptions)`, func() {
		editAdminCertsPath := "/ak/api/v2/kubernetes/components/testString/certs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(editAdminCertsPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"changes_made": 1, "set_admin_certs": [{"base_64_pem": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "issuer": "/C=US/ST=North Carolina/O=Hyperledger/OU=Fabric/CN=fabric-ca-server", "not_after_ts": 1597770420000, "not_before_ts": 1566234120000, "serial_number_hex": "649a1206fd0bc8be994886dd715cecb0a7a21276", "signature_algorithm": "SHA256withECDSA", "subject": "/OU=client/CN=admin", "X509_version": 3, "time_left": "TimeLeft"}]}`)
				}))
			})
			It(`Invoke EditAdminCerts successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.EditAdminCerts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EditAdminCertsOptions model
				editAdminCertsOptionsModel := new(blockchainv2.EditAdminCertsOptions)
				editAdminCertsOptionsModel.ID = core.StringPtr("testString")
				editAdminCertsOptionsModel.AppendAdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editAdminCertsOptionsModel.RemoveAdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
 				editAdminCertsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.EditAdminCerts(editAdminCertsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke EditAdminCerts with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the EditAdminCertsOptions model
				editAdminCertsOptionsModel := new(blockchainv2.EditAdminCertsOptions)
				editAdminCertsOptionsModel.ID = core.StringPtr("testString")
				editAdminCertsOptionsModel.AppendAdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editAdminCertsOptionsModel.RemoveAdminCerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				editAdminCertsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.EditAdminCerts(editAdminCertsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the EditAdminCertsOptions model with no property values
				editAdminCertsOptionsModelNew := new(blockchainv2.EditAdminCertsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.EditAdminCerts(editAdminCertsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
    Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
				URL: "https://blockchainv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BLOCKCHAIN_URL": "https://blockchainv2/api",
				"BLOCKCHAIN_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := blockchainv2.NewBlockchainV2UsingExternalConfig(&blockchainv2.BlockchainV2Options{
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := blockchainv2.NewBlockchainV2UsingExternalConfig(&blockchainv2.BlockchainV2Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := blockchainv2.NewBlockchainV2UsingExternalConfig(&blockchainv2.BlockchainV2Options{
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BLOCKCHAIN_URL": "https://blockchainv2/api",
				"BLOCKCHAIN_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := blockchainv2.NewBlockchainV2UsingExternalConfig(&blockchainv2.BlockchainV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BLOCKCHAIN_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := blockchainv2.NewBlockchainV2UsingExternalConfig(&blockchainv2.BlockchainV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListComponents(listComponentsOptions *ListComponentsOptions) - Operation response error`, func() {
		listComponentsPath := "/ak/api/v2/components"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listComponentsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["deployment_attrs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["parsed_certs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["cache"]).To(Equal([]string{"skip"}))

					Expect(req.URL.Query()["ca_attrs"]).To(Equal([]string{"included"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListComponents with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListComponentsOptions model
				listComponentsOptionsModel := new(blockchainv2.ListComponentsOptions)
				listComponentsOptionsModel.DeploymentAttrs = core.StringPtr("included")
				listComponentsOptionsModel.ParsedCerts = core.StringPtr("included")
				listComponentsOptionsModel.Cache = core.StringPtr("skip")
				listComponentsOptionsModel.CaAttrs = core.StringPtr("included")
				listComponentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListComponents(listComponentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListComponents(listComponentsOptions *ListComponentsOptions)`, func() {
		listComponentsPath := "/ak/api/v2/components"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listComponentsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["deployment_attrs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["parsed_certs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["cache"]).To(Equal([]string{"skip"}))

					Expect(req.URL.Query()["ca_attrs"]).To(Equal([]string{"included"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"components": [{"id": "myca-2", "type": "fabric-ca", "display_name": "Example CA", "grpcwp_url": "https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "operations_url": "https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443", "msp_id": "Org1", "location": "ibmcloud", "ca_name": "ca", "admin_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="], "node_ou": {"enabled": true}, "ecert": {"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "cacert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}, "state_db": "couchdb", "timestamp": 1537262855753, "resources": {"ca": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "peer": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "orderer": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "statedb": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"ca": {"size": "4GiB", "class": "default"}, "peer": {"size": "4GiB", "class": "default"}, "orderer": {"size": "4GiB", "class": "default"}, "statedb": {"size": "4GiB", "class": "default"}}, "tags": ["Tags"], "tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "version": "Version", "zone": "Zone"}]}`)
				}))
			})
			It(`Invoke ListComponents successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListComponents(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListComponentsOptions model
				listComponentsOptionsModel := new(blockchainv2.ListComponentsOptions)
				listComponentsOptionsModel.DeploymentAttrs = core.StringPtr("included")
				listComponentsOptionsModel.ParsedCerts = core.StringPtr("included")
				listComponentsOptionsModel.Cache = core.StringPtr("skip")
				listComponentsOptionsModel.CaAttrs = core.StringPtr("included")
 				listComponentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListComponents(listComponentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListComponents with error: Operation request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListComponentsOptions model
				listComponentsOptionsModel := new(blockchainv2.ListComponentsOptions)
				listComponentsOptionsModel.DeploymentAttrs = core.StringPtr("included")
				listComponentsOptionsModel.ParsedCerts = core.StringPtr("included")
				listComponentsOptionsModel.Cache = core.StringPtr("skip")
				listComponentsOptionsModel.CaAttrs = core.StringPtr("included")
				listComponentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListComponents(listComponentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetComponentsByType(getComponentsByTypeOptions *GetComponentsByTypeOptions) - Operation response error`, func() {
		getComponentsByTypePath := "/ak/api/v2/components/types/fabric-peer"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getComponentsByTypePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["deployment_attrs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["parsed_certs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["cache"]).To(Equal([]string{"skip"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetComponentsByType with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetComponentsByTypeOptions model
				getComponentsByTypeOptionsModel := new(blockchainv2.GetComponentsByTypeOptions)
				getComponentsByTypeOptionsModel.ComponentType = core.StringPtr("fabric-peer")
				getComponentsByTypeOptionsModel.DeploymentAttrs = core.StringPtr("included")
				getComponentsByTypeOptionsModel.ParsedCerts = core.StringPtr("included")
				getComponentsByTypeOptionsModel.Cache = core.StringPtr("skip")
				getComponentsByTypeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetComponentsByType(getComponentsByTypeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetComponentsByType(getComponentsByTypeOptions *GetComponentsByTypeOptions)`, func() {
		getComponentsByTypePath := "/ak/api/v2/components/types/fabric-peer"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getComponentsByTypePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["deployment_attrs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["parsed_certs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["cache"]).To(Equal([]string{"skip"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"components": [{"id": "myca-2", "type": "fabric-ca", "display_name": "Example CA", "grpcwp_url": "https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "operations_url": "https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443", "msp_id": "Org1", "location": "ibmcloud", "ca_name": "ca", "admin_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="], "node_ou": {"enabled": true}, "ecert": {"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "cacert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}, "state_db": "couchdb", "timestamp": 1537262855753, "resources": {"ca": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "peer": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "orderer": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "statedb": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"ca": {"size": "4GiB", "class": "default"}, "peer": {"size": "4GiB", "class": "default"}, "orderer": {"size": "4GiB", "class": "default"}, "statedb": {"size": "4GiB", "class": "default"}}, "tags": ["Tags"], "tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "version": "Version", "zone": "Zone"}]}`)
				}))
			})
			It(`Invoke GetComponentsByType successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetComponentsByType(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetComponentsByTypeOptions model
				getComponentsByTypeOptionsModel := new(blockchainv2.GetComponentsByTypeOptions)
				getComponentsByTypeOptionsModel.ComponentType = core.StringPtr("fabric-peer")
				getComponentsByTypeOptionsModel.DeploymentAttrs = core.StringPtr("included")
				getComponentsByTypeOptionsModel.ParsedCerts = core.StringPtr("included")
				getComponentsByTypeOptionsModel.Cache = core.StringPtr("skip")
 				getComponentsByTypeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetComponentsByType(getComponentsByTypeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetComponentsByType with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetComponentsByTypeOptions model
				getComponentsByTypeOptionsModel := new(blockchainv2.GetComponentsByTypeOptions)
				getComponentsByTypeOptionsModel.ComponentType = core.StringPtr("fabric-peer")
				getComponentsByTypeOptionsModel.DeploymentAttrs = core.StringPtr("included")
				getComponentsByTypeOptionsModel.ParsedCerts = core.StringPtr("included")
				getComponentsByTypeOptionsModel.Cache = core.StringPtr("skip")
				getComponentsByTypeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetComponentsByType(getComponentsByTypeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetComponentsByTypeOptions model with no property values
				getComponentsByTypeOptionsModelNew := new(blockchainv2.GetComponentsByTypeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetComponentsByType(getComponentsByTypeOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetComponentByTag(getComponentByTagOptions *GetComponentByTagOptions) - Operation response error`, func() {
		getComponentByTagPath := "/ak/api/v2/components/tags/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getComponentByTagPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["deployment_attrs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["parsed_certs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["cache"]).To(Equal([]string{"skip"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetComponentByTag with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetComponentByTagOptions model
				getComponentByTagOptionsModel := new(blockchainv2.GetComponentByTagOptions)
				getComponentByTagOptionsModel.Tag = core.StringPtr("testString")
				getComponentByTagOptionsModel.DeploymentAttrs = core.StringPtr("included")
				getComponentByTagOptionsModel.ParsedCerts = core.StringPtr("included")
				getComponentByTagOptionsModel.Cache = core.StringPtr("skip")
				getComponentByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetComponentByTag(getComponentByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetComponentByTag(getComponentByTagOptions *GetComponentByTagOptions)`, func() {
		getComponentByTagPath := "/ak/api/v2/components/tags/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getComponentByTagPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["deployment_attrs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["parsed_certs"]).To(Equal([]string{"included"}))

					Expect(req.URL.Query()["cache"]).To(Equal([]string{"skip"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"components": [{"id": "myca-2", "type": "fabric-ca", "display_name": "Example CA", "grpcwp_url": "https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084", "api_url": "grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051", "operations_url": "https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443", "msp_id": "Org1", "location": "ibmcloud", "ca_name": "ca", "admin_certs": ["LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="], "node_ou": {"enabled": true}, "ecert": {"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "cacert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}, "state_db": "couchdb", "timestamp": 1537262855753, "resources": {"ca": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "peer": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "orderer": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "proxy": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}, "statedb": {"requests": {"cpu": "40m", "memory": "40M"}, "limits": {"cpu": "8000m", "memory": "16384M"}}}, "scheme_version": "v1", "storage": {"ca": {"size": "4GiB", "class": "default"}, "peer": {"size": "4GiB", "class": "default"}, "orderer": {"size": "4GiB", "class": "default"}, "statedb": {"size": "4GiB", "class": "default"}}, "tags": ["Tags"], "tls_cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=", "version": "Version", "zone": "Zone"}]}`)
				}))
			})
			It(`Invoke GetComponentByTag successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetComponentByTag(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetComponentByTagOptions model
				getComponentByTagOptionsModel := new(blockchainv2.GetComponentByTagOptions)
				getComponentByTagOptionsModel.Tag = core.StringPtr("testString")
				getComponentByTagOptionsModel.DeploymentAttrs = core.StringPtr("included")
				getComponentByTagOptionsModel.ParsedCerts = core.StringPtr("included")
				getComponentByTagOptionsModel.Cache = core.StringPtr("skip")
 				getComponentByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetComponentByTag(getComponentByTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetComponentByTag with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetComponentByTagOptions model
				getComponentByTagOptionsModel := new(blockchainv2.GetComponentByTagOptions)
				getComponentByTagOptionsModel.Tag = core.StringPtr("testString")
				getComponentByTagOptionsModel.DeploymentAttrs = core.StringPtr("included")
				getComponentByTagOptionsModel.ParsedCerts = core.StringPtr("included")
				getComponentByTagOptionsModel.Cache = core.StringPtr("skip")
				getComponentByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetComponentByTag(getComponentByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetComponentByTagOptions model with no property values
				getComponentByTagOptionsModelNew := new(blockchainv2.GetComponentByTagOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetComponentByTag(getComponentByTagOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RemoveComponentsByTag(removeComponentsByTagOptions *RemoveComponentsByTagOptions) - Operation response error`, func() {
		removeComponentsByTagPath := "/ak/api/v2/components/tags/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(removeComponentsByTagPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RemoveComponentsByTag with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the RemoveComponentsByTagOptions model
				removeComponentsByTagOptionsModel := new(blockchainv2.RemoveComponentsByTagOptions)
				removeComponentsByTagOptionsModel.Tag = core.StringPtr("testString")
				removeComponentsByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.RemoveComponentsByTag(removeComponentsByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`RemoveComponentsByTag(removeComponentsByTagOptions *RemoveComponentsByTagOptions)`, func() {
		removeComponentsByTagPath := "/ak/api/v2/components/tags/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(removeComponentsByTagPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"removed": [{"message": "deleted", "type": "fabric-peer", "id": "component-1", "display_name": "My Peer"}]}`)
				}))
			})
			It(`Invoke RemoveComponentsByTag successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.RemoveComponentsByTag(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RemoveComponentsByTagOptions model
				removeComponentsByTagOptionsModel := new(blockchainv2.RemoveComponentsByTagOptions)
				removeComponentsByTagOptionsModel.Tag = core.StringPtr("testString")
 				removeComponentsByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.RemoveComponentsByTag(removeComponentsByTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke RemoveComponentsByTag with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the RemoveComponentsByTagOptions model
				removeComponentsByTagOptionsModel := new(blockchainv2.RemoveComponentsByTagOptions)
				removeComponentsByTagOptionsModel.Tag = core.StringPtr("testString")
				removeComponentsByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.RemoveComponentsByTag(removeComponentsByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RemoveComponentsByTagOptions model with no property values
				removeComponentsByTagOptionsModelNew := new(blockchainv2.RemoveComponentsByTagOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.RemoveComponentsByTag(removeComponentsByTagOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteComponentsByTag(deleteComponentsByTagOptions *DeleteComponentsByTagOptions) - Operation response error`, func() {
		deleteComponentsByTagPath := "/ak/api/v2/kubernetes/components/tags/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteComponentsByTagPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteComponentsByTag with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteComponentsByTagOptions model
				deleteComponentsByTagOptionsModel := new(blockchainv2.DeleteComponentsByTagOptions)
				deleteComponentsByTagOptionsModel.Tag = core.StringPtr("testString")
				deleteComponentsByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.DeleteComponentsByTag(deleteComponentsByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteComponentsByTag(deleteComponentsByTagOptions *DeleteComponentsByTagOptions)`, func() {
		deleteComponentsByTagPath := "/ak/api/v2/kubernetes/components/tags/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteComponentsByTagPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"deleted": [{"message": "deleted", "type": "fabric-peer", "id": "component-1", "display_name": "My Peer"}]}`)
				}))
			})
			It(`Invoke DeleteComponentsByTag successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteComponentsByTag(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteComponentsByTagOptions model
				deleteComponentsByTagOptionsModel := new(blockchainv2.DeleteComponentsByTagOptions)
				deleteComponentsByTagOptionsModel.Tag = core.StringPtr("testString")
 				deleteComponentsByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteComponentsByTag(deleteComponentsByTagOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteComponentsByTag with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteComponentsByTagOptions model
				deleteComponentsByTagOptionsModel := new(blockchainv2.DeleteComponentsByTagOptions)
				deleteComponentsByTagOptionsModel.Tag = core.StringPtr("testString")
				deleteComponentsByTagOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DeleteComponentsByTag(deleteComponentsByTagOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteComponentsByTagOptions model with no property values
				deleteComponentsByTagOptionsModelNew := new(blockchainv2.DeleteComponentsByTagOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.DeleteComponentsByTag(deleteComponentsByTagOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteAllComponents(deleteAllComponentsOptions *DeleteAllComponentsOptions) - Operation response error`, func() {
		deleteAllComponentsPath := "/ak/api/v2/kubernetes/components/purge"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteAllComponentsPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteAllComponents with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteAllComponentsOptions model
				deleteAllComponentsOptionsModel := new(blockchainv2.DeleteAllComponentsOptions)
				deleteAllComponentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.DeleteAllComponents(deleteAllComponentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteAllComponents(deleteAllComponentsOptions *DeleteAllComponentsOptions)`, func() {
		deleteAllComponentsPath := "/ak/api/v2/kubernetes/components/purge"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteAllComponentsPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"deleted": [{"message": "deleted", "type": "fabric-peer", "id": "component-1", "display_name": "My Peer"}]}`)
				}))
			})
			It(`Invoke DeleteAllComponents successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteAllComponents(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteAllComponentsOptions model
				deleteAllComponentsOptionsModel := new(blockchainv2.DeleteAllComponentsOptions)
 				deleteAllComponentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteAllComponents(deleteAllComponentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteAllComponents with error: Operation request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteAllComponentsOptions model
				deleteAllComponentsOptionsModel := new(blockchainv2.DeleteAllComponentsOptions)
				deleteAllComponentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DeleteAllComponents(deleteAllComponentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
    Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
				URL: "https://blockchainv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BLOCKCHAIN_URL": "https://blockchainv2/api",
				"BLOCKCHAIN_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := blockchainv2.NewBlockchainV2UsingExternalConfig(&blockchainv2.BlockchainV2Options{
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := blockchainv2.NewBlockchainV2UsingExternalConfig(&blockchainv2.BlockchainV2Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := blockchainv2.NewBlockchainV2UsingExternalConfig(&blockchainv2.BlockchainV2Options{
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BLOCKCHAIN_URL": "https://blockchainv2/api",
				"BLOCKCHAIN_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := blockchainv2.NewBlockchainV2UsingExternalConfig(&blockchainv2.BlockchainV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BLOCKCHAIN_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := blockchainv2.NewBlockchainV2UsingExternalConfig(&blockchainv2.BlockchainV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetSettings(getSettingsOptions *GetSettingsOptions) - Operation response error`, func() {
		getSettingsPath := "/ak/api/v2/settings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSettings with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(blockchainv2.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSettings(getSettingsOptions *GetSettingsOptions)`, func() {
		getSettingsPath := "/ak/api/v2/settings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"ACTIVITY_TRACKER_PATH": "/logs", "ATHENA_ID": "17v7e", "AUTH_SCHEME": "iam", "CALLBACK_URI": "/auth/cb", "CLUSTER_DATA": {"type": "paid"}, "CONFIGTXLATOR_URL": "https://n3a3ec3-configtxlator.ibp.us-south.containers.appdomain.cloud", "CRN": {"account_id": "a/abcd", "c_name": "staging", "c_type": "public", "instance_id": "abc123", "location": "us-south", "resource_id": "ResourceID", "resource_type": "ResourceType", "service_name": "blockchain", "version": "v1"}, "CRN_STRING": "crn:v1:staging:public:blockchain:us-south:a/abcd:abc123::", "CSP_HEADER_VALUES": ["CSPHEADERVALUES"], "DB_SYSTEM": "system", "DEPLOYER_URL": "https://api.dev.blockchain.cloud.ibm.com", "DOMAIN": "DOMAIN", "ENVIRONMENT": "ENVIRONMENT", "FABRIC_CAPABILITIES": {"application": ["V1_1"], "channel": ["V1_1"], "orderer": ["V1_1"]}, "FEATURE_FLAGS": {"anyKey": "anyValue"}, "FILE_LOGGING": {"server": {"client": {"enabled": true, "level": "silly", "unique_name": false}, "server": {"enabled": true, "level": "silly", "unique_name": false}}, "client": {"client": {"enabled": true, "level": "silly", "unique_name": false}, "server": {"enabled": true, "level": "silly", "unique_name": false}}}, "HOST_URL": "http://localhost:3000", "IAM_CACHE_ENABLED": true, "IAM_URL": "IAMURL", "IBM_ID_CALLBACK_URL": "IBMIDCALLBACKURL", "IGNORE_CONFIG_FILE": true, "INACTIVITY_TIMEOUTS": {"enabled": true, "max_idle_time": 60000}, "INFRASTRUCTURE": "ibmcloud", "LANDING_URL": "http://localhost:3000", "LOGIN_URI": "/auth/login", "LOGOUT_URI": "/auth/logout", "MAX_REQ_PER_MIN": 25, "MAX_REQ_PER_MIN_AK": 25, "MEMORY_CACHE_ENABLED": true, "PORT": "3000", "PROXY_CACHE_ENABLED": true, "PROXY_TLS_FABRIC_REQS": "PROXYTLSFABRICREQS", "PROXY_TLS_HTTP_URL": "PROXYTLSHTTPURL", "PROXY_TLS_WS_URL": "anyValue", "REGION": "REGION", "SESSION_CACHE_ENABLED": true, "TIMEOUTS": {"anyKey": "anyValue"}, "TIMESTAMPS": {"now": 1542746836056, "born": 1542746836056, "next_settings_update": "1.2 mins", "up_time": "30 days"}, "TRANSACTION_VISIBILITY": {"anyKey": "anyValue"}, "TRUST_PROXY": "loopback", "TRUST_UNKNOWN_CERTS": true, "VERSIONS": {"apollo": "65f3cbfd", "athena": "1198f94", "stitch": "0f1a0c6", "tag": "v0.4.31"}}`)
				}))
			})
			It(`Invoke GetSettings successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(blockchainv2.GetSettingsOptions)
 				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetSettings with error: Operation request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(blockchainv2.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`EditSettings(editSettingsOptions *EditSettingsOptions) - Operation response error`, func() {
		editSettingsPath := "/ak/api/v2/settings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(editSettingsPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke EditSettings with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the LoggingSettingsClient model
				loggingSettingsClientModel := new(blockchainv2.LoggingSettingsClient)
				loggingSettingsClientModel.Enabled = core.BoolPtr(true)
				loggingSettingsClientModel.Level = core.StringPtr("silly")
				loggingSettingsClientModel.UniqueName = core.BoolPtr(false)

				// Construct an instance of the LoggingSettingsServer model
				loggingSettingsServerModel := new(blockchainv2.LoggingSettingsServer)
				loggingSettingsServerModel.Enabled = core.BoolPtr(true)
				loggingSettingsServerModel.Level = core.StringPtr("silly")
				loggingSettingsServerModel.UniqueName = core.BoolPtr(false)

				// Construct an instance of the EditLogSettingsBody model
				editLogSettingsBodyModel := new(blockchainv2.EditLogSettingsBody)
				editLogSettingsBodyModel.Client = loggingSettingsClientModel
				editLogSettingsBodyModel.Server = loggingSettingsServerModel

				// Construct an instance of the EditSettingsBodyInactivityTimeouts model
				editSettingsBodyInactivityTimeoutsModel := new(blockchainv2.EditSettingsBodyInactivityTimeouts)
				editSettingsBodyInactivityTimeoutsModel.Enabled = core.BoolPtr(false)
				editSettingsBodyInactivityTimeoutsModel.MaxIdleTime = core.Float64Ptr(float64(90000))

				// Construct an instance of the EditSettingsOptions model
				editSettingsOptionsModel := new(blockchainv2.EditSettingsOptions)
				editSettingsOptionsModel.InactivityTimeouts = editSettingsBodyInactivityTimeoutsModel
				editSettingsOptionsModel.FileLogging = editLogSettingsBodyModel
				editSettingsOptionsModel.MaxReqPerMin = core.Float64Ptr(float64(25))
				editSettingsOptionsModel.MaxReqPerMinAk = core.Float64Ptr(float64(25))
				editSettingsOptionsModel.FabricGetBlockTimeoutMs = core.Float64Ptr(float64(10000))
				editSettingsOptionsModel.FabricInstantiateTimeoutMs = core.Float64Ptr(float64(300000))
				editSettingsOptionsModel.FabricJoinChannelTimeoutMs = core.Float64Ptr(float64(25000))
				editSettingsOptionsModel.FabricInstallCcTimeoutMs = core.Float64Ptr(float64(300000))
				editSettingsOptionsModel.FabricLcInstallCcTimeoutMs = core.Float64Ptr(float64(300000))
				editSettingsOptionsModel.FabricLcGetCcTimeoutMs = core.Float64Ptr(float64(180000))
				editSettingsOptionsModel.FabricGeneralTimeoutMs = core.Float64Ptr(float64(10000))
				editSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.EditSettings(editSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`EditSettings(editSettingsOptions *EditSettingsOptions)`, func() {
		editSettingsPath := "/ak/api/v2/settings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(editSettingsPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"ACTIVITY_TRACKER_PATH": "/logs", "ATHENA_ID": "17v7e", "AUTH_SCHEME": "iam", "CALLBACK_URI": "/auth/cb", "CLUSTER_DATA": {"type": "paid"}, "CONFIGTXLATOR_URL": "https://n3a3ec3-configtxlator.ibp.us-south.containers.appdomain.cloud", "CRN": {"account_id": "a/abcd", "c_name": "staging", "c_type": "public", "instance_id": "abc123", "location": "us-south", "resource_id": "ResourceID", "resource_type": "ResourceType", "service_name": "blockchain", "version": "v1"}, "CRN_STRING": "crn:v1:staging:public:blockchain:us-south:a/abcd:abc123::", "CSP_HEADER_VALUES": ["CSPHEADERVALUES"], "DB_SYSTEM": "system", "DEPLOYER_URL": "https://api.dev.blockchain.cloud.ibm.com", "DOMAIN": "DOMAIN", "ENVIRONMENT": "ENVIRONMENT", "FABRIC_CAPABILITIES": {"application": ["V1_1"], "channel": ["V1_1"], "orderer": ["V1_1"]}, "FEATURE_FLAGS": {"anyKey": "anyValue"}, "FILE_LOGGING": {"server": {"client": {"enabled": true, "level": "silly", "unique_name": false}, "server": {"enabled": true, "level": "silly", "unique_name": false}}, "client": {"client": {"enabled": true, "level": "silly", "unique_name": false}, "server": {"enabled": true, "level": "silly", "unique_name": false}}}, "HOST_URL": "http://localhost:3000", "IAM_CACHE_ENABLED": true, "IAM_URL": "IAMURL", "IBM_ID_CALLBACK_URL": "IBMIDCALLBACKURL", "IGNORE_CONFIG_FILE": true, "INACTIVITY_TIMEOUTS": {"enabled": true, "max_idle_time": 60000}, "INFRASTRUCTURE": "ibmcloud", "LANDING_URL": "http://localhost:3000", "LOGIN_URI": "/auth/login", "LOGOUT_URI": "/auth/logout", "MAX_REQ_PER_MIN": 25, "MAX_REQ_PER_MIN_AK": 25, "MEMORY_CACHE_ENABLED": true, "PORT": "3000", "PROXY_CACHE_ENABLED": true, "PROXY_TLS_FABRIC_REQS": "PROXYTLSFABRICREQS", "PROXY_TLS_HTTP_URL": "PROXYTLSHTTPURL", "PROXY_TLS_WS_URL": "anyValue", "REGION": "REGION", "SESSION_CACHE_ENABLED": true, "TIMEOUTS": {"anyKey": "anyValue"}, "TIMESTAMPS": {"now": 1542746836056, "born": 1542746836056, "next_settings_update": "1.2 mins", "up_time": "30 days"}, "TRANSACTION_VISIBILITY": {"anyKey": "anyValue"}, "TRUST_PROXY": "loopback", "TRUST_UNKNOWN_CERTS": true, "VERSIONS": {"apollo": "65f3cbfd", "athena": "1198f94", "stitch": "0f1a0c6", "tag": "v0.4.31"}}`)
				}))
			})
			It(`Invoke EditSettings successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.EditSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LoggingSettingsClient model
				loggingSettingsClientModel := new(blockchainv2.LoggingSettingsClient)
				loggingSettingsClientModel.Enabled = core.BoolPtr(true)
				loggingSettingsClientModel.Level = core.StringPtr("silly")
				loggingSettingsClientModel.UniqueName = core.BoolPtr(false)

				// Construct an instance of the LoggingSettingsServer model
				loggingSettingsServerModel := new(blockchainv2.LoggingSettingsServer)
				loggingSettingsServerModel.Enabled = core.BoolPtr(true)
				loggingSettingsServerModel.Level = core.StringPtr("silly")
				loggingSettingsServerModel.UniqueName = core.BoolPtr(false)

				// Construct an instance of the EditLogSettingsBody model
				editLogSettingsBodyModel := new(blockchainv2.EditLogSettingsBody)
				editLogSettingsBodyModel.Client = loggingSettingsClientModel
				editLogSettingsBodyModel.Server = loggingSettingsServerModel

				// Construct an instance of the EditSettingsBodyInactivityTimeouts model
				editSettingsBodyInactivityTimeoutsModel := new(blockchainv2.EditSettingsBodyInactivityTimeouts)
				editSettingsBodyInactivityTimeoutsModel.Enabled = core.BoolPtr(false)
				editSettingsBodyInactivityTimeoutsModel.MaxIdleTime = core.Float64Ptr(float64(90000))

				// Construct an instance of the EditSettingsOptions model
				editSettingsOptionsModel := new(blockchainv2.EditSettingsOptions)
				editSettingsOptionsModel.InactivityTimeouts = editSettingsBodyInactivityTimeoutsModel
				editSettingsOptionsModel.FileLogging = editLogSettingsBodyModel
				editSettingsOptionsModel.MaxReqPerMin = core.Float64Ptr(float64(25))
				editSettingsOptionsModel.MaxReqPerMinAk = core.Float64Ptr(float64(25))
				editSettingsOptionsModel.FabricGetBlockTimeoutMs = core.Float64Ptr(float64(10000))
				editSettingsOptionsModel.FabricInstantiateTimeoutMs = core.Float64Ptr(float64(300000))
				editSettingsOptionsModel.FabricJoinChannelTimeoutMs = core.Float64Ptr(float64(25000))
				editSettingsOptionsModel.FabricInstallCcTimeoutMs = core.Float64Ptr(float64(300000))
				editSettingsOptionsModel.FabricLcInstallCcTimeoutMs = core.Float64Ptr(float64(300000))
				editSettingsOptionsModel.FabricLcGetCcTimeoutMs = core.Float64Ptr(float64(180000))
				editSettingsOptionsModel.FabricGeneralTimeoutMs = core.Float64Ptr(float64(10000))
 				editSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.EditSettings(editSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke EditSettings with error: Operation request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the LoggingSettingsClient model
				loggingSettingsClientModel := new(blockchainv2.LoggingSettingsClient)
				loggingSettingsClientModel.Enabled = core.BoolPtr(true)
				loggingSettingsClientModel.Level = core.StringPtr("silly")
				loggingSettingsClientModel.UniqueName = core.BoolPtr(false)

				// Construct an instance of the LoggingSettingsServer model
				loggingSettingsServerModel := new(blockchainv2.LoggingSettingsServer)
				loggingSettingsServerModel.Enabled = core.BoolPtr(true)
				loggingSettingsServerModel.Level = core.StringPtr("silly")
				loggingSettingsServerModel.UniqueName = core.BoolPtr(false)

				// Construct an instance of the EditLogSettingsBody model
				editLogSettingsBodyModel := new(blockchainv2.EditLogSettingsBody)
				editLogSettingsBodyModel.Client = loggingSettingsClientModel
				editLogSettingsBodyModel.Server = loggingSettingsServerModel

				// Construct an instance of the EditSettingsBodyInactivityTimeouts model
				editSettingsBodyInactivityTimeoutsModel := new(blockchainv2.EditSettingsBodyInactivityTimeouts)
				editSettingsBodyInactivityTimeoutsModel.Enabled = core.BoolPtr(false)
				editSettingsBodyInactivityTimeoutsModel.MaxIdleTime = core.Float64Ptr(float64(90000))

				// Construct an instance of the EditSettingsOptions model
				editSettingsOptionsModel := new(blockchainv2.EditSettingsOptions)
				editSettingsOptionsModel.InactivityTimeouts = editSettingsBodyInactivityTimeoutsModel
				editSettingsOptionsModel.FileLogging = editLogSettingsBodyModel
				editSettingsOptionsModel.MaxReqPerMin = core.Float64Ptr(float64(25))
				editSettingsOptionsModel.MaxReqPerMinAk = core.Float64Ptr(float64(25))
				editSettingsOptionsModel.FabricGetBlockTimeoutMs = core.Float64Ptr(float64(10000))
				editSettingsOptionsModel.FabricInstantiateTimeoutMs = core.Float64Ptr(float64(300000))
				editSettingsOptionsModel.FabricJoinChannelTimeoutMs = core.Float64Ptr(float64(25000))
				editSettingsOptionsModel.FabricInstallCcTimeoutMs = core.Float64Ptr(float64(300000))
				editSettingsOptionsModel.FabricLcInstallCcTimeoutMs = core.Float64Ptr(float64(300000))
				editSettingsOptionsModel.FabricLcGetCcTimeoutMs = core.Float64Ptr(float64(180000))
				editSettingsOptionsModel.FabricGeneralTimeoutMs = core.Float64Ptr(float64(10000))
				editSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.EditSettings(editSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetFabVersions(getFabVersionsOptions *GetFabVersionsOptions) - Operation response error`, func() {
		getFabVersionsPath := "/ak/api/v2/kubernetes/fabric/versions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getFabVersionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["cache"]).To(Equal([]string{"skip"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetFabVersions with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetFabVersionsOptions model
				getFabVersionsOptionsModel := new(blockchainv2.GetFabVersionsOptions)
				getFabVersionsOptionsModel.Cache = core.StringPtr("skip")
				getFabVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetFabVersions(getFabVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetFabVersions(getFabVersionsOptions *GetFabVersionsOptions)`, func() {
		getFabVersionsPath := "/ak/api/v2/kubernetes/fabric/versions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getFabVersionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["cache"]).To(Equal([]string{"skip"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"versions": {"ca": {"1.4.6-2": {"default": true, "version": "1.4.6-2", "image": {"anyKey": "anyValue"}}, "2.1.0-0": {"default": true, "version": "1.4.6-2", "image": {"anyKey": "anyValue"}}}, "peer": {"1.4.6-2": {"default": true, "version": "1.4.6-2", "image": {"anyKey": "anyValue"}}, "2.1.0-0": {"default": true, "version": "1.4.6-2", "image": {"anyKey": "anyValue"}}}, "orderer": {"1.4.6-2": {"default": true, "version": "1.4.6-2", "image": {"anyKey": "anyValue"}}, "2.1.0-0": {"default": true, "version": "1.4.6-2", "image": {"anyKey": "anyValue"}}}}}`)
				}))
			})
			It(`Invoke GetFabVersions successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetFabVersions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetFabVersionsOptions model
				getFabVersionsOptionsModel := new(blockchainv2.GetFabVersionsOptions)
				getFabVersionsOptionsModel.Cache = core.StringPtr("skip")
 				getFabVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetFabVersions(getFabVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetFabVersions with error: Operation request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetFabVersionsOptions model
				getFabVersionsOptionsModel := new(blockchainv2.GetFabVersionsOptions)
				getFabVersionsOptionsModel.Cache = core.StringPtr("skip")
				getFabVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetFabVersions(getFabVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetHealth(getHealthOptions *GetHealthOptions) - Operation response error`, func() {
		getHealthPath := "/ak/api/v2/health"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getHealthPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetHealth with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetHealthOptions model
				getHealthOptionsModel := new(blockchainv2.GetHealthOptions)
				getHealthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetHealth(getHealthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetHealth(getHealthOptions *GetHealthOptions)`, func() {
		getHealthPath := "/ak/api/v2/health"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getHealthPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"OPTOOLS": {"instance_id": "p59ta", "now": 1542746836056, "born": 1542746836056, "up_time": "30 days", "memory_usage": {"rss": "56.1 MB", "heapTotal": "34.4 MB", "heapUsed": "28.4 MB", "external": "369.3 KB"}, "session_cache_stats": {"hits": 42, "misses": 11, "keys": 4, "cache_size": "CacheSize"}, "couch_cache_stats": {"hits": 42, "misses": 11, "keys": 4, "cache_size": "CacheSize"}, "iam_cache_stats": {"hits": 42, "misses": 11, "keys": 4, "cache_size": "CacheSize"}, "proxy_cache": {"hits": 42, "misses": 11, "keys": 4, "cache_size": "CacheSize"}}, "OS": {"arch": "x64", "type": "Windows_NT", "endian": "LE", "loadavg": "[0,0,0]", "cpus": [{"model": "Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz", "speed": "2592", "times": {"idle": 131397203, "irq": 6068640, "nice": 0, "sys": 9652328, "user": 4152187}}], "total_memory": "31.7 GB", "free_memory": "21.9 GB", "up_time": "4.9 days"}}`)
				}))
			})
			It(`Invoke GetHealth successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetHealth(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetHealthOptions model
				getHealthOptionsModel := new(blockchainv2.GetHealthOptions)
 				getHealthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetHealth(getHealthOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetHealth with error: Operation request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetHealthOptions model
				getHealthOptionsModel := new(blockchainv2.GetHealthOptions)
				getHealthOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetHealth(getHealthOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListNotifications(listNotificationsOptions *ListNotificationsOptions) - Operation response error`, func() {
		listNotificationsPath := "/ak/api/v2/notifications"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listNotificationsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for limit query parameter


					// TODO: Add check for skip query parameter

					Expect(req.URL.Query()["component_id"]).To(Equal([]string{"MyPeer"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListNotifications with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListNotificationsOptions model
				listNotificationsOptionsModel := new(blockchainv2.ListNotificationsOptions)
				listNotificationsOptionsModel.Limit = core.Float64Ptr(float64(1))
				listNotificationsOptionsModel.Skip = core.Float64Ptr(float64(1))
				listNotificationsOptionsModel.ComponentID = core.StringPtr("MyPeer")
				listNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListNotifications(listNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListNotifications(listNotificationsOptions *ListNotificationsOptions)`, func() {
		listNotificationsPath := "/ak/api/v2/notifications"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listNotificationsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for limit query parameter


					// TODO: Add check for skip query parameter

					Expect(req.URL.Query()["component_id"]).To(Equal([]string{"MyPeer"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"total": 10, "returning": 3, "notifications": [{"id": "60d84819bfa17adb4174ff3a1c52b5d6", "type": "notification", "status": "pending", "by": "By", "message": "Restarting application", "ts_display": 1537262855753}]}`)
				}))
			})
			It(`Invoke ListNotifications successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListNotifications(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListNotificationsOptions model
				listNotificationsOptionsModel := new(blockchainv2.ListNotificationsOptions)
				listNotificationsOptionsModel.Limit = core.Float64Ptr(float64(1))
				listNotificationsOptionsModel.Skip = core.Float64Ptr(float64(1))
				listNotificationsOptionsModel.ComponentID = core.StringPtr("MyPeer")
 				listNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListNotifications(listNotificationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListNotifications with error: Operation request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListNotificationsOptions model
				listNotificationsOptionsModel := new(blockchainv2.ListNotificationsOptions)
				listNotificationsOptionsModel.Limit = core.Float64Ptr(float64(1))
				listNotificationsOptionsModel.Skip = core.Float64Ptr(float64(1))
				listNotificationsOptionsModel.ComponentID = core.StringPtr("MyPeer")
				listNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListNotifications(listNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteSigTx(deleteSigTxOptions *DeleteSigTxOptions) - Operation response error`, func() {
		deleteSigTxPath := "/ak/api/v2/signature_collections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteSigTxPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteSigTx with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteSigTxOptions model
				deleteSigTxOptionsModel := new(blockchainv2.DeleteSigTxOptions)
				deleteSigTxOptionsModel.ID = core.StringPtr("testString")
				deleteSigTxOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.DeleteSigTx(deleteSigTxOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteSigTx(deleteSigTxOptions *DeleteSigTxOptions)`, func() {
		deleteSigTxPath := "/ak/api/v2/signature_collections/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteSigTxPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"message": "ok", "tx_id": "abcde"}`)
				}))
			})
			It(`Invoke DeleteSigTx successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteSigTx(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteSigTxOptions model
				deleteSigTxOptionsModel := new(blockchainv2.DeleteSigTxOptions)
				deleteSigTxOptionsModel.ID = core.StringPtr("testString")
 				deleteSigTxOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteSigTx(deleteSigTxOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteSigTx with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteSigTxOptions model
				deleteSigTxOptionsModel := new(blockchainv2.DeleteSigTxOptions)
				deleteSigTxOptionsModel.ID = core.StringPtr("testString")
				deleteSigTxOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DeleteSigTx(deleteSigTxOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteSigTxOptions model with no property values
				deleteSigTxOptionsModelNew := new(blockchainv2.DeleteSigTxOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.DeleteSigTx(deleteSigTxOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ArchiveNotifications(archiveNotificationsOptions *ArchiveNotificationsOptions) - Operation response error`, func() {
		archiveNotificationsPath := "/ak/api/v2/notifications/bulk"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(archiveNotificationsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ArchiveNotifications with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ArchiveNotificationsOptions model
				archiveNotificationsOptionsModel := new(blockchainv2.ArchiveNotificationsOptions)
				archiveNotificationsOptionsModel.NotificationIds = []string{"c9d00ebf849051e4f102008dc0be2488"}
				archiveNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ArchiveNotifications(archiveNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ArchiveNotifications(archiveNotificationsOptions *ArchiveNotificationsOptions)`, func() {
		archiveNotificationsPath := "/ak/api/v2/notifications/bulk"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(archiveNotificationsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"message": "ok", "details": "archived 3 notification(s)"}`)
				}))
			})
			It(`Invoke ArchiveNotifications successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ArchiveNotifications(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ArchiveNotificationsOptions model
				archiveNotificationsOptionsModel := new(blockchainv2.ArchiveNotificationsOptions)
				archiveNotificationsOptionsModel.NotificationIds = []string{"c9d00ebf849051e4f102008dc0be2488"}
 				archiveNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ArchiveNotifications(archiveNotificationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ArchiveNotifications with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ArchiveNotificationsOptions model
				archiveNotificationsOptionsModel := new(blockchainv2.ArchiveNotificationsOptions)
				archiveNotificationsOptionsModel.NotificationIds = []string{"c9d00ebf849051e4f102008dc0be2488"}
				archiveNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ArchiveNotifications(archiveNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ArchiveNotificationsOptions model with no property values
				archiveNotificationsOptionsModelNew := new(blockchainv2.ArchiveNotificationsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ArchiveNotifications(archiveNotificationsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Restart(restartOptions *RestartOptions) - Operation response error`, func() {
		restartPath := "/ak/api/v2/restart"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(restartPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Restart with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the RestartOptions model
				restartOptionsModel := new(blockchainv2.RestartOptions)
				restartOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.Restart(restartOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`Restart(restartOptions *RestartOptions)`, func() {
		restartPath := "/ak/api/v2/restart"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(restartPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"message": "restarting - give me 10 seconds"}`)
				}))
			})
			It(`Invoke Restart successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.Restart(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RestartOptions model
				restartOptionsModel := new(blockchainv2.RestartOptions)
 				restartOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.Restart(restartOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke Restart with error: Operation request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the RestartOptions model
				restartOptionsModel := new(blockchainv2.RestartOptions)
				restartOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.Restart(restartOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteAllSessions(deleteAllSessionsOptions *DeleteAllSessionsOptions) - Operation response error`, func() {
		deleteAllSessionsPath := "/ak/api/v2/sessions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteAllSessionsPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteAllSessions with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteAllSessionsOptions model
				deleteAllSessionsOptionsModel := new(blockchainv2.DeleteAllSessionsOptions)
				deleteAllSessionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.DeleteAllSessions(deleteAllSessionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteAllSessions(deleteAllSessionsOptions *DeleteAllSessionsOptions)`, func() {
		deleteAllSessionsPath := "/ak/api/v2/sessions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteAllSessionsPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"message": "ok", "deleted": 42}`)
				}))
			})
			It(`Invoke DeleteAllSessions successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteAllSessions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteAllSessionsOptions model
				deleteAllSessionsOptionsModel := new(blockchainv2.DeleteAllSessionsOptions)
 				deleteAllSessionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteAllSessions(deleteAllSessionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteAllSessions with error: Operation request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteAllSessionsOptions model
				deleteAllSessionsOptionsModel := new(blockchainv2.DeleteAllSessionsOptions)
				deleteAllSessionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DeleteAllSessions(deleteAllSessionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteAllNotifications(deleteAllNotificationsOptions *DeleteAllNotificationsOptions) - Operation response error`, func() {
		deleteAllNotificationsPath := "/ak/api/v2/notifications/purge"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteAllNotificationsPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteAllNotifications with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteAllNotificationsOptions model
				deleteAllNotificationsOptionsModel := new(blockchainv2.DeleteAllNotificationsOptions)
				deleteAllNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.DeleteAllNotifications(deleteAllNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteAllNotifications(deleteAllNotificationsOptions *DeleteAllNotificationsOptions)`, func() {
		deleteAllNotificationsPath := "/ak/api/v2/notifications/purge"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteAllNotificationsPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"message": "ok", "details": "deleted 101 notification(s)"}`)
				}))
			})
			It(`Invoke DeleteAllNotifications successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteAllNotifications(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteAllNotificationsOptions model
				deleteAllNotificationsOptionsModel := new(blockchainv2.DeleteAllNotificationsOptions)
 				deleteAllNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteAllNotifications(deleteAllNotificationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteAllNotifications with error: Operation request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteAllNotificationsOptions model
				deleteAllNotificationsOptionsModel := new(blockchainv2.DeleteAllNotificationsOptions)
				deleteAllNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DeleteAllNotifications(deleteAllNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ClearCaches(clearCachesOptions *ClearCachesOptions) - Operation response error`, func() {
		clearCachesPath := "/ak/api/v2/cache"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(clearCachesPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ClearCaches with error: Operation response processing error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ClearCachesOptions model
				clearCachesOptionsModel := new(blockchainv2.ClearCachesOptions)
				clearCachesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ClearCaches(clearCachesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ClearCaches(clearCachesOptions *ClearCachesOptions)`, func() {
		clearCachesPath := "/ak/api/v2/cache"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(clearCachesPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"message": "ok", "flushed": ["couch_cache"]}`)
				}))
			})
			It(`Invoke ClearCaches successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ClearCaches(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ClearCachesOptions model
				clearCachesOptionsModel := new(blockchainv2.ClearCachesOptions)
 				clearCachesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ClearCaches(clearCachesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ClearCaches with error: Operation request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ClearCachesOptions model
				clearCachesOptionsModel := new(blockchainv2.ClearCachesOptions)
				clearCachesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ClearCaches(clearCachesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
    Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
				URL: "https://blockchainv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BLOCKCHAIN_URL": "https://blockchainv2/api",
				"BLOCKCHAIN_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := blockchainv2.NewBlockchainV2UsingExternalConfig(&blockchainv2.BlockchainV2Options{
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := blockchainv2.NewBlockchainV2UsingExternalConfig(&blockchainv2.BlockchainV2Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := blockchainv2.NewBlockchainV2UsingExternalConfig(&blockchainv2.BlockchainV2Options{
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BLOCKCHAIN_URL": "https://blockchainv2/api",
				"BLOCKCHAIN_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := blockchainv2.NewBlockchainV2UsingExternalConfig(&blockchainv2.BlockchainV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"BLOCKCHAIN_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := blockchainv2.NewBlockchainV2UsingExternalConfig(&blockchainv2.BlockchainV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})

	Describe(`GetPostman(getPostmanOptions *GetPostmanOptions)`, func() {
		getPostmanPath := "/ak/api/v2/postman"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getPostmanPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["auth_type"]).To(Equal([]string{"bearer"}))

					Expect(req.URL.Query()["token"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["api_key"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["username"]).To(Equal([]string{"admin"}))

					Expect(req.URL.Query()["password"]).To(Equal([]string{"password"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetPostman successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.GetPostman(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GetPostmanOptions model
				getPostmanOptionsModel := new(blockchainv2.GetPostmanOptions)
				getPostmanOptionsModel.AuthType = core.StringPtr("bearer")
				getPostmanOptionsModel.Token = core.StringPtr("testString")
				getPostmanOptionsModel.ApiKey = core.StringPtr("testString")
				getPostmanOptionsModel.Username = core.StringPtr("admin")
				getPostmanOptionsModel.Password = core.StringPtr("password")
 				getPostmanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.GetPostman(getPostmanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke GetPostman with error: Operation validation and request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetPostmanOptions model
				getPostmanOptionsModel := new(blockchainv2.GetPostmanOptions)
				getPostmanOptionsModel.AuthType = core.StringPtr("bearer")
				getPostmanOptionsModel.Token = core.StringPtr("testString")
				getPostmanOptionsModel.ApiKey = core.StringPtr("testString")
				getPostmanOptionsModel.Username = core.StringPtr("admin")
				getPostmanOptionsModel.Password = core.StringPtr("password")
				getPostmanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.GetPostman(getPostmanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the GetPostmanOptions model with no property values
				getPostmanOptionsModelNew := new(blockchainv2.GetPostmanOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.GetPostman(getPostmanOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSwagger(getSwaggerOptions *GetSwaggerOptions)`, func() {
		getSwaggerPath := "/ak/api/v2/openapi"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getSwaggerPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "text/plain")
					res.WriteHeader(200)
					fmt.Fprintf(res, `"OperationResponse"`)
				}))
			})
			It(`Invoke GetSwagger successfully`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetSwagger(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSwaggerOptions model
				getSwaggerOptionsModel := new(blockchainv2.GetSwaggerOptions)
 				getSwaggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetSwagger(getSwaggerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetSwagger with error: Operation request error`, func() {
				testService, testServiceErr := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetSwaggerOptions model
				getSwaggerOptionsModel := new(blockchainv2.GetSwaggerOptions)
				getSwaggerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetSwagger(getSwaggerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			testService, _ := blockchainv2.NewBlockchainV2(&blockchainv2.BlockchainV2Options{
				URL:           "http://blockchainv2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewArchiveNotificationsOptions successfully`, func() {
				// Construct an instance of the ArchiveNotificationsOptions model
				archiveNotificationsOptionsNotificationIds := []string{"c9d00ebf849051e4f102008dc0be2488"}
				archiveNotificationsOptionsModel := testService.NewArchiveNotificationsOptions(archiveNotificationsOptionsNotificationIds)
				archiveNotificationsOptionsModel.SetNotificationIds([]string{"c9d00ebf849051e4f102008dc0be2488"})
				archiveNotificationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(archiveNotificationsOptionsModel).ToNot(BeNil())
				Expect(archiveNotificationsOptionsModel.NotificationIds).To(Equal([]string{"c9d00ebf849051e4f102008dc0be2488"}))
				Expect(archiveNotificationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewBccspPKCS11 successfully`, func() {
				label := "testString"
				pin := "testString"
				model, err := testService.NewBccspPKCS11(label, pin)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewBccspSW successfully`, func() {
				hash := "SHA2"
				security := float64(256)
				model, err := testService.NewBccspSW(hash, security)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewClearCachesOptions successfully`, func() {
				// Construct an instance of the ClearCachesOptions model
				clearCachesOptionsModel := testService.NewClearCachesOptions()
				clearCachesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(clearCachesOptionsModel).ToNot(BeNil())
				Expect(clearCachesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewConfigCACfgIdentities successfully`, func() {
				passwordattempts := float64(10)
				model, err := testService.NewConfigCACfgIdentities(passwordattempts)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCACreate successfully`, func() {
				var registry *blockchainv2.ConfigCARegistry = nil
				_, err := testService.NewConfigCACreate(registry)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewConfigCACsrKeyrequest successfully`, func() {
				algo := "ecdsa"
				size := float64(256)
				model, err := testService.NewConfigCACsrKeyrequest(algo, size)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCACsrNamesItem successfully`, func() {
				c := "US"
				st := "North Carolina"
				o := "Hyperledger"
				model, err := testService.NewConfigCACsrNamesItem(c, st, o)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCADbTlsClient successfully`, func() {
				certfile := "testString"
				keyfile := "testString"
				model, err := testService.NewConfigCADbTlsClient(certfile, keyfile)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCAIntermediateEnrollment successfully`, func() {
				hosts := "localhost"
				profile := "testString"
				label := "testString"
				model, err := testService.NewConfigCAIntermediateEnrollment(hosts, profile, label)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCAIntermediateParentserver successfully`, func() {
				url := "testString"
				caname := "testString"
				model, err := testService.NewConfigCAIntermediateParentserver(url, caname)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCAIntermediateTls successfully`, func() {
				certfiles := []string{"testString"}
				model, err := testService.NewConfigCAIntermediateTls(certfiles)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCAIntermediateTlsClient successfully`, func() {
				certfile := "testString"
				keyfile := "testString"
				model, err := testService.NewConfigCAIntermediateTlsClient(certfile, keyfile)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCARegistryIdentitiesItem successfully`, func() {
				name := "admin"
				pass := "password"
				typeVar := "client"
				model, err := testService.NewConfigCARegistryIdentitiesItem(name, pass, typeVar)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCATlsClientauth successfully`, func() {
				typeVar := "noclientcert"
				certfiles := []string{"testString"}
				model, err := testService.NewConfigCATlsClientauth(typeVar, certfiles)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCACfg successfully`, func() {
				var identities *blockchainv2.ConfigCACfgIdentities = nil
				_, err := testService.NewConfigCACfg(identities)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewConfigCACors successfully`, func() {
				enabled := true
				origins := []string{"*"}
				model, err := testService.NewConfigCACors(enabled, origins)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCACrl successfully`, func() {
				expiry := "24h"
				model, err := testService.NewConfigCACrl(expiry)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCACsr successfully`, func() {
				cn := "ca"
				names := []blockchainv2.ConfigCACsrNamesItem{}
				var ca *blockchainv2.ConfigCACsrCa = nil
				_, err := testService.NewConfigCACsr(cn, names, ca)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewConfigCADb successfully`, func() {
				typeVar := "postgres"
				datasource := "host=fake.databases.appdomain.cloud port=31941 user=ibm_cloud password=password dbname=ibmclouddb sslmode=verify-full"
				model, err := testService.NewConfigCADb(typeVar, datasource)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCAIdemix successfully`, func() {
				rhpoolsize := float64(100)
				nonceexpiration := "15s"
				noncesweepinterval := "15m"
				model, err := testService.NewConfigCAIdemix(rhpoolsize, nonceexpiration, noncesweepinterval)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCAIntermediate successfully`, func() {
				var parentserver *blockchainv2.ConfigCAIntermediateParentserver = nil
				_, err := testService.NewConfigCAIntermediate(parentserver)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewConfigCARegistry successfully`, func() {
				maxenrollments := float64(-1)
				identities := []blockchainv2.ConfigCARegistryIdentitiesItem{}
				model, err := testService.NewConfigCARegistry(maxenrollments, identities)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigCATls successfully`, func() {
				keyfile := "testString"
				certfile := "testString"
				model, err := testService.NewConfigCATls(keyfile, certfile)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigObjectEnrollment successfully`, func() {
				var component *blockchainv2.ConfigObjectEnrollmentComponent = nil
				var tls *blockchainv2.ConfigObjectEnrollmentTls = nil
				_, err := testService.NewConfigObjectEnrollment(component, tls)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewConfigObjectEnrollmentComponent successfully`, func() {
				cahost := "n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud"
				caport := float64(7054)
				caname := "ca"
				var catls *blockchainv2.ConfigObjectEnrollmentComponentCatls = nil
				enrollid := "admin"
				enrollsecret := "password"
				_, err := testService.NewConfigObjectEnrollmentComponent(cahost, caport, caname, catls, enrollid, enrollsecret)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewConfigObjectEnrollmentComponentCatls successfully`, func() {
				cacert := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
				model, err := testService.NewConfigObjectEnrollmentComponentCatls(cacert)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigObjectEnrollmentTls successfully`, func() {
				cahost := "n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud"
				caport := float64(7054)
				caname := "tlsca"
				var catls *blockchainv2.ConfigObjectEnrollmentTlsCatls = nil
				_, err := testService.NewConfigObjectEnrollmentTls(cahost, caport, caname, catls)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewConfigObjectEnrollmentTlsCatls successfully`, func() {
				cacert := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
				model, err := testService.NewConfigObjectEnrollmentTlsCatls(cacert)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigObjectEnrollmentTlsCsr successfully`, func() {
				hosts := []string{"testString"}
				model, err := testService.NewConfigObjectEnrollmentTlsCsr(hosts)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigObjectMsp successfully`, func() {
				var component *blockchainv2.MspConfigData = nil
				var tls *blockchainv2.MspConfigData = nil
				_, err := testService.NewConfigObjectMsp(component, tls)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewConfigPeerAdminService successfully`, func() {
				listenAddress := "0.0.0.0:7051"
				model, err := testService.NewConfigPeerAdminService(listenAddress)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigPeerAuthentication successfully`, func() {
				timewindow := "15m"
				model, err := testService.NewConfigPeerAuthentication(timewindow)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewConfigPeerClient successfully`, func() {
				connTimeout := "2s"
				model, err := testService.NewConfigPeerClient(connTimeout)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateCaBodyConfigOverride successfully`, func() {
				var ca *blockchainv2.ConfigCACreate = nil
				_, err := testService.NewCreateCaBodyConfigOverride(ca)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCreateCaBodyResources successfully`, func() {
				var ca *blockchainv2.ResourceObject = nil
				_, err := testService.NewCreateCaBodyResources(ca)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCreateCaBodyStorage successfully`, func() {
				var ca *blockchainv2.StorageObject = nil
				_, err := testService.NewCreateCaBodyStorage(ca)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCreateCaOptions successfully`, func() {
				// Construct an instance of the ConfigCADbTlsClient model
				configCaDbTlsClientModel := new(blockchainv2.ConfigCADbTlsClient)
				Expect(configCaDbTlsClientModel).ToNot(BeNil())
				configCaDbTlsClientModel.Certfile = core.StringPtr("testString")
				configCaDbTlsClientModel.Keyfile = core.StringPtr("testString")
				Expect(configCaDbTlsClientModel.Certfile).To(Equal(core.StringPtr("testString")))
				Expect(configCaDbTlsClientModel.Keyfile).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigCAIntermediateTlsClient model
				configCaIntermediateTlsClientModel := new(blockchainv2.ConfigCAIntermediateTlsClient)
				Expect(configCaIntermediateTlsClientModel).ToNot(BeNil())
				configCaIntermediateTlsClientModel.Certfile = core.StringPtr("testString")
				configCaIntermediateTlsClientModel.Keyfile = core.StringPtr("testString")
				Expect(configCaIntermediateTlsClientModel.Certfile).To(Equal(core.StringPtr("testString")))
				Expect(configCaIntermediateTlsClientModel.Keyfile).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the IdentityAttrs model
				identityAttrsModel := new(blockchainv2.IdentityAttrs)
				Expect(identityAttrsModel).ToNot(BeNil())
				identityAttrsModel.HfRegistrarRoles = core.StringPtr("*")
				identityAttrsModel.HfRegistrarDelegateRoles = core.StringPtr("*")
				identityAttrsModel.HfRevoker = core.BoolPtr(true)
				identityAttrsModel.HfIntermediateCA = core.BoolPtr(true)
				identityAttrsModel.HfGenCRL = core.BoolPtr(true)
				identityAttrsModel.HfRegistrarAttributes = core.StringPtr("*")
				identityAttrsModel.HfAffiliationMgr = core.BoolPtr(true)
				Expect(identityAttrsModel.HfRegistrarRoles).To(Equal(core.StringPtr("*")))
				Expect(identityAttrsModel.HfRegistrarDelegateRoles).To(Equal(core.StringPtr("*")))
				Expect(identityAttrsModel.HfRevoker).To(Equal(core.BoolPtr(true)))
				Expect(identityAttrsModel.HfIntermediateCA).To(Equal(core.BoolPtr(true)))
				Expect(identityAttrsModel.HfGenCRL).To(Equal(core.BoolPtr(true)))
				Expect(identityAttrsModel.HfRegistrarAttributes).To(Equal(core.StringPtr("*")))
				Expect(identityAttrsModel.HfAffiliationMgr).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv2.BccspPKCS11)
				Expect(bccspPkcS11Model).ToNot(BeNil())
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))
				Expect(bccspPkcS11Model.Label).To(Equal(core.StringPtr("testString")))
				Expect(bccspPkcS11Model.Pin).To(Equal(core.StringPtr("testString")))
				Expect(bccspPkcS11Model.Hash).To(Equal(core.StringPtr("SHA2")))
				Expect(bccspPkcS11Model.Security).To(Equal(core.Float64Ptr(float64(256))))

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv2.BccspSW)
				Expect(bccspSwModel).ToNot(BeNil())
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))
				Expect(bccspSwModel.Hash).To(Equal(core.StringPtr("SHA2")))
				Expect(bccspSwModel.Security).To(Equal(core.Float64Ptr(float64(256))))

				// Construct an instance of the ConfigCACfgIdentities model
				configCaCfgIdentitiesModel := new(blockchainv2.ConfigCACfgIdentities)
				Expect(configCaCfgIdentitiesModel).ToNot(BeNil())
				configCaCfgIdentitiesModel.Passwordattempts = core.Float64Ptr(float64(10))
				configCaCfgIdentitiesModel.Allowremove = core.BoolPtr(false)
				Expect(configCaCfgIdentitiesModel.Passwordattempts).To(Equal(core.Float64Ptr(float64(10))))
				Expect(configCaCfgIdentitiesModel.Allowremove).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the ConfigCACsrCa model
				configCaCsrCaModel := new(blockchainv2.ConfigCACsrCa)
				Expect(configCaCsrCaModel).ToNot(BeNil())
				configCaCsrCaModel.Expiry = core.StringPtr("131400h")
				configCaCsrCaModel.Pathlength = core.Float64Ptr(float64(0))
				Expect(configCaCsrCaModel.Expiry).To(Equal(core.StringPtr("131400h")))
				Expect(configCaCsrCaModel.Pathlength).To(Equal(core.Float64Ptr(float64(0))))

				// Construct an instance of the ConfigCACsrKeyrequest model
				configCaCsrKeyrequestModel := new(blockchainv2.ConfigCACsrKeyrequest)
				Expect(configCaCsrKeyrequestModel).ToNot(BeNil())
				configCaCsrKeyrequestModel.Algo = core.StringPtr("ecdsa")
				configCaCsrKeyrequestModel.Size = core.Float64Ptr(float64(256))
				Expect(configCaCsrKeyrequestModel.Algo).To(Equal(core.StringPtr("ecdsa")))
				Expect(configCaCsrKeyrequestModel.Size).To(Equal(core.Float64Ptr(float64(256))))

				// Construct an instance of the ConfigCACsrNamesItem model
				configCaCsrNamesItemModel := new(blockchainv2.ConfigCACsrNamesItem)
				Expect(configCaCsrNamesItemModel).ToNot(BeNil())
				configCaCsrNamesItemModel.C = core.StringPtr("US")
				configCaCsrNamesItemModel.ST = core.StringPtr("North Carolina")
				configCaCsrNamesItemModel.L = core.StringPtr("Raleigh")
				configCaCsrNamesItemModel.O = core.StringPtr("Hyperledger")
				configCaCsrNamesItemModel.OU = core.StringPtr("Fabric")
				Expect(configCaCsrNamesItemModel.C).To(Equal(core.StringPtr("US")))
				Expect(configCaCsrNamesItemModel.ST).To(Equal(core.StringPtr("North Carolina")))
				Expect(configCaCsrNamesItemModel.L).To(Equal(core.StringPtr("Raleigh")))
				Expect(configCaCsrNamesItemModel.O).To(Equal(core.StringPtr("Hyperledger")))
				Expect(configCaCsrNamesItemModel.OU).To(Equal(core.StringPtr("Fabric")))

				// Construct an instance of the ConfigCADbTls model
				configCaDbTlsModel := new(blockchainv2.ConfigCADbTls)
				Expect(configCaDbTlsModel).ToNot(BeNil())
				configCaDbTlsModel.Certfiles = []string{"testString"}
				configCaDbTlsModel.Client = configCaDbTlsClientModel
				configCaDbTlsModel.Enabled = core.BoolPtr(false)
				Expect(configCaDbTlsModel.Certfiles).To(Equal([]string{"testString"}))
				Expect(configCaDbTlsModel.Client).To(Equal(configCaDbTlsClientModel))
				Expect(configCaDbTlsModel.Enabled).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the ConfigCAIntermediateEnrollment model
				configCaIntermediateEnrollmentModel := new(blockchainv2.ConfigCAIntermediateEnrollment)
				Expect(configCaIntermediateEnrollmentModel).ToNot(BeNil())
				configCaIntermediateEnrollmentModel.Hosts = core.StringPtr("localhost")
				configCaIntermediateEnrollmentModel.Profile = core.StringPtr("testString")
				configCaIntermediateEnrollmentModel.Label = core.StringPtr("testString")
				Expect(configCaIntermediateEnrollmentModel.Hosts).To(Equal(core.StringPtr("localhost")))
				Expect(configCaIntermediateEnrollmentModel.Profile).To(Equal(core.StringPtr("testString")))
				Expect(configCaIntermediateEnrollmentModel.Label).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigCAIntermediateParentserver model
				configCaIntermediateParentserverModel := new(blockchainv2.ConfigCAIntermediateParentserver)
				Expect(configCaIntermediateParentserverModel).ToNot(BeNil())
				configCaIntermediateParentserverModel.URL = core.StringPtr("testString")
				configCaIntermediateParentserverModel.Caname = core.StringPtr("testString")
				Expect(configCaIntermediateParentserverModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(configCaIntermediateParentserverModel.Caname).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigCAIntermediateTls model
				configCaIntermediateTlsModel := new(blockchainv2.ConfigCAIntermediateTls)
				Expect(configCaIntermediateTlsModel).ToNot(BeNil())
				configCaIntermediateTlsModel.Certfiles = []string{"testString"}
				configCaIntermediateTlsModel.Client = configCaIntermediateTlsClientModel
				Expect(configCaIntermediateTlsModel.Certfiles).To(Equal([]string{"testString"}))
				Expect(configCaIntermediateTlsModel.Client).To(Equal(configCaIntermediateTlsClientModel))

				// Construct an instance of the ConfigCARegistryIdentitiesItem model
				configCaRegistryIdentitiesItemModel := new(blockchainv2.ConfigCARegistryIdentitiesItem)
				Expect(configCaRegistryIdentitiesItemModel).ToNot(BeNil())
				configCaRegistryIdentitiesItemModel.Name = core.StringPtr("admin")
				configCaRegistryIdentitiesItemModel.Pass = core.StringPtr("password")
				configCaRegistryIdentitiesItemModel.Type = core.StringPtr("client")
				configCaRegistryIdentitiesItemModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryIdentitiesItemModel.Affiliation = core.StringPtr("testString")
				configCaRegistryIdentitiesItemModel.Attrs = identityAttrsModel
				Expect(configCaRegistryIdentitiesItemModel.Name).To(Equal(core.StringPtr("admin")))
				Expect(configCaRegistryIdentitiesItemModel.Pass).To(Equal(core.StringPtr("password")))
				Expect(configCaRegistryIdentitiesItemModel.Type).To(Equal(core.StringPtr("client")))
				Expect(configCaRegistryIdentitiesItemModel.Maxenrollments).To(Equal(core.Float64Ptr(float64(-1))))
				Expect(configCaRegistryIdentitiesItemModel.Affiliation).To(Equal(core.StringPtr("testString")))
				Expect(configCaRegistryIdentitiesItemModel.Attrs).To(Equal(identityAttrsModel))

				// Construct an instance of the ConfigCATlsClientauth model
				configCaTlsClientauthModel := new(blockchainv2.ConfigCATlsClientauth)
				Expect(configCaTlsClientauthModel).ToNot(BeNil())
				configCaTlsClientauthModel.Type = core.StringPtr("noclientcert")
				configCaTlsClientauthModel.Certfiles = []string{"testString"}
				Expect(configCaTlsClientauthModel.Type).To(Equal(core.StringPtr("noclientcert")))
				Expect(configCaTlsClientauthModel.Certfiles).To(Equal([]string{"testString"}))

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv2.MetricsStatsd)
				Expect(metricsStatsdModel).ToNot(BeNil())
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")
				Expect(metricsStatsdModel.Network).To(Equal(core.StringPtr("udp")))
				Expect(metricsStatsdModel.Address).To(Equal(core.StringPtr("127.0.0.1:8125")))
				Expect(metricsStatsdModel.WriteInterval).To(Equal(core.StringPtr("10s")))
				Expect(metricsStatsdModel.Prefix).To(Equal(core.StringPtr("server")))

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv2.Bccsp)
				Expect(bccspModel).ToNot(BeNil())
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model
				Expect(bccspModel.Default).To(Equal(core.StringPtr("SW")))
				Expect(bccspModel.SW).To(Equal(bccspSwModel))
				Expect(bccspModel.PKCS11).To(Equal(bccspPkcS11Model))

				// Construct an instance of the ConfigCAAffiliations model
				configCaAffiliationsModel := new(blockchainv2.ConfigCAAffiliations)
				Expect(configCaAffiliationsModel).ToNot(BeNil())
				configCaAffiliationsModel.Org1 = []string{"department1"}
				configCaAffiliationsModel.Org2 = []string{"department1"}
				configCaAffiliationsModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(configCaAffiliationsModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))
				Expect(configCaAffiliationsModel.GetProperties()).ToNot(BeEmpty())
				Expect(configCaAffiliationsModel.Org1).To(Equal([]string{"department1"}))
				Expect(configCaAffiliationsModel.Org2).To(Equal([]string{"department1"}))

				// Construct an instance of the ConfigCACa model
				configCaCaModel := new(blockchainv2.ConfigCACa)
				Expect(configCaCaModel).ToNot(BeNil())
				configCaCaModel.Keyfile = core.StringPtr("testString")
				configCaCaModel.Certfile = core.StringPtr("testString")
				configCaCaModel.Chainfile = core.StringPtr("testString")
				Expect(configCaCaModel.Keyfile).To(Equal(core.StringPtr("testString")))
				Expect(configCaCaModel.Certfile).To(Equal(core.StringPtr("testString")))
				Expect(configCaCaModel.Chainfile).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigCACfg model
				configCaCfgModel := new(blockchainv2.ConfigCACfg)
				Expect(configCaCfgModel).ToNot(BeNil())
				configCaCfgModel.Identities = configCaCfgIdentitiesModel
				Expect(configCaCfgModel.Identities).To(Equal(configCaCfgIdentitiesModel))

				// Construct an instance of the ConfigCACors model
				configCaCorsModel := new(blockchainv2.ConfigCACors)
				Expect(configCaCorsModel).ToNot(BeNil())
				configCaCorsModel.Enabled = core.BoolPtr(true)
				configCaCorsModel.Origins = []string{"*"}
				Expect(configCaCorsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(configCaCorsModel.Origins).To(Equal([]string{"*"}))

				// Construct an instance of the ConfigCACrl model
				configCaCrlModel := new(blockchainv2.ConfigCACrl)
				Expect(configCaCrlModel).ToNot(BeNil())
				configCaCrlModel.Expiry = core.StringPtr("24h")
				Expect(configCaCrlModel.Expiry).To(Equal(core.StringPtr("24h")))

				// Construct an instance of the ConfigCACsr model
				configCaCsrModel := new(blockchainv2.ConfigCACsr)
				Expect(configCaCsrModel).ToNot(BeNil())
				configCaCsrModel.Cn = core.StringPtr("ca")
				configCaCsrModel.Keyrequest = configCaCsrKeyrequestModel
				configCaCsrModel.Names = []blockchainv2.ConfigCACsrNamesItem{*configCaCsrNamesItemModel}
				configCaCsrModel.Hosts = []string{"localhost"}
				configCaCsrModel.Ca = configCaCsrCaModel
				Expect(configCaCsrModel.Cn).To(Equal(core.StringPtr("ca")))
				Expect(configCaCsrModel.Keyrequest).To(Equal(configCaCsrKeyrequestModel))
				Expect(configCaCsrModel.Names).To(Equal([]blockchainv2.ConfigCACsrNamesItem{*configCaCsrNamesItemModel}))
				Expect(configCaCsrModel.Hosts).To(Equal([]string{"localhost"}))
				Expect(configCaCsrModel.Ca).To(Equal(configCaCsrCaModel))

				// Construct an instance of the ConfigCADb model
				configCaDbModel := new(blockchainv2.ConfigCADb)
				Expect(configCaDbModel).ToNot(BeNil())
				configCaDbModel.Type = core.StringPtr("postgres")
				configCaDbModel.Datasource = core.StringPtr("host=fake.databases.appdomain.cloud port=31941 user=ibm_cloud password=password dbname=ibmclouddb sslmode=verify-full")
				configCaDbModel.Tls = configCaDbTlsModel
				Expect(configCaDbModel.Type).To(Equal(core.StringPtr("postgres")))
				Expect(configCaDbModel.Datasource).To(Equal(core.StringPtr("host=fake.databases.appdomain.cloud port=31941 user=ibm_cloud password=password dbname=ibmclouddb sslmode=verify-full")))
				Expect(configCaDbModel.Tls).To(Equal(configCaDbTlsModel))

				// Construct an instance of the ConfigCAIdemix model
				configCaIdemixModel := new(blockchainv2.ConfigCAIdemix)
				Expect(configCaIdemixModel).ToNot(BeNil())
				configCaIdemixModel.Rhpoolsize = core.Float64Ptr(float64(100))
				configCaIdemixModel.Nonceexpiration = core.StringPtr("15s")
				configCaIdemixModel.Noncesweepinterval = core.StringPtr("15m")
				Expect(configCaIdemixModel.Rhpoolsize).To(Equal(core.Float64Ptr(float64(100))))
				Expect(configCaIdemixModel.Nonceexpiration).To(Equal(core.StringPtr("15s")))
				Expect(configCaIdemixModel.Noncesweepinterval).To(Equal(core.StringPtr("15m")))

				// Construct an instance of the ConfigCAIntermediate model
				configCaIntermediateModel := new(blockchainv2.ConfigCAIntermediate)
				Expect(configCaIntermediateModel).ToNot(BeNil())
				configCaIntermediateModel.Parentserver = configCaIntermediateParentserverModel
				configCaIntermediateModel.Enrollment = configCaIntermediateEnrollmentModel
				configCaIntermediateModel.Tls = configCaIntermediateTlsModel
				Expect(configCaIntermediateModel.Parentserver).To(Equal(configCaIntermediateParentserverModel))
				Expect(configCaIntermediateModel.Enrollment).To(Equal(configCaIntermediateEnrollmentModel))
				Expect(configCaIntermediateModel.Tls).To(Equal(configCaIntermediateTlsModel))

				// Construct an instance of the ConfigCARegistry model
				configCaRegistryModel := new(blockchainv2.ConfigCARegistry)
				Expect(configCaRegistryModel).ToNot(BeNil())
				configCaRegistryModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryModel.Identities = []blockchainv2.ConfigCARegistryIdentitiesItem{*configCaRegistryIdentitiesItemModel}
				Expect(configCaRegistryModel.Maxenrollments).To(Equal(core.Float64Ptr(float64(-1))))
				Expect(configCaRegistryModel.Identities).To(Equal([]blockchainv2.ConfigCARegistryIdentitiesItem{*configCaRegistryIdentitiesItemModel}))

				// Construct an instance of the ConfigCATls model
				configCaTlsModel := new(blockchainv2.ConfigCATls)
				Expect(configCaTlsModel).ToNot(BeNil())
				configCaTlsModel.Keyfile = core.StringPtr("testString")
				configCaTlsModel.Certfile = core.StringPtr("testString")
				configCaTlsModel.Clientauth = configCaTlsClientauthModel
				Expect(configCaTlsModel.Keyfile).To(Equal(core.StringPtr("testString")))
				Expect(configCaTlsModel.Certfile).To(Equal(core.StringPtr("testString")))
				Expect(configCaTlsModel.Clientauth).To(Equal(configCaTlsClientauthModel))

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv2.Metrics)
				Expect(metricsModel).ToNot(BeNil())
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel
				Expect(metricsModel.Provider).To(Equal(core.StringPtr("prometheus")))
				Expect(metricsModel.Statsd).To(Equal(metricsStatsdModel))

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv2.ResourceLimits)
				Expect(resourceLimitsModel).ToNot(BeNil())
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")
				Expect(resourceLimitsModel.Cpu).To(Equal(core.StringPtr("100m")))
				Expect(resourceLimitsModel.Memory).To(Equal(core.StringPtr("256MiB")))

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv2.ResourceRequests)
				Expect(resourceRequestsModel).ToNot(BeNil())
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")
				Expect(resourceRequestsModel.Cpu).To(Equal(core.StringPtr("100m")))
				Expect(resourceRequestsModel.Memory).To(Equal(core.StringPtr("256MiB")))

				// Construct an instance of the ConfigCACreate model
				configCaCreateModel := new(blockchainv2.ConfigCACreate)
				Expect(configCaCreateModel).ToNot(BeNil())
				configCaCreateModel.Cors = configCaCorsModel
				configCaCreateModel.Debug = core.BoolPtr(false)
				configCaCreateModel.Crlsizelimit = core.Float64Ptr(float64(512000))
				configCaCreateModel.Tls = configCaTlsModel
				configCaCreateModel.Ca = configCaCaModel
				configCaCreateModel.Crl = configCaCrlModel
				configCaCreateModel.Registry = configCaRegistryModel
				configCaCreateModel.Db = configCaDbModel
				configCaCreateModel.Affiliations = configCaAffiliationsModel
				configCaCreateModel.Csr = configCaCsrModel
				configCaCreateModel.Idemix = configCaIdemixModel
				configCaCreateModel.BCCSP = bccspModel
				configCaCreateModel.Intermediate = configCaIntermediateModel
				configCaCreateModel.Cfg = configCaCfgModel
				configCaCreateModel.Metrics = metricsModel
				Expect(configCaCreateModel.Cors).To(Equal(configCaCorsModel))
				Expect(configCaCreateModel.Debug).To(Equal(core.BoolPtr(false)))
				Expect(configCaCreateModel.Crlsizelimit).To(Equal(core.Float64Ptr(float64(512000))))
				Expect(configCaCreateModel.Tls).To(Equal(configCaTlsModel))
				Expect(configCaCreateModel.Ca).To(Equal(configCaCaModel))
				Expect(configCaCreateModel.Crl).To(Equal(configCaCrlModel))
				Expect(configCaCreateModel.Registry).To(Equal(configCaRegistryModel))
				Expect(configCaCreateModel.Db).To(Equal(configCaDbModel))
				Expect(configCaCreateModel.Affiliations).To(Equal(configCaAffiliationsModel))
				Expect(configCaCreateModel.Csr).To(Equal(configCaCsrModel))
				Expect(configCaCreateModel.Idemix).To(Equal(configCaIdemixModel))
				Expect(configCaCreateModel.BCCSP).To(Equal(bccspModel))
				Expect(configCaCreateModel.Intermediate).To(Equal(configCaIntermediateModel))
				Expect(configCaCreateModel.Cfg).To(Equal(configCaCfgModel))
				Expect(configCaCreateModel.Metrics).To(Equal(metricsModel))

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv2.ResourceObject)
				Expect(resourceObjectModel).ToNot(BeNil())
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel
				Expect(resourceObjectModel.Requests).To(Equal(resourceRequestsModel))
				Expect(resourceObjectModel.Limits).To(Equal(resourceLimitsModel))

				// Construct an instance of the StorageObject model
				storageObjectModel := new(blockchainv2.StorageObject)
				Expect(storageObjectModel).ToNot(BeNil())
				storageObjectModel.Size = core.StringPtr("4GiB")
				storageObjectModel.Class = core.StringPtr("default")
				Expect(storageObjectModel.Size).To(Equal(core.StringPtr("4GiB")))
				Expect(storageObjectModel.Class).To(Equal(core.StringPtr("default")))

				// Construct an instance of the CreateCaBodyConfigOverride model
				createCaBodyConfigOverrideModel := new(blockchainv2.CreateCaBodyConfigOverride)
				Expect(createCaBodyConfigOverrideModel).ToNot(BeNil())
				createCaBodyConfigOverrideModel.Ca = configCaCreateModel
				createCaBodyConfigOverrideModel.Tlsca = configCaCreateModel
				Expect(createCaBodyConfigOverrideModel.Ca).To(Equal(configCaCreateModel))
				Expect(createCaBodyConfigOverrideModel.Tlsca).To(Equal(configCaCreateModel))

				// Construct an instance of the CreateCaBodyResources model
				createCaBodyResourcesModel := new(blockchainv2.CreateCaBodyResources)
				Expect(createCaBodyResourcesModel).ToNot(BeNil())
				createCaBodyResourcesModel.Ca = resourceObjectModel
				Expect(createCaBodyResourcesModel.Ca).To(Equal(resourceObjectModel))

				// Construct an instance of the CreateCaBodyStorage model
				createCaBodyStorageModel := new(blockchainv2.CreateCaBodyStorage)
				Expect(createCaBodyStorageModel).ToNot(BeNil())
				createCaBodyStorageModel.Ca = storageObjectModel
				Expect(createCaBodyStorageModel.Ca).To(Equal(storageObjectModel))

				// Construct an instance of the Hsm model
				hsmModel := new(blockchainv2.Hsm)
				Expect(hsmModel).ToNot(BeNil())
				hsmModel.Pkcs11endpoint = core.StringPtr("tcp://example.com:666")
				Expect(hsmModel.Pkcs11endpoint).To(Equal(core.StringPtr("tcp://example.com:666")))

				// Construct an instance of the CreateCaOptions model
				createCaOptionsDisplayName := "My CA"
				var createCaOptionsConfigOverride *blockchainv2.CreateCaBodyConfigOverride = nil
				createCaOptionsModel := testService.NewCreateCaOptions(createCaOptionsDisplayName, createCaOptionsConfigOverride)
				createCaOptionsModel.SetDisplayName("My CA")
				createCaOptionsModel.SetConfigOverride(createCaBodyConfigOverrideModel)
				createCaOptionsModel.SetResources(createCaBodyResourcesModel)
				createCaOptionsModel.SetStorage(createCaBodyStorageModel)
				createCaOptionsModel.SetZone("testString")
				createCaOptionsModel.SetReplicas(float64(1))
				createCaOptionsModel.SetTags([]string{"testString"})
				createCaOptionsModel.SetHsm(hsmModel)
				createCaOptionsModel.SetRegion("testString")
				createCaOptionsModel.SetVersion("1.4.6-1")
				createCaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCaOptionsModel).ToNot(BeNil())
				Expect(createCaOptionsModel.DisplayName).To(Equal(core.StringPtr("My CA")))
				Expect(createCaOptionsModel.ConfigOverride).To(Equal(createCaBodyConfigOverrideModel))
				Expect(createCaOptionsModel.Resources).To(Equal(createCaBodyResourcesModel))
				Expect(createCaOptionsModel.Storage).To(Equal(createCaBodyStorageModel))
				Expect(createCaOptionsModel.Zone).To(Equal(core.StringPtr("testString")))
				Expect(createCaOptionsModel.Replicas).To(Equal(core.Float64Ptr(float64(1))))
				Expect(createCaOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(createCaOptionsModel.Hsm).To(Equal(hsmModel))
				Expect(createCaOptionsModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(createCaOptionsModel.Version).To(Equal(core.StringPtr("1.4.6-1")))
				Expect(createCaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateOrdererOptions successfully`, func() {
				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv2.BccspPKCS11)
				Expect(bccspPkcS11Model).ToNot(BeNil())
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))
				Expect(bccspPkcS11Model.Label).To(Equal(core.StringPtr("testString")))
				Expect(bccspPkcS11Model.Pin).To(Equal(core.StringPtr("testString")))
				Expect(bccspPkcS11Model.Hash).To(Equal(core.StringPtr("SHA2")))
				Expect(bccspPkcS11Model.Security).To(Equal(core.Float64Ptr(float64(256))))

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv2.BccspSW)
				Expect(bccspSwModel).ToNot(BeNil())
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))
				Expect(bccspSwModel.Hash).To(Equal(core.StringPtr("SHA2")))
				Expect(bccspSwModel.Security).To(Equal(core.Float64Ptr(float64(256))))

				// Construct an instance of the ConfigObjectEnrollmentComponentCatls model
				configObjectEnrollmentComponentCatlsModel := new(blockchainv2.ConfigObjectEnrollmentComponentCatls)
				Expect(configObjectEnrollmentComponentCatlsModel).ToNot(BeNil())
				configObjectEnrollmentComponentCatlsModel.Cacert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				Expect(configObjectEnrollmentComponentCatlsModel.Cacert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))

				// Construct an instance of the ConfigObjectEnrollmentTlsCatls model
				configObjectEnrollmentTlsCatlsModel := new(blockchainv2.ConfigObjectEnrollmentTlsCatls)
				Expect(configObjectEnrollmentTlsCatlsModel).ToNot(BeNil())
				configObjectEnrollmentTlsCatlsModel.Cacert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				Expect(configObjectEnrollmentTlsCatlsModel.Cacert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))

				// Construct an instance of the ConfigObjectEnrollmentTlsCsr model
				configObjectEnrollmentTlsCsrModel := new(blockchainv2.ConfigObjectEnrollmentTlsCsr)
				Expect(configObjectEnrollmentTlsCsrModel).ToNot(BeNil())
				configObjectEnrollmentTlsCsrModel.Hosts = []string{"testString"}
				Expect(configObjectEnrollmentTlsCsrModel.Hosts).To(Equal([]string{"testString"}))

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv2.Bccsp)
				Expect(bccspModel).ToNot(BeNil())
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model
				Expect(bccspModel.Default).To(Equal(core.StringPtr("SW")))
				Expect(bccspModel.SW).To(Equal(bccspSwModel))
				Expect(bccspModel.PKCS11).To(Equal(bccspPkcS11Model))

				// Construct an instance of the ConfigObjectEnrollmentComponent model
				configObjectEnrollmentComponentModel := new(blockchainv2.ConfigObjectEnrollmentComponent)
				Expect(configObjectEnrollmentComponentModel).ToNot(BeNil())
				configObjectEnrollmentComponentModel.Cahost = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				configObjectEnrollmentComponentModel.Caport = core.Float64Ptr(float64(7054))
				configObjectEnrollmentComponentModel.Caname = core.StringPtr("ca")
				configObjectEnrollmentComponentModel.Catls = configObjectEnrollmentComponentCatlsModel
				configObjectEnrollmentComponentModel.Enrollid = core.StringPtr("admin")
				configObjectEnrollmentComponentModel.Enrollsecret = core.StringPtr("password")
				configObjectEnrollmentComponentModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				Expect(configObjectEnrollmentComponentModel.Cahost).To(Equal(core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")))
				Expect(configObjectEnrollmentComponentModel.Caport).To(Equal(core.Float64Ptr(float64(7054))))
				Expect(configObjectEnrollmentComponentModel.Caname).To(Equal(core.StringPtr("ca")))
				Expect(configObjectEnrollmentComponentModel.Catls).To(Equal(configObjectEnrollmentComponentCatlsModel))
				Expect(configObjectEnrollmentComponentModel.Enrollid).To(Equal(core.StringPtr("admin")))
				Expect(configObjectEnrollmentComponentModel.Enrollsecret).To(Equal(core.StringPtr("password")))
				Expect(configObjectEnrollmentComponentModel.Admincerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))

				// Construct an instance of the ConfigObjectEnrollmentTls model
				configObjectEnrollmentTlsModel := new(blockchainv2.ConfigObjectEnrollmentTls)
				Expect(configObjectEnrollmentTlsModel).ToNot(BeNil())
				configObjectEnrollmentTlsModel.Cahost = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				configObjectEnrollmentTlsModel.Caport = core.Float64Ptr(float64(7054))
				configObjectEnrollmentTlsModel.Caname = core.StringPtr("tlsca")
				configObjectEnrollmentTlsModel.Catls = configObjectEnrollmentTlsCatlsModel
				configObjectEnrollmentTlsModel.Enrollid = core.StringPtr("admin")
				configObjectEnrollmentTlsModel.Enrollsecret = core.StringPtr("password")
				configObjectEnrollmentTlsModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				configObjectEnrollmentTlsModel.Csr = configObjectEnrollmentTlsCsrModel
				Expect(configObjectEnrollmentTlsModel.Cahost).To(Equal(core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")))
				Expect(configObjectEnrollmentTlsModel.Caport).To(Equal(core.Float64Ptr(float64(7054))))
				Expect(configObjectEnrollmentTlsModel.Caname).To(Equal(core.StringPtr("tlsca")))
				Expect(configObjectEnrollmentTlsModel.Catls).To(Equal(configObjectEnrollmentTlsCatlsModel))
				Expect(configObjectEnrollmentTlsModel.Enrollid).To(Equal(core.StringPtr("admin")))
				Expect(configObjectEnrollmentTlsModel.Enrollsecret).To(Equal(core.StringPtr("password")))
				Expect(configObjectEnrollmentTlsModel.Admincerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(configObjectEnrollmentTlsModel.Csr).To(Equal(configObjectEnrollmentTlsCsrModel))

				// Construct an instance of the ConfigOrdererAuthentication model
				configOrdererAuthenticationModel := new(blockchainv2.ConfigOrdererAuthentication)
				Expect(configOrdererAuthenticationModel).ToNot(BeNil())
				configOrdererAuthenticationModel.TimeWindow = core.StringPtr("15m")
				configOrdererAuthenticationModel.NoExpirationChecks = core.BoolPtr(false)
				Expect(configOrdererAuthenticationModel.TimeWindow).To(Equal(core.StringPtr("15m")))
				Expect(configOrdererAuthenticationModel.NoExpirationChecks).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the ConfigOrdererKeepalive model
				configOrdererKeepaliveModel := new(blockchainv2.ConfigOrdererKeepalive)
				Expect(configOrdererKeepaliveModel).ToNot(BeNil())
				configOrdererKeepaliveModel.ServerMinInterval = core.StringPtr("60s")
				configOrdererKeepaliveModel.ServerInterval = core.StringPtr("2h")
				configOrdererKeepaliveModel.ServerTimeout = core.StringPtr("20s")
				Expect(configOrdererKeepaliveModel.ServerMinInterval).To(Equal(core.StringPtr("60s")))
				Expect(configOrdererKeepaliveModel.ServerInterval).To(Equal(core.StringPtr("2h")))
				Expect(configOrdererKeepaliveModel.ServerTimeout).To(Equal(core.StringPtr("20s")))

				// Construct an instance of the ConfigOrdererMetricsStatsd model
				configOrdererMetricsStatsdModel := new(blockchainv2.ConfigOrdererMetricsStatsd)
				Expect(configOrdererMetricsStatsdModel).ToNot(BeNil())
				configOrdererMetricsStatsdModel.Network = core.StringPtr("udp")
				configOrdererMetricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				configOrdererMetricsStatsdModel.WriteInterval = core.StringPtr("10s")
				configOrdererMetricsStatsdModel.Prefix = core.StringPtr("server")
				Expect(configOrdererMetricsStatsdModel.Network).To(Equal(core.StringPtr("udp")))
				Expect(configOrdererMetricsStatsdModel.Address).To(Equal(core.StringPtr("127.0.0.1:8125")))
				Expect(configOrdererMetricsStatsdModel.WriteInterval).To(Equal(core.StringPtr("10s")))
				Expect(configOrdererMetricsStatsdModel.Prefix).To(Equal(core.StringPtr("server")))

				// Construct an instance of the MspConfigData model
				mspConfigDataModel := new(blockchainv2.MspConfigData)
				Expect(mspConfigDataModel).ToNot(BeNil())
				mspConfigDataModel.Keystore = core.StringPtr("testString")
				mspConfigDataModel.Signcerts = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspConfigDataModel.Cacerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				mspConfigDataModel.Intermediatecerts = []string{"testString"}
				mspConfigDataModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				Expect(mspConfigDataModel.Keystore).To(Equal(core.StringPtr("testString")))
				Expect(mspConfigDataModel.Signcerts).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(mspConfigDataModel.Cacerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(mspConfigDataModel.Intermediatecerts).To(Equal([]string{"testString"}))
				Expect(mspConfigDataModel.Admincerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv2.ResourceLimits)
				Expect(resourceLimitsModel).ToNot(BeNil())
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")
				Expect(resourceLimitsModel.Cpu).To(Equal(core.StringPtr("100m")))
				Expect(resourceLimitsModel.Memory).To(Equal(core.StringPtr("256MiB")))

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv2.ResourceRequests)
				Expect(resourceRequestsModel).ToNot(BeNil())
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")
				Expect(resourceRequestsModel.Cpu).To(Equal(core.StringPtr("100m")))
				Expect(resourceRequestsModel.Memory).To(Equal(core.StringPtr("256MiB")))

				// Construct an instance of the ConfigObjectEnrollment model
				configObjectEnrollmentModel := new(blockchainv2.ConfigObjectEnrollment)
				Expect(configObjectEnrollmentModel).ToNot(BeNil())
				configObjectEnrollmentModel.Component = configObjectEnrollmentComponentModel
				configObjectEnrollmentModel.Tls = configObjectEnrollmentTlsModel
				Expect(configObjectEnrollmentModel.Component).To(Equal(configObjectEnrollmentComponentModel))
				Expect(configObjectEnrollmentModel.Tls).To(Equal(configObjectEnrollmentTlsModel))

				// Construct an instance of the ConfigObjectMsp model
				configObjectMspModel := new(blockchainv2.ConfigObjectMsp)
				Expect(configObjectMspModel).ToNot(BeNil())
				configObjectMspModel.Component = mspConfigDataModel
				configObjectMspModel.Tls = mspConfigDataModel
				configObjectMspModel.Clientauth = mspConfigDataModel
				Expect(configObjectMspModel.Component).To(Equal(mspConfigDataModel))
				Expect(configObjectMspModel.Tls).To(Equal(mspConfigDataModel))
				Expect(configObjectMspModel.Clientauth).To(Equal(mspConfigDataModel))

				// Construct an instance of the ConfigOrdererDebug model
				configOrdererDebugModel := new(blockchainv2.ConfigOrdererDebug)
				Expect(configOrdererDebugModel).ToNot(BeNil())
				configOrdererDebugModel.BroadcastTraceDir = core.StringPtr("testString")
				configOrdererDebugModel.DeliverTraceDir = core.StringPtr("testString")
				Expect(configOrdererDebugModel.BroadcastTraceDir).To(Equal(core.StringPtr("testString")))
				Expect(configOrdererDebugModel.DeliverTraceDir).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigOrdererGeneral model
				configOrdererGeneralModel := new(blockchainv2.ConfigOrdererGeneral)
				Expect(configOrdererGeneralModel).ToNot(BeNil())
				configOrdererGeneralModel.Keepalive = configOrdererKeepaliveModel
				configOrdererGeneralModel.BCCSP = bccspModel
				configOrdererGeneralModel.Authentication = configOrdererAuthenticationModel
				Expect(configOrdererGeneralModel.Keepalive).To(Equal(configOrdererKeepaliveModel))
				Expect(configOrdererGeneralModel.BCCSP).To(Equal(bccspModel))
				Expect(configOrdererGeneralModel.Authentication).To(Equal(configOrdererAuthenticationModel))

				// Construct an instance of the ConfigOrdererMetrics model
				configOrdererMetricsModel := new(blockchainv2.ConfigOrdererMetrics)
				Expect(configOrdererMetricsModel).ToNot(BeNil())
				configOrdererMetricsModel.Provider = core.StringPtr("disabled")
				configOrdererMetricsModel.Statsd = configOrdererMetricsStatsdModel
				Expect(configOrdererMetricsModel.Provider).To(Equal(core.StringPtr("disabled")))
				Expect(configOrdererMetricsModel.Statsd).To(Equal(configOrdererMetricsStatsdModel))

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv2.ResourceObject)
				Expect(resourceObjectModel).ToNot(BeNil())
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel
				Expect(resourceObjectModel.Requests).To(Equal(resourceRequestsModel))
				Expect(resourceObjectModel.Limits).To(Equal(resourceLimitsModel))

				// Construct an instance of the StorageObject model
				storageObjectModel := new(blockchainv2.StorageObject)
				Expect(storageObjectModel).ToNot(BeNil())
				storageObjectModel.Size = core.StringPtr("4GiB")
				storageObjectModel.Class = core.StringPtr("default")
				Expect(storageObjectModel.Size).To(Equal(core.StringPtr("4GiB")))
				Expect(storageObjectModel.Class).To(Equal(core.StringPtr("default")))

				// Construct an instance of the ConfigObject model
				configObjectModel := new(blockchainv2.ConfigObject)
				Expect(configObjectModel).ToNot(BeNil())
				configObjectModel.Enrollment = configObjectEnrollmentModel
				configObjectModel.Msp = configObjectMspModel
				Expect(configObjectModel.Enrollment).To(Equal(configObjectEnrollmentModel))
				Expect(configObjectModel.Msp).To(Equal(configObjectMspModel))

				// Construct an instance of the ConfigOrdererCreate model
				configOrdererCreateModel := new(blockchainv2.ConfigOrdererCreate)
				Expect(configOrdererCreateModel).ToNot(BeNil())
				configOrdererCreateModel.General = configOrdererGeneralModel
				configOrdererCreateModel.Debug = configOrdererDebugModel
				configOrdererCreateModel.Metrics = configOrdererMetricsModel
				Expect(configOrdererCreateModel.General).To(Equal(configOrdererGeneralModel))
				Expect(configOrdererCreateModel.Debug).To(Equal(configOrdererDebugModel))
				Expect(configOrdererCreateModel.Metrics).To(Equal(configOrdererMetricsModel))

				// Construct an instance of the CreateOrdererRaftBodyResources model
				createOrdererRaftBodyResourcesModel := new(blockchainv2.CreateOrdererRaftBodyResources)
				Expect(createOrdererRaftBodyResourcesModel).ToNot(BeNil())
				createOrdererRaftBodyResourcesModel.Orderer = resourceObjectModel
				createOrdererRaftBodyResourcesModel.Proxy = resourceObjectModel
				Expect(createOrdererRaftBodyResourcesModel.Orderer).To(Equal(resourceObjectModel))
				Expect(createOrdererRaftBodyResourcesModel.Proxy).To(Equal(resourceObjectModel))

				// Construct an instance of the CreateOrdererRaftBodyStorage model
				createOrdererRaftBodyStorageModel := new(blockchainv2.CreateOrdererRaftBodyStorage)
				Expect(createOrdererRaftBodyStorageModel).ToNot(BeNil())
				createOrdererRaftBodyStorageModel.Orderer = storageObjectModel
				Expect(createOrdererRaftBodyStorageModel.Orderer).To(Equal(storageObjectModel))

				// Construct an instance of the Hsm model
				hsmModel := new(blockchainv2.Hsm)
				Expect(hsmModel).ToNot(BeNil())
				hsmModel.Pkcs11endpoint = core.StringPtr("tcp://example.com:666")
				Expect(hsmModel.Pkcs11endpoint).To(Equal(core.StringPtr("tcp://example.com:666")))

				// Construct an instance of the CreateOrdererOptions model
				createOrdererOptionsOrdererType := "raft"
				createOrdererOptionsMspID := "Org1"
				createOrdererOptionsDisplayName := "orderer"
				createOrdererOptionsConfig := []blockchainv2.ConfigObject{}
				createOrdererOptionsModel := testService.NewCreateOrdererOptions(createOrdererOptionsOrdererType, createOrdererOptionsMspID, createOrdererOptionsDisplayName, createOrdererOptionsConfig)
				createOrdererOptionsModel.SetOrdererType("raft")
				createOrdererOptionsModel.SetMspID("Org1")
				createOrdererOptionsModel.SetDisplayName("orderer")
				createOrdererOptionsModel.SetConfig([]blockchainv2.ConfigObject{*configObjectModel})
				createOrdererOptionsModel.SetClusterName("ordering service 1")
				createOrdererOptionsModel.SetClusterID("abcde")
				createOrdererOptionsModel.SetExternalAppend("false")
				createOrdererOptionsModel.SetConfigOverride([]blockchainv2.ConfigOrdererCreate{*configOrdererCreateModel})
				createOrdererOptionsModel.SetResources(createOrdererRaftBodyResourcesModel)
				createOrdererOptionsModel.SetStorage(createOrdererRaftBodyStorageModel)
				createOrdererOptionsModel.SetSystemChannelID("testchainid")
				createOrdererOptionsModel.SetZone([]string{"testString"})
				createOrdererOptionsModel.SetTags([]string{"testString"})
				createOrdererOptionsModel.SetRegion([]string{"testString"})
				createOrdererOptionsModel.SetHsm(hsmModel)
				createOrdererOptionsModel.SetVersion("1.4.6-1")
				createOrdererOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createOrdererOptionsModel).ToNot(BeNil())
				Expect(createOrdererOptionsModel.OrdererType).To(Equal(core.StringPtr("raft")))
				Expect(createOrdererOptionsModel.MspID).To(Equal(core.StringPtr("Org1")))
				Expect(createOrdererOptionsModel.DisplayName).To(Equal(core.StringPtr("orderer")))
				Expect(createOrdererOptionsModel.Config).To(Equal([]blockchainv2.ConfigObject{*configObjectModel}))
				Expect(createOrdererOptionsModel.ClusterName).To(Equal(core.StringPtr("ordering service 1")))
				Expect(createOrdererOptionsModel.ClusterID).To(Equal(core.StringPtr("abcde")))
				Expect(createOrdererOptionsModel.ExternalAppend).To(Equal(core.StringPtr("false")))
				Expect(createOrdererOptionsModel.ConfigOverride).To(Equal([]blockchainv2.ConfigOrdererCreate{*configOrdererCreateModel}))
				Expect(createOrdererOptionsModel.Resources).To(Equal(createOrdererRaftBodyResourcesModel))
				Expect(createOrdererOptionsModel.Storage).To(Equal(createOrdererRaftBodyStorageModel))
				Expect(createOrdererOptionsModel.SystemChannelID).To(Equal(core.StringPtr("testchainid")))
				Expect(createOrdererOptionsModel.Zone).To(Equal([]string{"testString"}))
				Expect(createOrdererOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(createOrdererOptionsModel.Region).To(Equal([]string{"testString"}))
				Expect(createOrdererOptionsModel.Hsm).To(Equal(hsmModel))
				Expect(createOrdererOptionsModel.Version).To(Equal(core.StringPtr("1.4.6-1")))
				Expect(createOrdererOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateOrdererRaftBodyResources successfully`, func() {
				var orderer *blockchainv2.ResourceObject = nil
				_, err := testService.NewCreateOrdererRaftBodyResources(orderer)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCreateOrdererRaftBodyStorage successfully`, func() {
				var orderer *blockchainv2.StorageObject = nil
				_, err := testService.NewCreateOrdererRaftBodyStorage(orderer)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCreatePeerBodyStorage successfully`, func() {
				var peer *blockchainv2.StorageObject = nil
				_, err := testService.NewCreatePeerBodyStorage(peer)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCreatePeerOptions successfully`, func() {
				// Construct an instance of the ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy model
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel := new(blockchainv2.ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy)
				Expect(configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel).ToNot(BeNil())
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.RequiredPeerCount = core.Float64Ptr(float64(0))
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.MaxPeerCount = core.Float64Ptr(float64(1))
				Expect(configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.RequiredPeerCount).To(Equal(core.Float64Ptr(float64(0))))
				Expect(configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.MaxPeerCount).To(Equal(core.Float64Ptr(float64(1))))

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv2.BccspPKCS11)
				Expect(bccspPkcS11Model).ToNot(BeNil())
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))
				Expect(bccspPkcS11Model.Label).To(Equal(core.StringPtr("testString")))
				Expect(bccspPkcS11Model.Pin).To(Equal(core.StringPtr("testString")))
				Expect(bccspPkcS11Model.Hash).To(Equal(core.StringPtr("SHA2")))
				Expect(bccspPkcS11Model.Security).To(Equal(core.Float64Ptr(float64(256))))

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv2.BccspSW)
				Expect(bccspSwModel).ToNot(BeNil())
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))
				Expect(bccspSwModel.Hash).To(Equal(core.StringPtr("SHA2")))
				Expect(bccspSwModel.Security).To(Equal(core.Float64Ptr(float64(256))))

				// Construct an instance of the ConfigObjectEnrollmentComponentCatls model
				configObjectEnrollmentComponentCatlsModel := new(blockchainv2.ConfigObjectEnrollmentComponentCatls)
				Expect(configObjectEnrollmentComponentCatlsModel).ToNot(BeNil())
				configObjectEnrollmentComponentCatlsModel.Cacert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				Expect(configObjectEnrollmentComponentCatlsModel.Cacert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))

				// Construct an instance of the ConfigObjectEnrollmentTlsCatls model
				configObjectEnrollmentTlsCatlsModel := new(blockchainv2.ConfigObjectEnrollmentTlsCatls)
				Expect(configObjectEnrollmentTlsCatlsModel).ToNot(BeNil())
				configObjectEnrollmentTlsCatlsModel.Cacert = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				Expect(configObjectEnrollmentTlsCatlsModel.Cacert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))

				// Construct an instance of the ConfigObjectEnrollmentTlsCsr model
				configObjectEnrollmentTlsCsrModel := new(blockchainv2.ConfigObjectEnrollmentTlsCsr)
				Expect(configObjectEnrollmentTlsCsrModel).ToNot(BeNil())
				configObjectEnrollmentTlsCsrModel.Hosts = []string{"testString"}
				Expect(configObjectEnrollmentTlsCsrModel.Hosts).To(Equal([]string{"testString"}))

				// Construct an instance of the ConfigPeerDeliveryclientAddressOverridesItem model
				configPeerDeliveryclientAddressOverridesItemModel := new(blockchainv2.ConfigPeerDeliveryclientAddressOverridesItem)
				Expect(configPeerDeliveryclientAddressOverridesItemModel).ToNot(BeNil())
				configPeerDeliveryclientAddressOverridesItemModel.From = core.StringPtr("n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.To = core.StringPtr("n3a3ec3-myorderer2.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.CaCertsFile = core.StringPtr("my-data/cert.pem")
				Expect(configPeerDeliveryclientAddressOverridesItemModel.From).To(Equal(core.StringPtr("n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")))
				Expect(configPeerDeliveryclientAddressOverridesItemModel.To).To(Equal(core.StringPtr("n3a3ec3-myorderer2.ibp.us-south.containers.appdomain.cloud:7050")))
				Expect(configPeerDeliveryclientAddressOverridesItemModel.CaCertsFile).To(Equal(core.StringPtr("my-data/cert.pem")))

				// Construct an instance of the ConfigPeerGossipElection model
				configPeerGossipElectionModel := new(blockchainv2.ConfigPeerGossipElection)
				Expect(configPeerGossipElectionModel).ToNot(BeNil())
				configPeerGossipElectionModel.StartupGracePeriod = core.StringPtr("15s")
				configPeerGossipElectionModel.MembershipSampleInterval = core.StringPtr("1s")
				configPeerGossipElectionModel.LeaderAliveThreshold = core.StringPtr("10s")
				configPeerGossipElectionModel.LeaderElectionDuration = core.StringPtr("5s")
				Expect(configPeerGossipElectionModel.StartupGracePeriod).To(Equal(core.StringPtr("15s")))
				Expect(configPeerGossipElectionModel.MembershipSampleInterval).To(Equal(core.StringPtr("1s")))
				Expect(configPeerGossipElectionModel.LeaderAliveThreshold).To(Equal(core.StringPtr("10s")))
				Expect(configPeerGossipElectionModel.LeaderElectionDuration).To(Equal(core.StringPtr("5s")))

				// Construct an instance of the ConfigPeerGossipPvtData model
				configPeerGossipPvtDataModel := new(blockchainv2.ConfigPeerGossipPvtData)
				Expect(configPeerGossipPvtDataModel).ToNot(BeNil())
				configPeerGossipPvtDataModel.PullRetryThreshold = core.StringPtr("60s")
				configPeerGossipPvtDataModel.TransientstoreMaxBlockRetention = core.Float64Ptr(float64(1000))
				configPeerGossipPvtDataModel.PushAckTimeout = core.StringPtr("3s")
				configPeerGossipPvtDataModel.BtlPullMargin = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileBatchSize = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileSleepInterval = core.StringPtr("1m")
				configPeerGossipPvtDataModel.ReconciliationEnabled = core.BoolPtr(true)
				configPeerGossipPvtDataModel.SkipPullingInvalidTransactionsDuringCommit = core.BoolPtr(false)
				configPeerGossipPvtDataModel.ImplicitCollectionDisseminationPolicy = configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel
				Expect(configPeerGossipPvtDataModel.PullRetryThreshold).To(Equal(core.StringPtr("60s")))
				Expect(configPeerGossipPvtDataModel.TransientstoreMaxBlockRetention).To(Equal(core.Float64Ptr(float64(1000))))
				Expect(configPeerGossipPvtDataModel.PushAckTimeout).To(Equal(core.StringPtr("3s")))
				Expect(configPeerGossipPvtDataModel.BtlPullMargin).To(Equal(core.Float64Ptr(float64(10))))
				Expect(configPeerGossipPvtDataModel.ReconcileBatchSize).To(Equal(core.Float64Ptr(float64(10))))
				Expect(configPeerGossipPvtDataModel.ReconcileSleepInterval).To(Equal(core.StringPtr("1m")))
				Expect(configPeerGossipPvtDataModel.ReconciliationEnabled).To(Equal(core.BoolPtr(true)))
				Expect(configPeerGossipPvtDataModel.SkipPullingInvalidTransactionsDuringCommit).To(Equal(core.BoolPtr(false)))
				Expect(configPeerGossipPvtDataModel.ImplicitCollectionDisseminationPolicy).To(Equal(configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel))

				// Construct an instance of the ConfigPeerGossipState model
				configPeerGossipStateModel := new(blockchainv2.ConfigPeerGossipState)
				Expect(configPeerGossipStateModel).ToNot(BeNil())
				configPeerGossipStateModel.Enabled = core.BoolPtr(true)
				configPeerGossipStateModel.CheckInterval = core.StringPtr("10s")
				configPeerGossipStateModel.ResponseTimeout = core.StringPtr("3s")
				configPeerGossipStateModel.BatchSize = core.Float64Ptr(float64(10))
				configPeerGossipStateModel.BlockBufferSize = core.Float64Ptr(float64(100))
				configPeerGossipStateModel.MaxRetries = core.Float64Ptr(float64(3))
				Expect(configPeerGossipStateModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(configPeerGossipStateModel.CheckInterval).To(Equal(core.StringPtr("10s")))
				Expect(configPeerGossipStateModel.ResponseTimeout).To(Equal(core.StringPtr("3s")))
				Expect(configPeerGossipStateModel.BatchSize).To(Equal(core.Float64Ptr(float64(10))))
				Expect(configPeerGossipStateModel.BlockBufferSize).To(Equal(core.Float64Ptr(float64(100))))
				Expect(configPeerGossipStateModel.MaxRetries).To(Equal(core.Float64Ptr(float64(3))))

				// Construct an instance of the ConfigPeerKeepaliveClient model
				configPeerKeepaliveClientModel := new(blockchainv2.ConfigPeerKeepaliveClient)
				Expect(configPeerKeepaliveClientModel).ToNot(BeNil())
				configPeerKeepaliveClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveClientModel.Timeout = core.StringPtr("20s")
				Expect(configPeerKeepaliveClientModel.Interval).To(Equal(core.StringPtr("60s")))
				Expect(configPeerKeepaliveClientModel.Timeout).To(Equal(core.StringPtr("20s")))

				// Construct an instance of the ConfigPeerKeepaliveDeliveryClient model
				configPeerKeepaliveDeliveryClientModel := new(blockchainv2.ConfigPeerKeepaliveDeliveryClient)
				Expect(configPeerKeepaliveDeliveryClientModel).ToNot(BeNil())
				configPeerKeepaliveDeliveryClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveDeliveryClientModel.Timeout = core.StringPtr("20s")
				Expect(configPeerKeepaliveDeliveryClientModel.Interval).To(Equal(core.StringPtr("60s")))
				Expect(configPeerKeepaliveDeliveryClientModel.Timeout).To(Equal(core.StringPtr("20s")))

				// Construct an instance of the ConfigPeerLimitsConcurrency model
				configPeerLimitsConcurrencyModel := new(blockchainv2.ConfigPeerLimitsConcurrency)
				Expect(configPeerLimitsConcurrencyModel).ToNot(BeNil())
				configPeerLimitsConcurrencyModel.EndorserService = map[string]interface{}{"anyKey": "anyValue"}
				configPeerLimitsConcurrencyModel.DeliverService = map[string]interface{}{"anyKey": "anyValue"}
				Expect(configPeerLimitsConcurrencyModel.EndorserService).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(configPeerLimitsConcurrencyModel.DeliverService).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv2.Bccsp)
				Expect(bccspModel).ToNot(BeNil())
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model
				Expect(bccspModel.Default).To(Equal(core.StringPtr("SW")))
				Expect(bccspModel.SW).To(Equal(bccspSwModel))
				Expect(bccspModel.PKCS11).To(Equal(bccspPkcS11Model))

				// Construct an instance of the ConfigObjectEnrollmentComponent model
				configObjectEnrollmentComponentModel := new(blockchainv2.ConfigObjectEnrollmentComponent)
				Expect(configObjectEnrollmentComponentModel).ToNot(BeNil())
				configObjectEnrollmentComponentModel.Cahost = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				configObjectEnrollmentComponentModel.Caport = core.Float64Ptr(float64(7054))
				configObjectEnrollmentComponentModel.Caname = core.StringPtr("ca")
				configObjectEnrollmentComponentModel.Catls = configObjectEnrollmentComponentCatlsModel
				configObjectEnrollmentComponentModel.Enrollid = core.StringPtr("admin")
				configObjectEnrollmentComponentModel.Enrollsecret = core.StringPtr("password")
				configObjectEnrollmentComponentModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				Expect(configObjectEnrollmentComponentModel.Cahost).To(Equal(core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")))
				Expect(configObjectEnrollmentComponentModel.Caport).To(Equal(core.Float64Ptr(float64(7054))))
				Expect(configObjectEnrollmentComponentModel.Caname).To(Equal(core.StringPtr("ca")))
				Expect(configObjectEnrollmentComponentModel.Catls).To(Equal(configObjectEnrollmentComponentCatlsModel))
				Expect(configObjectEnrollmentComponentModel.Enrollid).To(Equal(core.StringPtr("admin")))
				Expect(configObjectEnrollmentComponentModel.Enrollsecret).To(Equal(core.StringPtr("password")))
				Expect(configObjectEnrollmentComponentModel.Admincerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))

				// Construct an instance of the ConfigObjectEnrollmentTls model
				configObjectEnrollmentTlsModel := new(blockchainv2.ConfigObjectEnrollmentTls)
				Expect(configObjectEnrollmentTlsModel).ToNot(BeNil())
				configObjectEnrollmentTlsModel.Cahost = core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")
				configObjectEnrollmentTlsModel.Caport = core.Float64Ptr(float64(7054))
				configObjectEnrollmentTlsModel.Caname = core.StringPtr("tlsca")
				configObjectEnrollmentTlsModel.Catls = configObjectEnrollmentTlsCatlsModel
				configObjectEnrollmentTlsModel.Enrollid = core.StringPtr("admin")
				configObjectEnrollmentTlsModel.Enrollsecret = core.StringPtr("password")
				configObjectEnrollmentTlsModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				configObjectEnrollmentTlsModel.Csr = configObjectEnrollmentTlsCsrModel
				Expect(configObjectEnrollmentTlsModel.Cahost).To(Equal(core.StringPtr("n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud")))
				Expect(configObjectEnrollmentTlsModel.Caport).To(Equal(core.Float64Ptr(float64(7054))))
				Expect(configObjectEnrollmentTlsModel.Caname).To(Equal(core.StringPtr("tlsca")))
				Expect(configObjectEnrollmentTlsModel.Catls).To(Equal(configObjectEnrollmentTlsCatlsModel))
				Expect(configObjectEnrollmentTlsModel.Enrollid).To(Equal(core.StringPtr("admin")))
				Expect(configObjectEnrollmentTlsModel.Enrollsecret).To(Equal(core.StringPtr("password")))
				Expect(configObjectEnrollmentTlsModel.Admincerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(configObjectEnrollmentTlsModel.Csr).To(Equal(configObjectEnrollmentTlsCsrModel))

				// Construct an instance of the ConfigPeerAdminService model
				configPeerAdminServiceModel := new(blockchainv2.ConfigPeerAdminService)
				Expect(configPeerAdminServiceModel).ToNot(BeNil())
				configPeerAdminServiceModel.ListenAddress = core.StringPtr("0.0.0.0:7051")
				Expect(configPeerAdminServiceModel.ListenAddress).To(Equal(core.StringPtr("0.0.0.0:7051")))

				// Construct an instance of the ConfigPeerAuthentication model
				configPeerAuthenticationModel := new(blockchainv2.ConfigPeerAuthentication)
				Expect(configPeerAuthenticationModel).ToNot(BeNil())
				configPeerAuthenticationModel.Timewindow = core.StringPtr("15m")
				Expect(configPeerAuthenticationModel.Timewindow).To(Equal(core.StringPtr("15m")))

				// Construct an instance of the ConfigPeerChaincodeExternalBuildersItem model
				configPeerChaincodeExternalBuildersItemModel := new(blockchainv2.ConfigPeerChaincodeExternalBuildersItem)
				Expect(configPeerChaincodeExternalBuildersItemModel).ToNot(BeNil())
				configPeerChaincodeExternalBuildersItemModel.Path = core.StringPtr("/path/to/directory")
				configPeerChaincodeExternalBuildersItemModel.Name = core.StringPtr("descriptive-build-name")
				configPeerChaincodeExternalBuildersItemModel.EnvironmentWhitelist = []string{"GOPROXY"}
				Expect(configPeerChaincodeExternalBuildersItemModel.Path).To(Equal(core.StringPtr("/path/to/directory")))
				Expect(configPeerChaincodeExternalBuildersItemModel.Name).To(Equal(core.StringPtr("descriptive-build-name")))
				Expect(configPeerChaincodeExternalBuildersItemModel.EnvironmentWhitelist).To(Equal([]string{"GOPROXY"}))

				// Construct an instance of the ConfigPeerChaincodeGolang model
				configPeerChaincodeGolangModel := new(blockchainv2.ConfigPeerChaincodeGolang)
				Expect(configPeerChaincodeGolangModel).ToNot(BeNil())
				configPeerChaincodeGolangModel.DynamicLink = core.BoolPtr(false)
				Expect(configPeerChaincodeGolangModel.DynamicLink).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the ConfigPeerChaincodeLogging model
				configPeerChaincodeLoggingModel := new(blockchainv2.ConfigPeerChaincodeLogging)
				Expect(configPeerChaincodeLoggingModel).ToNot(BeNil())
				configPeerChaincodeLoggingModel.Level = core.StringPtr("info")
				configPeerChaincodeLoggingModel.Shim = core.StringPtr("warning")
				configPeerChaincodeLoggingModel.Format = core.StringPtr("%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}")
				Expect(configPeerChaincodeLoggingModel.Level).To(Equal(core.StringPtr("info")))
				Expect(configPeerChaincodeLoggingModel.Shim).To(Equal(core.StringPtr("warning")))
				Expect(configPeerChaincodeLoggingModel.Format).To(Equal(core.StringPtr("%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}")))

				// Construct an instance of the ConfigPeerChaincodeSystem model
				configPeerChaincodeSystemModel := new(blockchainv2.ConfigPeerChaincodeSystem)
				Expect(configPeerChaincodeSystemModel).ToNot(BeNil())
				configPeerChaincodeSystemModel.Cscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Lscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Escc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Vscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Qscc = core.BoolPtr(true)
				Expect(configPeerChaincodeSystemModel.Cscc).To(Equal(core.BoolPtr(true)))
				Expect(configPeerChaincodeSystemModel.Lscc).To(Equal(core.BoolPtr(true)))
				Expect(configPeerChaincodeSystemModel.Escc).To(Equal(core.BoolPtr(true)))
				Expect(configPeerChaincodeSystemModel.Vscc).To(Equal(core.BoolPtr(true)))
				Expect(configPeerChaincodeSystemModel.Qscc).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the ConfigPeerClient model
				configPeerClientModel := new(blockchainv2.ConfigPeerClient)
				Expect(configPeerClientModel).ToNot(BeNil())
				configPeerClientModel.ConnTimeout = core.StringPtr("2s")
				Expect(configPeerClientModel.ConnTimeout).To(Equal(core.StringPtr("2s")))

				// Construct an instance of the ConfigPeerDeliveryclient model
				configPeerDeliveryclientModel := new(blockchainv2.ConfigPeerDeliveryclient)
				Expect(configPeerDeliveryclientModel).ToNot(BeNil())
				configPeerDeliveryclientModel.ReconnectTotalTimeThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.ConnTimeout = core.StringPtr("2s")
				configPeerDeliveryclientModel.ReConnectBackoffThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.AddressOverrides = []blockchainv2.ConfigPeerDeliveryclientAddressOverridesItem{*configPeerDeliveryclientAddressOverridesItemModel}
				Expect(configPeerDeliveryclientModel.ReconnectTotalTimeThreshold).To(Equal(core.StringPtr("60m")))
				Expect(configPeerDeliveryclientModel.ConnTimeout).To(Equal(core.StringPtr("2s")))
				Expect(configPeerDeliveryclientModel.ReConnectBackoffThreshold).To(Equal(core.StringPtr("60m")))
				Expect(configPeerDeliveryclientModel.AddressOverrides).To(Equal([]blockchainv2.ConfigPeerDeliveryclientAddressOverridesItem{*configPeerDeliveryclientAddressOverridesItemModel}))

				// Construct an instance of the ConfigPeerDiscovery model
				configPeerDiscoveryModel := new(blockchainv2.ConfigPeerDiscovery)
				Expect(configPeerDiscoveryModel).ToNot(BeNil())
				configPeerDiscoveryModel.Enabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheEnabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheMaxSize = core.Float64Ptr(float64(1000))
				configPeerDiscoveryModel.AuthCachePurgeRetentionRatio = core.Float64Ptr(float64(0.75))
				configPeerDiscoveryModel.OrgMembersAllowedAccess = core.BoolPtr(false)
				Expect(configPeerDiscoveryModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(configPeerDiscoveryModel.AuthCacheEnabled).To(Equal(core.BoolPtr(true)))
				Expect(configPeerDiscoveryModel.AuthCacheMaxSize).To(Equal(core.Float64Ptr(float64(1000))))
				Expect(configPeerDiscoveryModel.AuthCachePurgeRetentionRatio).To(Equal(core.Float64Ptr(float64(0.75))))
				Expect(configPeerDiscoveryModel.OrgMembersAllowedAccess).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the ConfigPeerGossip model
				configPeerGossipModel := new(blockchainv2.ConfigPeerGossip)
				Expect(configPeerGossipModel).ToNot(BeNil())
				configPeerGossipModel.UseLeaderElection = core.BoolPtr(true)
				configPeerGossipModel.OrgLeader = core.BoolPtr(false)
				configPeerGossipModel.MembershipTrackerInterval = core.StringPtr("5s")
				configPeerGossipModel.MaxBlockCountToStore = core.Float64Ptr(float64(100))
				configPeerGossipModel.MaxPropagationBurstLatency = core.StringPtr("10ms")
				configPeerGossipModel.MaxPropagationBurstSize = core.Float64Ptr(float64(10))
				configPeerGossipModel.PropagateIterations = core.Float64Ptr(float64(3))
				configPeerGossipModel.PullInterval = core.StringPtr("4s")
				configPeerGossipModel.PullPeerNum = core.Float64Ptr(float64(3))
				configPeerGossipModel.RequestStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.PublishStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.StateInfoRetentionInterval = core.StringPtr("0s")
				configPeerGossipModel.PublishCertPeriod = core.StringPtr("10s")
				configPeerGossipModel.SkipBlockVerification = core.BoolPtr(false)
				configPeerGossipModel.DialTimeout = core.StringPtr("3s")
				configPeerGossipModel.ConnTimeout = core.StringPtr("2s")
				configPeerGossipModel.RecvBuffSize = core.Float64Ptr(float64(20))
				configPeerGossipModel.SendBuffSize = core.Float64Ptr(float64(200))
				configPeerGossipModel.DigestWaitTime = core.StringPtr("1s")
				configPeerGossipModel.RequestWaitTime = core.StringPtr("1500ms")
				configPeerGossipModel.ResponseWaitTime = core.StringPtr("2s")
				configPeerGossipModel.AliveTimeInterval = core.StringPtr("5s")
				configPeerGossipModel.AliveExpirationTimeout = core.StringPtr("25s")
				configPeerGossipModel.ReconnectInterval = core.StringPtr("25s")
				configPeerGossipModel.Election = configPeerGossipElectionModel
				configPeerGossipModel.PvtData = configPeerGossipPvtDataModel
				configPeerGossipModel.State = configPeerGossipStateModel
				Expect(configPeerGossipModel.UseLeaderElection).To(Equal(core.BoolPtr(true)))
				Expect(configPeerGossipModel.OrgLeader).To(Equal(core.BoolPtr(false)))
				Expect(configPeerGossipModel.MembershipTrackerInterval).To(Equal(core.StringPtr("5s")))
				Expect(configPeerGossipModel.MaxBlockCountToStore).To(Equal(core.Float64Ptr(float64(100))))
				Expect(configPeerGossipModel.MaxPropagationBurstLatency).To(Equal(core.StringPtr("10ms")))
				Expect(configPeerGossipModel.MaxPropagationBurstSize).To(Equal(core.Float64Ptr(float64(10))))
				Expect(configPeerGossipModel.PropagateIterations).To(Equal(core.Float64Ptr(float64(3))))
				Expect(configPeerGossipModel.PullInterval).To(Equal(core.StringPtr("4s")))
				Expect(configPeerGossipModel.PullPeerNum).To(Equal(core.Float64Ptr(float64(3))))
				Expect(configPeerGossipModel.RequestStateInfoInterval).To(Equal(core.StringPtr("4s")))
				Expect(configPeerGossipModel.PublishStateInfoInterval).To(Equal(core.StringPtr("4s")))
				Expect(configPeerGossipModel.StateInfoRetentionInterval).To(Equal(core.StringPtr("0s")))
				Expect(configPeerGossipModel.PublishCertPeriod).To(Equal(core.StringPtr("10s")))
				Expect(configPeerGossipModel.SkipBlockVerification).To(Equal(core.BoolPtr(false)))
				Expect(configPeerGossipModel.DialTimeout).To(Equal(core.StringPtr("3s")))
				Expect(configPeerGossipModel.ConnTimeout).To(Equal(core.StringPtr("2s")))
				Expect(configPeerGossipModel.RecvBuffSize).To(Equal(core.Float64Ptr(float64(20))))
				Expect(configPeerGossipModel.SendBuffSize).To(Equal(core.Float64Ptr(float64(200))))
				Expect(configPeerGossipModel.DigestWaitTime).To(Equal(core.StringPtr("1s")))
				Expect(configPeerGossipModel.RequestWaitTime).To(Equal(core.StringPtr("1500ms")))
				Expect(configPeerGossipModel.ResponseWaitTime).To(Equal(core.StringPtr("2s")))
				Expect(configPeerGossipModel.AliveTimeInterval).To(Equal(core.StringPtr("5s")))
				Expect(configPeerGossipModel.AliveExpirationTimeout).To(Equal(core.StringPtr("25s")))
				Expect(configPeerGossipModel.ReconnectInterval).To(Equal(core.StringPtr("25s")))
				Expect(configPeerGossipModel.Election).To(Equal(configPeerGossipElectionModel))
				Expect(configPeerGossipModel.PvtData).To(Equal(configPeerGossipPvtDataModel))
				Expect(configPeerGossipModel.State).To(Equal(configPeerGossipStateModel))

				// Construct an instance of the ConfigPeerKeepalive model
				configPeerKeepaliveModel := new(blockchainv2.ConfigPeerKeepalive)
				Expect(configPeerKeepaliveModel).ToNot(BeNil())
				configPeerKeepaliveModel.MinInterval = core.StringPtr("60s")
				configPeerKeepaliveModel.Client = configPeerKeepaliveClientModel
				configPeerKeepaliveModel.DeliveryClient = configPeerKeepaliveDeliveryClientModel
				Expect(configPeerKeepaliveModel.MinInterval).To(Equal(core.StringPtr("60s")))
				Expect(configPeerKeepaliveModel.Client).To(Equal(configPeerKeepaliveClientModel))
				Expect(configPeerKeepaliveModel.DeliveryClient).To(Equal(configPeerKeepaliveDeliveryClientModel))

				// Construct an instance of the ConfigPeerLimits model
				configPeerLimitsModel := new(blockchainv2.ConfigPeerLimits)
				Expect(configPeerLimitsModel).ToNot(BeNil())
				configPeerLimitsModel.Concurrency = configPeerLimitsConcurrencyModel
				Expect(configPeerLimitsModel.Concurrency).To(Equal(configPeerLimitsConcurrencyModel))

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv2.MetricsStatsd)
				Expect(metricsStatsdModel).ToNot(BeNil())
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")
				Expect(metricsStatsdModel.Network).To(Equal(core.StringPtr("udp")))
				Expect(metricsStatsdModel.Address).To(Equal(core.StringPtr("127.0.0.1:8125")))
				Expect(metricsStatsdModel.WriteInterval).To(Equal(core.StringPtr("10s")))
				Expect(metricsStatsdModel.Prefix).To(Equal(core.StringPtr("server")))

				// Construct an instance of the MspConfigData model
				mspConfigDataModel := new(blockchainv2.MspConfigData)
				Expect(mspConfigDataModel).ToNot(BeNil())
				mspConfigDataModel.Keystore = core.StringPtr("testString")
				mspConfigDataModel.Signcerts = core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				mspConfigDataModel.Cacerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				mspConfigDataModel.Intermediatecerts = []string{"testString"}
				mspConfigDataModel.Admincerts = []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				Expect(mspConfigDataModel.Keystore).To(Equal(core.StringPtr("testString")))
				Expect(mspConfigDataModel.Signcerts).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(mspConfigDataModel.Cacerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(mspConfigDataModel.Intermediatecerts).To(Equal([]string{"testString"}))
				Expect(mspConfigDataModel.Admincerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv2.ResourceLimits)
				Expect(resourceLimitsModel).ToNot(BeNil())
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")
				Expect(resourceLimitsModel.Cpu).To(Equal(core.StringPtr("100m")))
				Expect(resourceLimitsModel.Memory).To(Equal(core.StringPtr("256MiB")))

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv2.ResourceRequests)
				Expect(resourceRequestsModel).ToNot(BeNil())
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")
				Expect(resourceRequestsModel.Cpu).To(Equal(core.StringPtr("100m")))
				Expect(resourceRequestsModel.Memory).To(Equal(core.StringPtr("256MiB")))

				// Construct an instance of the ConfigObjectEnrollment model
				configObjectEnrollmentModel := new(blockchainv2.ConfigObjectEnrollment)
				Expect(configObjectEnrollmentModel).ToNot(BeNil())
				configObjectEnrollmentModel.Component = configObjectEnrollmentComponentModel
				configObjectEnrollmentModel.Tls = configObjectEnrollmentTlsModel
				Expect(configObjectEnrollmentModel.Component).To(Equal(configObjectEnrollmentComponentModel))
				Expect(configObjectEnrollmentModel.Tls).To(Equal(configObjectEnrollmentTlsModel))

				// Construct an instance of the ConfigObjectMsp model
				configObjectMspModel := new(blockchainv2.ConfigObjectMsp)
				Expect(configObjectMspModel).ToNot(BeNil())
				configObjectMspModel.Component = mspConfigDataModel
				configObjectMspModel.Tls = mspConfigDataModel
				configObjectMspModel.Clientauth = mspConfigDataModel
				Expect(configObjectMspModel.Component).To(Equal(mspConfigDataModel))
				Expect(configObjectMspModel.Tls).To(Equal(mspConfigDataModel))
				Expect(configObjectMspModel.Clientauth).To(Equal(mspConfigDataModel))

				// Construct an instance of the ConfigPeerChaincode model
				configPeerChaincodeModel := new(blockchainv2.ConfigPeerChaincode)
				Expect(configPeerChaincodeModel).ToNot(BeNil())
				configPeerChaincodeModel.Golang = configPeerChaincodeGolangModel
				configPeerChaincodeModel.ExternalBuilders = []blockchainv2.ConfigPeerChaincodeExternalBuildersItem{*configPeerChaincodeExternalBuildersItemModel}
				configPeerChaincodeModel.InstallTimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Startuptimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Executetimeout = core.StringPtr("30s")
				configPeerChaincodeModel.System = configPeerChaincodeSystemModel
				configPeerChaincodeModel.Logging = configPeerChaincodeLoggingModel
				Expect(configPeerChaincodeModel.Golang).To(Equal(configPeerChaincodeGolangModel))
				Expect(configPeerChaincodeModel.ExternalBuilders).To(Equal([]blockchainv2.ConfigPeerChaincodeExternalBuildersItem{*configPeerChaincodeExternalBuildersItemModel}))
				Expect(configPeerChaincodeModel.InstallTimeout).To(Equal(core.StringPtr("300s")))
				Expect(configPeerChaincodeModel.Startuptimeout).To(Equal(core.StringPtr("300s")))
				Expect(configPeerChaincodeModel.Executetimeout).To(Equal(core.StringPtr("30s")))
				Expect(configPeerChaincodeModel.System).To(Equal(configPeerChaincodeSystemModel))
				Expect(configPeerChaincodeModel.Logging).To(Equal(configPeerChaincodeLoggingModel))

				// Construct an instance of the ConfigPeerCreatePeer model
				configPeerCreatePeerModel := new(blockchainv2.ConfigPeerCreatePeer)
				Expect(configPeerCreatePeerModel).ToNot(BeNil())
				configPeerCreatePeerModel.ID = core.StringPtr("john-doe")
				configPeerCreatePeerModel.NetworkID = core.StringPtr("dev")
				configPeerCreatePeerModel.Keepalive = configPeerKeepaliveModel
				configPeerCreatePeerModel.Gossip = configPeerGossipModel
				configPeerCreatePeerModel.Authentication = configPeerAuthenticationModel
				configPeerCreatePeerModel.BCCSP = bccspModel
				configPeerCreatePeerModel.Client = configPeerClientModel
				configPeerCreatePeerModel.Deliveryclient = configPeerDeliveryclientModel
				configPeerCreatePeerModel.AdminService = configPeerAdminServiceModel
				configPeerCreatePeerModel.ValidatorPoolSize = core.Float64Ptr(float64(8))
				configPeerCreatePeerModel.Discovery = configPeerDiscoveryModel
				configPeerCreatePeerModel.Limits = configPeerLimitsModel
				Expect(configPeerCreatePeerModel.ID).To(Equal(core.StringPtr("john-doe")))
				Expect(configPeerCreatePeerModel.NetworkID).To(Equal(core.StringPtr("dev")))
				Expect(configPeerCreatePeerModel.Keepalive).To(Equal(configPeerKeepaliveModel))
				Expect(configPeerCreatePeerModel.Gossip).To(Equal(configPeerGossipModel))
				Expect(configPeerCreatePeerModel.Authentication).To(Equal(configPeerAuthenticationModel))
				Expect(configPeerCreatePeerModel.BCCSP).To(Equal(bccspModel))
				Expect(configPeerCreatePeerModel.Client).To(Equal(configPeerClientModel))
				Expect(configPeerCreatePeerModel.Deliveryclient).To(Equal(configPeerDeliveryclientModel))
				Expect(configPeerCreatePeerModel.AdminService).To(Equal(configPeerAdminServiceModel))
				Expect(configPeerCreatePeerModel.ValidatorPoolSize).To(Equal(core.Float64Ptr(float64(8))))
				Expect(configPeerCreatePeerModel.Discovery).To(Equal(configPeerDiscoveryModel))
				Expect(configPeerCreatePeerModel.Limits).To(Equal(configPeerLimitsModel))

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv2.Metrics)
				Expect(metricsModel).ToNot(BeNil())
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel
				Expect(metricsModel.Provider).To(Equal(core.StringPtr("prometheus")))
				Expect(metricsModel.Statsd).To(Equal(metricsStatsdModel))

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv2.ResourceObject)
				Expect(resourceObjectModel).ToNot(BeNil())
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel
				Expect(resourceObjectModel.Requests).To(Equal(resourceRequestsModel))
				Expect(resourceObjectModel.Limits).To(Equal(resourceLimitsModel))

				// Construct an instance of the ResourceObjectCouchDb model
				resourceObjectCouchDbModel := new(blockchainv2.ResourceObjectCouchDb)
				Expect(resourceObjectCouchDbModel).ToNot(BeNil())
				resourceObjectCouchDbModel.Requests = resourceRequestsModel
				resourceObjectCouchDbModel.Limits = resourceLimitsModel
				Expect(resourceObjectCouchDbModel.Requests).To(Equal(resourceRequestsModel))
				Expect(resourceObjectCouchDbModel.Limits).To(Equal(resourceLimitsModel))

				// Construct an instance of the ResourceObjectFabV1 model
				resourceObjectFabV1Model := new(blockchainv2.ResourceObjectFabV1)
				Expect(resourceObjectFabV1Model).ToNot(BeNil())
				resourceObjectFabV1Model.Requests = resourceRequestsModel
				resourceObjectFabV1Model.Limits = resourceLimitsModel
				Expect(resourceObjectFabV1Model.Requests).To(Equal(resourceRequestsModel))
				Expect(resourceObjectFabV1Model.Limits).To(Equal(resourceLimitsModel))

				// Construct an instance of the ResourceObjectFabV2 model
				resourceObjectFabV2Model := new(blockchainv2.ResourceObjectFabV2)
				Expect(resourceObjectFabV2Model).ToNot(BeNil())
				resourceObjectFabV2Model.Requests = resourceRequestsModel
				resourceObjectFabV2Model.Limits = resourceLimitsModel
				Expect(resourceObjectFabV2Model.Requests).To(Equal(resourceRequestsModel))
				Expect(resourceObjectFabV2Model.Limits).To(Equal(resourceLimitsModel))

				// Construct an instance of the StorageObject model
				storageObjectModel := new(blockchainv2.StorageObject)
				Expect(storageObjectModel).ToNot(BeNil())
				storageObjectModel.Size = core.StringPtr("4GiB")
				storageObjectModel.Class = core.StringPtr("default")
				Expect(storageObjectModel.Size).To(Equal(core.StringPtr("4GiB")))
				Expect(storageObjectModel.Class).To(Equal(core.StringPtr("default")))

				// Construct an instance of the ConfigObject model
				configObjectModel := new(blockchainv2.ConfigObject)
				Expect(configObjectModel).ToNot(BeNil())
				configObjectModel.Enrollment = configObjectEnrollmentModel
				configObjectModel.Msp = configObjectMspModel
				Expect(configObjectModel.Enrollment).To(Equal(configObjectEnrollmentModel))
				Expect(configObjectModel.Msp).To(Equal(configObjectMspModel))

				// Construct an instance of the ConfigPeerCreate model
				configPeerCreateModel := new(blockchainv2.ConfigPeerCreate)
				Expect(configPeerCreateModel).ToNot(BeNil())
				configPeerCreateModel.Peer = configPeerCreatePeerModel
				configPeerCreateModel.Chaincode = configPeerChaincodeModel
				configPeerCreateModel.Metrics = metricsModel
				Expect(configPeerCreateModel.Peer).To(Equal(configPeerCreatePeerModel))
				Expect(configPeerCreateModel.Chaincode).To(Equal(configPeerChaincodeModel))
				Expect(configPeerCreateModel.Metrics).To(Equal(metricsModel))

				// Construct an instance of the CreatePeerBodyStorage model
				createPeerBodyStorageModel := new(blockchainv2.CreatePeerBodyStorage)
				Expect(createPeerBodyStorageModel).ToNot(BeNil())
				createPeerBodyStorageModel.Peer = storageObjectModel
				createPeerBodyStorageModel.Statedb = storageObjectModel
				Expect(createPeerBodyStorageModel.Peer).To(Equal(storageObjectModel))
				Expect(createPeerBodyStorageModel.Statedb).To(Equal(storageObjectModel))

				// Construct an instance of the Hsm model
				hsmModel := new(blockchainv2.Hsm)
				Expect(hsmModel).ToNot(BeNil())
				hsmModel.Pkcs11endpoint = core.StringPtr("tcp://example.com:666")
				Expect(hsmModel.Pkcs11endpoint).To(Equal(core.StringPtr("tcp://example.com:666")))

				// Construct an instance of the PeerResources model
				peerResourcesModel := new(blockchainv2.PeerResources)
				Expect(peerResourcesModel).ToNot(BeNil())
				peerResourcesModel.Chaincodelauncher = resourceObjectFabV2Model
				peerResourcesModel.Couchdb = resourceObjectCouchDbModel
				peerResourcesModel.Statedb = resourceObjectModel
				peerResourcesModel.Dind = resourceObjectFabV1Model
				peerResourcesModel.Fluentd = resourceObjectFabV1Model
				peerResourcesModel.Peer = resourceObjectModel
				peerResourcesModel.Proxy = resourceObjectModel
				Expect(peerResourcesModel.Chaincodelauncher).To(Equal(resourceObjectFabV2Model))
				Expect(peerResourcesModel.Couchdb).To(Equal(resourceObjectCouchDbModel))
				Expect(peerResourcesModel.Statedb).To(Equal(resourceObjectModel))
				Expect(peerResourcesModel.Dind).To(Equal(resourceObjectFabV1Model))
				Expect(peerResourcesModel.Fluentd).To(Equal(resourceObjectFabV1Model))
				Expect(peerResourcesModel.Peer).To(Equal(resourceObjectModel))
				Expect(peerResourcesModel.Proxy).To(Equal(resourceObjectModel))

				// Construct an instance of the CreatePeerOptions model
				createPeerOptionsMspID := "Org1"
				createPeerOptionsDisplayName := "My Peer"
				var createPeerOptionsConfig *blockchainv2.ConfigObject = nil
				createPeerOptionsModel := testService.NewCreatePeerOptions(createPeerOptionsMspID, createPeerOptionsDisplayName, createPeerOptionsConfig)
				createPeerOptionsModel.SetMspID("Org1")
				createPeerOptionsModel.SetDisplayName("My Peer")
				createPeerOptionsModel.SetConfig(configObjectModel)
				createPeerOptionsModel.SetConfigOverride(configPeerCreateModel)
				createPeerOptionsModel.SetResources(peerResourcesModel)
				createPeerOptionsModel.SetStorage(createPeerBodyStorageModel)
				createPeerOptionsModel.SetZone("testString")
				createPeerOptionsModel.SetStateDb("couchdb")
				createPeerOptionsModel.SetTags([]string{"testString"})
				createPeerOptionsModel.SetHsm(hsmModel)
				createPeerOptionsModel.SetRegion("testString")
				createPeerOptionsModel.SetVersion("1.4.6-1")
				createPeerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createPeerOptionsModel).ToNot(BeNil())
				Expect(createPeerOptionsModel.MspID).To(Equal(core.StringPtr("Org1")))
				Expect(createPeerOptionsModel.DisplayName).To(Equal(core.StringPtr("My Peer")))
				Expect(createPeerOptionsModel.Config).To(Equal(configObjectModel))
				Expect(createPeerOptionsModel.ConfigOverride).To(Equal(configPeerCreateModel))
				Expect(createPeerOptionsModel.Resources).To(Equal(peerResourcesModel))
				Expect(createPeerOptionsModel.Storage).To(Equal(createPeerBodyStorageModel))
				Expect(createPeerOptionsModel.Zone).To(Equal(core.StringPtr("testString")))
				Expect(createPeerOptionsModel.StateDb).To(Equal(core.StringPtr("couchdb")))
				Expect(createPeerOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(createPeerOptionsModel.Hsm).To(Equal(hsmModel))
				Expect(createPeerOptionsModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(createPeerOptionsModel.Version).To(Equal(core.StringPtr("1.4.6-1")))
				Expect(createPeerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAllComponentsOptions successfully`, func() {
				// Construct an instance of the DeleteAllComponentsOptions model
				deleteAllComponentsOptionsModel := testService.NewDeleteAllComponentsOptions()
				deleteAllComponentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAllComponentsOptionsModel).ToNot(BeNil())
				Expect(deleteAllComponentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAllNotificationsOptions successfully`, func() {
				// Construct an instance of the DeleteAllNotificationsOptions model
				deleteAllNotificationsOptionsModel := testService.NewDeleteAllNotificationsOptions()
				deleteAllNotificationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAllNotificationsOptionsModel).ToNot(BeNil())
				Expect(deleteAllNotificationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAllSessionsOptions successfully`, func() {
				// Construct an instance of the DeleteAllSessionsOptions model
				deleteAllSessionsOptionsModel := testService.NewDeleteAllSessionsOptions()
				deleteAllSessionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAllSessionsOptionsModel).ToNot(BeNil())
				Expect(deleteAllSessionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteComponentOptions successfully`, func() {
				// Construct an instance of the DeleteComponentOptions model
				id := "testString"
				deleteComponentOptionsModel := testService.NewDeleteComponentOptions(id)
				deleteComponentOptionsModel.SetID("testString")
				deleteComponentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteComponentOptionsModel).ToNot(BeNil())
				Expect(deleteComponentOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteComponentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteComponentsByTagOptions successfully`, func() {
				// Construct an instance of the DeleteComponentsByTagOptions model
				tag := "testString"
				deleteComponentsByTagOptionsModel := testService.NewDeleteComponentsByTagOptions(tag)
				deleteComponentsByTagOptionsModel.SetTag("testString")
				deleteComponentsByTagOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteComponentsByTagOptionsModel).ToNot(BeNil())
				Expect(deleteComponentsByTagOptionsModel.Tag).To(Equal(core.StringPtr("testString")))
				Expect(deleteComponentsByTagOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteSigTxOptions successfully`, func() {
				// Construct an instance of the DeleteSigTxOptions model
				id := "testString"
				deleteSigTxOptionsModel := testService.NewDeleteSigTxOptions(id)
				deleteSigTxOptionsModel.SetID("testString")
				deleteSigTxOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteSigTxOptionsModel).ToNot(BeNil())
				Expect(deleteSigTxOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSigTxOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEditAdminCertsOptions successfully`, func() {
				// Construct an instance of the EditAdminCertsOptions model
				id := "testString"
				editAdminCertsOptionsModel := testService.NewEditAdminCertsOptions(id)
				editAdminCertsOptionsModel.SetID("testString")
				editAdminCertsOptionsModel.SetAppendAdminCerts([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="})
				editAdminCertsOptionsModel.SetRemoveAdminCerts([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="})
				editAdminCertsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(editAdminCertsOptionsModel).ToNot(BeNil())
				Expect(editAdminCertsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(editAdminCertsOptionsModel.AppendAdminCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(editAdminCertsOptionsModel.RemoveAdminCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(editAdminCertsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEditCaOptions successfully`, func() {
				// Construct an instance of the EditCaOptions model
				id := "testString"
				editCaOptionsModel := testService.NewEditCaOptions(id)
				editCaOptionsModel.SetID("testString")
				editCaOptionsModel.SetDisplayName("My CA")
				editCaOptionsModel.SetApiURL("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:7054")
				editCaOptionsModel.SetOperationsURL("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443")
				editCaOptionsModel.SetCaName("ca")
				editCaOptionsModel.SetLocation("ibmcloud")
				editCaOptionsModel.SetTags([]string{"testString"})
				editCaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(editCaOptionsModel).ToNot(BeNil())
				Expect(editCaOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(editCaOptionsModel.DisplayName).To(Equal(core.StringPtr("My CA")))
				Expect(editCaOptionsModel.ApiURL).To(Equal(core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:7054")))
				Expect(editCaOptionsModel.OperationsURL).To(Equal(core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443")))
				Expect(editCaOptionsModel.CaName).To(Equal(core.StringPtr("ca")))
				Expect(editCaOptionsModel.Location).To(Equal(core.StringPtr("ibmcloud")))
				Expect(editCaOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(editCaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEditMspOptions successfully`, func() {
				// Construct an instance of the EditMspOptions model
				id := "testString"
				editMspOptionsModel := testService.NewEditMspOptions(id)
				editMspOptionsModel.SetID("testString")
				editMspOptionsModel.SetMspID("Org1")
				editMspOptionsModel.SetDisplayName("My Peer")
				editMspOptionsModel.SetRootCerts([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="})
				editMspOptionsModel.SetIntermediateCerts([]string{"testString"})
				editMspOptionsModel.SetAdmins([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="})
				editMspOptionsModel.SetTlsRootCerts([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="})
				editMspOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(editMspOptionsModel).ToNot(BeNil())
				Expect(editMspOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(editMspOptionsModel.MspID).To(Equal(core.StringPtr("Org1")))
				Expect(editMspOptionsModel.DisplayName).To(Equal(core.StringPtr("My Peer")))
				Expect(editMspOptionsModel.RootCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(editMspOptionsModel.IntermediateCerts).To(Equal([]string{"testString"}))
				Expect(editMspOptionsModel.Admins).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(editMspOptionsModel.TlsRootCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(editMspOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEditOrdererOptions successfully`, func() {
				// Construct an instance of the EditOrdererOptions model
				id := "testString"
				editOrdererOptionsModel := testService.NewEditOrdererOptions(id)
				editOrdererOptionsModel.SetID("testString")
				editOrdererOptionsModel.SetClusterName("ordering service 1")
				editOrdererOptionsModel.SetDisplayName("orderer")
				editOrdererOptionsModel.SetApiURL("grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				editOrdererOptionsModel.SetOperationsURL("https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443")
				editOrdererOptionsModel.SetGrpcwpURL("https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443")
				editOrdererOptionsModel.SetMspID("Org1")
				editOrdererOptionsModel.SetConsenterProposalFin(true)
				editOrdererOptionsModel.SetLocation("ibmcloud")
				editOrdererOptionsModel.SetSystemChannelID("testchainid")
				editOrdererOptionsModel.SetTags([]string{"testString"})
				editOrdererOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(editOrdererOptionsModel).ToNot(BeNil())
				Expect(editOrdererOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(editOrdererOptionsModel.ClusterName).To(Equal(core.StringPtr("ordering service 1")))
				Expect(editOrdererOptionsModel.DisplayName).To(Equal(core.StringPtr("orderer")))
				Expect(editOrdererOptionsModel.ApiURL).To(Equal(core.StringPtr("grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")))
				Expect(editOrdererOptionsModel.OperationsURL).To(Equal(core.StringPtr("https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443")))
				Expect(editOrdererOptionsModel.GrpcwpURL).To(Equal(core.StringPtr("https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443")))
				Expect(editOrdererOptionsModel.MspID).To(Equal(core.StringPtr("Org1")))
				Expect(editOrdererOptionsModel.ConsenterProposalFin).To(Equal(core.BoolPtr(true)))
				Expect(editOrdererOptionsModel.Location).To(Equal(core.StringPtr("ibmcloud")))
				Expect(editOrdererOptionsModel.SystemChannelID).To(Equal(core.StringPtr("testchainid")))
				Expect(editOrdererOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(editOrdererOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEditPeerOptions successfully`, func() {
				// Construct an instance of the EditPeerOptions model
				id := "testString"
				editPeerOptionsModel := testService.NewEditPeerOptions(id)
				editPeerOptionsModel.SetID("testString")
				editPeerOptionsModel.SetDisplayName("My Peer")
				editPeerOptionsModel.SetApiURL("grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051")
				editPeerOptionsModel.SetOperationsURL("https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443")
				editPeerOptionsModel.SetGrpcwpURL("https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084")
				editPeerOptionsModel.SetMspID("Org1")
				editPeerOptionsModel.SetLocation("ibmcloud")
				editPeerOptionsModel.SetTags([]string{"testString"})
				editPeerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(editPeerOptionsModel).ToNot(BeNil())
				Expect(editPeerOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(editPeerOptionsModel.DisplayName).To(Equal(core.StringPtr("My Peer")))
				Expect(editPeerOptionsModel.ApiURL).To(Equal(core.StringPtr("grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051")))
				Expect(editPeerOptionsModel.OperationsURL).To(Equal(core.StringPtr("https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443")))
				Expect(editPeerOptionsModel.GrpcwpURL).To(Equal(core.StringPtr("https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084")))
				Expect(editPeerOptionsModel.MspID).To(Equal(core.StringPtr("Org1")))
				Expect(editPeerOptionsModel.Location).To(Equal(core.StringPtr("ibmcloud")))
				Expect(editPeerOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(editPeerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEditSettingsOptions successfully`, func() {
				// Construct an instance of the LoggingSettingsClient model
				loggingSettingsClientModel := new(blockchainv2.LoggingSettingsClient)
				Expect(loggingSettingsClientModel).ToNot(BeNil())
				loggingSettingsClientModel.Enabled = core.BoolPtr(true)
				loggingSettingsClientModel.Level = core.StringPtr("silly")
				loggingSettingsClientModel.UniqueName = core.BoolPtr(false)
				Expect(loggingSettingsClientModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(loggingSettingsClientModel.Level).To(Equal(core.StringPtr("silly")))
				Expect(loggingSettingsClientModel.UniqueName).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the LoggingSettingsServer model
				loggingSettingsServerModel := new(blockchainv2.LoggingSettingsServer)
				Expect(loggingSettingsServerModel).ToNot(BeNil())
				loggingSettingsServerModel.Enabled = core.BoolPtr(true)
				loggingSettingsServerModel.Level = core.StringPtr("silly")
				loggingSettingsServerModel.UniqueName = core.BoolPtr(false)
				Expect(loggingSettingsServerModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(loggingSettingsServerModel.Level).To(Equal(core.StringPtr("silly")))
				Expect(loggingSettingsServerModel.UniqueName).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the EditLogSettingsBody model
				editLogSettingsBodyModel := new(blockchainv2.EditLogSettingsBody)
				Expect(editLogSettingsBodyModel).ToNot(BeNil())
				editLogSettingsBodyModel.Client = loggingSettingsClientModel
				editLogSettingsBodyModel.Server = loggingSettingsServerModel
				Expect(editLogSettingsBodyModel.Client).To(Equal(loggingSettingsClientModel))
				Expect(editLogSettingsBodyModel.Server).To(Equal(loggingSettingsServerModel))

				// Construct an instance of the EditSettingsBodyInactivityTimeouts model
				editSettingsBodyInactivityTimeoutsModel := new(blockchainv2.EditSettingsBodyInactivityTimeouts)
				Expect(editSettingsBodyInactivityTimeoutsModel).ToNot(BeNil())
				editSettingsBodyInactivityTimeoutsModel.Enabled = core.BoolPtr(false)
				editSettingsBodyInactivityTimeoutsModel.MaxIdleTime = core.Float64Ptr(float64(90000))
				Expect(editSettingsBodyInactivityTimeoutsModel.Enabled).To(Equal(core.BoolPtr(false)))
				Expect(editSettingsBodyInactivityTimeoutsModel.MaxIdleTime).To(Equal(core.Float64Ptr(float64(90000))))

				// Construct an instance of the EditSettingsOptions model
				editSettingsOptionsModel := testService.NewEditSettingsOptions()
				editSettingsOptionsModel.SetInactivityTimeouts(editSettingsBodyInactivityTimeoutsModel)
				editSettingsOptionsModel.SetFileLogging(editLogSettingsBodyModel)
				editSettingsOptionsModel.SetMaxReqPerMin(float64(25))
				editSettingsOptionsModel.SetMaxReqPerMinAk(float64(25))
				editSettingsOptionsModel.SetFabricGetBlockTimeoutMs(float64(10000))
				editSettingsOptionsModel.SetFabricInstantiateTimeoutMs(float64(300000))
				editSettingsOptionsModel.SetFabricJoinChannelTimeoutMs(float64(25000))
				editSettingsOptionsModel.SetFabricInstallCcTimeoutMs(float64(300000))
				editSettingsOptionsModel.SetFabricLcInstallCcTimeoutMs(float64(300000))
				editSettingsOptionsModel.SetFabricLcGetCcTimeoutMs(float64(180000))
				editSettingsOptionsModel.SetFabricGeneralTimeoutMs(float64(10000))
				editSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(editSettingsOptionsModel).ToNot(BeNil())
				Expect(editSettingsOptionsModel.InactivityTimeouts).To(Equal(editSettingsBodyInactivityTimeoutsModel))
				Expect(editSettingsOptionsModel.FileLogging).To(Equal(editLogSettingsBodyModel))
				Expect(editSettingsOptionsModel.MaxReqPerMin).To(Equal(core.Float64Ptr(float64(25))))
				Expect(editSettingsOptionsModel.MaxReqPerMinAk).To(Equal(core.Float64Ptr(float64(25))))
				Expect(editSettingsOptionsModel.FabricGetBlockTimeoutMs).To(Equal(core.Float64Ptr(float64(10000))))
				Expect(editSettingsOptionsModel.FabricInstantiateTimeoutMs).To(Equal(core.Float64Ptr(float64(300000))))
				Expect(editSettingsOptionsModel.FabricJoinChannelTimeoutMs).To(Equal(core.Float64Ptr(float64(25000))))
				Expect(editSettingsOptionsModel.FabricInstallCcTimeoutMs).To(Equal(core.Float64Ptr(float64(300000))))
				Expect(editSettingsOptionsModel.FabricLcInstallCcTimeoutMs).To(Equal(core.Float64Ptr(float64(300000))))
				Expect(editSettingsOptionsModel.FabricLcGetCcTimeoutMs).To(Equal(core.Float64Ptr(float64(180000))))
				Expect(editSettingsOptionsModel.FabricGeneralTimeoutMs).To(Equal(core.Float64Ptr(float64(10000))))
				Expect(editSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetComponentByTagOptions successfully`, func() {
				// Construct an instance of the GetComponentByTagOptions model
				tag := "testString"
				getComponentByTagOptionsModel := testService.NewGetComponentByTagOptions(tag)
				getComponentByTagOptionsModel.SetTag("testString")
				getComponentByTagOptionsModel.SetDeploymentAttrs("included")
				getComponentByTagOptionsModel.SetParsedCerts("included")
				getComponentByTagOptionsModel.SetCache("skip")
				getComponentByTagOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getComponentByTagOptionsModel).ToNot(BeNil())
				Expect(getComponentByTagOptionsModel.Tag).To(Equal(core.StringPtr("testString")))
				Expect(getComponentByTagOptionsModel.DeploymentAttrs).To(Equal(core.StringPtr("included")))
				Expect(getComponentByTagOptionsModel.ParsedCerts).To(Equal(core.StringPtr("included")))
				Expect(getComponentByTagOptionsModel.Cache).To(Equal(core.StringPtr("skip")))
				Expect(getComponentByTagOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetComponentOptions successfully`, func() {
				// Construct an instance of the GetComponentOptions model
				id := "testString"
				getComponentOptionsModel := testService.NewGetComponentOptions(id)
				getComponentOptionsModel.SetID("testString")
				getComponentOptionsModel.SetDeploymentAttrs("included")
				getComponentOptionsModel.SetParsedCerts("included")
				getComponentOptionsModel.SetCache("skip")
				getComponentOptionsModel.SetCaAttrs("included")
				getComponentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getComponentOptionsModel).ToNot(BeNil())
				Expect(getComponentOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getComponentOptionsModel.DeploymentAttrs).To(Equal(core.StringPtr("included")))
				Expect(getComponentOptionsModel.ParsedCerts).To(Equal(core.StringPtr("included")))
				Expect(getComponentOptionsModel.Cache).To(Equal(core.StringPtr("skip")))
				Expect(getComponentOptionsModel.CaAttrs).To(Equal(core.StringPtr("included")))
				Expect(getComponentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetComponentsByTypeOptions successfully`, func() {
				// Construct an instance of the GetComponentsByTypeOptions model
				componentType := "fabric-peer"
				getComponentsByTypeOptionsModel := testService.NewGetComponentsByTypeOptions(componentType)
				getComponentsByTypeOptionsModel.SetComponentType("fabric-peer")
				getComponentsByTypeOptionsModel.SetDeploymentAttrs("included")
				getComponentsByTypeOptionsModel.SetParsedCerts("included")
				getComponentsByTypeOptionsModel.SetCache("skip")
				getComponentsByTypeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getComponentsByTypeOptionsModel).ToNot(BeNil())
				Expect(getComponentsByTypeOptionsModel.ComponentType).To(Equal(core.StringPtr("fabric-peer")))
				Expect(getComponentsByTypeOptionsModel.DeploymentAttrs).To(Equal(core.StringPtr("included")))
				Expect(getComponentsByTypeOptionsModel.ParsedCerts).To(Equal(core.StringPtr("included")))
				Expect(getComponentsByTypeOptionsModel.Cache).To(Equal(core.StringPtr("skip")))
				Expect(getComponentsByTypeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetFabVersionsOptions successfully`, func() {
				// Construct an instance of the GetFabVersionsOptions model
				getFabVersionsOptionsModel := testService.NewGetFabVersionsOptions()
				getFabVersionsOptionsModel.SetCache("skip")
				getFabVersionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getFabVersionsOptionsModel).ToNot(BeNil())
				Expect(getFabVersionsOptionsModel.Cache).To(Equal(core.StringPtr("skip")))
				Expect(getFabVersionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetHealthOptions successfully`, func() {
				// Construct an instance of the GetHealthOptions model
				getHealthOptionsModel := testService.NewGetHealthOptions()
				getHealthOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getHealthOptionsModel).ToNot(BeNil())
				Expect(getHealthOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetMspCertificateOptions successfully`, func() {
				// Construct an instance of the GetMspCertificateOptions model
				mspID := "testString"
				getMspCertificateOptionsModel := testService.NewGetMspCertificateOptions(mspID)
				getMspCertificateOptionsModel.SetMspID("testString")
				getMspCertificateOptionsModel.SetCache("skip")
				getMspCertificateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getMspCertificateOptionsModel).ToNot(BeNil())
				Expect(getMspCertificateOptionsModel.MspID).To(Equal(core.StringPtr("testString")))
				Expect(getMspCertificateOptionsModel.Cache).To(Equal(core.StringPtr("skip")))
				Expect(getMspCertificateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPostmanOptions successfully`, func() {
				// Construct an instance of the GetPostmanOptions model
				authType := "bearer"
				getPostmanOptionsModel := testService.NewGetPostmanOptions(authType)
				getPostmanOptionsModel.SetAuthType("bearer")
				getPostmanOptionsModel.SetToken("testString")
				getPostmanOptionsModel.SetApiKey("testString")
				getPostmanOptionsModel.SetUsername("admin")
				getPostmanOptionsModel.SetPassword("password")
				getPostmanOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPostmanOptionsModel).ToNot(BeNil())
				Expect(getPostmanOptionsModel.AuthType).To(Equal(core.StringPtr("bearer")))
				Expect(getPostmanOptionsModel.Token).To(Equal(core.StringPtr("testString")))
				Expect(getPostmanOptionsModel.ApiKey).To(Equal(core.StringPtr("testString")))
				Expect(getPostmanOptionsModel.Username).To(Equal(core.StringPtr("admin")))
				Expect(getPostmanOptionsModel.Password).To(Equal(core.StringPtr("password")))
				Expect(getPostmanOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSettingsOptions successfully`, func() {
				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := testService.NewGetSettingsOptions()
				getSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSettingsOptionsModel).ToNot(BeNil())
				Expect(getSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSwaggerOptions successfully`, func() {
				// Construct an instance of the GetSwaggerOptions model
				getSwaggerOptionsModel := testService.NewGetSwaggerOptions()
				getSwaggerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSwaggerOptionsModel).ToNot(BeNil())
				Expect(getSwaggerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewImportCaOptions successfully`, func() {
				// Construct an instance of the ImportCaOptions model
				importCaOptionsDisplayName := "Sample CA"
				importCaOptionsApiURL := "https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:7054"
				importCaOptionsCaName := "org1CA"
				importCaOptionsTlscaName := "org1CA"
				importCaOptionsTlsCert := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
				importCaOptionsModel := testService.NewImportCaOptions(importCaOptionsDisplayName, importCaOptionsApiURL, importCaOptionsCaName, importCaOptionsTlscaName, importCaOptionsTlsCert)
				importCaOptionsModel.SetDisplayName("Sample CA")
				importCaOptionsModel.SetApiURL("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:7054")
				importCaOptionsModel.SetCaName("org1CA")
				importCaOptionsModel.SetTlscaName("org1CA")
				importCaOptionsModel.SetTlsCert("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				importCaOptionsModel.SetLocation("ibmcloud")
				importCaOptionsModel.SetOperationsURL("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443")
				importCaOptionsModel.SetTags([]string{"testString"})
				importCaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(importCaOptionsModel).ToNot(BeNil())
				Expect(importCaOptionsModel.DisplayName).To(Equal(core.StringPtr("Sample CA")))
				Expect(importCaOptionsModel.ApiURL).To(Equal(core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:7054")))
				Expect(importCaOptionsModel.CaName).To(Equal(core.StringPtr("org1CA")))
				Expect(importCaOptionsModel.TlscaName).To(Equal(core.StringPtr("org1CA")))
				Expect(importCaOptionsModel.TlsCert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(importCaOptionsModel.Location).To(Equal(core.StringPtr("ibmcloud")))
				Expect(importCaOptionsModel.OperationsURL).To(Equal(core.StringPtr("https://n3a3ec3-myca.ibp.us-south.containers.appdomain.cloud:9443")))
				Expect(importCaOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(importCaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewImportMspOptions successfully`, func() {
				// Construct an instance of the ImportMspOptions model
				importMspOptionsMspID := "Org1"
				importMspOptionsDisplayName := "My Peer"
				importMspOptionsRootCerts := []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				importMspOptionsModel := testService.NewImportMspOptions(importMspOptionsMspID, importMspOptionsDisplayName, importMspOptionsRootCerts)
				importMspOptionsModel.SetMspID("Org1")
				importMspOptionsModel.SetDisplayName("My Peer")
				importMspOptionsModel.SetRootCerts([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="})
				importMspOptionsModel.SetIntermediateCerts([]string{"testString"})
				importMspOptionsModel.SetAdmins([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="})
				importMspOptionsModel.SetTlsRootCerts([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="})
				importMspOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(importMspOptionsModel).ToNot(BeNil())
				Expect(importMspOptionsModel.MspID).To(Equal(core.StringPtr("Org1")))
				Expect(importMspOptionsModel.DisplayName).To(Equal(core.StringPtr("My Peer")))
				Expect(importMspOptionsModel.RootCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(importMspOptionsModel.IntermediateCerts).To(Equal([]string{"testString"}))
				Expect(importMspOptionsModel.Admins).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(importMspOptionsModel.TlsRootCerts).To(Equal([]string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}))
				Expect(importMspOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewImportOrdererOptions successfully`, func() {
				// Construct an instance of the ImportOrdererOptions model
				importOrdererOptionsClusterName := "ordering service 1"
				importOrdererOptionsDisplayName := "orderer"
				importOrdererOptionsMspID := "Org1"
				importOrdererOptionsGrpcwpURL := "https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443"
				importOrdererOptionsTlsCaRootCert := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkRpZmZlcmVudCBkYXRhIGhlcmUgaWYgdGhpcyB3YXMgcmVhbAotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg=="
				importOrdererOptionsModel := testService.NewImportOrdererOptions(importOrdererOptionsClusterName, importOrdererOptionsDisplayName, importOrdererOptionsMspID, importOrdererOptionsGrpcwpURL, importOrdererOptionsTlsCaRootCert)
				importOrdererOptionsModel.SetClusterName("ordering service 1")
				importOrdererOptionsModel.SetDisplayName("orderer")
				importOrdererOptionsModel.SetMspID("Org1")
				importOrdererOptionsModel.SetGrpcwpURL("https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443")
				importOrdererOptionsModel.SetTlsCaRootCert("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkRpZmZlcmVudCBkYXRhIGhlcmUgaWYgdGhpcyB3YXMgcmVhbAotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==")
				importOrdererOptionsModel.SetLocation("ibmcloud")
				importOrdererOptionsModel.SetApiURL("grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				importOrdererOptionsModel.SetOperationsURL("https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443")
				importOrdererOptionsModel.SetSystemChannelID("testchainid")
				importOrdererOptionsModel.SetTags([]string{"testString"})
				importOrdererOptionsModel.SetTlsCert("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				importOrdererOptionsModel.SetServerTlsCert("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				importOrdererOptionsModel.SetClientTlsCert("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				importOrdererOptionsModel.SetClusterID("testString")
				importOrdererOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(importOrdererOptionsModel).ToNot(BeNil())
				Expect(importOrdererOptionsModel.ClusterName).To(Equal(core.StringPtr("ordering service 1")))
				Expect(importOrdererOptionsModel.DisplayName).To(Equal(core.StringPtr("orderer")))
				Expect(importOrdererOptionsModel.MspID).To(Equal(core.StringPtr("Org1")))
				Expect(importOrdererOptionsModel.GrpcwpURL).To(Equal(core.StringPtr("https://n3a3ec3-myorderer-proxy.ibp.us-south.containers.appdomain.cloud:443")))
				Expect(importOrdererOptionsModel.TlsCaRootCert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkRpZmZlcmVudCBkYXRhIGhlcmUgaWYgdGhpcyB3YXMgcmVhbAotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==")))
				Expect(importOrdererOptionsModel.Location).To(Equal(core.StringPtr("ibmcloud")))
				Expect(importOrdererOptionsModel.ApiURL).To(Equal(core.StringPtr("grpcs://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")))
				Expect(importOrdererOptionsModel.OperationsURL).To(Equal(core.StringPtr("https://n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:8443")))
				Expect(importOrdererOptionsModel.SystemChannelID).To(Equal(core.StringPtr("testchainid")))
				Expect(importOrdererOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(importOrdererOptionsModel.TlsCert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(importOrdererOptionsModel.ServerTlsCert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(importOrdererOptionsModel.ClientTlsCert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(importOrdererOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(importOrdererOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewImportPeerOptions successfully`, func() {
				// Construct an instance of the ImportPeerOptions model
				importPeerOptionsDisplayName := "My Peer"
				importPeerOptionsMspID := "Org1"
				importPeerOptionsGrpcwpURL := "https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084"
				importPeerOptionsTlsCaRootCert := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkRpZmZlcmVudCBkYXRhIGhlcmUgaWYgdGhpcyB3YXMgcmVhbAotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg=="
				importPeerOptionsModel := testService.NewImportPeerOptions(importPeerOptionsDisplayName, importPeerOptionsMspID, importPeerOptionsGrpcwpURL, importPeerOptionsTlsCaRootCert)
				importPeerOptionsModel.SetDisplayName("My Peer")
				importPeerOptionsModel.SetMspID("Org1")
				importPeerOptionsModel.SetGrpcwpURL("https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084")
				importPeerOptionsModel.SetTlsCaRootCert("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkRpZmZlcmVudCBkYXRhIGhlcmUgaWYgdGhpcyB3YXMgcmVhbAotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==")
				importPeerOptionsModel.SetLocation("ibmcloud")
				importPeerOptionsModel.SetApiURL("grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051")
				importPeerOptionsModel.SetOperationsURL("https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443")
				importPeerOptionsModel.SetTags([]string{"testString"})
				importPeerOptionsModel.SetTlsCert("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")
				importPeerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(importPeerOptionsModel).ToNot(BeNil())
				Expect(importPeerOptionsModel.DisplayName).To(Equal(core.StringPtr("My Peer")))
				Expect(importPeerOptionsModel.MspID).To(Equal(core.StringPtr("Org1")))
				Expect(importPeerOptionsModel.GrpcwpURL).To(Equal(core.StringPtr("https://n3a3ec3-mypeer-proxy.ibp.us-south.containers.appdomain.cloud:8084")))
				Expect(importPeerOptionsModel.TlsCaRootCert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkRpZmZlcmVudCBkYXRhIGhlcmUgaWYgdGhpcyB3YXMgcmVhbAotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==")))
				Expect(importPeerOptionsModel.Location).To(Equal(core.StringPtr("ibmcloud")))
				Expect(importPeerOptionsModel.ApiURL).To(Equal(core.StringPtr("grpcs://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:7051")))
				Expect(importPeerOptionsModel.OperationsURL).To(Equal(core.StringPtr("https://n3a3ec3-mypeer.ibp.us-south.containers.appdomain.cloud:9443")))
				Expect(importPeerOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(importPeerOptionsModel.TlsCert).To(Equal(core.StringPtr("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=")))
				Expect(importPeerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListComponentsOptions successfully`, func() {
				// Construct an instance of the ListComponentsOptions model
				listComponentsOptionsModel := testService.NewListComponentsOptions()
				listComponentsOptionsModel.SetDeploymentAttrs("included")
				listComponentsOptionsModel.SetParsedCerts("included")
				listComponentsOptionsModel.SetCache("skip")
				listComponentsOptionsModel.SetCaAttrs("included")
				listComponentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listComponentsOptionsModel).ToNot(BeNil())
				Expect(listComponentsOptionsModel.DeploymentAttrs).To(Equal(core.StringPtr("included")))
				Expect(listComponentsOptionsModel.ParsedCerts).To(Equal(core.StringPtr("included")))
				Expect(listComponentsOptionsModel.Cache).To(Equal(core.StringPtr("skip")))
				Expect(listComponentsOptionsModel.CaAttrs).To(Equal(core.StringPtr("included")))
				Expect(listComponentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListNotificationsOptions successfully`, func() {
				// Construct an instance of the ListNotificationsOptions model
				listNotificationsOptionsModel := testService.NewListNotificationsOptions()
				listNotificationsOptionsModel.SetLimit(float64(1))
				listNotificationsOptionsModel.SetSkip(float64(1))
				listNotificationsOptionsModel.SetComponentID("MyPeer")
				listNotificationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listNotificationsOptionsModel).ToNot(BeNil())
				Expect(listNotificationsOptionsModel.Limit).To(Equal(core.Float64Ptr(float64(1))))
				Expect(listNotificationsOptionsModel.Skip).To(Equal(core.Float64Ptr(float64(1))))
				Expect(listNotificationsOptionsModel.ComponentID).To(Equal(core.StringPtr("MyPeer")))
				Expect(listNotificationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewMetrics successfully`, func() {
				provider := "prometheus"
				model, err := testService.NewMetrics(provider)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewMetricsStatsd successfully`, func() {
				network := "udp"
				address := "127.0.0.1:8125"
				writeInterval := "10s"
				prefix := "server"
				model, err := testService.NewMetricsStatsd(network, address, writeInterval, prefix)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewMspConfigData successfully`, func() {
				keystore := "testString"
				signcerts := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
				cacerts := []string{"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCkNlcnQgZGF0YSB3b3VsZCBiZSBoZXJlIGlmIHRoaXMgd2FzIHJlYWwKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="}
				model, err := testService.NewMspConfigData(keystore, signcerts, cacerts)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRemoveComponentOptions successfully`, func() {
				// Construct an instance of the RemoveComponentOptions model
				id := "testString"
				removeComponentOptionsModel := testService.NewRemoveComponentOptions(id)
				removeComponentOptionsModel.SetID("testString")
				removeComponentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(removeComponentOptionsModel).ToNot(BeNil())
				Expect(removeComponentOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(removeComponentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRemoveComponentsByTagOptions successfully`, func() {
				// Construct an instance of the RemoveComponentsByTagOptions model
				tag := "testString"
				removeComponentsByTagOptionsModel := testService.NewRemoveComponentsByTagOptions(tag)
				removeComponentsByTagOptionsModel.SetTag("testString")
				removeComponentsByTagOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(removeComponentsByTagOptionsModel).ToNot(BeNil())
				Expect(removeComponentsByTagOptionsModel.Tag).To(Equal(core.StringPtr("testString")))
				Expect(removeComponentsByTagOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewResourceObject successfully`, func() {
				var requests *blockchainv2.ResourceRequests = nil
				_, err := testService.NewResourceObject(requests)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewResourceObjectCouchDb successfully`, func() {
				var requests *blockchainv2.ResourceRequests = nil
				_, err := testService.NewResourceObjectCouchDb(requests)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewResourceObjectFabV1 successfully`, func() {
				var requests *blockchainv2.ResourceRequests = nil
				_, err := testService.NewResourceObjectFabV1(requests)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewResourceObjectFabV2 successfully`, func() {
				var requests *blockchainv2.ResourceRequests = nil
				_, err := testService.NewResourceObjectFabV2(requests)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewRestartOptions successfully`, func() {
				// Construct an instance of the RestartOptions model
				restartOptionsModel := testService.NewRestartOptions()
				restartOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(restartOptionsModel).ToNot(BeNil())
				Expect(restartOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSubmitBlockOptions successfully`, func() {
				// Construct an instance of the SubmitBlockOptions model
				id := "testString"
				submitBlockOptionsModel := testService.NewSubmitBlockOptions(id)
				submitBlockOptionsModel.SetID("testString")
				submitBlockOptionsModel.SetB64Block("bWFzc2l2ZSBiaW5hcnkgb2YgYSBjb25maWcgYmxvY2sgd291bGQgYmUgaGVyZSBpZiB0aGlzIHdhcyByZWFs")
				submitBlockOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(submitBlockOptionsModel).ToNot(BeNil())
				Expect(submitBlockOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(submitBlockOptionsModel.B64Block).To(Equal(core.StringPtr("bWFzc2l2ZSBiaW5hcnkgb2YgYSBjb25maWcgYmxvY2sgd291bGQgYmUgaGVyZSBpZiB0aGlzIHdhcyByZWFs")))
				Expect(submitBlockOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateCaBodyConfigOverride successfully`, func() {
				var ca *blockchainv2.ConfigCAUpdate = nil
				_, err := testService.NewUpdateCaBodyConfigOverride(ca)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewUpdateCaBodyResources successfully`, func() {
				var ca *blockchainv2.ResourceObject = nil
				_, err := testService.NewUpdateCaBodyResources(ca)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewUpdateCaOptions successfully`, func() {
				// Construct an instance of the ConfigCADbTlsClient model
				configCaDbTlsClientModel := new(blockchainv2.ConfigCADbTlsClient)
				Expect(configCaDbTlsClientModel).ToNot(BeNil())
				configCaDbTlsClientModel.Certfile = core.StringPtr("testString")
				configCaDbTlsClientModel.Keyfile = core.StringPtr("testString")
				Expect(configCaDbTlsClientModel.Certfile).To(Equal(core.StringPtr("testString")))
				Expect(configCaDbTlsClientModel.Keyfile).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigCAIntermediateTlsClient model
				configCaIntermediateTlsClientModel := new(blockchainv2.ConfigCAIntermediateTlsClient)
				Expect(configCaIntermediateTlsClientModel).ToNot(BeNil())
				configCaIntermediateTlsClientModel.Certfile = core.StringPtr("testString")
				configCaIntermediateTlsClientModel.Keyfile = core.StringPtr("testString")
				Expect(configCaIntermediateTlsClientModel.Certfile).To(Equal(core.StringPtr("testString")))
				Expect(configCaIntermediateTlsClientModel.Keyfile).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the IdentityAttrs model
				identityAttrsModel := new(blockchainv2.IdentityAttrs)
				Expect(identityAttrsModel).ToNot(BeNil())
				identityAttrsModel.HfRegistrarRoles = core.StringPtr("*")
				identityAttrsModel.HfRegistrarDelegateRoles = core.StringPtr("*")
				identityAttrsModel.HfRevoker = core.BoolPtr(true)
				identityAttrsModel.HfIntermediateCA = core.BoolPtr(true)
				identityAttrsModel.HfGenCRL = core.BoolPtr(true)
				identityAttrsModel.HfRegistrarAttributes = core.StringPtr("*")
				identityAttrsModel.HfAffiliationMgr = core.BoolPtr(true)
				Expect(identityAttrsModel.HfRegistrarRoles).To(Equal(core.StringPtr("*")))
				Expect(identityAttrsModel.HfRegistrarDelegateRoles).To(Equal(core.StringPtr("*")))
				Expect(identityAttrsModel.HfRevoker).To(Equal(core.BoolPtr(true)))
				Expect(identityAttrsModel.HfIntermediateCA).To(Equal(core.BoolPtr(true)))
				Expect(identityAttrsModel.HfGenCRL).To(Equal(core.BoolPtr(true)))
				Expect(identityAttrsModel.HfRegistrarAttributes).To(Equal(core.StringPtr("*")))
				Expect(identityAttrsModel.HfAffiliationMgr).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the BccspPKCS11 model
				bccspPkcS11Model := new(blockchainv2.BccspPKCS11)
				Expect(bccspPkcS11Model).ToNot(BeNil())
				bccspPkcS11Model.Label = core.StringPtr("testString")
				bccspPkcS11Model.Pin = core.StringPtr("testString")
				bccspPkcS11Model.Hash = core.StringPtr("SHA2")
				bccspPkcS11Model.Security = core.Float64Ptr(float64(256))
				Expect(bccspPkcS11Model.Label).To(Equal(core.StringPtr("testString")))
				Expect(bccspPkcS11Model.Pin).To(Equal(core.StringPtr("testString")))
				Expect(bccspPkcS11Model.Hash).To(Equal(core.StringPtr("SHA2")))
				Expect(bccspPkcS11Model.Security).To(Equal(core.Float64Ptr(float64(256))))

				// Construct an instance of the BccspSW model
				bccspSwModel := new(blockchainv2.BccspSW)
				Expect(bccspSwModel).ToNot(BeNil())
				bccspSwModel.Hash = core.StringPtr("SHA2")
				bccspSwModel.Security = core.Float64Ptr(float64(256))
				Expect(bccspSwModel.Hash).To(Equal(core.StringPtr("SHA2")))
				Expect(bccspSwModel.Security).To(Equal(core.Float64Ptr(float64(256))))

				// Construct an instance of the ConfigCACfgIdentities model
				configCaCfgIdentitiesModel := new(blockchainv2.ConfigCACfgIdentities)
				Expect(configCaCfgIdentitiesModel).ToNot(BeNil())
				configCaCfgIdentitiesModel.Passwordattempts = core.Float64Ptr(float64(10))
				configCaCfgIdentitiesModel.Allowremove = core.BoolPtr(false)
				Expect(configCaCfgIdentitiesModel.Passwordattempts).To(Equal(core.Float64Ptr(float64(10))))
				Expect(configCaCfgIdentitiesModel.Allowremove).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the ConfigCACsrCa model
				configCaCsrCaModel := new(blockchainv2.ConfigCACsrCa)
				Expect(configCaCsrCaModel).ToNot(BeNil())
				configCaCsrCaModel.Expiry = core.StringPtr("131400h")
				configCaCsrCaModel.Pathlength = core.Float64Ptr(float64(0))
				Expect(configCaCsrCaModel.Expiry).To(Equal(core.StringPtr("131400h")))
				Expect(configCaCsrCaModel.Pathlength).To(Equal(core.Float64Ptr(float64(0))))

				// Construct an instance of the ConfigCACsrKeyrequest model
				configCaCsrKeyrequestModel := new(blockchainv2.ConfigCACsrKeyrequest)
				Expect(configCaCsrKeyrequestModel).ToNot(BeNil())
				configCaCsrKeyrequestModel.Algo = core.StringPtr("ecdsa")
				configCaCsrKeyrequestModel.Size = core.Float64Ptr(float64(256))
				Expect(configCaCsrKeyrequestModel.Algo).To(Equal(core.StringPtr("ecdsa")))
				Expect(configCaCsrKeyrequestModel.Size).To(Equal(core.Float64Ptr(float64(256))))

				// Construct an instance of the ConfigCACsrNamesItem model
				configCaCsrNamesItemModel := new(blockchainv2.ConfigCACsrNamesItem)
				Expect(configCaCsrNamesItemModel).ToNot(BeNil())
				configCaCsrNamesItemModel.C = core.StringPtr("US")
				configCaCsrNamesItemModel.ST = core.StringPtr("North Carolina")
				configCaCsrNamesItemModel.L = core.StringPtr("Raleigh")
				configCaCsrNamesItemModel.O = core.StringPtr("Hyperledger")
				configCaCsrNamesItemModel.OU = core.StringPtr("Fabric")
				Expect(configCaCsrNamesItemModel.C).To(Equal(core.StringPtr("US")))
				Expect(configCaCsrNamesItemModel.ST).To(Equal(core.StringPtr("North Carolina")))
				Expect(configCaCsrNamesItemModel.L).To(Equal(core.StringPtr("Raleigh")))
				Expect(configCaCsrNamesItemModel.O).To(Equal(core.StringPtr("Hyperledger")))
				Expect(configCaCsrNamesItemModel.OU).To(Equal(core.StringPtr("Fabric")))

				// Construct an instance of the ConfigCADbTls model
				configCaDbTlsModel := new(blockchainv2.ConfigCADbTls)
				Expect(configCaDbTlsModel).ToNot(BeNil())
				configCaDbTlsModel.Certfiles = []string{"testString"}
				configCaDbTlsModel.Client = configCaDbTlsClientModel
				configCaDbTlsModel.Enabled = core.BoolPtr(false)
				Expect(configCaDbTlsModel.Certfiles).To(Equal([]string{"testString"}))
				Expect(configCaDbTlsModel.Client).To(Equal(configCaDbTlsClientModel))
				Expect(configCaDbTlsModel.Enabled).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the ConfigCAIntermediateEnrollment model
				configCaIntermediateEnrollmentModel := new(blockchainv2.ConfigCAIntermediateEnrollment)
				Expect(configCaIntermediateEnrollmentModel).ToNot(BeNil())
				configCaIntermediateEnrollmentModel.Hosts = core.StringPtr("localhost")
				configCaIntermediateEnrollmentModel.Profile = core.StringPtr("testString")
				configCaIntermediateEnrollmentModel.Label = core.StringPtr("testString")
				Expect(configCaIntermediateEnrollmentModel.Hosts).To(Equal(core.StringPtr("localhost")))
				Expect(configCaIntermediateEnrollmentModel.Profile).To(Equal(core.StringPtr("testString")))
				Expect(configCaIntermediateEnrollmentModel.Label).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigCAIntermediateParentserver model
				configCaIntermediateParentserverModel := new(blockchainv2.ConfigCAIntermediateParentserver)
				Expect(configCaIntermediateParentserverModel).ToNot(BeNil())
				configCaIntermediateParentserverModel.URL = core.StringPtr("testString")
				configCaIntermediateParentserverModel.Caname = core.StringPtr("testString")
				Expect(configCaIntermediateParentserverModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(configCaIntermediateParentserverModel.Caname).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigCAIntermediateTls model
				configCaIntermediateTlsModel := new(blockchainv2.ConfigCAIntermediateTls)
				Expect(configCaIntermediateTlsModel).ToNot(BeNil())
				configCaIntermediateTlsModel.Certfiles = []string{"testString"}
				configCaIntermediateTlsModel.Client = configCaIntermediateTlsClientModel
				Expect(configCaIntermediateTlsModel.Certfiles).To(Equal([]string{"testString"}))
				Expect(configCaIntermediateTlsModel.Client).To(Equal(configCaIntermediateTlsClientModel))

				// Construct an instance of the ConfigCARegistryIdentitiesItem model
				configCaRegistryIdentitiesItemModel := new(blockchainv2.ConfigCARegistryIdentitiesItem)
				Expect(configCaRegistryIdentitiesItemModel).ToNot(BeNil())
				configCaRegistryIdentitiesItemModel.Name = core.StringPtr("admin")
				configCaRegistryIdentitiesItemModel.Pass = core.StringPtr("password")
				configCaRegistryIdentitiesItemModel.Type = core.StringPtr("client")
				configCaRegistryIdentitiesItemModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryIdentitiesItemModel.Affiliation = core.StringPtr("testString")
				configCaRegistryIdentitiesItemModel.Attrs = identityAttrsModel
				Expect(configCaRegistryIdentitiesItemModel.Name).To(Equal(core.StringPtr("admin")))
				Expect(configCaRegistryIdentitiesItemModel.Pass).To(Equal(core.StringPtr("password")))
				Expect(configCaRegistryIdentitiesItemModel.Type).To(Equal(core.StringPtr("client")))
				Expect(configCaRegistryIdentitiesItemModel.Maxenrollments).To(Equal(core.Float64Ptr(float64(-1))))
				Expect(configCaRegistryIdentitiesItemModel.Affiliation).To(Equal(core.StringPtr("testString")))
				Expect(configCaRegistryIdentitiesItemModel.Attrs).To(Equal(identityAttrsModel))

				// Construct an instance of the ConfigCATlsClientauth model
				configCaTlsClientauthModel := new(blockchainv2.ConfigCATlsClientauth)
				Expect(configCaTlsClientauthModel).ToNot(BeNil())
				configCaTlsClientauthModel.Type = core.StringPtr("noclientcert")
				configCaTlsClientauthModel.Certfiles = []string{"testString"}
				Expect(configCaTlsClientauthModel.Type).To(Equal(core.StringPtr("noclientcert")))
				Expect(configCaTlsClientauthModel.Certfiles).To(Equal([]string{"testString"}))

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv2.MetricsStatsd)
				Expect(metricsStatsdModel).ToNot(BeNil())
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")
				Expect(metricsStatsdModel.Network).To(Equal(core.StringPtr("udp")))
				Expect(metricsStatsdModel.Address).To(Equal(core.StringPtr("127.0.0.1:8125")))
				Expect(metricsStatsdModel.WriteInterval).To(Equal(core.StringPtr("10s")))
				Expect(metricsStatsdModel.Prefix).To(Equal(core.StringPtr("server")))

				// Construct an instance of the Bccsp model
				bccspModel := new(blockchainv2.Bccsp)
				Expect(bccspModel).ToNot(BeNil())
				bccspModel.Default = core.StringPtr("SW")
				bccspModel.SW = bccspSwModel
				bccspModel.PKCS11 = bccspPkcS11Model
				Expect(bccspModel.Default).To(Equal(core.StringPtr("SW")))
				Expect(bccspModel.SW).To(Equal(bccspSwModel))
				Expect(bccspModel.PKCS11).To(Equal(bccspPkcS11Model))

				// Construct an instance of the ConfigCAAffiliations model
				configCaAffiliationsModel := new(blockchainv2.ConfigCAAffiliations)
				Expect(configCaAffiliationsModel).ToNot(BeNil())
				configCaAffiliationsModel.Org1 = []string{"department1"}
				configCaAffiliationsModel.Org2 = []string{"department1"}
				configCaAffiliationsModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(configCaAffiliationsModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))
				Expect(configCaAffiliationsModel.GetProperties()).ToNot(BeEmpty())
				Expect(configCaAffiliationsModel.Org1).To(Equal([]string{"department1"}))
				Expect(configCaAffiliationsModel.Org2).To(Equal([]string{"department1"}))

				// Construct an instance of the ConfigCACa model
				configCaCaModel := new(blockchainv2.ConfigCACa)
				Expect(configCaCaModel).ToNot(BeNil())
				configCaCaModel.Keyfile = core.StringPtr("testString")
				configCaCaModel.Certfile = core.StringPtr("testString")
				configCaCaModel.Chainfile = core.StringPtr("testString")
				Expect(configCaCaModel.Keyfile).To(Equal(core.StringPtr("testString")))
				Expect(configCaCaModel.Certfile).To(Equal(core.StringPtr("testString")))
				Expect(configCaCaModel.Chainfile).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigCACfg model
				configCaCfgModel := new(blockchainv2.ConfigCACfg)
				Expect(configCaCfgModel).ToNot(BeNil())
				configCaCfgModel.Identities = configCaCfgIdentitiesModel
				Expect(configCaCfgModel.Identities).To(Equal(configCaCfgIdentitiesModel))

				// Construct an instance of the ConfigCACors model
				configCaCorsModel := new(blockchainv2.ConfigCACors)
				Expect(configCaCorsModel).ToNot(BeNil())
				configCaCorsModel.Enabled = core.BoolPtr(true)
				configCaCorsModel.Origins = []string{"*"}
				Expect(configCaCorsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(configCaCorsModel.Origins).To(Equal([]string{"*"}))

				// Construct an instance of the ConfigCACrl model
				configCaCrlModel := new(blockchainv2.ConfigCACrl)
				Expect(configCaCrlModel).ToNot(BeNil())
				configCaCrlModel.Expiry = core.StringPtr("24h")
				Expect(configCaCrlModel.Expiry).To(Equal(core.StringPtr("24h")))

				// Construct an instance of the ConfigCACsr model
				configCaCsrModel := new(blockchainv2.ConfigCACsr)
				Expect(configCaCsrModel).ToNot(BeNil())
				configCaCsrModel.Cn = core.StringPtr("ca")
				configCaCsrModel.Keyrequest = configCaCsrKeyrequestModel
				configCaCsrModel.Names = []blockchainv2.ConfigCACsrNamesItem{*configCaCsrNamesItemModel}
				configCaCsrModel.Hosts = []string{"localhost"}
				configCaCsrModel.Ca = configCaCsrCaModel
				Expect(configCaCsrModel.Cn).To(Equal(core.StringPtr("ca")))
				Expect(configCaCsrModel.Keyrequest).To(Equal(configCaCsrKeyrequestModel))
				Expect(configCaCsrModel.Names).To(Equal([]blockchainv2.ConfigCACsrNamesItem{*configCaCsrNamesItemModel}))
				Expect(configCaCsrModel.Hosts).To(Equal([]string{"localhost"}))
				Expect(configCaCsrModel.Ca).To(Equal(configCaCsrCaModel))

				// Construct an instance of the ConfigCADb model
				configCaDbModel := new(blockchainv2.ConfigCADb)
				Expect(configCaDbModel).ToNot(BeNil())
				configCaDbModel.Type = core.StringPtr("postgres")
				configCaDbModel.Datasource = core.StringPtr("host=fake.databases.appdomain.cloud port=31941 user=ibm_cloud password=password dbname=ibmclouddb sslmode=verify-full")
				configCaDbModel.Tls = configCaDbTlsModel
				Expect(configCaDbModel.Type).To(Equal(core.StringPtr("postgres")))
				Expect(configCaDbModel.Datasource).To(Equal(core.StringPtr("host=fake.databases.appdomain.cloud port=31941 user=ibm_cloud password=password dbname=ibmclouddb sslmode=verify-full")))
				Expect(configCaDbModel.Tls).To(Equal(configCaDbTlsModel))

				// Construct an instance of the ConfigCAIdemix model
				configCaIdemixModel := new(blockchainv2.ConfigCAIdemix)
				Expect(configCaIdemixModel).ToNot(BeNil())
				configCaIdemixModel.Rhpoolsize = core.Float64Ptr(float64(100))
				configCaIdemixModel.Nonceexpiration = core.StringPtr("15s")
				configCaIdemixModel.Noncesweepinterval = core.StringPtr("15m")
				Expect(configCaIdemixModel.Rhpoolsize).To(Equal(core.Float64Ptr(float64(100))))
				Expect(configCaIdemixModel.Nonceexpiration).To(Equal(core.StringPtr("15s")))
				Expect(configCaIdemixModel.Noncesweepinterval).To(Equal(core.StringPtr("15m")))

				// Construct an instance of the ConfigCAIntermediate model
				configCaIntermediateModel := new(blockchainv2.ConfigCAIntermediate)
				Expect(configCaIntermediateModel).ToNot(BeNil())
				configCaIntermediateModel.Parentserver = configCaIntermediateParentserverModel
				configCaIntermediateModel.Enrollment = configCaIntermediateEnrollmentModel
				configCaIntermediateModel.Tls = configCaIntermediateTlsModel
				Expect(configCaIntermediateModel.Parentserver).To(Equal(configCaIntermediateParentserverModel))
				Expect(configCaIntermediateModel.Enrollment).To(Equal(configCaIntermediateEnrollmentModel))
				Expect(configCaIntermediateModel.Tls).To(Equal(configCaIntermediateTlsModel))

				// Construct an instance of the ConfigCARegistry model
				configCaRegistryModel := new(blockchainv2.ConfigCARegistry)
				Expect(configCaRegistryModel).ToNot(BeNil())
				configCaRegistryModel.Maxenrollments = core.Float64Ptr(float64(-1))
				configCaRegistryModel.Identities = []blockchainv2.ConfigCARegistryIdentitiesItem{*configCaRegistryIdentitiesItemModel}
				Expect(configCaRegistryModel.Maxenrollments).To(Equal(core.Float64Ptr(float64(-1))))
				Expect(configCaRegistryModel.Identities).To(Equal([]blockchainv2.ConfigCARegistryIdentitiesItem{*configCaRegistryIdentitiesItemModel}))

				// Construct an instance of the ConfigCATls model
				configCaTlsModel := new(blockchainv2.ConfigCATls)
				Expect(configCaTlsModel).ToNot(BeNil())
				configCaTlsModel.Keyfile = core.StringPtr("testString")
				configCaTlsModel.Certfile = core.StringPtr("testString")
				configCaTlsModel.Clientauth = configCaTlsClientauthModel
				Expect(configCaTlsModel.Keyfile).To(Equal(core.StringPtr("testString")))
				Expect(configCaTlsModel.Certfile).To(Equal(core.StringPtr("testString")))
				Expect(configCaTlsModel.Clientauth).To(Equal(configCaTlsClientauthModel))

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv2.Metrics)
				Expect(metricsModel).ToNot(BeNil())
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel
				Expect(metricsModel.Provider).To(Equal(core.StringPtr("prometheus")))
				Expect(metricsModel.Statsd).To(Equal(metricsStatsdModel))

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv2.ResourceLimits)
				Expect(resourceLimitsModel).ToNot(BeNil())
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")
				Expect(resourceLimitsModel.Cpu).To(Equal(core.StringPtr("100m")))
				Expect(resourceLimitsModel.Memory).To(Equal(core.StringPtr("256MiB")))

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv2.ResourceRequests)
				Expect(resourceRequestsModel).ToNot(BeNil())
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")
				Expect(resourceRequestsModel.Cpu).To(Equal(core.StringPtr("100m")))
				Expect(resourceRequestsModel.Memory).To(Equal(core.StringPtr("256MiB")))

				// Construct an instance of the ConfigCAUpdate model
				configCaUpdateModel := new(blockchainv2.ConfigCAUpdate)
				Expect(configCaUpdateModel).ToNot(BeNil())
				configCaUpdateModel.Cors = configCaCorsModel
				configCaUpdateModel.Debug = core.BoolPtr(false)
				configCaUpdateModel.Crlsizelimit = core.Float64Ptr(float64(512000))
				configCaUpdateModel.Tls = configCaTlsModel
				configCaUpdateModel.Ca = configCaCaModel
				configCaUpdateModel.Crl = configCaCrlModel
				configCaUpdateModel.Registry = configCaRegistryModel
				configCaUpdateModel.Db = configCaDbModel
				configCaUpdateModel.Affiliations = configCaAffiliationsModel
				configCaUpdateModel.Csr = configCaCsrModel
				configCaUpdateModel.Idemix = configCaIdemixModel
				configCaUpdateModel.BCCSP = bccspModel
				configCaUpdateModel.Intermediate = configCaIntermediateModel
				configCaUpdateModel.Cfg = configCaCfgModel
				configCaUpdateModel.Metrics = metricsModel
				Expect(configCaUpdateModel.Cors).To(Equal(configCaCorsModel))
				Expect(configCaUpdateModel.Debug).To(Equal(core.BoolPtr(false)))
				Expect(configCaUpdateModel.Crlsizelimit).To(Equal(core.Float64Ptr(float64(512000))))
				Expect(configCaUpdateModel.Tls).To(Equal(configCaTlsModel))
				Expect(configCaUpdateModel.Ca).To(Equal(configCaCaModel))
				Expect(configCaUpdateModel.Crl).To(Equal(configCaCrlModel))
				Expect(configCaUpdateModel.Registry).To(Equal(configCaRegistryModel))
				Expect(configCaUpdateModel.Db).To(Equal(configCaDbModel))
				Expect(configCaUpdateModel.Affiliations).To(Equal(configCaAffiliationsModel))
				Expect(configCaUpdateModel.Csr).To(Equal(configCaCsrModel))
				Expect(configCaUpdateModel.Idemix).To(Equal(configCaIdemixModel))
				Expect(configCaUpdateModel.BCCSP).To(Equal(bccspModel))
				Expect(configCaUpdateModel.Intermediate).To(Equal(configCaIntermediateModel))
				Expect(configCaUpdateModel.Cfg).To(Equal(configCaCfgModel))
				Expect(configCaUpdateModel.Metrics).To(Equal(metricsModel))

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv2.ResourceObject)
				Expect(resourceObjectModel).ToNot(BeNil())
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel
				Expect(resourceObjectModel.Requests).To(Equal(resourceRequestsModel))
				Expect(resourceObjectModel.Limits).To(Equal(resourceLimitsModel))

				// Construct an instance of the UpdateCaBodyConfigOverride model
				updateCaBodyConfigOverrideModel := new(blockchainv2.UpdateCaBodyConfigOverride)
				Expect(updateCaBodyConfigOverrideModel).ToNot(BeNil())
				updateCaBodyConfigOverrideModel.Ca = configCaUpdateModel
				Expect(updateCaBodyConfigOverrideModel.Ca).To(Equal(configCaUpdateModel))

				// Construct an instance of the UpdateCaBodyResources model
				updateCaBodyResourcesModel := new(blockchainv2.UpdateCaBodyResources)
				Expect(updateCaBodyResourcesModel).ToNot(BeNil())
				updateCaBodyResourcesModel.Ca = resourceObjectModel
				Expect(updateCaBodyResourcesModel.Ca).To(Equal(resourceObjectModel))

				// Construct an instance of the UpdateCaOptions model
				id := "testString"
				updateCaOptionsModel := testService.NewUpdateCaOptions(id)
				updateCaOptionsModel.SetID("testString")
				updateCaOptionsModel.SetResources(updateCaBodyResourcesModel)
				updateCaOptionsModel.SetZone("testString")
				updateCaOptionsModel.SetConfigOverride(updateCaBodyConfigOverrideModel)
				updateCaOptionsModel.SetReplicas(float64(1))
				updateCaOptionsModel.SetVersion("1.4.6-1")
				updateCaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCaOptionsModel).ToNot(BeNil())
				Expect(updateCaOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateCaOptionsModel.Resources).To(Equal(updateCaBodyResourcesModel))
				Expect(updateCaOptionsModel.Zone).To(Equal(core.StringPtr("testString")))
				Expect(updateCaOptionsModel.ConfigOverride).To(Equal(updateCaBodyConfigOverrideModel))
				Expect(updateCaOptionsModel.Replicas).To(Equal(core.Float64Ptr(float64(1))))
				Expect(updateCaOptionsModel.Version).To(Equal(core.StringPtr("1.4.6-1")))
				Expect(updateCaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateOrdererOptions successfully`, func() {
				// Construct an instance of the ConfigOrdererAuthentication model
				configOrdererAuthenticationModel := new(blockchainv2.ConfigOrdererAuthentication)
				Expect(configOrdererAuthenticationModel).ToNot(BeNil())
				configOrdererAuthenticationModel.TimeWindow = core.StringPtr("15m")
				configOrdererAuthenticationModel.NoExpirationChecks = core.BoolPtr(false)
				Expect(configOrdererAuthenticationModel.TimeWindow).To(Equal(core.StringPtr("15m")))
				Expect(configOrdererAuthenticationModel.NoExpirationChecks).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the ConfigOrdererKeepalive model
				configOrdererKeepaliveModel := new(blockchainv2.ConfigOrdererKeepalive)
				Expect(configOrdererKeepaliveModel).ToNot(BeNil())
				configOrdererKeepaliveModel.ServerMinInterval = core.StringPtr("60s")
				configOrdererKeepaliveModel.ServerInterval = core.StringPtr("2h")
				configOrdererKeepaliveModel.ServerTimeout = core.StringPtr("20s")
				Expect(configOrdererKeepaliveModel.ServerMinInterval).To(Equal(core.StringPtr("60s")))
				Expect(configOrdererKeepaliveModel.ServerInterval).To(Equal(core.StringPtr("2h")))
				Expect(configOrdererKeepaliveModel.ServerTimeout).To(Equal(core.StringPtr("20s")))

				// Construct an instance of the ConfigOrdererMetricsStatsd model
				configOrdererMetricsStatsdModel := new(blockchainv2.ConfigOrdererMetricsStatsd)
				Expect(configOrdererMetricsStatsdModel).ToNot(BeNil())
				configOrdererMetricsStatsdModel.Network = core.StringPtr("udp")
				configOrdererMetricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				configOrdererMetricsStatsdModel.WriteInterval = core.StringPtr("10s")
				configOrdererMetricsStatsdModel.Prefix = core.StringPtr("server")
				Expect(configOrdererMetricsStatsdModel.Network).To(Equal(core.StringPtr("udp")))
				Expect(configOrdererMetricsStatsdModel.Address).To(Equal(core.StringPtr("127.0.0.1:8125")))
				Expect(configOrdererMetricsStatsdModel.WriteInterval).To(Equal(core.StringPtr("10s")))
				Expect(configOrdererMetricsStatsdModel.Prefix).To(Equal(core.StringPtr("server")))

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv2.ResourceLimits)
				Expect(resourceLimitsModel).ToNot(BeNil())
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")
				Expect(resourceLimitsModel.Cpu).To(Equal(core.StringPtr("100m")))
				Expect(resourceLimitsModel.Memory).To(Equal(core.StringPtr("256MiB")))

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv2.ResourceRequests)
				Expect(resourceRequestsModel).ToNot(BeNil())
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")
				Expect(resourceRequestsModel.Cpu).To(Equal(core.StringPtr("100m")))
				Expect(resourceRequestsModel.Memory).To(Equal(core.StringPtr("256MiB")))

				// Construct an instance of the ConfigOrdererDebug model
				configOrdererDebugModel := new(blockchainv2.ConfigOrdererDebug)
				Expect(configOrdererDebugModel).ToNot(BeNil())
				configOrdererDebugModel.BroadcastTraceDir = core.StringPtr("testString")
				configOrdererDebugModel.DeliverTraceDir = core.StringPtr("testString")
				Expect(configOrdererDebugModel.BroadcastTraceDir).To(Equal(core.StringPtr("testString")))
				Expect(configOrdererDebugModel.DeliverTraceDir).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConfigOrdererGeneralUpdate model
				configOrdererGeneralUpdateModel := new(blockchainv2.ConfigOrdererGeneralUpdate)
				Expect(configOrdererGeneralUpdateModel).ToNot(BeNil())
				configOrdererGeneralUpdateModel.Keepalive = configOrdererKeepaliveModel
				configOrdererGeneralUpdateModel.Authentication = configOrdererAuthenticationModel
				Expect(configOrdererGeneralUpdateModel.Keepalive).To(Equal(configOrdererKeepaliveModel))
				Expect(configOrdererGeneralUpdateModel.Authentication).To(Equal(configOrdererAuthenticationModel))

				// Construct an instance of the ConfigOrdererMetrics model
				configOrdererMetricsModel := new(blockchainv2.ConfigOrdererMetrics)
				Expect(configOrdererMetricsModel).ToNot(BeNil())
				configOrdererMetricsModel.Provider = core.StringPtr("disabled")
				configOrdererMetricsModel.Statsd = configOrdererMetricsStatsdModel
				Expect(configOrdererMetricsModel.Provider).To(Equal(core.StringPtr("disabled")))
				Expect(configOrdererMetricsModel.Statsd).To(Equal(configOrdererMetricsStatsdModel))

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv2.ResourceObject)
				Expect(resourceObjectModel).ToNot(BeNil())
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel
				Expect(resourceObjectModel.Requests).To(Equal(resourceRequestsModel))
				Expect(resourceObjectModel.Limits).To(Equal(resourceLimitsModel))

				// Construct an instance of the ConfigOrdererUpdate model
				configOrdererUpdateModel := new(blockchainv2.ConfigOrdererUpdate)
				Expect(configOrdererUpdateModel).ToNot(BeNil())
				configOrdererUpdateModel.General = configOrdererGeneralUpdateModel
				configOrdererUpdateModel.Debug = configOrdererDebugModel
				configOrdererUpdateModel.Metrics = configOrdererMetricsModel
				Expect(configOrdererUpdateModel.General).To(Equal(configOrdererGeneralUpdateModel))
				Expect(configOrdererUpdateModel.Debug).To(Equal(configOrdererDebugModel))
				Expect(configOrdererUpdateModel.Metrics).To(Equal(configOrdererMetricsModel))

				// Construct an instance of the UpdateOrdererBodyResources model
				updateOrdererBodyResourcesModel := new(blockchainv2.UpdateOrdererBodyResources)
				Expect(updateOrdererBodyResourcesModel).ToNot(BeNil())
				updateOrdererBodyResourcesModel.Orderer = resourceObjectModel
				updateOrdererBodyResourcesModel.Proxy = resourceObjectModel
				Expect(updateOrdererBodyResourcesModel.Orderer).To(Equal(resourceObjectModel))
				Expect(updateOrdererBodyResourcesModel.Proxy).To(Equal(resourceObjectModel))

				// Construct an instance of the UpdateOrdererOptions model
				id := "testString"
				updateOrdererOptionsModel := testService.NewUpdateOrdererOptions(id)
				updateOrdererOptionsModel.SetID("testString")
				updateOrdererOptionsModel.SetConfigOverride(configOrdererUpdateModel)
				updateOrdererOptionsModel.SetResources(updateOrdererBodyResourcesModel)
				updateOrdererOptionsModel.SetZone("testString")
				updateOrdererOptionsModel.SetVersion("1.4.6-1")
				updateOrdererOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateOrdererOptionsModel).ToNot(BeNil())
				Expect(updateOrdererOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateOrdererOptionsModel.ConfigOverride).To(Equal(configOrdererUpdateModel))
				Expect(updateOrdererOptionsModel.Resources).To(Equal(updateOrdererBodyResourcesModel))
				Expect(updateOrdererOptionsModel.Zone).To(Equal(core.StringPtr("testString")))
				Expect(updateOrdererOptionsModel.Version).To(Equal(core.StringPtr("1.4.6-1")))
				Expect(updateOrdererOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdatePeerOptions successfully`, func() {
				// Construct an instance of the ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy model
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel := new(blockchainv2.ConfigPeerGossipPvtDataImplicitCollectionDisseminationPolicy)
				Expect(configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel).ToNot(BeNil())
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.RequiredPeerCount = core.Float64Ptr(float64(0))
				configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.MaxPeerCount = core.Float64Ptr(float64(1))
				Expect(configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.RequiredPeerCount).To(Equal(core.Float64Ptr(float64(0))))
				Expect(configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel.MaxPeerCount).To(Equal(core.Float64Ptr(float64(1))))

				// Construct an instance of the ConfigPeerDeliveryclientAddressOverridesItem model
				configPeerDeliveryclientAddressOverridesItemModel := new(blockchainv2.ConfigPeerDeliveryclientAddressOverridesItem)
				Expect(configPeerDeliveryclientAddressOverridesItemModel).ToNot(BeNil())
				configPeerDeliveryclientAddressOverridesItemModel.From = core.StringPtr("n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.To = core.StringPtr("n3a3ec3-myorderer2.ibp.us-south.containers.appdomain.cloud:7050")
				configPeerDeliveryclientAddressOverridesItemModel.CaCertsFile = core.StringPtr("my-data/cert.pem")
				Expect(configPeerDeliveryclientAddressOverridesItemModel.From).To(Equal(core.StringPtr("n3a3ec3-myorderer.ibp.us-south.containers.appdomain.cloud:7050")))
				Expect(configPeerDeliveryclientAddressOverridesItemModel.To).To(Equal(core.StringPtr("n3a3ec3-myorderer2.ibp.us-south.containers.appdomain.cloud:7050")))
				Expect(configPeerDeliveryclientAddressOverridesItemModel.CaCertsFile).To(Equal(core.StringPtr("my-data/cert.pem")))

				// Construct an instance of the ConfigPeerGossipElection model
				configPeerGossipElectionModel := new(blockchainv2.ConfigPeerGossipElection)
				Expect(configPeerGossipElectionModel).ToNot(BeNil())
				configPeerGossipElectionModel.StartupGracePeriod = core.StringPtr("15s")
				configPeerGossipElectionModel.MembershipSampleInterval = core.StringPtr("1s")
				configPeerGossipElectionModel.LeaderAliveThreshold = core.StringPtr("10s")
				configPeerGossipElectionModel.LeaderElectionDuration = core.StringPtr("5s")
				Expect(configPeerGossipElectionModel.StartupGracePeriod).To(Equal(core.StringPtr("15s")))
				Expect(configPeerGossipElectionModel.MembershipSampleInterval).To(Equal(core.StringPtr("1s")))
				Expect(configPeerGossipElectionModel.LeaderAliveThreshold).To(Equal(core.StringPtr("10s")))
				Expect(configPeerGossipElectionModel.LeaderElectionDuration).To(Equal(core.StringPtr("5s")))

				// Construct an instance of the ConfigPeerGossipPvtData model
				configPeerGossipPvtDataModel := new(blockchainv2.ConfigPeerGossipPvtData)
				Expect(configPeerGossipPvtDataModel).ToNot(BeNil())
				configPeerGossipPvtDataModel.PullRetryThreshold = core.StringPtr("60s")
				configPeerGossipPvtDataModel.TransientstoreMaxBlockRetention = core.Float64Ptr(float64(1000))
				configPeerGossipPvtDataModel.PushAckTimeout = core.StringPtr("3s")
				configPeerGossipPvtDataModel.BtlPullMargin = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileBatchSize = core.Float64Ptr(float64(10))
				configPeerGossipPvtDataModel.ReconcileSleepInterval = core.StringPtr("1m")
				configPeerGossipPvtDataModel.ReconciliationEnabled = core.BoolPtr(true)
				configPeerGossipPvtDataModel.SkipPullingInvalidTransactionsDuringCommit = core.BoolPtr(false)
				configPeerGossipPvtDataModel.ImplicitCollectionDisseminationPolicy = configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel
				Expect(configPeerGossipPvtDataModel.PullRetryThreshold).To(Equal(core.StringPtr("60s")))
				Expect(configPeerGossipPvtDataModel.TransientstoreMaxBlockRetention).To(Equal(core.Float64Ptr(float64(1000))))
				Expect(configPeerGossipPvtDataModel.PushAckTimeout).To(Equal(core.StringPtr("3s")))
				Expect(configPeerGossipPvtDataModel.BtlPullMargin).To(Equal(core.Float64Ptr(float64(10))))
				Expect(configPeerGossipPvtDataModel.ReconcileBatchSize).To(Equal(core.Float64Ptr(float64(10))))
				Expect(configPeerGossipPvtDataModel.ReconcileSleepInterval).To(Equal(core.StringPtr("1m")))
				Expect(configPeerGossipPvtDataModel.ReconciliationEnabled).To(Equal(core.BoolPtr(true)))
				Expect(configPeerGossipPvtDataModel.SkipPullingInvalidTransactionsDuringCommit).To(Equal(core.BoolPtr(false)))
				Expect(configPeerGossipPvtDataModel.ImplicitCollectionDisseminationPolicy).To(Equal(configPeerGossipPvtDataImplicitCollectionDisseminationPolicyModel))

				// Construct an instance of the ConfigPeerGossipState model
				configPeerGossipStateModel := new(blockchainv2.ConfigPeerGossipState)
				Expect(configPeerGossipStateModel).ToNot(BeNil())
				configPeerGossipStateModel.Enabled = core.BoolPtr(true)
				configPeerGossipStateModel.CheckInterval = core.StringPtr("10s")
				configPeerGossipStateModel.ResponseTimeout = core.StringPtr("3s")
				configPeerGossipStateModel.BatchSize = core.Float64Ptr(float64(10))
				configPeerGossipStateModel.BlockBufferSize = core.Float64Ptr(float64(100))
				configPeerGossipStateModel.MaxRetries = core.Float64Ptr(float64(3))
				Expect(configPeerGossipStateModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(configPeerGossipStateModel.CheckInterval).To(Equal(core.StringPtr("10s")))
				Expect(configPeerGossipStateModel.ResponseTimeout).To(Equal(core.StringPtr("3s")))
				Expect(configPeerGossipStateModel.BatchSize).To(Equal(core.Float64Ptr(float64(10))))
				Expect(configPeerGossipStateModel.BlockBufferSize).To(Equal(core.Float64Ptr(float64(100))))
				Expect(configPeerGossipStateModel.MaxRetries).To(Equal(core.Float64Ptr(float64(3))))

				// Construct an instance of the ConfigPeerKeepaliveClient model
				configPeerKeepaliveClientModel := new(blockchainv2.ConfigPeerKeepaliveClient)
				Expect(configPeerKeepaliveClientModel).ToNot(BeNil())
				configPeerKeepaliveClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveClientModel.Timeout = core.StringPtr("20s")
				Expect(configPeerKeepaliveClientModel.Interval).To(Equal(core.StringPtr("60s")))
				Expect(configPeerKeepaliveClientModel.Timeout).To(Equal(core.StringPtr("20s")))

				// Construct an instance of the ConfigPeerKeepaliveDeliveryClient model
				configPeerKeepaliveDeliveryClientModel := new(blockchainv2.ConfigPeerKeepaliveDeliveryClient)
				Expect(configPeerKeepaliveDeliveryClientModel).ToNot(BeNil())
				configPeerKeepaliveDeliveryClientModel.Interval = core.StringPtr("60s")
				configPeerKeepaliveDeliveryClientModel.Timeout = core.StringPtr("20s")
				Expect(configPeerKeepaliveDeliveryClientModel.Interval).To(Equal(core.StringPtr("60s")))
				Expect(configPeerKeepaliveDeliveryClientModel.Timeout).To(Equal(core.StringPtr("20s")))

				// Construct an instance of the ConfigPeerLimitsConcurrency model
				configPeerLimitsConcurrencyModel := new(blockchainv2.ConfigPeerLimitsConcurrency)
				Expect(configPeerLimitsConcurrencyModel).ToNot(BeNil())
				configPeerLimitsConcurrencyModel.EndorserService = map[string]interface{}{"anyKey": "anyValue"}
				configPeerLimitsConcurrencyModel.DeliverService = map[string]interface{}{"anyKey": "anyValue"}
				Expect(configPeerLimitsConcurrencyModel.EndorserService).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(configPeerLimitsConcurrencyModel.DeliverService).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))

				// Construct an instance of the ConfigPeerAdminService model
				configPeerAdminServiceModel := new(blockchainv2.ConfigPeerAdminService)
				Expect(configPeerAdminServiceModel).ToNot(BeNil())
				configPeerAdminServiceModel.ListenAddress = core.StringPtr("0.0.0.0:7051")
				Expect(configPeerAdminServiceModel.ListenAddress).To(Equal(core.StringPtr("0.0.0.0:7051")))

				// Construct an instance of the ConfigPeerAuthentication model
				configPeerAuthenticationModel := new(blockchainv2.ConfigPeerAuthentication)
				Expect(configPeerAuthenticationModel).ToNot(BeNil())
				configPeerAuthenticationModel.Timewindow = core.StringPtr("15m")
				Expect(configPeerAuthenticationModel.Timewindow).To(Equal(core.StringPtr("15m")))

				// Construct an instance of the ConfigPeerChaincodeExternalBuildersItem model
				configPeerChaincodeExternalBuildersItemModel := new(blockchainv2.ConfigPeerChaincodeExternalBuildersItem)
				Expect(configPeerChaincodeExternalBuildersItemModel).ToNot(BeNil())
				configPeerChaincodeExternalBuildersItemModel.Path = core.StringPtr("/path/to/directory")
				configPeerChaincodeExternalBuildersItemModel.Name = core.StringPtr("descriptive-build-name")
				configPeerChaincodeExternalBuildersItemModel.EnvironmentWhitelist = []string{"GOPROXY"}
				Expect(configPeerChaincodeExternalBuildersItemModel.Path).To(Equal(core.StringPtr("/path/to/directory")))
				Expect(configPeerChaincodeExternalBuildersItemModel.Name).To(Equal(core.StringPtr("descriptive-build-name")))
				Expect(configPeerChaincodeExternalBuildersItemModel.EnvironmentWhitelist).To(Equal([]string{"GOPROXY"}))

				// Construct an instance of the ConfigPeerChaincodeGolang model
				configPeerChaincodeGolangModel := new(blockchainv2.ConfigPeerChaincodeGolang)
				Expect(configPeerChaincodeGolangModel).ToNot(BeNil())
				configPeerChaincodeGolangModel.DynamicLink = core.BoolPtr(false)
				Expect(configPeerChaincodeGolangModel.DynamicLink).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the ConfigPeerChaincodeLogging model
				configPeerChaincodeLoggingModel := new(blockchainv2.ConfigPeerChaincodeLogging)
				Expect(configPeerChaincodeLoggingModel).ToNot(BeNil())
				configPeerChaincodeLoggingModel.Level = core.StringPtr("info")
				configPeerChaincodeLoggingModel.Shim = core.StringPtr("warning")
				configPeerChaincodeLoggingModel.Format = core.StringPtr("%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}")
				Expect(configPeerChaincodeLoggingModel.Level).To(Equal(core.StringPtr("info")))
				Expect(configPeerChaincodeLoggingModel.Shim).To(Equal(core.StringPtr("warning")))
				Expect(configPeerChaincodeLoggingModel.Format).To(Equal(core.StringPtr("%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}")))

				// Construct an instance of the ConfigPeerChaincodeSystem model
				configPeerChaincodeSystemModel := new(blockchainv2.ConfigPeerChaincodeSystem)
				Expect(configPeerChaincodeSystemModel).ToNot(BeNil())
				configPeerChaincodeSystemModel.Cscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Lscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Escc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Vscc = core.BoolPtr(true)
				configPeerChaincodeSystemModel.Qscc = core.BoolPtr(true)
				Expect(configPeerChaincodeSystemModel.Cscc).To(Equal(core.BoolPtr(true)))
				Expect(configPeerChaincodeSystemModel.Lscc).To(Equal(core.BoolPtr(true)))
				Expect(configPeerChaincodeSystemModel.Escc).To(Equal(core.BoolPtr(true)))
				Expect(configPeerChaincodeSystemModel.Vscc).To(Equal(core.BoolPtr(true)))
				Expect(configPeerChaincodeSystemModel.Qscc).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the ConfigPeerClient model
				configPeerClientModel := new(blockchainv2.ConfigPeerClient)
				Expect(configPeerClientModel).ToNot(BeNil())
				configPeerClientModel.ConnTimeout = core.StringPtr("2s")
				Expect(configPeerClientModel.ConnTimeout).To(Equal(core.StringPtr("2s")))

				// Construct an instance of the ConfigPeerDeliveryclient model
				configPeerDeliveryclientModel := new(blockchainv2.ConfigPeerDeliveryclient)
				Expect(configPeerDeliveryclientModel).ToNot(BeNil())
				configPeerDeliveryclientModel.ReconnectTotalTimeThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.ConnTimeout = core.StringPtr("2s")
				configPeerDeliveryclientModel.ReConnectBackoffThreshold = core.StringPtr("60m")
				configPeerDeliveryclientModel.AddressOverrides = []blockchainv2.ConfigPeerDeliveryclientAddressOverridesItem{*configPeerDeliveryclientAddressOverridesItemModel}
				Expect(configPeerDeliveryclientModel.ReconnectTotalTimeThreshold).To(Equal(core.StringPtr("60m")))
				Expect(configPeerDeliveryclientModel.ConnTimeout).To(Equal(core.StringPtr("2s")))
				Expect(configPeerDeliveryclientModel.ReConnectBackoffThreshold).To(Equal(core.StringPtr("60m")))
				Expect(configPeerDeliveryclientModel.AddressOverrides).To(Equal([]blockchainv2.ConfigPeerDeliveryclientAddressOverridesItem{*configPeerDeliveryclientAddressOverridesItemModel}))

				// Construct an instance of the ConfigPeerDiscovery model
				configPeerDiscoveryModel := new(blockchainv2.ConfigPeerDiscovery)
				Expect(configPeerDiscoveryModel).ToNot(BeNil())
				configPeerDiscoveryModel.Enabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheEnabled = core.BoolPtr(true)
				configPeerDiscoveryModel.AuthCacheMaxSize = core.Float64Ptr(float64(1000))
				configPeerDiscoveryModel.AuthCachePurgeRetentionRatio = core.Float64Ptr(float64(0.75))
				configPeerDiscoveryModel.OrgMembersAllowedAccess = core.BoolPtr(false)
				Expect(configPeerDiscoveryModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(configPeerDiscoveryModel.AuthCacheEnabled).To(Equal(core.BoolPtr(true)))
				Expect(configPeerDiscoveryModel.AuthCacheMaxSize).To(Equal(core.Float64Ptr(float64(1000))))
				Expect(configPeerDiscoveryModel.AuthCachePurgeRetentionRatio).To(Equal(core.Float64Ptr(float64(0.75))))
				Expect(configPeerDiscoveryModel.OrgMembersAllowedAccess).To(Equal(core.BoolPtr(false)))

				// Construct an instance of the ConfigPeerGossip model
				configPeerGossipModel := new(blockchainv2.ConfigPeerGossip)
				Expect(configPeerGossipModel).ToNot(BeNil())
				configPeerGossipModel.UseLeaderElection = core.BoolPtr(true)
				configPeerGossipModel.OrgLeader = core.BoolPtr(false)
				configPeerGossipModel.MembershipTrackerInterval = core.StringPtr("5s")
				configPeerGossipModel.MaxBlockCountToStore = core.Float64Ptr(float64(100))
				configPeerGossipModel.MaxPropagationBurstLatency = core.StringPtr("10ms")
				configPeerGossipModel.MaxPropagationBurstSize = core.Float64Ptr(float64(10))
				configPeerGossipModel.PropagateIterations = core.Float64Ptr(float64(3))
				configPeerGossipModel.PullInterval = core.StringPtr("4s")
				configPeerGossipModel.PullPeerNum = core.Float64Ptr(float64(3))
				configPeerGossipModel.RequestStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.PublishStateInfoInterval = core.StringPtr("4s")
				configPeerGossipModel.StateInfoRetentionInterval = core.StringPtr("0s")
				configPeerGossipModel.PublishCertPeriod = core.StringPtr("10s")
				configPeerGossipModel.SkipBlockVerification = core.BoolPtr(false)
				configPeerGossipModel.DialTimeout = core.StringPtr("3s")
				configPeerGossipModel.ConnTimeout = core.StringPtr("2s")
				configPeerGossipModel.RecvBuffSize = core.Float64Ptr(float64(20))
				configPeerGossipModel.SendBuffSize = core.Float64Ptr(float64(200))
				configPeerGossipModel.DigestWaitTime = core.StringPtr("1s")
				configPeerGossipModel.RequestWaitTime = core.StringPtr("1500ms")
				configPeerGossipModel.ResponseWaitTime = core.StringPtr("2s")
				configPeerGossipModel.AliveTimeInterval = core.StringPtr("5s")
				configPeerGossipModel.AliveExpirationTimeout = core.StringPtr("25s")
				configPeerGossipModel.ReconnectInterval = core.StringPtr("25s")
				configPeerGossipModel.Election = configPeerGossipElectionModel
				configPeerGossipModel.PvtData = configPeerGossipPvtDataModel
				configPeerGossipModel.State = configPeerGossipStateModel
				Expect(configPeerGossipModel.UseLeaderElection).To(Equal(core.BoolPtr(true)))
				Expect(configPeerGossipModel.OrgLeader).To(Equal(core.BoolPtr(false)))
				Expect(configPeerGossipModel.MembershipTrackerInterval).To(Equal(core.StringPtr("5s")))
				Expect(configPeerGossipModel.MaxBlockCountToStore).To(Equal(core.Float64Ptr(float64(100))))
				Expect(configPeerGossipModel.MaxPropagationBurstLatency).To(Equal(core.StringPtr("10ms")))
				Expect(configPeerGossipModel.MaxPropagationBurstSize).To(Equal(core.Float64Ptr(float64(10))))
				Expect(configPeerGossipModel.PropagateIterations).To(Equal(core.Float64Ptr(float64(3))))
				Expect(configPeerGossipModel.PullInterval).To(Equal(core.StringPtr("4s")))
				Expect(configPeerGossipModel.PullPeerNum).To(Equal(core.Float64Ptr(float64(3))))
				Expect(configPeerGossipModel.RequestStateInfoInterval).To(Equal(core.StringPtr("4s")))
				Expect(configPeerGossipModel.PublishStateInfoInterval).To(Equal(core.StringPtr("4s")))
				Expect(configPeerGossipModel.StateInfoRetentionInterval).To(Equal(core.StringPtr("0s")))
				Expect(configPeerGossipModel.PublishCertPeriod).To(Equal(core.StringPtr("10s")))
				Expect(configPeerGossipModel.SkipBlockVerification).To(Equal(core.BoolPtr(false)))
				Expect(configPeerGossipModel.DialTimeout).To(Equal(core.StringPtr("3s")))
				Expect(configPeerGossipModel.ConnTimeout).To(Equal(core.StringPtr("2s")))
				Expect(configPeerGossipModel.RecvBuffSize).To(Equal(core.Float64Ptr(float64(20))))
				Expect(configPeerGossipModel.SendBuffSize).To(Equal(core.Float64Ptr(float64(200))))
				Expect(configPeerGossipModel.DigestWaitTime).To(Equal(core.StringPtr("1s")))
				Expect(configPeerGossipModel.RequestWaitTime).To(Equal(core.StringPtr("1500ms")))
				Expect(configPeerGossipModel.ResponseWaitTime).To(Equal(core.StringPtr("2s")))
				Expect(configPeerGossipModel.AliveTimeInterval).To(Equal(core.StringPtr("5s")))
				Expect(configPeerGossipModel.AliveExpirationTimeout).To(Equal(core.StringPtr("25s")))
				Expect(configPeerGossipModel.ReconnectInterval).To(Equal(core.StringPtr("25s")))
				Expect(configPeerGossipModel.Election).To(Equal(configPeerGossipElectionModel))
				Expect(configPeerGossipModel.PvtData).To(Equal(configPeerGossipPvtDataModel))
				Expect(configPeerGossipModel.State).To(Equal(configPeerGossipStateModel))

				// Construct an instance of the ConfigPeerKeepalive model
				configPeerKeepaliveModel := new(blockchainv2.ConfigPeerKeepalive)
				Expect(configPeerKeepaliveModel).ToNot(BeNil())
				configPeerKeepaliveModel.MinInterval = core.StringPtr("60s")
				configPeerKeepaliveModel.Client = configPeerKeepaliveClientModel
				configPeerKeepaliveModel.DeliveryClient = configPeerKeepaliveDeliveryClientModel
				Expect(configPeerKeepaliveModel.MinInterval).To(Equal(core.StringPtr("60s")))
				Expect(configPeerKeepaliveModel.Client).To(Equal(configPeerKeepaliveClientModel))
				Expect(configPeerKeepaliveModel.DeliveryClient).To(Equal(configPeerKeepaliveDeliveryClientModel))

				// Construct an instance of the ConfigPeerLimits model
				configPeerLimitsModel := new(blockchainv2.ConfigPeerLimits)
				Expect(configPeerLimitsModel).ToNot(BeNil())
				configPeerLimitsModel.Concurrency = configPeerLimitsConcurrencyModel
				Expect(configPeerLimitsModel.Concurrency).To(Equal(configPeerLimitsConcurrencyModel))

				// Construct an instance of the MetricsStatsd model
				metricsStatsdModel := new(blockchainv2.MetricsStatsd)
				Expect(metricsStatsdModel).ToNot(BeNil())
				metricsStatsdModel.Network = core.StringPtr("udp")
				metricsStatsdModel.Address = core.StringPtr("127.0.0.1:8125")
				metricsStatsdModel.WriteInterval = core.StringPtr("10s")
				metricsStatsdModel.Prefix = core.StringPtr("server")
				Expect(metricsStatsdModel.Network).To(Equal(core.StringPtr("udp")))
				Expect(metricsStatsdModel.Address).To(Equal(core.StringPtr("127.0.0.1:8125")))
				Expect(metricsStatsdModel.WriteInterval).To(Equal(core.StringPtr("10s")))
				Expect(metricsStatsdModel.Prefix).To(Equal(core.StringPtr("server")))

				// Construct an instance of the ResourceLimits model
				resourceLimitsModel := new(blockchainv2.ResourceLimits)
				Expect(resourceLimitsModel).ToNot(BeNil())
				resourceLimitsModel.Cpu = core.StringPtr("100m")
				resourceLimitsModel.Memory = core.StringPtr("256MiB")
				Expect(resourceLimitsModel.Cpu).To(Equal(core.StringPtr("100m")))
				Expect(resourceLimitsModel.Memory).To(Equal(core.StringPtr("256MiB")))

				// Construct an instance of the ResourceRequests model
				resourceRequestsModel := new(blockchainv2.ResourceRequests)
				Expect(resourceRequestsModel).ToNot(BeNil())
				resourceRequestsModel.Cpu = core.StringPtr("100m")
				resourceRequestsModel.Memory = core.StringPtr("256MiB")
				Expect(resourceRequestsModel.Cpu).To(Equal(core.StringPtr("100m")))
				Expect(resourceRequestsModel.Memory).To(Equal(core.StringPtr("256MiB")))

				// Construct an instance of the ConfigPeerChaincode model
				configPeerChaincodeModel := new(blockchainv2.ConfigPeerChaincode)
				Expect(configPeerChaincodeModel).ToNot(BeNil())
				configPeerChaincodeModel.Golang = configPeerChaincodeGolangModel
				configPeerChaincodeModel.ExternalBuilders = []blockchainv2.ConfigPeerChaincodeExternalBuildersItem{*configPeerChaincodeExternalBuildersItemModel}
				configPeerChaincodeModel.InstallTimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Startuptimeout = core.StringPtr("300s")
				configPeerChaincodeModel.Executetimeout = core.StringPtr("30s")
				configPeerChaincodeModel.System = configPeerChaincodeSystemModel
				configPeerChaincodeModel.Logging = configPeerChaincodeLoggingModel
				Expect(configPeerChaincodeModel.Golang).To(Equal(configPeerChaincodeGolangModel))
				Expect(configPeerChaincodeModel.ExternalBuilders).To(Equal([]blockchainv2.ConfigPeerChaincodeExternalBuildersItem{*configPeerChaincodeExternalBuildersItemModel}))
				Expect(configPeerChaincodeModel.InstallTimeout).To(Equal(core.StringPtr("300s")))
				Expect(configPeerChaincodeModel.Startuptimeout).To(Equal(core.StringPtr("300s")))
				Expect(configPeerChaincodeModel.Executetimeout).To(Equal(core.StringPtr("30s")))
				Expect(configPeerChaincodeModel.System).To(Equal(configPeerChaincodeSystemModel))
				Expect(configPeerChaincodeModel.Logging).To(Equal(configPeerChaincodeLoggingModel))

				// Construct an instance of the ConfigPeerUpdatePeer model
				configPeerUpdatePeerModel := new(blockchainv2.ConfigPeerUpdatePeer)
				Expect(configPeerUpdatePeerModel).ToNot(BeNil())
				configPeerUpdatePeerModel.ID = core.StringPtr("john-doe")
				configPeerUpdatePeerModel.NetworkID = core.StringPtr("dev")
				configPeerUpdatePeerModel.Keepalive = configPeerKeepaliveModel
				configPeerUpdatePeerModel.Gossip = configPeerGossipModel
				configPeerUpdatePeerModel.Authentication = configPeerAuthenticationModel
				configPeerUpdatePeerModel.Client = configPeerClientModel
				configPeerUpdatePeerModel.Deliveryclient = configPeerDeliveryclientModel
				configPeerUpdatePeerModel.AdminService = configPeerAdminServiceModel
				configPeerUpdatePeerModel.ValidatorPoolSize = core.Float64Ptr(float64(8))
				configPeerUpdatePeerModel.Discovery = configPeerDiscoveryModel
				configPeerUpdatePeerModel.Limits = configPeerLimitsModel
				Expect(configPeerUpdatePeerModel.ID).To(Equal(core.StringPtr("john-doe")))
				Expect(configPeerUpdatePeerModel.NetworkID).To(Equal(core.StringPtr("dev")))
				Expect(configPeerUpdatePeerModel.Keepalive).To(Equal(configPeerKeepaliveModel))
				Expect(configPeerUpdatePeerModel.Gossip).To(Equal(configPeerGossipModel))
				Expect(configPeerUpdatePeerModel.Authentication).To(Equal(configPeerAuthenticationModel))
				Expect(configPeerUpdatePeerModel.Client).To(Equal(configPeerClientModel))
				Expect(configPeerUpdatePeerModel.Deliveryclient).To(Equal(configPeerDeliveryclientModel))
				Expect(configPeerUpdatePeerModel.AdminService).To(Equal(configPeerAdminServiceModel))
				Expect(configPeerUpdatePeerModel.ValidatorPoolSize).To(Equal(core.Float64Ptr(float64(8))))
				Expect(configPeerUpdatePeerModel.Discovery).To(Equal(configPeerDiscoveryModel))
				Expect(configPeerUpdatePeerModel.Limits).To(Equal(configPeerLimitsModel))

				// Construct an instance of the Metrics model
				metricsModel := new(blockchainv2.Metrics)
				Expect(metricsModel).ToNot(BeNil())
				metricsModel.Provider = core.StringPtr("prometheus")
				metricsModel.Statsd = metricsStatsdModel
				Expect(metricsModel.Provider).To(Equal(core.StringPtr("prometheus")))
				Expect(metricsModel.Statsd).To(Equal(metricsStatsdModel))

				// Construct an instance of the ResourceObject model
				resourceObjectModel := new(blockchainv2.ResourceObject)
				Expect(resourceObjectModel).ToNot(BeNil())
				resourceObjectModel.Requests = resourceRequestsModel
				resourceObjectModel.Limits = resourceLimitsModel
				Expect(resourceObjectModel.Requests).To(Equal(resourceRequestsModel))
				Expect(resourceObjectModel.Limits).To(Equal(resourceLimitsModel))

				// Construct an instance of the ResourceObjectCouchDb model
				resourceObjectCouchDbModel := new(blockchainv2.ResourceObjectCouchDb)
				Expect(resourceObjectCouchDbModel).ToNot(BeNil())
				resourceObjectCouchDbModel.Requests = resourceRequestsModel
				resourceObjectCouchDbModel.Limits = resourceLimitsModel
				Expect(resourceObjectCouchDbModel.Requests).To(Equal(resourceRequestsModel))
				Expect(resourceObjectCouchDbModel.Limits).To(Equal(resourceLimitsModel))

				// Construct an instance of the ResourceObjectFabV1 model
				resourceObjectFabV1Model := new(blockchainv2.ResourceObjectFabV1)
				Expect(resourceObjectFabV1Model).ToNot(BeNil())
				resourceObjectFabV1Model.Requests = resourceRequestsModel
				resourceObjectFabV1Model.Limits = resourceLimitsModel
				Expect(resourceObjectFabV1Model.Requests).To(Equal(resourceRequestsModel))
				Expect(resourceObjectFabV1Model.Limits).To(Equal(resourceLimitsModel))

				// Construct an instance of the ResourceObjectFabV2 model
				resourceObjectFabV2Model := new(blockchainv2.ResourceObjectFabV2)
				Expect(resourceObjectFabV2Model).ToNot(BeNil())
				resourceObjectFabV2Model.Requests = resourceRequestsModel
				resourceObjectFabV2Model.Limits = resourceLimitsModel
				Expect(resourceObjectFabV2Model.Requests).To(Equal(resourceRequestsModel))
				Expect(resourceObjectFabV2Model.Limits).To(Equal(resourceLimitsModel))

				// Construct an instance of the ConfigPeerUpdate model
				configPeerUpdateModel := new(blockchainv2.ConfigPeerUpdate)
				Expect(configPeerUpdateModel).ToNot(BeNil())
				configPeerUpdateModel.Peer = configPeerUpdatePeerModel
				configPeerUpdateModel.Chaincode = configPeerChaincodeModel
				configPeerUpdateModel.Metrics = metricsModel
				Expect(configPeerUpdateModel.Peer).To(Equal(configPeerUpdatePeerModel))
				Expect(configPeerUpdateModel.Chaincode).To(Equal(configPeerChaincodeModel))
				Expect(configPeerUpdateModel.Metrics).To(Equal(metricsModel))

				// Construct an instance of the PeerResources model
				peerResourcesModel := new(blockchainv2.PeerResources)
				Expect(peerResourcesModel).ToNot(BeNil())
				peerResourcesModel.Chaincodelauncher = resourceObjectFabV2Model
				peerResourcesModel.Couchdb = resourceObjectCouchDbModel
				peerResourcesModel.Statedb = resourceObjectModel
				peerResourcesModel.Dind = resourceObjectFabV1Model
				peerResourcesModel.Fluentd = resourceObjectFabV1Model
				peerResourcesModel.Peer = resourceObjectModel
				peerResourcesModel.Proxy = resourceObjectModel
				Expect(peerResourcesModel.Chaincodelauncher).To(Equal(resourceObjectFabV2Model))
				Expect(peerResourcesModel.Couchdb).To(Equal(resourceObjectCouchDbModel))
				Expect(peerResourcesModel.Statedb).To(Equal(resourceObjectModel))
				Expect(peerResourcesModel.Dind).To(Equal(resourceObjectFabV1Model))
				Expect(peerResourcesModel.Fluentd).To(Equal(resourceObjectFabV1Model))
				Expect(peerResourcesModel.Peer).To(Equal(resourceObjectModel))
				Expect(peerResourcesModel.Proxy).To(Equal(resourceObjectModel))

				// Construct an instance of the UpdatePeerOptions model
				id := "testString"
				updatePeerOptionsModel := testService.NewUpdatePeerOptions(id)
				updatePeerOptionsModel.SetID("testString")
				updatePeerOptionsModel.SetConfigOverride(configPeerUpdateModel)
				updatePeerOptionsModel.SetResources(peerResourcesModel)
				updatePeerOptionsModel.SetZone("testString")
				updatePeerOptionsModel.SetVersion("1.4.6-1")
				updatePeerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updatePeerOptionsModel).ToNot(BeNil())
				Expect(updatePeerOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updatePeerOptionsModel.ConfigOverride).To(Equal(configPeerUpdateModel))
				Expect(updatePeerOptionsModel.Resources).To(Equal(peerResourcesModel))
				Expect(updatePeerOptionsModel.Zone).To(Equal(core.StringPtr("testString")))
				Expect(updatePeerOptionsModel.Version).To(Equal(core.StringPtr("1.4.6-1")))
				Expect(updatePeerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewHsm successfully`, func() {
				pkcs11endpoint := "tcp://example.com:666"
				model, err := testService.NewHsm(pkcs11endpoint)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
