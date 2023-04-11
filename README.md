## 安裝golang
### linux
```sh
apt-get install -y curl && \
    curl -O https://storage.googleapis.com/golang/go1.16.3.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.16.3.linux-amd64.tar.gz && \
    echo "export PATH=/usr/local/go/bin:$HOME/Projects/Proj1/bin:$PATH" >> ~/.zshrc && \
    source ~/.zshrc
```
### window
https://go.dev/dl/

```sh
go build -ldflags "-s -w" -o main *.go 
```

## 執行前:


`vim env/.env`
- TARGET_FileName:選定版本請至os_version查看
- SSH_PASSWORD: ssh密碼,以root方式進入 -port:20
- VNC_PASSWORD: vnc密碼


```conf
TARGET_FileName=Dockerfile-ubuntu-22.04
OUTPUT_FileName=Dockerfile

SSH_PASSWORD=16313302
VNC_PASSWORD=16313302
```

## shell_conf/rund
環境重啟過後需要執行的腳本放在這裡,將會自動執行腳本

## 執行
```sh
sh build.sh
```