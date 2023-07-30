# mycaserver

## 一些有用的命令
```bash  
openssl rsa -in .\rootCA\ca.private.key -text  
openssl pkey -in .\rootCA\ca.private.key -text  
```  
第一条命令查看生成的rsa 私钥内容  
第二条命令从私钥中提取rsa 公钥并打印

```bash
openssl x509 -in rootCA\root.crt -noout -text
```  
查看x509证书内容
```bash
openssl verify -CAfile .\rootCA\root.crt .\clientCA\2022-07-13_15-19-17.crt
```
验证后面的那个crt是否可信
