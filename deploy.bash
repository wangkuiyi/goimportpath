GOOS=linux GOARCH=amd64 go install github.com/wangkuiyi/goimportpath &&
    ssh root@topic.ai 'killall goimportpath' ;
    scp $GOPATH/bin/linux_amd64/goimportpath root@topic.ai:/root/ &&
    scp ./server.* root@topic.ai:/root/ &&
    ssh root@topic.ai 'nohup /root/goimportpath'
