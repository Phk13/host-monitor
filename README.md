# Host monitor
ICMP monitoring and notification tool (Gmail OAuth)


# Usage
./host-monitor -ip ip/range -mail example@gmail.com [-i interval] [-t timeout] [-debug] [-mailTest]
### Options
```
  -debug
        Debug logging
  -i int
        Interval in seconds between each check (default 60)
  -ip value
        IP range (CIDR or single IP) to monitor - e.g. 192.168.1.1 or 192.168.1.0/24 (can be repeated to specify multiple ip/ranges)
  -mail string
        Mail to notify
  -mailTest
        Do not send any mail, only try to connect and trigger notifications.
  -t int
        Timeout in milliseconds for ICMP (default 1000ms)
```

### Generate credentials for Google OAuth
1. Access [Google API Console](https://console.developers.google.com/)
2. Create a new project (or select an existing one)
3. Go to OAuth Consent Screen and add a new user using your email
4. Go to Credentials and create a new OAuth client ID:
    -   Application type: Web application
    -   Authorized Javascript Origins: `http://IP:9090`
    -   Authorized Redirect URIs: `http://IP:9090/callback-gl`
    -   **Use the IP of the system that will launch the tool. It can be localhost if the system has a graphical browser (required to login the first time)**
5. Save the credential and download its JSON as `credentials.json`
6. Put `credentials.json` in the same folder as its executable

### First launch
When launching the tool for the first time, it will output a URL that will need to be opened in a browser. **It can be opened in another computer if the credential was created with an IP instead of localhost**

The URL redirects to Google Auth. After login into the corresponding account and authorizing the app, it will create a token.json and start the app.

Afterwards, the token will be refreshed automatically.

### Token.json was removed or any other issues with token verification
If the token is removed, the authorization process will not set up automatic token refresh. To enable it again, the authorization needs to be revoked.

In case of any other errors regarding google auth, the same process applies.

1. Access [Google Account Security](https://myaccount.google.com/security)
2. Go to Third Party apps -> Manage third-party access
3. Remove access to the app
4. Repeat the First launch process 
