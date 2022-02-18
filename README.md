# Doka
This tool adds new features to Dockerfiles. It does not reimplement parts of the Docker API nor it replaces Dockerfile functionalities, it's just a preprocessor that generates a Dockerfile from various instructions.

# Usage
When running the `doka` command, it will try to get the `Dokafile` file in the current working directory. If a `Dokafile` doesn't exist, you can specify a custom path with `--file`/`-f` option or create one in the current working directory.

## Variables
Variables can be specified using `--var`/`-v` option.
```shell
doka -v my_var1=value1 -v my_var2=value2
```

You can also let doka fetch the value of a variable from the environment by not manually giving a value.
```shell
doka -v MY_ENV_VAR
```

Variables are called using `$my_var` syntax in Dokafiles.

By default, `os` and `arch` variables are set depending on the current system.

## `!IF`, `!ELIF` and `!ELSE` preprocessors
`!IF` and `!ELIF` preprocessors can be used to conditionally include or not instructions. Right now you can use it to compare variables and values between them.

**Compare to string:**
```dockerfile
!IF $my_var == "value1"
    RUN echo "Hi there"
!END
```
**Compare to another variable:**
```dockerfile
!IF $var1 != $var2
    RUN echo "var1 is not equal to var2"
!END
```
**`!ELIF` and `!ELSE`:**
```dockerfile
!IF $my_var == "value1"
    RUN echo "1"
!ELIF $my_var == "value2"
    RUN echo "2"
!ELSE
    RUN echo "3"
!END
```

You **can** nest `!IF` instructions if wanted.

## `!INCLUDE` preprocessor
You can include other Dokafiles and Dockerfiles using `!INCLUDE` preprocessor.
```dockerfile
# Dokafile.a
FROM ubuntu:18.04

!INCLUDE "./Dokafile.deps"

COPY myapp .
CMD ["./myapp"]
```
```dockerfile
# Dokafile.deps
RUN apt-get update && apt-get install -y curl
```

You **can** nest `!INCLUDE` instructions in included files.
`!INCLUDE` can be used inside `!IF` blocks.

# Examples
**Optionally installing `curl`**
```dockerfile
# Dokafile
FROM ubuntu:18.04

WORKDIR /home/ubuntu

!IF $enable_curl == "1"
    RUN apt-get update && apt-get install -y curl
!END

COPY myapp .
CMD ["./myapp"]
```
