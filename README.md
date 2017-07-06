# WorkLog

I have a terrible time remembering what I was working on or what I have done recently.
This is a small pet project to track things I do. If you find it useful, awesome.

## Framework

For the CLI framework I chose to use [Cobra](https://github.com/spf13/cobra). It is super simple to use and easy to  expand. A lot of interesting CLI tools have been developed using this framework. Kubectl, hugo, docker, and others have used Cobra to provide the framework for their CLI tools.



## Testing

TBD



## Config Options

Currently there are only three config options, the config file is loaded from 

```
$HOME/.worklog.yaml
```



LogDir: Where to store your current work log

CurrentLog: The name of the file that is your current log

ArchiveLogDir: Where you want to store archived logs (created with new command)
