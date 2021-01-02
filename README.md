# Gontti

Gontti is container deployment util written with Go language.

## Commands

**gh-publish**: download a repository from github, builds a Dockerfile located
in the root and then publishes it into Docker hub. Usage:

```bash
gontti gh-publish https://github.com/ahojukka5/Gontti
```

After a while, you should have `ahojukka5/Gontti` published in Docker hub.
