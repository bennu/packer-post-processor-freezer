# Deprecated 

This postprocessor was merged to  [hashicorp/packer](https://github.com/hashicorp/packer).
 
# Freezer is a Packer Postprocessor for VMWare vSphere

It allows you to create a template for the [vmware-iso builder](https://www.packer.io/docs/builders/vmware-iso.html).

Go ahead and download a pre-built Freezer binary for your operating system from [Releases page](https://github.com/bennu/packer-post-processor-freezer/releases)
Then you need choose one folder for the magic to begins:

1. The directory where packer is, or the executable directory.

2. ~/.packer.d/plugins on Unix systems or %APPDATA%/packer.d/plugins on Windows.

3. The current working directory.

For more details check the [install plugins page](https://www.packer.io/docs/extending/plugins.html#installing-plugins)

## Basic Example

```json
{  
   "type":"freezer",
   "host":"{{user `host`}}",
   "insecure": true,
   "username":"{{user `username`}}",
   "password":"{{user `password`}}",
   "vm_name":"{{user `vm_name`}}"
}
```

* `insecure` - If it's true Skip verification of server certificate. Default is false.   

## Configuration Reference

Optional parameters:
* `folder` - Where the template is created.
* `datacenter` - If you have more than one, you will need to specified one.

## Full Example

```json
{  
   "type":"freezer",
   "host":"vcenter.local",
   "insecure": true,
   "username":"root",
   "password":"sssh_is_a_secret",
   "datacenter":"murlock",
   "vm_name":"centos-7.3",
   "folder":"/packer-templates/os/centos-7"
}
```

