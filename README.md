![pepipostlogo](https://pepipost.com/assets/img/pepipost-footLogo.png)

[![Twitter Follow](https://img.shields.io/twitter/follow/pepi_post.svg?style=social&label=Follow)](https://twitter.com/pepi_post)

# Official Go Library for [Pepipost](http://www.pepipost.com/?utm_campaign=GitHubSDK&utm_medium=GithubSDK&utm_source=GithubSDK)

This library contains methods for easily interacting with the Pepipost Email Sending API to send emails within few seconds.

We are trying to make our libraries Community Driven- which means we need your help in building the right things in proper order we would request you to help us by sharing comments, creating new [issues](https://github.com/pepipost/pepipost-sdk-go/issues) or [pull requests](https://github.com/pepipost/pepipost-sdk-go/pulls).

We welcome any sort of contribution to this library.

The latest 2.5.0 version of this library provides is fully compatible with the latest Pepipost v2.0 API.

For any update of this library check [Releases](https://github.com/pepipost/pepipost-sdk-go/releases).

# Table of Content
  
* [Installation](#installation)
* [Quick Start](#quick-start)
* [Sample Example](#example)
* [Announcements](#announcements)
* [Roadmap](#roadmap)
* [About](#about)
* [License](#license)

<a name="installation"></a>
# Installation

* In order to successfully build and run your SDK files, you are required to have the following Prerequisites in you system:

<a name="prereq"></a>
## Prerequisites

* **GO** (visit [Go installation page](https://golang.org/doc/install) for details on how to install Go)
* Ensure that **```GOROOT```** & **```GOPATH```** enviroment variable is set in the system variables. if not, set it to your working directory where you add your Go projects.
* This library uses unirest-go http library. Therefore it will require internet access to resolve the dependency.
  
  * If Go is properly installed and configure, copy and run the following command to pull the dependency.
  
    ```Go
    go get github.com/apimatic/unirest-go
    ```
    
    ![pic1.5](http://app1.falconide.com/integration_imgs/goimg/capture(25).png)
    
  * This will install unirest-go in the ```GOPATH``` you specified in the system variables.

<a name='quick-start'></a>
## Quick Start

Quick guide for installing Pepipost Go library 

* Check **GOPATH** using below command

  ``` echo $GOPATH ```
  
  ![pic1](http://app1.falconide.com/integration_imgs/goimg/capture(24).png)

* Change your directory to **$GOPATH** 
  
  ``` cd $GOPATH```
  
* Clone or download the repository using below command

  ``` git clone https://github.com/pepipost/pepipost-sdk-go.git ```
  
  **OR**
  
  Download **[zip](https://github.com/pepipost/pepipost-sdk-go/archive/master.zip)**

* Run below command to make directory and copy library file ```pepipost_lib``` to your GOPATH **$GOPATH/src** directory 

  ``` mkdir $GOPATH/src/pepipost_lib ```

  ``` cp -r pepipost-sdk-go/src/pepipost_lib/* $GOPATH/src/pepipost_lib/ ```
  
  ![pic4](http://app1.falconide.com/integration_imgs/goimg/capture(26).png)
  
  ![pic5](http://app1.falconide.com/integration_imgs/goimg/capture(27).png)

* make project using below command 

  ```  mkdir $GOPATH/src/pepipost_test && cd $GOPATH/src/pepipost_test ```
  
  ![pic7](http://app1.falconide.com/integration_imgs/goimg/capture(28).png)

* Create sample Go file named **main.go**

  copy the below [sample example](#example) in your ```main.go``` file.
  
* Change your api-key and sending domain respectively

    * **apikey** will be available under Login to Pepipost -> Settings -> Integration  
    * **Sending Domain** will be available under Login to Pepiost -> Settings -> Sending Domains 
  
    ```
  *Note :: Domains showing with Active status on Sending Domain dashboard are only allowed to send any sort of emails.* In case there are no Sending Domain added under your account, then first add the domain, get the DNS (SPF/DKIM) settings done and get it reviewed by our compliance team for approval. Once the domain is approved, it will be in ACTIVE status and will be ready to send any sort of emails. 
    ```
* Sending first email using Go library

  ```go build main.go```
  
  ```./main.exe ``` (for windows) **OR** ```./main```
  
  ![img](http://app1.falconide.com/integration_imgs/goimg/capture(30).png)
 

<a name='example'></a>
## Sample Example  

```Go
package  main
import (
        "fmt"
        "pepipost_lib/pepipost_pkg"
        "pepipost_lib/models_pkg"
)

func main() {
        client := PepipostClient.NewPEPIPOST()
        email := client.Email()
        ApiKey := "your apikey here"
        username := "to recipient here"
        fromEmail := "fromEmail@your verified domain"
        subject := "GO-SDK Pepipost "
        body := "This is Go-SDK"
        Body := &models_pkg.EmailBody{}
        Body.Personalizations = make([]*models_pkg.Personalizations,3)

        Body.Personalizations[0] = &models_pkg.Personalizations{}
        Body.Personalizations[0].Recipient = &username

        Body.From = models_pkg.From{}
        Body.From.FromEmail = &fromEmail
        Body.Subject = &subject
        Body.Content = &body

        var err error
        var result *models_pkg.SendEmailResponse
        result, err = email.CreateSendEmail(&ApiKey, Body)

        if err != nil{
                //TODO: Use err variable here
                fmt.Println(result)
        }else{
                //TODO: Use result variable here
                fmt.Println(result)
        }
}
```

<a name="announcements"></a>
# Announcements

v2.5.0 has been released! Please see the [release notes](https://github.com/pepipost/pepipost-sdk-go/releases/) for details.

All updates to this library are documented in our [releases](https://github.com/pepipost/pepipost-sdk-go/releases). For any queries, feel free to reach out us at dx@pepipost.com

<a name="roadmap"></a>
## Roadmap

If you are interested in the future direction of this project, please take a look at our open [issues](https://github.com/pepipost/pepipost-sdk-go/issues) and [pull requests](https://github.com/pepipost/pepipost-sdk-go/pulls). We would love to hear your feedback.

<a name="about"></a>
## About
pepipost-sdk-go library is guided and supported by the [Pepipost Developer Experience Team](https://github.com/orgs/pepipost/teams/pepis/members) .
This pepipost library is maintained and funded by Pepipost Ltd. The names and logos for pepipost are trademarks of Pepipost Ltd.

<a name="license"></a>
## License
This code library was semi-automatically generated by APIMATIC v2.0 and licensed under The MIT License (MIT).



