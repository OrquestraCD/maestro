Maestro
=======
Maestro serves as an easier way to orchestrate AWS EC2 instances
using the native SSM tooling that AWS provides. This project is
currently in heavy development.

![alt text](demo.gif)

### Install
#### Stable
* [Windows](https://maestro.rax.io/v0.1.0/maestro-windows.zip)
* [Linux](https://maestro.rax.io/v0.1.0/maestro-linux.zip)
* [OSX](https://maestro.rax.io/v0.1.0/maestro-darwin.zip)

#### Latest
The latest build off the master branch:

* [Windows](https://maestro.rax.io/latest/maestro-windows.zip)
* [Linux](https://maestro.rax.io/latest/maestro-linux.zip)
* [OSX](https://maestro.rax.io/latest/maestro-darwin.zip)

### Commands
#### create
Provides sub commands for creating SSM Documents locally

##### command
Create a new SSM runCommand document locally from a script.

```shell
maestro create command mytest.sh
```

###### Options
* `--output, -o` - Name of the resulting SSM document file. (Optional)
* `--type, -t` - Type of script. (ie: powershell, bash) (Optional but may be required if
  the script does not use a common file extension)

#### list
Provides sub commands for listing tarets, documents, etc on an account.

##### aliases
List all aliases in the Maestro configuration file..

```shell
maestro list aliases
```

##### asgs
List all available ASGs to orchestrate.

```shell
maestro list asgs
```

###### Options
* `--fields, -f` - A comma delimited list of fields to include in the output.
* `--no-header, -H` - Output will not include a table header when set

##### documents
List all available documents on an account/region.

```shell
maestro list documents
```

###### Options
* `--fields, -f` - A comma delimited list of fields to include in the output.

##### instances
List all available instances on an account/region. Only shows instances that have
checked in with SSM.

```shell
maestro list instances
```

###### Options
* `--fields, -f` - A comma delimited list of fields to include in the output.
* `--filters, -F` - Filter instances based on key values. (eg: PlatformTypes=Linux)
* `--list, -l` - Print a comma delimited list of instances. Ex: `maestro run command -i $(maestro list instances -l) apt-get update`

#### run
Provides sub commands for running SSM documents against specified targets.

##### command
Run an inline command against specified targets.

```shell
maestro run command echo hello world
```

###### Options
* `--alias, -A` - Tells maestro the command being run is an alias set in the maestro config.
* `--autoscale-group, -a` - Autoscaling Group Name to execute command on.
* `--bucket-name, -B` - Name of the S3 Bucket to use for Maestro Output. Maestro will create
  a random bucket if no name is provided.
* `--instances, -i` - Target Instance IDs for SSM document
* `--no-clean, -N` - Do not clean up temporary resources after execution
* `--platform, -P` - Specify the platform type of the instances. (Optional, maestro will attempt
  auto detection)
* `--tag-key, -K` - Target tag key to execute SSM document against
* `--tag-value, -V` - Target tag Value to execute SSM document against. (Requires --tag-key)

##### document
Run a published SSM document against specified targets.

```shell
maestro run document AWS-UpdateSSMAgent
```

###### Options
* `--autoscale-group, -a` - Autoscaling Group Name to execute command on.
* `--bucket-name, -B` - Name of the S3 Bucket to use for Maestro Output. Maestro will create
  a random bucket if no name is provided.
* `--instances, -i` - Target Instance IDs for SSM document
* `--no-clean, -N` - Do not clean up temporary resources after execution
* `--parameters, -p` - Parameters to pass to the SSM doc. (Key1=Value1 Key2=Value2)
* `--parameters-delimiter, -d` - Parameters delimiter to split on. Defaults to splitting on a " ".
  Ex with a / Param=Value1/Param2=Value2
* `--tag-key, -K` - Target tag key to execute SSM document against
* `--tag-value, -V` - Target tag Value to execute SSM document against. (Requires --tag-key)

##### script
Run a local script against specified targets.

```shell
maestro run script ./echo-hello.sh
```

###### Options
* `--autoscale-group, -a` - Autoscaling Group Name to execute command on.
* `--bucket-name, -B` - Name of the S3 Bucket to use for Maestro Output. Maestro will create
  a random bucket if no name is provided.
* `--instances, -i` - Target Instance IDs for SSM document
* `--no-clean, -N` - Do not clean up temporary resources after execution
* `--platform, -P` - Specify the platform type of the instances. (Optional, maestro will attempt
  auto detection)
* `--tag-key, -K` - Target tag key to execute SSM document against
* `--tag-value, -V` - Target tag Value to execute SSM document against. (Requires --tag-key)

For Development information see [the Contributing guide](CONTRIBUTING.md).

### Configuration File
Maestro does have a configuration file available, it currently is only used to set command
aliases.

Example:
```json
{
  "aliases": {
    "sysinfo": {
      "command": "echo '===============Processes=============' && ps aux && echo '========================Filesystem===================' && df -h",
      "platform": "Linux",
      "description": "Get process and filesystem information.",
      "type": "bash"
    }
  }
}
```

#### Aliases
Aliases alias a one liner to a human readable name.

##### Attributes
* `command` - The command or run liner that should be run by maestro. **Required**
* `platform` - Platform the command can be run on. **Not Required**
* `description` - A description of the alias, used when listing aliases with maestro. **Not Required**
* `type` - Type of script/command (ie: bash or powershell). This is not currently in use.**Not Required**
