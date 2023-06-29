# Overview

# Authentication

Before starting the authentication process via CLI, please make sure that you have a valid account and valid tenant name.

## **Authenticating**
The authentication process will go through 4 steps:

### **1. Input authentication information via CLI:**

Use this command to begin the authn process:

```
onqlave auth login your_email_account@domain.com -t your-verified-tenant-name
```

After entering the above command, the output should appear as follows:
```
Login instruction is sent to email address 'your_email_count@domain.com'.
Please be mindful that the link provided in email is only valid for 30 minutes.
```

### **2. Wait for an email from Onqlave**

You should shortly receive an email at the address you provided (in the previous step).

Please note that the link contained within the email is only valid for 30 minutes.

### **3. Confirm your email**
Open your email and click on the provided link to complete the email confirmation prompt
![Email confirmation](https://t36712295.p.clickup-attachments.com/t36712295/629391bb-7432-442c-9dd0-f24f91cbae7b/image.png)

### **4. Return to the CLI and work with an authenticated session**
After confirming your email, the CLI output should look like this:
```
🎉 Done!  You successfully login to Onqlave platform.

For more information, read our documentation at https://docs.onqlave.com

```

From here you can explore the interaction options provided by onqlave-cli using one of the following commands:

```plain
onqlave
```

or

```plain
onqlave help
```

or

```plain
onqlave --help
```

Each will produce instructions for the supported commands of onqlave-cli:

```plain
Usage:
  onqlave [command]

Examples:
onqlave

Available Commands:
  application application management
  arx         arx management
  auth        authentication
  completion  Generate the autocompletion script for the specified shell
  config      config environment variables
  help        Help about any command
  key         api key management
  tenant      tenant management
  user        user management

Flags:
  -h, --help      help for onqlave
      --json      JSON Output. Set to true if stdout is not a TTY.
  -v, --version   version for onqlave

Use "onqlave [command] --help" for more information about a command


```

# Administration

# Arx

## **Inspiration**

If you skipped the section about the meaning of Arx, you can have a look at the previous page.

## **Before you start**

If you are familiar with allocating cloud computing resources for your company, you will be familiar with the concept of creating and assigning Arx to support your expected workload. With Onqlave, we follow a similar approach to allow you to optimise for speed and availability.



### **Review the provided permission/role**

Our current release includes 3 defined roles: Platform Owner, Platform Admin and Developer. Each role has its own set of permissions and supported operations. You may need to first skim through the documentation on roles and supported operations

Next, you can explore a list of **[available commands](#explore-the-supported-interaction-commands-with-arx)** or explore the **[base value of supported configuration of Arx](#get-base-configuration-information-for-arx).**

When you are ready to interact, you can go through the below list of commands to perform various operations.

## **Create an Arx**

**Who can perform this operation?**

- Platform Owner

To create an Arx, simply input the following command into the CLI. Please pay attention to the flags and their available assigned values.
There are several configurable attributes of an Arx that are grouped into 4 sections:
- **Planning**: We support you in segregating the development, testing, staging and production by providing single purposed Arx for each of your desired environments, including: development, testing, staging, and production.

- **Cloud Provider**: The choice of cloud provider determines which service is used to store your information. This allows for you to choose a cloud provider that your organisation already uses. Currently, we only support Google, but more providers are coming soon.

- **Region**: The choice of region allows you to determine within which geography you would like the data to reside. This may an important factor for data localisation / data residency requirements for sensitive data, whilst there can also be additional [speed and efficiency] benefits from having the data reside in the same geography as the rest of your information.

- **Encryption mechanism**: We only offer encryption services based on the highest performance encryption algorithms. You have the choice of AES-GCM-128, AES-GCM-256 or XCHACHA20-POLY1305, with the latter offering stronger encryption but at a lower processing speed. The key rotation frequency determines how regularly the encryption keys are changed. More regular rotations increase the level of security to ensure that your information remains safe.

You can append the *-h* or *--help* flag to the end of the add command to see the available flags:

```
onqlave arx add -h
```
The suggested output should look like this:

```
This command is used to add arx. Valid arx name, arx provider,  arx type, arx purpose, arx region, arx description, arx encryption method, arx rotation cycle, arx owner, arx spend limit and arx is default are required.

Usage:
  onqlave arx add [flags]

Examples:
onqlave arx add

Flags:
  -d, --arx_description string         enter arx description
  -e, --arx_encryption_method string   enter arx encryption method
  -i, --arx_is_default                 enter arx is default
  -o, --arx_owner string               enter arx owner
  -p, --arx_provider string            enter arx cloud provider
  -u, --arx_purpose string             enter arx purpose
  -r, --arx_region string              enter arx region
  -c, --arx_rotation_cycle string      enter arx rotation cycle
  -l, --arx_spend_limit uint           enter arx spend limit
  -t, --arx_type string                enter arx type
  -h, --help                           help for add

Global Flags:
      --json   JSON Output. Set to true if stdout is not a TTY.
```
Next, you can explore the **[base value of supported configuration of Arx](#get-base-configuration-information-for-arx).**

After that, when you are ready, just enter the **add** command:

```
onqlave arx add my_1st_arx -d 'this is my first arx' -e aes-gcm-128 -i true -o your_login_email@domain.com -p gcp -u development -r au -c monthly -t serverless
```

... and wait for the output logs.

```
Arx creation sometime takes up to 10 minutes.
```
If the creation is successful, the output will look similar to:
```
🎉 Done!  Arx created successfully.
Arx ID: {your-arx-id}
For more information, read our documentation at https://docs.onqlave.com/
```

## **Set a default Arx**

**Who can perform this operation?**

- Platform Owner

Input command:
```
onqlave arx default your_arx_id
```
Expected output:
```
Setting default arx sometime takes up to 10 minutes.
🎉 Done!  Arx set default successfully.
Arx ID: {your-arx-id}
For more information, read our documentation at https://docs.onqlave.com/
```

## **Describe an Arx**

**Who can perform this operation?**

- Platform Owner

You can retrieve all information of your Arx by using this command:
```
onqlave arx describe your_arx_id
```

```
┌─────────────────────────────────────────┐
│ Key                 Value               │
│─────────────────────────────────────────│
│ Name                arx-1               │
│ SpendLimit          0                   │
│ Description                             │
│ Purpose             development         │
│ PlanID              serverless          │
│ ProviderID          gcp                 │
│ EncryptionMethodID  aes-gcm-128         │
│ RotationCycleID     3-monthly           │
│ Owner               owner_id            │
│ IsDefault           false               │
└─────────────────────────────────────────┘

```

## **View a list of Arx**

**Who can perform this operation?**

- Platform Owner

To see a list of your Arx from CLI, simply use this command:


```
onqlave arx list
```
The output will be formatted in tabular format
![](https://t36712295.p.clickup-attachments.com/t36712295/909f1eed-e921-4371-860b-a85e0d1293b1/image.png)


<!-- There is another JSON output if you append the flag **--json** to the end of the above comand
```
List Arx =>
{
    "arx": [
        {
            "acl": {
                "can": {
                    "delete": true,
                    "edit": true,
                    "seal": true,
                    "set_as_default": true,
                    "unseal": false
                },
                "can_not": {
                    "unseal_reason": "You don't have access to unseal the arx as you are not the owner | only sealed arx can be unsealed!"
                }
            },
            "availability_message": "",
            "description": "",
            "encryption_method": "aes-gcm-128",
            "id": "arx--id_string_here",
            "is_default": false,
            "name": "my_1st_arx",
            "owner": "your_email@here.com",
            "plan": "serverless",
            "provider": "gcp",
            "purpose": "development",
            "regions": [
                "au"
            ],
            "rotation_cycle": "monthly",
            "spend_limit": 0,
            "status": "active"
        }
    ]
}
```
-->

## **Update an arx**

**Who can perform this operation?**

- Platform Owner

Firstly, to get help on which fields can be updated in an Arx, use this command:

```
onqlave arx update -h
```

and see the help output:

```
This command is used to update arx by ID. Arx id, arx name, arx region, arx encryption method, arx rotation cycle, arx owner, arx spend limit and arx is default are required.

Usage:
  onqlave arx update [flags]

Examples:
onqlave arx update

Flags:
  -i, --arx_is_default              Enter Arx Is Default
  -n, --arx_name string             Enter Arx Name (default "test")
  -o, --arx_owner string            Enter Arx Owner (default "Default")
  -r, --arx_region string           Enter Arx Region - (AUS-EAST, AUS-WEST)
  -c, --arx_rotation_cycle string   Enter Arx Rotation Cycle (default "Default")
  -l, --arx_spend_limit uint        Enter Arx Spend Limit
  -h, --help                        help for update

Global Flags:
      --json   Output logs as JSON.  Set to true if stdout is not a TTY
```

Once you decide which field to update, use this command and append your list of flags and values:

```
onqlave arx update arx_id <your_list_of_flags_and_values>
```


Output should look like this:
```
🎉 Done! Arx updated successfully.
Arx ID: your_arx_id
For more information, read our documentation at https://docs.onqlave.com
```

## **Seal an Arx**

**Who can perform this operation?**

- Platform Owner

This is a feature supporting you to temporarily disable an Arx and enable it in the future without having to reconfigure everything from scratch. To seal an Arx, include its ID in this command:

```
onqlave arx seal arx_id_here
```

And see the logged result:
```
Arx seal sometime takes up to 10 minutes.
🎉 Done!  Arx sealed successfully.
For more information, read our documentation at https://docs.onqlave.com/

```


## **Unseal an Arx**

**Who can perform this operation?**

- Platform Owner

In contrary to seal, we just need to alter the command:
```
onqlave arx unseal arx_id_here
```

And see the result:
```
Arx unseal sometime takes up to 10 minutes.
🎉 Done!  Arx unsealed successfully.
For more information, read our documentation at https://docs.onqlave.com/
```

## **Delete an Arx**

**Who can perform this operation?**

- Platform Owner

```
onqlave arx delete your_arx_id
```
```
🎉 Done!  Arx deleted successfully.
For more information, read our documentation at https://docs.onqlave.com/
```


## **Get base configuration information for Arx**

This information may be useful when you need to input the required flags during the creation of an Arx. To get this information, use the following command:
```
onqlave arx base
```
The result will be organized in a tabular format.

<!-- If you want plain JSON, just append **--json** to the end of the command. -->

![arx-base-command](https://t36712295.p.clickup-attachments.com/t36712295/70e44a5e-ae74-4820-9bf9-0ae125697a82/image.png)

<!-- JSON Output:
```
Arx Base Information =>
{
    "encryption_methods": [
        {
            "description": "Relatively strong \u0026 faster algorithm",
            "enable": true,
            "icon": "LockIcon",
            "id": "aes-gcm-128",
            "is_default": false,
            "name": "AES-GCM-128",
            "order": 0
        },
        {
            "description": "Strong \u0026 fast algorithm",
            "enable": true,
            "icon": "LockIcon",
            "id": "aes-gcm-256",
            "is_default": true,
            "name": "AES-GCM-256",
            "order": 1
        },
        {
            "description": "Strong \u0026 high performing algorithm",
            "enable": true,
            "icon": "LockIcon",
            "id": "cha-cha-20-poly-1305",
            "is_default": false,
            "name": "CHACHA20-POLY1305",
            "order": 2
        }
    ],
    "plans": [
        {
            "description": "Highly available arx that scale instantly",
            "enable": true,
            "icon": "ServerIcon",
            "id": "serverless",
            "is_default": true,
            "name": "Serverless",
            "order": 0
        },
        {
            "description": "Didicated single tenant arx. For more information, contact sales@onqlave.com",
            "enable": false,
            "icon": "BoxIcon",
            "id": "dedicated",
            "is_default": false,
            "name": "Dedicated",
            "order": 1
        }
    ],
    "providers": [
        {
            "description": "",
            "enable": true,
            "id": "gcp",
            "image": "image-gcp.svg",
            "is_default": true,
            "name": "Google Cloud",
            "order": 0,
            "regions": [
                {
                    "enable": true,
                    "icon": "icon-australia.svg",
                    "id": "au",
                    "is_default": false,
                    "name": "Australia",
                    "optimisation": {
                        "message": "Highly Available",
                        "value": 100
                    },
                    "order": 0
                },
                {
                    "enable": false,
                    "icon": "icon-singapore.svg",
                    "id": "sg",
                    "is_default": false,
                    "name": "Singapore",
                    "optimisation": {
                        "message": "Highly Available",
                        "value": 100
                    },
                    "order": 0
                },
                {
                    "enable": false,
                    "icon": "icon-gb.svg",
                    "id": "gb",
                    "is_default": false,
                    "name": "United Kingdom",
                    "optimisation": {
                        "message": "Highly Available",
                        "value": 100
                    },
                    "order": 0
                },
                {
                    "enable": false,
                    "icon": "icon-usa.svg",
                    "id": "us",
                    "is_default": false,
                    "name": "USA",
                    "optimisation": {
                        "message": "Highly Available",
                        "value": 100
                    },
                    "order": 0
                }
            ]
        },
        {
            "description": "",
            "enable": false,
            "id": "aws",
            "image": "image-aws.svg",
            "is_default": false,
            "name": "AWS",
            "order": 1,
            "regions": [
                {
                    "enable": true,
                    "icon": "icon-australia.svg",
                    "id": "au",
                    "is_default": false,
                    "name": "Australia",
                    "optimisation": {
                        "message": "Highly Available",
                        "value": 100
                    },
                    "order": 0
                },
                {
                    "enable": false,
                    "icon": "icon-singapore.svg",
                    "id": "sg",
                    "is_default": false,
                    "name": "Singapore",
                    "optimisation": {
                        "message": "Highly Available",
                        "value": 100
                    },
                    "order": 0
                },
                {
                    "enable": false,
                    "icon": "icon-gb.svg",
                    "id": "gb",
                    "is_default": false,
                    "name": "United Kingdom",
                    "optimisation": {
                        "message": "Highly Available",
                        "value": 100
                    },
                    "order": 0
                },
                {
                    "enable": false,
                    "icon": "icon-usa.svg",
                    "id": "us",
                    "is_default": false,
                    "name": "USA",
                    "optimisation": {
                        "message": "Highly Available",
                        "value": 100
                    },
                    "order": 0
                }
            ]
        },
        {
            "description": "",
            "enable": false,
            "id": "azure",
            "image": "image-azure.svg",
            "is_default": false,
            "name": "Azure",
            "order": 2,
            "regions": [
                {
                    "enable": true,
                    "icon": "icon-australia.svg",
                    "id": "au",
                    "is_default": false,
                    "name": "Australia",
                    "optimisation": {
                        "message": "Highly Available",
                        "value": 100
                    },
                    "order": 0
                },
                {
                    "enable": false,
                    "icon": "icon-singapore.svg",
                    "id": "sg",
                    "is_default": false,
                    "name": "Singapore",
                    "optimisation": {
                        "message": "Highly Available",
                        "value": 100
                    },
                    "order": 0
                },
                {
                    "enable": false,
                    "icon": "icon-gb.svg",
                    "id": "gb",
                    "is_default": false,
                    "name": "United Kingdom",
                    "optimisation": {
                        "message": "Highly Available",
                        "value": 100
                    },
                    "order": 0
                },
                {
                    "enable": false,
                    "icon": "icon-usa.svg",
                    "id": "us",
                    "is_default": false,
                    "name": "USA",
                    "optimisation": {
                        "message": "Highly Available",
                        "value": 100
                    },
                    "order": 0
                }
            ]
        }
    ],
    "purposes": [
        {
            "enable": true,
            "id": "development",
            "is_default": true,
            "name": "Development",
            "order": 0
        },
        {
            "enable": true,
            "id": "testing",
            "is_default": false,
            "name": "Testing",
            "order": 1
        },
        {
            "enable": true,
            "id": "production",
            "is_default": false,
            "name": "Production",
            "order": 2
        },
        {
            "enable": true,
            "id": "staging",
            "is_default": false,
            "name": "Staging",
            "order": 3
        }
    ],
    "rotation_cycles": [
        {
            "enable": true,
            "id": "monthly",
            "is_default": false,
            "name": "Monthly",
            "order": 0
        },
        {
            "enable": true,
            "id": "3-monthly",
            "is_default": true,
            "name": "3 Monthly",
            "order": 1
        },
        {
            "enable": true,
            "id": "6-monthly",
            "is_default": false,
            "name": "6 Monthly",
            "order": 2
        },
        {
            "enable": true,
            "id": "annually",
            "is_default": false,
            "name": "Annually",
            "order": 3
        }
    ]
}
```
-->

## **Explore the supported interaction commands with Arx**

```
onqlave arx
This command is used to manage arx resources.

Usage:
  onqlave arx [command]

Examples:
onqlave arx

Available Commands:
  add         add arx by name and attributes
  base        get base arx info
  default     set default arx by ID
  delete      delete arx by ID
  describe    describe arx by ID
  list        list arx
  retry       retry adding arx by ID and attributes
  seal        seal arx by ID
  unseal      unseal arx by ID
  update      update arx by ID and attributes

Flags:
  -h, --help   help for arx

Global Flags:
      --json   Output logs as JSON.  Set to true if stdout is not a TTY.

Use "onqlave arx [command] --help" for more information about a command.
```

# Application

## **Before you start**

The CLI commands for your Application will allow you to create and allocate the unique identifiers for your front and backend applications to be used when creating API Keys. This separated Application workflow ensures you have easy access to enabling, disabling and archiving applications as needed.


When the application reference is created, an API token and encryption key is established. Note that Onqlave does not allow you to permanently delete any applications, however they can be archived, which will then disable the respective API token and encryption key.

You may need to look at the supported commands for an application

## **Create an application**

**Who can perform this operation?**

- Platform Owner

```
onqlave application add your_app_name -d 'this is my application' -t application_technology -o application_owner_id
```

Then the returned output should include your created application ID
```
🎉 Done! Application created successfully.
Application ID: your-app-id-here
For more information, read our documentation at https://docs.onqlave.com
```

## **Describe an application**

**Who can perform this operation?**

- Platform Owner

```
onqlave application describe your_app_id
```

The output is formatted as a table:
 <!-- or JSON depends on your choice of appending **--json** flag -->
```
┌──────────────────────────────────────┐
│ Key          Value                   │
│──────────────────────────────────────│
│ ID           your_app_id_here        │
│ Name         app-2                   │
│ Description                          │
│ Technology   server                  │
│ Owner        owner_id                │
│ APIKeys      0                       │
│ Status       active                  │
└──────────────────────────────────────┘
```
<!--
```
{
   "Application": {
       "acl": {
           "can": {
               "archive": false,
               "disabled": true,
               "edit": true
           },
           "can_not": {
               "archive_reason": "Application is not disabled yet!"
           }
       },
       "api_keys": 0,
       "application_id": "your_app_id_here",
       "cors": [],
       "description": "",
       "name": "app-2",
       "owner": "owner_id",
       "status": "active",
       "technology": "server"
   }
}

``` 
-->

## **List all applications**

**Who can perform this operation?**

- Platform Owner

```
onqlave application list
```
The output will be displayed as a table by default.
<!-- And you can show the JSON output by appending the **--json** to the end of the above command. -->

![application-list](https://t36712295.p.clickup-attachments.com/t36712295/32dbd08e-a5ec-4758-8770-2a40c1359ab6/image.png)


## **Update an application**

**Who can perform this operation?**

- Platform Owner

Currently, the Onqlave platform supports updating an application via its ID

```
onqlave application update your_application_id your_list_of_flags_and_values
```

To see the available flags, you can use this command:
```
onqlave application update
```

And explore all the flags:

```
Usage:
 onqlave application update [flags]

Examples:
onqlave application update

Flags:
 -c, --application_cors string          Enter Application Cors
 -d, --application_description string   Enter Application Description
 -n, --application_name string          Enter Application Name
 -o, --application_owner string         Enter Application Owner
 -t, --application_technology string    Enter Application Technology
 -h, --help                             help for update

Global Flags:
     --json   Output logs as JSON.  Set to true if stdout is not a TTY.
```

## **Disable an application**

**Who can perform this operation?**

- Platform Owner

```
onqlave application disable your_app_id
```

## **Enable an application**

**Who can perform this operation?**

- Platform Owner

```
onqlave application enable your_app_id
```

## **Archive an application**

**Who can perform this operation?**

- Platform Owner

Since we do not support deleting applications, you can archive it. Before archiving an application, you have to disable it like the previous step.


```
onqlave application archive your_app_id
```

## **Get base configuration information for Application**

This information may be useful when you need to input the required flags during the creation of an application. To get these information, use the following command:


```
onqlave application base
```

Result will be organized either in tabular format by default or can be converted to JSON by appending **--json** flag at the end of the command

![application-base](https://t36712295.p.clickup-attachments.com/t36712295/6fb8663c-bccb-4362-a6a5-043668b2233b/image.png)

<!-- JSON output:
```
Application Base Information =>
{
   "technologies": [
       {
           "cors": false,
           "description": "Application which contains backend",
           "enable": false,
           "icon": "ServerIcon",
           "id": "server",
           "is_default": false,
           "name": "Server",
           "order": 0
       },
       {
           "cors": true,
           "description": "Application which contains frontend",
           "enable": false,
           "icon": "ChromeIcon",
           "id": "client",
           "is_default": false,
           "name": "Client",
           "order": 1
       }
   ]
}
```
-->

## **Explore available commands**

```
onqlave application
```
```
This command is used to manage applications resources.

Usage:
 onqlave application [command]

Examples:
onqlave application

Available Commands:
 add         add application by name and attributes
 archive     archive application by ID
 describe    describe application by ID
 disable     disable application by ID
 enable      enable application by ID
 list        list applications
 update      update application by ID and attributes

Flags:
 -h, --help   help for application

Global Flags:
     --json   Output logs as JSON.  Set to true if stdout is not a TTY.

Use "onqlave application [command] --help" for more information about a command.
```

# API Key

## **Before you start**

The CLI commands to manage API Keys allow you to bring your Arx and applications together. The API Keys created here will draw on all of the unique inputs you have created for the Arx and application that you chose.

When you create your key it is critical that you store it as **it will only be displayed once**. To preserve the integrity of the key, Onqlave does not keep a record of this. Do not close the final window  until you have made this record!

You may need to look at the supported available commands for api keys and **[base configuration information for API Keys](#get-base-configuration-information-for-api-keys)** before working with API Keys via CLI.

## **Create an API key**

**Who can perform this operation?**

- Platform Owner
- Developer

To create your api key, your have to specify your application_id, application_technology and arx_id

```
onqlave key add -a your_application_id -c your_arx_id -t your_application_technology
```

The output log will include the ID of the newly created APIKey
```
🎉 Done! API Key created successfully.
API Key ID: apikey--your_api_key
For more information, read our documentation at https://docs.onqlave.com
```

## **Describe an API key**

**Who can perform this operation?**

- Platform Owner
- Developer

```
onqlave key describe your_app_key_id
```
Output will be formatted in tabular format

![apikey-describe](https://t36712295.p.clickup-attachments.com/t36712295/a30ee2ac-3d8f-4077-be7d-13fc95e8d209/image.png)

## **List all API keys**

**Who can perform this operation?**

- Platform Owner
- Developer

```
onqlave key list
```

By default, the result will be formatted in a table.
 <!-- If you want JSON format, simply appending the **--json** at the end of the above command -->

![apikey-list](https://t36712295.p.clickup-attachments.com/t36712295/7c28851f-1fa3-4f26-a8bd-910ba668916d/image.png)

## **Delete an API key**

**Who can perform this operation?**

- Platform Owner
- Developer

```
onqlave key delete your_app_key_id
```

## **Get base configuration information for API Keys**

Before interacting with the API key, you may need to retrieve all the base information about your Arx and application. The most frequently used information when interacting with API key via CLI are IDs of Arx, application and owner


```
onqlave key base
```

The output should be similar to to the following:
 <!-- you can alternate the format into JSON by appending **--json** into the end of the command: -->

![api-key-base](https://t36712295.p.clickup-attachments.com/t36712295/ce103cc9-2d3a-4bce-9d30-2d6bc0795fdf/image.png)


## **Explore available commands**

```
onqlave key
```

```
This command is used to manage api key resources.

Usage:
 onqlave key [command]

Examples:
onqlave key

Available Commands:
 add         add api key by attributes
 base        get base
 delete      delete api key by ID
 describe    describe api key by ID
 list        list api key

Flags:
 -h, --help   help for key

Global Flags:
     --json   Output logs as JSON.  Set to true if stdout is not a TTY.

Use "onqlave key [command] --help" for more information about a command.
```



<!-- ```
API Key Base Information =>
{
   "applications": [
       {
           "application_technology": {
               "cors": false,
               "description": "",
               "enable": false,
               "icon": "ServerIcon",
               "id": "server",
               "is_default": false,
               "name": "Server",
               "order": 0
           },
           "id": "app--zovKQ3NESVrtHMoPJgr5m",
           "label": "",
           "name": "Thelma Schmidt"
       }
   ],
   "arx": [
       {
           "created_by": {
               "avatar": "",
               "email_address": "your_email@gmail.com",
               "id": "Qtc2e7LVjSO7Q1tJm0PwaoeVFUS2",
               "name": "Platform"
           },
           "encryption": {
               "icon": "LockIcon",
               "id": "aes-gcm-256",
               "name": "AES-GCM-256"
           },
           "id": "arx---CSjvl4DuOygw8sYXJntm",
           "label": "",
           "name": "Marquise Douglas",
           "plan": {
               "icon": "ServerIcon",
               "name": "Serverless"
           },
           "provider": {
               "image": "image-gcp.svg",
               "name": "Google Cloud"
           },
           "purpose": {
               "name": "Development"
           },
           "regions": [
               {
                   "icon": "icon-australia.svg",
                   "name": "Australia"
               }
           ],
           "rotation_cycle": {
               "id": "3-monthly",
               "name": "3 Monthly"
           }
       },
       {
           "created_by": {
               "avatar": "",
               "email_address": "your_email@gmail.com",
               "id": "Qtc2e7LVjSO7Q1tJm0PwaoeVFUS2",
               "name": "Platform"
           },
           "encryption": {
               "icon": "LockIcon",
               "id": "aes-gcm-256",
               "name": "AES-GCM-256"
           },
           "id": "arx--YcvryC1LtCGmTLXQUJJyP",
           "label": "",
           "name": "Horace Gibson",
           "plan": {
               "icon": "ServerIcon",
               "name": "Serverless"
           },
           "provider": {
               "image": "image-gcp.svg",
               "name": "Google Cloud"
           },
           "purpose": {
               "name": "Development"
           },
           "regions": [
               {
                   "icon": "icon-australia.svg",
                   "name": "Australia"
               }
           ],
           "rotation_cycle": {
               "id": "monthly",
               "name": "Monthly"
           }
       },
       {
           "created_by": {
               "avatar": "",
               "email_address": "your_email@gmail.com",
               "id": "Qtc2e7LVjSO7Q1tJm0PwaoeVFUS2",
               "name": "Platform"
           },
           "encryption": {
               "icon": "LockIcon",
               "id": "aes-gcm-256",
               "name": "AES-GCM-256"
           },
           "id": "arx--v8eE3s1EnT8bYpjU2C6Wj",
           "label": "",
           "name": "Sister Jerde",
           "plan": {
               "icon": "ServerIcon",
               "name": "Serverless"
           },
           "provider": {
               "image": "image-gcp.svg",
               "name": "Google Cloud"
           },
           "purpose": {
               "name": "Development"
           },
           "regions": [
               {
                   "icon": "icon-australia.svg",
                   "name": "Australia"
               }
           ],
           "rotation_cycle": {
               "id": "monthly",
               "name": "Monthly"
           }
       }
   ]
}
```
-->

# Platform

# Account

## **Before you start**

You can refer to this section on the dashboard for more details.

Or you can refer to this section for getting help about available commands.

## **Retrieving tenant information**

You can use this command to retrieve the organisation/tenant information:

```
# onqlave tenant describe
```

The result should be displayed in tabular format by default and look like this:

```
┌────────────────────────────────────────────┐
│ Key          Value                         │
│────────────────────────────────────────────│
│ Id           tenant--cr7ZHkeyjdWycfsPF     │
│ Name         dc-tenant-1                   │
│ Label        theirs                        │
│ OwnerEmail   your_email@host.com           │
└────────────────────────────────────────────┘
```
<!-- You can also append **--json** to the end of the command to get a JSON output like this

```
Tenant 'tenant--cr7ZHkeyjdWycfsPF' Information =>
{
   "data": {
       "acl": {
           "can": {
               "edit": true
           },
           "can_not": null
       },
       "created_on": "2023-04-06T07:16:44.169507Z",
       "owner_email": "your_email@host.com",
       "tenant_id": "tenant--cr7ZHkeyjdWycfsPF",
       "tenant_label": "dc",
       "tenant_name": "dc-tenant-1"
   },
   "error": {
       "code": 0,
       "correlation_id": "",
       "details": null,
       "message": "",
       "status": ""
   }
}
```
-->

## **Updating tenant information**
There are two fields that you can update: **label** and **name**

```
# onqlave tenants update -l dcm -n dc-tenant-1-updated
```

The successful output appears as follows:

```
🎉 Done!  Tenant updated successfully
```

You can double check the updated result by using the **describe** command above:

```
┌────────────────────────────────────────────┐
│ Key          Value                         │
│────────────────────────────────────────────│
│ Id           tenant--B8dxGtiyx2CWG8mYtvpfr │
│ Name         dev-tenant-1-updated          │
│ Label        your_label                    │
│ OwnerEmail   po.onclave@gmail.com          │
└────────────────────────────────────────────┘
```

## **Explore available commands**

```
# onqlave tenant
```

```
This command is used to manage tenants' resources.

Usage:
 onqlave tenant [command]

Examples:
onqlave tenant

Available Commands:
 describe    describe tenant
 update      update tenant by name and label

Flags:
 -h, --help   help for tenant

Global Flags:
     --json   JSON Output. Set to true if stdout is not a TTY.

Use "onqlave tenant [command] --help" for more information about a command.
```

## Access


## **Retrieving information**

Currently, from the CLI, we only support retrieving information about participating users. To get all the information of your participating users, you can use this command:

```
# onqlave user list
```

The output should look like this:

![user-list](https://t36712295.p.clickup-attachments.com/t36712295/685e6c2b-83e6-4c35-8df9-ecc97b7f5417/access%2Baccount-2.png)

<!-- If you want JSON output, simply append **--json** flag to the end of the above command
```
List Users =>
{
    "users": [
        {
            "avatar": "",
            "country_code": "au",
            "createdAt": "2023-04-06T07:16:43.295363Z",
            "disable": false,
            "email_address": "your_email@host.com",
            "full_name": "John Doe",
            "id": "user_id",
            "role": [
                "platform_owner"
            ],
            "status": "active",
            "tenant_id": "your_tenant_id",
            "updatedAt": "2023-04-06T07:17:55.579855Z"
        }
    ]
}
=====
``` -->

To add or invite new users, you need to do so in the dashboard, as explained here
