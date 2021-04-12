CHANGELOG
=========

## HEAD (Unreleased)
* Upgrade to pulumi-terraform-bridge v2.23.0

---

## 0.8.1 (2021-03-23)
* Upgrade to pulumi-terraform-bridge v2.22.1  
  **Please Note:** This includes a bug fix to the refresh operation regarding arrays

## 0.8.0 (2021-03-15)
* Upgrade to pulumi-terraform-bridge v2.21.0
* Release macOS arm64 binary

## 0.7.2 (2021-03-09)
* Upgrade to v2.8.3 of the Wavefront Terraform Provider

## 0.7.1 (2021-02-16)
* Upgrade to pulumi-terraform-bridge v2.19.0  
  **Please Note:** This includes a bug fix that stops mutating resources options in the nodejs provider
* Avoid storing config from the environment into the state

## 0.7.0 (2021-02-03)
* Upgrade to v2.8.1 of the Wavefront Terraform Provider

## 0.6.0 (2021-01-29)
* Upgrading pulumi-terraform-bridge to v2.18.0

## 0.5.5 (2021-01-13)
* Upgrade to pulumi-terraform-bridge v2.17.0
* Upgrade to Pulumi v2.17.0

## 0.5.4 (2021-01-05)
* Upgrade to pulumi-terraform-bridge v2.13.2
  * This adds support for import specific examples in documentation

## 0.5.3 (2020-12-21)
* Upgrade to v2.7.3 of the Wavefront Terraform Provider

## 0.5.2 (2020-12-10)
* Upgrade to v2.7.2 of the Wavefront Terraform Provider

## 0.5.1 (2020-10-28)
* Upgrade to v2.7.1 of the Wavefront Terraform Provider

## 0.5.0 (2020-10-26)
* Upgrade to Pulumi v2.12.0 and pulumi-terraform-bridge v2.11.0
* Improving the accuracy of previews leading to a more accurate understanding of what will actually change rather than assuming all output properties will change.  
  ** PLEASE NOTE:**  
  This new preview functionality can be disabled by setting `PULUMI_DISABLE_PROVIDER_PREVIEW` to `1` or `false`.

## 0.4.0 (2020-10-06)
* Upgrade to v2.7.0 of the Wavefront Terraform Provider

## 0.3.0 (2020-09-29)
* Upgrade to v2.6.0 of the Wavefront Terraform Provider

## 0.2.1 (2020-09-21)
* Upgrade to v2.5.1 of the Wavefront Terraform Provider

## 0.2.0 (2020-09-14)
* Upgrade to v2.5.0 of the Wavefront Terraform Provider
* Upgrade to pulumi-terraform-bridge v2.8.0
* Upgrade to Pulumi v2.10.0

## 0.1.0 (2020-08-31)
* Upgrade to pulumi-terraform-bridge v2.7.3
* Upgrade to Pulumi v2.9.0, which adds type annotations and input/output classes to Python

## 0.0.1 (2020-08-12)
* Initial creation of the provider against v2.3.1 of the Wavefront Terraform Provider
