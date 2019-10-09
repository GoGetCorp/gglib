# GoGet Task library

This Go v1.13 package is to be used in preforming the Go programming task

To download this package so it can be used in your code:

`go get -u github.com/GoGetCorp/gglib`

## Import

You need to import the library into your code before you can use it. (see above)
Importing 3rd party libs is an exercise as part of the candidates task.

## Usage

```
errChannel := make(chan error)
GenerateAnnoyingErrors(errChannel)
```

OBS!: For the task the call to `GenerateAnnoyingErrors` must be called in a concurrent fashion :)
