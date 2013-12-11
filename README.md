# mmfa

mmfa is a command line tool to help manage MFA secrets and generate one time MFA codes on OSX.

It is heavily inspired by [Bradly Feeley's ruby mmfa project](https://github.com/bradly/mmfa).

## Quick Start

First, download a pre-built binary from the desired [release](https://github.com/thbishop/mmfa-go/releases).

### Add a Secret

You can add a MFA secret with the add command (it's stored in the OSX login keychain):

```
$ ./mmfa add -n my_service$
```

### Get the Current MFA Code

You can then grab the current MFA code with:

```
$ ./mmfa get -n my_service
```

## Contribute
* Fork the project
* Make your feature addition or bug fix (with tests and docs) in a topic branch
* Send a pull request and I'll get it integrated

## License
See LICENSE
