# ZCHAT2 Protocol

## User ID

```plaintext
user.name@example.com
```

## Discovery

The zchat2 discovery protocol allows several different methods to discover a zchat2 server.

The priority of the discovery methods is as follows:

1. DNS TXT
2. .well-known Endpoint

### DNS TXT

If the user ID is `user.name@example.com` then the DNS TXT record should be:

```plaintext
;; zchat2 discovery records

;; endpoint
;; multiple endpoints are allowed
;; 
zc2f719f IN TXT "v=zchat2ep1; url=https://ap-northeast-2.example.org:443/zchat2;"
zc2f719f IN TXT "v=zchat2ep1; url=https://eu-north-1.example.org:443/zchat2;"

;; server fingerprints
;; multiple fingerprints are allowed
;;
zc2a377d IN TXT "v=zchat2fp1; alg=1; fp=c203384b983e18398d3e0785fdfb716b1475a0b8235ad815c4b76b7e4ff9053a;"
zc2a377d IN TXT "v=zchat2fp1; alg=1; fp=f801225583b5367374b9bc62f7fae5262c822fb388cfd0a05b28cfc14b2bdccd;"
zc2a377d IN TXT "v=zchat2fp1; alg=1; fp=a0ff18fcf63764bf61ab3da823d5c100e41ab40128a447099e5f75f300cf94eb;"
zc2a377d IN TXT "v=zchat2fp1; alg=1; fp=019ddeebbd9c2e82d97d4f1f6f14509af31136881d9cd8a7bde35c40eaf1e815;"
```

## Session Protocol

