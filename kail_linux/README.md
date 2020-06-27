```bash
docker build . -t kail_linux:v1
docker run -it kail_linux bin/bash
```

bin/bashにて
```bash
service postgresql start && msfdb init
```