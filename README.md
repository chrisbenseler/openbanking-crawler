# Open Banking Crawler

## Running in docker

```ssh
docker build -t openbankingcrawler .
docker run -p 3000:3000 -e DBHOST=host.docker.internal -t openbankingcrawler:latest
```
