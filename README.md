English | [简体中文](https://github.com/betterfor/action-send-mail/blob/main/README_ZH.md)

# Send Mail Github Action

An action that simply sends a mail to multiple recipients.

Some feature:

- Plain text body
- HTML body
- Multipart body (plain text + HTML)
- Markdown to HTML
- File attachments

## Usage

```yaml
- name: Send mail
  uses: betterfor/action-sned-mail@main
  with:
    # Reqired mail server address 
    server_address: smtp.qq.com
    # Optional Server port, default 25 (if server_port is 465 this connection use TLS)
    server_port: 465
    # Optional (recommended): mail server username
    username: ${{secrets.MAIL_USERNAME}}
    # Optional (recommended): mail server password
    password: ${{secrets.MAIL_PASSWORD}}
    # Required mail subject
    subject: Github Actions job results
    # Required recipients address
    to: alice@example.com, bob@example.com
    # Required sender full name (address can be skipped)
    from: alice # <alice@example.com>
    # Optional plain body
    body: Build job of ${{github.repository}} completed successfully!
    # Optional HTML body read from file
    html_body: file://README.md
    # Optional carbon copy recipients
    cc: a@example.com,b@example.com
    # Optional blind carbon copy recipients
    bcc: c@example.com,d@example.com
    # Optional recipient of the email response
    reply_to: luke@example.com
    # Optional Message ID this message is replying to
    in_reply_to: <random-luke@example.com>
    # Optional converting Markdown to HTML (set content_type to text/html too)
    convert_mardown: true
    # Optional attachments
    attachments: README.md
    # Optional priority: 'high', 'normal' (default) or 'low'
    priority: low
```

## Troubleshooting

### Unauthenticated login (username/password field)

The parameters `username` and `password` are set as optional to support self-hosted runners access to on-premise infrastructrue. If you are accessing public email servers make sure you provide a username/password authentication through [Github Secrets](https://docs.github.com/en/actions/security-guides/encrypted-secrets) to make the email delivery secure.