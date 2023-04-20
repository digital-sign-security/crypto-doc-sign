# crypto-doc-sign

- POST /docs
- GET /docs/<doc-id>
- GET /docs/available

- GET /public_key/my (without my???)
- GET /public_key/<user-id>
- POST /public_key/my

скорее всего правильнее будет отправлять в двух режимах, 
1 - доверительном через сервер (по открытому ключу сервака)
2 - отложенному (по открытому ключу другого юзера, но придется его дозапрашивать)