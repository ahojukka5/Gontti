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
cat auth.json | base64
```

XRA is read from environment variable DOCKER_XRA during deployment.

## Some curl commands to understand /var/run/docker.sock

List containers:

```bash
jukka@jukka-XPS-13-9380:~$ curl -s --unix-socket /var/run/docker.sock http://localhost/v1.40/containers/json | jq
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
