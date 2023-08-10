// +build integration

/**
 * (C) Copyright IBM Corp. 2023.
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

package usagereportsv4_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/usagereportsv4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the usagereportsv4 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`UsageReportsV4 Integration Tests`, func() {
	const externalConfigFile = "../usage_reports_v4.env"

	var (
		err          error
		usageReportsService *usagereportsv4.UsageReportsV4
		serviceURL   string
		config       map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(usagereportsv4.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			usageReportsServiceOptions := &usagereportsv4.UsageReportsV4Options{}

			usageReportsService, err = usagereportsv4.NewUsageReportsV4UsingExternalConfig(usageReportsServiceOptions)
			Expect(err).To(BeNil())
			Expect(usageReportsService).ToNot(BeNil())
			Expect(usageReportsService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			usageReportsService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`GetAccountSummary - Get account summary`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAccountSummary(getAccountSummaryOptions *GetAccountSummaryOptions)`, func() {
			getAccountSummaryOptions := &usagereportsv4.GetAccountSummaryOptions{
				AccountID: core.StringPtr("testString"),
				Billingmonth: core.StringPtr("testString"),
				Accept: core.StringPtr("application/json"),
				Format: core.StringPtr("csv"),
			}

			accountSummary, response, err := usageReportsService.GetAccountSummary(getAccountSummaryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSummary).ToNot(BeNil())
		})
	})

	Describe(`GetAccountUsage - Get account usage`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAccountUsage(getAccountUsageOptions *GetAccountUsageOptions)`, func() {
			getAccountUsageOptions := &usagereportsv4.GetAccountUsageOptions{
				AccountID: core.StringPtr("testString"),
				Billingmonth: core.StringPtr("testString"),
				Names: core.BoolPtr(true),
				AcceptLanguage: core.StringPtr("testString"),
			}

			accountUsage, response, err := usageReportsService.GetAccountUsage(getAccountUsageOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountUsage).ToNot(BeNil())
		})
	})

	Describe(`GetResourceGroupUsage - Get resource group usage`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourceGroupUsage(getResourceGroupUsageOptions *GetResourceGroupUsageOptions)`, func() {
			getResourceGroupUsageOptions := &usagereportsv4.GetResourceGroupUsageOptions{
				AccountID: core.StringPtr("testString"),
				ResourceGroupID: core.StringPtr("testString"),
				Billingmonth: core.StringPtr("testString"),
				Names: core.BoolPtr(true),
				AcceptLanguage: core.StringPtr("testString"),
			}

			resourceGroupUsage, response, err := usageReportsService.GetResourceGroupUsage(getResourceGroupUsageOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceGroupUsage).ToNot(BeNil())
		})
	})

	Describe(`GetResourceUsageAccount - Get resource instance usage in an account`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourceUsageAccount(getResourceUsageAccountOptions *GetResourceUsageAccountOptions) with pagination`, func(){
			getResourceUsageAccountOptions := &usagereportsv4.GetResourceUsageAccountOptions{
				AccountID: core.StringPtr("testString"),
				Billingmonth: core.StringPtr("testString"),
				Accept: core.StringPtr("application/json"),
				Format: core.StringPtr("csv"),
				Names: core.BoolPtr(true),
				Tags: core.BoolPtr(true),
				AcceptLanguage: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(30)),
				Start: core.StringPtr("testString"),
				ResourceGroupID: core.StringPtr("testString"),
				OrganizationID: core.StringPtr("testString"),
				ResourceInstanceID: core.StringPtr("testString"),
				ResourceID: core.StringPtr("testString"),
				PlanID: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
			}

			getResourceUsageAccountOptions.Start = nil

			var allResults []usagereportsv4.InstanceUsage
			for {
				instancesUsage, response, err := usageReportsService.GetResourceUsageAccount(getResourceUsageAccountOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(instancesUsage).ToNot(BeNil())
				allResults = append(allResults, instancesUsage.Resources...)

				getResourceUsageAccountOptions.Start, err = instancesUsage.GetNextStart()
				Expect(err).To(BeNil())

				if getResourceUsageAccountOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`GetResourceUsageAccount(getResourceUsageAccountOptions *GetResourceUsageAccountOptions) using GetResourceUsageAccountPager`, func(){
			getResourceUsageAccountOptions := &usagereportsv4.GetResourceUsageAccountOptions{
				AccountID: core.StringPtr("testString"),
				Billingmonth: core.StringPtr("testString"),
				Accept: core.StringPtr("application/json"),
				Format: core.StringPtr("csv"),
				Names: core.BoolPtr(true),
				Tags: core.BoolPtr(true),
				AcceptLanguage: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(30)),
				ResourceGroupID: core.StringPtr("testString"),
				OrganizationID: core.StringPtr("testString"),
				ResourceInstanceID: core.StringPtr("testString"),
				ResourceID: core.StringPtr("testString"),
				PlanID: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
			}

			// Test GetNext().
			pager, err := usageReportsService.NewGetResourceUsageAccountPager(getResourceUsageAccountOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []usagereportsv4.InstanceUsage
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = usageReportsService.NewGetResourceUsageAccountPager(getResourceUsageAccountOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "GetResourceUsageAccount() returned a total of %d item(s) using GetResourceUsageAccountPager.\n", len(allResults))
		})
	})

	Describe(`GetResourceUsageResourceGroup - Get resource instance usage in a resource group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourceUsageResourceGroup(getResourceUsageResourceGroupOptions *GetResourceUsageResourceGroupOptions) with pagination`, func(){
			getResourceUsageResourceGroupOptions := &usagereportsv4.GetResourceUsageResourceGroupOptions{
				AccountID: core.StringPtr("testString"),
				ResourceGroupID: core.StringPtr("testString"),
				Billingmonth: core.StringPtr("testString"),
				Names: core.BoolPtr(true),
				Tags: core.BoolPtr(true),
				AcceptLanguage: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(30)),
				Start: core.StringPtr("testString"),
				ResourceInstanceID: core.StringPtr("testString"),
				ResourceID: core.StringPtr("testString"),
				PlanID: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
			}

			getResourceUsageResourceGroupOptions.Start = nil

			var allResults []usagereportsv4.InstanceUsage
			for {
				instancesUsage, response, err := usageReportsService.GetResourceUsageResourceGroup(getResourceUsageResourceGroupOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(instancesUsage).ToNot(BeNil())
				allResults = append(allResults, instancesUsage.Resources...)

				getResourceUsageResourceGroupOptions.Start, err = instancesUsage.GetNextStart()
				Expect(err).To(BeNil())

				if getResourceUsageResourceGroupOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`GetResourceUsageResourceGroup(getResourceUsageResourceGroupOptions *GetResourceUsageResourceGroupOptions) using GetResourceUsageResourceGroupPager`, func(){
			getResourceUsageResourceGroupOptions := &usagereportsv4.GetResourceUsageResourceGroupOptions{
				AccountID: core.StringPtr("testString"),
				ResourceGroupID: core.StringPtr("testString"),
				Billingmonth: core.StringPtr("testString"),
				Names: core.BoolPtr(true),
				Tags: core.BoolPtr(true),
				AcceptLanguage: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(30)),
				ResourceInstanceID: core.StringPtr("testString"),
				ResourceID: core.StringPtr("testString"),
				PlanID: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
			}

			// Test GetNext().
			pager, err := usageReportsService.NewGetResourceUsageResourceGroupPager(getResourceUsageResourceGroupOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []usagereportsv4.InstanceUsage
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = usageReportsService.NewGetResourceUsageResourceGroupPager(getResourceUsageResourceGroupOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "GetResourceUsageResourceGroup() returned a total of %d item(s) using GetResourceUsageResourceGroupPager.\n", len(allResults))
		})
	})

	Describe(`GetResourceUsageOrg - Get resource instance usage in an organization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourceUsageOrg(getResourceUsageOrgOptions *GetResourceUsageOrgOptions) with pagination`, func(){
			getResourceUsageOrgOptions := &usagereportsv4.GetResourceUsageOrgOptions{
				AccountID: core.StringPtr("testString"),
				OrganizationID: core.StringPtr("testString"),
				Billingmonth: core.StringPtr("testString"),
				Names: core.BoolPtr(true),
				Tags: core.BoolPtr(true),
				AcceptLanguage: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(30)),
				Start: core.StringPtr("testString"),
				ResourceInstanceID: core.StringPtr("testString"),
				ResourceID: core.StringPtr("testString"),
				PlanID: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
			}

			getResourceUsageOrgOptions.Start = nil

			var allResults []usagereportsv4.InstanceUsage
			for {
				instancesUsage, response, err := usageReportsService.GetResourceUsageOrg(getResourceUsageOrgOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(instancesUsage).ToNot(BeNil())
				allResults = append(allResults, instancesUsage.Resources...)

				getResourceUsageOrgOptions.Start, err = instancesUsage.GetNextStart()
				Expect(err).To(BeNil())

				if getResourceUsageOrgOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`GetResourceUsageOrg(getResourceUsageOrgOptions *GetResourceUsageOrgOptions) using GetResourceUsageOrgPager`, func(){
			getResourceUsageOrgOptions := &usagereportsv4.GetResourceUsageOrgOptions{
				AccountID: core.StringPtr("testString"),
				OrganizationID: core.StringPtr("testString"),
				Billingmonth: core.StringPtr("testString"),
				Names: core.BoolPtr(true),
				Tags: core.BoolPtr(true),
				AcceptLanguage: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(30)),
				ResourceInstanceID: core.StringPtr("testString"),
				ResourceID: core.StringPtr("testString"),
				PlanID: core.StringPtr("testString"),
				Region: core.StringPtr("testString"),
			}

			// Test GetNext().
			pager, err := usageReportsService.NewGetResourceUsageOrgPager(getResourceUsageOrgOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []usagereportsv4.InstanceUsage
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = usageReportsService.NewGetResourceUsageOrgPager(getResourceUsageOrgOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "GetResourceUsageOrg() returned a total of %d item(s) using GetResourceUsageOrgPager.\n", len(allResults))
		})
	})

	Describe(`GetOrgUsage - Get organization usage`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetOrgUsage(getOrgUsageOptions *GetOrgUsageOptions)`, func() {
			getOrgUsageOptions := &usagereportsv4.GetOrgUsageOptions{
				AccountID: core.StringPtr("testString"),
				OrganizationID: core.StringPtr("testString"),
				Billingmonth: core.StringPtr("testString"),
				Names: core.BoolPtr(true),
				AcceptLanguage: core.StringPtr("testString"),
			}

			orgUsage, response, err := usageReportsService.GetOrgUsage(getOrgUsageOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(orgUsage).ToNot(BeNil())
		})
	})

	Describe(`CreateReportsSnapshotConfig - Setup the snapshot configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateReportsSnapshotConfig(createReportsSnapshotConfigOptions *CreateReportsSnapshotConfigOptions)`, func() {
			createReportsSnapshotConfigOptions := &usagereportsv4.CreateReportsSnapshotConfigOptions{
				AccountID: core.StringPtr("abc"),
				Interval: core.StringPtr("daily"),
				CosBucket: core.StringPtr("bucket_name"),
				CosLocation: core.StringPtr("us-south"),
				CosReportsFolder: core.StringPtr("IBMCloud-Billing-Reports"),
				ReportTypes: []string{"account_summary", "enterprise_summary", "account_resource_instance_usage"},
				Versioning: core.StringPtr("new"),
			}

			snapshotConfig, response, err := usageReportsService.CreateReportsSnapshotConfig(createReportsSnapshotConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(snapshotConfig).ToNot(BeNil())
		})
	})

	Describe(`GetReportsSnapshotConfig - Fetch the snapshot configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportsSnapshotConfig(getReportsSnapshotConfigOptions *GetReportsSnapshotConfigOptions)`, func() {
			getReportsSnapshotConfigOptions := &usagereportsv4.GetReportsSnapshotConfigOptions{
				AccountID: core.StringPtr("abc"),
			}

			snapshotConfig, response, err := usageReportsService.GetReportsSnapshotConfig(getReportsSnapshotConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshotConfig).ToNot(BeNil())
		})
	})

	Describe(`UpdateReportsSnapshotConfig - Update the snapshot configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateReportsSnapshotConfig(updateReportsSnapshotConfigOptions *UpdateReportsSnapshotConfigOptions)`, func() {
			updateReportsSnapshotConfigOptions := &usagereportsv4.UpdateReportsSnapshotConfigOptions{
				AccountID: core.StringPtr("abc"),
				Interval: core.StringPtr("daily"),
				CosBucket: core.StringPtr("bucket_name"),
				CosLocation: core.StringPtr("us-south"),
				CosReportsFolder: core.StringPtr("IBMCloud-Billing-Reports"),
				ReportTypes: []string{"account_summary", "enterprise_summary", "account_resource_instance_usage"},
				Versioning: core.StringPtr("new"),
			}

			snapshotConfig, response, err := usageReportsService.UpdateReportsSnapshotConfig(updateReportsSnapshotConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshotConfig).ToNot(BeNil())
		})
	})

	Describe(`GetReportsSnapshot - Fetch the current or past snapshots`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportsSnapshot(getReportsSnapshotOptions *GetReportsSnapshotOptions)`, func() {
			getReportsSnapshotOptions := &usagereportsv4.GetReportsSnapshotOptions{
				AccountID: core.StringPtr("abc"),
				Month: core.StringPtr("2023-02"),
				DateFrom: core.Float64Ptr(float64(1675209600000)),
				DateTo: core.Float64Ptr(float64(1675987200000)),
			}

			snapshotList, response, err := usageReportsService.GetReportsSnapshot(getReportsSnapshotOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshotList).ToNot(BeNil())
		})
	})

	Describe(`DeleteReportsSnapshotConfig - Delete the snapshot configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteReportsSnapshotConfig(deleteReportsSnapshotConfigOptions *DeleteReportsSnapshotConfigOptions)`, func() {
			deleteReportsSnapshotConfigOptions := &usagereportsv4.DeleteReportsSnapshotConfigOptions{
				AccountID: core.StringPtr("abc"),
			}

			response, err := usageReportsService.DeleteReportsSnapshotConfig(deleteReportsSnapshotConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
