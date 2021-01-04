# Gontti

Gontti is container deployment util written with Go language.

## How to use

```bash
gontti ahojukka5/anecdotes
```

After a while, you should have `ahojukka5/anecdotes` published in Docker hub.

During pushing to container registry, X-Registry-Auth must be set in request
header. Basically it can be generated in the following way:

```json
{
  "username": "myusername",
  "password": "mypassword",
}
```

```bash
cat auth.json | base64 -w 0
```

Result is base64-encoded string containing credentials. Keep in mind that it's
unsafe to share it. XRA is read from environment variable `DOCKER_XRA` during
deployment.

This tool can also be used with docker. Just map volume `/var/run/docker.sock`
and add environment variable `DOCKER_XRA`.

```bash
alias gontti='docker run --rm -it -v /var/run/docker.sock:/var/run/docker.sock -e DOCKER_XRA=xxyyzz ahojukka5/gontti'
gontti ahojukka5/gontti # build and upload ahojukka5/gontti to Docker hub
```

This is actually how docker image for this repository is created. GitHub action is:

```yml
name: Build and push to Docker registry

on: push

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - name: Build and push to Docker registry
        run: docker run --rm -v /var/run/docker.sock:/var/run/docker.sock -e DOCKER_XRA=${{ secrets.DOCKER_XRA }} ahojukka5/gontti ahojukka5/gontti
 ```

If you want to utilize this script in your own workflow, replace the latter
`ahojukka5/gontti` with your own repository.

## Some curl commands to understand /var/run/docker.sock

List containers:

```bash
curl -s --unix-socket /var/run/docker.sock http://localhost/v1.40/containers/json | jq
```

Download source from Internet build it:

```bash
curl -s --unix-socket /var/run/docker.sock -X POST "http://localhost/v1.40/build?dockerfile=anecdotes-master/Dockerfile&remote=https://github.com/ahojukka5/anecdotes/archive/master.tar.gz"
```

```bash
curl -s --unix-socket /var/run/docker.sock -X POST "http://localhost/v1.40/build?remote=https://github.com/ahojukka5/anecdotes.git#master:/"
```

Tag image `9e7244cf5587`:

```bash
curl -s --unix-socket /var/run/docker.sock -X POST "http://localhost/v1.40/images/9e7244cf5587/tag?repo=ahojukka5/anecdotes"
```

Push image `9e7244cf5587`:

```bash
curl -s --unix-socket /var/run/docker.sock -X POST "http://localhost/v1.40/images/9e7244cf5587/push?tag=ahojukka5/anecdotes" -H "X-Registry-Auth" -d "$XRA"
```
