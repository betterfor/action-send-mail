[![Build Status](https://github.com/betterfor/action-send-mail/workflows/Test%20action/badge.svg)](https://github.com/betterfor/action-send-mail/commits/main)

[English](https://github.com/betterfor/action-send-mail/blob/main/README.md) | 简体中文

# 发送邮件 Github Action

> 参考 [action-send-mail](https://github.com/dawidd6/action-send-mail) 

通过 action 可以发送多个收件人。

一些特性：
- 文本内容
- HTML内容
- 支持文本+HTML
- Markdown转HTML
- 附加文件

## 用法

```yaml
- name: Send mail
  uses: betterfor/action-sned-mail@main
  with:
    # 必需，邮件服务器地址
    server_address: smtp.qq.com
    # 必需，邮件服务器端口，默认25 (如果端口为465，则会使用TLS连接)
    server_port: 465
    # 可选 (建议): 邮件服务器用户
    username: ${{secrets.MAIL_USERNAME}}
    # 可选 (建议): 邮件服务器密码
    password: ${{secrets.MAIL_PASSWORD}}
    # 必需，邮件主题
    subject: Github Actions job results
    # 必需，收件人地址
    to: alice@example.com, bob@example.com
    # 必需，发送人全名 (地址可以省略)
    from: alice # <alice@example.com>
    # 可选，文本内容
    body: Build job of ${{github.repository}} completed successfully!
    # 可选，HTML内容，可从文件读取
    html_body: file://README.md
    # 可选，抄送人
    cc: a@example.com,b@example.com
    # 可选，密抄送人
    bcc: c@example.com,d@example.com
    # 可选，邮件回执
    reply_to: luke@example.com
    # 可选，回执邮件消息ID
    in_reply_to: <random-luke@example.com>
    # 可选，markdown转HTML (会设置内容格式为text/html)
    convert_mardown: true
    # 可选，附件
    attachments: README.md
    # 可选，邮件优先级设置: 'high', 'normal' (default) or 'low'
    priority: low
```

## 问题定位

### Unauthenticated login 

参数`username`和`password`被设置为可选参数，以支持自我托管的运行者访问内部基础设施。如果你要访问公共电子邮件服务器，请确保你通过[Github Secrets](https://docs.github.com/en/actions/security-guides/encrypted-secrets)提供用户名/密码认证，以使电子邮件的发送安全。